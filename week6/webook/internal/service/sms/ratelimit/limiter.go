package ratelimit

import (
	"basic-go/webook/internal/service/sms"
	"basic-go/webook/pkg/limiter"
	"context"
	"errors"
)

var errLimited = errors.New("触发限流")

var _ sms.Service = &RateLimitSMSService{}

type RateLimitSMSService struct {
	// 被装饰的
	svc     sms.Service
	limiter limiter.Limiter
	key     string
}

type RateLimitSMSServiceV1 struct {
	sms.Service
	limiter limiter.Limiter
	key     string
}

func (r *RateLimitSMSService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	limited, err := r.limiter.Limit(ctx, r.key)
	if err != nil {
		return err
	}
	if limited {
		return errLimited
	}
	return r.svc.Send(ctx, tplId, args, numbers...)
}

// NewRateLimitSMSService 我也是面向接口编程
func NewRateLimitSMSService(svc sms.Service,
	l limiter.Limiter) *RateLimitSMSService {
	return &RateLimitSMSService{
		svc:     svc,
		limiter: l,
		key:     "sms-limiter",
	}
}