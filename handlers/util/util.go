package util

type Event struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"created_at"`
}
