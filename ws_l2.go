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

// The WebSocket l2 data callback. This function must return ASAP because
// it will block receiving additional messages. The best practice is to place the message
// on an ordered queue and process asynchronously.
type WebSockeL2Callback func(event *WebSocketOrdersMessage)

func (c *clientImpl) WebSocketL2Subscribe(
	productIds []string,
	callback WebSockeL2Callback,
) error {

	return nil
}

// Unsubscribe to the L2 data. Pass in specific product IDs or none to
// unsubscribe from all.
func (c *clientImpl) WebSocketL2Unsubscribe(productIds []string) error {

	return nil
}
