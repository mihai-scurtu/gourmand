package gourmand

import (
	"os"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrawlerCrawlsHtml(t *testing.T) {
	f, _ := os.Open("./fixtures/page.html")
	defer f.Close()

	c := NewCrawler()

	var err error

	err = c.Parse(f)
	assert.Nil(t, err)

	var menu *Menu
	menu, err = c.Menu()

	assert.Nil(t, err)
	assert.Equal(t, 8, len(menu.Items))
	assert.Equal(t, time.Now().Format("2 Jan 2006"), menu.Date)

	item := menu.Items[0]
	assert.Equal(t, "Supa crema de legume si naut", item.Name)
	assert.Equal(t, "Supa vegetariana cremoasa cu arome usor marocane combina legumele cu condimente calde ca ienibahar, cardamom si nucsoara adaugand valoare proteica prin nautul delicios.",
		item.Description)
	assert.Equal(t, "https://www.sectorgurmand.ro/public/products/supa-crema-de-cartof-dulce-si-turmeric-4155336.jpg",
		item.ImageUrl)
	assert.Equal(t, "Vegetarian", item.Note)
	assert.True(t, item.LimitedAt.IsZero())
	assert.True(t, item.ExpiredAt.IsZero())

	item = menu.Items[2]
	assert.False(t, item.ExpiredAt.IsZero())

	item = menu.Items[3]
	assert.False(t, item.LimitedAt.IsZero())

}
