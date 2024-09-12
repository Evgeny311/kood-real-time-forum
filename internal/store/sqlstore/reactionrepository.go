package sqlstore

import (
	"fmt"

	"forum/internal/models"

	"github.com/google/uuid"
)

type ReactionRepository struct {
	store *Store
}

func (r *ReactionRepository) Create(reaction *models.Reaction) error {
	panic("implement me")
}

func (r *ReactionRepository) AddToParent(parentId string, reactionId string, userId string) error {
	id := uuid.New().String()

	query := `
	INSERT INTO parentReaction (id, user_id, reaction_id, parent_id, parent_type) VALUES (?, ?, ?, ?, ?);`
	_, err := r.store.Db.Exec(query, id, userId, reactionId, parentId, "")
	if err != nil {
		return err
	}

	return nil
}

func (r *ReactionRepository) RemoveFromParent(postId string, reactionId string, userId string) error {
	query := `
DELETE FROM parentReaction
WHERE user_id = ?
AND reaction_id = ?
AND parent_id = ?;
	`
	resp, err := r.store.Db.Exec(query, userId, reactionId, postId)
	if err != nil {
		return err
	}
	num, err := resp.RowsAffected()
	if num == 0 {
		return fmt.Errorf("nothing was removed from database, paramaters don't match a result in DB")
	}

	return nil
}

func (r *ReactionRepository) GetByParentID(parentId string) (*[]models.Reaction, error) {
	query := ` 
SELECT r.id, r.emoji, r.description
FROM parentReaction pr
INNER JOIN reaction r ON pr.reaction_id = r.id
WHERE pr.parent_id = ?;
	`

	var reactions []models.Reaction

	resp, err := r.store.Db.Query(query, parentId)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	for resp.Next() {
		var reaction models.Reaction
		if err := resp.Scan(&reaction.ID, &reaction.Emoji, &reaction.Description); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	return &reactions, nil
}

func (r *ReactionRepository) GetByUserParentID(parentId string, userId string) (*[]models.Reaction, error) {
	query := ` 
SELECT r.id, r.emoji, r.description
FROM parentReaction pr
INNER JOIN reaction r ON pr.reaction_id = r.id
WHERE pr.user_id = ? AND pr.parent_id = ?;
	`

	var reactions []models.Reaction

	resp, err := r.store.Db.Query(query, userId, parentId)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	for resp.Next() {
		var reaction models.Reaction
		if err := resp.Scan(&reaction.ID, &reaction.Emoji, &reaction.Description); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	return &reactions, nil
}

func (r *ReactionRepository) GetAll() (*[]models.Reaction, error) {
	query := ` 
	SELECT * FROM reaction
	`

	var reactions []models.Reaction

	resp, err := r.store.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	for resp.Next() {
		var reaction models.Reaction
		if err := resp.Scan(&reaction.ID, &reaction.Emoji, &reaction.Description); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	return &reactions, nil
}
