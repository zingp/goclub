package main

import(
	"reflect"
	"fmt"
)
// 定义一个People结构体
type People struct {
	Name string
	Age int
}

// 方法必须大写外部才可见
func (p *People)TalkSelf(){
	fmt.Printf("My name is %s", p.Name)
}

func (p People)SetAge(a int)People{
	p.Age = a
	return p
}

//  reflect.TypeOf()
func testTypeOf(a interface{}) {

	t := reflect.TypeOf(a)  // 根据对象反射类型
	fmt.Println("TypeOf:", t)
	fmt.Println("Kind:", t.Kind())	
}

func testValueOf(a interface{}) {
	v := reflect.ValueOf(a)   // 根据对象反射值
	fmt.Println("ValueOf:", v)
	k1 := v.Kind()
	fmt.Println("Kind:", k1)
}

func testNumField(a interface{}) {
	v := reflect.ValueOf(a)
	n := v.NumField()
	fmt.Println("field nums:", n)
}

func testNumMethod(a interface{}) {
	v := reflect.ValueOf(a)
	n := v.NumMethod()
	fmt.Println("method nums:", n)
}

func testMethodByName(a interface{}) {
	v := reflect.ValueOf(a)
	r := v.MethodByName("TalkSelf")
	fmt.Println(r)

	// var b []reflect.Value
	// b = append(b,)
	r.Call(nil)

}

func testElemCall(a interface{}) {
	v := reflect.ValueOf(a)
	r := v.MethodByName("SetAge")
	fmt.Println(r)

	var b []reflect.Value
	age := 99
	newAge := reflect.ValueOf(age)
	b = append(b, newAge)

	p := r.Call(b)     // 返回的是一个切片
	fmt.Println(p)
	fmt.Println(p[0])
}
func main() {
	var p = People {
		Name: "liu",
		Age: 23,
	}
	//  testTypeOf(p)
	//  testValueOf(p)

	//  var name = "zingp"
	//  testTypeOf(name)
	//  testValueOf(name)

	//  testNumField(p)
	
	// testNumMethod(p)    // 对象其实并没有实现方法，所以输出结果是0
	// testNumMethod(&p)   // 指针类型实现了一个方法，所以输出结果是1

	// testMethodByName(&p)
	// p2 := p.SetAge(100)
	// fmt.Println(p2)

	testElemCall(p)

}