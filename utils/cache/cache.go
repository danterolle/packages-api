package cache

// data is a map that stores key-value pairs
var data = make(map[string]interface{})

// Get function retrieves the value associated with a given key
// It returns the value and a boolean indicating if the key exists
func Get(key string) (interface{}, bool) {
	val, ok := data[key]
	return val, ok
}

// Set function associates a given value with a given key
func Set(key string, value interface{}) {
	data[key] = value
}
