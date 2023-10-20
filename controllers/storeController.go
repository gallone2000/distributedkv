package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	httptools "distributedkv/httpTools"
	"distributedkv/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetValue(w http.ResponseWriter, r *http.Request, app *models.App) {
	params := mux.Vars(r)
	key := params["key"]
	value := app.Store.Get(key)
	message := ""
	if value == "" {
		message = fmt.Sprintf("no value for key %s", key)
	}
	message = fmt.Sprintf("value: %s for key: %s", value, key)
	httptools.SendResponse(w, http.StatusNotFound, message)
}

func SetValue(w http.ResponseWriter, r *http.Request, app *models.App) {
	var req httptools.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request parameters
	if req.Key == "" || req.Value == "" {
		message := "Key and value must be provided"
		httptools.SendResponse(w, http.StatusBadRequest, message)
		return
	}

	if req.ID != "" {
		// store by uuid the key/value pair in the temporary storage
		value := fmt.Sprintf("%s:%s", req.Key, req.Value)
		app.TmpStore.Put(req.ID, value)
		message := "Value has been successfully set to temporary location"
		httptools.SendResponse(w, http.StatusOK, message)
		return
	}

	// store by uuid the key/value pair in the temporary storage
	req.ID = uuid.New().String()
	value := fmt.Sprintf("%s:%s", req.Key, req.Value)
	app.TmpStore.Put(req.ID, value)
	// collecting votes from the other nodes
	hasConsensus := hasConsensus(req, app)

	//check for the consensus
	if !hasConsensus {
		message := "Consensus forbidden"
		httptools.SendResponse(w, http.StatusForbidden, message)
		return
	}
	//send set confirmation request to other nodes
	confirmSuccessfullySet(req, app)

	//update the distributed key/value store
	app.Store.Put(req.Key, req.Value)
	//remove the key/value pairs from temporary store
	app.TmpStore.Delete(req.ID)

	// Respond to the requester indicating successful set confirmation
	message := fmt.Sprintf("Value has been successfully set")
	httptools.SendResponse(w, http.StatusOK, message)
}

// Broadcast the request to all nodes
func hasConsensus(req httptools.Request, app *models.App) bool {
	hasConsensus := false
	//votes := 0
	//halfPlusOne := int(math.Ceil(float64(countNodes) / 2))
	// for _, node := range app.ActiveNodes.Nodes
	// check the response and if the node gave the consensus update the votes variable => votes = votes + 1
	// if/when the number of positive votes is >= halfPlusOne stop to send other requests and return true
	//
	//
	// SOMETHING LIKE THIS MAYBE A go routine with parallelism would be better
	// votes := 0
	// for _, node := range app.ActiveNodes.Nodes {
	// 	url := fmt.Sprintf("%s/set", node.URL)
	// 	err := sendRequestToNode(req, url)
	// 	if err != nil {
	// 		votes = votes + 1
	// 	}
	// }

	return hasConsensus
}

// Send the request to a specific node
func sendRequestToNode(req httptools.Request, nodeURL string) error {
	jsonData, err := json.Marshal(req)
	if err != nil {
		fmt.Println("Error encoding request:", err)
		return err
	}

	_, err = http.Post(nodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request to", nodeURL, ":", err)
		return err
	}
	return nil
}

func confirmSuccessfullySet(req httptools.Request, app *models.App) {
	//I used a for loop here but it's more efficient to send the requests in parallel
	for _, node := range app.ActiveNodes.Nodes {
		url := fmt.Sprintf("%s/confirm", node.URL)
		sendRequestToNode(req, url)
	}
}

func ConfirmSet(w http.ResponseWriter, r *http.Request, app *models.App) {
	var req httptools.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	keyValue := app.TmpStore.Get(req.ID)
	if keyValue == "" {
		message := "no request found"
		httptools.SendResponse(w, http.StatusNotFound, message)
		return
	}
	splitString := strings.Split(keyValue, ":")
	key := splitString[0]
	value := splitString[1]
	//  update the distrubuted key/value store
	app.Store.Put(key, value)
	//remove the key/value pair from the temporary store using the uuid request
	app.TmpStore.Delete(req.ID)
	message := "stope properly updated"
	httptools.SendResponse(w, http.StatusOK, message)
	return
}
