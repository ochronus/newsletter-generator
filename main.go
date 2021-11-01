package main

import (
	"fmt"
	"io/ioutil"

	"github.com/motemen/go-pocket/api"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func getAllItems(client *api.Client, ai *authInfo) {
	return
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

		res, err := AddTags(client, ai, item.ItemID, []string{"yolo"})
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}

func traverseAST(node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		if node.Kind() == ast.KindLink {
			n := node.(*ast.Link)
			fmt.Println(string(n.Destination))
		}
	}

	return ast.WalkContinue, nil
}

func main() {
	consumerKey := getConsumerKey()

	accessToken, err := restoreAccessToken(consumerKey)
	if err != nil {
		panic(err)
	}

	client := api.NewClient(consumerKey, accessToken.AccessToken)

	ai := &authInfo{
		ConsumerKey: consumerKey,
		AccessToken: accessToken.AccessToken,
	}
	getAllItems(client, ai)
	fname := "..\\newsletter.leadership.garden\\content\\posts\\2021-09-10-waldeinsamkeit.md"
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	doc := goldmark.DefaultParser().Parse(text.NewReader(b))

	err = ast.Walk(doc, traverseAST)
	if err != nil {
		panic(err)
	}
}
