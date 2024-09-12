package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"forum/internal/models"

	"github.com/google/uuid"
)

type PostRepository struct {
	store *Store
}

func (r *PostRepository) Create(post *models.Post, categories []models.Category, userId string) error {
	// Add other necessary information for posts
	post.ID = uuid.New().String()
	post.Timestamp = time.Now()
	// Get user id from Sessions/Cookeis
	post.UserID = userId

	insertQuery := `INSERT INTO post (id, title, content, user_id, image_url, timestamp) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.store.Db.Exec(insertQuery, post.ID, post.Title, post.Content, post.UserID, post.ImageURL, post.Timestamp)
	if err != nil {
		return fmt.Errorf("database SQL query error: %v", err)
	}

	// Insert info to join tabel
	insertQuery = `INSERT INTO postCategory (id, post_id, category_id) VALUES (?, ?, ?)`
	for _, category := range categories {
		categoryToCheck, err1 := r.store.Category().Get(category.Name)
		// If there is not added category to db, then add
		if errors.Is(err1, sql.ErrNoRows) {
			_, err = r.store.Db.Exec(insertQuery, uuid.New().String(), post.ID, category.ID)
			if err != nil {
				return fmt.Errorf("database SQL query error: %v", err)
			}
		} else if err1 != nil {
			return fmt.Errorf(err.Error())
			// If there is category added to db, then just take the existing category and add it
		} else {
			_, err1 = r.store.Db.Exec(insertQuery, uuid.New().String(), post.ID, &categoryToCheck.ID)
			if err != nil {
				return fmt.Errorf("database SQL query error: %v", err)
			}
		}
	}

	return nil
}

func (r *PostRepository) Delete(id string) error {
	query := `DELETE FROM post WHERE id = ?`

	if _, err := r.store.Db.Exec(query, id); err != nil {
		return err
	}

	if err := r.store.commentRepository.DeleteAllUnderPost(id); err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Get(id string) (*models.Post, error) {
	query := `SELECT p.*, u.username AS user_name
	FROM post p
	JOIN user u ON p.user_id = u.id
	WHERE p.id = ?`

	var post models.Post

	err := r.store.Db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.ImageURL, &post.Timestamp, &post.UserNickname)
	if err != nil {
		return nil, fmt.Errorf(`error occured while checking posts: %v`, err)
	}
	return &post, nil
}

func (r *PostRepository) GetFeed(offset, limit int, timeStamp time.Time, categoryID string) ([]models.Post, error) {
	query := `SELECT DISTINCT p.*, u.username AS user_name
FROM post p
JOIN user u ON p.user_id = u.id
JOIN postCategory pc ON p.id = pc.post_id
WHERE p.timestamp <= ? AND (pc.category_id = ? OR ? IS NULL OR ? = '')
ORDER BY p.timestamp DESC
LIMIT ? OFFSET ?`

	dateTimeParts := strings.Split(timeStamp.String(), " ")

	timeParts := strings.Split(dateTimeParts[1], ".")

	nearestSecond := dateTimeParts[0] + "T" + timeParts[0]

	rows, err := r.store.Db.Query(query, nearestSecond, categoryID, categoryID, categoryID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts from the database: %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.ImageURL, &post.Timestamp, &post.UserNickname); err != nil {
			return nil, fmt.Errorf("failed to scan post row: %v", err)
		}
		posts = append(posts, post)
	}
	fmt.Println(posts)

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading posts rows: %v", err)
	}

	return posts, nil
}

func (r *PostRepository) GetCommentNumber(postId string) (int, error) {
	query := `WITH RECURSIVE CommentIDs AS (
        SELECT
            c.id
        FROM
            comment c
        WHERE
            c.post_id = ? AND (c.parent_id IS NULL OR c.parent_id = '')
        
        UNION ALL
        
        SELECT
            c.id
        FROM
            comment c
        INNER JOIN
            CommentIDs ch ON c.parent_id = ch.id
    )
    SELECT
        COUNT(*) FROM CommentIDs;`

	var commentCount int
	err := r.store.Db.QueryRow(query, postId).Scan(&commentCount)
	if err != nil {
		return 0, fmt.Errorf("error while executing db query comments - %v", err)
	}

	return commentCount, nil
}
