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

func bgpParseNextHop(attrs []map[string]interface{}) string {
	for _, attr := range attrs {
		if attr["nexthop"] != nil {
			return attr["nexthop"].(string)
		}
	}
	return ""
}

func bgpParseAsPath(attrs []map[string]interface{}) []int {
	for _, attr := range attrs {
		if attr["as_paths"] != nil {
			// build aspath.
			path := attr["as_paths"].([]interface{})[0].(map[string]interface{})
			asns := mustIntList(path["asns"])
			return asns
		}
	}

	return []int{}
}

func parseBgpInfo(
	info map[string]interface{},
	attrs []map[string]interface{},
	config Config,
) api.BgpInfo {

	bgpInfo := api.BgpInfo{
		Origin: mustString(info["source-id"], "unknown source"),

		AsPath:  bgpParseAsPath(attrs),
		NextHop: bgpParseNextHop(attrs),

		// Communities: bgpParseCommunities(attrs),
		// LargeCommunities: bgpParseLargeCommunities(attrs)
	}

	return bgpInfo
}

func parseRoute(prefix string, info map[string]interface{}, config Config) api.Route {

	attrs := mustStringMapList(info["attrs"])
	bgpInfo := parseBgpInfo(info, attrs, config)

	route := api.Route{
		Id: prefix,

		NeighbourId: mustString(info["neighbor-ip"], ""),
		Network:     prefix,

		Bgp: bgpInfo,
		Age: mustDurationMs(info["age"], 0),

		Details: info,
	}
	return route
}

func parseRoutesImported(gobgp ClientResponse, config Config) ([]api.Route, error) {
	routes := []api.Route{}
	res := mustStringMap(gobgp["result"])

	for prefix, prefixRoutes := range res {
		for _, info := range prefixRoutes.([]interface{}) {
			routeInfo := mustStringMap(info)
			route := parseRoute(prefix, routeInfo, config)
			routes = append(routes, route)
		}
	}

	return routes, nil
}
