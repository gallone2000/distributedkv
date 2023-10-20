package controllers

import (
	"encoding/json"
	"net/http"
	httptools "distributedkv/httpTools"
	"distributedkv/models"
)

func GetNodes(w http.ResponseWriter, r *http.Request, app *models.App) {
	jsonData, err := json.Marshal(app.ActiveNodes.GetAllNodes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetHealtStatus(w http.ResponseWriter, r *http.Request, app *models.App) {
	message := "healthy"
	httptools.SendResponse(w, http.StatusOK, message)
	return
}

func GetHeartBeat(w http.ResponseWriter, r *http.Request, app *models.App) {
	message := "I am online"
	httptools.SendResponse(w, http.StatusOK, message)
	return
}

func Rejoin(w http.ResponseWriter, r *http.Request, app *models.App) {
	var req httptools.SimpleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	node := models.Node{URL: req.Message}
	// check if node is in the failed nodes list
	// send to the node a request to /health endpoint
	// if the response is successful so the node is healthy
	// add the node to the list of active nodes
	// remove the node from the list of failed nodes
	app.ActiveNodes.Put(node)
	httptools.SendResponse(w, http.StatusOK, "rejoined")
}

func GetFailedNodes(w http.ResponseWriter, r *http.Request, app *models.App) {
	//returns the list of failed nodes in App.FailedNodes.Nodes
}
