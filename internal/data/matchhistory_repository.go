package data

import (
	"context"

	"github.com/DevGrohl/dota-tracker-tui/pkg/matchhistory"
)

type MatchHistoryRepository struct {
	Data *Data
}

func (mhr *MatchHistoryRepository) Create(mh *matchhistory.MatchHistory) error {
	query := "INSERT INTO matches (match_id, start_time, lobby_type) VALUES ($1, $2, $3) RETURNING match_id"

	row := mhr.Data.DB.QueryRow(
		query, mh.MatchID, mh.StartTime, mh.LobbyType,
	)

	err := row.Scan(&mh.MatchID)
	if err != nil {
		return err
	}

	return nil
}

func (mhr *MatchHistoryRepository) GetAll(ctx context.Context) ([]matchhistory.MatchHistory, error) {
	query := "SELECT * FROM match_history"

	rows, err := mhr.Data.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []matchhistory.MatchHistory

	for rows.Next() {
		var mh matchhistory.MatchHistory
		rows.Scan(&mh.MatchID, &mh.StartTime, &mh.LobbyType)
		matches = append(matches, mh)
	}

	return matches, nil
}
