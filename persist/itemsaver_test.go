package persist

import (
	"context"
	"crawler/engine"
	"crawler/models"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expect := engine.Item{
		Id:   "3903982005871861481",
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/3903982005871861481",
		Type: "zhenai",
		Payload: models.Profile{
			Name:       "wayee",
			Gender:     "男",
			Age:        28,
			Height:     173,
			Weight:     52,
			Income:     "1000-2000",
			Marriage:   "已婚",
			Education:  "大专",
			Occupation: "程序",
			Hokou:      "四川中江",
			Xinzuo:     "狮子",
			House:      "有房",
			Car:        "有车",
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	const index = "dating_test"
	err = save(client, index, expect)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index(index).Type(expect.Type).Id(expect.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item

	json.Unmarshal(resp.Source, &actual)

	if err != nil {
		panic(err)
	}

	actualProfile, _ := models.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if actual != expect {
		t.Errorf("got %v; expected %v", actual, expect)
	}

}
