package gobgpwatcher

import (
	"fmt"
	"time"

	"github.com/ecix/alice-lg/backend/api"
)

/*
 Parse GoBGP watcher api responses
*/

func parseApiStatus(gobgp ClientResponse, config Config) (api.ApiStatus, error) {
	res := mustStringMap(gobgp["result"])
	data := mustStringMap(res["api"])

	status := api.ApiStatus{
		Version:         mustString(data["version"], "0.0.0"),
		CacheStatus:     api.CacheStatus{},
		ResultFromCache: false,
		Ttl:             time.Now().Add(time.Minute * 5),
	}

	return status, nil
}

func parseServerStatus(gobgp ClientResponse, config Config) (api.Status, error) {

	res := mustStringMap(gobgp["result"])
	data := mustStringMap(res["gobgp"])
	cfg := mustStringMap(data["config"])

	routerId := mustString(cfg["router-id"], "unknown")

	status := api.Status{
		ServerTime: time.Now(),
		Version:    "GoBGP",
		RouterId:   routerId,
		Message:    "GoBGP is up and running",
	}

	return status, nil
}

func parseSessionState(state string) string {
	if state == "active" {
		return "start"
	}

	if state == "established" {
		return "up"
	}

	return state
}

func parseNeighbour(info map[string]interface{}, config Config) api.Neighbour {

	state := mustStringMap(info["state"])

	// Make description
	asn := mustInt(state["peer-as"], 0)
	addr := mustString(state["neighbor-address"], "0.0.0.0")

	descr := fmt.Sprintf("AS%d %s", asn, addr)

	peerState := parseSessionState(mustString(state["session-state"], "down"))

	timers := mustStringMap(info["timers"])
	uptime := mustDurationMs(timers["uptime"], 0)

	rTable := mustStringMap(state["adj-table"])
	nReceived := mustInt(rTable["received"], 0)
	nAccepted := mustInt(rTable["accepted"], 0)
	nFiltered := nReceived - nAccepted

	neighbour := api.Neighbour{
		Id: addr,

		Address: addr,
		Asn:     asn,

		Description: descr,

		State: peerState,

		RoutesReceived:  nAccepted,
		RoutesFiltered:  nFiltered,
		RoutesExported:  0,
		RoutesPreferred: 0,

		Uptime: uptime,

		LastError: "",

		Details: info,
	}

	return neighbour
}

func parseNeighbours(gobgp ClientResponse, config Config) (api.Neighbours, error) {
	res := mustStringMapList(gobgp["result"])
	neighbours := make(api.Neighbours, 0)

	// Build neighbours list
	for _, data := range res {
		info := mustStringMap(data)
		neighbour := parseNeighbour(info, config)

		neighbours = append(neighbours, neighbour)
	}

	return neighbours, nil
}
