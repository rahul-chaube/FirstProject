package main

import iris "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	second := app.Party("second")
	{
		second.Get("/", func(ctx iris.Context) {
			ctx.Write([]byte("Hello world ******** "))
		})
	}

	app.Listen(":9003")
}
