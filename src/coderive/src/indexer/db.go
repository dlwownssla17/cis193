package indexer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"coderive/src/crawler"
	"log"
	"coderive/src/common"
)

// SaveQueryTextSearch stores into the query text search collection the given query text search.
func SaveQueryTextSearch(q *QueryTextSearch) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesTextSearch := common.GetCollection(session, "queries.textsearch")

	err = collQueriesTextSearch.Insert(q)
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

	collRepositories := common.GetCollection(session, "repositories")

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

	collRepositories := common.GetCollection(session, "repositories")

	_, err = collRepositories.UpdateAll(bson.M{"processed": false}, bson.M{"$set": bson.M{"processed": true}})
	if err != nil {
		log.Fatal(err)
	}
}