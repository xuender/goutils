package goutils

import (
	"sort"
	"sync"
)

type result struct {
	data  interface{}
	point interface{}
}

// Results 最优结果集.
type Results struct {
	Len   int
	datas []result
	size  int
	less  func(i, j interface{}) bool
	mutex sync.Mutex
}

// Add 增加结果.
func (r *Results) Add(data, point interface{}) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.datas[r.Len] = result{
		data:  data,
		point: point,
	}
	r.Len++
	sort.Slice(r.datas[:r.Len], func(i, j int) bool { return r.less(r.datas[i].point, r.datas[j].point) })
	if r.Len > r.size {
		r.Len = r.size
	}
}

// AddResults 增加结果集.
func (r *Results) AddResults(results *Results) {
	for i := 0; i < results.Len; i++ {
		r.Add(results.Get(i))
	}
}

// Get 获取数据.
func (r *Results) Get(i int) (interface{}, interface{}) {
	return r.datas[i].data, r.datas[i].point
}

// GetData 数据.
func (r *Results) GetData(i int) interface{} {
	return r.datas[i].data
}

// GetPoint 得分.
func (r *Results) GetPoint(i int) interface{} {
	return r.datas[i].point
}

// GetSize 尺寸.
func (r *Results) GetSize() int {
	return r.size
}

// NewResults 新建结果集.
func NewResults(size int, less func(i, j interface{}) bool) *Results {
	return &Results{
		Len:   0,
		datas: make([]result, size+1),
		size:  size,
		less:  less,
	}
}
