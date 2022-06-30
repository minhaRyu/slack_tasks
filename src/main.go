package main

import (
	"log"
	"sample-app/src/slackTest"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// slackTest 샘플 코드
	slackTest.SendMessage("C03LLF4EQQ4", "테스트 메시지")

	t := time.Now()
	curYear, curMonth, curDay := t.Date()
	todayOnTime := time.Date(curYear, curMonth, curDay, 0, 0, 0, 0, t.Location())
	slackTest.GetConversationHistory("C03LLF4EQQ4", todayOnTime)

	slackTest.GetUsersInConversation("C03LLF4EQQ4")

	slackTest.GetUsers()
	slackTest.ListenAndServe()


	/*
	// webdriverTest 샘플 코드
	webdriverTest.Test()

	*/

	/*
	// google tasks 샘플 코드
	tasksTest.Insert()

	*/
}
