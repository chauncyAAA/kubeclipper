package ipvsutil

import (
	"testing"
)

func TestClear(t *testing.T) {
	t.Skip("can't run in github action")

	type args struct {
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "base",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Clear(tt.args.dryRun); err != nil {
				t.Errorf("Clear() error = %v", err)
			}
		})
	}
}

func TestCreateIPVS(t *testing.T) {
	t.Skip("can't run in github action")

	type args struct {
		vs     *VirtualServer
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				vs: &VirtualServer{
					Address: "169.254.169.100",
					Port:    6443,
					RealServers: []RealServer{
						{Address: "172.18.94.111", Port: 6433},
						{Address: "172.18.94.200", Port: 6433},
						{Address: "172.18.94.114", Port: 6433},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateIPVS(tt.args.vs, tt.args.dryRun); err != nil {
				t.Errorf("CreateIPVS() error = %v", err)
			}
		})
	}
}

func TestDeleteIPVS(t *testing.T) {
	t.Skip("can't run in github action")

	type args struct {
		vs     *VirtualServer
		dryRun bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				vs: &VirtualServer{
					Address: "169.254.169.100",
					Port:    6443,
					RealServers: []RealServer{
						{Address: "172.18.94.111", Port: 6433},
						{Address: "172.18.94.200", Port: 6433},
						{Address: "172.18.94.114", Port: 6433},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteIPVS(tt.args.vs, tt.args.dryRun); err != nil {
				t.Errorf("DeleteIPVS() error = %v", err)
			}
		})
	}
}
