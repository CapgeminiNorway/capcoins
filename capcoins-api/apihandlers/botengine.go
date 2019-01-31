package apihandlers

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"strings"
)

var commands Commands
type Commands struct {
	Commands []Command `json:"commands,omitempty"`
}
type Command struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Guide []string `json:"guide,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
	Response string `json:"response,omitempty"`
}

func LoadBotDialect() (commands Commands) {
	dialectFile := "./config/bot-dialect.json"
	jsonFile, err := os.Open(dialectFile)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Successfully opened %v file \n", dialectFile)

		//var commands Commands
		jsonBlob, _ := ioutil.ReadAll(jsonFile)
		err = json.Unmarshal(jsonBlob, &commands)
		if err != nil {
			fmt.Println(err)
		}

		for ind := range commands.Commands {
		//for i:=0; i<len(commands.Commands); i++ {
			commandJson, _ := json.Marshal(commands.Commands[ind])
			fmt.Printf("command #%d: %v \n", ind, string(commandJson))
		}
	}

	return
}

func ProcessTeamsRequest(teamsRequest TeamsRequest) (teamsResponse TeamsResponse) {

	if len(commands.Commands) == 0 {
		commands = LoadBotDialect()
	}
// TODO: implement business logic for responses

	for ind := range commands.Commands {
		//strings.Contains(commands.Commands[ind].Keywords, teamsRequest.Text)
		for ind2 := range commands.Commands[ind].Keywords {
			if strings.Contains(html.UnescapeString(teamsRequest.Text), commands.Commands[ind].Keywords[ind2]) {
				teamsResponse.Text = fmt.Sprintf(commands.Commands[ind].Response, "TBD")
				teamsResponse.Type = "message"
				teamsResponse.ReplyToId = teamsRequest.ID
				goto foundit
			}
		}
	}

	foundit : {}

	return
}