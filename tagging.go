package main

import (
	"strings"

	"github.com/motemen/go-pocket/api"
)

type authInfo struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// Action represents one action in a bulk modify requests.
type AddTagsAction struct {
	Action string `json:"action"`
	ItemID int    `json:"item_id,string"`
	Tags   string `json:"tags"`
}

// NewArchiveAction creates an acrhive action.
func NewAddTagsAction(itemID int, tags []string) *AddTagsAction {
	return &AddTagsAction{
		Action: "tags_add",
		ItemID: itemID,
		Tags:   strings.Join(tags, ","),
	}
}

// ModifyResult represents the modify API's result.
type ModifyResult struct {
	// The results for each of the requested actions.
	ActionResults []bool
	Status        int
}

type addTagsActionWithAuth struct {
	Actions []*AddTagsAction `json:"actions"`
	authInfo
}

func AddTags(c *api.Client, authInfo *authInfo, itemID int, tags []string) (*ModifyResult, error) {
	res := &ModifyResult{}
	action := NewAddTagsAction(itemID, tags)
	data := addTagsActionWithAuth{
		authInfo: *authInfo,
		Actions:  []*AddTagsAction{action},
	}
	err := api.PostJSON("/v3/send", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
