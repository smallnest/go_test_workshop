package s1

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"普通的两个数相加", args{10, 20}, 30},
		{"一个参数是0", args{10, 0}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args // int
		want int  // out
	}{
		{"普通的两个数相减", args{10, 20}, -10},
		{"一个参数是0", args{10, 0}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"普通的两个数相乘", args{10, 20}, 200},
		{"一个参数是0", args{10, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"普通的两个数相除", args{10, 20}, 0},
		// {"一个参数是0", args{10, 0}, int(math.NaN())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}

}
