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

// The WebSocket heartbeats listener callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type webSockeHeartbeatsListener func(event *WebSocketHeartbeatMessage)

// Heartbeat channels are opened for product IDs subscribed via the Orders or L2 channels.
func (c *clientImpl) webSocketHeartbeatsSubscribe(
	productIds []string,
	callback webSockeHeartbeatsListener,
) error {

	return nil
}

// This is called internally when there are no Orders or L2 subscriptions for
// particular product IDs.
func (c *clientImpl) webSocketHeartbeatsUnsubscribe(productIds []string) error {
	return nil
}
