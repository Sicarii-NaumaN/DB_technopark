package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
)

type ForumApp struct {
	f repository.ForumRepository
}

func NewForumApp(f repository.ForumRepository) *ForumApp {
	return &ForumApp{f: f}
}

type ForumAppInterface interface {
	CreateForum(forumInput *entity.Forum) error
	GetForumDetails(slug string) (*entity.Forum, error)
	CreateThread(threadInput *entity.Thread) error
	GetForumUsers(slug string, limit int32, since string, desc bool) ([]entity.User, error)
	GetForumThreads(slug string, limit int32, since string, desc bool) (*entity.Thread, error)
	CheckForum(slug string) error
}

func (f *ForumApp) CreateForum(forumInput *entity.Forum) error {
	return nil
}

func (f *ForumApp) GetForumDetails(slug string) (*entity.Forum, error) {
	return nil, nil
}

func (f *ForumApp) CreateThread(threadInput *entity.Thread) error {
	return nil
}

func (f *ForumApp) GetForumUsers(slug string, limit int32, since string, desc bool) ([]entity.User, error) {
	return nil, nil
}

func (f *ForumApp) GetForumThreads(slug string, limit int32, since string, desc bool) (*entity.Thread, error) {
	return nil, nil
}

func (f *ForumApp) CheckForum(slug string) error {
	return nil
}

