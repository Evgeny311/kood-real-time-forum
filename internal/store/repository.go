package store

import (
	"time"

	"forum/internal/models"
)

type ChatRepository interface {
	WriteLine(line models.Line) error
	GetLinesInit(chat_id string, time time.Time, limit int, offset int) ([]models.Line, error)
	CheckChatExists(user1 string, user2 string) ([]string, error)
	Create(user1 string, user2 string) (models.Chat, error)
	Get(id string) (models.Chat, error)
}

type UserRepository interface {
	IsUser(id string) (bool, error)
	Create(*models.User) error
	// FindByID is used to send the client user data
	// (password from db has been sanitized)
	FindByID(string) (*models.User, error)
	// Check can be used to get user data not meant to go the client
	// (password from db has NOT BEEN SANITIZED)
	Check(string) (*models.User, error)
	GetAllOtherUsers(user_id string) ([]models.User, error)
	Delete(string) error
}

type ReactionRepository interface {
	Create(*models.Reaction) error
	AddToParent(postId string, reactionId string, userId string) error
	RemoveFromParent(postId string, reactionId string, userId string) error
	GetByUserParentID(parentId string, userId string) (*[]models.Reaction, error)
	GetByParentID(parentId string) (*[]models.Reaction, error)
	GetAll() (*[]models.Reaction, error)
}

type PostRepository interface {
	Create(*models.Post, []models.Category, string) error
	Delete(id string) error
	Get(string) (*models.Post, error)
	GetFeed(offset, limit int, timeStamp time.Time, categoryID string) ([]models.Post, error)
	GetCommentNumber(postId string) (int, error)
}

type CategoryRepository interface {
	Create(*models.Category) error
	Get(string) (*models.Category, error)
	GetForPost(postId string) (*[]models.Category, error)
	GetAll() (*[]models.Category, error)
}

type CommentRepository interface {
	Create(*models.Comment, string) error
	Get(id string) (*[]models.Comment, error)
	Delete(id string) error
	DeleteAllUnderPost(post_id string) error
}
