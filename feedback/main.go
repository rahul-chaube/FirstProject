package main

import (
	model "feedback/model"
	"log"

	iris "github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	feedback := app.Party("feedback")
	{
		feedback.Get("/", listOfFile)
		feedback.Post("/", addFeedback)
	}

	app.Listen("9004")
}

func listOfFile(ctx iris.Context) {
	log.Println("List of file is called ")
}

func addFeedback(ctx iris.Context) {
	log.Println("Create File is called .... ")
	feedback := model.Feedback{}

}
