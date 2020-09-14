package ontario

import (
	"encoding/json"
	"fmt"
	beer "github.com/CodelyTV/golang-examples/02-refactor-to-cobra/internal"
	"io/ioutil"
	"net/http"
)

const ontarioUrl = "http://ontariobeerapi.ca/"
const endpoint = "beers"

type ontarioRepository struct {
	url string
}

func NewOntarioRepository() beer.BeersRepository {
	return &ontarioRepository{url: ontarioUrl}
}

func (r ontarioRepository) GetBeers() (beers []beer.Beer, er error) {
	response, err := http.Get(fmt.Sprintf("%v%v", ontarioUrl, endpoint))

	if nil != err {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)

	if nil != err {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)

	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}

	return
}
