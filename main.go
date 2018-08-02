package main

import (
	"time"
	"demo/testcrypto"
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
	//language.TestChannel(5)
	testcrypto.TestCrypto()
	//i := make(chan int)
	//log.Println(<- i)
	//log.Println("------------------end---------------------")
	time.Sleep(50 * time.Second)
}
