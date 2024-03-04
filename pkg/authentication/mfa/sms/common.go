package sms

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kubeclipper/kubeclipper/pkg/simple/client/cache"
)

var (
	ErrSMSRateLimitExceeded = fmt.Errorf("verification code was sent too frequently. Please try again later")
	ErrSMSSendLimitExceeded = fmt.Errorf("SMS sending limit reached")
)

const (
	smsSendInterval = time.Minute
)

func smsCacheKey(Type, phone string) string {
	return fmt.Sprintf("%s-%s", Type, phone)
}

func generateNumberCode(n int) string {
	if n <= 0 {
		return ""
	}
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = '0' + byte(random.Intn(10))
	}
	return string(buf)
}

func rateLimit(kv cache.Interface, key string, expire time.Duration) error {
	exist, err := kv.Exist(key)
	if err != nil {
		return err
	}
	if exist {
		return ErrSMSRateLimitExceeded
	}
	err = kv.Set(key, time.Now().String(), expire)
	if err != nil {
		return err
	}
	return nil
}
