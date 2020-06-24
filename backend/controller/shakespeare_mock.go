package controller

import "github.com/tmp202006/shakespoke/backend/model"

// ShakespeareControllerMock is the mock for ShakespeareControl
type ShakespeareControllerMock struct {
	Translation *model.ShakespeareTranslation
	Error       error
}

// Translate mocks the translate method by returning dummy data specified in the mock struct
func (mock ShakespeareControllerMock) Translate(s string) (*model.ShakespeareTranslation, error) {
	return mock.Translation, mock.Error
}
