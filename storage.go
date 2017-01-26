package gourmand

import (
	"log"

	"github.com/zabawaba99/firego"
)

type Storage interface {
	FindAllMenus() ([]*Menu, error)
	FindMenu(string) (*Menu, error)
	SaveMenu(*Menu) error
	DeleteMenu(*Menu) error
}

const FIREBASE_URL = "https://gourmand-6fe46.firebaseio.com/"

var URIS = map[string]string{
	"menu": "menus",
}

type firebaseStorage struct {
	url string
	uri map[string]string
}

func NewFirebaseStorage() *firebaseStorage {
	return &firebaseStorage{
		url: FIREBASE_URL,
		uri: URIS,
	}
}

func (s *firebaseStorage) FindMenu(key string) (menu *Menu, err error) {
	fb := firego.New(s.MenuUrl(key), nil)

	err = fb.Value(&menu)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return menu, nil
}

func (s *firebaseStorage) SaveMenu(menu *Menu) error {
	fb := firego.New(s.MenuUrl(menu.Id), nil)

	err := fb.Set(menu)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s firebaseStorage) FindAllMenus() ([]*Menu, error) {
	fb := firego.New(s.MenuUrl(""), nil)

	var hash map[string]*Menu
	if err := fb.Value(&hash); err != nil {
		log.Printf("Cannot get all menus: %s", err.Error())
		return nil, err
	}

	menus := make([]*Menu, 0, len(hash))
	for _, menu := range hash {
		menus = append(menus, menu)
	}

	return menus, nil
}

func (s firebaseStorage) DeleteMenu(menu *Menu) error {
	fb := firego.New(s.MenuUrl(menu.Id), nil)

	return fb.Remove()
}

func (s firebaseStorage) MenuUrl(key string) string {
	return s.url + s.uri["menu"] + "/" + key
}
