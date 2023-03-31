package helper

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		if s == "*" {
			return true
		}
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
