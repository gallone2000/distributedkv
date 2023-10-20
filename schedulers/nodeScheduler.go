package schedulers

import (
	"time"
	"distributedkv/models"
)

func GetActiveNodesScheduler(app *models.App, isStartUp bool) {
	// if isStartUp is true send a request to /nodes to the node in the .env file
	// save the list of active nodes in app.ActiveNodes.Nodes
	for {
		// send request to /nodes endpoint to few active nodes available in app.ActiveNodes.Nodes
		//compare the result and update accordingly the list of active nodes
		// if the node that sent the request is not in the list of active nodes
		// if the status is healthy send a request to /rejoin endpoint of few nodes available in app.ActiveNodes.Nodes
		time.Sleep(5 * time.Minute)
	}
}

func GetFailedNodesScheduler(app *models.App, isStartUp bool) {
	// if isStartUp is true send a request to /failed-nodes to the node in the .env file
	// save the list of failed nodes in app.FailedNodes.Nodes
	for {
		// send request to /failed-nodes endpoint to few active nodes available in app.ActiveNodes.Nodes
		//compare the result and update accordingly the list of failed nodes
		// if the node that sent the request is in the list of failed nodes
		// if the status is healthy send a request to /rejoin endpoint of few nodes available in app.ActiveNodes.Nodes
		time.Sleep(5 * time.Minute)
	}
}

func HeartBeatScheduler(app *models.App) {
	for {
		// send requests to /heartbeat endpoint of few active nodes available in app.ActiveNodes.Nodes
		// take into account possible failures or network timeouts => this package to send http requests can be a valid option resty.Client
		// if a request to /heartbeat endpoint fails move the node from App.ActiveNodes.Nodes to App.FailedNodes.Nodes
		// it can be an option to have a map where to save for each node the timestamp of the first failed request
		// to set a max number of attempts and time window.
		// if the last attempt is equal to MAX_ATTEMPTS and time.now() - last attempt < TIME_WINDOW mode the node to App.FailedNodes.Nodes
		time.Sleep(5 * time.Second)
	}
}

func RejoinScheduler(app *models.App) {
	//send a rejoin request to few nodes in app.ActiveNodes and return
}
