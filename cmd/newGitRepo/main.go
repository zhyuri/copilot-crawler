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

func init() {
	log = logging.MustGetLogger("cli-newGitRepo")
	logging.SetLevel(logging.INFO, "newGitRepo")
	logging.SetLevel(logging.INFO, "lib")
}

func handler(ctx context.Context, snsEvent events.SNSEvent) {
	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		log.Debugf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
		var event NewRepoEvent
		err := json.Unmarshal([]byte(snsRecord.Message), &event)
		if err != nil {
			log.Errorf("parse sns event json got error: %v", err)
			continue
		}
		log.Debugf("get new repo name: %s of owner %s", event.Name, event.Owner)
		repo, err := lib.NewGithubRepo(event.Name, event.Owner)
		if err != nil {
			log.Errorf("cannot fetch new repo %s, got err %s", event.Name, err)
			continue
		}
		log.Infof("get new repo %v", repo)
	}
}

func main() {
	lambda.Start(handler)
}
