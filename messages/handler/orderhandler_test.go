package handler

import (
	"github.com/jacob-elektronik/gofact/messages/model"
	"github.com/jacob-elektronik/gofact/messages/model/segments"
	"github.com/jacob-elektronik/gofact/parser"
	"github.com/jacob-elektronik/gofact/segment"
	"reflect"
	"testing"
)

func TestUnmarshalOrder(t *testing.T) {
	type args struct {
		messageSegments []segment.Segment
	}
	p := parser.NewParser("../../example/edi_messages/testorder.edi", "")
	err := p.ParseEdiFactMessageConcurrent()
	if err != nil {
		t.Errorf("Error parsing file")
		return
	}
	orderWant := model.OrderMessage{
		InterchangeHeader: segments.InterchangeHeader{
			SyntaxIdentifier: segments.SyntaxIdentifier{
				SyntaxIdentifier:    "UNOC",
				SyntaxVersionNumber: "3",
			},
			InterchangeSender: segments.InterchangeSender{
				SenderIdentification:               "1234567891234",
				PartnerIdentificationCodeQualifier: "14",
				AddressReverseRouting:              "",
			},
			InterchangeRecipient: segments.InterchangeRecipient{
				RecipientIdentification:            "9876543219876",
				PartnerIdentificationCodeQualifier: "92",
				RoutingAddress:                     "",
			},
			DateTime: segments.DateTime{
				DateOfPreparation: "180907",
				TimeOfPreparation: "1418",
			},
			InterchangeControlReference: "ORD12345678",
			RecipientsReferencePassword: segments.RecipientsReferencePassword{
				RecipientsReferencePassword:          "",
				RecipientsReferencePasswordQualifier: "",
			},
			ApplicationReference:      "",
			ProcessingPriorityCode:    "",
			AcknowledgementRequest:    "",
			CommunicationsAgreementID: "",
			TestIndicator:             "",
		},
		MessageHeader: segments.MessageHeader{
			MessageReferenceNumber: "1",
			MessageIdentifier: segments.MessageIdentifier{
				MessageType:                          "ORDERS",
				MessageVersionNumber:                 "D",
				MessageReleaseNumber:                 "96A",
				ControllingAgency:                    "UN",
				AssociationAssignedCode:              "EAN008",
				CodeLisDirectoryVersionNumber:        "",
				MessageTypeSubFunctioIdentificationn: "",
			},
			CommonAccessReference: "",
			StatusOfTransfer: segments.StatusOfTransfer{
				SequenceOfTransfers:  "",
				FirstAndLastTransfer: "",
			},
			MessageSubsetIdentification: segments.MessageSubsetIdentification{
				MessageSubsetIdentification: "",
				MessageSubsetVersionNumber:  "",
				MessageSubsetReleaseNumber:  "",
				ControllingAgency:           "",
			},
			MessageImplementationGuidelineIdentification: segments.MessageImplementationGuidelineIdentification{
				MessageImplementationGuidelineIdentification: "",
				MessageImplementationGuidelineVersionNumber:  "",
				MessageImplementationGuidelineReleaseNumber:  "",
				ControllingAgency: "",
			},
			ScenarioIdentification: segments.ScenarioIdentification{
				ScenarioIdentification: "",
				ScenarioVersionNumber:  "",
				ScenarioReleaseNumber:  "",
				ControllingAgency:      "",
			},
		},
		BeginningOfMessage: segments.BeginningOfMessage{
			MessageName: segments.MessageName{
				DocumentNameCode:              "226",
				CodeListIdentificationCode:    "",
				CodeListResponsibleAgencyCode: "",
				DocumentName:                  "",
			},
			MessageIdentification: segments.MessageIdentification{
				DocumentIdentifier: "12345678",
				VersionIdentifier:  "",
				RevisionIdentifier: "",
			},
			MessageFunctionCode: "",
			ResponseTypeCode:    "",
		},
		DateTimePeriod: segments.DateTimePeriod{
			DTMFunctionCode: "137",
			DTMValue:        "20180823",
			DTMFormatCode:   "102",
		},
		ReferenceNumbersOrders: []model.ReferenceNumber{{
			Reference: segments.Reference{
				ReferenceCodeQualifier: "CR",
				ReferenceIdentifier:    "CustomerRefNumber",
				DocumentLineIdentifier: "",
				VersionIdentifier:      "",
				RevisionIdentifier:     "",
			},
			DateTimePeriod: segments.DateTimePeriod{
				DTMFunctionCode: "",
				DTMValue:        "",
				DTMFormatCode:   "",
			},
		}, {
			Reference: segments.Reference{
				ReferenceCodeQualifier: "IT",
				ReferenceIdentifier:    "InternalOrderNumber",
				DocumentLineIdentifier: "",
				VersionIdentifier:      "",
				RevisionIdentifier:     "",
			},
			DateTimePeriod: segments.DateTimePeriod{
				DTMFunctionCode: "",
				DTMValue:        "",
				DTMFormatCode:   "",
			},
		}, {
			Reference: segments.Reference{
				ReferenceCodeQualifier: "CT",
				ReferenceIdentifier:    "ContractNumber",
				DocumentLineIdentifier: "",
				VersionIdentifier:      "",
				RevisionIdentifier:     "",
			},
			DateTimePeriod: segments.DateTimePeriod{
				DTMFunctionCode: "",
				DTMValue:        "",
				DTMFormatCode:   "",
			},
		}},
		Parties: []model.Party{{
			NameAddress: segments.NameAddress{
				PartyFunctionCodeQualifier: "SU",
				PartyIdenNameAndAddressDescriptiontificationDetails: segments.PartyIdentificationDetails{
					PartyIdentifier:               "9876543219876",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "92",
				},
				NameAndAddressDescription: segments.NameAndAddressDescription{
					NameAndAddressDescription:                 "",
					NameAndAddressDescriptionConditionalOne:   "",
					NameAndAddressDescriptionConditionalTwo:   "",
					NameAndAddressDescriptionConditionalThree: "",
					NameAndAddressDescriptionConditionalFour:  "",
				},
				PartyName: segments.PartyName{
					PartyName:                 "party name",
					PartyNameConditionalOne:   "",
					PartyNameConditionalTwo:   "",
					PartyNameConditionalThree: "",
					PartyNameConditionalFour:  "",
					PartyNameConditionalFive:  "",
					PartyNameCode:             "",
				},
				Street: segments.Street{
					Street:                 "teststreet",
					StreetConditionalOne:   "",
					StreetConditionalTwo:   "",
					StreetConditionalThree: "",
				},
				CityName:    "testcity",
				Postal:      "12345",
				CountryCode: "DE",
			},
			ReferenceNumbersParties: model.ReferenceNumber{
				Reference: segments.Reference{
					ReferenceCodeQualifier: "VA",
					ReferenceIdentifier:    "9876543219876",
					DocumentLineIdentifier: "",
					VersionIdentifier:      "",
					RevisionIdentifier:     "",
				},
				DateTimePeriod: segments.DateTimePeriod{
					DTMFunctionCode: "",
					DTMValue:        "",
					DTMFormatCode:   "",
				},
			},
			ContactDetails: model.ContactDetails{
				ContactInformation: segments.ContactInformation{
					ContactFunctionCode: "",
					ContactDetails: segments.ContactDetails{
						ContactIdentifier: "",
						ContactName:       "",
					},
				},
				CommunicationContact: segments.CommunicationContact{
					CommunicationAddressIdentifier:    "",
					CommunicationAddressCodeQualifier: "",
				},
			},
		}, {
			NameAddress: segments.NameAddress{
				PartyFunctionCodeQualifier: "BY",
				PartyIdenNameAndAddressDescriptiontificationDetails: segments.PartyIdentificationDetails{
					PartyIdentifier:               "1234567891234",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "9",
				},
				NameAndAddressDescription: segments.NameAndAddressDescription{},
				PartyName: segments.PartyName{
					PartyName:                 "party name",
					PartyNameConditionalOne:   "party name2",
					PartyNameConditionalTwo:   "",
					PartyNameConditionalThree: "",
					PartyNameConditionalFour:  "",
					PartyNameConditionalFive:  "",
					PartyNameCode:             "",
				},
				Street:      segments.Street{},
				CityName:    "",
				Postal:      "",
				CountryCode: "DE",
			},
			ReferenceNumbersParties: model.ReferenceNumber{
				Reference: segments.Reference{
					ReferenceCodeQualifier: "VA",
					ReferenceIdentifier:    "56789123456789",
					DocumentLineIdentifier: "",
					VersionIdentifier:      "",
					RevisionIdentifier:     "",
				},
				DateTimePeriod: segments.DateTimePeriod{},
			},
			ContactDetails: model.ContactDetails{
				ContactInformation: segments.ContactInformation{
					ContactFunctionCode: "PD",
					ContactDetails: segments.ContactDetails{
						ContactIdentifier: "",
						ContactName:       "Contact name",
					},
				},
				CommunicationContact: segments.CommunicationContact{
					CommunicationAddressIdentifier:    "some.mail@email.com",
					CommunicationAddressCodeQualifier: "EM",
				},
			},
		}, {
			NameAddress: segments.NameAddress{
				PartyFunctionCodeQualifier: "IV",
				PartyIdenNameAndAddressDescriptiontificationDetails: segments.PartyIdentificationDetails{
					PartyIdentifier:               "1233411223219",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "9",
				},
				NameAndAddressDescription: segments.NameAndAddressDescription{},
				PartyName: segments.PartyName{
					PartyName:                 "party name",
					PartyNameConditionalOne:   "",
					PartyNameConditionalTwo:   "",
					PartyNameConditionalThree: "",
					PartyNameConditionalFour:  "",
					PartyNameConditionalFive:  "",
					PartyNameCode:             "",
				},
				Street: segments.Street{
					Street:                 "teststreet",
					StreetConditionalOne:   "",
					StreetConditionalTwo:   "",
					StreetConditionalThree: "",
				},
				CityName:    "testcity",
				Postal:      "12345",
				CountryCode: "DE",
			},
			ReferenceNumbersParties: model.ReferenceNumber{},
			ContactDetails:          model.ContactDetails{},
		}, {
			NameAddress: segments.NameAddress{
				PartyFunctionCodeQualifier: "DP",
				PartyIdenNameAndAddressDescriptiontificationDetails: segments.PartyIdentificationDetails{
					PartyIdentifier:               "1233411666666",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "9",
				},
				NameAndAddressDescription: segments.NameAndAddressDescription{},
				PartyName: segments.PartyName{
					PartyName:                 "party name",
					PartyNameConditionalOne:   "party name2",
					PartyNameConditionalTwo:   "",
					PartyNameConditionalThree: "",
					PartyNameConditionalFour:  "",
					PartyNameConditionalFive:  "",
					PartyNameCode:             "",
				},
				Street: segments.Street{
					Street:                 "teststreet",
					StreetConditionalOne:   "",
					StreetConditionalTwo:   "",
					StreetConditionalThree: "",
				},
				CityName:    "testcity",
				Postal:      "12345",
				CountryCode: "DE",
			},
			ReferenceNumbersParties: model.ReferenceNumber{},
			ContactDetails:          model.ContactDetails{},
		}},
		Currencies: model.Currencies{
			Currencies: segments.Currencies{
				CurrencyUsageCodeQualifier:           "2",
				CurrencyIdentificationCode:           "EUR",
				CurrencyTypeCodeQualifier:            "9",
				CurrencyRateValue:                    "",
				ExchangeRateCurrencyMarketIdentifier: "",
			},
			DateTimePeriod: segments.DateTimePeriod{
				DTMFunctionCode: "",
				DTMValue:        "",
				DTMFormatCode:   "",
			},
		},
		Items: []model.Item{{
			LineItem: segments.LineItem{
				LineItemIdentifier:         "1",
				ActionCode:                 "",
				ItemNumberIdentification:   segments.ItemNumberIdentification{},
				SublineInformation:         segments.SublineInformation{},
				ConfigurationLevelNumber:   "",
				ConfigurationOperationCode: "",
			},
			AdditionalProductID: segments.AdditionalProductID{
				ProductIdentifierCodeQualifier: "555",
				ItemNumberIdentification: segments.ItemNumberIdentification{
					ItemIdentifier:                "123456",
					ItemTypeIdentificationCode:    "SA",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "",
				},
				ItemNumberIdentificationConditionalOne:   segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalTwo:   segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalThree: segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalFour:  segments.ItemNumberIdentification{},
			},
			ItemDescription: segments.ItemDescription{
				DescriptionFormatCode: "F",
				Description:           segments.Description{Description: "Item description"},
			},
			Quantity: segments.Quantity{QuantityDetails: segments.QuantityDetails{
				QuantityTypeCodeQualifier: "21",
				Quantity:                  "1",
				MeasurementUnitCode:       "PCE",
			}},
			PriceInformation: segments.PriceInformation{
				PriceCodeQualifier:  "AAA",
				PriceAmount:         "40.98765",
				UnitPriceBasisValue: "",
				MeasurementUnitCode: "",
			},
		}, {
			LineItem: segments.LineItem{
				LineItemIdentifier:         "2",
				ActionCode:                 "",
				ItemNumberIdentification:   segments.ItemNumberIdentification{},
				SublineInformation:         segments.SublineInformation{},
				ConfigurationLevelNumber:   "",
				ConfigurationOperationCode: "",
			},
			AdditionalProductID: segments.AdditionalProductID{
				ProductIdentifierCodeQualifier: "555",
				ItemNumberIdentification: segments.ItemNumberIdentification{
					ItemIdentifier:                "654321",
					ItemTypeIdentificationCode:    "SA",
					CodeListIdentificationCode:    "",
					CodeListResponsibleAgencyCode: "",
				},
				ItemNumberIdentificationConditionalOne:   segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalTwo:   segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalThree: segments.ItemNumberIdentification{},
				ItemNumberIdentificationConditionalFour:  segments.ItemNumberIdentification{},
			},
			ItemDescription: segments.ItemDescription{
				DescriptionFormatCode: "F",
				Description:           segments.Description{Description: "Item descrition"},
			},
			Quantity: segments.Quantity{QuantityDetails: segments.QuantityDetails{
				QuantityTypeCodeQualifier: "21",
				Quantity:                  "1",
				MeasurementUnitCode:       "PCE",
			}},
			PriceInformation: segments.PriceInformation{
				PriceCodeQualifier:  "AAA",
				PriceAmount:         "20.12346",
				UnitPriceBasisValue: "",
				MeasurementUnitCode: "",
			},
		}},
		SectionControl: segments.SectionControl{SectionIdentification: "S"},
		ControlTotal: []segments.ControlTotal{{
			ControlTotalTypeCodeQualifier: "1",
			ControlTotalQuantity:          "2",
			MeasurementUnitcode:           "",
		}, {
			ControlTotalTypeCodeQualifier: "2",
			ControlTotalQuantity:          "2",
			MeasurementUnitcode:           "",
		}},
		MessageTrailer: segments.MessageTrailer{
			NumberOfSegmentsInMessage: "29",
			MessageReferenceNumber:    "1",
		},
		InterchangeTrailer: segments.InterchangeTrailer{
			InterchangeControlCount:     "1",
			InterchangeControlReference: "ORD12345678",
		},
	}
	tests := []struct {
		name    string
		args    args
		want    *model.OrderMessage
		wantErr bool
	}{
		{name: "Testorder", args: args{messageSegments: p.Segments}, want: &orderWant, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalOrder(tt.args.messageSegments)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
