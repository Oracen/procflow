package topo

func getAllItems[T comparable](inputCollection map[string]T, copy bool) (collection map[string]T) {
	if !copy {
		return inputCollection
	}
	collection = map[string]T{}
	for key, value := range inputCollection {
		collection[key] = value
	}
	return
}

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

func mergeMapItems[T comparable](map1, map2 map[string]T, errType error) (merged map[string]T, err error) {
	merged = map[string]T{}
	for key, value := range map1 {
		// Create copy for safety
		merged[key] = value
	}

	for key, value := range map2 {
		err = addGraphItem(merged, key, value, errType)
		if err != nil {
			// TODO: Better error communication
			return map[string]T{}, err
		}
	}
	return
}
