package apihandlers

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var dbName = "capcoins"
var collectionName = ""
func DbConnect() (session *mgo.Session) {
	connectURL := os.Getenv("MONGODB_URL") //"localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return
}

func SaveTeamsRequest(teamsReq TeamsRequest) (err error) {

	// init DB connection
	dbSession := DbConnect()
	defer dbSession.Close()

	/*if len(teamsReq._id) == 0 {
		teamsReq._id = teamsReq.ID
		teamsReq.ID = ""
	}*/

	collectionName = "requests"
	collection := dbSession.DB(dbName).C(collectionName)
	err = collection.Insert(teamsReq)
	if err != nil {
		log.Fatal(err)
	}
	return
}