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
	"sync"

	"github.com/coinbase-samples/core-go"
)

const (
	webSocketHeartbeats = "heartbeats"
	webSocketOrders     = "orders"
	webSocketL2         = "l2_data"
)

type webSocketState struct {
	sync.Mutex
	conn          *core.WebSocketConnection
	subscriptions []*webSocketSubscription
}

/*
func (s *webSocketState) connect() error {
	s.Lock()
	defer s.Unlock()

	if s.conn != nil {
		c = true
	}

	return
}
*/

func newWebSocketState() *webSocketState {
	return &webSocketState{}
}

type webSocketSubscription struct {
	channel    string
	productIds []string
}
