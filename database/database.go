package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/write-blog-every-week/write-blog-every-week-remind/config"
	"github.com/write-blog-every-week/write-blog-every-week-remind/slack"
)

// WriteBlogEveryWeek DynamoDBからのデータを格納する構造体
type WriteBlogEveryWeek struct {
	UserID       string `dynamo:"user_id"`
	UserName     string `dynamo:"user_name"`
	FeedURL      string `dynamo:"feed_url"`
	RequireCount int    `dynamo:"require_count"`
}

// FindAll DynamoDBからデータを全取得する
func FindAll(configData config.ConfigData) []WriteBlogEveryWeek {
	var writeBlogEveryWeek []WriteBlogEveryWeek
	table := getTableObject(configData)
	err := table.Scan().All(&writeBlogEveryWeek)
	if err != nil {
		panic("データの読み込みエラー => " + err.Error())
	}

	return writeBlogEveryWeek
}

// FindByPK Pkを指定して1件取得
func FindByPK(configData config.ConfigData, pk string) WriteBlogEveryWeek {
	var writeBlogEveryWeek WriteBlogEveryWeek
	table := getTableObject(configData)
	table.Get("user_id", pk).One(&writeBlogEveryWeek)
	return writeBlogEveryWeek
}

// ResetRequireCount ブログの必要記事数をリフレッシュする
func ResetRequireCount(configData config.ConfigData, userList map[string]int) {
	table := getTableObject(configData)
	for userID, requireCount := range userList {
		err := table.Update("user_id", userID).Set("require_count", requireCount).Run()
		if err != nil {
			panic("データ保存エラー => " + err.Error())
		}
	}
}

// CreateUser 新しいユーザーデータを作成する
func CreateUser(configData config.ConfigData, slackParams *slack.SlackParams) {
	writeBlogEveryWeek := WriteBlogEveryWeek{
		UserID:       slackParams.UserID,
		UserName:     slackParams.UserName,
		FeedURL:      slackParams.Text,
		RequireCount: 1,
	}
	table := getTableObject(configData)
	if err := table.Put(writeBlogEveryWeek).Run(); err != nil {
		panic("登録エラー => " + err.Error())
	}
}

// DeleteUser ユーザーデータを削除する
func DeleteUser(configData config.ConfigData, member WriteBlogEveryWeek) {
	table := getTableObject(configData)
	if err := table.Delete("user_id", member.UserID).Run(); err != nil {
		panic("削除エラー => " + err.Error())
	}
}

/**
 * DynamoDBのテーブルオブジェクトを取得する
 */
func getTableObject(configData config.ConfigData) dynamo.Table {
	credential := credentials.NewStaticCredentials(configData.AWS.AccessKey, configData.AWS.SecretKey, "")
	db := dynamo.New(session.New(), &aws.Config{
		Credentials: credential,
		Region:      aws.String(configData.AWS.Region),
	})

	table := db.Table(configData.AWS.DataBase)

	return table
}
