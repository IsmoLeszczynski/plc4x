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

// BACnetConstructedDataDateValueRelinquishDefault is the corresponding interface of BACnetConstructedDataDateValueRelinquishDefault
type BACnetConstructedDataDateValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
}

// BACnetConstructedDataDateValueRelinquishDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDateValueRelinquishDefault.
// This is useful for switch cases.
type BACnetConstructedDataDateValueRelinquishDefaultExactly interface {
	BACnetConstructedDataDateValueRelinquishDefault
	isBACnetConstructedDataDateValueRelinquishDefault() bool
}

// _BACnetConstructedDataDateValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataDateValueRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetApplicationTagDate
}

var _ BACnetConstructedDataDateValueRelinquishDefault = (*_BACnetConstructedDataDateValueRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataDateValueRelinquishDefault)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_DATE_VALUE
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagDate {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetActualValue() BACnetApplicationTagDate {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDate(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDateValueRelinquishDefault factory function for _BACnetConstructedDataDateValueRelinquishDefault
func NewBACnetConstructedDataDateValueRelinquishDefault(relinquishDefault BACnetApplicationTagDate, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDateValueRelinquishDefault {
	_result := &_BACnetConstructedDataDateValueRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateValueRelinquishDefault(structType any) BACnetConstructedDataDateValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataDateValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataDateValueRelinquishDefault"
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataDateValueRelinquishDefault BACnetConstructedDataDateValueRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "relinquishDefault", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetApplicationTagDate](ctx, "actualValue", (*BACnetApplicationTagDate)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateValueRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateValueRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateValueRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) isBACnetConstructedDataDateValueRelinquishDefault() bool {
	return true
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
