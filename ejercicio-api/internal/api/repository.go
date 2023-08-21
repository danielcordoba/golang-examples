package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	peopleEndpoint = "/people"
	starWarsUrl    = "https://swapi.dev/api"
)

type repository struct {
	url string
}

func NewStarWarsRepo() StarWarsRepo {
	return &repository{starWarsUrl}
}

func (b *repository) GetPeople() (people []Person, err error) {

	var peopleResponse PeopleResponse

	endpoint := fmt.Sprintf("%v%v", b.url, peopleEndpoint)

	response, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &peopleResponse)

	people = peopleResponse.Results

	if err != nil {
		return nil, err
	}

	return
}

func (b *repository) GetPerson(id string) (people []Person, err error) {

	endpoint := fmt.Sprintf("%v%v/%v", b.url, peopleEndpoint, id)
	response, err := http.Get(endpoint)
	var person Person

	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &person)

	if err != nil {
		return nil, err
	}

	return []Person{person}, nil

}
