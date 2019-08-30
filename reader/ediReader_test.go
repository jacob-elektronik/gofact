package reader

import (
	"reflect"
	"testing"
)


func TestNewEdiReader(t *testing.T) {
	type args struct {
		fileStr string
	}
	tests := []struct {
		name string
		args args
		want *EdiReader
	}{
		{name:"test1", args: args{fileStr:""}, want:NewEdiReader("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEdiReader(tt.args.fileStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEdiReader() = %v, want %v", got, tt.want)
			}
		})
	}
}