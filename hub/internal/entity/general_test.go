package entity

import "testing"

func TestValidateGtin(t *testing.T) {
	type args struct {
		gtin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal 1", args{"01234567890123"}, false},
		{"normal 2", args{"00000000000000"}, false},
		{"short", args{"0000000000003"}, true},
		{"long", args{"000000000000005"}, true},
		{"empty", args{""}, true},
		{"unacept symbols 1", args{" 0000000000003"}, true},
		{"unacept symbols 2", args{"A0000000000003"}, true},
		{"unacept symbols 3", args{"ы0000000000003"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateGtin(tt.args.gtin); (err != nil) != tt.wantErr {
				t.Errorf("ValidateGtin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
