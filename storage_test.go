package gourmand

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func newTestFirebaseStorage() *firebaseStorage {
	s := NewFirebaseStorage()

	s.uri = map[string]string{
		"menu": "test-menus",
	}

	return s
}

func saveTestMenu(s *firebaseStorage) (*Menu, error) {
	now := time.Now()
	menu := &Menu{
		Id:        "foo",
		Items:     []*MenuItem{},
		CrawledAt: now,
	}

	err := s.SaveMenu(menu)

	return menu, err
}

func TestFirebaseStorage_FindMenu(t *testing.T) {
	s := newTestFirebaseStorage()

	var menu *Menu

	menu, err := s.FindMenu("test")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(menu.Items))
	assert.Equal(t, "Test", menu.Items[0].Name)

	menu, err = s.FindMenu("inexistent")
	assert.Nil(t, err)
	assert.Nil(t, menu)
}
func TestFirebaseStorage_SaveMenu(t *testing.T) {
	s := newTestFirebaseStorage()
	menu, err := saveTestMenu(s)
	assert.Nil(t, err)

	newMenu, err := s.FindMenu("foo")
	assert.Nil(t, err)
	assert.Equal(t, menu.Id, newMenu.Id)
}

func TestFirebaseStorage_FindAllMenus(t *testing.T) {
	s := newTestFirebaseStorage()

	menus, err := s.FindAllMenus()

	assert.Nil(t, err)
	assert.True(t, len(menus) >= 1)

	found := false
	for _, menu := range menus {
		if menu.Id == "test" {
			found = true
			break
		}
	}
	assert.True(t, found)
}

func TestFirebaseStorage_DeleteMenu(t *testing.T) {
	s := newTestFirebaseStorage()

	menu, err := saveTestMenu(s)
	assert.Nil(t, err)

	err = s.DeleteMenu(menu)
	assert.Nil(t, err)

	menu, err = s.FindMenu(menu.Id)
	assert.Nil(t, err)

	assert.Nil(t, menu)
}
