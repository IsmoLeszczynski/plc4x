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

// BACnetConstructedDataAccessEventAuthenticationFactor is the corresponding interface of BACnetConstructedDataAccessEventAuthenticationFactor
type BACnetConstructedDataAccessEventAuthenticationFactor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEventAuthenticationFactor returns AccessEventAuthenticationFactor (property field)
	GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAuthenticationFactor
}

// BACnetConstructedDataAccessEventAuthenticationFactorExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessEventAuthenticationFactor.
// This is useful for switch cases.
type BACnetConstructedDataAccessEventAuthenticationFactorExactly interface {
	BACnetConstructedDataAccessEventAuthenticationFactor
	isBACnetConstructedDataAccessEventAuthenticationFactor() bool
}

// _BACnetConstructedDataAccessEventAuthenticationFactor is the data-structure of this message
type _BACnetConstructedDataAccessEventAuthenticationFactor struct {
	BACnetConstructedDataContract
	AccessEventAuthenticationFactor BACnetAuthenticationFactor
}

var _ BACnetConstructedDataAccessEventAuthenticationFactor = (*_BACnetConstructedDataAccessEventAuthenticationFactor)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessEventAuthenticationFactor)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT_AUTHENTICATION_FACTOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetAccessEventAuthenticationFactor() BACnetAuthenticationFactor {
	return m.AccessEventAuthenticationFactor
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetActualValue() BACnetAuthenticationFactor {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAuthenticationFactor(m.GetAccessEventAuthenticationFactor())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEventAuthenticationFactor factory function for _BACnetConstructedDataAccessEventAuthenticationFactor
func NewBACnetConstructedDataAccessEventAuthenticationFactor(accessEventAuthenticationFactor BACnetAuthenticationFactor, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEventAuthenticationFactor {
	_result := &_BACnetConstructedDataAccessEventAuthenticationFactor{
		BACnetConstructedDataContract:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessEventAuthenticationFactor: accessEventAuthenticationFactor,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEventAuthenticationFactor(structType any) BACnetConstructedDataAccessEventAuthenticationFactor {
	if casted, ok := structType.(BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEventAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetTypeName() string {
	return "BACnetConstructedDataAccessEventAuthenticationFactor"
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accessEventAuthenticationFactor)
	lengthInBits += m.AccessEventAuthenticationFactor.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessEventAuthenticationFactor BACnetConstructedDataAccessEventAuthenticationFactor, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEventAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEventAuthenticationFactor, err := ReadSimpleField[BACnetAuthenticationFactor](ctx, "accessEventAuthenticationFactor", ReadComplex[BACnetAuthenticationFactor](BACnetAuthenticationFactorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEventAuthenticationFactor' field"))
	}
	m.AccessEventAuthenticationFactor = accessEventAuthenticationFactor

	actualValue, err := ReadVirtualField[BACnetAuthenticationFactor](ctx, "actualValue", (*BACnetAuthenticationFactor)(nil), accessEventAuthenticationFactor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEventAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEventAuthenticationFactor")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEventAuthenticationFactor"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEventAuthenticationFactor")
		}

		if err := WriteSimpleField[BACnetAuthenticationFactor](ctx, "accessEventAuthenticationFactor", m.GetAccessEventAuthenticationFactor(), WriteComplex[BACnetAuthenticationFactor](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEventAuthenticationFactor' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEventAuthenticationFactor"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEventAuthenticationFactor")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) isBACnetConstructedDataAccessEventAuthenticationFactor() bool {
	return true
}

func (m *_BACnetConstructedDataAccessEventAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
