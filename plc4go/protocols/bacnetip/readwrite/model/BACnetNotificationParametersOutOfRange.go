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

// BACnetNotificationParametersOutOfRange is the corresponding interface of BACnetNotificationParametersOutOfRange
type BACnetNotificationParametersOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagReal
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagReal
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagReal
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersOutOfRange()
}

// _BACnetNotificationParametersOutOfRange is the data-structure of this message
type _BACnetNotificationParametersOutOfRange struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagReal
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagReal
	ExceededLimit   BACnetContextTagReal
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersOutOfRange = (*_BACnetNotificationParametersOutOfRange)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersOutOfRange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersOutOfRange) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersOutOfRange) GetExceedingValue() BACnetContextTagReal {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersOutOfRange) GetDeadband() BACnetContextTagReal {
	return m.Deadband
}

func (m *_BACnetNotificationParametersOutOfRange) GetExceededLimit() BACnetContextTagReal {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersOutOfRange factory function for _BACnetNotificationParametersOutOfRange
func NewBACnetNotificationParametersOutOfRange(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagReal, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagReal, exceededLimit BACnetContextTagReal, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersOutOfRange {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersOutOfRange must not be nil")
	}
	if exceedingValue == nil {
		panic("exceedingValue of type BACnetContextTagReal for BACnetNotificationParametersOutOfRange must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagReal for BACnetNotificationParametersOutOfRange must not be nil")
	}
	if exceededLimit == nil {
		panic("exceededLimit of type BACnetContextTagReal for BACnetNotificationParametersOutOfRange must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersOutOfRange must not be nil")
	}
	_result := &_BACnetNotificationParametersOutOfRange{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		ExceedingValue:                       exceedingValue,
		StatusFlags:                          statusFlags,
		Deadband:                             deadband,
		ExceededLimit:                        exceededLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersOutOfRange(structType any) BACnetNotificationParametersOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersOutOfRange"
}

func (m *_BACnetNotificationParametersOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersOutOfRange BACnetNotificationParametersOutOfRange, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	exceedingValue, err := ReadSimpleField[BACnetContextTagReal](ctx, "exceedingValue", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceedingValue' field"))
	}
	m.ExceedingValue = exceedingValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	deadband, err := ReadSimpleField[BACnetContextTagReal](ctx, "deadband", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	exceededLimit, err := ReadSimpleField[BACnetContextTagReal](ctx, "exceededLimit", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}
	m.ExceededLimit = exceededLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersOutOfRange")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "exceedingValue", m.GetExceedingValue(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceedingValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "exceededLimit", m.GetExceededLimit(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceededLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersOutOfRange")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersOutOfRange) IsBACnetNotificationParametersOutOfRange() {}

func (m *_BACnetNotificationParametersOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
