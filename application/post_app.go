package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
)

type PostApp struct {
	p repository.PostRepository
}

func NewPostApp(p repository.PostRepository) *PostApp {
	return &PostApp{p: p}
}

type PostAppInterface interface {
	GetPostDetails(postID int) (*entity.Post, error)
	ChangePostMessage(post *entity.Post) (*entity.Post, error)
}

func (p *PostApp) GetPostDetails(postID int) (*entity.Post, error) {
	return p.p.GetPostDetails(postID)
}

func (p *PostApp) ChangePostMessage(post *entity.Post) (*entity.Post, error) {
	return p.p.ChangePostMessage(post)
}
