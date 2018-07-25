package makeandnew

import (
"fmt"
"strconv"
)

type Student struct {
	name string
	age int
}

func testBasic() {
	//定义普通变量
	var i1 int
	i2 := 3
	fmt.Println(fmt.Sprintf("i1:%d, i2:%d", i1, i2))

	//定义结构体变量（不涉及指针等）
	var stu1 Student
	stu1.name = "wang"
	stu1.age = 1
	fmt.Println(fmt.Sprintf("i1:%s, i2:%d", stu1.name, stu1.age))

	var stu2 = Student{"lili", 2}
	fmt.Println(fmt.Sprintf("i1:%s, i2:%d", stu2.name, stu2.age))

	stu3 := Student{"lili", 2}
	fmt.Println(fmt.Sprintf("i1:%s, i2:%d", stu3.name, stu3.age))
}

// 指针和引用
func pointerAndRef() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var i *int 		//定义一个指针
	*i = 10	   		//panic  没有分配地址
	fmt.Println(i)  //panic  没有分配地址

	//同样的Studeng结构体 也会出现这个问题，所以引入make 和 直接初始化
	var i2 *int
	i2 = new(int)  //也可以 i2 := 2  //直接分配内存且初始化
	*i2 = 10
	fmt.Println(i2)  //正确

	var stu1 *Student	//指针类型
	// 也可以使用 stu1 := Student{"zhao, 1} 这种方式初始化的为引用类型
	stu1 = new(Student)
	stu1.name = "zhao"
	stu1.age = 1
	fmt.Println(fmt.Sprintf("i1:%s, i2:%d", stu1.name, stu1.age))
	fmt.Println("new: 一般情况下不使用new 直接初始化")
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d  %v\n", len(s), cap(s), s)
}

//make slice
func testMakeSlice() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("make只用于chan、map以及切片的内存创建, 并且重置这块内存数据为0")

	//切片 slice
	var numbers1 [] int = make([]int, 0, 10)
	printSlice(numbers1)

	//分配了len=5 cap（容量）=10  这样索引最多到第五个元素
	//分配10的cap意义在于当需要扩容的时候，因为预先分配了16大小，将不需要再分配在16范围内的大小，节约时间和内存
	numbers2 := make([]int, 5, 16)  //
	printSlice(numbers2)
	numbers3 := []int{1, 2, 3, 4}
	printSlice(numbers3)

	fmt.Println("numbers2[3] :" + strconv.Itoa(numbers2[3]))
	//这将会panic 因为索引超过len
	//fmt.Println("numbers2[5] :" + strconv.Itoa(numbers2[5]))
}

func printMap(m map[string]string) {
	for key, value := range m { fmt.Println("Key:", key, "Value:", value) }
}

//make map
func testMakeMap() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var m1 map[string]string  //只定义了一指针
	var m2 map[string]string = map[string]string{"li":"1"}
	var m3 = make(map[string]string)

	m4 := map[string]string{"li":"1"}
	m5 := make(map[string]string)
	//panic 因为m1是没有分配内存的
	//m1["li"] = "1"
	printMap(m1)
	printMap(m2)
	printMap(m3)
	printMap(m4)
	printMap(m5)

	//判断不存在的值  因为不存的值 也是返回""
	//同样的value是int，string类型都存在这个问题
	key := "wang"
	if v, ok := m3[key]; ok {
		fmt.Println("key" + " is" + v)
	} else {
		fmt.Println("key: " + key + " is not exist" )
	}

	m3[key] = ""

	if v, ok := m3[key]; ok {
		fmt.Println(key + " is:" + v)
	} else {
		fmt.Println("key:" + key + "is not exist" )
	}
}


func TestMakeAndNew() {
	testBasic()

	pointerAndRef()

	testMakeSlice()

	testMakeMap()
}