package apihandlers

import (
	"fmt"
	"net/http"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

/*
// WebHook represents the interface needed to handle Microsoft Teams WebHook Requests.
type WebHook interface {
	OnMessage(TeamsRequest) (TeamsResponse, error)
}
*/

// TeamsRequest data representing an inbound WebHook request from Microsoft Teams.
type TeamsRequest struct {
	CreatedAt string  `json:",omitempty"`

	Type           string    `json:"type,omitempty"`
	ID             string    `json:"id,omitempty"`
	Timestamp      string    `json:"timestamp,omitempty"`
	LocalTimestamp string    `json:"localTimestamp,omitempty"`
	ServiceURL     string    `json:"serviceUrl,omitempty"`
	ChannelID      string    `json:"channelId,omitempty"`
	FromUser       TeamsUser `json:"from,omitempty"`
	Conversation   struct {
		ID string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"conversation"`
	RecipientUser TeamsUser `json:"recipient,omitempty"`
	TextFormat    string    `json:"textFormat,omitempty"`
	Text          string    `json:"text,omitempty"`
	Attachments   []struct {
		ContentType string `json:"contentType,omitempty"`
		Content     string `json:"Content,omitempty"`
	} `json:"attachments"`
	Entities    []interface{} `json:"entities,omitempty"`
	ChannelData struct {
		TeamsChannelID string `json:"teamsChannelId,omitempty"`
		TeamsTeamID    string `json:"teamsTeamId,omitempty"`
	}
}

// TeamsResponse represents the data to return to Microsoft Teams.
type TeamsResponse struct {
	CreatedAt string  `json:",omitempty"`

	Type          string    `json:"type"`
	Text          string    `json:"text"`
	ReplyToId     string    `json:"replyToId,omitempty"`
	FromUser      TeamsUser `json:"from,omitempty"`
	RecipientUser TeamsUser `json:"recipient,omitempty"`
}
// BuildTeamsResponse is a helper method to build a TeamsResponse
func BuildTeamsResponse(teamsRequest TeamsRequest, fields ...string) (teamsResponse TeamsResponse) {
	teamsResponse = TeamsResponse{}
	teamsResponse.Type = "message"
	teamsResponse.Text = teamsRequest.Text

	if fields != nil {
		teamsResponse.RecipientUser = teamsRequest.FromUser
		teamsResponse.ReplyToId = teamsRequest.ID
	}

	return
}

// TeamsUser represents data for a Microsoft Teams user.
type TeamsUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func DisplayTokenWarning() {
	//fmt.Println("!!!securityToken is NOT valid!!!")
	fmt.Println("!!!you MUST provide securityToken to be able to use MS Teams API!!!")
	fmt.Println("get it from https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/outgoingwebhook ")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
