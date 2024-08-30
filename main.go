package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DevGrohl/dota-tracker-tui/internal/data"
	"github.com/DevGrohl/dota-tracker-tui/pkg/matchhistory"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

func main() {
	// Get match history
	matchHistory := getMatchHistory(&optionalMatchHistory{matches_requested: "1"})
	fmt.Printf("%+v\n", matchHistory.Result.Matches[0])

	// Access to data from defined user
	getMatchDetails(matchHistory.Result.Matches[0].MatchID)

	// Store relevant information
	db := data.New()
	if err := db.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	var mh matchhistory.MatchHistory
	mh.MatchID = matchHistory.Result.Matches[0].MatchID
	mh.StartTime = time.Unix(int64(matchHistory.Result.Matches[0].StartTime), 0) // matchHistory.Result.Matches[0].StartTime
	mh.LobbyType = matchHistory.Result.Matches[0].LobbyType

	repo := data.MatchHistoryRepository{Data: db}
	err := repo.Create(&mh)
	if err != nil {
		log.Fatal(err)
	}

	data.Close()

	// Display data in TUI
}
