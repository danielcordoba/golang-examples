package api

type Person struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	Gender    string `json:"gender"`
	BirthYear string `json:"birth_year"`
}

type PeopleResponse struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Person `json:"results"`
}

// StarWarsRepo definition of methods to access a data people
type StarWarsRepo interface {
	GetPeople() (people []Person, err error)
	GetPerson(id string) (people []Person, err error)
}
