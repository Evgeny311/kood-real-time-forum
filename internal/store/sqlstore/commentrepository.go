package sqlstore

import (
	"fmt"
	"time"

	"forum/internal/models"

	"github.com/google/uuid"
)

type CommentRepository struct {
	store *Store
}

func (c *CommentRepository) Create(comment *models.Comment, userId string) error {
	// Add other neccessary information for posts
	comment.ID = uuid.New().String()
	comment.Timestamp = time.Now()
	comment.UserID = userId

	insertQuery := `INSERT INTO comment (id, user_Id, post_id, parent_id, content, timestamp) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := c.store.Db.Exec(insertQuery, comment.ID, comment.UserID, comment.PostID, comment.ParentID, comment.Content, comment.Timestamp)
	if err != nil {
		return fmt.Errorf("Database SQL query error: %v", err)
	}

	return nil
}

func (c *CommentRepository) Delete(id string) error {
	query := `WITH RECURSIVE CommentHierarchy AS (
    -- Anchor member: Start with the comments to be deleted
    SELECT id
    FROM comment
    WHERE id = ?

    UNION ALL

    -- Recursive member: Join with sub-comments
    SELECT c.id
    FROM comment c
    JOIN CommentHierarchy ch ON c.parent_id = ch.id
)
DELETE FROM comment
WHERE id IN (SELECT id FROM CommentHierarchy);
`

	if _, err := c.store.Db.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (c *CommentRepository) DeleteAllUnderPost(post_id string) error {
	query := `DELETE FROM comment WHERE post_id = ?;
`

	if _, err := c.store.Db.Exec(query, post_id); err != nil {
		return err
	}

	return nil
}

func (c *CommentRepository) Get(id string) (*[]models.Comment, error) {
	query := `WITH RECURSIVE CommentHierarchy AS (
		-- Anchor member: Start with the top-level comments for the post
		SELECT
			c.id,
			c.user_id,
			c.post_id,
			c.parent_id,
			c.content,
			c.timestamp,
			(SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS subcomment_count
		FROM
			comment c
		WHERE
			c.post_id = ? AND (c.parent_id IS NULL OR c.parent_id = '')
	
		UNION ALL

		-- Recursive member: Join with sub-comments
		SELECT
			c.id,
			c.user_id,
			c.post_id,
			c.parent_id,
			c.content,
			c.timestamp,
			(SELECT COUNT(*) FROM comment subc WHERE subc.parent_id = c.id) AS subcomment_count
		FROM
			comment c
		JOIN
			CommentHierarchy ch ON c.parent_id = ch.id
	)
	SELECT * FROM CommentHierarchy;`

	rows, err := c.store.Db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf(`Error while executing db query comments - %v`, err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.Timestamp, &comment.SubcommentCount)
		if err != nil {
			return nil, fmt.Errorf(`Error while looping through comments - %v`, err)
		}
		// get user
		user, err := c.store.User().FindByID(comment.UserID)
		if err != nil {
			return nil, fmt.Errorf(`Error while retrieving user information for comments - %v`, err)
		}

		comment.Nickname = user.Username
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(`Looping error - %v`, err)
	}
	return &comments, nil
}
