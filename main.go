package main

import (
	"time"
	"./makeandnew"
)

func main() {

	//httpsever.StartHttpServer()
	//language.TestPanic()
	//language.TestDefer()

	// := redismanager.GetInstance()
	//value := rMgr.Get("jiang")
	//fmt.Println(value)

	//testNoDistributedLock()

	//testDistributedLock()
	makeandnew.TestMakeAndNew()
	//i := make(chan int)
	//log.Println(<- i)
	//log.Println("------------------end---------------------")
	time.Sleep(50 * time.Second)
}
