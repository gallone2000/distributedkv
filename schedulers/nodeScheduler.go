package schedulers

import (
	"distributedkv/models"
	"time"
)

func GetActiveNodesScheduler(app *models.App, isStartUp bool) {
	// if isStartUp is true send a request to /nodes to the node in the .env file
	// save the list of active nodes in app.ActiveNodes.Nodes with app.ActiveNodes.Put() to avoid race conditions

	for {
		// send request to /nodes endpoint to few active nodes available in app.ActiveNodes.Nodes
		// please use app.ActiveNodes.GetAllNodes() to avoid race conditions
		//compare the result and update accordingly the list of active nodes
		// if the node that sent the request is not in the list of active nodes
		// if the status is healthy send a request to /rejoin endpoint of few nodes available in app.ActiveNodes.Nodes
		// for this demo I set the sleep time to 5 minutes but, in production, it would be better to have a config variable in .env file
		// in this case this function will be called every 5 minutes
		time.Sleep(5 * time.Minute)
	}
}

func GetFailedNodesScheduler(app *models.App, isStartUp bool) {
	// if isStartUp is true send a request to /failed-nodes to the node in the .env file
	// save the list of failed nodes in app.FailedNodes.Nodes with app.FailedNodes.Put() to avoid race conditions
	for {
		// send request to /failed-nodes endpoint to few active nodes available in app.ActiveNodes.Nodes
		// please use app.ActiveNodes.GetAllNodes() to avoid race conditions
		// compare the result and update accordingly the list of failed nodes
		// if the node that sent the request is in the list of failed nodes
		// if the status is healthy send a request to /rejoin endpoint of few nodes available in app.ActiveNodes.Nodes
		// for this demo I set the sleep time to 5 minutes but, in production, it would be better to have a config variable in .env file
		// in this case this function will be called every 5 minutes
		time.Sleep(5 * time.Minute)
	}
}

func HeartBeatScheduler(app *models.App) {
	for {
		// send requests to /heartbeat endpoint of few active nodes available in app.ActiveNodes.Nodes
		// please use app.ActiveNodes.GetAllNodes() to avoid race conditions
		// take into account possible failures or network timeouts => this package to send http requests can be a valid option resty.Client
		// if a request to /heartbeat endpoint fails move the node from App.ActiveNodes.Nodes to App.FailedNodes.Nodes
		// please use app.FailedNodes.Put() to avoid race conditions
		// it can be an option to have a map where to save, for each node, the timestamp of the first failed request
		// to set a const for max number of attempts and time window.
		// if the last attempt is equal to MAX_ATTEMPTS and time.now() - last attempt < TIME_WINDOW move the node to App.FailedNodes.Nodes
		// e.g MAX_ATTEMPTS=3 TIME_WINDOW=15 minutes => if 3rd attempts fails in the past 15 minutes reset MAX_ATTEMPTS and move the node
		// reset MAX_ATTEMPTS after the last attempt allowed  
		// for this demo I set the sleep time to 5 minutes but, in production, it would be better to have a config variable in .env file
		// in this case this function will be called every 5 minutes
		time.Sleep(5 * time.Second)
	}
}
