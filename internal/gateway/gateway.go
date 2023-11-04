package gateway

import (
	"movie-service/internal/domain"
	imdbgw "movie-service/internal/gateway/imdb"
)

type gateway struct {
}

func InitGateway() domain.IGateway {
	return &gateway{}
}

func (g *gateway) ImdbGateway() domain.IImdbGateway {
	return &imdbgw.Imdb{}
}
