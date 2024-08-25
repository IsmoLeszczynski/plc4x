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

// BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
	isBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValue
	DoubleValue BACnetApplicationTagDouble
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(doubleValue BACnetApplicationTagDouble, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble{
		DoubleValue: doubleValue,
		_BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble, error) {
	return BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doubleValue")
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field of BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
	}
	doubleValue := _doubleValue.(BACnetApplicationTagDouble)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doubleValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble{
		_BACnetFaultParameterFaultOutOfRangeMinNormalValue: &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{
			TagNumber: tagNumber,
		},
		DoubleValue: doubleValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
		}

		// Simple Field (doubleValue)
		if pushErr := writeBuffer.PushContext("doubleValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doubleValue")
		}
		_doubleValueErr := writeBuffer.WriteSerializable(ctx, m.GetDoubleValue())
		if popErr := writeBuffer.PopContext("doubleValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doubleValue")
		}
		if _doubleValueErr != nil {
			return errors.Wrap(_doubleValueErr, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) isBACnetFaultParameterFaultOutOfRangeMinNormalValueDouble() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
