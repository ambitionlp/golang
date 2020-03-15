package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan int)
	// go sendData(ch1)
	// /*
	// 	子goroutine，写出数据10个
	// 			每写一个，阻塞一次，主程序读取一次，解除阻塞

	// 	主goroutine：循环读
	// 			每次读取一个，堵塞一次，子程序，写出一个，解除阻塞

	// 	发送发，关闭通道的--->接收方，接收到的数据是该类型的零值，以及false
	// */
	// //主程序中获取通道的数据
	// for {
	// 	// time.Sleep(1 * time.Second)
	// 	v, ok := <-ch1 //其他goroutine，显示的调用close方法关闭通道。
	// 	if !ok {
	// 		fmt.Println("已经读取了所有的数据，", ok)
	// 		break
	// 	}
	// 	fmt.Println("取出数据：", v, ok)
	// }

	// fmt.Println("main...over....")
	// ch1 := make(chan int)
	// go sendData(ch1)
	// for i := range ch1 {
	// 	fmt.Println(i)
	// }

	//缓存通道
	// ch2 := make(chan string, 6)
	//ch2 := make(chan <- int) // 单向，只写，不能读
	//ch3 := make(<- chan int) //单向，只读，不能写
	// go sendString(ch2)
	// for value := range ch2 {
	// 	fmt.Println(value)
	// }
	// fmt.Println("main...over")

	// 计时器
	// timer := time.NewTimer(5 * time.Second)
	// //新开启一个线程来处理触发后的事件
	// go func() {
	// 	//等触发时的信号
	// 	<-timer.C
	// 	fmt.Println("Timer  结束。。")
	// }()
	// //由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器
	// time.Sleep(3 * time.Second)
	// stop := timer.Stop()
	// if stop {
	// 	fmt.Println("Timer  停止。。")
	// }

	//after()
	/*
		func After(d Duration) <-chan Time
			返回一个通道：chan，存储的是d时间间隔后的当前时间。
	*/
	// ch1 := time.After(3 * time.Second) //3s后
	// fmt.Printf("%T\n", ch1)            // <-chan time.Time
	// fmt.Println(time.Now())            //2019-08-15 09:56:41.529883 +0800 CST m=+0.000465158
	// time2 := <-ch1
	// fmt.Println(time2) //2019-08-15 09:56:44.532047 +0800 CST m=+3.002662179

	//select语句
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("获取 ch1 数据", num1)
	case num2 := <-ch2:
		fmt.Println("获取 ch2 数据", num2)
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	}

	fmt.Println("main...over")
}

// func sendData(ch1 chan int) {
// 	// 发送方：10条数据
// 	for i := 0; i < 10; i++ {
// 		ch1 <- i //将i写入通道中
// 	}
// 	fmt.Println("main...over")
// 	close(ch1) //将ch1通道关闭了。
// }
func sendString(ch2 chan string) {
	for i := 0; i < 6; i++ {
		ch2 <- "数据"
		fmt.Println("子程序写入", i, "数据")
	}
	close(ch2)
}
