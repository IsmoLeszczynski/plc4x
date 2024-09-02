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

// BACnetSpecialEventPeriodCalendarReference is the corresponding interface of BACnetSpecialEventPeriodCalendarReference
type BACnetSpecialEventPeriodCalendarReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetSpecialEventPeriod
	// GetCalendarReference returns CalendarReference (property field)
	GetCalendarReference() BACnetContextTagObjectIdentifier
}

// BACnetSpecialEventPeriodCalendarReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetSpecialEventPeriodCalendarReference.
// This is useful for switch cases.
type BACnetSpecialEventPeriodCalendarReferenceExactly interface {
	BACnetSpecialEventPeriodCalendarReference
	isBACnetSpecialEventPeriodCalendarReference() bool
}

// _BACnetSpecialEventPeriodCalendarReference is the data-structure of this message
type _BACnetSpecialEventPeriodCalendarReference struct {
	BACnetSpecialEventPeriodContract
	CalendarReference BACnetContextTagObjectIdentifier
}

var _ BACnetSpecialEventPeriodCalendarReference = (*_BACnetSpecialEventPeriodCalendarReference)(nil)
var _ BACnetSpecialEventPeriodRequirements = (*_BACnetSpecialEventPeriodCalendarReference)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetSpecialEventPeriodCalendarReference) GetParent() BACnetSpecialEventPeriodContract {
	return m.BACnetSpecialEventPeriodContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventPeriodCalendarReference) GetCalendarReference() BACnetContextTagObjectIdentifier {
	return m.CalendarReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventPeriodCalendarReference factory function for _BACnetSpecialEventPeriodCalendarReference
func NewBACnetSpecialEventPeriodCalendarReference(calendarReference BACnetContextTagObjectIdentifier, peekedTagHeader BACnetTagHeader) *_BACnetSpecialEventPeriodCalendarReference {
	_result := &_BACnetSpecialEventPeriodCalendarReference{
		BACnetSpecialEventPeriodContract: NewBACnetSpecialEventPeriod(peekedTagHeader),
		CalendarReference:                calendarReference,
	}
	_result.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventPeriodCalendarReference(structType any) BACnetSpecialEventPeriodCalendarReference {
	if casted, ok := structType.(BACnetSpecialEventPeriodCalendarReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventPeriodCalendarReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetTypeName() string {
	return "BACnetSpecialEventPeriodCalendarReference"
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod).getLengthInBits(ctx))

	// Simple field (calendarReference)
	lengthInBits += m.CalendarReference.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventPeriodCalendarReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetSpecialEventPeriodCalendarReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetSpecialEventPeriod) (__bACnetSpecialEventPeriodCalendarReference BACnetSpecialEventPeriodCalendarReference, err error) {
	m.BACnetSpecialEventPeriodContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventPeriodCalendarReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventPeriodCalendarReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	calendarReference, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "calendarReference", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'calendarReference' field"))
	}
	m.CalendarReference = calendarReference

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventPeriodCalendarReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventPeriodCalendarReference")
	}

	return m, nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventPeriodCalendarReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetSpecialEventPeriodCalendarReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventPeriodCalendarReference")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "calendarReference", m.GetCalendarReference(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'calendarReference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetSpecialEventPeriodCalendarReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetSpecialEventPeriodCalendarReference")
		}
		return nil
	}
	return m.BACnetSpecialEventPeriodContract.(*_BACnetSpecialEventPeriod).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetSpecialEventPeriodCalendarReference) isBACnetSpecialEventPeriodCalendarReference() bool {
	return true
}

func (m *_BACnetSpecialEventPeriodCalendarReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
