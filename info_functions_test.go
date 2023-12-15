package main

import (
	"reflect"
	"testing"

	mocksCanellaRest "bitbucket.org/bakingthecookie/bi-bel3-canella-rest-go/creditcards/mocks"
	"bitbucket.org/bakingthecookie/bi-bel3-credit-cards-go/db"
)

func Test_getBenefitsTermsAndConditionsCreditCards(t *testing.T) {
	loadPersonalizationDBMock()
	tests := []struct {
		name         string
		args         TestContextGetList
		wantResponse GetBenefitsTermsAndConditionsResponse
		wantErr      bool
	}{
		{
			name: "CASE 1: When service GetList return error",
			args: MapTestContextGetList["case1"],
			wantResponse: GetBenefitsTermsAndConditionsResponse{
				CreditCards: []CreditCardInformation{},
			},
			wantErr: true,
		},
		{
			name: "CASE 2 : When service return status different to 0",
			args: MapTestContextGetList["case2"],
			wantResponse: GetBenefitsTermsAndConditionsResponse{
				CreditCards: []CreditCardInformation{},
			},
			wantErr: false,
		},
		{
			name: "CASE 3 : When service return OK",
			args: MapTestContextGetList["case3"],
			wantResponse: GetBenefitsTermsAndConditionsResponse{
				CreditCards: []CreditCardInformation{
					{
						Name:            "DIEGO PULIDO ARAGON",
						Number:          "5524-XXXX-XXXX-1112",
						LinkInformation: "https://tarjetasbi.com/visa-signature/",
					},
					{
						Name:            "DIEGO PULIDO ARAGON",
						Number:          "4236-XXXX-XXXX-2110",
						LinkInformation: "https://tarjetasbi.com/visa-signature/",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Register mock
			mocks := mocksCanellaRest.Consumer{}
			mocks.On("GetList", "", test.args.User).Return(test.args.GetListMock.Status, test.args.GetListMock.Message, test.args.GetListMock.ListCreditCards, test.args.GetListMock.Err)
			// creditCards.RegisterConsumer(&mocks)

			gotResponse, err := getBenefitsTermsAndConditionsCreditCards(test.args.TraceId, test.args.User)
			if (err != nil) != test.wantErr {
				t.Errorf("getBenefitsTermsAndConditionsCreditCards() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, test.wantResponse) {
				t.Errorf("getBenefitsTermsAndConditionsCreditCards() = %v, want %v", gotResponse, test.wantResponse)
			}
		})
	}
}

func loadPersonalizationDBMock() {
	creditCardPersonalization = map[string]db.PersonalizationCreditCard{
		"423616": {
			DBID:                    31,
			UrlBackgroundHorizontal: "https://www.clubdecomprasbi.com/Imagenes/PAT/visaorolocal.png",
			CardType:                "VISA",
			FeesText:                "Permitir Visa Cuotas",
			UrlBackgroundRear:       "https://www.clubdecomprasbi.com/Imagenes/PAT/visalocalorobackside.png",
			LinkInformation:         "https://tarjetasbi.com/visa-signature/",
		},
		"552457": {
			DBID:                    57,
			UrlBackgroundHorizontal: "https://www.clubdecomprasbi.com/Imagenes/PAT/mcblackn.png",
			CardType:                "MASTERCARD",
			FeesText:                "Permitir Master Cuotas",
			UrlBackgroundRear:       "https://www.clubdecomprasbi.com/Imagenes/PAT/mastercardblackbackside.png",
			LinkInformation:         "https://tarjetasbi.com/visa-signature/",
		},
	}
}
