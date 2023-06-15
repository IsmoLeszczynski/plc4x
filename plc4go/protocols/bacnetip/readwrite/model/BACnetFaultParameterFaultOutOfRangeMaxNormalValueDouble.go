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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMaxNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(doubleValue BACnetApplicationTagDouble, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		DoubleValue: doubleValue,
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, error) {
	return BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doubleValue)
	if pullErr := readBuffer.PullContext("doubleValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doubleValue")
	}
	_doubleValue, _doubleValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _doubleValueErr != nil {
		return nil, errors.Wrap(_doubleValueErr, "Error parsing 'doubleValue' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}
	doubleValue := _doubleValue.(BACnetApplicationTagDouble)
	if closeErr := readBuffer.CloseContext("doubleValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doubleValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &_BACnetFaultParameterFaultOutOfRangeMaxNormalValue{
			TagNumber: tagNumber,
		},
		DoubleValue: doubleValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) isBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
