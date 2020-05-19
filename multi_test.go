package mescon

import (
	"reflect"
	"testing"
)

func Test_getMessageLengthInMessages(t *testing.T) {
	tests := []struct {
		name string
		m    []string
		want int
	}{
		{
			name: "Example 1",
			m:    []string{"abc", "abcd", "abcde"},
			want: 5,
		},
		{
			name: "Example 2",
			m:    []string{"abc", "abc", "abc"},
			want: 3,
		},
		{
			name: "Example 3",
			m:    []string{"abcd"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessageLengthInMessages(tt.m); got != tt.want {
				t.Errorf("getMessageLengthInMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessagesFromString(t *testing.T) {
	tests := []struct {
		name string
		mes  string
		want []string
	}{
		{
			name: "Example 1",
			mes:  "test",
			want: []string{"test"},
		},
		{
			name: "Example 2",
			mes:  "first\nsecond\nthird",
			want: []string{"first", "second", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessagesFromString(tt.mes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMessagesFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessagesAndLength(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    []string
		want1   int
	}{
		{
			name:    "Example 1",
			message: "testing",
			want:    []string{"testing"},
			want1:   7,
		},
		{
			name:    "Example 2",
			message: "test\ntest\ntest",
			want:    []string{"test", "test", "test"},
			want1:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getMessagesAndLength(tt.message)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMessagesAndLength() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getMessagesAndLength() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_multiLineMessage_generateMultiLineMessage(t *testing.T) {
	type fields struct {
		width            int
		maxMessageLength int
		messages         []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				width:            15,
				maxMessageLength: 4,
				messages:         []string{"test", "test", "test"},
			},
			want: `
***************
*             *
*    test     *
*    test     *
*    test     *
*             *
***************
`,
		},
		{
			name: "Example 2",
			fields: fields{
				width:            10,
				maxMessageLength: 4,
				messages:         []string{"lolg", "test"},
			},
			want: `
***********
*         *
*  lolg   *
*  test   *
*         *
***********
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := multiLineMessage{
				width:            tt.fields.width,
				maxMessageLength: tt.fields.maxMessageLength,
				messages:         tt.fields.messages,
			}
			if got := mm.generateMultiLineMessage(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMultiLineMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
