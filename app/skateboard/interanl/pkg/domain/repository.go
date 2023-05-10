package domain

func OffsetLimit(page, size int) (offset int, limit int) {
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 20
	}
	offset = (page - 1) * size
	limit = size
	return
}

type QueryOptions map[string]interface{}

func NewQueryOptions() QueryOptions {
	options := make(map[string]interface{})
	return options
}
func (options QueryOptions) WithOffsetLimit(page, size int) QueryOptions {
	offset, limit := OffsetLimit(page, size)
	options["offset"] = offset
	options["limit"] = limit
	return options
}

func (options QueryOptions) WithKV(key string, value interface{}) QueryOptions {
	options[key] = value
	return options
}
