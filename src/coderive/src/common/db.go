package common

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

// Dummy struct
type Dummy struct {
	Dummy bool
}

// GetDummy gets a dummy for initializing a database.
func GetDummy() *Dummy {
	return &Dummy{
		Dummy: true,
	}
}

// GetDatabase gets the database given the session.
func GetDatabase(session *mgo.Session) *mgo.Database {
	return session.DB("coderive")
}

// GetCollection gets the specified collection.
func GetCollection(session *mgo.Session, collectionName string) *mgo.Collection {
	return GetDatabase(session).C(collectionName)
}

// DBCollectionInit initializes the specified collection.
func DBCollectionInit(collectionName string) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := GetCollection(session, collectionName)

	dummy := GetDummy()

	err = collRepositories.Insert(dummy)
	if err != nil {
		log.Fatal(err)
	}

	err = collRepositories.Remove(bson.M{"dummy": true})
	if err != nil {
		log.Fatal(err)
	}
}

// DBDrop drops the entire coderive database.
func DBDrop() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := GetDatabase(session)

	err = db.DropDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
