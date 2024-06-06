package models

type Movie struct {
	Name string `json:"name"`
	ID   string `json:"ID"`
}
type WebParam struct {
	ID    string `json:"ID"`
	param string
}

// ScrapeRequest represents the user input for scraping
type ScrapeRequest struct {
	URL        string            `json:"url" form:"url"`
	Parameters map[string]string `json:"parameters" form:"parameters"`
}

// ScrapeResult represents the result of the scraping process
type ScrapeResult struct {
	Data map[string]string `json:"data"`
}
