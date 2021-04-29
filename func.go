package main

import "fmt"

//引数の型をまとめられる．わかりづらくなりそう
func Plus(x, y int) int {
	return x + y
}

//返り値を先に宣言して返すときはreturnだけ
//引数を先に記述して可読性あげる？
func Double(price int) (result int) {
	result = price * 2
	return
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Noreturn() {
	fmt.Println("unko")
}

//関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I am a function")
	}
}

func ReturnFunc2() func() string {
	//ここの返す関数は名前つけられないの？
	return func() string {
		return "I am a function"
	}
}

//関数を引数に取る関数
func CallFunction(f func()) {
	f()
}

//クロージャー 一個前に渡したやつが帰ってくる関数
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレータ どんなジェネレータかを最初に簡単でいいから説明してよ！
// 呼ばれるたびに1増えるんかiが保持される
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	q, r := Div(9, 2)
	// q, _ := Div(9, 2)
	fmt.Println(q, r)

	ans := Double(1000)
	fmt.Println(ans)

	Noreturn()

	//無名関数
	f := func(x, y int) int {
		return x + y
	}
	i = f(1, 2)
	fmt.Println(i)

	//関数のあとに引数を渡す方法
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)

	//関数を返す関数 wtf??
	//f2に関数が入ってそのなかの関数を実行できる
	f2 := ReturnFunc()
	f2()

	//f3に文字列を返す関数が入ってるからそれを呼び出す
	f3 := ReturnFunc2()
	fmt.Println(f3())

	//関数を引数に取る関数
	CallFunction(func() {
		//これはCallfunctionのf()で実行される
		fmt.Println("I am function")
	})

	//クロージャー 非常に重要とか言って説明テキトーだし！FUCK
	g := Later()
	fmt.Println(g("1"))
	fmt.Println(g("2"))
	fmt.Println(g("3"))
	fmt.Println(g("4"))

	// ジェネレータ 連続した値を返し続ける
	// くろーじゃをおうよう
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	//intsのiを初期化する方法とかあるん？

}
