package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hwllo world")
	// resitng 100 miliseconds
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("sayHello function ended")
}

func sayHi() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("hi world")
}

func main() {
	// 1. first execute main
	// 2. execute line by line
	// 3. with the help of go switching to other function
	go sayHello()
	go sayHi()

	// witing this till this time, if runned then good, if not then also good
	time.Sleep(3000 * time.Millisecond)
}
