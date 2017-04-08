package revelation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/pkg/errors"

	"github.com/shiimaxx/pocket"
)

var (
	pocketKey   = os.Getenv("POCKET_CONSUMER_KEY")
	pocketToken = os.Getenv("POCKET_ACCESS_TOKEN")
	slackURL    = os.Getenv("SLACK_WEBHOOK_URL")
)

type PostItem struct {
	Title string
	URL   string
}

type slackInput struct {
	Text string `json:"text"`
}

func ToSlack(postItems []PostItem) error {
	client := new(http.Client)

	postText := "本日のキツネ様からのお告げです。\n\n"
	for i := range postItems {
		postText += postItems[i].Title + "\n" + postItems[i].URL + "\n\n"
	}
	jsonData, err := json.Marshal(&slackInput{Text: postText})
	if err != nil {
		return errors.Wrap(err, "failed to parse json")
	}

	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return errors.Wrap(err, "failed to create http request")
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to http request")
	}
	return nil
}

func Random() ([]PostItem, error) {
	client, err := pocket.NewClient(pocketKey, pocketToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create pocket client")
	}

	items, err := client.Retrieve(&pocket.RetrieveOpts{
		State: "unread",
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get items")
	}

	var PostItems []PostItem
	counter := 0
	for _, item := range items.List {
		if counter < 5 {
			PostItems = append(PostItems, PostItem{
				Title: item.ResolvedTitle,
				URL:   item.ResolvedURL})
		}
		counter++
	}

	return PostItems, nil
}
