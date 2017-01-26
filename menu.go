package gourmand

import (
	"log"
	"time"
)

type Menu struct {
	Id        string      `json:"id"`
	Items     []*MenuItem `json:"items"`
	CrawledAt time.Time   `json:"crawled_at,string"`
}

func (m Menu) Date() time.Time {
	date, err := time.Parse(SQL_DATE_FORMAT, m.Id)
	if err != nil {
		log.Printf("Could not parse date from menu id '%s': %s", m.Id, err.Error())
	}

	return date
}

type MenuItem struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Note        string    `json:"note"`
	LimitedAt   time.Time `json:"limited_at,string"`
	ExpiredAt   time.Time `json:"expired_at,string"`
}
