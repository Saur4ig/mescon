package mescon

import (
	"reflect"
	"testing"
)

func Test_singleMessage_generateSingleLineMessage(t *testing.T) {
	type fields struct {
		width         int
		message       string
		messageLength int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Example 1",
			fields: fields{
				width:         10,
				message:       "test",
				messageLength: 4,
			},
			want: `
**********
*        *
*  test  *
*        *
**********
`,
		},
		{
			name: "Example 2",
			fields: fields{
				width:         30,
				message:       "test is - test",
				messageLength: 14,
			},
			want: `
******************************
*                            *
*       test is - test       *
*                            *
******************************
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := singleMessage{
				width:         tt.fields.width,
				message:       tt.fields.message,
				messageLength: tt.fields.messageLength,
			}
			if got := sm.generateSingleLineMessage(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSingleLineMessageMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
