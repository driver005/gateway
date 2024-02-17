package services

import (
	"reflect"
	"testing"
)

func TestNewFlagRouter(t *testing.T) {
	type args struct {
		r Registry
	}
	tests := []struct {
		name string
		args args
		want *FlagRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFlagRouter(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFlagRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRouter_IsFeatureEnabled(t *testing.T) {
	type args struct {
		flag []string
	}
	tests := []struct {
		name string
		m    *FlagRouter
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsFeatureEnabled(tt.args.flag); got != tt.want {
				t.Errorf("FlagRouter.IsFeatureEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRouter_SetFlag(t *testing.T) {
	type args struct {
		flag  string
		value bool
	}
	tests := []struct {
		name string
		m    *FlagRouter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.SetFlag(tt.args.flag, tt.args.value)
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		s []string
		r string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remove(tt.args.s, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}
