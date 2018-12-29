package rss

import (
	"testing"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/write-blog-every-week/write-blog-every-week-remind/date"
)

func parse(published string) time.Time {
	// なぜかRFC3339でうまくパースできないのでRFC1123Z
	parsed, _ := time.ParseInLocation(time.RFC1123Z, published, time.Local)
	return parsed
}

func item(published string) *gofeed.Item {
	parsed := parse(published)
	return &gofeed.Item {
		Published: published,
		PublishedParsed: &parsed,
	}
}

func TestGetLatestFeedPubDate(t *testing.T) {
	date.SetFakeTime(time.Date(2018, 12, 27, 0, 0, 0, 0, time.Local))
	thisMonday := date.GetThisMonday()
	tests := []struct {
		name		 string
		feed		 *gofeed.Feed
		requireCount int
		want		 time.Time
	}{
		{
			name: "not enough feeds",
			feed: &gofeed.Feed{},
			requireCount: 0,
			want: thisMonday,
		},
		{
			name: "1 feed this week",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 26 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 0,
			want: parse("Wed, 26 Dec 2018 19:00:00 +0900"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLatestFeedPubDate(tt.feed, tt.requireCount, nil, time.Local); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}