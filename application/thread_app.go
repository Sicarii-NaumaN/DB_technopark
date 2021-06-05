package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
	"strconv"
)

type ThreadApp struct {
	t repository.ThreadRepository
}

func NewThreadApp(f repository.ThreadRepository) *ThreadApp {
	return &ThreadApp{t: f}
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
	return t.t.CreatePosts(thread, posts)
}

func (t *ThreadApp) CreateThread(thread *entity.Thread) error {
	return t.t.CreateThread(thread)
}

func (t *ThreadApp) GetThreadPosts(slug string, limit int32, since string, sort string, desc bool) ([]entity.Post, error) {
	order := "ASC"
	switch desc {
	case true:
		order = "DESC"
	}

	switch sort {
	case "flat":
		return t.t.GetThreadPosts(slug, limit, order, since)
	case "tree":
		return t.t.GetThreadPostsTree(slug, limit, order, since)
	case "parent_tree":
		return t.t.GetThreadPostsParentTree(slug, limit, order, since)
	default:
		return t.t.GetThreadPosts(slug, limit, order, since)
	}
}

func (t *ThreadApp) CheckThread(slugOrID string) error {
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		return t.t.CheckThreadBySlug(slugOrID)
	}

	return 	t.t.CheckThreadByID(id)
}

func (t *ThreadApp) VoteForThread(vote *entity.Vote) (*entity.Thread, error) {
	return t.t.VoteForThread(vote)
}

func (t *ThreadApp) GetThread(slugOrID string) (*entity.Thread, error) {
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		return t.t.GetThreadBySlug(slugOrID)
	}
	return t.t.GetThreadByID(id)
}

func (t *ThreadApp) GetThreadForumAndID(slugOrID string) (*entity.Thread, error) {
	return t.t.GetThreadForumAndID(slugOrID)
}

func (t *ThreadApp) GetThreadsByForumSlug(slug string, limit int32, since string, desc bool) ([]entity.Thread, error) {
	return t.t.GetThreadsByForumSlug(slug , limit, since , desc)
}

func (t *ThreadApp) UpdateThread(slugOrID string, newThreadData *entity.Thread) error {
	newThreadData.Slug = &slugOrID
	id, err := strconv.Atoi(slugOrID)
	if err != nil {
		id = 0
	}

	newThreadData.ID = id
	return t.t.UpdateThread(newThreadData)
}