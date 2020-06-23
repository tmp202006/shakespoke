package model

// ShakespeareTranslation is the response for a Shakespeare Translation request
type ShakespeareTranslation struct {
	Success  ShakespeareTranslationSuccess  `json:"success"`
	Contents ShakespeareTranslationContents `json:"contents"`
}

// ShakespeareTranslationSuccess describes the successful number of requests
type ShakespeareTranslationSuccess struct {
	Total int64 `json:"total"`
}

// ShakespeareTranslationContents includes the translation details
type ShakespeareTranslationContents struct {
	Translation string `json:"translation"`
	Text        string `json:"text"`
	Translated  string `json:"translated"`
}
