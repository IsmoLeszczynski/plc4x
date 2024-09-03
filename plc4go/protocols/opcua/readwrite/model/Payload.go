/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Payload is the corresponding interface of Payload
type Payload interface {
	PayloadContract
	PayloadRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// IsPayload is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPayload()
}

// PayloadContract provides a set of functions which can be overwritten by a sub struct
type PayloadContract interface {
	// GetSequenceHeader returns SequenceHeader (property field)
	GetSequenceHeader() SequenceHeader
	// GetByteCount() returns a parser argument
	GetByteCount() uint32
	// IsPayload is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPayload()
}

// PayloadRequirements provides a set of functions which need to be implemented by a sub struct
type PayloadRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetExtensible returns Extensible (discriminator field)
	GetExtensible() bool
}

// _Payload is the data-structure of this message
type _Payload struct {
	_SubType       Payload
	SequenceHeader SequenceHeader

	// Arguments.
	ByteCount uint32
}

var _ PayloadContract = (*_Payload)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Payload) GetSequenceHeader() SequenceHeader {
	return m.SequenceHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPayload factory function for _Payload
func NewPayload(sequenceHeader SequenceHeader, byteCount uint32) *_Payload {
	if sequenceHeader == nil {
		panic("sequenceHeader of type SequenceHeader for Payload must not be nil")
	}
	return &_Payload{SequenceHeader: sequenceHeader, ByteCount: byteCount}
}

// Deprecated: use the interface for direct cast
func CastPayload(structType any) Payload {
	if casted, ok := structType.(Payload); ok {
		return casted
	}
	if casted, ok := structType.(*Payload); ok {
		return *casted
	}
	return nil
}

func (m *_Payload) GetTypeName() string {
	return "Payload"
}

func (m *_Payload) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (sequenceHeader)
	lengthInBits += m.SequenceHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_Payload) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func PayloadParse[T Payload](ctx context.Context, theBytes []byte, extensible bool, byteCount uint32) (T, error) {
	return PayloadParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), extensible, byteCount)
}

func PayloadParseWithBufferProducer[T Payload](extensible bool, byteCount uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := PayloadParseWithBuffer[T](ctx, readBuffer, extensible, byteCount)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func PayloadParseWithBuffer[T Payload](ctx context.Context, readBuffer utils.ReadBuffer, extensible bool, byteCount uint32) (T, error) {
	v, err := (&_Payload{ByteCount: byteCount}).parse(ctx, readBuffer, extensible, byteCount)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_Payload) parse(ctx context.Context, readBuffer utils.ReadBuffer, extensible bool, byteCount uint32) (__payload Payload, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Payload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sequenceHeader, err := ReadSimpleField[SequenceHeader](ctx, "sequenceHeader", ReadComplex[SequenceHeader](SequenceHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceHeader' field"))
	}
	m.SequenceHeader = sequenceHeader

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child Payload
	switch {
	case extensible == bool(true): // ExtensiblePayload
		if _child, err = (&_ExtensiblePayload{}).parse(ctx, readBuffer, m, extensible, byteCount); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type ExtensiblePayload for type-switch of Payload")
		}
	case extensible == bool(false): // BinaryPayload
		if _child, err = (&_BinaryPayload{}).parse(ctx, readBuffer, m, extensible, byteCount); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BinaryPayload for type-switch of Payload")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [extensible=%v]", extensible)
	}

	if closeErr := readBuffer.CloseContext("Payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Payload")
	}

	return _child, nil
}

func (pm *_Payload) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child Payload, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("Payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Payload")
	}

	if err := WriteSimpleField[SequenceHeader](ctx, "sequenceHeader", m.GetSequenceHeader(), WriteComplex[SequenceHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'sequenceHeader' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Payload")
	}
	return nil
}

////
// Arguments Getter

func (m *_Payload) GetByteCount() uint32 {
	return m.ByteCount
}

//
////

func (m *_Payload) IsPayload() {}
