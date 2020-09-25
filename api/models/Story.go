package models

import (
	"encoding/json"
	"time"
)

type Story struct {
	// The description of the story.
	Description string `json:"description"`
	// The ID of the group associated with the story.
	GroupId *string `json:"group_id"`
	// The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time"`
	// The time/date the Story was last changed workflow-state.
	MovedAt        *time.Time      `json:"moved_at"`
	SupportTickets []SupportTicket `json:"support_tickets"`
	// An array of commits attached to the story.
	Commits []Commit `json:"commits"`
	// The time/date the Story was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// An array of external link (strings) associated with a Story
	ExternalLinks []string `json:"external_links"`
	// An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The name of the story.
	Name string `json:"name"`
	// A number representing the position of the story in relation to every other story in the current project.
	Position int64      `json:"position"`
	Stats    StoryStats `json:"stats"`
	// An array of tasks connected to the story.
	Tasks []Task `json:"tasks"`
	// The ID of the workflow state the story is currently in.
	WorkflowStateId int64 `json:"workflow_state_id"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// An array of External Tickets associated with a Story
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id"`
	// The ID of the Member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	// True if the story has been archived or not.
	Archived bool `json:"archived"`
	// The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// The unique ID of the Story.
	Id int64 `json:"id"`
	// An array of files attached to the story.
	Files []File `json:"files"`
	// An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// An array of labels attached to the story.
	Labels []Label `json:"labels"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// A true/false boolean indicating if the Story is currently a blocker of another story.
	Blocker bool `json:"blocker"`
	// An array of Git branches attached to the story.
	Branches []Branch `json:"branches"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id"`
	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// The IDs of the iteration the story belongs to.
	PreviousIterationIds []int64 `json:"previous_iteration_ids"`
	// An array of Pull/Merge Requests attached to the story.
	PullRequests []PullRequest `json:"pull_requests"`
	// A true/false boolean indicating if the Story has been started.
	Started bool `json:"started"`
	// The type of story (feature, bug, chore).
	StoryType string `json:"story_type"`
	// A true/false boolean indicating if the Story has been completed.
	Completed bool `json:"completed"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// An array of linked files attached to the story.
	LinkedFiles []LinkedFile `json:"linked_files"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// An array of story links attached to the Story.
	StoryLinks []TypedStoryLink `json:"story_links"`
	// The time/date the Story was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The Clubhouse application url for the Story.
	AppUrl string `json:"app_url"`
	// A true/false boolean indicating if the Story is currently blocked.
	Blocked bool `json:"blocked"`
	// An array of comments attached to the story.
	Comments []Comment `json:"comments"`
	// The time/date the Story was started.
	StartedAt *time.Time `json:"started_at"`
}

func (m *Story) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *Story) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
