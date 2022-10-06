package bcd

import (
	"reflect"
	"testing"
)

var one = Number{
	Sign:     0,
	Exponent: 0,
	Digits:   [12]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var two = Number{
	Sign:     0,
	Exponent: 0,
	Digits:   [12]byte{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var ten = Number{
	Sign:     0,
	Exponent: 1,
	Digits:   [12]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var twelve = Number{
	Sign:     0,
	Exponent: 1,
	Digits:   [12]byte{1, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var ninetyNine = Number{
	Sign:     0,
	Exponent: 1,
	Digits:   [12]byte{9, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func TestAdd(t *testing.T) {
	type args struct {
		x Number
		y Number
	}
	tests := []struct {
		name string
		args args
		want Number
	}{
		{
			name: "1+1",
			args: args{x: one, y: one},
			want: two,
		},
		{
			name: "10+2",
			args: args{x: ten, y: two},
			want: twelve,
		},
		{
			name: "2+10",
			args: args{x: two, y: ten},
			want: twelve,
		},
		{
			name: "12+0",
			args: args{x: twelve, y: Number{}},
			want: twelve,
		},
		{
			name: "0+12",
			args: args{x: Number{}, y: twelve},
			want: twelve,
		},
		{
			name: "10+12",
			args: args{x: ten, y: twelve},
			want: Number{Sign: 0, Exponent: 1, Digits: [12]byte{2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
		{
			name: "99+1",
			args: args{x: ninetyNine, y: one},
			want: Number{Sign: 0, Exponent: 2, Digits: [12]byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
