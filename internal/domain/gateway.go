package domain

type IGateway interface {
	ImdbGateway() IImdbGateway
}

type IImdbGateway interface {
}
