package gourmand

import (
	"encoding/json"
	"testing"
	"time"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestCrawledAtMarshalling(t *testing.T) {
	now := time.Now()

	m := &Menu{
		CrawledAt: now,
	}

	j, err := json.Marshal(m)

	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`{"items":null,"date":"","crawled_at":"%s"}`, now.Format(time.RFC3339Nano)), string(j))
}

func TestCrawledAtUnmarshalling(t *testing.T) {
	m := &Menu{}
	now := time.Now()
	j := fmt.Sprintf(`{"items":null,"date":"","crawled_at":"%s"}`, now.Format(time.RFC3339Nano))

	json.Unmarshal([]byte(j), m)

	assert.Equal(t, now, m.CrawledAt)
}
