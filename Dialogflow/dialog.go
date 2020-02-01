package Dialogflow

type DialogflowRequest struct {
	ProjectID    string
	SessionID    string
	Text         string
	LanguageCode string
}

type DialogflowResponse struct {
}
