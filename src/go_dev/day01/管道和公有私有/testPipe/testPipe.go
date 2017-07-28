package testPipe
/*
公有成员首字母大写，首字母小写为私有成员；
往管道里放入值是pipe <- value.
*/
func AddFunc(a int, b int, c chan int){
	c <- a
	c <- b
	c <- (a+b)
}
