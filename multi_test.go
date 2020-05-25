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
		sep  string
		want []string
	}{
		{
			name: "Example 1",
			mes:  "test",
			sep:  "\n",
			want: []string{"test"},
		},
		{
			name: "Example 2",
			mes:  "first\nsecond\nthird",
			sep:  "\n",
			want: []string{"first", "second", "third"},
		},
		{
			name: "Example 3",
			mes:  "test||reft||third",
			sep:  "||",
			want: []string{"test", "reft", "third"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessagesFromString(tt.mes, tt.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMessagesFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessagesAndLength(t *testing.T) {
	tests := []struct {
		name      string
		message   string
		separator string
		want      []string
		want1     int
	}{
		{
			name:      "Example 1",
			message:   "testing",
			separator: "\n",
			want:      []string{"testing"},
			want1:     7,
		},
		{
			name:      "Example 2",
			message:   "test\ntest\ntest",
			separator: "\n",
			want:      []string{"test", "test", "test"},
			want1:     4,
		},
		{
			name:      "Example 3",
			message:   "first/second/third",
			separator: "/",
			want:      []string{"first", "second", "third"},
			want1:     6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getMessagesAndLength(tt.message, tt.separator)
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
