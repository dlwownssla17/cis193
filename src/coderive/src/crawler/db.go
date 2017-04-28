package crawler

import (
	"coderive/src/common"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// ExistsRepository checks whether the specified repository already exists in the database.
func ExistsRepository(username, repositoryName string) bool {
	bsonMatch := make(bson.M)
	bsonMatch["username"] = username
	bsonMatch["name"] = repositoryName

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := common.GetCollection(session, "repositories")

	count, err := collRepositories.Find(bsonMatch).Count()
	if err != nil {
		panic(err)
	}
	return count > 0
}

// SaveRepository stores into the repository collection the given crawled repository.
func SaveRepository(repo *Repository) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := common.GetCollection(session, "repositories")

	err = collRepositories.Insert(repo)
	if err != nil {
		log.Fatal(err)
	}
}
