type TimeMap struct {
	store map[string][]struct {
		timestamp int
		value     string
	}
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]struct {
			timestamp int
			value     string
		}),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], struct {
		timestamp int
		value     string
	}{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.store[key]
	if !ok {
		return ""
	}

	left, right := 0, len(entries)-1
	res := ""

	for left <= right {
		mid := left + (right-left)/2

		if entries[mid].timestamp <= timestamp {
			res = entries[mid].value
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
