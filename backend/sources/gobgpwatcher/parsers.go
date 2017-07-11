package gobgpwatcher

import (
	"time"

	"github.com/ecix/alice-lg/backend/api"
)

/*
 Parse GoBGP watcher api responses
*/

func parseApiStatus(gobgp ClientResponse, config Config) (api.ApiStatus, error) {
	data := mustStringMap(gobgp["api"])

	status := api.ApiStatus{
		Version:         mustString(data["version"], "0.0.0"),
		CacheStatus:     api.CacheStatus{},
		ResultFromCache: false,
		Ttl:             time.Now().Add(time.Minute * 5),
	}

	return status, nil
}

func parseServerStatus(gobgp ClientResponse, config Config) (api.Status, error) {
	data := mustStringMap(gobgp["gobgp"])
	cfg := mustStringMap(data["config"])

	routerId := mustString(cfg["router-id"], "unknown")

	status := api.Status{
		ServerTime: time.Now(),
		Version:    "",
		RouterId:   routerId,
	}

	return status, nil
}
