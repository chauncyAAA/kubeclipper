package template

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kubeclipper/kubeclipper/pkg/utils/strutil"
)

func TestRender(t *testing.T) {
	tests := []struct {
		tmpl, expected string
		vars           interface{}
	}{
		{
			tmpl:     `{{ b64enc . }}`,
			expected: `YVJNSXdwM05KQzNHMzI0YQ==`,
			vars:     `aRMIwp3NJC3G324a`,
		},
		{
			tmpl:     `{{ b64dec . }}`,
			expected: `aRMIwp3NJC3G324a`,
			vars:     "YVJNSXdwM05KQzNHMzI0YQ==",
		},
	}

	at := New()
	for _, tt := range tests {
		result, err := at.Render(tt.tmpl, tt.vars)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, result, tt.tmpl)
	}
}

func TestRegisterFunc(t *testing.T) {
	at := New()
	at.RegisterFunc("b64enc_", strutil.Base64Encode) // the default Base64 function is `b64enc`

	tests := []struct {
		tmpl, expected string
		vars           interface{}
	}{
		{
			tmpl:     `{{ b64enc_ . }}`,
			expected: `YVJNSXdwM05KQzNHMzI0YQ==`,
			vars:     `aRMIwp3NJC3G324a`,
		},
		{
			tmpl:     `{{ b64dec . }}`,
			expected: `aRMIwp3NJC3G324a`,
			vars:     "YVJNSXdwM05KQzNHMzI0YQ==",
		},
	}

	for _, tt := range tests {
		result, err := at.Render(tt.tmpl, tt.vars)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, result, tt.tmpl)
	}
}
