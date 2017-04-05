/**
 * 切片
 */
package main

import (
	"fmt"
)

func main() {

	// 存放int类型切片
	var s1 []int
	s1 = append(s1, 11, 22, 33, 44)
	fmt.Println(s1)

	// 存放复合类型切片
	var s2 []interface{}
	s2 = append(s2, 11, "aa", 3.5, true, nil, 8/4, 2<<1, s1)
	fmt.Println(s2)

	// 存放字符类型切片
	s3 := []string{"a", "b", "c", "d"}
	fmt.Println(s3)

	// 切片的索引以及长度
	s5 := []string{"a", 2: "b", "c", 6: "d"}
	fmt.Println(s5)
	fmt.Printf("s5.len:%d\n", len(s5))

	// 切片的长度以及容量
	s4 := s3[1:]
	fmt.Println(s4)
	fmt.Printf("s4.len:%d, s4.cap:%d\n", len(s4), cap(s4))

	// 基于数组生成切片
	arr := [6]int{10, 20, 30, 40, 50, 60}
	fmt.Println(arr)

	s7 := arr[1:4]
	fmt.Println(s7)

	s8 := arr[2:4]
	fmt.Println(s8)

	// 数组元素改变, 切片也变
	arr[3] = 33
	fmt.Println(s7, s8)

	fmt.Printf("arr:%p, s7:%p\n", &arr, &s7)

	var s9 []string
	fmt.Printf("s9.ptr:%p ptr:%p\n", s9, &s9)

}
