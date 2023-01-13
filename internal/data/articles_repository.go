package data

import (
	"context"
	"go-postgres/pkg/article"
	"time"
)

type ArticlesRepository struct {
	Data *Data
}

/*
	GET ONE

*/

func (pr *ArticlesRepository) GetOne(ctx context.Context, id uint) (article.ArticleGame, error) {
	q := `SELECT id, name, price, picture, description, user_id, created_at, updated_at  FROM articles WHERE id = $1;`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)
	var art article.ArticleGame

	err := row.Scan(&art.ID, &art.Name, &art.Price, &art.Description, &art.UserID, &art.CreatedAt, &art.UpdatedAt)

	if err != nil {
		return article.ArticleGame{}, err
	}

	return art, nil
}

/*
	GET ALL
*/

func (pr *ArticlesRepository) GetAll(ctx context.Context) ([]article.ArticleGame, error) {
	q := `SELECT id, name, price, picture, description, user_id, created_at, updated_at FROM articles; `

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []article.ArticleGame
	for rows.Next() {
		var p article.ArticleGame
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Picture, &p.Description, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		articles = append(articles, p)
	}

	return articles, nil
}

/*
	BY ID USER
*/

func (pr *ArticlesRepository) GetByUser(ctx context.Context, userID uint) ([]article.ArticleGame, error) {
	q := `
    SELECT id, name, price, picture, description, user_id, created_at, updated_at
        FROM articles
        WHERE user_id = $1;
    `

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []article.ArticleGame
	for rows.Next() {
		var p article.ArticleGame
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description, &p.UserID, &p.CreatedAt, &p.UpdatedAt)
		articles = append(articles, p)
	}

	return articles, nil
}

/*
	INSERTAR
*/

func (pr *ArticlesRepository) Create(ctx context.Context, p *article.ArticleGame) error {
	q := `
    INSERT INTO ARTICLES (name, price, picture, description, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Name, p.Price, p.Description, p.UserID, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

/*
	UPDATE
*/

func (pr *ArticlesRepository) Update(ctx context.Context, id uint, p article.ArticleGame) (article.ArticleGame, error) {
	q := `UPDATE articles set name=$1, price=$2, description=$3, updated_at=$4 WHERE id=$5; `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return p, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Name, p.Price, p.Description, time.Now(), id,
	)
	if err != nil {
		return p, err
	}

	return p, nil
}

/*
	DELETE
*/

func (pr *ArticlesRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM ARTICLES WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
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
