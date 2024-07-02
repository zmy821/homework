package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// 用本地缓存实现
type LocalCodeCache struct {
	cache sync.Map
}

func NewLocalCodeCache() CodeCache {
	return &LocalCodeCache{}
}

func (c *LocalCodeCache) Set(ctx context.Context, biz, phone, code string) error {
	// 模拟设置操作，根据业务需求调整过期逻辑
	c.cache.Store(c.key(biz, phone), code)
	return nil
}

func (c *LocalCodeCache) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	value, ok := c.cache.Load(c.key(biz, phone))
	if !ok {
		return false, nil // 未找到验证码
	}

	storedCode, ok := value.(string)
	if !ok {
		return false, errors.New("缓存数据格式错误")
	}

	if storedCode == code {
		return true, nil
	} else {
		return false, ErrCodeVerifyTooMany
	}
}

func (c *LocalCodeCache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, phone)
}
