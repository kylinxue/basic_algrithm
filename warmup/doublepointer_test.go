package warmup

import "testing"

func Test_reverseWords(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		// TODO: Add test cases.
		{"test-1", " hello world! ", "world! hello"},
		{"test-2", " hello   world! ", "world! hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
