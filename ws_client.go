/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prime

import (
	"time"

	"github.com/coinbase-samples/core-go"
)

const defaultWebSocketUrl = "wss://ws-feed.prime.coinbase.com"

func (c *clientImpl) SetWebSocketDialerConfig(conf core.DialerConfig) Client {
	if len(conf.Url) == 0 || conf.Url != c.webSocket.url {
		conf.Url = c.webSocket.url
	}
	c.webSocket.dialerConfig = conf
	return c
}

func (c *clientImpl) WebSocketUrl() string {
	return c.webSocket.url
}

func (c *clientImpl) SetWebSocketUrl(url string) Client {
	c.webSocket.url = url
	c.webSocket.dialerConfig.Url = url
	return c
}

func (c *clientImpl) SetWebSocketDialTimeout(t time.Duration) Client {
	c.webSocket.dialTimeout = t
	return c
}

func (c *clientImpl) WebSocketDialTimeout() time.Duration {
	return c.webSocket.dialTimeout
}
