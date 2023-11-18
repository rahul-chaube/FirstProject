package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	movies := app.Party("movies")
	{
		movies.Get("/list", movieList)
	}
	app.Listen(":" + os.Getenv("PORT"))
	// app.Listen(":9005")
}

func movieList(ctx iris.Context) {
	//https://api.publicapis.org/entries

	req, err := http.NewRequest("GET", "https://api.publicapis.org/entries", nil)
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte("Error while preparing request " + err.Error()))
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte("Error while requesting " + err.Error()))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.StatusCode(400)
		ctx.Write([]byte("Error while reading response " + err.Error()))
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.StatusCode(200)
	ctx.Write(body)

}
