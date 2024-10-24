/**
 * Copyright 2023-present Coinbase Global, Inc.
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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/coinbase-samples/core-go"
)

const defaultV1ApiBaseUrl = "https://api.prime.coinbase.com/v1"

var defaultHeadersFunc = AddPrimeHeaders

var successStatusCodes = []int{http.StatusOK}

type Client interface {

	// ==========================================================================
	// Client
	// ==========================================================================

	HttpBaseUrl() string
	SetBaseUrl(u string) Client

	HttpClient() *http.Client
	SetHeadersFunc(hf core.HttpHeaderFunc) Client

	Credentials() *Credentials

	// ==========================================================================
	// WebSockets
	// ==========================================================================

	// SetWebSocketUrl sets the URL for the WebSocket connection. It overrides
	// anything passed in the DialerConfig if that is set.
	SetWebSocketUrl(u string) Client
	WebSocketUrl() string

	SetWebSocketDialTimeout(seconds time.Duration) Client
	WebSocketDialTimeout() time.Duration

	// SetWebSocketDialerConfig sets the WebSocket dialer config. You must call
	// SetWebSocketUrl if you to override the default URL.
	SetWebSocketDialerConfig(c core.DialerConfig) Client

	WebSocketL2Subscribe(productIds []string, callback WebSockeL2Callback) error
	WebSocketL2Unsubscribe(productIds []string) error

	WebSocketOrdersSubscribe(productIds []string, callback WebSockeOrderCallback) error
	WebSocketOrdersUnsubscribe(productIds []string) error

	// ==========================================================================
	// Allocations
	// ==========================================================================

	CreatePortfolioAllocations(ctx context.Context, request *CreatePortfolioAllocationsRequest) (*CreatePortfolioAllocationsResponse, error)
	CreatePortfolioNetAllocations(ctx context.Context, request *CreatePortfolioNetAllocationsRequest) (*CreatePortfolioNetAllocationsResponse, error)
	ListPortfolioAllocations(ctx context.Context, request *ListPortfolioAllocationsRequest) (*ListPortfolioAllocationsResponse, error)
	GetPortfolioAllocation(ctx context.Context, request *GetPortfolioAllocationRequest) (*GetPortfolioAllocationResponse, error)
	GetPortfolioNetAllocation(ctx context.Context, request *GetPortfolioNetAllocationRequest) (*GetPortfolioNetAllocationResponse, error)

	// ==========================================================================
	// Invoices
	// ==========================================================================

	ListInvoices(ctx context.Context, request *ListInvoicesRequest) (*ListInvoicesResponse, error)

	// ==========================================================================
	// Assets
	// ==========================================================================

	ListAssets(ctx context.Context, request *ListAssetsRequest) (*ListAssetsResponse, error)

	// ==========================================================================
	// Payment Methods
	// ==========================================================================

	ListEntityPaymentMethods(ctx context.Context, request *ListEntityPaymentMethodsRequest) (*ListEntityPaymentMethodsResponse, error)
	GetEntityPaymentMethod(ctx context.Context, request *GetEntityPaymentMethodRequest) (*GetEntityPaymentMethodResponse, error)

	// ==========================================================================
	// Users
	// ==========================================================================

	ListEntityUsers(ctx context.Context, request *ListEntityUsersRequest) (*ListEntityUsersResponse, error)
	ListPortfolioUsers(ctx context.Context, request *ListPortfolioUsersRequest) (*ListPortfolioUsersResponse, error)

	// ==========================================================================
	// Portfolios
	// ==========================================================================

	ListPortfolios(ctx context.Context, request *ListPortfoliosRequest) (*ListPortfoliosResponse, error)
	GetPortfolio(ctx context.Context, request *GetPortfolioRequest) (*GetPortfolioResponse, error)
	GetPortfolioCredit(ctx context.Context, request *GetPortfolioCreditRequest) (*GetPortfolioCreditResponse, error)

	// ==========================================================================
	// Activities
	// ==========================================================================

	ListActivities(ctx context.Context, request *ListActivitiesRequest) (*ListActivitiesResponse, error)
	GetActivity(ctx context.Context, request *GetActivityRequest) (*GetActivityResponse, error)

	// ==========================================================================
	// Address Book
	// ==========================================================================

	GetAddressBook(ctx context.Context, request *GetAddressBookRequest) (*GetAddressBookResponse, error)
	CreateAddressBookEntry(ctx context.Context, request *CreateAddressBookEntryRequest) (*CreateAddressBookEntryResponse, error)

	// ==========================================================================
	// Balances
	// ==========================================================================

	ListPortfolioBalances(ctx context.Context, request *ListPortfolioBalancesRequest) (*ListPortfolioBalancesResponse, error)
	GetWalletBalance(ctx context.Context, request *GetWalletBalanceRequest) (*GetWalletBalanceResponse, error)
	ListOnchainWalletBalances(ctx context.Context, request *ListOnchainWalletBalancesRequest) (*ListOnchainWalletBalancesResponse, error)

	// ==========================================================================
	// Commission
	// ==========================================================================

	GetPortfolioCommission(ctx context.Context, request *GetPortfolioCommissionRequest) (*GetPortfolioCommissionResponse, error)

	// ==========================================================================
	// Orders
	// ==========================================================================

	ListOpenOrders(ctx context.Context, request *ListOpenOrdersRequest) (*ListOpenOrdersResponse, error)
	CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrderPreview(ctx context.Context, request *CreateOrderRequest) (*CreateOrderPreviewResponse, error)
	ListOrders(ctx context.Context, request *ListOrdersRequest) (*ListOrdersResponse, error)
	GetOrder(ctx context.Context, request *GetOrderRequest) (*GetOrderResponse, error)
	CancelOrder(ctx context.Context, request *CancelOrderRequest) (*CancelOrderResponse, error)
	ListOrderFills(ctx context.Context, request *ListOrderFillsRequest) (*ListOrderFillsResponse, error)
	ListPortfolioFills(ctx context.Context, request *ListPortfolioFillsRequest) (*ListPortfolioFillsResponse, error)

	// ==========================================================================
	// Products
	// ==========================================================================

	ListProducts(ctx context.Context, request *ListProductsRequest) (*ListProductsResponse, error)

	// ==========================================================================
	// Transactions
	// ==========================================================================

	ListPortfolioTransactions(ctx context.Context, request *ListPortfolioTransactionsRequest) (*ListPortfolioTransactionsResponse, error)
	GetTransaction(ctx context.Context, request *GetTransactionRequest) (*GetTransactionResponse, error)
	CreateConversion(ctx context.Context, request *CreateConversionRequest) (*CreateConversionResponse, error)
	ListWalletTransactions(ctx context.Context, request *ListWalletTransactionsRequest) (*ListWalletTransactionsResponse, error)
	CreateWalletTransfer(ctx context.Context, request *CreateWalletTransferRequest) (*CreateWalletTransferResponse, error)
	CreateWalletWithdrawal(ctx context.Context, request *CreateWalletWithdrawalRequest) (*CreateWalletWithdrawalResponse, error)

	// ==========================================================================
	// Wallets
	// ==========================================================================

	ListWallets(ctx context.Context, request *ListWalletsRequest) (*ListWalletsResponse, error)
	CreateWallet(ctx context.Context, request *CreateWalletRequest) (*CreateWalletResponse, error)
	GetWallet(ctx context.Context, request *GetWalletRequest) (*GetWalletResponse, error)
	GetWalletDepositInstructions(ctx context.Context, request *GetWalletDepositInstructionsRequest) (*GetWalletDepositInstructionsResponse, error)
}

type clientImpl struct {
	httpClient  http.Client
	httpBaseUrl string

	headersFunc core.HttpHeaderFunc
	credentials *Credentials
	webSocket   *webSocket
}

func (c *clientImpl) HttpBaseUrl() string {
	return c.httpBaseUrl
}

func (c *clientImpl) SetBaseUrl(u string) Client {
	c.httpBaseUrl = u
	return c
}

func (c *clientImpl) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *clientImpl) Credentials() *Credentials {
	return c.credentials
}

func (c *clientImpl) SetHeadersFunc(hf core.HttpHeaderFunc) Client {
	c.headersFunc = hf
	return c
}

func NewClient(credentials *Credentials, httpClient http.Client) Client {
	return &clientImpl{
		httpBaseUrl: defaultV1ApiBaseUrl,
		credentials: credentials,
		httpClient:  httpClient,
		headersFunc: defaultHeadersFunc,
		webSocket:   newWebSocket(defaultWebSocketUrl),
	}
}

func AddPrimeHeaders(req *http.Request, path string, body []byte, client core.Client, t time.Time) {
	c := client.(*clientImpl)
	timestamp := strconv.FormatInt(t.Unix(), 10)
	signature := sign(req.Method, path, timestamp, c.Credentials().SigningKey, string(body))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CB-ACCESS-KEY", c.Credentials().AccessKey)
	req.Header.Add("X-CB-ACCESS-PASSPHRASE", c.Credentials().Passphrase)
	req.Header.Add("X-CB-ACCESS-SIGNATURE", signature)
	req.Header.Add("X-CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("User-Agent", fmt.Sprintf("prime-sdk-go/%s", sdkVersion))
}

func sign(method, path, timestamp, signingKey, body string) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	h.Write([]byte(fmt.Sprintf("%s%s%s%s", timestamp, method, path, body)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func signWebSocket(channel, portfolioId, svcAccountId, timestamp, key, signingKey string, productIds []string) string {
	h := hmac.New(sha256.New, []byte(signingKey))
	message := fmt.Sprintf("%s%s%s%s%s%s", channel, key, svcAccountId, timestamp, portfolioId, strings.Join(productIds, ""))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
