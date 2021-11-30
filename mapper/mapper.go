package mapper

import (
	"context"
	"fmt"

	"github.com/ONSdigital/dp-frontend-area-profiles/config"
)

// TODO: remove hello world example mapper and models

type HelloModel struct {
	Greeting string `json:"greeting"`
	Who      string `json:"who"`
}

type HelloWorldModel struct {
	HelloWho string `json:"hello-who"`
}

func HelloWorld(ctx context.Context, hm HelloModel, cfg config.Config) HelloWorldModel {
	var hwm HelloWorldModel
	hwm.HelloWho = fmt.Sprintf("%s %s", hm.Greeting, hm.Who)
	return hwm
}
