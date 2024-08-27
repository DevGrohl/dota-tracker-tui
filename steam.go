package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const (
	steamURL = "https://api.steampowered.com/"
	appID    = "570"
)

var (
	steamKey  string
	accountID string
)

type MatchData struct {
	Result struct {
		Status       int `json:"status"`
		TotalResults int `json:"total_results"`
		Matches      []struct {
			MatchID   int `json:"match_id"`
			StartTime int `json:"start_time"`
			LobbyType int `json:"lobby_type"`
			Players   []struct {
				AccountId int `json:"account_id"`
				HeroId    int `json:"hero_id"`
			}
		}
	}
}

type optionalMatchHistory struct {
	start_at_match_id string
	matches_requested string
}

// load env
func loadEnv() {
	steamKey = os.Getenv("STEAM_KEY")
	accountID = os.Getenv("ACCOUNT_ID")
}

func getMatchHistory(opts *optionalMatchHistory) MatchData {
	loadEnv()
	accessPoint := "IDOTA2Match_570/GetMatchHistory/v1/?&key=" + steamKey + "&account_id=" + accountID

	if (&opts) != nil {
		if opts.start_at_match_id != "" {
			accessPoint += "&start_at_match_id=" + opts.start_at_match_id
		}
		if opts.matches_requested != "" {
			accessPoint += "&matches_requested=" + opts.matches_requested
		}
	}

	// HTTP GET
	resp, err := http.Get(steamURL + accessPoint)
	// fmt.Println("Response from Get: ", resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	// fmt.Println("Response body: ", string(body))
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON
	matchHistory := MatchData{}
	err = json.Unmarshal(body, &matchHistory)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", matchHistory)

	return matchHistory
}
