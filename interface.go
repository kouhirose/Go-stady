package main

import "fmt"

//異なる型に共通の性質を付与？？
// type Stringfy interface {
// 	Tostring() string
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p *Person) ToString() string {
// 	return fmt.Sprintf("Name=%v,Age=%v", p.Name, p.Age)
// }

// type Car struct {
// 	Number string
// 	Model  string
// }

// func (c *Car) ToString() string {
// 	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
// }

// func main() {
// 	vs := []Stringfy{
// 		&Person{Name: "Taro", Age: 21},
// 		&Car{Numer: "123-456", Model: "AB-unko"},
// 	}

// 	for _, v := range vs {
// 		fmt.Println(v.ToString())
// 	}
// }

// カスタムエラー
type MyError struct {
	Message string
	ErrCode int
}

//継承してる？
func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラー発生", ErrCode: 12334}
}

//ストリンガー

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	//型アサーションで復元して使用
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

	//ストリンガー
}
