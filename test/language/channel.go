package language

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"math/rand"
)

//chan(信道) 是goroutine之间互相通讯的东西。 类似我们Unix上的管道（可以在进程间传递消息），
// 用来goroutine之间发消息和接收消息。其实，就是在做goroutine之间的内存共享

/*知识点：
1. chann的创建
2. chann的类型
   2.1 存储的数据类型
   2.2 有无缓冲类型
       默认的信道在存，取消息的时候都是阻塞的
   2.3 单方向
     定义只读和只写的channel意义不大，一般用于在参数传递中
	 //定义只读的channel  从chan中读数据
	 read_only := make (<-chan int)
	 //定义只写的channel  向chan中写数据
	 write_only := make (chan<- int)
3.select
	3.1 可以跳过阻塞，且可以对多个channel进行接收和发送

注意事项：
	1. 如何优雅地关闭Go channel https://www.jianshu.com/p/d24dfbb33781
*/

//创建chann
func createChan() {
	//都需要使用make
	var channel1 chan int = make(chan int) //使用var 定义
	channel2 := make(chan int)             //使用 := 定义

	//信道支持不同数据类型
	channel3 := make(chan string)
	channel4 := make(chan map[string]string)
	channel5 := make(chan []int)

	fmt.Println(channel1, channel2, channel3, channel4, channel5)
	defer func() {
		close(channel1)
		close(channel2)
		close(channel3)
		close(channel4)
	}()
}

//channel的有无缓冲例子
func testChanCache() {
	//无缓冲阻塞, 直到channel被写入数据才有效，且缓冲是一个队列，先进先出
	channel1 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- 1
		channel1 <- 2
		channel1 <- 3
	}()

	log.Println(fmt.Sprintf("channel1 value is:%d, %d, %d", <-channel1, <-channel1, <-channel1))
	defer close(channel1)

	time.Sleep(2 * time.Second)

	//缓冲阻塞
	channel2 := make(chan int, 3) //这个3表示channel2的容量表示 在未达到3个int数据放入channel2中，读取channel2将不会阻塞

	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			channel2 <- i
		}
	}()

	s := make([]int, 0, 3)
	for {
		s = append(s, <-channel2)
		log.Println(s)

		if len(s) == 3 {
			log.Println("channel len equals 3")
			close(channel2)
			break
		}
	}

	time.Sleep(4 * time.Second)
}

//channel的阻塞例子
func testChannelBlock() {
	//输入阻塞
	channel3 := make(chan int, 3)
	go func() {
		for i := 1; i <= 50; i++ {
			log.Println("start push data in channel3 i:" + strconv.Itoa(i))
			channel3 <- i
			log.Println("end push data in channel3 i:" + strconv.Itoa(i))
		}
		defer close(channel3)
	}()

	go func() {
		var isClosed = false
		for i := 1; i <= 100; i++ {
			time.Sleep(1 * time.Second)
			for j := 1; j <= 2; j++ {
				if v, ok := <-channel3; ok {
					log.Print("i:", strconv.Itoa(v))
				} else {
					isClosed = true
					break
				}
			}

			if isClosed {
				break
			}
		}
	}()
}

//channel单方向
func testSingleDirChannel() {
	channel1 := make(chan string)
	go func(out chan<- string) {
		out <- "hello"
		//i := <-in 这样是会panic的  chan<-只允许向chan写数据
	}(channel1)

	go func(in <-chan string) {
		str := <-in
		log.Println("str is:" + str)
		//out<-"hello" 这样是会panic的  <-chan
	}(channel1)
}

//channel多路复用
type chatServer struct {
	personASend chan string
	personARecv chan string
	personBSend chan string
	personBRecv chan string
}

func (*chatServer) speck(word string) {

}

func (*chatServer) start() {

}

func testSelectChannel() {

	personASend := make(chan string)
	personARecv := make(chan string)
	personBSend := make(chan string)
	personBRecv := make(chan string)

	chatSvr := chatServer{
		personASend, personARecv, personBSend, personBRecv,
	}

	chatSvr.start()

	chatCli := func(name string) {
		for i:= 1; i <= 30; i++ {
			randTime := rand.Intn(10)
			k,_ := strconv.ParseInt(strconv.Itoa(randTime * time.Millisecond), 10, 64)
			chatSvr.speck(fmt.Sprintf("person:%s, say:%d", name, randTime))
		}
	}

	
	go chatCli("A")
	go chatCli("B")
}

func TestChannel() {
	createChan()

	testChanCache()

	testChannelBlock()

	testSingleDirChannel()

}
