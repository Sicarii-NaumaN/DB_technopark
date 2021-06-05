package persistence

import (
	"forum/domain/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostRepo struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepo {
	return &PostRepo{db: db}
}

func (p *PostRepo) GetPostDetails(postID int) (*entity.Post, error) {
	return nil, nil
}

func (p *PostRepo) ChangePostMessage(post *entity.Post) (*entity.Post, error) {
	return nil, nil
}
