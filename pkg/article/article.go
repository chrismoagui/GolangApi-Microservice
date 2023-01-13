package article

import "time"

type ArticleGame struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Price       uint      `json:"price,omitempty"`
	Picture     string    `json:"picture,omitempty"`
	Description string    `json:"description,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
