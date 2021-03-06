package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDS(t *testing.T) {
	assert := assert.New(t)
	ids := IDS{}
	id := NewID('A')
	ids.Add(id)
	assert.Equal(len(ids), 1)

	b := IDS{}
	b.Add(NewID('A'), NewID('A'), NewID('A'))
	ids.Add(b...)
	assert.Equal(len(ids), 4)
}

func ExampleIDS() {
	ids := IDS{}
	ids.Add(NewID('A'))

	fmt.Println(len(ids))

	// Output:
	// 1
}

func ExampleIDS_Add() {
	ids := IDS{}
	ids.Add(NewID('A'), NewID('A'))

	fmt.Println(len(ids))

	// Output:
	// 2
}

func ExampleIDS_IndexOf() {
	ids := IDS{}
	id := NewID('K')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)

	fmt.Println(ids.IndexOf(id))

	// Output:
	// 2
}

func ExampleIDS_Contains() {
	ids := IDS{}
	id := NewID('K')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)

	fmt.Println(ids.Contains(id))
	fmt.Println(ids.Contains(NewID('A')))

	// Output:
	// true
	// false
}
func ExampleIDS_Delete() {
	ids := IDS{}
	id := NewID('K')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)

	fmt.Println(len(ids))
	fmt.Println(ids.Contains(id))
	fmt.Println(len(ids.Delete(id)))
	fmt.Println(ids.Contains(id))

	// Output:
	// 3
	// true
	// 2
	// false
}

func ExampleIDS_DeleteIndex() {
	ids := IDS{}
	id := NewID('K')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)

	fmt.Println(len(ids))
	fmt.Println(ids.Contains(id))
	fmt.Println(len(ids.DeleteIndex(1, 2)))
	fmt.Println(ids.Contains(id))

	// Output:
	// 3
	// true
	// 1
	// false
}

func ExampleIDS_Empty() {
	ids := IDS{}

	fmt.Println(ids.Empty())
	ids.Add(NewID('A'))
	fmt.Println(ids.Empty())

	// Output:
	// true
	// false
}
func ExampleIDS_Intersect() {
	ids := IDS{}
	id := NewID('K')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)
	ids.Intersect(id, NewID('A'))

	fmt.Println(len(ids))

	// Output:
	// 1
}
func ExampleIDS_Reverse() {
	ids := IDS{}
	id := NewID('A')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)

	ids.Reverse()

	fmt.Println(ids[0] == id)

	// Output:
	// true
}
func ExampleIDS_Distinct() {
	ids := IDS{}
	id := NewID('A')
	ids.Add(NewID('A'), NewID('A'))
	ids.Add(id)
	ids.Add(NewID('A'), NewID('A'), NewID('A'), NewID('A'))
	ids.Add(id)

	fmt.Println(len(ids))
	ids.Distinct()
	fmt.Println(len(ids))
	fmt.Println(ids[2] == id)

	// Output:
	// 8
	// 7
	// true
}
func BenchmarkContains(b *testing.B) {
	b.StopTimer()
	ids := IDS{}
	id := NewID('A')
	for i := 0; i < 1000; i++ {
		ids.Add(NewID('A'))
	}
	ids.Add(id)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ids.Contains(id)
	}
}

func BenchmarkDistinct(b *testing.B) {
	b.StopTimer()
	ids := IDS{}
	id := NewID('A')
	ids.Add(id)
	for i := 0; i < 1000; i++ {
		ids.Add(NewID('A'))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ids.Distinct()
		ids.Add(id)
	}
}
