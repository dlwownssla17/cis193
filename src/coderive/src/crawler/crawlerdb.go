package crawler

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Dummy struct {
	Dummy bool
}

func getDummy() Dummy {
	return Dummy{
		Dummy: true,
	}
}

func getDatabase(session *mgo.Session) *mgo.Database {
	return session.DB("coderive")
}

// GetCollRepositories gets the repository collection.
func GetCollRepositories(session *mgo.Session) *mgo.Collection {
	return getDatabase(session).C("repositories")
}

// DBRepositoriesInit initializes the repository collection.
func DBRepositoriesInit() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := GetCollRepositories(session)

	dummy := getDummy()

	err = collRepositories.Insert(&dummy)
	if err != nil {
		log.Fatal(err)
	}

	err = collRepositories.Remove(bson.M{"dummy": true})
	if err != nil {
		log.Fatal(err)
	}
}

// DBRepositoryDrop drops the entire coderive database.
func DBDrop() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	db := getDatabase(session)

	err = db.DropDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

// SaveRepository stores into the repository collection the given crawled repository.
func SaveRepository(repo Repository) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := GetCollRepositories(session)

	err = collRepositories.Insert(&repo)
	if err != nil {
		log.Fatal(err)
	}
}