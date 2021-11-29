package gamesettings

func safeMapBoolAdd(m map[string]bool, k string, v bool) map[string]bool {
	if m == nil {
		m = make(map[string]bool)
	}

	m[k] = v
	return m
}
