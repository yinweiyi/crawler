package engine

import (
	"log"
)

type SimpleEngine struct {}

func (s SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching Url: %s", r.Url)
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("get item %v", item)
		}

	}
}


