package bcd

import (
	"reflect"
	"testing"
)

func TestAton(t *testing.T) {
	type args struct {
		stringrep string
	}
	tests := []struct {
		name    string
		args    args
		want    Number
		wantErr bool
	}{
		{
			name: "1.23456",
			args: args{stringrep: "1.23456"},
			want: Number{Sign: 0, Exponent: 0, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "0.123456",
			args: args{stringrep: "0.123456"},
			want: Number{Sign: 0, Exponent: -1, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: ".123456",
			args: args{stringrep: ".123456"},
			want: Number{Sign: 0, Exponent: -1, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "123456",
			args: args{stringrep: "123456"},
			want: Number{Sign: 0, Exponent: 5, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "123456000",
			args: args{stringrep: "123456000"},
			want: Number{Sign: 0, Exponent: 8, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "12345678901234567890",
			args: args{stringrep: "12345678901234567890"},
			want: Number{Sign: 0, Exponent: 19, Digits: [12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}},
		},
		{
			name: "1.23456E5",
			args: args{stringrep: "1.23456E5"},
			want: Number{Sign: 0, Exponent: 5, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "1.23456E-1",
			args: args{stringrep: "1.23456E-1"},
			want: Number{Sign: 0, Exponent: -1, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "3E2",
			args: args{stringrep: "3E2"},
			want: Number{Sign: 0, Exponent: 2, Digits: [12]byte{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Aton(tt.args.stringrep)
			if (err != nil) != tt.wantErr {
				t.Errorf("Aton() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumber_String(t *testing.T) {
	type fields struct {
		Sign     int8
		Exponent int8
		Digits   [12]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "1.23456",
			fields: fields{Sign: 0, Exponent: 0, Digits: [12]byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0}},
			want:   "1.23456E0",
		},
		{
			name:   "1.2345E6",
			fields: fields{Sign: 0, Exponent: 6, Digits: [12]byte{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0}},
			want:   "1.2345E6",
		},
		{
			name:   "9.99999999999E99",
			fields: fields{Sign: 0, Exponent: 99, Digits: [12]byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
			want:   "9.99999999999E99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Number{
				Sign:     tt.fields.Sign,
				Exponent: tt.fields.Exponent,
				Digits:   tt.fields.Digits,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("Number.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
