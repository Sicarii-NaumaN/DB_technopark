package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
)

type ThreadApp struct {
	f repository.ThreadRepository
}

func NewThreadApp(f repository.ThreadRepository) *ThreadApp {
	return &ThreadApp{f: f}
}

type ThreadAppInterface interface {
	CreatePosts(thread *entity.Thread, posts []entity.Post) error
	CreateThread(thread *entity.Thread) error
	GetThreadPosts(slug string, limit int32, since string, sort string, desc bool) ([]entity.Post, error)
	CheckThread(slugOrID string) error
	VoteForThread(vote *entity.Vote) (*entity.Thread, error)
	GetThread(slugOrID string) (*entity.Thread, error)
	GetThreadForumAndID(slugOrID string) (*entity.Thread, error)
	GetThreadsByForumSlug(slug string, limit int32, since string, desc bool) ([]entity.Thread, error)
	UpdateThread(slugOrID string, newThreadData *entity.Thread) error
}

func (t *ThreadApp) CreatePosts(thread *entity.Thread, posts []entity.Post) error {
	return nil
}

func (t *ThreadApp) CreateThread(thread *entity.Thread) error {
	return nil
}

func (t *ThreadApp) GetThreadPosts(slug string, limit int32, since string, sort string, desc bool) ([]entity.Post, error) {
	//switch desc {
	//case true:
	//	order = "DESC"
	//case false:
	//	order = "ASC"
	//}
	return nil, nil
}

func (t *ThreadApp) CheckThread(slugOrID string) error {
	return nil
}

func (t *ThreadApp) VoteForThread(vote *entity.Vote) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadApp) GetThread(slugOrID string) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadApp) GetThreadForumAndID(slugOrID string) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadApp) GetThreadsByForumSlug(slug string, limit int32, since string, desc bool) ([]entity.Thread, error) {
	return nil, nil
}

func (t *ThreadApp) UpdateThread(slugOrID string, newThreadData *entity.Thread) error {
	return nil
}