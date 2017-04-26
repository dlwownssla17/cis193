package indexer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"coderive/src/crawler"
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

// GetCollQueriesTextSearch gets the query text search collection.
func GetCollQueriesTextSearch(session *mgo.Session) *mgo.Collection {
	return getDatabase(session).C("queries.textsearch")
}

// DBQueriesTextSearchInit initializes the query text search collection.
func DBQueriesTextSearchInit() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesTextSearch := GetCollQueriesTextSearch(session)

	dummy := getDummy()

	err = collQueriesTextSearch.Insert(&dummy)
	if err != nil {
		log.Fatal(err)
	}

	err = collQueriesTextSearch.Remove(bson.M{"dummy": true})
	if err != nil {
		log.Fatal(err)
	}
}

// SaveQueryTextSearch stores into the query text search collection the given query text search.
func SaveQueryTextSearch(q QueryTextSearch) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesTextSearch := GetCollQueriesTextSearch(session)

	err = collQueriesTextSearch.Insert(&q)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllRepositoriesToProcess gets all the repositories to process for indexing.
func GetAllRepositoriesToProcess() []crawler.Repository {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := crawler.GetCollRepositories(session)

	var results []crawler.Repository
	err = collRepositories.Find(bson.M{"processed": false}).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

// UpdateAllRepositoriesProcessed updates the processed field to true for all repositories whose processed field was
// originally false.
func UpdateAllRepositoriesProcessed() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := crawler.GetCollRepositories(session)

	_, err = collRepositories.UpdateAll(bson.M{"processed": false}, bson.M{"$set": bson.M{"processed": true}})
	if err != nil {
		log.Fatal(err)
	}
}