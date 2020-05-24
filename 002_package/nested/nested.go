package main

import (
	"github.com/KapaMiki/GolangLearning/002_package/nested/say"
	"github.com/KapaMiki/GolangLearning/002_package/nested/hello"
	"fmt"
)


func main() {
	fmt.Println("Start program...")
	fmt.Println(say.CallFromSay() + hello.CallFromHello())
}