package matchhistory

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]MatchHistory, error)
	GetOne(ctx context.Context, id int) (MatchHistory, error)
	GetByAccountID(ctx context.Context, accountID int) ([]MatchHistory, error)
	Create(ctx context.Context, mh *MatchHistory) error
	Delete(ctx context.Context, id int) error
}
