package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	config "../config"
	slack "../slack"
)

type WriteBlogEveryWeek struct {
	UserID       string `dynamo:"user_id"`
	UserName     string `dynamo:"user_name"`
	FeedURL      string `dynamo:"feed_url"`
	RequireCount int    `dynamo:"require_count"`
}

/**
 * DynamoDBからデータを取得する
 * RequireCountが1以上のもののみ
 */
func FindByMemberData(configData config.ConfigData) []WriteBlogEveryWeek {
	var writeBlogEveryWeek []WriteBlogEveryWeek
	table := getTableObject(configData)
	err := table.Scan().Filter("require_count >= ?", 1).All(&writeBlogEveryWeek)
	if err != nil {
		panic("データの読み込みエラー => " + err.Error())
	}

	return writeBlogEveryWeek
}

/**
 * Pkを指定して1件取得
 */
func FindByPK(configData config.ConfigData, pk string) WriteBlogEveryWeek {
	var writeBlogEveryWeek WriteBlogEveryWeek
	table := getTableObject(configData)
	table.Get("user_id", pk).One(&writeBlogEveryWeek)
	return writeBlogEveryWeek
}

/**
 * ブログの必要記事数を更新する
 */
func UpdateRequireCount(configData config.ConfigData, allMemberData []WriteBlogEveryWeek, targetUserList map[string]int) {
	table := getTableObject(configData)
	for i := 0; i < len(allMemberData); i++ {
		findRequireCount := allMemberData[i].RequireCount
		currentRequireCount := targetUserList[allMemberData[i].UserID]
		if findRequireCount != currentRequireCount {
			// 食い違っている = 少なくとも1記事以上は書いているはず
			allMemberData[i].RequireCount = currentRequireCount
			err := table.Put(allMemberData[i]).Run()
			if err != nil {
				panic("データ保存エラー => " + err.Error())
			}
		}
	}
}

/**
 * ブログの必要記事数をリフレッシュする
 */
func ResetRequireCount(configData config.ConfigData, allMemberData []WriteBlogEveryWeek, targetUserList map[string]int) map[string]int {
	table := getTableObject(configData)
	for i := 0; i < len(allMemberData); i++ {
		// 0の人は1になり、1以上の人は1記事増える
		targetUserList[allMemberData[i].UserID]++
		allMemberData[i].RequireCount = targetUserList[allMemberData[i].UserID]
		err := table.Put(allMemberData[i]).Run()
		if err != nil {
			panic("データ保存エラー => " + err.Error())
		}
	}

	return targetUserList
}

/**
 * 新しいユーザーデータを作成する
 */
func CreateUser(configData config.ConfigData, slackParams *slack.SlackParams) {
	var writeBlogEveryWeek WriteBlogEveryWeek
	writeBlogEveryWeek.UserID = slackParams.UserID
	writeBlogEveryWeek.UserName = slackParams.UserName
	writeBlogEveryWeek.FeedURL = slackParams.Text
	writeBlogEveryWeek.RequireCount = 1
	table := getTableObject(configData)
	err := table.Put(writeBlogEveryWeek).Run()
	if err != nil {
		panic("登録エラー => " + err.Error())
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
