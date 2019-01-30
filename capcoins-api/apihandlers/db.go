package apihandlers

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"time"
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
	if len(teamsReq.CreatedAt) == 0 {
		teamsReq.CreatedAt = jodaTime.Format("YYYY-MM-ddTHH:mm:ss", time.Now())
	}

	collectionName = "teams_request"
	collection := dbSession.DB(dbName).C(collectionName)
	err = collection.Insert(teamsReq)
	if err != nil {
		log.Fatal(err)
	}
	return
}
func SaveTeamResponse(teamResp TeamsResponse) (err error) {
	// init DB connection
	dbSession := DbConnect()
	defer dbSession.Close()

	if len(teamResp.CreatedAt) == 0 {
		teamResp.CreatedAt = jodaTime.Format("YYYY-MM-ddTHH:mm:ss", time.Now())
	}

	collectionName = "teams_response"
	collection := dbSession.DB(dbName).C(collectionName)
	err = collection.Insert(teamResp)
	if err != nil {
		log.Fatal(err)
	}
	return
}