package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/tmp202006/shakespoke/backend/model"
)

// ShakespeareControl is the interface defining the methods for interacting
// with Shakespeare external API
type ShakespeareControl interface {
	Translate(s string) (*model.ShakespeareTranslation, error)
}

// ShakespeareController implements ShakespeareControl by actually
// connecting to the external API
type ShakespeareController struct {
	apiURL string
}

// NewShakespeareController init a new ShakespeareController instance
func NewShakespeareController(apiURL string) ShakespeareController {
	return ShakespeareController{
		apiURL: apiURL,
	}
}

// Translate implements the request to the external API
func (sc ShakespeareController) Translate(s string) (*model.ShakespeareTranslation, error) {
	values := url.Values{}
	values.Add("text", s)
	reqURL := fmt.Sprintf("%v/translate/shakespeare.json?%v", sc.apiURL, values.Encode())
	log.Println(reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	var translation model.ShakespeareTranslation
	if err := json.Unmarshal(respBody, &translation); err != nil {
		return nil, err
	}
	return &translation, nil
}
