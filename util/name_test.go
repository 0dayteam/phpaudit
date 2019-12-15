package util

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/name"
	"testing"
)

func TestParserName(t *testing.T) {
	type args struct {
		n node.Node
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
		wantErr bool
	}{
		{
			name:    "a name part",
			args:    args{n: &name.Name{Parts: []node.Node{&name.NamePart{Value: "hello"}}}},
			wantRet: "hello",
		},
		{
			name:    "error node type",
			args:    args{n: &node.Root{}},
			wantErr: true,
		},
		{
			name:    "mult name part",
			args:    args{n: &name.Name{Parts: []node.Node{&name.NamePart{Value: "foo"}, &name.NamePart{Value: "bor"}}}},
			wantRet: "foo/bor",
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := ParseName(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("ParserName() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
