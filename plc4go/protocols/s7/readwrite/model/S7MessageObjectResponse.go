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

// S7MessageObjectResponse is the corresponding interface of S7MessageObjectResponse
type S7MessageObjectResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7DataAlarmMessage
	// GetReturnCode returns ReturnCode (property field)
	GetReturnCode() DataTransportErrorCode
	// GetTransportSize returns TransportSize (property field)
	GetTransportSize() DataTransportSize
}

// _S7MessageObjectResponse is the data-structure of this message
type _S7MessageObjectResponse struct {
	S7DataAlarmMessageContract
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	// Reserved Fields
	reservedField0 *uint8
}

var _ S7MessageObjectResponse = (*_S7MessageObjectResponse)(nil)
var _ S7DataAlarmMessageRequirements = (*_S7MessageObjectResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageObjectResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageObjectResponse) GetParent() S7DataAlarmMessageContract {
	return m.S7DataAlarmMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7MessageObjectResponse) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

func (m *_S7MessageObjectResponse) GetTransportSize() DataTransportSize {
	return m.TransportSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7MessageObjectResponse factory function for _S7MessageObjectResponse
func NewS7MessageObjectResponse(returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7MessageObjectResponse {
	_result := &_S7MessageObjectResponse{
		S7DataAlarmMessageContract: NewS7DataAlarmMessage(),
		ReturnCode:                 returnCode,
		TransportSize:              transportSize,
	}
	_result.S7DataAlarmMessageContract.(*_S7DataAlarmMessage)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageObjectResponse(structType any) S7MessageObjectResponse {
	if casted, ok := structType.(S7MessageObjectResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageObjectResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageObjectResponse) GetTypeName() string {
	return "S7MessageObjectResponse"
}

func (m *_S7MessageObjectResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7DataAlarmMessageContract.(*_S7DataAlarmMessage).getLengthInBits(ctx))

	// Simple field (returnCode)
	lengthInBits += 8

	// Simple field (transportSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7MessageObjectResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7MessageObjectResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7DataAlarmMessage, cpuFunctionType uint8) (__s7MessageObjectResponse S7MessageObjectResponse, err error) {
	m.S7DataAlarmMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageObjectResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageObjectResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	returnCode, err := ReadEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", ReadEnum(DataTransportErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnCode' field"))
	}
	m.ReturnCode = returnCode

	transportSize, err := ReadEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", ReadEnum(DataTransportSizeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSize' field"))
	}
	m.TransportSize = transportSize

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	if closeErr := readBuffer.CloseContext("S7MessageObjectResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageObjectResponse")
	}

	return m, nil
}

func (m *_S7MessageObjectResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7MessageObjectResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageObjectResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageObjectResponse")
		}

		if err := WriteSimpleEnumField[DataTransportErrorCode](ctx, "returnCode", "DataTransportErrorCode", m.GetReturnCode(), WriteEnum[DataTransportErrorCode, uint8](DataTransportErrorCode.GetValue, DataTransportErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'returnCode' field")
		}

		if err := WriteSimpleEnumField[DataTransportSize](ctx, "transportSize", "DataTransportSize", m.GetTransportSize(), WriteEnum[DataTransportSize, uint8](DataTransportSize.GetValue, DataTransportSize.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSize' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if popErr := writeBuffer.PopContext("S7MessageObjectResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageObjectResponse")
		}
		return nil
	}
	return m.S7DataAlarmMessageContract.(*_S7DataAlarmMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7MessageObjectResponse) IsS7MessageObjectResponse() {}

func (m *_S7MessageObjectResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
