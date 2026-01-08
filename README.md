# go-github-activity-tracker
GitHub Activity Tracker CLI

GitHub Activity Tracker CLI
A sleek, lightweight command-line tool built with Go that fetches and displays GitHub user activity in real-time. Perfect for monitoring your contributions, tracking team member activities, or exploring what top developers are working on.

Features
Core Functionality

🔍 Real-time Activity Fetching - Retrieves up to 30 most recent GitHub events using the official GitHub REST API
📊 Comprehensive Event Support - Tracks 15+ different GitHub activities including:

Push events with commit counts and branch information
Repository creation, forking, and starring
Issue creation, updates, and comments
Pull request activities and reviews



User Experience

🎨 Clean Output - Well-formatted, easy-to-scan terminal display
⚡ Blazing Fast - Written in Go for optimal performance
🛡️ Robust Error Handling - Graceful handling of network issues, rate limits, and invalid users

Technical Highlights

🔒 No External Dependencies - Pure Go standard library implementation
🌐 REST API Integration - Direct interaction with GitHub's public API
📝 Type-Safe JSON Parsing - Structured data handling with Go structs
⏱️ Built-in Timeout Protection - 10-second request timeout prevents hanging
🔄 Proper Resource Management - Deferred cleanup and connection handling

🎯 Use Cases

Developers: Monitor your daily contributions and maintain your coding streak
Team Leads: Track team member activities and contributions
Open Source Enthusiasts: Follow your favorite developers and discover new projects
Recruiters: Verify candidate activity and engagement on GitHub
Students: Learn from active developers by observing their workflow

📦 Installation
Prerequisites

Go 1.21 or higher
Internet connection for API access
GitHub account (for checking your own activity)

Quick Start
bash# Clone the repository
git clone https://github.com/AboloreDev/github-activity-tracker.git
cd github-activity-tracker

# Initialize Go module
go mod init github-activity

# Build the binary
go build -o github-activity

# Run it!
./go-github-activity-tracker <username>
Alternative: Run Without Building
bashgo run main.go <username>
🎮 Usage
Basic Usage
bash# Check any GitHub user's activity
./github-activity torvalds

# Check your own activity
./go-github-activity-tracker <username>
Debug Mode
bash# Enable debug mode for detailed API response inspection
./go-github-activity-tracker <username> --debug
Sample Output
Fetching activity for GitHub user: aboloredev

Recent Activity (showing 10 events):
----------------------------------------------------------------------
- Pushed 3 commits to 'main' branch in aboloredev/github-activity-tracker (2 hours ago)
- Created repository aboloredev/go-task-cli (5 hours ago)
- Starred torvalds/linux (1 day ago)
- Opened issue #42 in aboloredev/project-name (2 days ago)
- Forked octocat/Hello-World to aboloredev/Hello-World (3 days ago)
- Commented on pull request #15 in golang/go (4 days ago)
🏗️ Architecture
Project Structure
github-activity-tracker/
├── main.go           # Main application code
├── go.mod           # Go module definition
└── README.md        # Documentation
Core Components

Event Structures: Type-safe representations of GitHub API responses
HTTP Client: Configured with timeouts and proper headers
Event Parser: JSON unmarshaling with error handling
Formatter: Converts raw events into human-readable messages
Time Calculator: Intelligent relative time display

API Integration
The tool uses GitHub's REST API v3:
Endpoint: https://api.github.com/users/{username}/events
Method: GET
Rate Limit: 60 requests/hour (unauthenticated)
🔧 Technical Details
Supported Event Types
Event TypeDescriptionPushEventCode pushes with commit informationCreateEventRepository, branch, or tag creationDeleteEventBranch or tag deletionForkEventRepository forkingWatchEventRepository starringIssuesEventIssue creation and updatesIssueCommentEventComments on issuesPullRequestEventPull request actionsPullRequestReviewEventPR reviewsPullRequestReviewCommentEventComments on PR reviewsReleaseEventRelease publicationsPublicEventRepository made publicMemberEventCollaborator additionsCommitCommentEventCommit commentsGollumEventWiki updates
Error Handling

404 Not Found: User doesn't exist
403 Forbidden: API rate limit exceeded
Network Errors: Connection timeout or failure
JSON Parsing: Malformed API response
Empty Response: User has no public activity

🔮 Future Enhancements

 Authentication support for higher rate limits
 Activity filtering by event type
 Export to JSON/CSV formats
 Colorized output for better readability
 Pagination support for older events
 Cache layer for repeated queries
 Webhook integration for real-time notifications
 Multiple user comparison
 Statistical analysis (commits per day, most active repos)

🤝 Contributing
Contributions are welcome! Here's how you can help:

Fork the repository
Create a feature branch (git checkout -b feature/AmazingFeature)
Commit your changes (git commit -m 'Add some AmazingFeature')
Push to the branch (git push origin feature/AmazingFeature)
Open a Pull Request

📝 License
This project is licensed under the MIT License - see the LICENSE file for details.
🙏 Acknowledgments

GitHub REST API v3 Documentation
Go Programming Language
The amazing open-source community

📧 Contact
Abolore - @AboloreDev
Project Link: https://github.com/AboloreDev/go-github-activity-tracker

🎯 Project Origin
This project was inspired by the roadmap.sh Backend Projects - a fantastic resource for hands-on learning and skill development.
Challenge: Build a CLI tool to fetch and display GitHub user activity
Goal: Practice API integration, JSON parsing, and CLI development
Result: A production-ready tool with comprehensive error handling and user-friendly output

⭐ If you found this project helpful, please consider giving it a star on GitHub!
