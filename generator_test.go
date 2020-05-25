package mescon

import (
	"reflect"
	"testing"
)

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

func Test_wrapMessage(t *testing.T) {
	type args struct {
		width      int
		messageLen int
		message    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				width:      20,
				messageLen: 4,
				message:    "test",
			},
			want: "*       test       *",
		},
		{
			name: "Example 2",
			args: args{
				width:      10,
				messageLen: 2,
				message:    "aa",
			},
			want: "*   aa   *",
		},
		{
			name: "Example 3",
			args: args{
				width:      7,
				messageLen: 2,
				message:    "aa",
			},
			want: "* aa  *",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapMessage(tt.args.width, tt.args.messageLen, tt.args.message); got != tt.want {
				t.Errorf("wrapMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMessageLength(t *testing.T) {
	tests := []struct {
		name string
		mes  string
		want int
	}{
		{
			name: "Example 1",
			mes:  "12345",
			want: 5,
		},
		{
			name: "Example 2",
			mes:  "test",
			want: 4,
		},
		{
			name: "Example 3",
			mes:  "тест",
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMessageLength(tt.mes); got != tt.want {
				t.Errorf("getMessageLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenMultiLineMessage(t *testing.T) {
	type args struct {
		width     int
		message   string
		separator string
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
				message: "lets test this",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Example 2",
			args: args{
				width:     20,
				message:   "lets test this\ntest\nmessage",
				separator: "",
			},
			want: `
********************
*                  *
*  lets test this  *
*       test       *
*     message      *
********************
`,
			wantErr: false,
		},
		{
			name: "Example 3",
			args: args{
				width:     20,
				message:   "test message<<one more<<and one more<<and the last one",
				separator: "<<",
			},
			want: `
********************
*                  *
*   test message   *
*     one more     *
*   and one more   *
* and the last one *
********************
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenMultiLineMessage(tt.args.width, tt.args.message, tt.args.separator)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenMultiLineMessage() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenMultiLineMessage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenAny(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    string
		wantErr bool
	}{
		{
			name:    "Example 1",
			message: "test",
			want: `
************
*          *
*   test   *
*          *
************
`,
			wantErr: false,
		},
		{
			name:    "Example 2",
			message: "sdfajsdhfhasfhgsahdfghjasdgfhjsagfjhgdhfgasdjkfkasdvfjsvdfvdnvfjdfasdfhsdjfjkashfljahsdjfhalshflkahsflkhsakfhkalhdfklahsdklfhasldhf",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Example 3",
			message: "test\ntest",
			want: `
************
*          *
*   test   *
*   test   *
*          *
************
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenAny(tt.message)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GenAny() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenAny() got = %v, want %v", got, tt.want)
			}
		})
	}
}
