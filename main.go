package revelation

import (
	"fmt"
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
		Count: 5,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get items")
	}

	var PostItems []PostItem
	for _, item := range items.List {
		PostItems = append(PostItems, PostItem{
			Title: item.ResolvedTitle,
			URL:   item.ResolvedURL})
	}

	return PostItems, nil
}
