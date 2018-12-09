package repository_test

import (
	"testing"

	mgo "gopkg.in/mgo.v2"
)

func TestViewDataJapanShouldBeJAPAN(t *testing.T) {
	expected := Country{
		Name:   "JAPAN",
		IsEuro: false,
	}
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Fatal("database cannot connect")
	}

	repository := RepositoryMongo{
		Session: session,
	}

	country, err := repository.FindCountryByName("JAPAN")
	if err != nil {
		t.Fatal(err)
	}
	if country.Name != expected.Name {
		t.Fatalf("expected %v but it got %v", expected, country)
	}
}

func TestCreateCountryThaiShouldBeNoError(t *testing.T) {
	input := Country{
		Name:   "THAILAND",
		IsEuro: false,
	}
	session, _ := mgo.Dial("localhost:27017")

	repository := RepositoryMongo{
		Session: session,
	}
	id, err := repository.Insert(input)
	if err != nil {
		t.Fatal(err)
	}
	if id.Hex() == "" {
		t.Fatal("error insert")
	}

}
