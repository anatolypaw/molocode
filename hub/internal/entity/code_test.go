package entity

import (
	"testing"
)

func TestValidateSerial(t *testing.T) {
	type args struct {
		serial string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal 1", args{"abcdef"}, false},
		{"normal 2", args{"ABC123"}, false},
		{"normal 3", args{"aAb:23"}, false},
		{"normal 4", args{`!"%&'*`}, false},
		{"normal 5", args{`+-./_,`}, false},
		{"normal 6", args{`:;=<>?`}, false},
		{"short", args{"1234"}, true},
		{"long", args{"abcdefG"}, true},
		{"empty", args{""}, true},
		{"unacept symbols 1", args{"123 45"}, true},
		{"unacept symbol 2", args{`12345\`}, true},
		{"unacept symbol 3", args{`12345(`}, true},
		{"unacept symbol 4", args{`12345)`}, true},
		{"unacept symbol 5", args{`12345й`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSerial(tt.args.serial); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSerial() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCrypto(t *testing.T) {
	type args struct {
		crypto string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"normal 1", args{"abcd"}, false},
		{"normal 2", args{"ABC1"}, false},
		{"normal 3", args{"aAb:"}, false},
		{"normal 4", args{`!"%&`}, false},
		{"normal 5", args{`+-./`}, false},
		{"normal 6", args{`:;=<`}, false},
		{"normal 7", args{`'*_,`}, false},
		{"normal 8", args{`ab>?`}, false},
		{"short", args{"123"}, true},
		{"long", args{"abcdefG"}, true},
		{"empty", args{""}, true},
		{"unacept symbols 1", args{"3 45"}, true},
		{"unacept symbol 2", args{`345\`}, true},
		{"unacept symbol 3", args{`345(`}, true},
		{"unacept symbol 4", args{`345)`}, true},
		{"unacept symbol 5", args{`345й`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCrypto(tt.args.crypto); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCrypto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCode_Validate(t *testing.T) {
	type fields struct {
		Gtin   string
		Serial string
		Crypto string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"nornmal", fields{Gtin: "00000000000000", Serial: "abc123", Crypto: "1234"}, false},
		{"incorrect gtin", fields{Gtin: "0000000000000", Serial: "abc123", Crypto: "1234"}, true},
		{"incorrect serial", fields{Gtin: "00000000000000", Serial: "abc12", Crypto: "1234"}, true},
		{"incorrect crypto", fields{Gtin: "00000000000000", Serial: "abc123", Crypto: "123"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := &Code{
				Gtin:   tt.fields.Gtin,
				Serial: tt.fields.Serial,
				Crypto: tt.fields.Crypto,
			}
			if err := code.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Code.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
