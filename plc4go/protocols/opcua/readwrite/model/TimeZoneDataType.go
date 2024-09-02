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

// TimeZoneDataType is the corresponding interface of TimeZoneDataType
type TimeZoneDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetOffset returns Offset (property field)
	GetOffset() int16
	// GetDaylightSavingInOffset returns DaylightSavingInOffset (property field)
	GetDaylightSavingInOffset() bool
}

// _TimeZoneDataType is the data-structure of this message
type _TimeZoneDataType struct {
	ExtensionObjectDefinitionContract
	Offset                 int16
	DaylightSavingInOffset bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ TimeZoneDataType = (*_TimeZoneDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TimeZoneDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TimeZoneDataType) GetIdentifier() string {
	return "8914"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TimeZoneDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TimeZoneDataType) GetOffset() int16 {
	return m.Offset
}

func (m *_TimeZoneDataType) GetDaylightSavingInOffset() bool {
	return m.DaylightSavingInOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTimeZoneDataType factory function for _TimeZoneDataType
func NewTimeZoneDataType(offset int16, daylightSavingInOffset bool) *_TimeZoneDataType {
	_result := &_TimeZoneDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Offset:                            offset,
		DaylightSavingInOffset:            daylightSavingInOffset,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTimeZoneDataType(structType any) TimeZoneDataType {
	if casted, ok := structType.(TimeZoneDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TimeZoneDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TimeZoneDataType) GetTypeName() string {
	return "TimeZoneDataType"
}

func (m *_TimeZoneDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (offset)
	lengthInBits += 16

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (daylightSavingInOffset)
	lengthInBits += 1

	return lengthInBits
}

func (m *_TimeZoneDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TimeZoneDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__timeZoneDataType TimeZoneDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TimeZoneDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TimeZoneDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	offset, err := ReadSimpleField(ctx, "offset", ReadSignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'offset' field"))
	}
	m.Offset = offset

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	daylightSavingInOffset, err := ReadSimpleField(ctx, "daylightSavingInOffset", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'daylightSavingInOffset' field"))
	}
	m.DaylightSavingInOffset = daylightSavingInOffset

	if closeErr := readBuffer.CloseContext("TimeZoneDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TimeZoneDataType")
	}

	return m, nil
}

func (m *_TimeZoneDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TimeZoneDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TimeZoneDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TimeZoneDataType")
		}

		if err := WriteSimpleField[int16](ctx, "offset", m.GetOffset(), WriteSignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'offset' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "daylightSavingInOffset", m.GetDaylightSavingInOffset(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'daylightSavingInOffset' field")
		}

		if popErr := writeBuffer.PopContext("TimeZoneDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TimeZoneDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TimeZoneDataType) IsTimeZoneDataType() {}

func (m *_TimeZoneDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
