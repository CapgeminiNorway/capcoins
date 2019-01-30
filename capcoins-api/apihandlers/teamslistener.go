package apihandlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func TeamsListener(resp http.ResponseWriter, req *http.Request) {

	var respStatus int
	var respData []byte

	authenticated, reqBody := isRequestAuthenticated(req)
	fmt.Printf("authenticated: %v \n", authenticated)
	if !authenticated {
		//resp.WriteHeader(http.StatusForbidden)
		//resp.Write([]byte("Invalid Authentication"))
		respStatus = http.StatusForbidden
		respData = []byte("Invalid Authentication")
	}

	//reqBody, err := ioutil.ReadAll(req.Body)
	fmt.Printf("reqBody: %v \n", string(reqBody))
	/*if err != nil {
		//resp.WriteHeader(http.StatusInternalServerError)
		respStatus = http.StatusInternalServerError
		respData = []byte("")
	} */

	var teamsRequest = TeamsRequest{}
	err := json.Unmarshal(reqBody, &teamsRequest)
	//err := json.NewDecoder(strings.NewReader(string(reqBody))).Decode(&teamsRequest)
	if err != nil {
		//resp.WriteHeader(http.StatusBadRequest)
		//resp.Write([]byte(err.Error()))
		respStatus = http.StatusBadRequest
		respData = []byte(err.Error())
	}
	//fmt.Printf("teamsRequest: %v \n", teamsRequest)
	fmt.Printf("teamsRequest.text: %v \n", teamsRequest.Text)
	err = SaveTeamsRequest(teamsRequest)
	if err != nil {
		log.Fatal(err)
	}
	/*
teamsRequest: {message 1548765074233 2019-01-29T12:31:14.236Z 2019-01-29T13:31:14.236+01:00 https://smba.trafficmanager.net/emea/ msteams {29:14_tQdN2ZuAfTS1EheqkJ_v2sCNUNijEZuG7yw3rTvA9Tk3Jra3WybOctP0s3i_kGvO2vAoIGwzctN4slngHmMQ Guleryuz, Yilmaz} {19:a694a70b0a7f4bfd9572b2d912c5d160@thread.skype;messageid=1548765074233} { } plain <at>CapCoins</at> goooo
[{text/html <div><div><span itemscope="" itemtype="http://schema.skype.com/Mention" itemid="0">CapCoins</span> goooo</div>
</div>}] [map[type:clientInfo locale:nb-NO country:NO platform:Mac]] {19:a694a70b0a7f4bfd9572b2d912c5d160@thread.skype 19:f49d648a29104431a8b662c4182b8bf2@thread.skype}}
	*/

	// TODO: process requests by following bot-comm dialect

	// TODO: respond according to processed request
	teamsResponse := BuildTeamsResponse("Echo: " + teamsRequest.Text)
	err = SaveTeamResponse(teamsResponse)
	if err != nil {
		log.Fatal(err)
	}
	respData, err = json.Marshal(teamsResponse)
	if err == nil {
		respStatus = http.StatusOK
	} else {
		respStatus = http.StatusInternalServerError
	}
	ResponseWithJSON(resp, respData, respStatus)
	/*
	fmt.Println(respStatus)
	fmt.Println(string(respData))
	resp.Write(respData)
	*/
}

func isRequestAuthenticated(lreq *http.Request) (isAuth bool, reqBody []byte) {

	teamsSecret := os.Getenv("TEAMS_SECRET")
	teamsSecretBytes, _ := base64.StdEncoding.DecodeString(teamsSecret)

	reqBody, _ = ioutil.ReadAll(lreq.Body)
	authHeader := lreq.Header.Get("authorization")
	fmt.Printf("authHeader: %v \n", authHeader)
	messageMAC, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "HMAC "))

	mac := hmac.New(sha256.New, teamsSecretBytes)
	mac.Write(reqBody)
	expectedMAC := mac.Sum(nil)
	isAuth = hmac.Equal(messageMAC, expectedMAC)

	return
}

/*
reqBody: {
	"type":"message","id":"1548770278830","timestamp":"2019-01-29T13:57:58.834Z",
	"localTimestamp":"2019-01-29T14:57:58.834+01:00",
	"serviceUrl":"https://smba.trafficmanager.net/emea/",
	"channelId":"msteams",
	"from":{
		"id":"29:14_tQdN2ZuAfTS1EheqkJ_v2sCNUNijEZuG7yw3rTvA9Tk3Jra3WybOctP0s3i_kGvO2vAoIGwzctN4slngHmMQ",
		"name":"Guleryuz, Yilmaz",
	"aadObjectId":"87994953-d622-4c2c-a425-28e58c73b2b2"
	},
	"conversation":{
		"isGroup":true,"id":"19:a694a70b0a7f4bfd9572b2d912c5d160@thread.skype;messageid=1548770278830",
		"name":null,"conversationType":"channel"
	},
	"recipient":null,"textFormat":"plain","attachmentLayout":null,
	"membersAdded":[],"membersRemoved":[],"topicName":null,"historyDisclosed":null,"locale":null,
	"text":"<at>CapCoins</at>&nbsp;testing go go go <at><img alt=\"🍺\" itemid=\"beer\" itemscope=\"\" itemtype=\"http://schema.skype.com/Emoji\" src=\"https://statics.teams.microsoft.com/evergreen-assets/funstuff/skype-emoticons-f/beer/default_20.png\" style=\"width:20px; height:20px\"></at>&nbsp;&nbsp;\n","speak":null,"inputHint":null,"summary":null,"suggestedActions":null,"attachments":[{"contentType":"image/*","contentUrl":"https://statics.teams.microsoft.com/evergreen-assets/funstuff/skype-emoticons-f/beer/default_20.png","content":null,"name":null,"thumbnailUrl":null},{"contentType":"text/html","contentUrl":null,"content":"<div><div><span itemscope=\"\" itemtype=\"http://schema.skype.com/Mention\" itemid=\"0\">CapCoins</span>&nbsp;testing go go go <span class=\"animated-emoticon-20-beer\" title=\"Øl\" type=\"(beer)\"><img alt=\"🍺\" itemid=\"beer\" itemscope=\"\" itemtype=\"http://schema.skype.com/Emoji\" src=\"https://statics.teams.microsoft.com/evergreen-assets/funstuff/skype-emoticons-f/beer/default_20.png\" style=\"width:20px; height:20px\"></span>&nbsp;&nbsp;</div>\n</div>","name":null,"thumbnailUrl":null}],"entities":[{"type":"clientInfo","locale":"nb-NO","country":"NO","platform":"Mac"}],"channelData":{"teamsChannelId":"19:a694a70b0a7f4bfd9572b2d912c5d160@thread.skype","teamsTeamId":"19:f49d648a29104431a8b662c4182b8bf2@thread.skype","channel":{"id":"19:a694a70b0a7f4bfd9572b2d912c5d160@thread.skype"},"team":{"id":"19:f49d648a29104431a8b662c4182b8bf2@thread.skype"},"tenant":{"id":"76a2ae5a-9f00-4f6b-95ed-5d33d77c4d61"}},"action":null,"replyToId":null,"value":null,"name":null,"relatesTo":null,"code":null}


*/