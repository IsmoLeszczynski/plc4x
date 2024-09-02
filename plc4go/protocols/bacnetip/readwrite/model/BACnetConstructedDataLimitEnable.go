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

// BACnetConstructedDataLimitEnable is the corresponding interface of BACnetConstructedDataLimitEnable
type BACnetConstructedDataLimitEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLimitEnable returns LimitEnable (property field)
	GetLimitEnable() BACnetLimitEnableTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLimitEnableTagged
}

// BACnetConstructedDataLimitEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLimitEnable.
// This is useful for switch cases.
type BACnetConstructedDataLimitEnableExactly interface {
	BACnetConstructedDataLimitEnable
	isBACnetConstructedDataLimitEnable() bool
}

// _BACnetConstructedDataLimitEnable is the data-structure of this message
type _BACnetConstructedDataLimitEnable struct {
	BACnetConstructedDataContract
	LimitEnable BACnetLimitEnableTagged
}

var _ BACnetConstructedDataLimitEnable = (*_BACnetConstructedDataLimitEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLimitEnable)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLimitEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIMIT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetLimitEnable() BACnetLimitEnableTagged {
	return m.LimitEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLimitEnable) GetActualValue() BACnetLimitEnableTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLimitEnableTagged(m.GetLimitEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLimitEnable factory function for _BACnetConstructedDataLimitEnable
func NewBACnetConstructedDataLimitEnable(limitEnable BACnetLimitEnableTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLimitEnable {
	_result := &_BACnetConstructedDataLimitEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LimitEnable:                   limitEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLimitEnable(structType any) BACnetConstructedDataLimitEnable {
	if casted, ok := structType.(BACnetConstructedDataLimitEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLimitEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLimitEnable) GetTypeName() string {
	return "BACnetConstructedDataLimitEnable"
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (limitEnable)
	lengthInBits += m.LimitEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLimitEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLimitEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLimitEnable BACnetConstructedDataLimitEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLimitEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLimitEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	limitEnable, err := ReadSimpleField[BACnetLimitEnableTagged](ctx, "limitEnable", ReadComplex[BACnetLimitEnableTagged](BACnetLimitEnableTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limitEnable' field"))
	}
	m.LimitEnable = limitEnable

	actualValue, err := ReadVirtualField[BACnetLimitEnableTagged](ctx, "actualValue", (*BACnetLimitEnableTagged)(nil), limitEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLimitEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLimitEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLimitEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLimitEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLimitEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLimitEnable")
		}

		if err := WriteSimpleField[BACnetLimitEnableTagged](ctx, "limitEnable", m.GetLimitEnable(), WriteComplex[BACnetLimitEnableTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limitEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLimitEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLimitEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLimitEnable) isBACnetConstructedDataLimitEnable() bool {
	return true
}

func (m *_BACnetConstructedDataLimitEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
