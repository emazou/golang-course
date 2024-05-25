package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) { // any is a type that can be any type like int, string, float64, bool, etc.
	for _, item := range list {
		fmt.Println(item)
	}
}

// for to create restrictions, we can use interfaces
type Numbers interface {
	~int | ~float64 | ~uint | ~uint64
}

// we can use restrictions of packages that GO

/*
Arbitrary type parameter T is declared in square brackets after the function name and union operator
~ is used to specify the type of T. In this case, T can be either int or float64.
*/
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// comparable is a type that can be compared with == operator
func Includes[T comparable](list []T, item T) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	var result []T
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

// generic struct
type Product[T uint | string] struct {
	Id       T
	Name     string
	Price    float64
	Quantity int
}

func main() {
	PrintList(1, 2, 3, 4, 5)
	PrintList("a", "b", "c", "d", "e")
	PrintList(1.1, 2.2, 3.3, 4.4, 5.5)
	PrintList(true, false, true, false, true)
	fmt.Println(Sum(1, 2.3, 3, 4, 5))
	fmt.Println(Includes([]int{1, 2, 3, 4, 5}, 3))

	fmt.Println(Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n%2 == 0
	}))

	product1 := Product[uint]{Id: 1, Name: "Product 1", Price: 10.5, Quantity: 5}
	fmt.Println(product1)
	product2 := Product[string]{Id: "sdafhkjsdfh", Name: "Product 2", Price: 20.5, Quantity: 10}
	fmt.Println(product2)
}
