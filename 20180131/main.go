package part01

import (
	"fmt"
	"unsafe"
)

var (
	u1 uint8
	u2 uint16
	u3 uint32
	u4 uint64
	u5 uint

	i1 int8
	i2 int16
	i3 int32
	i4 int64
	i5 int

	f1 float32
	f2 float64

	s string

	b  byte
	r  rune
	p  uintptr
	bl bool
)

func main() {
	fmt.Println("uint8:", u1, "size:", unsafe.Sizeof(u1))
	fmt.Println("uint16:", u2, "size:", unsafe.Sizeof(u2))
	fmt.Println("uint32:", u3, "size:", unsafe.Sizeof(u3))
	fmt.Println("uint64:", u4, "size:", unsafe.Sizeof(u4))
	fmt.Println("uint:", u5, "size:", unsafe.Sizeof(u5))
	fmt.Println("int8:", i1, "size:", unsafe.Sizeof(i1))
	fmt.Println("int16:", i2, "size:", unsafe.Sizeof(i2))
	fmt.Println("int32:", i3, "size:", unsafe.Sizeof(i3))
	fmt.Println("int64:", i4, "size:", unsafe.Sizeof(i4))
	fmt.Println("int:", i5, "size:", unsafe.Sizeof(i5))
	fmt.Println("float32:", f1, "size:", unsafe.Sizeof(f1))
	fmt.Println("float64:", f2, "size:", unsafe.Sizeof(f2))
	fmt.Println("string:", s, "size:", unsafe.Sizeof(s))
	fmt.Println("byte:", b, "size:", unsafe.Sizeof(b))
	fmt.Println("rune:", r, "size:", unsafe.Sizeof(r))
	fmt.Println("uintptr:", p, "size:", unsafe.Sizeof(p))
	fmt.Println("bool:", bl, "size:", unsafe.Sizeof(bl))
}
