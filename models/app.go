package models

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sagikazarmark/slog-shim"
	"github.com/spf13/viper"
)

type App struct {
	Store       *KeyValueStore
	TmpStore    *KeyValueStore
	Nodes       *NodesRepository
	ActiveNodes *NodesRepository
	FailedNodes *NodesRepository
	Viper       *viper.Viper
	Logger      *slog.Logger
	Healthy     *bool
}

func (a *App) InitApp() App {
	kvs := KeyValueStore{Store: make(map[string]string)}
	tmpKvs := KeyValueStore{Store: make(map[string]string)}
	viper := initViper()
	nodesList := viper.GetString("NODES")
	nodes := initNodeRepository(nodesList)
	activeNodes := initNodeRepository()
	failedNodes := initNodeRepository()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	healthy := true

	return App{
		Store:       &kvs,
		TmpStore:    &tmpKvs,
		Nodes:       nodes,
		ActiveNodes: activeNodes,
		FailedNodes: failedNodes,
		Viper:       viper,
		Logger:      logger,
		Healthy:     &healthy,
	}
}

func initViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	// Read the configuration file
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	return v
}

func initNodeRepository(optInput ...string) *NodesRepository {
	input := ""
	if len(optInput) > 0 {
		input = optInput[0]
	}

	nodes := make(map[string]*Node, 0)
	repository := NodesRepository{
		Nodes: nodes,
		mutex: sync.RWMutex{},
	}
	if input == "" {
		return &repository
	}
	nodesURLs := strings.Split(input, ",")
	for _, nodeURL := range nodesURLs {
		repository.Put(Node{
			URL: nodeURL,
		})
	}
	return &repository
}
