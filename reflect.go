package main

import (
	"fmt"
	"reflect"
)

type People struct {
	name string
	age  int
	sex  string
}

func Say(str string) {
	fmt.Println("hello ", str)
}
func (p People) Say(str string) {
	fmt.Println("hello ", str)
}
func (p People) Game() {
	fmt.Println(" I'm playing")
}
func (p People) Study(x, y int) {
	fmt.Println(x, y)
}
func main() {
	// //反射 reflect
	// var x int = 3
	// fmt.Println(reflect.ValueOf(x))
	// fmt.Println(reflect.TypeOf(x))
	// t := reflect.ValueOf(x)
	// fmt.Println("kind is int", t.Kind() == reflect.Int)
	// fmt.Println(t.Int())
	// fmt.Println(t.Type())

	//reflect 调用方法
	p := People{"lp", 21, "boy"}
	value := reflect.ValueOf(p)
	methon := value.MethodByName("Game")
	arg1 := make([]reflect.Value, 0)
	methon.Call(arg1)

	methon1 := value.MethodByName("Study")
	arg2 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200)}
	methon1.Call(arg2)

	methon2 := value.MethodByName("Say")
	arg3 := []reflect.Value{reflect.ValueOf("my love!")}
	methon2.Call(arg3)

	//reflect 调用函数
	fun := Say
	value1 := reflect.ValueOf(fun)
	arg4 := []reflect.Value{reflect.ValueOf("反射机制，调用函数")}
	value1.Call(arg4)
}
