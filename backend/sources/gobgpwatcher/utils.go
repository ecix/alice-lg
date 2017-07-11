package gobgpwatcher

import (
	"time"
)

// Assert string, provide default
func mustString(value interface{}, fallback string) string {
	sval, ok := value.(string)
	if !ok {
		return fallback
	}
	return sval
}

func mustInt(value interface{}, fallback int) int {
	sval, ok := value.(float64)
	if !ok {
		return fallback
	}
	return int(sval)
}

func mustStringMap(value interface{}) map[string]interface{} {
	sval, ok := value.(map[string]interface{})

	if !ok {
		return make(map[string]interface{})
	}

	return sval
}

func mustStringMapList(value interface{}) []map[string]interface{} {
	res := []map[string]interface{}{}
	list, ok := value.([]interface{})
	if !ok {
		return res
	}
	for _, v := range list {
		res = append(res, mustStringMap(v))
	}
	return res
}

func mustDurationMs(value interface{}, fallback time.Duration) time.Duration {
	sval, ok := value.(time.Duration)
	if !ok {
		return fallback
	}

	return sval
}
