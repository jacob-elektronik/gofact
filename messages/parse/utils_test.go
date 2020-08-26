package parse

import (
	"reflect"
	"testing"
)

func Test_handleIndicator(t *testing.T) {
	type args struct {
		s         string
		indicator string
		delimiter string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				s:         "aa?+aa+aa",
				indicator: "?",
				delimiter: "+",
			},
			want: []string{"aa+aa", "aa"},
		},
		{
			name: "test2",
			args: args{
				s:         "aa+aa+aa",
				indicator: "?",
				delimiter: "+",
			},
			want: []string{"aa", "aa", "aa"},
		},

		{
			name: "test3",
			args: args{
				s:         "aa?+aa?+aa+bb",
				indicator: "?",
				delimiter: "+",
			},
			want: []string{"aa+aa+aa", "bb"},
		},

		{
			name: "test4",
			args: args{
				s:         "aa?+aa+aa+aa?+aa",
				indicator: "?",
				delimiter: "+",
			},
			want: []string{"aa+aa", "aa", "aa+aa",},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleIndicator(tt.args.s, tt.args.indicator, tt.args.delimiter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleIndicator() = %v, want %v", got, tt.want)
			}
		})
	}
}
