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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusPDUWriteMultipleCoilsResponse is the corresponding interface of ModbusPDUWriteMultipleCoilsResponse
type ModbusPDUWriteMultipleCoilsResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
}

// ModbusPDUWriteMultipleCoilsResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteMultipleCoilsResponse.
// This is useful for switch cases.
type ModbusPDUWriteMultipleCoilsResponseExactly interface {
	ModbusPDUWriteMultipleCoilsResponse
	isModbusPDUWriteMultipleCoilsResponse() bool
}

// _ModbusPDUWriteMultipleCoilsResponse is the data-structure of this message
type _ModbusPDUWriteMultipleCoilsResponse struct {
	*_ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetFunctionFlag() uint8 {
	return 0x0F
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUWriteMultipleCoilsResponse factory function for _ModbusPDUWriteMultipleCoilsResponse
func NewModbusPDUWriteMultipleCoilsResponse(startingAddress uint16, quantity uint16) *_ModbusPDUWriteMultipleCoilsResponse {
	_result := &_ModbusPDUWriteMultipleCoilsResponse{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		_ModbusPDU:      NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleCoilsResponse(structType any) ModbusPDUWriteMultipleCoilsResponse {
	if casted, ok := structType.(ModbusPDUWriteMultipleCoilsResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleCoilsResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetTypeName() string {
	return "ModbusPDUWriteMultipleCoilsResponse"
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteMultipleCoilsResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUWriteMultipleCoilsResponse, error) {
	return ModbusPDUWriteMultipleCoilsResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUWriteMultipleCoilsResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteMultipleCoilsResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleCoilsResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleCoilsResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUWriteMultipleCoilsResponse")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUWriteMultipleCoilsResponse")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleCoilsResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleCoilsResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteMultipleCoilsResponse{
		_ModbusPDU:      &_ModbusPDU{},
		StartingAddress: startingAddress,
		Quantity:        quantity,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleCoilsResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleCoilsResponse")
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.GetStartingAddress())
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.GetQuantity())
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleCoilsResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleCoilsResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) isModbusPDUWriteMultipleCoilsResponse() bool {
	return true
}

func (m *_ModbusPDUWriteMultipleCoilsResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
