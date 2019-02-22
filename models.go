package graphql_poc

type User struct {
	ID       string     `json:"id,omitempty"`
	Username string     `json:"username,omitempty"`
	Messages []*Message `json:"messages,omitempty"`
}

type Message struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
	User string `json:"user,omitempty"`
}
