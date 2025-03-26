package qxDao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type (
	RedisDao interface {
		GetRD() *redis.Client
		Ping() error
		Close() error
		Set(ctx context.Context, key string, value interface{}) error
		SetEx(ctx context.Context, key string, value interface{}, seconds int) error
		Get(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)
		NewWatcher(ctx context.Context, channel string, fn func(msg *redis.Message)) error
		Publish(ctx context.Context, channel string, msg string) error
	}
	defaultRedisDao struct {
		rd *redis.Client
	}
)

func NewRedisDao(rd *redis.Client) RedisDao {
	return &defaultRedisDao{
		rd: rd,
	}
}

func (d *defaultRedisDao) GetRD() *redis.Client {
	return d.rd
}

func (d *defaultRedisDao) Ping() error {
	// note: 设置5秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return d.rd.Ping(ctx).Err()
	}
	return nil
}

func (d *defaultRedisDao) Close() error {
	return d.rd.Close()
}

func (d *defaultRedisDao) Set(ctx context.Context, key string, value interface{}) error {
	if err := d.rd.Set(ctx, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (d *defaultRedisDao) SetEx(ctx context.Context, key string, value interface{}, seconds int) error {
	if err := d.rd.SetEx(ctx, key, value, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}
	return nil
}

func (d *defaultRedisDao) Get(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	rdCmd := d.rd.Get(ctx, key)
	if err := rdCmd.Err(); err != nil {
		return "", err
	}
	return rdCmd.String(), nil
}

func (d *defaultRedisDao) NewWatcher(ctx context.Context, channel string, fn func(msg *redis.Message)) error {
	sub := d.rd.Subscribe(ctx, channel)
	ch := sub.Channel()
	for msg := range ch {
		// 处理消息
		fn(msg)
	}
	return nil
}

func (d *defaultRedisDao) Publish(ctx context.Context, channel string, msg string) error {
	if err := d.rd.Publish(ctx, channel, msg).Err(); err != nil {
		return err
	}
	return nil
}
