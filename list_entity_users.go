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
	"fmt"

	"github.com/coinbase-samples/core-go"
)

type ListEntityUsersRequest struct {
	EntityId   string            `json:"entity_id"`
	Pagination *PaginationParams `json:"pagination_params"`
}

type ListEntityUsersResponse struct {
	Users      []*User                 `json:"users"`
	Request    *ListEntityUsersRequest `json:"request"`
	Pagination *Pagination             `json:"pagination"`
}

func (c *clientImpl) ListEntityUsers(
	ctx context.Context,
	request *ListEntityUsersRequest,
) (*ListEntityUsersResponse, error) {

	path := fmt.Sprintf("/entities/%s/users", request.EntityId)

	queryParams := appendPaginationParams(core.EmptyQueryParams, request.Pagination)

	response := &ListEntityUsersResponse{Request: request}

	if err := core.HttpGet(ctx, c, path, queryParams, successStatusCodes, request, response, c.headersFunc); err != nil {
		return nil, err
	}

	return response, nil
}
