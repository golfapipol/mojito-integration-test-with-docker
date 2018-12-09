package model

import "gopkg.in/mgo.v2/bson"

type Country struct {
	ID     bson.ObjectId `json:"-" bson:"_id"`
	Name   string        `json:"name" bson:"country_name"`
	IsEuro bool          `json:"is_euro" bson:"schengen_flag"`
}
