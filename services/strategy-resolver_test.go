package services

import (
	"context"
	"reflect"
	"testing"
)

func TestNewStrategyResolverService(t *testing.T) {
	tests := []struct {
		name string
		want *StrategyResolverService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStrategyResolverService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStrategyResolverService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrategyResolverService_SetContext(t *testing.T) {
	type args struct {
		context context.Context
	}
	tests := []struct {
		name string
		s    *StrategyResolverService
		args args
		want *StrategyResolverService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetContext(tt.args.context); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrategyResolverService.SetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
