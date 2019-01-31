package main

import (
	"./apihandlers"
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Printf("starting CapCoins-API..! on port: %v \n", port())

	// try to get teamsSecret from env-vars
	teamsSecret := os.Getenv("TEAMS_SECRET")
	//fmt.Printf("token from ENV: %v \n", vimeoToken)
	teamsSecret = strings.TrimSpace(teamsSecret)

	if len(teamsSecret) == 0 || len(teamsSecret) < 10 {
		// if not found then ask the user to enter it manually
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your security token for MS Teams: ")
		teamsSecret, _ := reader.ReadString('\n')
		teamsSecret = strings.TrimSpace(teamsSecret)
		if len(teamsSecret) == 0 || len(teamsSecret) < 20 {
			apihandlers.DisplayTokenWarning()
			panic("!!!missing teamsSecret!!!")
		}
		os.Setenv("TEAMS_SECRET", teamsSecret)
	}

	//apihandlers.LoadBotDialect()

	// init DB connection
	//dbSession := apihandlers.DbConnect()
	//defer dbSession.Close()

	// using Gorilla mux, http://www.gorillatoolkit.org/pkg/mux#overview
	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/", index)
	// dummy testing
	muxRouter.HandleFunc("/echo", apihandlers.ExampleEchoFunc)
	muxRouter.HandleFunc("/hello", apihandlers.ExampleHelloFunc)

	// --- CapCoins endpoints ---
	// MS Teams interacts with this... bot-comm
	muxRouter.HandleFunc("/teams", apihandlers.TeamsListener)
	// WebApp interacts with this... api
	// muxRouter.HandleFunc("/api", apihandlers.ApiListener)

	// Negroni for middleware funcs, see https://github.com/urfave/negroni
	negroniDefaults := negroni.Classic() // include default middleware
	/*
	negroniCustom := negroni.New()
	logger := negroni.NewLogger()
	logger.SetFormat("{{.StartTime}} | [{{.Status}} - {{.Duration}}] | {{.Hostname}} | {{.Method}} {{.Path}} | {{.TeamsRequest.UserAgent}}")
	negroniCustom.Use(logger)
	*/
	negroniDefaults.UseHandler(muxRouter)

	err := http.ListenAndServe(port(), negroniDefaults)
	checkError(err)

}


func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8088"
	}
	return ":" + port
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to CapCoins API crafted with Golang!")
}

