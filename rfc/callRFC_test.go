package main

import (
	"reflect"
	"testing"
)

func Test_abapSystem(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abapSystem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("abapSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}
