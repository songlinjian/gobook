package main

import (
	"time"
	"fmt"
)

func ready(w string, sec int) {
	time.Sleep(int64(sec) * 1e9)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(5 * 1e9)
}
