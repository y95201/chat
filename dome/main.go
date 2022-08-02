package main

import "fmt"

func main() {
	//var arr1 [5]int
	//arr2 := [3]int{1, 2, 3}
	//arr3 := [...]int{2, 4, 6, 8, 10}
	//
	//fmt.Println(arr1, arr2, arr3)
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n")
		}
	}
	fmt.Printf("a 值为 : %d\n", a)
	fmt.Printf("b 值为 : %d\n", b)
}
