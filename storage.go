package gourmand

import (
	"log"

	"github.com/zabawaba99/firego"
)

type Storage interface {
	FindMenu(string) (*Menu, error)
	SaveMenu(*Menu) error
}

const FIREBASE_URL = "https://gourmand-6fe46.firebaseio.com/"

type firebaseStorage struct{}

func (s *firebaseStorage) FindMenu(key string) (menu *Menu, err error) {
	fb := firego.New(menuUrl(key), nil)

	err = fb.Value(&menu)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return menu, nil
}

func (s *firebaseStorage) SaveMenu(menu *Menu) error {
	fb := firego.New(menuUrl(menu.Date), nil)

	err := fb.Set(menu)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func menuUrl(key string) string {
	return FIREBASE_URL + "menus/" + key
}
