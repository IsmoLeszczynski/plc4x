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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ErrorReportingSystemCategoryTypeReserved is the corresponding interface of ErrorReportingSystemCategoryTypeReserved
type ErrorReportingSystemCategoryTypeReserved interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetReservedValue returns ReservedValue (property field)
	GetReservedValue() uint8
}

// ErrorReportingSystemCategoryTypeReservedExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryTypeReserved.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeReservedExactly interface {
	ErrorReportingSystemCategoryTypeReserved
	isErrorReportingSystemCategoryTypeReserved() bool
}

// _ErrorReportingSystemCategoryTypeReserved is the data-structure of this message
type _ErrorReportingSystemCategoryTypeReserved struct {
	*_ErrorReportingSystemCategoryType
	ReservedValue uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) InitializeParent(parent ErrorReportingSystemCategoryType) {
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetParent() ErrorReportingSystemCategoryType {
	return m._ErrorReportingSystemCategoryType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeReserved) GetReservedValue() uint8 {
	return m.ReservedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategoryTypeReserved factory function for _ErrorReportingSystemCategoryTypeReserved
func NewErrorReportingSystemCategoryTypeReserved(reservedValue uint8) *_ErrorReportingSystemCategoryTypeReserved {
	_result := &_ErrorReportingSystemCategoryTypeReserved{
		ReservedValue:                     reservedValue,
		_ErrorReportingSystemCategoryType: NewErrorReportingSystemCategoryType(),
	}
	_result._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeReserved(structType any) ErrorReportingSystemCategoryTypeReserved {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeReserved); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeReserved); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeReserved"
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reservedValue)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeReserved) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ErrorReportingSystemCategoryTypeReservedParse(ctx context.Context, theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeReserved, error) {
	return ErrorReportingSystemCategoryTypeReservedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeReservedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeReserved, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reservedValue)
	_reservedValue, _reservedValueErr := readBuffer.ReadUint8("reservedValue", 4)
	if _reservedValueErr != nil {
		return nil, errors.Wrap(_reservedValueErr, "Error parsing 'reservedValue' field of ErrorReportingSystemCategoryTypeReserved")
	}
	reservedValue := _reservedValue

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeReserved")
	}

	// Create a partially initialized instance
	_child := &_ErrorReportingSystemCategoryTypeReserved{
		_ErrorReportingSystemCategoryType: &_ErrorReportingSystemCategoryType{},
		ReservedValue:                     reservedValue,
	}
	_child._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _child
	return _child, nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeReserved) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeReserved")
		}

		// Simple Field (reservedValue)
		reservedValue := uint8(m.GetReservedValue())
		_reservedValueErr := writeBuffer.WriteUint8("reservedValue", 4, uint8((reservedValue)))
		if _reservedValueErr != nil {
			return errors.Wrap(_reservedValueErr, "Error serializing 'reservedValue' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeReserved")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeReserved) isErrorReportingSystemCategoryTypeReserved() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryTypeReserved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
