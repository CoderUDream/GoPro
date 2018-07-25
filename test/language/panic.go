package language

import (
	"fmt"
)

func funcA() {
	fmt.Println("funcA~~~~~~")
}

func funcB() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("funcB catch a Panic! is:", err)
		}
	}()

	fmt.Println("funcB start a Panic!")
	panic("B is panic")

	fmt.Println("after panic i'm working")
}

func funcC() {
	fmt.Println("funcC~~~~~~")
}

func TestPanic() {
	funcA()

	funcB()

	funcC()
}
