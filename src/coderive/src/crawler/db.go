package crawler

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"coderive/src/common"
)

// ExistsRepository checks whether the specified repository already exists in the database.
func ExistsRepository(username, repositoryName string) bool {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collRepositories := common.GetCollection(session, "repositories")

	count, err := collRepositories.Find(bson.M{"username": username, "repositoryname": repositoryName}).Count()
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