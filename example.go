package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	//1
	fmt.Println("Hello, World!")

	//2
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	//3
	if x > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("Less than 2")
	} else {
		fmt.Println("Between 2 and 6 :)")
	}

	//4 arrays (fix size only)
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	//5 Slices ("dynamic arrays")
	c := []int{5, 4, 3, 2, 1}
	fmt.Println(c)
	c = append(c, 13)
	fmt.Println(c)
	c = append(c, 13, 33)
	fmt.Println(c)

	//6 Map
	fmt.Println("Map")
	figures := make(map[string]int)
	figures["triangle"] = 2
	figures["square"] = 3
	figures["dodecagon"] = 12
	fmt.Println(figures)
	fmt.Println(figures["square"])
	delete(figures, "square")
	fmt.Println(figures)

	//7 Loops. Only "for" loop
	fmt.Println("For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("For loop Range")
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index : ", index, "value : ", value)
	}

	fmt.Println("For loop Range in Map")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key : ", key, "value : ", value)
	}

	//8 Functions
	result := summ(22, 55)
	fmt.Println("Result of sum() func : ", result)

	result2, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	//9 Struct
	p := person{name: "Ilnur", age: 29}
	fmt.Println(p)
	fmt.Println(p.age)

	//10 Pointers
	i := 7
	fmt.Println("Var :", i, " and address: ", &i)
	inc(&i)
	fmt.Println("Var :", i, " and address: ", &i)

}

func summ(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func inc(i *int) {
	*i++
}
