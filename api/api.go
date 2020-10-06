package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pec1985/go-clubhouse.io/v1/api/models"
)

// API the api interface
type API interface {
	// List Categories returns a list of all Categories and their attributes.
	ListCategories() (*[]models.Category, error)
	// Create Category allows you to create a new Category in Clubhouse.
	CreateCategory(createCategory *models.CreateCategory) error
	// Delete Category can be used to delete any Category.
	DeleteCategory(categoryPublicId int64) error
	// Get Category returns information about the selected Category.
	GetCategory(categoryPublicId int64) (*models.Category, error)
	// Update Category allows you to replace a Category name with another name. If you try to name a Category something that already exists, you will receive a 422 response.
	UpdateCategory(categoryPublicId int64, updateCategory *models.UpdateCategory) (*models.Category, error)
	// List Category Milestones returns a list of all Milestones with the Category.
	ListCategoryMilestones(categoryPublicId int64) (*[]models.Milestone, error)
	// List all the entity templates for an organization.
	ListEntityTemplates() (*[]models.EntityTemplate, error)
	// Create a new entity template for your organization.
	CreateEntityTemplate(createEntityTemplate *models.CreateEntityTemplate) error
	// Disables the Story Template feature for the given Organization.
	DisableStoryTemplates() error
	// Enables the Story Template feature for the given Organization.
	EnableStoryTemplates() error
	DeleteEntityTemplate(entityTemplatePublicId string) error
	// Get Entity Template returns information about a given entity template.
	GetEntityTemplate(entityTemplatePublicId string) (*models.EntityTemplate, error)
	// Update an entity template's name or its contents.
	UpdateEntityTemplate(entityTemplatePublicId string, updateEntityTemplate *models.UpdateEntityTemplate) (*models.EntityTemplate, error)
	// Get Epic Workflow returns the Epic Workflow for the organization.
	GetEpicWorkflow() (*models.EpicWorkflow, error)
	// List Epics returns a list of all Epics and their attributes.
	ListEpics(listEpics *models.ListEpics) (*[]models.EpicSlim, error)
	// Create Epic allows you to create a new Epic in Clubhouse.
	CreateEpic(createEpic *models.CreateEpic) error
	// Delete Epic can be used to delete the Epic. The only required parameter is Epic ID.
	DeleteEpic(epicPublicId int64) error
	// Get Epic returns information about the selected Epic.
	GetEpic(epicPublicId int64) (*models.Epic, error)
	// Update Epic can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Clubhouse UI.
	UpdateEpic(epicPublicId int64, updateEpic *models.UpdateEpic) (*models.Epic, error)
	// Get a list of all Comments on an Epic.
	ListEpicComments(epicPublicId int64) (*[]models.ThreadedComment, error)
	// This endpoint allows you to create a threaded Comment on an Epic.
	CreateEpicComment(epicPublicId int64, createEpicComment *models.CreateEpicComment) error
	// This endpoint allows you to delete a Comment from an Epic.
	DeleteEpicComment(epicPublicId int64, commentPublicId int64) error
	// This endpoint returns information about the selected Epic Comment.
	GetEpicComment(epicPublicId int64, commentPublicId int64) (*models.ThreadedComment, error)
	// This endpoint allows you to create a nested Comment reply to an existing Epic Comment.
	CreateEpicCommentComment(epicPublicId int64, commentPublicId int64, createCommentComment *models.CreateCommentComment) error
	// This endpoint allows you to update a threaded Comment on an Epic.
	UpdateEpicComment(epicPublicId int64, commentPublicId int64, updateComment *models.UpdateComment) (*models.ThreadedComment, error)
	// Get a list of all Stories in an Epic.
	ListEpicStories(epicPublicId int64, getEpicStories *models.GetEpicStories) (*[]models.StorySlim, error)
	// List Files returns a list of all Files and related attributes in your Clubhouse.
	ListFiles() (*[]models.File, error)
	CreateFiles(createFiles *models.CreateFiles) error
	// Delete File can be used to delete any previously attached File.
	DeleteFile(filePublicId int64) error
	// Get File returns information about the selected File.
	GetFile(filePublicId int64) (*models.File, error)
	// Update File can used to update the properties of a file uploaded to Clubhouse.
	UpdateFile(filePublicId int64, updateFile *models.UpdateFile) (*models.File, error)
	ListGroups() (*[]models.Group, error)
	CreateGroup(createGroup *models.CreateGroup) error
	// Disables Groups for the current workspace2
	DisableGroups() error
	// Enables Groups for the current workspace2
	EnableGroups() error
	GetGroup(groupPublicId string) (*models.Group, error)
	UpdateGroup(groupPublicId string, updateGroup *models.UpdateGroup) (*models.Group, error)
	ListIterations() (*[]models.IterationSlim, error)
	CreateIteration(createIteration *models.CreateIteration) error
	// Disables Iterations for the current workspace
	DisableIterations() error
	// Enables Iterations for the current workspace
	EnableIterations() error
	DeleteIteration(iterationPublicId int64) error
	GetIteration(iterationPublicId int64) (*models.Iteration, error)
	UpdateIteration(iterationPublicId int64, updateIteration *models.UpdateIteration) (*models.Iteration, error)
	// Get a list of all Stories in an Iteration.
	ListIterationStories(iterationPublicId int64, getIterationStories *models.GetIterationStories) (*[]models.StorySlim, error)
	// List Labels returns a list of all Labels and their attributes.
	ListLabels(listLabels *models.ListLabels) (*[]models.Label, error)
	// Create Label allows you to create a new Label in Clubhouse.
	CreateLabel(createLabelParams *models.CreateLabelParams) error
	// Delete Label can be used to delete any Label.
	DeleteLabel(labelPublicId int64) error
	// Get Label returns information about the selected Label.
	GetLabel(labelPublicId int64) (*models.Label, error)
	// Update Label allows you to replace a Label name with another name. If you try to name a Label something that already exists, you will receive a 422 response.
	UpdateLabel(labelPublicId int64, updateLabel *models.UpdateLabel) (*models.Label, error)
	// List all of the Epics with the Label.
	ListLabelEpics(labelPublicId int64) (*[]models.EpicSlim, error)
	// List all of the Stories with the Label.
	ListLabelStories(labelPublicId int64, getLabelStories *models.GetLabelStories) (*[]models.StorySlim, error)
	// List Linked Files returns a list of all Linked-Files and their attributes.
	ListLinkedFiles() (*[]models.LinkedFile, error)
	// Create Linked File allows you to create a new Linked File in Clubhouse.
	CreateLinkedFile(createLinkedFile *models.CreateLinkedFile) error
	// Delete Linked File can be used to delete any previously attached Linked-File.
	DeleteLinkedFile(linkedFilePublicId int64) error
	// Get File returns information about the selected Linked File.
	GetLinkedFile(linkedFilePublicId int64) (*models.LinkedFile, error)
	// Updated Linked File allows you to update properties of a previously attached Linked-File.
	UpdateLinkedFile(linkedFilePublicId int64, updateLinkedFile *models.UpdateLinkedFile) (*models.LinkedFile, error)
	// Returns information about the authenticated member.
	GetCurrentMemberInfo() (*models.MemberInfo, error)
	// List Members returns information about members of the organization.
	ListMembers(listMembers *models.ListMembers) (*[]models.Member, error)
	// Returns information about a Member.
	GetMember(memberPublicId string, getMember *models.GetMember) (*models.Member, error)
	// List Milestones returns a list of all Milestones and their attributes.
	ListMilestones() (*[]models.Milestone, error)
	// Create Milestone allows you to create a new Milestone in Clubhouse.
	CreateMilestone(createMilestone *models.CreateMilestone) error
	// Delete Milestone can be used to delete any Milestone.
	DeleteMilestone(milestonePublicId int64) error
	// Get Milestone returns information about a chosen Milestone.
	GetMilestone(milestonePublicId int64) (*models.Milestone, error)
	// Update Milestone can be used to update Milestone properties.
	UpdateMilestone(milestonePublicId int64, updateMilestone *models.UpdateMilestone) (*models.Milestone, error)
	// List all of the Epics within the Milestone.
	ListMilestoneEpics(milestonePublicId int64) (*[]models.EpicSlim, error)
	// List Projects returns a list of all Projects and their attributes.
	ListProjects() (*[]models.Project, error)
	// Create Project is used to create a new Clubhouse Project.
	CreateProject(createProject *models.CreateProject) error
	// Delete Project can be used to delete a Project. Projects can only be deleted if all associated Stories are moved or deleted. In the case that the Project cannot be deleted, you will receive a 422 response.
	DeleteProject(projectPublicId int64) error
	// Get Project returns information about the selected Project.
	GetProject(projectPublicId int64) (*models.Project, error)
	// Update Project can be used to change properties of a Project.
	UpdateProject(projectPublicId int64, updateProject *models.UpdateProject) (*models.Project, error)
	// List Stories returns a list of all Stories in a selected Project and their attributes.
	ListStories(projectPublicId int64, getProjectStories *models.GetProjectStories) (*[]models.StorySlim, error)
	// List Repositories returns a list of all Repositories and their attributes.
	ListRepositories() (*[]models.Repository, error)
	// Get Repository returns information about the selected Repository.
	GetRepository(repoPublicId int64) (*models.Repository, error)
	// Search lets you search Epics and Stories based on desired parameters. Since ordering of the results can change over time (due to search ranking decay, new Epics and Stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	Search(search *models.Search) (*models.SearchResults, error)
	// Search Epics lets you search Epics based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new Epics being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	SearchEpics(search *models.Search) (*models.EpicSearchResults, error)
	// Search Stories lets you search Stories based on desired parameters. Since ordering of stories can change over time (due to search ranking decay, new stories being created), the `next` value from the previous response can be used as the path and query string for the next page to ensure stable ordering.
	SearchStories(search *models.Search) (*models.StorySearchResults, error)
	// Create Story is used to add a new story to your Clubhouse.
	CreateStory(createStoryParams *models.CreateStoryParams) error
	// Delete Multiple Stories allows you to delete multiple archived stories at once.
	DeleteMultipleStories(deleteStories *models.DeleteStories) error
	// Create Multiple Stories allows you to create multiple stories in a single request using the same syntax as [Create Story](https://clubhouse.io/api/#create-story).
	CreateMultipleStories(createStories *models.CreateStories) error
	// Update Multiple Stories allows you to make changes to numerous stories at once.
	UpdateMultipleStories(updateStories *models.UpdateStories) (*[]models.StorySlim, error)
	// Search Stories lets you search Stories based on desired parameters.
	SearchStoriesOld(searchStories *models.SearchStories) error
	// Delete Story can be used to delete any Story.
	DeleteStory(storyPublicId int64) error
	// Get Story returns information about a chosen Story.
	GetStory(storyPublicId int64) (*models.Story, error)
	// Update Story can be used to update Story properties.
	UpdateStory(storyPublicId int64, updateStory *models.UpdateStory) (*models.Story, error)
	// Create Comment allows you to create a Comment on any Story.
	CreateComment(storyPublicId int64, createComment *models.CreateComment) error
	// Delete a Comment from any story.
	DeleteComment(storyPublicId int64, commentPublicId int64) error
	// Get Comment is used to get Comment information.
	GetComment(storyPublicId int64, commentPublicId int64) (*models.Comment, error)
	// Update Comment replaces the text of the existing Comment.
	UpdateComment(storyPublicId int64, commentPublicId int64, updateComment *models.UpdateComment) (*models.Comment, error)
	// Delete a Reaction from any comment.
	DeleteReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error
	// Create a reaction to a comment.
	CreateReaction(storyPublicId int64, commentPublicId int64, createOrDeleteReaction *models.CreateOrDeleteReaction) error
	// Create Task is used to create a new task in a Story.
	CreateTask(storyPublicId int64, createTask *models.CreateTask) error
	// Delete Task can be used to delete any previously created Task on a Story.
	DeleteTask(storyPublicId int64, taskPublicId int64) error
	// Returns information about a chosen Task.
	GetTask(storyPublicId int64, taskPublicId int64) (*models.Task, error)
	// Update Task can be used to update Task properties.
	UpdateTask(storyPublicId int64, taskPublicId int64, updateTask *models.UpdateTask) (*models.Task, error)
	// Story Links (called Story Relationships in the UI) allow you create semantic relationships between two stories. The parameters read like an active voice grammatical sentence:  subject -> verb -> object.
	//
	// The subject story acts on the object Story; the object story is the direct object of the sentence.
	//
	// The subject story "blocks", "duplicates", or "relates to" the object story.  Examples:
	// - "story 5 blocks story 6” -- story 6 is now "blocked" until story 5 is moved to a Done workflow state.
	// - "story 2 duplicates story 1” -- Story 2 represents the same body of work as Story 1 (and should probably be archived).
	// - "story 7 relates to story 3”
	CreateStoryLink(createStoryLink *models.CreateStoryLink) error
	// Removes the relationship between the stories for the given Story Link.
	DeleteStoryLink(storyLinkPublicId int64) error
	// Returns the stories and their relationship for the given Story Link.
	GetStoryLink(storyLinkPublicId int64) (*models.StoryLink, error)
	// Updates the stories and/or the relationship for the given Story Link.
	UpdateStoryLink(storyLinkPublicId int64, updateStoryLink *models.UpdateStoryLink) (*models.StoryLink, error)
	// List Teams returns a list of all Teams in the organization.
	ListTeams() (*[]models.Team, error)
	// Get Team is used to get Team information.
	GetTeam(teamPublicId int64) (*models.Team, error)
	// List Workflows returns a list of all Workflows in the organization.
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

func (a *api) request(method string, endpoint string, params url.Values, data io.Reader, out interface{}) error {

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
