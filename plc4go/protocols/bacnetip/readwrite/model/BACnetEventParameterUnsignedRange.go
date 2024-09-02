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

// BACnetEventParameterUnsignedRange is the corresponding interface of BACnetEventParameterUnsignedRange
type BACnetEventParameterUnsignedRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagUnsignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// _BACnetEventParameterUnsignedRange is the data-structure of this message
type _BACnetEventParameterUnsignedRange struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagUnsignedInteger
	HighLimit  BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterUnsignedRange = (*_BACnetEventParameterUnsignedRange)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterUnsignedRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterUnsignedRange) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterUnsignedRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterUnsignedRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterUnsignedRange) GetLowLimit() BACnetContextTagUnsignedInteger {
	return m.LowLimit
}

func (m *_BACnetEventParameterUnsignedRange) GetHighLimit() BACnetContextTagUnsignedInteger {
	return m.HighLimit
}

func (m *_BACnetEventParameterUnsignedRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterUnsignedRange factory function for _BACnetEventParameterUnsignedRange
func NewBACnetEventParameterUnsignedRange(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetEventParameterUnsignedRange {
	_result := &_BACnetEventParameterUnsignedRange{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		LowLimit:                     lowLimit,
		HighLimit:                    highLimit,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterUnsignedRange(structType any) BACnetEventParameterUnsignedRange {
	if casted, ok := structType.(BACnetEventParameterUnsignedRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterUnsignedRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterUnsignedRange) GetTypeName() string {
	return "BACnetEventParameterUnsignedRange"
}

func (m *_BACnetEventParameterUnsignedRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterUnsignedRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterUnsignedRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterUnsignedRange BACnetEventParameterUnsignedRange, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterUnsignedRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterUnsignedRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(11))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	lowLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	highLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(11))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterUnsignedRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterUnsignedRange")
	}

	return m, nil
}

func (m *_BACnetEventParameterUnsignedRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterUnsignedRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterUnsignedRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterUnsignedRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterUnsignedRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterUnsignedRange")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterUnsignedRange) IsBACnetEventParameterUnsignedRange() {}

func (m *_BACnetEventParameterUnsignedRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
