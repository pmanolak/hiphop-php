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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/facebook/fbthrift/thrift/lib/go/thrift/types"
)

func TestWriteJSONProtocolBool(t *testing.T) {
	thetype := "boolean"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range boolValues {
		if e := p.WriteBool(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		expected := ""
		if value {
			expected = "1"
		} else {
			expected = "0"
		}
		if s != expected {
			t.Fatalf("Bad value for %s %v: %s expected", thetype, value, s)
		}
		v := -1
		if err := json.Unmarshal([]byte(s), &v); err != nil || (v != 0) != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolBool(t *testing.T) {
	thetype := "boolean"
	for _, value := range boolValues {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		if value {
			trans.Write([]byte{'1'}) // not JSON_TRUE
		} else {
			trans.Write([]byte{'0'}) // not JSON_FALSE
		}
		s := trans.String()
		v, e := p.ReadBool()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		vv := -1
		if err := json.Unmarshal([]byte(s), &vv); err != nil || (vv != 0) != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, vv)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolByte(t *testing.T) {
	thetype := "byte"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range byteValues {
		if e := p.WriteByte(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if s != fmt.Sprint(value) {
			t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
		}
		v := byte(0)
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolByte(t *testing.T) {
	thetype := "byte"
	for _, value := range byteValues {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		trans.WriteString(strconv.Itoa(int(value)))
		s := trans.String()
		v, e := p.ReadByte()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolI16(t *testing.T) {
	thetype := "int16"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range int16Values {
		if e := p.WriteI16(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if s != fmt.Sprint(value) {
			t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
		}
		v := int16(0)
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolI16(t *testing.T) {
	thetype := "int16"
	for _, value := range int16Values {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		trans.WriteString(strconv.Itoa(int(value)))
		s := trans.String()
		v, e := p.ReadI16()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolI32(t *testing.T) {
	thetype := "int32"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range int32Values {
		if e := p.WriteI32(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if s != fmt.Sprint(value) {
			t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
		}
		v := int32(0)
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolI32(t *testing.T) {
	thetype := "int32"
	for _, value := range int32Values {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		trans.WriteString(strconv.Itoa(int(value)))
		s := trans.String()
		v, e := p.ReadI32()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolI64(t *testing.T) {
	thetype := "int64"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range int64Values {
		if e := p.WriteI64(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if s != fmt.Sprint(value) {
			t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
		}
		v := int64(0)
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolI64(t *testing.T) {
	thetype := "int64"
	for _, value := range int64Values {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		trans.WriteString(strconv.FormatInt(value, 10))
		s := trans.String()
		v, e := p.ReadI64()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolDouble(t *testing.T) {
	thetype := "double"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range doubleValues {
		if e := p.WriteDouble(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if math.IsInf(value, 1) {
			if s != jsonQuote(JSON_INFINITY) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_INFINITY))
			}
		} else if math.IsInf(value, -1) {
			if s != jsonQuote(JSON_NEGATIVE_INFINITY) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_NEGATIVE_INFINITY))
			}
		} else if math.IsNaN(value) {
			if s != jsonQuote(JSON_NAN) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_NAN))
			}
		} else {
			if s != fmt.Sprint(value) {
				t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
			}
			v := float64(0)
			if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
			}
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolDouble(t *testing.T) {
	thetype := "double"
	for _, value := range doubleValues {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		n := NewNumericFromDouble(value)
		trans.WriteString(n.String())
		s := trans.String()
		v, e := p.ReadDouble()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if math.IsInf(value, 1) {
			if !math.IsInf(v, 1) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else if math.IsInf(value, -1) {
			if !math.IsInf(v, -1) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else if math.IsNaN(value) {
			if !math.IsNaN(v) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else {
			if v != value {
				t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
			}
			if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
			}
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolFloat(t *testing.T) {
	thetype := "float"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range floatValues {
		if e := p.WriteFloat(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if math.IsInf(float64(value), 1) {
			if s != jsonQuote(JSON_INFINITY) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_INFINITY))
			}
		} else if math.IsInf(float64(value), -1) {
			if s != jsonQuote(JSON_NEGATIVE_INFINITY) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_NEGATIVE_INFINITY))
			}
		} else if math.IsNaN(float64(value)) {
			if s != jsonQuote(JSON_NAN) {
				t.Fatalf("Bad value for %s %v, wrote: %v, expected: %v", thetype, value, s, jsonQuote(JSON_NAN))
			}
		} else {
			if s != fmt.Sprint(value) {
				t.Fatalf("Bad value for %s %v: %s", thetype, value, s)
			}
			v := float32(0)
			if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
			}
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolFloat(t *testing.T) {
	thetype := "float"
	for _, value := range floatValues {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		n := NewNumericFromFloat(value)
		trans.WriteString(n.String())
		s := trans.String()
		v, e := p.ReadFloat()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if math.IsInf(float64(value), 1) {
			if !math.IsInf(float64(v), 1) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else if math.IsInf(float64(value), -1) {
			if !math.IsInf(float64(v), -1) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else if math.IsNaN(float64(value)) {
			if !math.IsNaN(float64(v)) {
				t.Fatalf("Bad value for %s %v, wrote: %v, received: %v", thetype, value, s, v)
			}
		} else {
			if v != value {
				t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
			}
			if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
			}
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolString(t *testing.T) {
	thetype := "string"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	for _, value := range stringValues {
		if e := p.WriteString(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
		if e := p.Flush(); e != nil {
			t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
		}
		s := trans.String()
		if s[0] != '"' || s[len(s)-1] != '"' {
			t.Fatalf("Bad value for %s '%v', wrote '%v', expected: %v", thetype, value, s, fmt.Sprint("\"", value, "\""))
		}
		v := new(string)
		if err := json.Unmarshal([]byte(s), v); err != nil || *v != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, *v)
		}
		trans.Reset()
	}
}

func TestReadJSONProtocolString(t *testing.T) {
	thetype := "string"
	for _, value := range stringValues {
		trans := new(bytes.Buffer)
		p := NewCompactJSONFormat(trans)
		trans.WriteString(jsonQuote(value))
		s := trans.String()
		v, e := p.ReadString()
		if e != nil {
			t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
		}
		if v != value {
			t.Fatalf("Bad value for %s value %v, wrote: %v, received: %v", thetype, value, s, v)
		}
		v1 := new(string)
		if err := json.Unmarshal([]byte(s), v1); err != nil || *v1 != value {
			t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, *v1)
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolBinary(t *testing.T) {
	thetype := "binary"
	value := protocolBdata
	b64value := make([]byte, base64.StdEncoding.EncodedLen(len(protocolBdata)))
	base64.StdEncoding.Encode(b64value, value)
	b64String := string(b64value)
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	if e := p.WriteBinary(value); e != nil {
		t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
	}
	if e := p.Flush(); e != nil {
		t.Fatalf("Unable to write %s value %v due to error flushing: %s", thetype, value, e.Error())
	}
	s := trans.String()
	expectedString := fmt.Sprint("\"", b64String, "\"")
	if s != expectedString {
		t.Fatalf("Bad value for %s %v\n  wrote:  \"%v\"\nexpected: \"%v\"", thetype, value, s, expectedString)
	}
	v1, err := p.ReadBinary()
	if err != nil {
		t.Fatalf("Unable to read binary: %s", err.Error())
	}
	if len(v1) != len(value) {
		t.Fatalf("Invalid value for binary\nexpected: \"%v\"\n   read: \"%v\"", value, v1)
	}
	for k, v := range value {
		if v1[k] != v {
			t.Fatalf("Invalid value for binary at %v\nexpected: \"%v\"\n   read: \"%v\"", k, v, v1[k])
		}
	}
}

func TestReadJSONProtocolBinary(t *testing.T) {
	thetype := "binary"
	value := protocolBdata
	b64value := make([]byte, base64.StdEncoding.EncodedLen(len(protocolBdata)))
	base64.StdEncoding.Encode(b64value, value)
	b64String := string(b64value)
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	trans.WriteString(jsonQuote(b64String))
	s := trans.String()
	v, e := p.ReadBinary()
	if e != nil {
		t.Fatalf("Unable to read %s value %v due to error: %s", thetype, value, e.Error())
	}
	if len(v) != len(value) {
		t.Fatalf("Bad value for %s value length %v, wrote: %v, received length: %v", thetype, len(value), s, len(v))
	}
	for i := 0; i < len(v); i++ {
		if v[i] != value[i] {
			t.Fatalf("Bad value for %s at index %d value %v, wrote: %v, received: %v", thetype, i, value[i], s, v[i])
		}
	}
	v1 := new(string)
	if err := json.Unmarshal([]byte(s), v1); err != nil || *v1 != b64String {
		t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, *v1)
	}
	trans.Reset()
}

func TestWriteJSONProtocolList(t *testing.T) {
	thetype := "list"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	p.WriteListBegin(types.DOUBLE, len(doubleValues))
	for _, value := range doubleValues {
		if e := p.WriteDouble(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
	}
	p.WriteListEnd()
	if e := p.Flush(); e != nil {
		t.Fatalf("Unable to write %s due to error flushing: %s", thetype, e.Error())
	}
	str := trans.String()
	str1 := new([]any)
	err := json.Unmarshal([]byte(str), str1)
	if err != nil {
		t.Fatalf("Unable to decode %s, wrote: %s", thetype, str)
	}
	l := *str1
	if len(l) < 2 {
		t.Fatalf("List must be at least of length two to include metadata")
	}
	if l[0] != "dbl" {
		t.Fatal("Invalid type for list, expected: ", types.STRING, ", but was: ", l[0])
	}
	if int(l[1].(float64)) != len(doubleValues) {
		t.Fatal("Invalid length for list, expected: ", len(doubleValues), ", but was: ", l[1])
	}
	for k, value := range doubleValues {
		s := l[k+2]
		if math.IsInf(value, 1) {
			if s.(string) != JSON_INFINITY {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_INFINITY), str)
			}
		} else if math.IsInf(value, 0) {
			if s.(string) != JSON_NEGATIVE_INFINITY {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_NEGATIVE_INFINITY), str)
			}
		} else if math.IsNaN(value) {
			if s.(string) != JSON_NAN {
				t.Fatalf("Bad value for %s at index %v  %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_NAN), str)
			}
		} else {
			if s.(float64) != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s'", thetype, value, s)
			}
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolSet(t *testing.T) {
	thetype := "set"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	p.WriteSetBegin(types.DOUBLE, len(doubleValues))
	for _, value := range doubleValues {
		if e := p.WriteDouble(value); e != nil {
			t.Fatalf("Unable to write %s value %v due to error: %s", thetype, value, e.Error())
		}
	}
	p.WriteSetEnd()
	if e := p.Flush(); e != nil {
		t.Fatalf("Unable to write %s due to error flushing: %s", thetype, e.Error())
	}
	str := trans.String()
	str1 := new([]any)
	err := json.Unmarshal([]byte(str), str1)
	if err != nil {
		t.Fatalf("Unable to decode %s, wrote: %s", thetype, str)
	}
	l := *str1
	if len(l) < 2 {
		t.Fatalf("Set must be at least of length two to include metadata")
	}
	if l[0] != "dbl" {
		t.Fatal("Invalid type for set, expected: ", types.DOUBLE, ", but was: ", l[0])
	}
	if int(l[1].(float64)) != len(doubleValues) {
		t.Fatal("Invalid length for set, expected: ", len(doubleValues), ", but was: ", l[1])
	}
	for k, value := range doubleValues {
		s := l[k+2]
		if math.IsInf(value, 1) {
			if s.(string) != JSON_INFINITY {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_INFINITY), str)
			}
		} else if math.IsInf(value, 0) {
			if s.(string) != JSON_NEGATIVE_INFINITY {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_NEGATIVE_INFINITY), str)
			}
		} else if math.IsNaN(value) {
			if s.(string) != JSON_NAN {
				t.Fatalf("Bad value for %s at index %v  %v, wrote: %q, expected: %q, originally wrote: %q", thetype, k, value, s, jsonQuote(JSON_NAN), str)
			}
		} else {
			if s.(float64) != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s'", thetype, value, s)
			}
		}
		trans.Reset()
	}
}

func TestWriteJSONProtocolMap(t *testing.T) {
	thetype := "map"
	trans := new(bytes.Buffer)
	p := NewCompactJSONFormat(trans)
	p.WriteMapBegin(types.I32, types.DOUBLE, len(doubleValues))
	for k, value := range doubleValues {
		if e := p.WriteI32(int32(k)); e != nil {
			t.Fatalf("Unable to write %s key int32 value %v due to error: %s", thetype, k, e.Error())
		}
		if e := p.WriteDouble(value); e != nil {
			t.Fatalf("Unable to write %s value float64 value %v due to error: %s", thetype, value, e.Error())
		}
	}
	p.WriteMapEnd()
	if e := p.Flush(); e != nil {
		t.Fatalf("Unable to write %s due to error flushing: %s", thetype, e.Error())
	}
	str := trans.String()
	if str[0] != '[' || str[len(str)-1] != ']' {
		t.Fatalf("Bad value for %s, wrote: %q, in go: '%f'", thetype, str, doubleValues)
	}
	expectedKeyType, expectedValueType, expectedSize, err := p.ReadMapBegin()
	if err != nil {
		t.Fatalf("Error while reading map begin: %s", err.Error())
	}
	if expectedKeyType != types.I32 {
		t.Fatal("Expected map key type ", types.I32, ", but was ", expectedKeyType)
	}
	if expectedValueType != types.DOUBLE {
		t.Fatal("Expected map value type ", types.DOUBLE, ", but was ", expectedValueType)
	}
	if expectedSize != len(doubleValues) {
		t.Fatal("Expected map size of ", len(doubleValues), ", but was ", expectedSize)
	}
	for k, value := range doubleValues {
		ik, err := p.ReadI32()
		if err != nil {
			t.Fatalf("Bad key for %s index %v, wrote: %v, expected: %v, error: %s", thetype, k, ik, strconv.Itoa(k), err.Error())
		}
		if int(ik) != k {
			t.Fatalf("Bad key for %s index %v, wrote: %v, expected: %v", thetype, k, ik, k)
		}
		dv, err := p.ReadDouble()
		if err != nil {
			t.Fatalf("Bad value for %s index %v, wrote: %v, expected: %v, error: %s", thetype, k, dv, value, err.Error())
		}
		s := strconv.FormatFloat(dv, 'g', 10, 64)
		if math.IsInf(value, 1) {
			if !math.IsInf(dv, 1) {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %v, expected: %v", thetype, k, value, s, jsonQuote(JSON_INFINITY))
			}
		} else if math.IsInf(value, 0) {
			if !math.IsInf(dv, 0) {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %v, expected: %v", thetype, k, value, s, jsonQuote(JSON_NEGATIVE_INFINITY))
			}
		} else if math.IsNaN(value) {
			if !math.IsNaN(dv) {
				t.Fatalf("Bad value for %s at index %v  %v, wrote: %v, expected: %v", thetype, k, value, s, jsonQuote(JSON_NAN))
			}
		} else {
			expected := strconv.FormatFloat(value, 'g', 10, 64)
			if s != expected {
				t.Fatalf("Bad value for %s at index %v %v, wrote: %v, expected %v", thetype, k, value, s, expected)
			}
			v := float64(0)
			if err := json.Unmarshal([]byte(s), &v); err != nil || v != value {
				t.Fatalf("Bad json-decoded value for %s %v, wrote: '%s', expected: '%v'", thetype, value, s, v)
			}
		}
	}
	err = p.ReadMapEnd()
	if err != nil {
		t.Fatalf("Error while reading map end: %s", err.Error())
	}
}

func TestCompactJSONProtocolUnmatchedBeginEnd(t *testing.T) {
	UnmatchedBeginEndProtocolTest(t, NewCompactJSONFormat)
}
