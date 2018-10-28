package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	User struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name     string        `json:"name"`
		LastName string        `json:"lastname"`
	}

	Person struct {
		Name  string
		Phone string
	}
)

type Repository struct {
	C *mgo.Collection
}

func (r *Repository) Create(docs ...interface{}) error {
	return r.C.Insert(docs)
}

func (r *Repository) GetAll() []interface{} {
	var users []interface{}
	iter := r.C.Find(nil).Iter()

	var result interface{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *Repository) Delete(id string) error {
	return r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	r := Repository{
		C: session.DB("test").C("people"),
	}

	r.C.Insert(&Person{"Heln", "+55 53 8116 9639"},
		&Person{"Ana", "+51 53 8116 9639"},
		&Person{"A3a", "+51 53 8116 9639"},
		&Person{"Awa", "+51 53 8116 9623"})

	r.Create(&User{
		Name:     "test",
		LastName: "unknow",
	})

	fmt.Println(r.GetAll())

	result := Person{}
	err = r.C.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("result :", result)
}
