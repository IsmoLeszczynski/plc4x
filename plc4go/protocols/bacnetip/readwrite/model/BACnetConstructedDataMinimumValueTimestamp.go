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

// BACnetConstructedDataMinimumValueTimestamp is the corresponding interface of BACnetConstructedDataMinimumValueTimestamp
type BACnetConstructedDataMinimumValueTimestamp interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinimumValueTimestamp returns MinimumValueTimestamp (property field)
	GetMinimumValueTimestamp() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataMinimumValueTimestampExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMinimumValueTimestamp.
// This is useful for switch cases.
type BACnetConstructedDataMinimumValueTimestampExactly interface {
	BACnetConstructedDataMinimumValueTimestamp
	isBACnetConstructedDataMinimumValueTimestamp() bool
}

// _BACnetConstructedDataMinimumValueTimestamp is the data-structure of this message
type _BACnetConstructedDataMinimumValueTimestamp struct {
	BACnetConstructedDataContract
	MinimumValueTimestamp BACnetDateTime
}

var _ BACnetConstructedDataMinimumValueTimestamp = (*_BACnetConstructedDataMinimumValueTimestamp)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMinimumValueTimestamp)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_VALUE_TIMESTAMP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetMinimumValueTimestamp() BACnetDateTime {
	return m.MinimumValueTimestamp
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetMinimumValueTimestamp())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumValueTimestamp factory function for _BACnetConstructedDataMinimumValueTimestamp
func NewBACnetConstructedDataMinimumValueTimestamp(minimumValueTimestamp BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMinimumValueTimestamp {
	_result := &_BACnetConstructedDataMinimumValueTimestamp{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MinimumValueTimestamp:         minimumValueTimestamp,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMinimumValueTimestamp(structType any) BACnetConstructedDataMinimumValueTimestamp {
	if casted, ok := structType.(BACnetConstructedDataMinimumValueTimestamp); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumValueTimestamp); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetTypeName() string {
	return "BACnetConstructedDataMinimumValueTimestamp"
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (minimumValueTimestamp)
	lengthInBits += m.MinimumValueTimestamp.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMinimumValueTimestamp BACnetConstructedDataMinimumValueTimestamp, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumValueTimestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumValueTimestamp")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	minimumValueTimestamp, err := ReadSimpleField[BACnetDateTime](ctx, "minimumValueTimestamp", ReadComplex[BACnetDateTime](BACnetDateTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'minimumValueTimestamp' field"))
	}
	m.MinimumValueTimestamp = minimumValueTimestamp

	actualValue, err := ReadVirtualField[BACnetDateTime](ctx, "actualValue", (*BACnetDateTime)(nil), minimumValueTimestamp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumValueTimestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumValueTimestamp")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumValueTimestamp"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumValueTimestamp")
		}

		if err := WriteSimpleField[BACnetDateTime](ctx, "minimumValueTimestamp", m.GetMinimumValueTimestamp(), WriteComplex[BACnetDateTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'minimumValueTimestamp' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumValueTimestamp"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumValueTimestamp")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) isBACnetConstructedDataMinimumValueTimestamp() bool {
	return true
}

func (m *_BACnetConstructedDataMinimumValueTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
