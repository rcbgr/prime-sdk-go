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

type WebSockeErrorCallback func(event *WebSocketErrorMessage)

// The WebSocket l2 data callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type WebSockeL2Callback func(event *WebSocketOrdersMessage)

// The WebSocket orders listener callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type WebSockeOrderCallback func(event *WebSocketL2Message)

// The WebSocket heartbeat callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type webSockeHeartbeatCallback func(event *WebSocketHeartbeatMessage)

func (c *clientImpl) WebSocketSubscribe() error {
	return nil
}

func (c *clientImpl) WebSocketUnsubscribe() error {

	return nil
}

// Heartbeat channels are opened for product IDs subscribed via the Orders or L2 channels.
func (c *clientImpl) webSocketHeartbeatsSubscribe(callback webSockeHeartbeatCallback) error {

	// TODO
	return nil
}

// This is called internally when there are Orders or L2 subscriptions for
// particular product IDs.
func (c *clientImpl) webSocketHeartbeatsUnsubscribe() error {
	// TODO
	return nil
}

func (c *clientImpl) SetErrorCallback(callback WebSockeErrorCallback) Client {
	c.webSocket.errorCallback = callback
	return c
}

func (c *clientImpl) SetWebSockeL2Callback(callback WebSockeL2Callback) Client {
	c.webSocket.l2Callback = callback
	return c

}

func (c *clientImpl) SetWebSockeOrderCallback(callback WebSockeOrderCallback) Client {
	c.webSocket.orderCallback = callback
	return c

}

func (c *clientImpl) SetWebSocketProductIds(productIds []string) Client {
	c.webSocket.productIds = productIds
	return c
}

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
