package rss

import (
	time "time"

	database "../database"
	gofeed "github.com/mmcdole/gofeed"
)

// FindTargetUserList ブログを書いていないユーザーを取得する
func FindTargetUserList(allMemberDataList []database.WriteBlogEveryWeek, targetMonday time.Time) map[string]int {
	// 日本時間に合わせる
	locale, _ := time.LoadLocation("Asia/Tokyo")
	parser := gofeed.NewParser()

	results := make(map[string]int)
	for i := 0; i < len(allMemberDataList); i++ {
		for j := 0; j < allMemberDataList[i].RequireCount; j++ {
			// 最新フィードの公開日を取得する
			latestPublishDate := getLatestFeedPubDate(allMemberDataList[i].FeedURL, j, parser, locale)

			// 今週の月曜日がAfterになる = 今週ブログを書いていない
			if targetMonday.After(latestPublishDate) {
				if _, ok := results[allMemberDataList[i].UserID]; !ok {
					// データがない場合は初期化
					results[allMemberDataList[i].UserID] = 0
				}

				results[allMemberDataList[i].UserID]++
			}
		}
	}

	return results
}

// getLatestFeedPubDate 最新フィードの公開日を取得する
func getLatestFeedPubDate(feedURL string, requireCount int, parser *gofeed.Parser, locale *time.Location) time.Time {
	// フィードを取得
	feed, err := parser.ParseURL(feedURL)
	if err != nil {
		panic("フィードが取得できませんでした。失敗したフィードURL => " + feedURL)
	}

	// 最新日を取得
	published := feed.Items[requireCount].Published
	latest, err := time.ParseInLocation(time.RFC3339, published, locale)
	if err != nil {
		// 取得できない = フォーマットを変えれば取得できる可能性がある
		latest2, err := time.ParseInLocation(time.RFC1123Z, published, locale)
		if err != nil {
			// それでも取得できない場合は、フィードで取得した生データをもらう
			latest = *feed.Items[requireCount].PublishedParsed
		} else {
			latest = latest2
		}
	}

	return latest
}
