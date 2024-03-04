package sysutil

import (
	"testing"
)

func TestDiskInfo(t *testing.T) {
	type args struct {
		byteSize ByteSize
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				byteSize: 2048,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DiskInfo(tt.args.byteSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DiskInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetSysInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "get system info test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetSysInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSysInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHostInfo(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "get host info",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HostInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("HostInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMemoryInfo(t *testing.T) {
	type args struct {
		byteSize ByteSize
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get memory info test",
			args: args{
				byteSize: 2048,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MemoryInfo(tt.args.byteSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
