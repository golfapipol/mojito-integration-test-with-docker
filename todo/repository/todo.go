package repository

import (
	"mojito/todo/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	FindCountryByName(name string) (model.Country, error)
}
type RepositoryMongo struct {
	Session *mgo.Session
}

func (repo RepositoryMongo) FindCountryByName(name string) (model.Country, error) {
	var country model.Country
	repo.Session.DB("mojito").
		C("countries_en").
		Find(bson.M{"country_name": name}).
		One(&country)
	return country, nil
}
func (repo RepositoryMongo) Insert(country model.Country) (bson.ObjectId, error) {
	var id bson.ObjectId
	id = bson.NewObjectId()
	country.ID = id
	err := repo.Session.DB("mojito").C("countries_en").Insert(country)
	return id, err
}
