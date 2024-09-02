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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// S7Address is the corresponding interface of S7Address
type S7Address interface {
	S7AddressContract
	S7AddressRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// S7AddressContract provides a set of functions which can be overwritten by a sub struct
type S7AddressContract interface {
	// IsS7Address() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7Address()
}

// S7AddressRequirements provides a set of functions which need to be implemented by a sub struct
type S7AddressRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetAddressType returns AddressType (discriminator field)
	GetAddressType() uint8
}

// _S7Address is the data-structure of this message
type _S7Address struct {
	_SubType S7Address
}

var _ S7AddressContract = (*_S7Address)(nil)

// NewS7Address factory function for _S7Address
func NewS7Address() *_S7Address {
	return &_S7Address{}
}

// Deprecated: use the interface for direct cast
func CastS7Address(structType any) S7Address {
	if casted, ok := structType.(S7Address); ok {
		return casted
	}
	if casted, ok := structType.(*S7Address); ok {
		return *casted
	}
	return nil
}

func (m *_S7Address) GetTypeName() string {
	return "S7Address"
}

func (m *_S7Address) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (addressType)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7Address) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func S7AddressParse[T S7Address](ctx context.Context, theBytes []byte) (T, error) {
	return S7AddressParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func S7AddressParseWithBufferProducer[T S7Address]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := S7AddressParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func S7AddressParseWithBuffer[T S7Address](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_S7Address{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_S7Address) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__s7Address S7Address, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Address")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	addressType, err := ReadDiscriminatorField[uint8](ctx, "addressType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'addressType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child S7Address
	switch {
	case addressType == 0x10: // S7AddressAny
		if _child, err = (&_S7AddressAny{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type S7AddressAny for type-switch of S7Address")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [addressType=%v]", addressType)
	}

	if closeErr := readBuffer.CloseContext("S7Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Address")
	}

	return _child, nil
}

func (pm *_S7Address) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child S7Address, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("S7Address"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Address")
	}

	if err := WriteDiscriminatorField(ctx, "addressType", m.GetAddressType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'addressType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Address"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Address")
	}
	return nil
}

func (m *_S7Address) IsS7Address() {}
