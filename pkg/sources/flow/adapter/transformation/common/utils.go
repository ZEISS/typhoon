package common

// ReadValue returns the source object item located at the requested path.
// nolint:gocyclo
func ReadValue(source interface{}, path map[string]interface{}) interface{} {
	var result interface{}
	for k, v := range path {
		switch value := v.(type) {
		case float64, bool, string:
			sourceMap, ok := source.(map[string]interface{})
			if !ok {
				break
			}
			result = sourceMap[k]
		case []interface{}:
			if k != "" {
				// array is inside the object
				// {"foo":[{},{},{}]}
				sourceMap, ok := source.(map[string]interface{})
				if !ok {
					break
				}
				source, ok = sourceMap[k]
				if !ok {
					break
				}
			}
			// array is a root object
			// [{},{},{}]
			sourceArr, ok := source.([]interface{})
			if !ok {
				break
			}

			index := len(value) - 1
			if index >= len(sourceArr) {
				break
			}
			result = ReadValue(sourceArr[index], value[index].(map[string]interface{}))
		case map[string]interface{}:
			if k == "" {
				result = source
				break
			}
			sourceMap, ok := source.(map[string]interface{})
			if !ok {
				break
			}
			if _, ok := sourceMap[k]; !ok {
				break
			}
			result = ReadValue(sourceMap[k], value)
		}
	}
	return result
}
