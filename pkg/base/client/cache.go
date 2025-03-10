/*
Copyright 2024 The west2-online Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/west2-online/domtok/config"
	"github.com/west2-online/domtok/pkg/errno"
	"github.com/west2-online/domtok/pkg/logger"
)

// NewRedisClient 传入dbName，具体参考 constants 包
func NewRedisClient(db int) (*redis.Client, error) {
	if config.Redis == nil {
		return nil, errors.New("redis config is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       db,
	})
	l := logger.GetRedisLogger()
	redis.SetLogger(l)
	client.AddHook(l)
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("client.NewRedisClient: ping redis failed: %v", err))
	}
	return client, nil
}
