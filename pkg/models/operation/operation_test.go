package operation

import (
	"reflect"
	"testing"

	v1 "github.com/kubeclipper/kubeclipper/pkg/scheme/core/v1"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/kubeclipper/kubeclipper/pkg/query"
)

func Test_operationOperator_operationFilter(t *testing.T) {
	args := struct {
		obj runtime.Object
		in1 *query.Query
	}{
		obj: &v1.OperationList{
			Items: []v1.Operation{
				{
					Steps: nil,
					Status: v1.OperationStatus{
						Status: v1.OperationStatusRunning,
					},
				},
				{
					Steps: nil,
					Status: v1.OperationStatus{
						Status: v1.OperationStatusFailed,
					},
				},
			},
		},
		in1: nil,
	}
	want := []runtime.Object{
		&v1.Operation{
			Steps: nil,
			Status: v1.OperationStatus{
				Status: v1.OperationStatusRunning,
			},
		},
		&v1.Operation{
			Steps: nil,
			Status: v1.OperationStatus{
				Status: v1.OperationStatusFailed,
			},
		},
	}

	l := &operationOperator{
		storage: nil,
	}
	if got := l.operationFilter(args.obj, args.in1); !reflect.DeepEqual(got, want) {
		t.Errorf("operationFilter() = %v, want %v", got, want)
	}
}
