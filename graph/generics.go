package graph

func getGraphItem[T comparable](collection map[string]T, name string, errType error) (item T, err error) {
	item, ok := collection[name]
	if !ok {
		return item, errType
	}
	return

}

func addGraphItem[T comparable](collection map[string]T, name string, item T, errType error) (err error) {
	var value T

	value, ok := collection[name]
	if ok && (value != item) {
		return errType
	}
	if !ok {
		collection[name] = item
	}
	return
}
