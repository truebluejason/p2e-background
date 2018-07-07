package poll

import (
	"bytes"
	"github.com/truebluejason/p2e-background/internal/conf"
	"github.com/truebluejason/p2e-background/internal/db"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"github.com/truebluejason/p2e-background/internal/syncLog"
)

type Payload struct {
	UserID string 		`json:"userID"`
	ContentID string 	`json:"contentID"`
	ContentType string 	`json:"contentType"`
	Payload string 		`json:"payload"`
}

func PollBot(formattedTime string) {
	pollParallel(formattedTime) // could've used goroutine for this
}

func pollParallel(formattedTime string) {
	userIDs, err := db.GetUsersFromTime(formattedTime)
	if err != nil {
		syncLog.Println(err.Error())
		return
	}
	if len(userIDs) == 0 {
		return
	}

	content, err := db.GetRandomContent()
	if err != nil {
		syncLog.Println(err.Error())
		return
	}

	if !botIsUp() {
		syncLog.Println("[ERROR]: P2E-Bot seems to be down.")
		return
	}
	
	var wg sync.WaitGroup
	wg.Add(len(userIDs))

	syncLog.Println("[INFO]: Poll WaitGroup Results========")
	for _, userID := range userIDs {
		go poll(&wg, userID, content)
	}

	wg.Wait()
	syncLog.Println("======================================")
}

func poll(wg *sync.WaitGroup, userID string, content db.Content) {
	defer wg.Done()

	if content.Author != "" {
		content.Payload = content.Payload + "\n- " + content.Author
	}

	url := conf.Configs.BotURL
	msg := Payload{userID, strconv.Itoa(content.Id), content.Type, content.Payload}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		syncLog.Println("[ERROR]: " + err.Error())
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonMsg))

	if err != nil {
		syncLog.Println("[ERROR]: " + err.Error())
		return
	}

	defer resp.Body.Close()
	botResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		syncLog.Println("[ERROR]: " + err.Error())
		return
	}
	if string(botResponse) != "Poll received" {
		syncLog.Println("[ERROR]: Bot's response wasn't 'Poll received'")
		return
	}
	syncLog.Println("[INFO]: Polling successful for userID: " + userID + " :)")
}

func botIsUp() bool {
	url := conf.Configs.BotURL
	baseUrl := strings.TrimSuffix(url, "poll")

	syncLog.Println("Pinging " + baseUrl)

	resp, err := http.Get(baseUrl)
	if err != nil {
		syncLog.Println("[ERROR]: " + err.Error())
		return false
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil || string(contents) != "hello world" {
		syncLog.Println("string content is: " + string(contents))
		return false
	}
	return true
}