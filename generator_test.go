package mescon

import (
	"log"
	"reflect"
	"testing"
)

func Test_generateHollowLine(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   string
	}{
		{
			name:   "Example 1",
			length: 5,
			want:   "*   *",
		},
		{
			name:   "Example 2",
			length: 2,
			want:   "**",
		},
		{
			name:   "Example 3",
			length: 0,
			want:   "**",
		},
		{
			name:   "Example 4",
			length: 10,
			want:   "*        *",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateHollowLine(tt.length); got != tt.want {
				log.Println()
				t.Errorf("generateHollowLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateHollowLineWithNL(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   string
	}{
		{
			name:   "Example 1",
			length: 5,
			want:   "*   *\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateHollowLineWithNL(tt.length); got != tt.want {
				t.Errorf("generateHollowLineWithNL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateFullLineWithNL(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   string
	}{
		{
			name:   "Example 1",
			length: 5,
			want:   "*****\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFullLineWithNL(tt.length); got != tt.want {
				t.Errorf("generateFullLineWithNL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMultiline(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "Example 1",
			str:  "",
			want: false,
		},
		{
			name: "Example 2",
			str:  "asdajshda \n sdfjsa",
			want: true,
		},
		{
			name: "Example 3",
			str: `sdfa
dfasdf`,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMultiline(tt.str); got != tt.want {
				t.Errorf("isMultiline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleMessage_generateSingleLineMessageMessage(t *testing.T) {
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

func Test_generateFullLine(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   string
	}{
		{
			name:   "Example 1",
			length: 5,
			want:   "*****",
		},
		{
			name:   "Example 2",
			length: 10,
			want:   "**********",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateFullLine(tt.length); got != tt.want {
				t.Errorf("generateFullLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenSingleLineMessage(t *testing.T) {
	type args struct {
		width   int
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Example 1",
			args: args{
				width:   5,
				message: "test test test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Example 2",
			args: args{
				width:   35,
				message: "test test test \n test",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Example 3",
			args: args{
				width:   20,
				message: "test",
			},
			want: `
********************
*                  *
*       test       *
*                  *
********************
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenSingleLineMessage(tt.args.width, tt.args.message)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenSingleLineMessage() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenSingleLineMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
