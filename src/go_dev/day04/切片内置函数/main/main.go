/*切片的拷贝
string 底层就是一个byte数组，因此也可以进行切片操作
string 不可变;底层就是ptr;len
改变string中的值
*/

package main

import "fmt"

func copySlice() {
	s1 := []int{1,2,3,4}
	s2 := make([]int, 8)

	copy(s2, s1)          // s1拷贝到s3，拷贝到前面
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	s3 := []int{1,2,3}
	s3 = append(s3, s2...)
	fmt.Println("s3=", s3)
	s3 = append(s3,4,5,6)   //可变参数
	fmt.Println("s3=", s3)
}


func modStr(){
	str := "Hello World"
	st := str[:4]
	fmt.Println(st)

	s2 := []rune(str)    //有中文用rune()
	s2[1] = 'E'
	s3 := string(s2)
	fmt.Println("s3=", s3)

}
func main() {
	copySlice()
	modStr()
}

