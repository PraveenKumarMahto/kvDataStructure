package main

import (
	"testing"
)

var keyValueMap [MAP_SIZE]*ArrayNode

func TestInsertOrUpdate(t *testing.T) {
	type args struct {
		key         int64
		value       string
		keyValueMap *[MAP_SIZE]*ArrayNode
	}
	tests := []struct {
		name string
		args args
	}{

		{
			args: args{
				key:         1,
				value:       "hello world",
				keyValueMap: &keyValueMap,
			}},

		{
			args: args{
				key:         1001,
				value:       "foo bar",
				keyValueMap: &keyValueMap,
			}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertOrUpdate(tt.args.key, tt.args.value, tt.args.keyValueMap)
		})
	}
}

func TestFind(t *testing.T) {
	hello := "hello world"
	foo := "foo bar"
	type args struct {
		key         int64
		keyValueMap *[MAP_SIZE]*ArrayNode
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			args: args{
				key:         1,
				keyValueMap: &keyValueMap,
			},
			want:  hello,
			want1: true,
		},
		{
			args: args{
				key:         1001,
				keyValueMap: &keyValueMap,
			},
			want:  foo,
			want1: true,
		},
		{
			args: args{
				key:         -1002,
				keyValueMap: &keyValueMap,
			},
			want:  foo,
			want1: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Find(tt.args.key, tt.args.keyValueMap)
			if got != nil && *got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPrintAllMapKeyValue(t *testing.T) {
	type args struct {
		keyValueMap *[MAP_SIZE]*ArrayNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				keyValueMap: &keyValueMap,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAllMapKeyValue(tt.args.keyValueMap)
		})
	}
}
