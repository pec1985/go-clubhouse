package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/pec1985/go-clubhouse/api/models"
)

// Story Links (called Story Relationships in the UI) allow you create semantic relationships between two stories. The parameters read like an active voice grammatical sentence:  subject -> verb -> object.
//
// The subject story acts on the object Story; the object story is the direct object of the sentence.
//
// The subject story "blocks", "duplicates", or "relates to" the object story.  Examples:
// - "story 5 blocks story 6” -- story 6 is now "blocked" until story 5 is moved to a Done workflow state.
// - "story 2 duplicates story 1” -- Story 2 represents the same body of work as Story 1 (and should probably be archived).
// - "story 7 relates to story 3”
func (a *api) CreateStoryLink(createStoryLink *models.CreateStoryLink) error {
	params := url.Values{}
	var body *bytes.Buffer
	if createStoryLink != nil {
		jsonbody, _ := json.Marshal(createStoryLink)
		body = bytes.NewBuffer(jsonbody)
	}
	var out interface{}
	if err := a.request("POST", "/api/v3/story-links", params, body, &out); err != nil {
		return err
	}
	return nil
}

// Removes the relationship between the stories for the given Story Link.
func (a *api) DeleteStoryLink(storyLinkPublicId int64) error {
	params := url.Values{}
	var out interface{}
	if err := a.request("DELETE", "/api/v3/story-links/"+fmt.Sprint(storyLinkPublicId)+"", params, nil, &out); err != nil {
		return err
	}
	return nil
}

// Returns the stories and their relationship for the given Story Link.
func (a *api) GetStoryLink(storyLinkPublicId int64) (*models.StoryLink, error) {
	params := url.Values{}
	var out models.StoryLink
	if err := a.request("GET", "/api/v3/story-links/"+fmt.Sprint(storyLinkPublicId)+"", params, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Updates the stories and/or the relationship for the given Story Link.
func (a *api) UpdateStoryLink(storyLinkPublicId int64, updateStoryLink *models.UpdateStoryLink) (*models.StoryLink, error) {
	params := url.Values{}
	var body *bytes.Buffer
	if updateStoryLink != nil {
		jsonbody, _ := json.Marshal(updateStoryLink)
		body = bytes.NewBuffer(jsonbody)
	}
	var out models.StoryLink
	if err := a.request("PUT", "/api/v3/story-links/"+fmt.Sprint(storyLinkPublicId)+"", params, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
