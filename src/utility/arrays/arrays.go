package arrays

import (
	"math/rand"
	"sort"
	"time"
)

const (
	ASC = iota
	DESC
)

// Sort 将data表示的slice排序， sortFlag表示升序还是降序
func Sort(data interface{}, sortFlag int) {
	switch tmpSlice := data.(type) {
	case []int:
		if sortFlag == ASC {
			sort.IntSlice(tmpSlice).Sort()
		} else {
			slice := intSlice{sort.IntSlice(tmpSlice), true}
			sort.Sort(slice)
		}
	case []float64:
		if sortFlag == ASC {
			sort.Float64Slice{tmpSlice}.Sort()
		} else {
			slice := float64Slice{sort.Float64Slice(tmpSlice), true}
			sort.Sort(slice)
		}
	case []string:
		if sortFlag == ASC {
			sort.StringSlice(tmpSlice).Sort()
		} else {
			slice := stringSlice{sort.StringSlice(tmpSlice), true}
			sort.Sort(slice)
		}
	}
}

var initRand bool

// Shuffle 将data表示的slice中的元素打乱
func Shuffle(data interface{}) {
	if initRand == false {
		rand.Seed(time.Now().UnixNano())
		initRand = true
	}
	switch tmpSlice := data.(type) {
	case []int:
		slice := intSlice{sort.IntSlice(tmpSlice), false}
		sort.Sort(slice)
	case []float64:
		slice := float64Slice{sort.Float64Slice(tmpSlice), false}
		sort.Sort(slice)
	case []string:
		slice := stringSlice{sort.StringSlice(tmpSlice), false}
		sort.Sort(slice)
	}
}

type intSlice struct {
	slice  sort.IntSlice
	isSort bool
}

func (is intSlice) Len() int {
	return len(is.slice)
}

func (is intSlice) Less(i, j int) bool {
	if is.isSort {
		return is.slice[i] < is.slice[j]
	}
	return rand.Intn(2) == 1
}

func (is intSlice) Swap(i, j int) {
	is.slice[i], is.slice[j] = is.slice[j], is.slice[i]
}

type float64Slice struct {
	slice  sort.Float64Slice
	isSort bool
}

func (fs float64Slice) Len() int {
	return len(fs.slice)
}

func (fs float64Slice) Less(i, j int) bool {
	if fs.isSort {
		return fs.slice[i] < fs.slice[j]
	}
	return rand.Intn(2) == 1
}

func (fs float64Slice) Swap(i, j int) {
	fs.slice[i], fs.slice[j] = fs.slice[j], fs.slice[i]
}

type stringSlice struct {
	slice  sort.StringSlice
	isSort bool
}

func (ss stringSlice) Len() int {
	return len(ss.slice)
}

func (ss stringSlice) Less(i, j int) bool {
	if ss.isSort {
		return ss.slice[i] < ss.slice[j]
	}
}

func (ss stringSlice) Swap(i, j int) {
	ss.slice[i], ss.slice[j] = ss.slice[j], ss.slice[i]
}
