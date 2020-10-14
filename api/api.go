package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pec1985/go-clubhouse/api/models"
)

// API the api interface
type API interface {
	// Request you can use this method to call any api that's missing from the sdk
	Request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error
	// ListCategories returns a list of all Categories and their attributes.
	ListCategories() (*[]models.Category, error)
	// CreateCategory allows you to create a new Category in Clubhouse.
	CreateCategory(category *models.CreateCategory) error
	// DeleteCategory can be used to delete any Category.
	DeleteCategory(categoryID int64) error
	// GetCategory returns information about the selected Category.
	GetCategory(categoryID int64) (*models.Category, error)
	// UpdateCategory allows you to replace a Category name with another name. If you try to name a Category something that already exists, you will receive a 422 response.
	UpdateCategory(categoryID int64, category *models.UpdateCategory) (*models.Category, error)
	// ListCategoryMilestones returns a list of all Milestones with the Category.
	ListCategoryMilestones(categoryID int64) (*[]models.Milestone, error)
	// ListEntityTemplates List all the entity templates for an organization.
	ListEntityTemplates() (*[]models.EntityTemplate, error)
	// CreateEntityTemplate Create a new entity template for your organization.
	CreateEntityTemplate(entityTemplate *models.CreateEntityTemplate) error
	// DisableStoryTemplates Disables the Story Template feature for the given Organization.
	DisableStoryTemplates() error
	// EnableStoryTemplates Enables the Story Template feature for the given Organization.
	EnableStoryTemplates() error
	DeleteEntityTemplate(entityTemplateID string) error
	// GetEntityTemplate returns information about a given entity template.
	GetEntityTemplate(entityTemplateID string) (*models.EntityTemplate, error)
	// UpdateEntityTemplate Update an entity template's name or its contents.
	UpdateEntityTemplate(entityTemplateID string, entityTemplate *models.UpdateEntityTemplate) (*models.EntityTemplate, error)
	// GetEpicWorkflow returns the Epic Workflow for the organization.
	GetEpicWorkflow() (*models.EpicWorkflow, error)
	// ListEpics returns a list of all Epics and their attributes.
	ListEpics(listEpics *models.ListEpics) (*[]models.EpicSlim, error)
	// CreateEpic allows you to create a new Epic in Clubhouse.
	CreateEpic(epic *models.CreateEpic) error
	// DeleteEpic can be used to delete the Epic. The only required parameter is Epic ID.
	DeleteEpic(epicID int64) error
	// GetEpic returns information about the selected Epic.
	GetEpic(epicID int64) (*models.Epic, error)
	// UpdateEpic can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Clubhouse UI.
	UpdateEpic(epicID int64, epic *models.UpdateEpic) (*models.Epic, error)
	// ListEpicComments Get a list of all Comments on an Epic.
	ListEpicComments(epicID int64) (*[]models.ThreadedComment, error)
	// CreateEpicComment This endpoint allows you to create a threaded Comment on an Epic.
	CreateEpicComment(epicID int64, epicComment *models.CreateEpicComment) error
	// DeleteEpicComment This endpoint allows you to delete a Comment from an Epic.
	DeleteEpicComment(epicID int64, commentID int64) error
	// GetEpicComment This endpoint returns information about the selected Epic Comment.
	GetEpicComment(epicID int64, commentID int64) (*models.ThreadedComment, error)
	// CreateEpicCommentComment This endpoint allows you to create a nested Comment reply to an existing Epic Comment.
	CreateEpicCommentComment(epicID int64, commentID int64, commentComment *models.CreateCommentComment) error
	// UpdateEpicComment This endpoint allows you to update a threaded Comment on an Epic.
	UpdateEpicComment(epicID int64, commentID int64, comment *models.UpdateComment) (*models.ThreadedComment, error)
	// ListEpicStories Get a list of all Stories in an Epic.
	ListEpicStories(epicID int64, getEpicStories *models.GetEpicStories) (*[]models.StorySlim, error)
	// ListFiles returns a list of all Files and related attributes in your Clubhouse.
	ListFiles() (*[]models.File, error)
	CreateFiles(files *models.CreateFiles) error
	// DeleteFile can be used to delete any previously attached File.
	DeleteFile(fileID int64) error
	// GetFile returns information about the selected File.
	GetFile(fileID int64) (*models.File, error)
	// UpdateFile can used to update the properties of a file uploaded to Clubhouse.
	UpdateFile(fileID int64, file *models.UpdateFile) (*models.File, error)
	ListGroups() (*[]models.Group, error)
	CreateGroup(group *models.CreateGroup) error
	// DisableGroups Disables Groups for the current workspace2
	DisableGroups() error
	// EnableGroups Enables Groups for the current workspace2
	EnableGroups() error
	GetGroup(groupID string) (*models.Group, error)
	UpdateGroup(groupID string, group *models.UpdateGroup) (*models.Group, error)
	ListIterations() (*[]models.IterationSlim, error)
	CreateIteration(iteration *models.CreateIteration) error
	// DisableIterations Disables Iterations for the current workspace
	DisableIterations() error
	// EnableIterations Enables Iterations for the current workspace
	EnableIterations() error
	DeleteIteration(iterationID int64) error
	GetIteration(iterationID int64) (*models.Iteration, error)
	UpdateIteration(iterationID int64, iteration *models.UpdateIteration) (*models.Iteration, error)
	// ListIterationStories Get a list of all Stories in an Iteration.
	ListIterationStories(iterationID int64, getIterationStories *models.GetIterationStories) (*[]models.StorySlim, error)
	// ListLabels returns a list of all Labels and their attributes.
	ListLabels(listLabels *models.ListLabels) (*[]models.Label, error)
	// CreateLabel allows you to create a new Label in Clubhouse.
	CreateLabel(labelParams *models.CreateLabelParams) error
	// DeleteLabel can be used to delete any Label.
	DeleteLabel(labelID int64) error
	// GetLabel returns information about the selected Label.
	GetLabel(labelID int64) (*models.Label, error)
	// UpdateLabel allows you to replace a Label name with another name. If you try to name a Label something that already exists, you will receive a 422 response.
	UpdateLabel(labelID int64, label *models.UpdateLabel) (*models.Label, error)
	// ListLabelEpics List all of the Epics with the Label.
	ListLabelEpics(labelID int64) (*[]models.EpicSlim, error)
	// ListLabelStories List all of the Stories with the Label.
	ListLabelStories(labelID int64, getLabelStories *models.GetLabelStories) (*[]models.StorySlim, error)
	// ListLinkedFiles returns a list of all Linked-Files and their attributes.
	ListLinkedFiles() (*[]models.LinkedFile, error)
	// CreateLinkedFile allows you to create a new Linked File in Clubhouse.
	CreateLinkedFile(linkedFile *models.CreateLinkedFile) error
	// DeleteLinkedFile can be used to delete any previously attached Linked-File.
	DeleteLinkedFile(linkedFileID int64) error
	// GetLinkedFile Get File returns information about the selected Linked File.
	GetLinkedFile(linkedFileID int64) (*models.LinkedFile, error)
	// UpdateLinkedFile Updated Linked File allows you to update properties of a previously attached Linked-File.
	UpdateLinkedFile(linkedFileID int64, linkedFile *models.UpdateLinkedFile) (*models.LinkedFile, error)
	// GetCurrentMemberInfo Returns information about the authenticated member.
	GetCurrentMemberInfo() (*models.MemberInfo, error)
	// ListMembers returns information about members of the organization.
	ListMembers(listMembers *models.ListMembers) (*[]models.Member, error)
	// GetMember Returns information about a Member.
	GetMember(memberID string, getMember *models.GetMember) (*models.Member, error)
	// ListMilestones returns a list of all Milestones and their attributes.
	ListMilestones() (*[]models.Milestone, error)
	// CreateMilestone allows you to create a new Milestone in Clubhouse.
	CreateMilestone(milestone *models.CreateMilestone) error
	// DeleteMilestone can be used to delete any Milestone.
	DeleteMilestone(milestoneID int64) error
	// GetMilestone returns information about a chosen Milestone.
	GetMilestone(milestoneID int64) (*models.Milestone, error)
	// UpdateMilestone can be used to update Milestone properties.
	UpdateMilestone(milestoneID int64, milestone *models.UpdateMilestone) (*models.Milestone, error)
	// ListMilestoneEpics List all of the Epics within the Milestone.
	ListMilestoneEpics(milestoneID int64) (*[]models.EpicSlim, error)
	// ListProjects returns a list of all Projects and their attributes.
	ListProjects() (*[]models.Project, error)
	// CreateProject is used to create a new Clubhouse Project.
	CreateProject(project *models.CreateProject) error
	// DeleteProject can be used to delete a Project. Projects can only be deleted if all associated Stories are moved or deleted. In the case that the Project cannot be deleted, you will receive a 422 response.
	DeleteProject(projectID int64) error
	// GetProject returns information about the selected Project.
	GetProject(projectID int64) (*models.Project, error)
	// UpdateProject can be used to change properties of a Project.
	UpdateProject(projectID int64, project *models.UpdateProject) (*models.Project, error)
	// ListStories returns a list of all Stories in a selected Project and their attributes.
	ListStories(projectID int64, getProjectStories *models.GetProjectStories) (*[]models.StorySlim, error)
	// ListRepositories returns a list of all Repositories and their attributes.
	ListRepositories() (*[]models.Repository, error)
	// GetRepository returns information about the selected Repository.
	GetRepository(repoID int64) (*models.Repository, error)
	// Search lets you search Epics and Stories based on desired parameters. Since ordering of the results can change over time (due to search ranking decay, new Epics and Stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	Search(search *models.Search) (*models.SearchResults, error)
	// SearchEpics lets you search Epics based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new Epics being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	SearchEpics(search *models.Search) (*models.EpicSearchResults, error)
	// SearchStories lets you search Stories based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	SearchStories(search *models.Search) (*models.StorySearchResults, error)
	// CreateStory is used to add a new story to your Clubhouse.
	CreateStory(storyParams *models.CreateStoryParams) error
	// DeleteMultipleStories allows you to delete multiple archived stories at once.
	DeleteMultipleStories(stories *models.DeleteStories) error
	// CreateMultipleStories allows you to create multiple stories in a single request using the same syntax as [Create Story](https://clubhouse.io/api/#create-story).
	CreateMultipleStories(stories *models.CreateStories) error
	// UpdateMultipleStories allows you to make changes to numerous stories at once.
	UpdateMultipleStories(stories *models.UpdateStories) (*[]models.StorySlim, error)
	// SearchStoriesOld Search Stories lets you search Stories based on desired parameters.
	SearchStoriesOld(searchStories *models.SearchStories) error
	// DeleteStory can be used to delete any Story.
	DeleteStory(storyID int64) error
	// GetStory returns information about a chosen Story.
	GetStory(storyID int64) (*models.Story, error)
	// UpdateStory can be used to update Story properties.
	UpdateStory(storyID int64, story *models.UpdateStory) (*models.Story, error)
	// CreateComment allows you to create a Comment on any Story.
	CreateComment(storyID int64, comment *models.CreateComment) error
	// DeleteComment Delete a Comment from any story.
	DeleteComment(storyID int64, commentID int64) error
	// GetComment is used to get Comment information.
	GetComment(storyID int64, commentID int64) (*models.Comment, error)
	// UpdateComment replaces the text of the existing Comment.
	UpdateComment(storyID int64, commentID int64, comment *models.UpdateComment) (*models.Comment, error)
	// DeleteReaction Delete a Reaction from any comment.
	DeleteReaction(storyID int64, commentID int64, orDeleteReaction *models.CreateOrDeleteReaction) error
	// CreateReaction Create a reaction to a comment.
	CreateReaction(storyID int64, commentID int64, orDeleteReaction *models.CreateOrDeleteReaction) error
	// CreateTask is used to create a new task in a Story.
	CreateTask(storyID int64, task *models.CreateTask) error
	// DeleteTask can be used to delete any previously created Task on a Story.
	DeleteTask(storyID int64, taskID int64) error
	// GetTask Returns information about a chosen Task.
	GetTask(storyID int64, taskID int64) (*models.Task, error)
	// UpdateTask can be used to update Task properties.
	UpdateTask(storyID int64, taskID int64, task *models.UpdateTask) (*models.Task, error)
	// CreateStoryLink Story Links (called Story Relationships in the UI) allow you create semantic relationships between two stories. The parameters read like an active voice grammatical sentence:  subject -> verb -> object.
	//
	// The subject story acts on the object Story; the object story is the direct object of the sentence.
	//
	// The subject story "blocks", "duplicates", or "relates to" the object story.  Examples:
	// - "story 5 blocks story 6” -- story 6 is now "blocked" until story 5 is moved to a Done workflow state.
	// - "story 2 duplicates story 1” -- Story 2 represents the same body of work as Story 1 (and should probably be archived).
	// - "story 7 relates to story 3”
	CreateStoryLink(storyLink *models.CreateStoryLink) error
	// DeleteStoryLink Removes the relationship between the stories for the given Story Link.
	DeleteStoryLink(storyLinkID int64) error
	// GetStoryLink Returns the stories and their relationship for the given Story Link.
	GetStoryLink(storyLinkID int64) (*models.StoryLink, error)
	// UpdateStoryLink Updates the stories and/or the relationship for the given Story Link.
	UpdateStoryLink(storyLinkID int64, storyLink *models.UpdateStoryLink) (*models.StoryLink, error)
	// ListTeams returns a list of all Teams in the organization.
	ListTeams() (*[]models.Team, error)
	// GetTeam is used to get Team information.
	GetTeam(teamID int64) (*models.Team, error)
	// ListWorkflows returns a list of all Workflows in the organization.
	ListWorkflows() (*[]models.Workflow, error)
}

// New creates a new instance of the api
func New(client *http.Client, token string) API {
	return &api{
		client: client,
		token:  token,
		url:    "https://api.clubhouse.io/",
	}
}

type api struct {
	client *http.Client
	token  string
	url    string
}

// Request you can use this method to call any api that's missing from the sdk
func (a *api) Request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error {

	if params == nil {
		params = url.Values{}
	}
	ur := strings.TrimRight(a.url, "/")
	endpoint = strings.TrimPrefix(endpoint, "/")
	var req *http.Request
	var err error
	if data == nil {
		req, err = http.NewRequest(http.MethodGet, ur+"/"+endpoint+"?"+params.Encode(), nil)
	} else {
		req, err = http.NewRequest(http.MethodGet, ur+"/"+endpoint+"?"+params.Encode(), data)
	}
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Clubhouse-Token", a.token)
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("status code: %v. message: %v", resp.StatusCode, string(b))
}
