package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.Concurrent{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

}
