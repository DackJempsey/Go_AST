package main
import (
"fmt"
"time"
)
func new_func() string {
	fmt.Println("new func running")
	return "This is new funcs return"
}
func main() {
	new_func()
	time.Sleep(10 * time.Second)
	fmt.Println("Hello from dynamically compiled code!")
}

