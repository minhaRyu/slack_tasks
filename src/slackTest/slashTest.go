package slackTest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv(("SLACK_SLASH_VERIFICATION_TOKEN"))) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/todo":
		params := &slack.Msg{Text: s.Text}
		response := fmt.Sprintf("You asked for the weather for %v", params.Text)
		w.Write([]byte(response))

	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ListenAndServe() {
	http.HandleFunc("/todo", slashCommandHandler)

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)
}
