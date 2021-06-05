package persistence

import (
	"forum/domain/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ForumRepo struct {
	db *pgxpool.Pool
}

func NewForumRepository(db *pgxpool.Pool) *ForumRepo {
	return &ForumRepo{db}
}

func (f *ForumRepo) CreateForum(forumInput *entity.Forum) error {
	return nil
}

func (f *ForumRepo) GetForumDetails(slug string) (*entity.Forum, error) {
	return nil, nil
}

func (f *ForumRepo) CreateThread(threadInput *entity.Thread) error {
	return nil
}

func (f *ForumRepo) GetForumUsers(slug string, limit int32, since string, desc bool) ([]entity.User, error) {
	return nil, nil
}

func (f *ForumRepo) GetForumThreads(slug string, limit int32, since string, desc bool) (*entity.Thread, error) {
	return nil, nil
}

func (f *ForumRepo) CheckForum(slug string) error {
	return nil
}