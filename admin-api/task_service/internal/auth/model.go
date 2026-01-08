package auth

type Assignee struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Attachment struct {
	Type string `json:"type"`
	Url string `json:"url"`
}

type HistoryItem struct {
	Timestamp string `json:"timestamp"`
	Action string `json:"action"`
	User string `json:"user"`
}

type Task struct {
	ID string `json:"id"`
	AccountId string `json:"account_id"`
	Category string `json:"category"`
	Priority string `json:"priority"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
	ExecutorID  string       `json:"executor_id"`
	Assignee    Assignee      `json:"assignee"`
	Attachments []Attachment  `json:"attachments"`
	History     []HistoryItem `json:"history"`       
}