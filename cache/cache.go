package cache

var data = make(map[string]interface{})

func Get(key string) (interface{}, bool) {
	val, ok := data[key]
	return val, ok
}

func Set(key string, value interface{}) {
	data[key] = value
}
