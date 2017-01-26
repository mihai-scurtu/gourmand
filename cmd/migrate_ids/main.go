package main

import (
	"time"

	"log"

	"github.com/mihai-scurtu/gourmand"
	"gopkg.in/zabawaba99/firego.v1"
)

type oldMenu struct {
	Date      string               `json:"date"`
	CrawledAt time.Time            `json:"crawled_at"`
	Items     []*gourmand.MenuItem `json:"items"`
}

func main() {
	s := gourmand.NewFirebaseStorage()
	fb := firego.New(s.MenuUrl(""), nil)

	var menus map[string]*oldMenu
	if err := fb.Value(&menus); err != nil {
		log.Fatal("Could not parse old menus: ", err.Error())
	}

	log.Printf("%+v", menus)

	for _, om := range menus {
		log.Printf("%+v", *om)
		t, err := time.Parse("2 Jan 2006", om.Date)
		if err != nil {
			log.Println("Cannot parse old date: ", err.Error())
			continue
		}

		menu := &gourmand.Menu{
			Id:        t.Format(gourmand.SQL_DATE_FORMAT),
			CrawledAt: om.CrawledAt,
			Items:     om.Items,
		}

		if err = s.SaveMenu(menu); err != nil {
			log.Println("Cannot save new menu: ", err.Error())
		}

		fb = firego.New(s.MenuUrl(om.Date), nil)
		if err = fb.Remove(); err != nil {
			log.Println("Cannot delete old menu", err.Error())
		}
	}
}
