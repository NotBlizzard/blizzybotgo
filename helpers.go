package main

func Keys(maps map[string]func(string) string) []string {
	keys := make([]string, 0, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
