package sqlstore

import (
	"database/sql"

	"forum/internal/store"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Db                 *sql.DB
	chatRepository     *ChatRepository
	userRepository     *UserRepository
	postRepository     *PostRepository
	categoryRepository *CategoryRepository
	commentRepository  *CommentRepository
	reactionRepository *ReactionRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		Db: db,
	}
}

func (s *Store) Chat() store.ChatRepository {
	if s.chatRepository != nil {
		return s.chatRepository
	}

	s.chatRepository = &ChatRepository{
		store: s,
	}

	return s.chatRepository
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository != nil {
		return s.postRepository
	}

	s.commentRepository = &CommentRepository{
		store: s,
	}

	s.postRepository = &PostRepository{
		store: s,
	}

	return s.postRepository
}

func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}

	s.categoryRepository = &CategoryRepository{
		store: s,
	}

	return s.categoryRepository
}

func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store: s,
	}

	return s.commentRepository
}

func (s *Store) Reaction() store.ReactionRepository {
	if s.reactionRepository != nil {
		return s.reactionRepository
	}

	s.reactionRepository = &ReactionRepository{
		store: s,
	}

	return s.reactionRepository
}
