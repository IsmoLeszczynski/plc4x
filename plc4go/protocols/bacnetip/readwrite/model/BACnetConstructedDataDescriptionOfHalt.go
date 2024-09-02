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

// BACnetConstructedDataDescriptionOfHalt is the corresponding interface of BACnetConstructedDataDescriptionOfHalt
type BACnetConstructedDataDescriptionOfHalt interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDescriptionForHalt returns DescriptionForHalt (property field)
	GetDescriptionForHalt() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// _BACnetConstructedDataDescriptionOfHalt is the data-structure of this message
type _BACnetConstructedDataDescriptionOfHalt struct {
	BACnetConstructedDataContract
	DescriptionForHalt BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataDescriptionOfHalt = (*_BACnetConstructedDataDescriptionOfHalt)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDescriptionOfHalt)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DESCRIPTION_OF_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetDescriptionForHalt() BACnetApplicationTagCharacterString {
	return m.DescriptionForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDescriptionOfHalt) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetDescriptionForHalt())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDescriptionOfHalt factory function for _BACnetConstructedDataDescriptionOfHalt
func NewBACnetConstructedDataDescriptionOfHalt(descriptionForHalt BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDescriptionOfHalt {
	_result := &_BACnetConstructedDataDescriptionOfHalt{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		DescriptionForHalt:            descriptionForHalt,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDescriptionOfHalt(structType any) BACnetConstructedDataDescriptionOfHalt {
	if casted, ok := structType.(BACnetConstructedDataDescriptionOfHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDescriptionOfHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetTypeName() string {
	return "BACnetConstructedDataDescriptionOfHalt"
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (descriptionForHalt)
	lengthInBits += m.DescriptionForHalt.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDescriptionOfHalt) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDescriptionOfHalt) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDescriptionOfHalt BACnetConstructedDataDescriptionOfHalt, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDescriptionOfHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDescriptionOfHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	descriptionForHalt, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "descriptionForHalt", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'descriptionForHalt' field"))
	}
	m.DescriptionForHalt = descriptionForHalt

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), descriptionForHalt)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDescriptionOfHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDescriptionOfHalt")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDescriptionOfHalt) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDescriptionOfHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDescriptionOfHalt")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "descriptionForHalt", m.GetDescriptionForHalt(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'descriptionForHalt' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDescriptionOfHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDescriptionOfHalt")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDescriptionOfHalt) IsBACnetConstructedDataDescriptionOfHalt() {}

func (m *_BACnetConstructedDataDescriptionOfHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
