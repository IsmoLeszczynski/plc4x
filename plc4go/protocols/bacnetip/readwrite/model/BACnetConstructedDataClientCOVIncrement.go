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

// BACnetConstructedDataClientCOVIncrement is the corresponding interface of BACnetConstructedDataClientCOVIncrement
type BACnetConstructedDataClientCOVIncrement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetClientCOV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetClientCOV
}

// BACnetConstructedDataClientCOVIncrementExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataClientCOVIncrement.
// This is useful for switch cases.
type BACnetConstructedDataClientCOVIncrementExactly interface {
	BACnetConstructedDataClientCOVIncrement
	isBACnetConstructedDataClientCOVIncrement() bool
}

// _BACnetConstructedDataClientCOVIncrement is the data-structure of this message
type _BACnetConstructedDataClientCOVIncrement struct {
	BACnetConstructedDataContract
	CovIncrement BACnetClientCOV
}

var _ BACnetConstructedDataClientCOVIncrement = (*_BACnetConstructedDataClientCOVIncrement)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataClientCOVIncrement)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CLIENT_COV_INCREMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetCovIncrement() BACnetClientCOV {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetActualValue() BACnetClientCOV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetClientCOV(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataClientCOVIncrement factory function for _BACnetConstructedDataClientCOVIncrement
func NewBACnetConstructedDataClientCOVIncrement(covIncrement BACnetClientCOV, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataClientCOVIncrement {
	_result := &_BACnetConstructedDataClientCOVIncrement{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CovIncrement:                  covIncrement,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataClientCOVIncrement(structType any) BACnetConstructedDataClientCOVIncrement {
	if casted, ok := structType.(BACnetConstructedDataClientCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataClientCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataClientCOVIncrement"
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataClientCOVIncrement) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataClientCOVIncrement BACnetConstructedDataClientCOVIncrement, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataClientCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataClientCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	covIncrement, err := ReadSimpleField[BACnetClientCOV](ctx, "covIncrement", ReadComplex[BACnetClientCOV](BACnetClientCOVParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'covIncrement' field"))
	}
	m.CovIncrement = covIncrement

	actualValue, err := ReadVirtualField[BACnetClientCOV](ctx, "actualValue", (*BACnetClientCOV)(nil), covIncrement)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataClientCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataClientCOVIncrement")
	}

	return m, nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataClientCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataClientCOVIncrement")
		}

		if err := WriteSimpleField[BACnetClientCOV](ctx, "covIncrement", m.GetCovIncrement(), WriteComplex[BACnetClientCOV](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'covIncrement' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataClientCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataClientCOVIncrement")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataClientCOVIncrement) isBACnetConstructedDataClientCOVIncrement() bool {
	return true
}

func (m *_BACnetConstructedDataClientCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
