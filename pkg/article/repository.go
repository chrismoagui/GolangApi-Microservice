package article

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]ArticleGame, error)
	GetOne(ctx context.Context, id uint) (ArticleGame, error)
	GetByUser(ctx context.Context, userID uint) ([]ArticleGame, error)
	Create(ctx context.Context, articlegame *ArticleGame) error
	Update(ctx context.Context, id uint, articlegame ArticleGame) (ArticleGame, error)
	Delete(ctx context.Context, id uint) error
}
