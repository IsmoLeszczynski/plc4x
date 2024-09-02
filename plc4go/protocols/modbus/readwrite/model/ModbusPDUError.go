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

// ModbusPDUError is the corresponding interface of ModbusPDUError
type ModbusPDUError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetExceptionCode returns ExceptionCode (property field)
	GetExceptionCode() ModbusErrorCode
}

// ModbusPDUErrorExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUError.
// This is useful for switch cases.
type ModbusPDUErrorExactly interface {
	ModbusPDUError
	isModbusPDUError() bool
}

// _ModbusPDUError is the data-structure of this message
type _ModbusPDUError struct {
	ModbusPDUContract
	ExceptionCode ModbusErrorCode
}

var _ ModbusPDUError = (*_ModbusPDUError)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUError)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUError) GetErrorFlag() bool {
	return bool(true)
}

func (m *_ModbusPDUError) GetFunctionFlag() uint8 {
	return 0
}

func (m *_ModbusPDUError) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUError) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUError) GetExceptionCode() ModbusErrorCode {
	return m.ExceptionCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUError factory function for _ModbusPDUError
func NewModbusPDUError(exceptionCode ModbusErrorCode) *_ModbusPDUError {
	_result := &_ModbusPDUError{
		ModbusPDUContract: NewModbusPDU(),
		ExceptionCode:     exceptionCode,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUError(structType any) ModbusPDUError {
	if casted, ok := structType.(ModbusPDUError); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUError); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUError) GetTypeName() string {
	return "ModbusPDUError"
}

func (m *_ModbusPDUError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Simple field (exceptionCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ModbusPDUError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUError ModbusPDUError, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	exceptionCode, err := ReadEnumField[ModbusErrorCode](ctx, "exceptionCode", "ModbusErrorCode", ReadEnum(ModbusErrorCodeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceptionCode' field"))
	}
	m.ExceptionCode = exceptionCode

	if closeErr := readBuffer.CloseContext("ModbusPDUError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUError")
	}

	return m, nil
}

func (m *_ModbusPDUError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUError")
		}

		if err := WriteSimpleEnumField[ModbusErrorCode](ctx, "exceptionCode", "ModbusErrorCode", m.GetExceptionCode(), WriteEnum[ModbusErrorCode, uint8](ModbusErrorCode.GetValue, ModbusErrorCode.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'exceptionCode' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUError")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUError) isModbusPDUError() bool {
	return true
}

func (m *_ModbusPDUError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
