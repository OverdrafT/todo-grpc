package main

import (
	"context"
	"log"

	"github.com/olezhek28/microservices_course/week_4/internal/app"
)

func main() {
	ctx := context.Background()

	myapp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = myapp.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
