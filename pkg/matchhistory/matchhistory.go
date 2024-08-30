package matchhistory

import "time"

type MatchHistory struct {
	MatchID   int       `json:"match_id,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	LobbyType int       `json:"lobby_type,omitempty"`
}
