package api

// Things
type Things []Thing

func (slice Things) Len() int {
	return len(slice)
}

func (slice Things) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice Things) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
