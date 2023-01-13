package data

import (
	"context"
	"go-postgres/pkg/pedido"
	"time"
)

type PedidosRepository struct {
	Data *Data
}

func (p *PedidosRepository) GetOne(ctx context.Context, id uint) (pedido.PedidoGame, error) {
	q := `SELECT id, estado,user_id, article_id, created_at, updated_at  FROM pedido WHERE id = $1;`

	row := p.Data.DB.QueryRowContext(ctx, q, id)
	var pe pedido.PedidoGame

	err := row.Scan(&pe.ID, &pe.Estado, &pe.UserID, &pe.CreatedAt, &pe.UpdatedAt)

	if err != nil {
		return pedido.PedidoGame{}, err
	}

	return pe, nil
}

func (p *PedidosRepository) GetAll(ctx context.Context) ([]pedido.PedidoGame, error) {
	q := `SELECT id, estado,user_id, article_id, created_at, updated_at  FROM pedido; `

	rows, err := p.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pedidos []pedido.PedidoGame
	for rows.Next() {
		var pe pedido.PedidoGame
		rows.Scan(&pe.ID, &pe.Estado, &pe.UserID, &pe.ArticleId, &pe.CreatedAt, &pe.UpdatedAt)
		pedidos = append(pedidos, pe)
	}

	return pedidos, nil
}

/*
	BY ID USER
*/

func (p *PedidosRepository) GetByUser(ctx context.Context, userID uint) ([]pedido.PedidoGame, error) {
	q := `
    SELECT id, estado,user_id, article_id, created_at, updated_at  
		FROM pedido
		WHERE user_id = $1; 
    `

	rows, err := p.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pedidos []pedido.PedidoGame
	for rows.Next() {
		var pe pedido.PedidoGame
		rows.Scan(&pe.ID, &pe.Estado, &pe.UserID, &pe.ArticleId, &pe.CreatedAt, &pe.UpdatedAt)
		pedidos = append(pedidos, pe)
	}

	return pedidos, nil
}

/*
	INSERTAR
*/

func (p *PedidosRepository) Create(ctx context.Context, pe *pedido.PedidoGame) error {
	q := `
    INSERT INTO PEDIDO (estado, user_id,article_id,created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id;
    `

	stmt, err := p.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, pe.Estado, pe.UserID, pe.ArticleId, time.Now(), time.Now())

	err = row.Scan(&pe.ID)
	if err != nil {
		return err
	}

	return nil
}

/*
	UPDATE
*/

func (p *PedidosRepository) Update(ctx context.Context, id uint, pe pedido.PedidoGame) (pedido.PedidoGame, error) {
	q := `UPDATE pedido set estado=$1, user_id=$2, article_id=$3, updated_at=$4 WHERE id=$5; `

	stmt, err := p.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return pe, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, pe.Estado, pe.UserID, pe.ArticleId, time.Now(), id,
	)
	if err != nil {
		return pe, err
	}

	return pe, nil
}

/*
	DELETE
*/

func (p *PedidosRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM PEDIDO WHERE id=$1;`

	stmt, err := p.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
