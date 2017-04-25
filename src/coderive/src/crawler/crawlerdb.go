package crawler

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type Dummy struct {
	Dummy bool
}

func getDatabase(session *mgo.Session) *mgo.Database {
	return session.DB("coderive")
}

func getCollRepository(session *mgo.Session) *mgo.Collection {
	return getDatabase(session).C("repository")
}

func getDummy() Dummy {
	return Dummy{
		Dummy: true,
	}
}

// DBRepositoryInit initializes the repository collection.
func DBRepositoryInit() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepository := getCollRepository(session)

	dummy := getDummy()

	err = collRepository.Insert(&dummy)
	if err != nil {
		log.Fatal(err)
	}

	err = collRepository.Remove(bson.M{"dummy": true})
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

	collRepository := getCollRepository(session)

	err = collRepository.Insert(&repo)
	if err != nil {
		log.Fatal(err)
	}
}