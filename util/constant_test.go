package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/scalar"
	"github.com/z7zmey/php-parser/position"
	"reflect"
	"testing"
)

func TestParseConstant(t *testing.T) {
	type args struct {
		node node.Node
		vars map[string]node.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "number",
			args: args{node: &scalar.Dnumber{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    69,
				},
				Value: "123456",
			},
			},
			want: int64(123456),
		},
		{
			name: "Long number",
			args: args{
				node: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    22,
					},
					Value: "1234567890123456789",
				},

				vars: nil,
			},
			want: int64(1234567890123456789),
		},

		{
			name: "bin number",
			args: args{
				node: &scalar.Dnumber{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    69,
					},
					Value: "0b11111",
				},
				vars: nil,
			},
			want: int64(31),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConstant(tt.args.node, tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConstant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConstant() got = %v, want %v", got, tt.want)
			}
		})
	}
}
