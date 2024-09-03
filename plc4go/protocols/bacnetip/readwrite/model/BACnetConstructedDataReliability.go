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

// BACnetConstructedDataReliability is the corresponding interface of BACnetConstructedDataReliability
type BACnetConstructedDataReliability interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReliability returns Reliability (property field)
	GetReliability() BACnetReliabilityTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetReliabilityTagged
	// IsBACnetConstructedDataReliability is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataReliability()
}

// _BACnetConstructedDataReliability is the data-structure of this message
type _BACnetConstructedDataReliability struct {
	BACnetConstructedDataContract
	Reliability BACnetReliabilityTagged
}

var _ BACnetConstructedDataReliability = (*_BACnetConstructedDataReliability)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataReliability)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReliability) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELIABILITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReliability) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetReliability() BACnetReliabilityTagged {
	return m.Reliability
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReliability) GetActualValue() BACnetReliabilityTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetReliabilityTagged(m.GetReliability())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReliability factory function for _BACnetConstructedDataReliability
func NewBACnetConstructedDataReliability(reliability BACnetReliabilityTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReliability {
	if reliability == nil {
		panic("reliability of type BACnetReliabilityTagged for BACnetConstructedDataReliability must not be nil")
	}
	_result := &_BACnetConstructedDataReliability{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Reliability:                   reliability,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReliability(structType any) BACnetConstructedDataReliability {
	if casted, ok := structType.(BACnetConstructedDataReliability); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliability); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReliability) GetTypeName() string {
	return "BACnetConstructedDataReliability"
}

func (m *_BACnetConstructedDataReliability) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (reliability)
	lengthInBits += m.Reliability.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReliability) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataReliability) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataReliability BACnetConstructedDataReliability, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliability"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReliability")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reliability, err := ReadSimpleField[BACnetReliabilityTagged](ctx, "reliability", ReadComplex[BACnetReliabilityTagged](BACnetReliabilityTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reliability' field"))
	}
	m.Reliability = reliability

	actualValue, err := ReadVirtualField[BACnetReliabilityTagged](ctx, "actualValue", (*BACnetReliabilityTagged)(nil), reliability)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliability"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReliability")
	}

	return m, nil
}

func (m *_BACnetConstructedDataReliability) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReliability) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliability"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReliability")
		}

		if err := WriteSimpleField[BACnetReliabilityTagged](ctx, "reliability", m.GetReliability(), WriteComplex[BACnetReliabilityTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reliability' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReliability"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReliability")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReliability) IsBACnetConstructedDataReliability() {}

func (m *_BACnetConstructedDataReliability) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
