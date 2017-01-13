package gourmand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppFetchPage(t *testing.T) {
	app := NewApp()

	r, err := app.fetchPage()

	assert.Nil(t, err)
	assert.NotNil(t, r)
}
