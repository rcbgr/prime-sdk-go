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
	"fmt"
	"time"

	"github.com/coinbase-samples/core-go"
)

type ListPortfolioFillsRequest struct {
	PortfolioId string            `json:"portfolio_id"` // required
	Start       time.Time         `json:"start_date"`   // required
	End         time.Time         `json:"end_date"`     // required
	Pagination  *PaginationParams `json:"pagination_params"`
}

type ListPortfolioFillsResponse struct {
	Fills      []*OrderFill               `json:"fills"`
	Pagination *Pagination                `json:"pagination"`
	Request    *ListPortfolioFillsRequest `json:"request"`
}

func (c *ClientImpl) ListPortfolioFills(
	ctx context.Context,
	request *ListPortfolioFillsRequest,
) (*ListPortfolioFillsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/fills", request.PortfolioId)

	queryParams := appendPaginationParams(core.EmptyQueryParams, request.Pagination)

	queryParams = core.AppendHttpQueryParam(queryParams, "start_date", TimeToStr(request.Start))

	if !request.End.IsZero() {
		queryParams = core.AppendHttpQueryParam(queryParams, "end_datae", TimeToStr(request.End))
	}

	response := &ListPortfolioFillsResponse{Request: request}

	if err := core.HttpGet(ctx, c, path, queryParams, successStatusCodes, request, response, c.headersFunc); err != nil {
		return nil, err
	}

	return response, nil
}
