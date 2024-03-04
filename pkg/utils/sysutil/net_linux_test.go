package sysutil

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_netInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    []Net
		wantErr bool
	}{
		{
			name:    "get net info test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := netInfo()
			fmt.Printf("got: %#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("netInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("netInfo() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestCUPInfo1(t *testing.T) {
	tests := []struct {
		name    string
		want    CPU
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CUPInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("CUPInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CUPInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_netInfo1(t *testing.T) {
	tests := []struct {
		name    string
		want    []Net
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := netInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("netInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("netInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
