package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"udemy.com/serverless/code/chatter/chatsess"
	"udemy.com/serverless/code/chatter/chatsess/usersess"
)

type Event struct {
	Sessid string
	Text   string
}

type Response struct {
	Job string
	Err string
}

func handler(c context.Context, ev Event) (Response, error) {
	sess := session.Must(session.NewSession())

	lg, err := usersess.GetLogin(ev.Sessid, sess)

	if err != nil {
		return Response{
			Job: "Say " + ev.Text,
			Err: "Not Logged in:" + err.Error(),
		}, nil
	}

	ch := chatsess.NewChat(lg.Username, ev.Text)
	err = ch.Put(sess)
	if err != nil {
		return Response{
			Job: "Say " + ev.Text,
			Err: "Could Not:" + err.Error(),
		}, nil
	}

	return Response{
		Job: "Say " + ev.Text,
		Err: "",
	}, nil
}

func main() {
	lambda.Start(handler)
}
