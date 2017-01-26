package gourmand

import (
	"errors"
	"io"
	"time"

	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Crawler interface {
	Parse(io.Reader) error
	Menu() (*Menu, error)
}

type goQueryCrawler struct {
	doc  *goquery.Document
	html string
	menu *Menu
}

func NewCrawler() *goQueryCrawler {
	return &goQueryCrawler{}
}

func (c *goQueryCrawler) Parse(r io.Reader) (err error) {
	c.doc, err = goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}

	c.parseMenu()

	return nil
}

func (c *goQueryCrawler) Menu() (*Menu, error) {
	if c.doc == nil {
		return nil, errors.New("No document was parsed.")
	}

	return c.menu, nil
}

func (c *goQueryCrawler) parseMenu() error {
	items := []*MenuItem{}
	itemNodes := c.doc.Find(".menu-item")

	if itemNodes == nil {
		return errors.New("Cannot find items in menu.")
	}

	itemNodes.Each(func(i int, node *goquery.Selection) {
		item := &MenuItem{}

		item.Name = node.Find("h3 > a").Text()
		item.ImageUrl, _ = node.Find(".image-wrapper a img").Attr("src")
		item.ImageUrl = "https://www.sectorgurmand.ro" + item.ImageUrl
		item.Description = node.Find("p.max_lines").Text()
		item.Note = strings.TrimSpace(node.Find(".meta-cat").Text())

		if node.Find(".image-wrapper .overlay span").Size() > 0 {
			item.ExpiredAt = time.Now()
		}

		if node.Find(".image-wrapper .overlay_up span").Size() > 0 {
			item.LimitedAt = time.Now()
		}

		items = append(items, item)
	})

	c.menu = &Menu{
		Id:        time.Now().Format(SQL_DATE_FORMAT),
		Items:     items,
		CrawledAt: time.Now(),
	}

	return nil
}
