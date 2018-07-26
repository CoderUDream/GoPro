package language

import "fmt"

// defer越在后面定义越先执行
func testRunOrder1() {
	fmt.Println("test defer run order1")

	var i int = 0
	defer func() {
		i++
		fmt.Println("defer A i:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer B i:", i)
	}()
}

// defer定义在return运行后面的将不会执行
func testRunOrder2() {
	fmt.Println("test defer run order2")

	var i int = 0
	defer func() {
		i++
		fmt.Println("defer A i:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer B i:", i)
	}()

	return

	defer func() {
		fmt.Println("defer C is called")
	}()
}

//无名返回值
func testDeferReturn1() int {
	fmt.Println("test defer run order")

	var i int = 0
	defer func() {
		i++
		fmt.Println("defer A i:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer B i:", i)
	}()

	return i
}

//有名返回值
func testDeferReturn2() (i int) {
	fmt.Println("test defer run order")

	defer func() {
		i++
		fmt.Println("defer A i:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer B i:", i)
	}()

	return i
}

//有名返回值
func testDeferReturn3() *int {
	fmt.Println("test defer run order")

	var i = 0
	defer func() {
		i++
		fmt.Println("defer A i:", i)
	}()

	defer func() {
		i++
		fmt.Println("defer B i:", i)
	}()

	return &i
}

func TestDefer() {
	testRunOrder1()
	fmt.Println("---------------------------")
	testRunOrder2()
	fmt.Println("---------------------------")
	fmt.Println(testDeferReturn1())
	fmt.Println("---------------------------")
	fmt.Println(testDeferReturn2())
	fmt.Println("---------------------------")
	fmt.Println(*testDeferReturn3())
	fmt.Println("结论：\n" +
		"defer的用途：1.类似于c++析构函数，用于函数最后资源释放等\n" +
		"defer的执行顺序：1.是越在后面定义的越先执行\n" +
		"				2.defer在return运行之后定义的将不会执行\n" +
		"defer对返回值的修改：1.如果返回是无名且非地址  修改不生效 2.如果返回是有名或者地址 修改生效")
}
