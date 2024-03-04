package sms

import (
	"io"
	"net/url"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apiserver/pkg/authentication/user"

	"github.com/kubeclipper/kubeclipper/pkg/authentication/mfa"
	"github.com/kubeclipper/kubeclipper/pkg/authentication/oauth"
	"github.com/kubeclipper/kubeclipper/pkg/simple/client/cache"
)

func TestFakeSMSProvider(t *testing.T) {
	kv, err := cache.NewMemory()
	require.NoError(t, err)
	err = mfa.SetupWithOptions(kv, &mfa.Options{
		Enabled: true,
		MFAProviders: []mfa.ProviderOptions{
			{Type: FakeSMSProvider, Options: oauth.DynamicOptions{"ttl": "5m"}},
		},
	})
	require.NoError(t, err)

	sms, err := mfa.GetProvider(FakeSMSProvider)
	require.NoError(t, err)

	// intercept stdout
	tempStdout, err := os.CreateTemp("", "")
	require.NoError(t, err)
	defer func() {
		_ = tempStdout.Close()
		_ = os.Remove(tempStdout.Name())
	}()
	stdout := os.Stdout
	os.Stdout = tempStdout
	defer func() {
		os.Stdout = stdout
	}()

	userInfo := &user.DefaultInfo{
		Name:   "",
		UID:    "",
		Groups: nil,
		Extra: map[string][]string{
			"phone": {"13888888888"},
		},
	}

	err = sms.Request(userInfo)
	require.NoError(t, err)

	req := make(url.Values)
	req.Set("code", "xxxxxx")

	err = sms.Verify(req, userInfo)
	require.Error(t, err)

	code := findCodeInFile(tempStdout)
	require.NotEmpty(t, code)
	req.Set("code", code)

	err = sms.Verify(req, userInfo)
	require.NoError(t, err)
}

func findCodeInFile(f *os.File) string {
	_, _ = f.Seek(0, io.SeekStart)
	buf, _ := io.ReadAll(f)
	reg := regexp.MustCompile(`code:(\d+) `)
	res := reg.FindSubmatch(buf)
	if len(res) == 0 {
		return ""
	}
	return string(res[1])
}
