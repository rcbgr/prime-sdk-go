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
	"context"
	"sync"
	"time"

	"github.com/coinbase-samples/core-go"
)

const (
	webSocketHeartbeats = "heartbeats"
	webSocketOrders     = "orders"
	webSocketL2         = "l2_data"
)

type webSocket struct {
	sync.Mutex
	conn          *core.WebSocketConnection
	subscriptions []*webSocketSubscription
	connected     bool
	dialerConfig  core.DialerConfig
	url           string
	dialTimeout   time.Duration
}

func (s *webSocket) closeHandler(code int, text string) error {
	s.Lock()
	defer s.Unlock()

	// TODO: Send a close error message to the listener
	s.connected = false
	return nil
}

func (s *webSocket) subscribe(productIds []string, channel string, msg interface{}) error {
	if err := s.connect(s.dialTimeout); err != nil {
		return err
	}

	// check to see if the product ids are already subscribed

	// ensure the listener is running

	// Send the message

}

// connect creates a new WebSocket connection. If the connection is already
// open, it does nothing.
func (s *webSocket) connect(timeout time.Duration) error {
	s.Lock()
	defer s.Unlock()

	if s.connected {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := core.DialWebSocket(ctx, s.dialerConfig)
	if err != nil {
		return err
	}

	s.conn.SetCloseHandler(s.closeHandler)

	s.conn = conn
	s.connected = true

	return nil
}

func newWebSocket(url string) *webSocket {
	return &webSocket{
		dialerConfig: core.DefaultDialerConfig(url),
		dialTimeout:  5 * time.Second,
	}
}

type webSocketSubscription struct {
	channel    string
	productIds []string
}
