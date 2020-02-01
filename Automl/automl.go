package Automl

type AutomlRequest struct {
	ProjectID    string
	SessionID    string
	Text         string
	LanguageCode string
}

type AutomlResponse struct {
}
