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

// BridgeAddress is the corresponding interface of BridgeAddress
type BridgeAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAddress returns Address (property field)
	GetAddress() byte
}

// _BridgeAddress is the data-structure of this message
type _BridgeAddress struct {
	Address byte
}

var _ BridgeAddress = (*_BridgeAddress)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BridgeAddress) GetAddress() byte {
	return m.Address
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBridgeAddress factory function for _BridgeAddress
func NewBridgeAddress(address byte) *_BridgeAddress {
	return &_BridgeAddress{Address: address}
}

// Deprecated: use the interface for direct cast
func CastBridgeAddress(structType any) BridgeAddress {
	if casted, ok := structType.(BridgeAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BridgeAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BridgeAddress) GetTypeName() string {
	return "BridgeAddress"
}

func (m *_BridgeAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (address)
	lengthInBits += 8

	return lengthInBits
}

func (m *_BridgeAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BridgeAddressParse(ctx context.Context, theBytes []byte) (BridgeAddress, error) {
	return BridgeAddressParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BridgeAddressParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BridgeAddress, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BridgeAddress, error) {
		return BridgeAddressParseWithBuffer(ctx, readBuffer)
	}
}

func BridgeAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BridgeAddress, error) {
	v, err := (&_BridgeAddress{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BridgeAddress) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bridgeAddress BridgeAddress, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BridgeAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BridgeAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	if closeErr := readBuffer.CloseContext("BridgeAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BridgeAddress")
	}

	return m, nil
}

func (m *_BridgeAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BridgeAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BridgeAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BridgeAddress")
	}

	if err := WriteSimpleField[byte](ctx, "address", m.GetAddress(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'address' field")
	}

	if popErr := writeBuffer.PopContext("BridgeAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BridgeAddress")
	}
	return nil
}

func (m *_BridgeAddress) IsBridgeAddress() {}

func (m *_BridgeAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
