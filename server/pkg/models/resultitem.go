package models

type ResultItem struct {
	ResultID   string   `json:"result_id"`
	ResultName string   `json:"result_name"`
	Text       []string `json:"text"`
	Keywords   []string `json:"keywords"`
	Link       string   `json:"link"`
	Timestamp  string   `json:"timestamp"`
}
