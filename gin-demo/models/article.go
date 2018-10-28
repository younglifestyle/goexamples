package models

import "github.com/globalsign/mgo/bson"

const (
	// CollectionArticle holds the name of the articles collection
	CollectionArticle = "articles"
	TableArticle      = "articles"
)

// Article model
type Article struct {
	Id        bson.ObjectId `json:"_id,omitempty" gorm:"id" bson:"_id,omitempty"`
	Title     string        `json:"title" form:"title" binding:"required" gorm:"title" bson:"title"`
	Body      string        `json:"body" form:"body" binding:"required" gorm:"body" bson:"body"`
	CreatedOn int64         `json:"created_on" gorm:"created_on" bson:"created_on"`
	UpdatedOn int64         `json:"updated_on" gorm:"updated_on" bson:"updated_on"`
	// User      bson.ObjectId `json:"user"`
}
