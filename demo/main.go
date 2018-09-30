package main

import (
	"demo/router"
	// "demo/controller"
	"demo/database"
	// "fmt"
)

var index int32 = 1

func main() {
	database.InitDB()
	router.Init()

	// for i := 0; i < 100000; i++ {
	// 	test()
	// 	fmt.Println(index)
	// }

}

// func test() {
// 	index = index + 1
// 	user := &controller.User{PassWord: index}
// 	user.Save()

// }
