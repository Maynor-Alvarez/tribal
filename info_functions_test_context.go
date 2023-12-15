package main

import (
	"errors"

	"bitbucket.org/bakingthecookie/bi-bel3-canella-rest-go/creditcards"
	security "bitbucket.org/bakingthecookie/bi-bel3-security-go"
)

type GetListMock struct {
	Status          int
	Message         string
	ListCreditCards []creditcards.CreditCard
	Err             error
}
type TestContextGetList struct {
	TraceId     string
	User        security.User
	GetListMock GetListMock
}

var contextGetList = TestContextGetList{
	TraceId: "",
	User: security.User{
		CodigoInstalacion: "833",
		Usuario:           "cobos",
		TieneToken:        true,
	},
	GetListMock: GetListMock{
		Status:          0,
		Message:         "Registro consultado con éxito",
		ListCreditCards: []creditcards.CreditCard{},
		Err:             nil,
	},
}

var MapTestContextGetList = map[string]TestContextGetList{
	"case1": {
		TraceId: contextGetList.TraceId,
		User:    contextGetList.User,
		GetListMock: GetListMock{
			Status:          0,
			Message:         "Ocurrio un error inesperado",
			ListCreditCards: []creditcards.CreditCard{},
			Err:             errors.New("Ocurrurrio un error inesperado"),
		},
	},
	"case2": {
		TraceId: contextGetList.TraceId,
		User:    contextGetList.User,
		GetListMock: GetListMock{
			Status:          1,
			Message:         "NO tiene tarjetas",
			ListCreditCards: []creditcards.CreditCard{},
			Err:             nil,
		},
	},
	"case3": {
		TraceId: contextGetList.TraceId,
		User:    contextGetList.User,
		GetListMock: GetListMock{
			Status:          0,
			Message:         "Se consulto exitosamente las tarjetas",
			ListCreditCards: buildResponseCreditCards(),
			Err:             nil,
		},
	},
}

func buildResponseCreditCards() []creditcards.CreditCard {
	return []creditcards.CreditCard{
		{
			Number: "5524570000501112",
			Name:   "DIEGO PULIDO ARAGON",
			Status: 1,
		},
		{
			Number: "4236161924202110",
			Name:   "DIEGO PULIDO ARAGON",
			Status: 1,
		},
	}
}
