package main

import "fmt"

func one(x *int) {
	*x = 1
}

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	// structの場合は自動的に参照してくれるので下記のような記述にしなくても良い
	// (*v).X = 1000
}

func main() {
	// ポインタ

	var n int = 100
	n = 200
	// one(&n)
	fmt.Println(n)

	/*
		fmt.Println(n)
		fmt.Println(&n)

		var p *int = &n

		fmt.Println(p)

		fmt.Println(*p)
	*/

	// newとmakeの違い

	/*
		var p *int = new(int)
		fmt.Println(p)
		fmt.Println(*p)
		*p++
		fmt.Println(*p)

		var p2 *int
		fmt.Println(p2)
		*p2++
		fmt.Println(ps)
	*/

	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	fmt.Println(s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)
	fmt.Println(m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)
	fmt.Println(ch)

	var p *int = new(int)
	fmt.Printf("%T\n", p)
	fmt.Println(p)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
	fmt.Println(st)

	// struct

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	// この記述方法だと要素の全てを記述しないといけない
	// v3 := Vertex{1}
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
	fmt.Printf("%T %v\n", v4, v4)

	// sliceとかmapとかだとnilになる
	var v5 Vertex
	fmt.Println(v5)
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Println(v6)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Println(v7)
	fmt.Printf("%T %v\n", v7, v7)

	// ss := make([]int, 0)
	// ss := []int{}
	// fmt.Println(ss)

	vv := Vertex{1, 2, "test"}
	changeVertex(vv)
	fmt.Println(vv)

	vv2 := &Vertex{1, 2, "test"}
	changeVertex2(vv2)
	fmt.Println(vv2)
	fmt.Println(*vv2)
}
