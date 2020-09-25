package models

import (
	"encoding/json"
	"time"
)

type StorySlim struct {
	// The Clubhouse application url for the Story.
	AppUrl string `json:"app_url"`
	// True if the story has been archived or not.
	Archived bool `json:"archived"`
	// A true/false boolean indicating if the Story is currently blocked.
	Blocked bool `json:"blocked"`
	// A true/false boolean indicating if the Story is currently a blocker of another story.
	Blocker bool `json:"blocker"`
	// An array of IDs of Comments attached to the story.
	CommentIds []int64 `json:"comment_ids"`
	// A true/false boolean indicating if the Story has been completed.
	Completed bool `json:"completed"`
	// The time/date the Story was completed.
	CompletedAt *time.Time `json:"completed_at"`
	// A manual override for the time/date the Story was completed.
	CompletedAtOverride *time.Time `json:"completed_at_override"`
	// The time/date the Story was created.
	CreatedAt time.Time `json:"created_at"`
	// The cycle time (in seconds) of this story when complete.
	CycleTime int64 `json:"cycle_time"`
	// The due date of the story.
	Deadline *time.Time `json:"deadline"`
	// The description of the Story.
	Description string `json:"description"`
	// A string description of this resource.
	EntityType string `json:"entity_type"`
	// The ID of the epic the story belongs to.
	EpicId *int64 `json:"epic_id"`
	// The numeric point estimate of the story. Can also be null, which means unestimated.
	Estimate *int64 `json:"estimate"`
	// This field can be set to another unique ID. In the case that the Story has been imported from another tool, the ID in the other tool can be indicated here.
	ExternalId *string `json:"external_id"`
	// An array of external link (strings) associated with a Story
	ExternalLinks []string `json:"external_links"`
	// An array of External Tickets associated with a Story
	ExternalTickets []ExternalTicket `json:"external_tickets"`
	// An array of IDs of Files attached to the story.
	FileIds []int64 `json:"file_ids"`
	// An array of UUIDs for any Members listed as Followers.
	FollowerIds []string `json:"follower_ids"`
	// The ID of the group associated with the story.
	GroupId *string `json:"group_id"`
	// An array of Group IDs that have been mentioned in the Story description.
	GroupMentionIds []string `json:"group_mention_ids"`
	// The unique ID of the Story.
	Id int64 `json:"id"`
	// The ID of the iteration the story belongs to.
	IterationId *int64 `json:"iteration_id"`
	// An array of labels attached to the story.
	Labels []Label `json:"labels"`
	// The lead time (in seconds) of this story when complete.
	LeadTime int64 `json:"lead_time"`
	// An array of IDs of LinkedFiles attached to the story.
	LinkedFileIds []int64 `json:"linked_file_ids"`
	// An array of Member IDs that have been mentioned in the Story description.
	MemberMentionIds []string `json:"member_mention_ids"`
	// Deprecated: use member_mention_ids.
	MentionIds []string `json:"mention_ids"`
	// The time/date the Story was last changed workflow-state.
	MovedAt *time.Time `json:"moved_at"`
	// The name of the story.
	Name string `json:"name"`
	// The number of tasks on the story which are complete.
	NumTasksCompleted int64 `json:"num_tasks_completed"`
	// An array of UUIDs of the owners of this story.
	OwnerIds []string `json:"owner_ids"`
	// A number representing the position of the story in relation to every other story in the current project.
	Position int64 `json:"position"`
	// The IDs of the iteration the story belongs to.
	PreviousIterationIds []int64 `json:"previous_iteration_ids"`
	// The ID of the project the story belongs to.
	ProjectId int64 `json:"project_id"`
	// The ID of the Member that requested the story.
	RequestedById string `json:"requested_by_id"`
	// A true/false boolean indicating if the Story has been started.
	Started bool `json:"started"`
	// The time/date the Story was started.
	StartedAt *time.Time `json:"started_at"`
	// A manual override for the time/date the Story was started.
	StartedAtOverride *time.Time `json:"started_at_override"`
	Stats             StoryStats `json:"stats"`
	// An array of story links attached to the Story.
	StoryLinks []TypedStoryLink `json:"story_links"`
	// The type of story (feature, bug, chore).
	StoryType      string          `json:"story_type"`
	SupportTickets []SupportTicket `json:"support_tickets"`
	// An array of IDs of Tasks attached to the story.
	TaskIds []int64 `json:"task_ids"`
	// The time/date the Story was updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// The ID of the workflow state the story is currently in.
	WorkflowStateId int64 `json:"workflow_state_id"`
}

func (m *StorySlim) Stringify() string {
	b, _ := json.Marshal(m)
	return string(b)
}
func (m *StorySlim) StringifyPretty() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}
