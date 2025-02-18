package models

import "time"

// tip : 帖子结构体

type Post struct {
	ID          int64     `db:"post_id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	AuthorID    int64     `db:"author_id" json:"author_id"`
	CommunityID int64     `db:"community_id" json:"community_id"`
	Status      int8      `db:"status" json:"status"`
	CreateTime  time.Time `db:"create_time" json:"create_time"`
	UpdateTime  time.Time `db:"update_time" json:"update_time"`
}
