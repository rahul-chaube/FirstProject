package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"networking/constant"
	"networking/model"
	"time"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	app := iris.New()

	favorite := app.Party("favorite")
	{
		favorite.Get("/", getFavList)
		favorite.Post("/", addFavorite)

	}
	app.Get("/people", getPeopleList)
	app.Get("/movies", getMoviesList)
	// app.Listen(":"+os.Getenv("PORT"))
	app.Listen(":9006")
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017").SetConnectTimeout(30*time.Second))

	if err != nil {
		panic(err)
	}
	fmt.Println("client connection successful")
}

func addFavorite(ctx iris.Context) {

	user := model.User{}

	err := ctx.ReadBody(&user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Write([]byte("Error occurred " + err.Error()))
		return
	}
	collection := client.Database(constant.UserDatabase).Collection(constant.UserTable)
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Write([]byte("Error occurred " + err.Error()))
		return
	}
	ctx.StatusCode(http.StatusOK)
	user.Id = fmt.Sprint(res.InsertedID)
	resp, _ := json.Marshal(user)
	ctx.Header("Content-Type", "application/json")
	ctx.Write(resp)
}

func getFavList(ctx iris.Context) {

	collection := client.Database(constant.UserDatabase).Collection(constant.UserTable)
	courser, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Write([]byte("Error occurred " + err.Error()))
		return
	}
	var users []model.User

	if err := courser.All(ctx, &users); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Write([]byte("Error occurred " + err.Error()))
		return
	}
	ctx.StatusCode(http.StatusOK)
	ctx.Header("Content-Type", "application/json")

	resp, _ := json.Marshal(users)
	ctx.Write(resp)
}

func getMoviesList(ctx iris.Context) {
	req, err := http.NewRequest("GET", "https://swapi.dev/api/films", nil)
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

func getPeopleList(ctx iris.Context) {
	req, err := http.NewRequest("GET", "https://swapi.dev/api/people", nil)
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
