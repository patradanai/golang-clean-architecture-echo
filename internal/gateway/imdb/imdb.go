package imdbgw

import (
	"encoding/json"
	"movie-service/pkg/api"
	errs "movie-service/pkg/errors"
	"net/http"
)

type Imdb struct {
	client *api.RestClient
}

func InitImdb() *Imdb {
	client := api.InitRestClient(api.RestOptions{
		Host:    "http://www.omdbapi.com",
		Headers: map[string]string{},
	})
	return &Imdb{
		client: client,
	}
}

func (i *Imdb) GetMovieByImdbID() (*MovieSearchGWReponse, errs.Errors) {
	res, err := i.client.GET("search", nil)
	if err != nil {
		return nil, errs.WrapDError(err, errs.InternalServerError)
	}

	if res.StatusCode() != http.StatusOK {
		return nil, api.WrapRestError(res)
	}

	movie := &MovieSearchGWReponse{}
	err = json.Unmarshal(res.Body(), movie)
	if err != nil {
		return nil, errs.WrapDError(err, errs.InternalServerError)
	}

	return movie, nil
}
