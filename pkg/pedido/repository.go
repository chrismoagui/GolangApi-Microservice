package pedido

import (
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) ([]PedidoGame, error)
	GetOne(ctx context.Context, id uint) (PedidoGame, error)
	GetByUser(ctx context.Context, userID uint) ([]PedidoGame, error)
	Create(ctx context.Context, pedidogame *PedidoGame) error
	Update(ctx context.Context, id uint, pedidogame PedidoGame) (PedidoGame, error)
	Delete(ctx context.Context, id uint) error
}
