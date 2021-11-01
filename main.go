package main

import (
	"fmt"

	"github.com/motemen/go-pocket/api"
)

func main() {
	consumerKey := getConsumerKey()

	accessToken, err := restoreAccessToken(consumerKey)
	if err != nil {
		panic(err)
	}

	client := api.NewClient(consumerKey, accessToken.AccessToken)
	res, err := client.Retrieve(&api.RetrieveOption{
		State:      "all",
		Tag:        "newsletter-material",
		Sort:       "newest",
		DetailType: "complete",
	})
	if err != nil {
		panic(err)
	}
	for _, item := range res.List {
		fmt.Println(item.GivenURL)
		for tag, _ := range item.Tags {
			fmt.Println("\t", tag)
		}
	}
}
