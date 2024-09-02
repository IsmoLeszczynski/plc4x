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

// BACnetConstructedDataStartTime is the corresponding interface of BACnetConstructedDataStartTime
type BACnetConstructedDataStartTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStartTime returns StartTime (property field)
	GetStartTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataStartTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStartTime.
// This is useful for switch cases.
type BACnetConstructedDataStartTimeExactly interface {
	BACnetConstructedDataStartTime
	isBACnetConstructedDataStartTime() bool
}

// _BACnetConstructedDataStartTime is the data-structure of this message
type _BACnetConstructedDataStartTime struct {
	BACnetConstructedDataContract
	StartTime BACnetDateTime
}

var _ BACnetConstructedDataStartTime = (*_BACnetConstructedDataStartTime)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataStartTime)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStartTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_START_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStartTime) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetStartTime() BACnetDateTime {
	return m.StartTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStartTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetStartTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStartTime factory function for _BACnetConstructedDataStartTime
func NewBACnetConstructedDataStartTime(startTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStartTime {
	_result := &_BACnetConstructedDataStartTime{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		StartTime:                     startTime,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStartTime(structType any) BACnetConstructedDataStartTime {
	if casted, ok := structType.(BACnetConstructedDataStartTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStartTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStartTime) GetTypeName() string {
	return "BACnetConstructedDataStartTime"
}

func (m *_BACnetConstructedDataStartTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStartTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataStartTime) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataStartTime BACnetConstructedDataStartTime, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStartTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStartTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startTime, err := ReadSimpleField[BACnetDateTime](ctx, "startTime", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), startTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStartTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStartTime")
	}

	return m, nil
}

func (m *_BACnetConstructedDataStartTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStartTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStartTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStartTime")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "startTime", m.GetStartTime(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStartTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStartTime")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStartTime) isBACnetConstructedDataStartTime() bool {
	return true
}

func (m *_BACnetConstructedDataStartTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
