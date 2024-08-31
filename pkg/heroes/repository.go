package heroes

import "context"

type Repository interface {
	GetAll(ctx context.Context) ([]Heroes, error)
	GetOne(ctx context.Context, id int) (Heroes, error)
	Create(ctx context.Context, mh *Heroes) error
	Delete(ctx context.Context, id int) error
}
