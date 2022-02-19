package main

import (
	"fmt"
	"unsafe"
)

// 空struct不占用内存空间，内存逃匿后都指向同一个地址

func main() {
	a := struct{ a int }{
		a: 3,
	}
	fmt.Println(unsafe.Pointer(&a))
	fmt.Println(unsafe.Sizeof(a))
	b := struct{}{}
	fmt.Println(unsafe.Pointer(&b))
	fmt.Println(unsafe.Sizeof(b))

	c := struct{}{}
	fmt.Println(unsafe.Pointer(&c))
	fmt.Println(unsafe.Sizeof(c))

	d := new(int)
	d = nil
	e := new(int)
	e = nil
	fmt.Println(unsafe.Pointer(&d))
	fmt.Println(unsafe.Pointer(&e))

}
