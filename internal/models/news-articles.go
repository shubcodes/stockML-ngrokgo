package models

type NewsArticle struct {
	Headline       string `json:"headline"`
	Abstract       string `json:"abstract"`
	LeadParagraph  string `json:"lead_paragraph"`
	WebURL         string `json:"web_url"`
	PublishedDate  string `json:"pub_date"`
	ThumbnailURL   string `json:"thumbnail_url"`
	ThumbnailWidth int    `json:"thumbnail_width"`
	ThumbnailHight int    `json:"thumbnail_height"`
}
