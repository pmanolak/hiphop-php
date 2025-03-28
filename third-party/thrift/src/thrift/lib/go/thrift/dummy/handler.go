/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dummy

import (
	"context"

	"github.com/facebook/fbthrift/thrift/test/go/if/dummy"
)

type DummyHandler struct {
	// Testing helper to verify that oneway RPCs get delivered
	OnewayRPCRequests chan<- string
}

// Compile time interface enforcer
var _ dummy.Dummy = (*DummyHandler)(nil)

func (h *DummyHandler) Echo(ctx context.Context, value string) (string, error) {
	return value, nil
}

func (h *DummyHandler) OnewayRPC(ctx context.Context, value string) error {
	if h.OnewayRPCRequests != nil {
		h.OnewayRPCRequests <- value
	}
	return nil
}
