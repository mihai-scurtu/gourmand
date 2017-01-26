package gourmand

import (
	"encoding/json"
	"testing"
	"time"

	"fmt"

	"regexp"

	"github.com/stretchr/testify/assert"
)

func TestCrawledAtMarshalling(t *testing.T) {
	now := time.Now()

	m := &Menu{
		CrawledAt: now,
	}

	j, err := json.Marshal(m)

	assert.Nil(t, err)

	assert.Regexp(t, fmt.Sprintf(`.*"crawled_at":"%s".*`, regexp.QuoteMeta(now.Format(time.RFC3339Nano))), string(j))
}

func TestCrawledAtUnmarshalling(t *testing.T) {
	m := &Menu{}
	now := time.Now()
	j := fmt.Sprintf(`{"items":null,"id":"","crawled_at":"%s"}`, now.Format(time.RFC3339Nano))

	err := json.Unmarshal([]byte(j), m)

	assert.Nil(t, err)
	assert.Equal(t, now, m.CrawledAt)
}

func TestMenu_Date(t *testing.T) {
	date := time.Date(1990, 4, 28, 0, 0, 0, 0, time.UTC)
	m := &Menu{
		Id: "1990-04-28",
	}

	assert.True(t, date.Equal(m.Date()))
}
