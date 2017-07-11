package gobgpwatcher

// Assert string, provide default
func mustString(value interface{}, fallback string) string {
	sval, ok := value.(string)
	if !ok {
		return fallback
	}
	return sval
}

func mustStringMap(value interface{}) map[string]interface{} {
	sval, ok := value.(map[string]interface{})

	if !ok {
		return make(map[string]interface{})
	}

	return sval
}
