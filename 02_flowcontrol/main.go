package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sumFunc(x int) (sum int) {
	sum = x
	// for文
	for sum < 100 {
		sum += sum
	}
	return
}

func sqrt(x float64) string {
	// if文
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if文の場合、評価するステートメントも記述できる
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func printOS() {
	// switch文
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Macintosh.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	// defer 遅延実行（この関数の一番最後に実行される）
	defer fmt.Println("遅延実行")
	// Stacking defers defer へ渡した関数は LIFO(last-in-first-out) の順番で実行される
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(sumFunc(1))

	// 無限ループ
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	printOS()

	// switch caseは、上から下へcaseを評価します。 caseの条件が一致すれば、そこで停止(自動的にbreak)します。
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(today + 1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// 条件のないswitchは、 switch true と書くことと同じです。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
