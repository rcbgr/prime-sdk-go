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
	"encoding/json"
	"time"
)

type WebSocketSubscribeMessage struct {
	Type         string   `json:"type"`
	Channel      string   `json:"channel"`
	AccessKey    string   `json:"access_key"`
	SvcAccountId string   `json:"api_key_id"`
	Timestamp    string   `json:"timestamp"`
	Passphrase   string   `json:"passphrase"`
	Signature    string   `json:"signature"`
	PortfolioId  string   `json:"portfolio_id,omitempty"`
	ProductIds   []string `json:"product_ids,omitempty"`
}

type WebSocketUnsubscribeMessage struct {
	Type     string   `json:"type"`
	Channels []string `json:"channels"`
}

type HeartbeatMessage struct {
	Channel     string `json:"channel"`
	Timestamp   string `json:"timestamp"`
	SequenceNum int    `json:"sequence_num"`
}

func SubscriptionMsg(credentials *Credentials, channel string, productIds []string) ([]byte, error) {
	t := time.Now().UTC().Format(time.RFC3339)

	msg := &WebSocketSubscribeMessage{
		Type:         "subscribe",
		Channel:      channel,
		AccessKey:    credentials.AccessKey,
		SvcAccountId: credentials.SvcAccountId,
		Timestamp:    t,
		Passphrase:   credentials.Passphrase,
		Signature: signWebSocket(
			channel,
			credentials.PortfolioId,
			credentials.SvcAccountId,
			t,
			credentials.AccessKey,
			credentials.SigningKey,
			productIds,
		),
		PortfolioId: credentials.PortfolioId,
	}

	if len(productIds) > 0 {
		msg.ProductIds = productIds
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func HeartbeatSubscriptionMsg(credentials *Credentials) ([]byte, error) {
	return SubscriptionMsg(credentials, "heartbeats", nil)
}

func OrderSubscriptionMsg(credentials *Credentials, productIds []string) ([]byte, error) {
	return SubscriptionMsg(credentials, "orders", productIds)
}

func L2OrderDataSubscriptionMsg(credentials *Credentials, productIds []string) ([]byte, error) {
	return SubscriptionMsg(credentials, "l2_data", productIds)
}
