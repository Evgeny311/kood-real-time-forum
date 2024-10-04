package store

type Store interface {
	Chat() ChatRepository
	User() UserRepository
	Post() PostRepository
	Category() CategoryRepository
	Reaction() ReactionRepository
	Comment() CommentRepository
}
