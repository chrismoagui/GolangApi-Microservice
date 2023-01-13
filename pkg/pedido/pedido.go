package pedido

import "time"

type PedidoGame struct {
	ID        uint      `json:"id,omitempty"`
	Estado    string    `json:"estado,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	ArticleId uint      `json:"article_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
