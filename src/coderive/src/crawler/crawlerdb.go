package crawler

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

type Dummy struct {
	Dummy bool
}

func getCollRepository(session *mgo.Session) *mgo.Collection {
	return session.DB("coderive").C("repository")
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