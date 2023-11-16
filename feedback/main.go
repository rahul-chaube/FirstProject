package main

import (
	"encoding/json"
	"errors"
	"feedback/model"
	"fmt"
	"log"
	"os"

	iris "github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	feedback := app.Party("feedback")
	{
		feedback.Get("/", listOfFile)
		feedback.Post("/", addFeedback)
	}

	app.Listen(":9004")
}

func listOfFile(ctx iris.Context) {
	log.Println("List of file is called ")
	entries, err := os.ReadDir("./feedback_dir")
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte("No File Exists " + err.Error()))
		return
	}
	data := []model.FileData{}
	for _, e := range entries {

		fmt.Println(e.Name())
		info, _ := e.Info()
		fmt.Println(e.Name(), info.Size(), info.ModTime())
		file := model.FileData{
			Name:         e.Name(),
			Size:         fmt.Sprint(info.Size()),
			ModifiedTime: info.ModTime().Format("2006-01-02 15:04:05"),
		}
		data = append(data, file)
	}
	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(200)
	resp, _ := json.Marshal(data)
	ctx.Write(resp)
}

func addFeedback(ctx iris.Context) {
	log.Println("Create File is called .... ")
	feedback := model.Feedback{}

	err := ctx.ReadBody(&feedback)
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte(err.Error()))
	}
	isExist := isFileExists(feedback.Title)
	if isExist {
		ctx.StatusCode(400)
		ctx.Write([]byte("File is already Exists "))
		return
	}
	err = os.WriteFile(fmt.Sprintf("feedback_dir/%s.txt", feedback.Title), []byte(feedback.Description), os.ModeAppend)
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte("Failed to write in file " + err.Error()))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(201)
	ctx.Write([]byte("File is created  "))

}

func isFileExists(fileName string) bool {
	_, err := os.Stat(fmt.Sprintf("feedback_dir/%s.txt", fileName))
	return !errors.Is(err, os.ErrNotExist)
}
