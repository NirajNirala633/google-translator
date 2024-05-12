package model

type Translate struct {
	SourceLanguage string `json:"source"`
	TargetLanguage string `json:"target"`
	Text           string `json:"text"`
}
