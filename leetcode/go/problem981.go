type TimeMap struct {
	m map[Key]string
	t map[string][]int
}

type Key struct {
	k string
	t int
}

func Constructor() TimeMap {
	a := make(map[Key]string)
	b := make(map[string][]int)
	return TimeMap{a, b}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[Key{key, timestamp}] = value
	this.t[key] = append(this.t[key], timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	l := 0
	r := len(this.t[key])
	arr := this.t[key]
	fmt.Println("done for", timestamp)
	for l < r {
		m := (r + l) / 2
		if arr[m] > timestamp {
			if r == m {
				break
			}
			r = m
		} else if arr[m] < timestamp {
			if l == m {
				break
			}
			l = m
		} else {
			return this.m[Key{key, arr[m]}]
		}
	}
	if l >= 0 && l < len(arr) && arr[l] <= timestamp {
		return this.m[Key{key, arr[l]}]
	} else {
		return ""
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
