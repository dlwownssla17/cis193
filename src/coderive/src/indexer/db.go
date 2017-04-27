package indexer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"coderive/src/crawler"
	"log"
	"coderive/src/common"
	"coderive/src/tokenizer"
	"regexp"
	"fmt"
)

/* * */

// SaveQueryTextSearch stores into the query text search collection the given query.
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

func FindQueryTextSearches(qMap *tokenizer.QueryMap) []QueryTextSearch {
	bsonMatch := make(bson.M)
	for qType, qVal := range *qMap {
		switch qType {
		case "lines":
			linesMap, ok := qVal.(map[string]int)
			if !ok || len(linesMap) != 1 {
				return nil
			}

			for linesOp, linesThreshold := range linesMap {
				bsonMatch["fileinfo.numlines"] = bson.M{fmt.Sprintf("$%s", linesOp): linesThreshold}
			}

		case "text":
			textMap, ok := qVal.(map[string]interface{})
			if !ok {
				return nil
			}

			qTextRegex, ok := textMap["val"].(string)
			if !ok {
				return nil
			}

			isRegex, ok := textMap["regex"].(bool)
			if !ok {
				return nil
			}

			if !isRegex {
				qTextRegex = regexp.QuoteMeta(qTextRegex)
			}

			bsonMatch["fileinfo.formatteddata"] = bson.M{"$regex": qTextRegex, "$options": "si"}

		default:
			return nil
		}
	}

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesTextSearch := common.GetCollection(session, "queries.textsearch")

	var results []QueryTextSearch
	err = collQueriesTextSearch.Find(bsonMatch).Sort("fileinfo.repositoryname fileinfo.filepath").All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

/* * */

// SaveQueryTextWordMatch stores into the query text word match collection the given query.
func SaveQueryTextWordMatch(q *QueryTextWordMatch) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesTextWordMatch := common.GetCollection(session, "queries.textwordmatch")

	err = collQueriesTextWordMatch.Insert(q)
	if err != nil {
		log.Fatal(err)
	}
}

/* * */

// SaveQuerySemVarType stores into the query sem var type collection the given query.
func SaveQuerySemVarType(q *QuerySemVarType) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collQueriesSemVarType := common.GetCollection(session, "queries.semvartype")

	err = collQueriesSemVarType.Insert(q)
	if err != nil {
		log.Fatal(err)
	}
}

/* * */

// GetAllRepositoriesToProcess gets all the repositories to process for indexing.
func GetAllRepositoriesToProcess() []crawler.Repository {
	bsonMatch := make(bson.M)
	bsonMatch["processed"] = false

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := common.GetCollection(session, "repositories")

	var results []crawler.Repository
	err = collRepositories.Find(bsonMatch).All(&results)
	if err != nil {
		log.Fatal(err)
	}

	return results
}

// UpdateAllRepositoriesProcessed updates the processed field to true for all repositories whose processed field was
// originally false.
func UpdateAllRepositoriesProcessed() {
	bsonMatch, bsonSet := make(bson.M), make(bson.M)
	bsonMatch["processed"] = false
	bsonSet["$set"] = bson.M{"processed": true}

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := common.GetCollection(session, "repositories")

	_, err = collRepositories.UpdateAll(bsonMatch, bsonSet)
	if err != nil {
		log.Fatal(err)
	}
}