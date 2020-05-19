package mescon

import (
	"log"
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
