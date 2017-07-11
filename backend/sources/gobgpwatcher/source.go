package gobgpwatcher

import (
	"fmt"

	"github.com/ecix/alice-lg/backend/api"
)

type Gobgpwatcher struct {
	config Config
	client *Client
}

func NewGobgpwatcher(config Config) *Gobgpwatcher {
	client := NewClient(config.Api)

	watcher := &Gobgpwatcher{
		config: config,
		client: client,
	}
	return watcher
}

func (self *Gobgpwatcher) Status() (api.StatusResponse, error) {
	gobgp, err := self.client.GetJson("/v1/status")
	if err != nil {
		return api.StatusResponse{}, err
	}

	serverStatus, err := parseServerStatus(gobgp, self.config)
	if err != nil {
		return api.StatusResponse{}, err
	}

	apiStatus, err := parseApiStatus(gobgp, self.config)
	if err != nil {
		return api.StatusResponse{}, err
	}

	response := api.StatusResponse{
		Api:    apiStatus,
		Status: serverStatus,
	}

	return response, nil
}

func (self *Gobgpwatcher) Neighbours() (api.NeighboursResponse, error) {
	gobgp, err := self.client.GetJson("/v1/neighbours")
	if err != nil {
		return api.NeighboursResponse{}, err
	}

	neighbours, err := parseNeighbours(gobgp, self.config)
	response := api.NeighboursResponse{
		Neighbours: neighbours,
	}
	return response, err
}

func (self *Gobgpwatcher) Routes(neighbourId string) (api.RoutesResponse, error) {
	return api.RoutesResponse{}, nil
}

func (self *Gobgpwatcher) AllRoutes() (api.RoutesResponse, error) {
	return api.RoutesResponse{}, fmt.Errorf("Routes dumping not implemented")
}
