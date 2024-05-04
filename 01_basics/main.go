package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
)

// 何も代入しない限り、false扱い
var c, python, java bool

// func add(x, y int) int { （関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。）
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。これを "naked" return と呼びます。
// この場合だと、x,yをintで返すことを明示しているため、returnを書いたらその時点での引数を明示できる
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 変数をまとめて宣言
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// 標準出力
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(42, 13))

	a, b := swap("world", "hello")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int = 1
	// 関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
	j := "string"
	fmt.Println(i, j, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 定数
	// 定数は、文字(character)、文字列(string)、boolean、数値(numeric)のみで使えます. なお、定数は := を使って宣言できません。
	const World = "世界"
	fmt.Println("Hello", World)

	// 浮動小数点の関係で実行できない
	// fmt.Println(Big)
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
