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

const defaultWebSocketUrl = "wss://ws-feed.prime.coinbase.com"

type WebSockeL2Listener func(event *WebSocketOrdersMessage)

type WebSockeOrderListener func(event *WebSocketL2Message)

type WebSockeHeartbeatListener func(event *WebSocketHeartbeatMessage)

type WebSockeErrorListener func(event *WebSocketErrorMessage)

func (c *clientImpl) WebSocketL2Subscribe(
	productIds []string,
	callback WebSockeL2Listener,
) error {

	return nil
}

// Unsubscribe to the L2 data. Pass in specific product IDs or none to
// unsubscribe from all.
func (c *clientImpl) WebSocketL2Unsubscribe(productIds []string) error {

	return nil
}

func (c *clientImpl) WebSocketOrdersSubscribe(
	productIds []string,
	callback WebSockeOrderListener,
) error {

	return nil
}

// Unsubscribe to the orders data. Pass in specific product IDs or none to
// unsubscribe from all.
func (c *clientImpl) WebSocketOrdersUnsubscribe(productIds []string) error {

	return nil
}

func (c *clientImpl) WebSocketHeartbeatSubscribe(
	productIds []string,
	callback WebSockeHeartbeatListener,
) error {

	return nil
}

// Unsubscribe to the heartbeat channel. Pass in specific product IDs or none to
// unsubscribe from all.
func (c *clientImpl) WebSocketHeartbeatUnsubscribe(productIds []string) error {

	return nil
}

func (c *clientImpl) AddErrorListener(callback WebSockeErrorListener) error {

	return nil
}

func (c *clientImpl) WebSocketUrl() string {
	return c.webSocketUrl
}

func (c *clientImpl) SetWebSocketUrl(u string) Client {
	c.webSocketUrl = u
	return c
}
