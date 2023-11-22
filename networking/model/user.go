package model

type User struct {
	Name string `bson:"name" json:"name"`
	Id   string `bson:"_id" json:"id"`
	Age  string `bson:"age" json:"age"`
}
