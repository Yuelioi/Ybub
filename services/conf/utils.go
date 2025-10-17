package conf

func findIndexByID[T any](items []T, id string, getID func(T) string) (int, bool) {
	for i, v := range items {
		if getID(v) == id {
			return i, true
		}
	}
	return -1, false
}
