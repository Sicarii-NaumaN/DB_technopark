package repository

import "forum/domain/entity"

type ForumRepository interface {
	CreateForum(forumInput *entity.Forum) error
	GetForumDetails(slug string) (*entity.Forum, error)
	CreateThread(threadInput *entity.Thread) error
	GetForumUsers(slug string, limit int32, since string, desc bool) ([]entity.User, error)
	GetForumThreads(slug string, limit int32, since string, desc bool) (*entity.Thread, error)
	CheckForum(slug string) error
}

