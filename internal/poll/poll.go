package poll

import (
	"../db"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"../syncLog"
)

func PollBot(formattedTime string, url string) {
	go pollParallel(formattedTime, url)
}

func pollParallel(formattedTime string, url string) {
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
	
	var wg sync.WaitGroup
	wg.Add(len(userIDs))

	syncLog.Println("[INFO]: Poll WaitGroup Results ========")
	for _, userID := range userIDs {
		go poll(&wg, url, userID, content)
	}

	wg.Wait()
	syncLog.Println("======================================")
}

func poll(wg *sync.WaitGroup, url string, userID string, content db.Content) {
	defer wg.Done()

	if content.Author != nil {
		content.Payload = content.Payload + "\n- " + content.Author
	}

	msg := fmt.Sprint(
`{
	"userID": "`, strconv.Itoa(userID),`",
	"contentID": "`, content.Id,`",
	"contentType": "`, content.Type,`",
	"payload": "`, content.Payload,`",
}`)
	jsonMsg := []byte(msg)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

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
	if botResponse != "Poll received" {
		syncLog.Println("[ERROR]: Bot's response wasn't 'Poll received'")
		return
	}
}