package check

import "testing"

func TestCheckIsMobile(t *testing.T) {
	type args struct {
		mobileNum string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIsMobile(tt.args.mobileNum); got != tt.want {
				t.Errorf("CheckIsMobile() = %v, want %v", got, tt.want)
			}
		})
	}
}
