package gourmand

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppFetchPage(t *testing.T) {
	return
	app := NewApp()

	r, err := app.fetchPage()

	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestMenuUpdate(t *testing.T) {
	now := time.Now()

	oldMenu := &Menu{
		Items: []*MenuItem{&MenuItem{
			ImageUrl: NO_PICTURE_YET_URL,
		}},
	}

	newMenu := &Menu{
		Items: []*MenuItem{&MenuItem{
			ImageUrl:  "foo",
			ExpiredAt: now,
			LimitedAt: now,
		}},
	}

	updateMenu(oldMenu, newMenu)

	item := oldMenu.Items[0]
	assert.Equal(t, "foo", item.ImageUrl)
	assert.False(t, item.ExpiredAt.IsZero())
	assert.False(t, item.LimitedAt.IsZero())

	oldMenu.Items[0] = &MenuItem{
		ImageUrl:  "foo",
		ExpiredAt: now,
		LimitedAt: now,
	}

	oldMenu.Items[0] = &MenuItem{
		ImageUrl:  NO_PICTURE_YET_URL,
		ExpiredAt: time.Now(),
		LimitedAt: time.Now(),
	}

	updateMenu(oldMenu, newMenu)

	item = oldMenu.Items[0]
	assert.Equal(t, "foo", item.ImageUrl)
	assert.True(t, now.Equal(item.ExpiredAt))
	assert.True(t, now.Equal(item.LimitedAt))
}
