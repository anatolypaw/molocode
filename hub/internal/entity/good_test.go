package entity

import (
	"testing"
	"time"
)

func TestGood_ValidateDesc(t *testing.T) {
	type fields struct {
		Gtin            string
		Desc            string
		StoreCount      uint
		GetCodeForPrint bool
		AllowProduce    bool
		AllowPrint      bool
		Upload          bool
		CreatedAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"normal 1", fields{Desc: "Молоко"}, false},
		{"normal 2", fields{Desc: "Молоко 5% красное красивое"}, false},
		{"short", fields{Desc: "Мо"}, true},
		{"long", fields{Desc: "Молоко творог сметана кефир ряженка вся"}, true},
		{"empty", fields{Desc: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ths := &Good{
				Gtin:            tt.fields.Gtin,
				Desc:            tt.fields.Desc,
				StoreCount:      tt.fields.StoreCount,
				GetCodeForPrint: tt.fields.GetCodeForPrint,
				AllowProduce:    tt.fields.AllowProduce,
				AllowPrint:      tt.fields.AllowPrint,
				Upload:          tt.fields.Upload,
				CreatedAt:       tt.fields.CreatedAt,
			}
			if err := ths.ValidateDesc(); (err != nil) != tt.wantErr {
				t.Errorf("Good.ValidateDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGood_ValidateGtin(t *testing.T) {
	type fields struct {
		Gtin            string
		Desc            string
		StoreCount      uint
		GetCodeForPrint bool
		AllowProduce    bool
		AllowPrint      bool
		Upload          bool
		CreatedAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"normal 1", fields{Gtin: "01234567890123"}, false},
		{"normal 2", fields{Gtin: "00000000000000"}, false},
		{"short", fields{Gtin: "0000000000003"}, true},
		{"long", fields{Gtin: "000000000000005"}, true},
		{"empty", fields{Gtin: ""}, true},
		{"unacept symbols 1", fields{Gtin: " 0000000000003"}, true},
		{"unacept symbols 2", fields{Gtin: "A0000000000003"}, true},
		{"unacept symbols 3", fields{Gtin: "ы0000000000003"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ths := &Good{
				Gtin:            tt.fields.Gtin,
				Desc:            tt.fields.Desc,
				StoreCount:      tt.fields.StoreCount,
				GetCodeForPrint: tt.fields.GetCodeForPrint,
				AllowProduce:    tt.fields.AllowProduce,
				AllowPrint:      tt.fields.AllowPrint,
				Upload:          tt.fields.Upload,
				CreatedAt:       tt.fields.CreatedAt,
			}
			if err := ths.ValidateGtin(); (err != nil) != tt.wantErr {
				t.Errorf("Good.ValidateGtin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
