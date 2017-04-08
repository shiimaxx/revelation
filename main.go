package revelation

import (
	"os"

	"github.com/pkg/errors"

	"github.com/shiimaxx/pocket"
)

var (
	pocketKey   = os.Getenv("POCKET_CONSUMER_KEY")
	pocketToken = os.Getenv("POCKET_ACCESS_TOKEN")
)

type PostItem struct {
	Title string
	URL   string
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
