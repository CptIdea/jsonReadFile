package utils

import (
	"reflect"
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	type args struct {
		path string
		out  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"1,2,3",
			args{
				path: "examples/array.json",
				out:  new([]int),
			},
			&[]int{1, 2, 3},
			false,
		},
		{
			"object",
			args{
				path: "examples/object.json",
				out:  new(map[string]string),
			},
			&map[string]string{"testing": "cool"},
			false,
		},
		{
			"file not found",
			args{
				path: "examples/404.json",
				out:  nil,
			},
			nil,
			true,
		},
		{
			"unmarshal error",
			args{
				path: "examples/array.json",
				out:  nil,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if err := ReadJsonFile(tt.args.path, tt.args.out); (err != nil) != tt.wantErr {
					t.Errorf("ReadJsonFile() error = %v, wantErr %v", err, tt.wantErr)
				}

				if !reflect.DeepEqual(tt.want, tt.args.out) {
					t.Errorf("ReadJsonFile() want = %v, got %v", tt.want, tt.args.out)
				}
			},
		)
	}
}
