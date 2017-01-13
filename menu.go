package gourmand

import "time"

type Menu struct {
	Items     []*MenuItem `json:"items"`
	Date      string      `json:"date"`
	CrawledAt time.Time   `json:"crawled_at,string"`
}

type MenuItem struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Note        string    `json:"note"`
	LimitedAt   time.Time `json:"limited_at,string"`
	ExpiredAt   time.Time `json:"expired_at,string"`
}
