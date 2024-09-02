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

// UnitAddress is the corresponding interface of UnitAddress
type UnitAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAddress returns Address (property field)
	GetAddress() byte
}

// _UnitAddress is the data-structure of this message
type _UnitAddress struct {
	Address byte
}

var _ UnitAddress = (*_UnitAddress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnitAddress) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnitAddress factory function for _UnitAddress
func NewUnitAddress(address byte) *_UnitAddress {
	return &_UnitAddress{Address: address}
}

// Deprecated: use the interface for direct cast
func CastUnitAddress(structType any) UnitAddress {
	if casted, ok := structType.(UnitAddress); ok {
		return casted
	}
	if casted, ok := structType.(*UnitAddress); ok {
		return *casted
	}
	return nil
}

func (m *_UnitAddress) GetTypeName() string {
	return "UnitAddress"
}

func (m *_UnitAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	return lengthInBits
}

func (m *_UnitAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func UnitAddressParse(ctx context.Context, theBytes []byte) (UnitAddress, error) {
	return UnitAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func UnitAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
		return UnitAddressParseWithBuffer(ctx, readBuffer)
	}
}

func UnitAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (UnitAddress, error) {
	v, err := (&_UnitAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_UnitAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__unitAddress UnitAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnitAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnitAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("UnitAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnitAddress")
	}

	return m, nil
}

func (m *_UnitAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnitAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("UnitAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for UnitAddress")
	}

	if err := WriteSimpleField[byte](ctx, "address", m.GetAddress(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}

	if popErr := writeBuffer.PopContext("UnitAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for UnitAddress")
	}
	return nil
}

func (m *_UnitAddress) IsUnitAddress() {}

func (m *_UnitAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
