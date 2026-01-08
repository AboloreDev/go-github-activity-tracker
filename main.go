package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Structs
type Event struct {
	Type      string    `json:"type"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action      string       `json:"action"`
	Ref         string       `json:"ref"`
	RefType     string       `json:"ref_type"`
	Commits     []Commit     `json:"commit"`
	Forkee      *Forkee      `json:"forkee"`
	Issue       *Issue       `json:"issue"`
	PullRequest *PullRequest `json:"pull_request"`
	Comment     *Comment     `json:"comment"`
	Size        int          `json:"size"`
}

type Commit struct {
	Message string `json:"message"`
	Sha     string `json:"sha"`
}

type Forkee struct {
	FullName string `json:"full-name"`
}

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type Comment struct {
	Body string `json:"body"`
}

// Main Function
func main() {
	// validation
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-github-activity-tracker <username>")
		fmt.Println("\nExample:")
		fmt.Println("  github-activity johndoe")
	}

	// Set the username
	username := os.Args[1]

	// Validate against empty string
	if strings.TrimSpace(username) == "" {
		fmt.Println("Username cannot be empty")
		os.Exit(1)
	}

	fmt.Printf("Fetching activity for GitHub user: %s\n\n", username)

	// check for Activity
	events, err := FetchGithubActivity(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found for this user.")
		return
	}

	// Display events
	DisplayActivity(events)
}

// Fetch Github Activity
func FetchGithubActivity(username string) ([]Event, error) {
	// Url to call
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	// Create a client timeout that last for 10 seconds and return to avoid being stuck
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send a request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the header
	request.Header.Set("User-Agent", "github-activity-cli")

	// Make the request
	// This is like a dispatch that performs the action
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}

	// Close to avoid memory leaks
	// Close when done
	defer response.Body.Close()

	// Status Codes
	// Check response status
	if response.StatusCode == 404 {
		return nil, fmt.Errorf("user '%s' not found", username)
	}
	if response.StatusCode == 403 {
		return nil, fmt.Errorf("API rate limit exceeded. Please try again later")
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API returned status code: %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Unmarshal the json into readable struct
	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// return
	return events, nil
}

// Display actiivty
// Range over each event and siaplay them as message
func DisplayActivity(events []Event) {
	fmt.Println("Recent Activity:")
	fmt.Println(strings.Repeat("-", 70))

	for _, event := range events {
		message := FormatEvent(event)
		if message != "" {
			fmt.Println(message)
		}
	}
}

// Format the event
// Conversion of the json into human readable code
func FormatEvent(event Event) string {
	repo := event.Repo.Name

	switch event.Type {
	case "PushEvent":
		commitCount := len(event.Payload.Commits)
		if commitCount == 1 {
			return fmt.Sprintf("- Pushed 1 commit to %s", repo)
		}
		return fmt.Sprintf("- Pushed %d commits to %s", commitCount, repo)
	case "CreateEvent":
		refType := event.Payload.RefType
		switch refType {
	case "repository":
			return fmt.Sprintf("- Created repository %s", repo)
		case "branch":
			return fmt.Sprintf("- Created branch '%s' in %s", event.Payload.Ref, repo)
		case "tag":
			return fmt.Sprintf("- Created tag '%s' in %s", event.Payload.Ref, repo)
		}
		return fmt.Sprintf("- Created %s in %s", refType, repo)

	case "DeleteEvent":
		refType := event.Payload.RefType
		return fmt.Sprintf("- Deleted %s '%s' in %s", refType, event.Payload.Ref, repo)

	case "ForkEvent":
		if event.Payload.Forkee != nil {
			return fmt.Sprintf("- Forked %s to %s", repo, event.Payload.Forkee.FullName)
		}
		return fmt.Sprintf("- Forked %s", repo)

	case "WatchEvent":
		return fmt.Sprintf("- Starred %s", repo)

	case "IssuesEvent":
		if event.Payload.Issue != nil {
			action := event.Payload.Action
			issueNum := event.Payload.Issue.Number
			return fmt.Sprintf("- %s issue #%d in %s", capitalizeFirst(action), issueNum, repo)
		}
		return fmt.Sprintf("- %s an issue in %s", capitalizeFirst(event.Payload.Action), repo)

	case "IssueCommentEvent":
		if event.Payload.Issue != nil {
			issueNum := event.Payload.Issue.Number
			return fmt.Sprintf("- Commented on issue #%d in %s", issueNum, repo)
		}
		return fmt.Sprintf("- Commented on an issue in %s", repo)

	case "PullRequestEvent":
		if event.Payload.PullRequest != nil {
			action := event.Payload.Action
			prNum := event.Payload.PullRequest.Number
			return fmt.Sprintf("- %s pull request #%d in %s", capitalizeFirst(action), prNum, repo)
		}
		return fmt.Sprintf("- %s a pull request in %s", capitalizeFirst(event.Payload.Action), repo)

	case "PullRequestReviewEvent":
		if event.Payload.PullRequest != nil {
			prNum := event.Payload.PullRequest.Number
			return fmt.Sprintf("- Reviewed pull request #%d in %s", prNum, repo)
		}
		return fmt.Sprintf("- Reviewed a pull request in %s", repo)

	case "PullRequestReviewCommentEvent":
		if event.Payload.PullRequest != nil {
			prNum := event.Payload.PullRequest.Number
			return fmt.Sprintf("- Commented on pull request #%d in %s", prNum, repo)
		}
		return fmt.Sprintf("- Commented on a pull request in %s", repo)

	case "ReleaseEvent":
		return fmt.Sprintf("- Published a release in %s", repo)

	case "PublicEvent":
		return fmt.Sprintf("- Made %s public", repo)

	case "MemberEvent":
		return fmt.Sprintf("- Added a collaborator to %s", repo)

	case "CommitCommentEvent":
		return fmt.Sprintf("- Commented on a commit in %s", repo)

	case "GollumEvent":
		return fmt.Sprintf("- Updated the wiki in %s", repo)

	default:
		// Return empty string for unknown event types
		return ""
	}

}

// capitalizeFirst capitalizes the first letter of a string
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
