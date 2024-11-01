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

package thrift

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
	"github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

type headerServerTestProcessor struct {
	requests chan<- *MyTestStruct
}

func (t *headerServerTestProcessor) ProcessorFunctionMap() map[string]types.ProcessorFunction {
	return map[string]types.ProcessorFunction{"test": &headerServerTestProcessorFunction{&testProcessorFunction{}, t.requests}}
}

func (t *headerServerTestProcessor) GetThriftMetadata() *metadata.ThriftMetadata {
	return nil
}

type headerServerTestProcessorFunction struct {
	types.ProcessorFunction
	requests chan<- *MyTestStruct
}

func (p *headerServerTestProcessorFunction) RunContext(ctx context.Context, reqStruct types.Struct) (types.WritableStruct, types.ApplicationException) {
	if p.requests != nil {
		p.requests <- reqStruct.(*MyTestStruct)
	}
	return reqStruct, nil
}

// Test that header server stops serving if listener is closed.
func TestHeaderServerCloseListener(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errChan := make(chan error)
	defer close(errChan)
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	server := NewSimpleServer(&headerServerTestProcessor{}, listener, TransportIDHeader)
	go func() {
		errChan <- server.ServeContext(ctx)
	}()
	addr := listener.Addr()
	conn, err := net.Dial(addr.Network(), addr.String())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	proto, err := NewHeaderProtocol(conn)
	if err != nil {
		t.Fatalf("could not create client protocol: %s", err)
	}
	client := NewSerialChannel(proto)
	req := &MyTestStruct{
		St: "hello",
	}
	resp := &MyTestStruct{}
	if err := client.Call(context.Background(), "test", req, resp); err != nil {
		t.Fatalf("could not complete call: %v", err)
	}
	if resp.St != "hello" {
		t.Fatalf("expected response to be a hello, got %s", resp.St)
	}
	listener.Close()
	select {
	case <-errChan:
		break
	case <-time.After(3 * time.Second):
		t.Fatalf("listener did not close")
	}
	cancel()
}
