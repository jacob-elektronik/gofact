package parser

import (
	"igitlab.jacob.de/ftomasetti/gofact/utils"
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	type args struct {
		message       string
		subSet        string
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		{name:"test1", args:args{message:"", subSet:""}, want:NewParser("", utils.SubSetDefault)},
		{name:"test2", args:args{message:"", subSet:"eancom"}, want:NewParser("", utils.SubSetEancom)},
		{name:"test3", args:args{message:"", subSet:"something"}, want:nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.message, tt.args.subSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseEdiFactMessageConcurrent(t *testing.T) {
	type fields struct {
		EdiFactMessage            string
		subSet                    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"test1", fields{
			EdiFactMessage:            "../edi_messages/eancom_ord.edi",
			subSet:                    "",
		}, true},
		{"test2", fields{
			EdiFactMessage:            "../edi_messages/eancom_ord.edi",
			subSet:                    "eancom",
		}, false},
		{"test3", fields{
			EdiFactMessage:            "../edi_messages/message.edi",
			subSet:                    "",
		}, false},
		{"test4", fields{
			EdiFactMessage:            "../edi_messages/message_err.edi",
			subSet:                    "",
		}, true},
		{"test5", fields{
			EdiFactMessage:            "../edi_messages/message_err2.edi",
			subSet:                    "",
		}, true},
		{"test6", fields{
			EdiFactMessage:            "../edi_messages/message_err3.edi",
			subSet:                    "",
		}, true},
		{"test7", fields{
			EdiFactMessage:            "../edi_messages/message_err4.edi",
			subSet:                    "",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(tt.fields.EdiFactMessage, tt.fields.subSet)
			if err := p.ParseEdiFactMessageConcurrent(); (err != nil) != tt.wantErr {
				t.Errorf("ParseEdiFactMessageConcurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}