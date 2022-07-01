package main

import "fmt"

func main() {
	// var a []int = []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(cap(a[:3])),

	// var add, sub = calFunc(5, 6)
	// fmt.Println("add result= ", add, ", sub result= ", sub)
	// var a, s = calFunc1(5, 6)
	// fmt.Println("add result= ", a, ", sub result= ", s)

	// rangs()

	// maps()

	// datatypes()

	// pointers()

	// structType()

	structMethods()
}

func calFunc(x int, y int) (int, int) {
	return x + y, x - y
}

func calFunc1(x int, y int) (add int, sub int) {
	defer fmt.Println("defer useful for cleaning purpose")
	add = x + y
	sub = x - y
	fmt.Println("before return")
	return
}

func rangs() {
	var a []int = []int{1, 3, 24, 65, 7, 98, 3}
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// for i, element := range a {
	// 	fmt.Printf("%d: %d\n", i, element)
	// }

	// for _, element := range a {
	// 	fmt.Printf("%d\n", element)
	// }

	// for i, element1 := range a {
	// 	for j, element2 := range a[i+1:] {

	// 		if element1 == element2 && j > i {
	// 			fmt.Println("duplicate= ", element1)
	// 		}
	// 	}
	// 	fmt.Println("round finish")
	// }

	for i, element1 := range a {
		for j := i + 1; j < len(a); j++ {
			if element1 == a[j] && j > i {
				fmt.Println("duplicate= ", element1)
			}
		}
		fmt.Println("round finish")
	}
}

func maps() {
	var mp map[string]int = map[string]int{
		"Mani":    1,
		"Manisha": 2,
		"Manish":  3,
	}

	fmt.Println(mp)

	mp["Hani"] = 9

	fmt.Println("after adding Hani ", mp)
	delete(mp, "Manish")
	fmt.Println("after deltetion of Manish ", mp)
	val, ok := mp["Manish"]
	fmt.Println("val= ", val, ", is Manish exist= ", ok)
	fmt.Println("length of map is= ", len(mp))

}

//mutable and immutable
func datatypes() { //slices and maps only mutable, arrays are not
	var x []int = []int{1, 2, 3}
	y := x
	y[0] = 100
	fmt.Println("slice=> ", x, y)

	var a map[string]int = map[string]int{"hello": 1}
	b := a
	b["hi"] = 2
	a["bye"] = 3
	fmt.Println("map=> ", a, b)

	var x1 [3]int = [3]int{1, 2, 3}
	y1 := x1
	y1[0] = 100
	fmt.Println("array=> ", x1, y1)

	var x2 []int = []int{4, 5, 6}
	fmt.Println("(slice)x2= ", x2)
	makeChange(x2)
	fmt.Println("x2 after calling makeCahnge func= ", x2)
}

func makeChange(slice []int) {
	slice[0] = 999
}

// pointers and derefrence operators (& ans *)
func pointers() {
	x := 9
	y := &x
	fmt.Println(x, y)
	*y = 99
	fmt.Println(x, y)

	toChange := "hello"
	fmt.Println("toChange= ", toChange)
	changeValue1(&toChange)
	fmt.Println("toChange(changeValue1)= ", toChange)
	changeValue2(toChange)
	fmt.Println("toChange(changeValue2)= ", toChange)
	var pointer *string = &toChange
	fmt.Println(pointer, *pointer)

}

func changeValue1(str *string) {
	*str = "changed(1)!"
}

func changeValue2(str string) {
	str = "changed(2)!"
}

func structType() {
	var p1 Emp = Emp{"Manisha", 9, Dept{"Tech", 200}}
	var p2 Emp = Emp{"Mani", 1, Dept{"IT", 50}}
	fmt.Println(p1.name, p1.dept.name)
	fmt.Println(p1.name, p2.name)
	changeName(&p2)
	fmt.Println("p2= ", p2)
}

type Emp struct {
	name string
	id   int32
	dept Dept
}

type Dept struct {
	name  string
	count int
}

func changeName(emp *Emp) {
	emp.name = "Manish"
}

type Student struct {
	name  string
	grand []int
	age   int32
}

func (s *Student) setAge(age int32) {
	s.age = age
}

func structMethods() {
	s1 := Student{"Manisha", []int{70, 80, 90}, 20}
	fmt.Println(s1)
	s1.setAge(24)
	fmt.Println(s1)
}

func interfaces() {

}
