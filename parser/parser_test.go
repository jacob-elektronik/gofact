package parser

import (
	"gofact/utils"
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	type args struct {
		message       string
		printSegments bool
		printTokens   bool
		subSet        string
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		{name:"test1", args:args{message:"", printSegments:false, printTokens:false, subSet:""}, want:NewParser("", false, false, utils.SubSetDefault)},
		{name:"test2", args:args{message:"", printSegments:false, printTokens:false, subSet:"eancom"}, want:NewParser("", false, false, utils.SubSetEancom)},
		{name:"test3", args:args{message:"", printSegments:false, printTokens:false, subSet:"something"}, want:nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.message, tt.args.printSegments, tt.args.printTokens, tt.args.subSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseEdiFactMessageConcurrent(t *testing.T) {
	type fields struct {
		EdiFactMessage            string
		printSegments             bool
		printTokens               bool
		subSet                    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"test1", fields{
			EdiFactMessage:            "../edi_messages/eancom_ord.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "",
		}, true},
		{"test2", fields{
			EdiFactMessage:            "../edi_messages/eancom_ord.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "eancom",
		}, false},
		{"test3", fields{
			EdiFactMessage:            "../edi_messages/message.edi",
			printSegments:             true,
			printTokens:               true,
			subSet:                    "",
		}, false},
		{"test4", fields{
			EdiFactMessage:            "../edi_messages/message_err.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "",
		}, true},
		{"test5", fields{
			EdiFactMessage:            "../edi_messages/message_err2.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "",
		}, true},
		{"test6", fields{
			EdiFactMessage:            "../edi_messages/message_err3.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "",
		}, true},
		{"test7", fields{
			EdiFactMessage:            "../edi_messages/message_err4.edi",
			printSegments:             false,
			printTokens:               false,
			subSet:                    "",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.fields.EdiFactMessage, tt.fields.printSegments, tt.fields.printTokens, tt.fields.subSet)
			if err := p.ParseEdiFactMessageConcurrent(); (err != nil) != tt.wantErr {
				t.Errorf("ParseEdiFactMessageConcurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}