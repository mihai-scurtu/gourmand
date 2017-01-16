package gourmand

import (
	"io"
	"net/http"
)

const URL = "https://www.sectorgurmand.ro"

const NO_PICTURE_YET_URL = "https://www.sectorgurmand.ro/images/produs_poza_in_curand.jpg"

type App struct {
	crawler Crawler
	storage Storage
}

func NewApp() *App {
	return &App{
		crawler: NewCrawler(),
		storage: &firebaseStorage{},
	}
}

func (app *App) Run() error {
	var err error

	newMenu, err := app.fetchMenu()
	if err != nil {
		return err
	}

	existingMenu, err := app.storage.FindMenu(newMenu.Date)
	if err != nil {
		return err
	}

	if existingMenu == nil {
		existingMenu = newMenu
	} else {
		updateMenu(existingMenu, newMenu)
	}

	app.storage.SaveMenu(existingMenu)

	return nil
}

func (app *App) fetchPage() (io.Reader, error) {
	resp, err := http.Get(URL)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (app *App) fetchMenu() (*Menu, error) {
	body, err := app.fetchPage()
	if err != nil {
		return nil, err
	}

	err = app.crawler.Parse(body)
	if err != nil {
		return nil, err
	}

	return app.crawler.Menu()
}

func updateMenu(currentMenu *Menu, newMenu *Menu) {
	currentMenu.CrawledAt = newMenu.CrawledAt

	for i, item := range currentMenu.Items {
		newItem := newMenu.Items[i]

		if item.ExpiredAt.IsZero() {
			item.ExpiredAt = newItem.ExpiredAt
		}

		if item.LimitedAt.IsZero() {
			item.ExpiredAt = newItem.ExpiredAt
		}

		if item.ImageUrl == NO_PICTURE_YET_URL {
			item.ImageUrl = newItem.ImageUrl
		}
	}
}
