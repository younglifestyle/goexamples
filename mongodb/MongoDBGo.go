package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// This is a type to hold our word definitions in
// we specifiy both bson (for MongoDB) and json (for web)
// naming for marshalling and unmarshalling
type item struct {
	ID         objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Word       string            `bson:"word" json:"word"`
	Definition string            `bson:"definition" json:"definition"`
}

var client *mongo.Client

//func wordHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "GET":
//		c := client.Database("grand_tour").Collection("words")
//
//		sort, err := mongo.Opt.Sort(bson.NewDocument(bson.EC.Int32("word", 1)))
//
//		if err != nil {
//			log.Fatal("Sort error ", err)
//		}
//
//		cur, err := c.Find(nil, nil, sort)
//
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		defer cur.Close(context.Background())
//
//		var items []item
//
//		for cur.Next(nil) {
//			item := item{}
//			err := cur.Decode(&item)
//			if err != nil {
//				log.Fatal("Decode error ", err)
//			}
//			items = append(items, item)
//		}
//
//		if err := cur.Err(); err != nil {
//			log.Fatal("Cursor error ", err)
//		}
//
//		jsonstr, err := json.Marshal(items)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(jsonstr)
//		return
//	case "PUT":
//		r.ParseForm()
//		c := client.Database("grand_tour").Collection("words")
//		// newItem := item{ID: objectid.New(), Word: r.Form.Get("word"), Definition: r.Form.Get("definition")}
//		// _, err := c.InsertOne(nil, newItem)
//		newItemDoc := bson.NewDocument(bson.EC.String("word", r.Form.Get("word")), bson.EC.String("definition", r.Form.Get("definition")))
//		_, err := c.InsertOne(nil, newItemDoc)
//		if err != nil {
//			log.Println(err)
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		w.WriteHeader(http.StatusAccepted)
//		return
//	}
//
//	return
//}

func main() {
	var err error
	client, err = mongo.NewClient("mongodb://@localhost:27017/test")
	if err != nil {
		log.Println("conn :", err)
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	coll := client.Database("test").Collection("people")
	cursor, err := coll.Find(
		context.Background(),
		bson.NewDocument(bson.EC.String("name", "Ale")),
	)
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	for cursor.Next(context.Background()) {
		fmt.Println("ID :", cursor.ID())
	}

	res, err := coll.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	fmt.Println("IDs :", id)

	cur, err := coll.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		// do something with elem....
		fmt.Println("document :", elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
