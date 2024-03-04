//go:build !linux
// +build !linux

package ipvsutil

import (
	"reflect"
	"testing"
)

func TestClear(t *testing.T) {
	type args struct {
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Clear(tt.args.dryRun); (err != nil) != tt.wantErr {
				t.Errorf("Clear() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateIPVS(t *testing.T) {
	type args struct {
		vs     *VirtualServer
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateIPVS(tt.args.vs, tt.args.dryRun); (err != nil) != tt.wantErr {
				t.Errorf("CreateIPVS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteIPVS(t *testing.T) {
	type args struct {
		vs     *VirtualServer
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteIPVS(tt.args.vs, tt.args.dryRun); (err != nil) != tt.wantErr {
				t.Errorf("DeleteIPVS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultIPVS(t *testing.T) {
	type args struct {
		address string
		port    uint16
	}
	tests := []struct {
		name string
		args args
		want *VirtualServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defaultIPVS(tt.args.address, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultIPVS() = %v, want %v", got, tt.want)
			}
		})
	}
}
