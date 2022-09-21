package models

type PostContents struct {
	Id         int64  `json:"id"`
	ContentUrl string `json:"content_url"`
	PostId     int64  `json:"post_id"`
}

func (m PostContents) TableName() string {
	return "post_contents"
}
