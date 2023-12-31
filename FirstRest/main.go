package main

import (
	"fmt"
	"math/rand"

	iris "github.com/kataras/iris/v12"
)

func main() {
	fmt.Println("Staring FirstRest Module ")

	app := iris.New()

	user := app.Party("/user")
	{
		user.Get("/", list)
		user.Post("/", create)
	}
	app.Listen(":9002")
}
func list(ctx iris.Context) {
	ctx.Write([]byte("Hello world " + fmt.Sprint(generateNumber())))

}
func create(ctx iris.Context) {
	ctx.Write([]byte("Post Request is called "))

}

func generateNumber() int {
	return rand.Intn(100)
}
