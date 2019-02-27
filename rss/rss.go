package rss

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/write-blog-every-week/write-blog-every-week-remind/database"
	"github.com/write-blog-every-week/write-blog-every-week-remind/date"
)

var asiaTokyo, _ = time.LoadLocation("Asia/Tokyo")

// Parser RSSの解析情報を返す
type Parser interface {
	ParseURL(url string) (feed *gofeed.Feed, err error)
}

type rssParser struct {
	parser *gofeed.Parser
}

func (rp *rssParser) ParseURL(url string) (feed *gofeed.Feed, err error) {
	return rp.parser.ParseURL(url)
}

// FindTargetUserList ブログを書いていないユーザーを取得する
func FindTargetUserList(allMemberDataList []database.WriteBlogEveryWeek, targetMonday time.Time) (map[string]int, []*database.WriteBlogEveryWeek) {
	rssParser := &rssParser{gofeed.NewParser()}
	return findTargetUserList(allMemberDataList, targetMonday, rssParser)
}

func findTargetUserList(allMemberDataList []database.WriteBlogEveryWeek, targetMonday time.Time, parser Parser) (map[string]int, []*database.WriteBlogEveryWeek) {
	var errMembers []*database.WriteBlogEveryWeek
	results := make(map[string]int)
	for i, wbem := range allMemberDataList {
		// フィードを取得
		feed, err := parser.ParseURL(wbem.FeedURL)
		if err != nil {
			fmt.Printf(
				"フィードが取得できませんでした。失敗したフィードURL => %s:%s",
				wbem.FeedURL, err.Error(),
			)
			errMembers = append(errMembers, &allMemberDataList[i])
			// ひとまずSKIPしておく
			continue
		}

		// 全ユーザーの情報を入れるため初期化
		results[wbem.UserID] = 0

		for i := 0; i < wbem.RequireCount; i++ {
			// 最新フィードの公開日を取得する
			latestPublishDate := getLatestFeedPubDate(feed, i, asiaTokyo)

			// 今週の月曜日が過去ではない場合は、まだ今週ブログを書いていない
			if !targetMonday.Before(latestPublishDate) {
				results[wbem.UserID]++
			}
		}
	}

	return results, errMembers
}

// getLatestFeedPubDate 最新フィードの公開日を取得する
func getLatestFeedPubDate(feed *gofeed.Feed, requireCount int, locale *time.Location) time.Time {
	if (requireCount + 1) > len(feed.Items) {
		// そもそも記事数が足りない場合は公開日を取得できないのでlatestは、必ず通知対象となる今週の月曜日と合わせる
		return date.GetThisMonday()
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
