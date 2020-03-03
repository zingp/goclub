package statmentAndInit


//声明一个数组，并设置为零值
var Array [5]int

//使用数组字面量声明数组 需要赋值 用=
var Array2 = [5]int{10, 20, 30, 40, 50}

//让 Go 自动计算声明数组的长度
var Array3 = [...]int{1, 2, 3, 4, 5}

// 声明一个有 5 个元素的数组
// 用具体值初始化索引为 1 和 2 的元素
// 其余元素保持零值
var Array4 = [5]int{1: 10, 2: 20}
