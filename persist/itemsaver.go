package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver:Got item #%d : %v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("Item saver: error saving item %v : %v", item, err)
				continue
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("Must supply type")
	}

	indexService := client.Index().Index(index).Type(item.Type).Id(item.Id).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)
	return nil
}
