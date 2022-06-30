package slackTest

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

// 메시지 보내기
func SendMessage(channelId string, msg string) {
	api := slack.New(os.Getenv("SLACK_APP_OAUTH_TOKEN"))
	channelID, timestamp, err := api.PostMessage(
		channelId, // ex. 채널 세부정보 -> 아래쪽에 채널 ID
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(false), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("메시지 보내기\nchannelID: %s timestamp: %s\n", channelID, timestamp)
}

// 대화 정보 들고 오기
func GetConversationHistory(channelId string, Oldest time.Time) *slack.GetConversationHistoryResponse {
	api := slack.New(os.Getenv("SLACK_APP_OAUTH_TOKEN"))
	historyParameters := slack.GetConversationHistoryParameters{
		ChannelID: channelId,
		Oldest:    strconv.FormatInt(Oldest.Unix(), 10),
	}
	history, err := api.GetConversationHistory(&historyParameters)

	if err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}

	fmt.Println("대화 정보 들고 오기")
	for index, msg := range (*history).Messages {
		fmt.Println(index, msg)
	}

	return history
}

// 채널의 멤버 정보 들고 오기
func GetUsersInConversation(channelId string) []string {
	api := slack.New(os.Getenv("SLACK_APP_OAUTH_TOKEN"))
	memberParameters := slack.GetUsersInConversationParameters{
		ChannelID: channelId,
	}

	members, _, _ := api.GetUsersInConversation(&memberParameters)

	fmt.Println("채널의 멤버 정보 들고 오기\n", members)

	return members
}

// 슬랙팀의 전체 유저 들고 오기
func GetUsers() []slack.User {
	api := slack.New(os.Getenv("SLACK_APP_OAUTH_TOKEN"))
	users, _ := api.GetUsers()

	fmt.Println("슬랙팀의 전체 유저 들고 오기\n", users)

	return users
}
