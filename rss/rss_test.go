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
			name: "1 feed required and written",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 26 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 0,
			want: parse("Wed, 26 Dec 2018 19:00:00 +0900"),
		},
		{
			name: "1 feed required and not written",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 19 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 0,
			want: parse("Wed, 19 Dec 2018 19:00:00 +0900"),
		},
		{
			name: "2 feeds required and only 1 feed exists",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 26 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 1,
			want: thisMonday,
		},
		{
			name: "2 feeds required and only 1 feed written this week",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 26 Dec 2018 19:00:00 +0900"),
					item("Tue, 18 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 1,
			want: parse("Tue, 18 Dec 2018 19:00:00 +0900"),
		},
		{
			name: "2 feeds required and written",
			feed: &gofeed.Feed{
				Items: []*gofeed.Item{
					item("Wed, 26 Dec 2018 19:00:00 +0900"),
					item("Tue, 25 Dec 2018 19:00:00 +0900"),
				},
			},
			requireCount: 1,
			want: parse("Tue, 25 Dec 2018 19:00:00 +0900"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLatestFeedPubDate(tt.feed, tt.requireCount, time.Local); got != tt.want {
				t.Errorf("want \n%s\n, but got \n%s\n", tt.want, got)
			}
		})
	}
}