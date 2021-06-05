package persistence

import (
	"forum/domain/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ThreadRepo struct {
	db *pgxpool.Pool
}

func NewThreadRepository(db *pgxpool.Pool) *ThreadRepo {
	return &ThreadRepo{db: db}
}


func (t *ThreadRepo) CreatePosts(thread *entity.Thread, posts []entity.Post) error {
	return nil
}

func (t *ThreadRepo) CreateThread(thread *entity.Thread) error {
	return nil
}

func (t *ThreadRepo) GetThreadPosts(slug string, limit int32, since string, order string) ([]entity.Post, error) {
	return nil, nil
}

func (t *ThreadRepo) GetThreadPostsTree(slug string, limit int32, since string, order string) ([]entity.Post, error) {
	return nil, nil
}

func (t *ThreadRepo) GetThreadPostsParentTree(slug string, limit int32, since string, order string) ([]entity.Post, error) {
	return nil, nil
}


func (t *ThreadRepo) CheckThreadBySlug(slug string) error {
	return nil
}

func (t *ThreadRepo) CheckThreadByID(ID int) error {
	return nil
}

func (t *ThreadRepo) GetThreadForumAndID(slugOrID string) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadRepo) GetThreadsByForumSlug(slug string, limit int32, since string, desc bool) ([]entity.Thread, error) {
	return nil, nil
}

func (t *ThreadRepo) VoteForThread(vote *entity.Vote) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadRepo) GetThreadBySlug(slug string) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadRepo) GetThreadByID(ID int) (*entity.Thread, error) {
	return nil, nil
}

func (t *ThreadRepo) UpdateThread(newThreadData *entity.Thread) error {
	return nil
}