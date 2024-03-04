package auditing

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/kubeclipper/kubeclipper/pkg/auditing/option"

	"k8s.io/apiserver/pkg/apis/audit"

	"github.com/kubeclipper/kubeclipper/pkg/server/request"
)

func Test_auditing_AddBackend(t *testing.T) {

	type args struct {
		backend Backend
	}
	tests := []struct {
		name string
		args args
		want []Backend
	}{
		{
			name: "addbackend test",
			args: args{
				backend: ConsoleBackend{},
			},
			want: []Backend{
				ConsoleBackend{},
				ConsoleBackend{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auditing{
				backends: []Backend{
					ConsoleBackend{},
				},
			}
			a.AddBackend(tt.args.backend)
			if !reflect.DeepEqual(a.backends, tt.want) {
				t.Errorf("AddBackend() = %v, want = %v,", a.backends, tt.want)
			}
		})
	}
}

func Test_auditing_Enabled(t *testing.T) {
	type fields struct {
		level    audit.Level
		backends []Backend
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "enabled testing",
			fields: fields{
				level:    audit.LevelMetadata,
				backends: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				level:    "other",
				backends: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auditing{
				backends: tt.fields.backends,
				auditOptions: &option.AuditOptions{
					AuditLevel: tt.fields.level,
				},
			}
			if got := a.Enabled(); got != tt.want {
				t.Errorf("Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auditing_LogRequestObject(t *testing.T) {
	type fields struct {
		level    audit.Level
		backends []Backend
	}
	type args struct {
		req  *http.Request
		info *request.Info
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "LogRequestObject test",
			fields: fields{
				level:    audit.LevelRequest,
				backends: nil,
			},
			args: args{
				req: &http.Request{
					Body: io.NopCloser(bytes.NewBufferString("testBody")),
					Header: map[string][]string{
						"Content-Type": {"application/json"},
					},
					URL: &url.URL{
						RawQuery: "testQuery=test",
					},
					ContentLength: 10,
				},
				info: &request.Info{
					IsResourceRequest: false,
					Path:              "testPath",
					Verb:              "create",
					APIGroup:          "testAPIGroup",
					APIVersion:        "testVersion",
					Resource:          "testResource",
					Subresource:       "testSubresource",
					Name:              "testname",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auditing{
				backends: tt.fields.backends,
				auditOptions: &option.AuditOptions{
					AuditLevel: tt.fields.level,
				},
			}
			_ = a.LogRequestObject(tt.args.req, tt.args.info)
		})
	}
}

func Test_auditing_LogResponseObject(t *testing.T) {
	type fields struct {
		level    audit.Level
		backends []Backend
	}
	type args struct {
		e    *audit.Event
		resp *ResponseCapture
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "LogResponse test",
			fields: fields{
				level:    audit.LevelMetadata,
				backends: make([]Backend, 0),
			},
			args: args{
				e: &audit.Event{},
				resp: &ResponseCapture{
					status: 200,
					body:   bytes.NewBuffer([]byte("")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &auditing{
				backends: tt.fields.backends,
				auditOptions: &option.AuditOptions{
					AuditLevel: tt.fields.level,
				},
			}
			a.LogResponseObject(tt.args.e, tt.args.resp)
		})
	}
}
