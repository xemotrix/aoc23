package main

import (
	"fmt"
	"unsafe"
)

type Named interface {
	GetName() string
}

type Lol int

func (l Lol) GetName() string {
	return "lol"
}

func main() {

	s := "hello"
	printRepr(s)

	// s[0] = byte('H')
	// printRepr(s)

	s += " world"
	printRepr(s)

	fmt.Println("Bytes")
	b := []byte(s)
	printRepr(b)

	b[0] = byte('H')
	printRepr(b)

	b = append(b, '!')
	printRepr(b)

	b = append(b, []byte("!!!!!!!!")...)
	printRepr(b)

	sb := string(b)
	printRepr(sb)

	fmt.Println("Runes")

	printRepr(s)
	r := []rune(s)
	printRepr(r)

	r[0] = rune('H')
	printRepr(r)

	r = append(r, rune('!'))
	printRepr(r)

	// l := Lol(0xdeadbeef)
	// printRepr(l)
	// uint64
	// t := reflect.TypeOf(l)
	// fmt.Println(t)
	// fmt.Println(t.Name())

	// var nl Named = l
	// printRepr(nl)

	// x := *(*[2]uint64)(unsafe.Pointer(&nl))
	// // printRepr(x)

	// r := *(*uint64)(unsafe.Pointer(uintptr(x[1])))

	// typeData := *(*[128]uint64)(unsafe.Pointer(uintptr(x[0])))
	// typeDataStr := *(*[128 * 8]byte)(unsafe.Pointer(&typeData))
	// fmt.Printf("typeDataStr:\n%s\n", typeDataStr)

	// printRepr(r)
	// printRepr(typeData)

	// printRepr(l.GetName)
	// fpx := *(*uintptr)(unsafe.Pointer(&fp))
	// fmt.Printf("fp %x, size: %d\n", fpx, unsafe.Sizeof(fp))

	// t := *(*[2]uint64)(unsafe.Pointer(&Lol))

}
func printRepr[T any](thing T) {
	nums := *(*[128]uint64)(unsafe.Pointer(&thing))

	fmt.Printf("%T [%d]: %v\n", thing, unsafe.Sizeof(thing), thing)

	for i := 0; i < len(nums); i++ {
		if uintptr(i) >= unsafe.Sizeof(thing)/8 {
			break
		}
		fmt.Printf("  %016x\n", nums[i])
	}
	fmt.Println()
}
