package main

import (
	"fmt"
	"strings"
)

// struct (構造体)は、フィールド( field )の集まりです。
type Vertex struct {
	X int
	Y int
}

var (
	v1    = Vertex{1, 2}  // has type Vertex
	v2    = Vertex{X: 1}  // Y:0 is implicit
	v3    = Vertex{}      // X:0 and Y:0
	point = &Vertex{1, 2} // has type *Vertex
)

// Goはポインタを扱います。 ポインタは値のメモリアドレスを指します。
// & オペレータは、そのオペランド( operand )へのポインタを引き出します。
// * オペレータは、ポインタの指す先の変数を示します。
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// 構造体
	fmt.Println(Vertex{1, 2})
	// structのフィールドは、ドット( . )を用いてアクセスします。
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	fmt.Println(v1, point, v2, v3)

	// 配列（固定長）
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice（可変長）スライスは長さ( length )と容量( capacity )の両方を持っています。
	slices := [6]int{2, 3, 5, 7, 11, 13}

	// 配列内の要素1~3番目まで（4つ目は含まない）
	var s []int = slices[1:4]
	fmt.Println(s)

	// すべて同じスライス
	sliceNumbers := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sliceNumbers[0:6])
	fmt.Println(sliceNumbers[:6])
	fmt.Println(sliceNumbers[0:])
	fmt.Println(sliceNumbers[:])

	var sl []int
	fmt.Println(sl, len(sl), cap(sl))
	if sl == nil {
		fmt.Println("nil!")
	}

	// スライスは、組み込みの make 関数を使用して作成することができます。 これは、動的サイズの配列を作成する方法です.
	// make 関数はゼロ化された配列を割り当て、その配列を指すスライスを返します。
	aaa := make([]int, 5)
	printSlice("aaa", aaa)
	bbb := make([]int, 0, 5)
	printSlice("bbb", bbb)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
