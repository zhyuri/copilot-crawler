package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/op/go-logging"
	"github.com/zhyuri/copilot-crawler/lib"
)

var log *logging.Logger

type NewRepoEvent struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		log.Infof("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		var event NewRepoEvent
		err := json.Unmarshal([]byte(snsRecord.Message), &event)
		if err != nil {
			log.Infof("error: %v", err)
			continue
		}
		log.Infof("get new repo name: %s", event.Name)
		repo, err := lib.NewGithubRepo(event.Name, event.Owner)
		if err != nil {
			log.Warningf("cannot fetch new repo %s, got err %s", event.Name, err)
			continue
		}
		log.Infof("get new repo %v", repo)
	}
}

func main() {
	log = lib.Log
	lambda.Start(handler)
}
