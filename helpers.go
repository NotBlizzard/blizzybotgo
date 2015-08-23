package main

func Keys(maps map[string]func(string) string) []string {
	keys := make([]string, 0, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

func Contains(needle string, haystack []string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
