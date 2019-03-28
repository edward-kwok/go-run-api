package mongo_test

import (
	"github.com/edward-kwok/go-run-api/pkg/mongo"
	"log"
	"testing"
)

const (
	mongoUrl           = "localhost:27017"
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(){
	session, err := mongo.NewSession("localhost:27017")
	if err != nil {
		log.Fatalf("Cannot connect to mongo %s", err)
	}
	defer session.Close()
	mongo.NewRunnerService((session.Copy(), "test","runner")


}
