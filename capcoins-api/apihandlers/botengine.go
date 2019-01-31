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
	fileName := "./config/bot-dialect.json"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Successfully opened %v file \n", fileName)

		//var commands Commands
		fileBlob, _ := ioutil.ReadAll(file)
		err = json.Unmarshal(fileBlob, &commands)
		if err != nil {
			fmt.Println(err)
		}

		for ind := range commands.Commands {
		//for ind:=0; ind<len(commands.Commands); ind++ {
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
// make use of html templates

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