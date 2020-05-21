package v2

import (
	"encoding/json"
	"fmt"

	"github.com/acaciamoney/basiq-sdk-golang/errors"
)

type Institution struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	ShortName       string `json:"shortName"`
	Country         string `json:"country"`
	ServiceName     string `json:"serviceName"`
	ServiceType     string `json:"serviceType"`
	LoginIdCaption  string `json:"loginIdCaption"`
	PasswordCaption string `json:"PasswordCaption"`
	Logo            struct {
		Type   string `json:"type"`
		Colors string `json:"colors"`
		Links  struct {
			Square string `json:"square"`
			Full   string `json:"full"`
		} `json:"links"`
	} `json:"logo"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type InstitutionsList struct {
	Count int           `json:"count"`
	Data  []Institution `json:"data"`
}

type InstitutionService struct {
	Session Session
}

func NewInstitutionService(session *Session) *InstitutionService {
	return &InstitutionService{
		Session: *session,
	}
}

func (is *InstitutionService) GetInstitutions() (InstitutionsList, *errors.APIError) {
	var data InstitutionsList

	body, _, err := is.Session.Api.Send("GET", "institutions", nil)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (is *InstitutionService) GetInstitution(institutionId string) (Institution, *errors.APIError) {
	var data Institution

	body, _, err := is.Session.Api.Send("GET", "institutions/"+institutionId, nil)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}
