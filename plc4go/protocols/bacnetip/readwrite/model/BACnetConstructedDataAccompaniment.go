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

// BACnetConstructedDataAccompaniment is the corresponding interface of BACnetConstructedDataAccompaniment
type BACnetConstructedDataAccompaniment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccompaniment returns Accompaniment (property field)
	GetAccompaniment() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataAccompanimentExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccompaniment.
// This is useful for switch cases.
type BACnetConstructedDataAccompanimentExactly interface {
	BACnetConstructedDataAccompaniment
	isBACnetConstructedDataAccompaniment() bool
}

// _BACnetConstructedDataAccompaniment is the data-structure of this message
type _BACnetConstructedDataAccompaniment struct {
	BACnetConstructedDataContract
	Accompaniment BACnetDeviceObjectReference
}

var _ BACnetConstructedDataAccompaniment = (*_BACnetConstructedDataAccompaniment)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccompaniment)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccompaniment) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCOMPANIMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetAccompaniment() BACnetDeviceObjectReference {
	return m.Accompaniment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccompaniment) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetAccompaniment())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccompaniment factory function for _BACnetConstructedDataAccompaniment
func NewBACnetConstructedDataAccompaniment(accompaniment BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccompaniment {
	_result := &_BACnetConstructedDataAccompaniment{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Accompaniment:                 accompaniment,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccompaniment(structType any) BACnetConstructedDataAccompaniment {
	if casted, ok := structType.(BACnetConstructedDataAccompaniment); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccompaniment); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccompaniment) GetTypeName() string {
	return "BACnetConstructedDataAccompaniment"
}

func (m *_BACnetConstructedDataAccompaniment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accompaniment)
	lengthInBits += m.Accompaniment.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccompaniment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccompaniment) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccompaniment BACnetConstructedDataAccompaniment, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccompaniment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccompaniment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accompaniment, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "accompaniment", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accompaniment' field"))
	}
	m.Accompaniment = accompaniment

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), accompaniment)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccompaniment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccompaniment")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccompaniment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccompaniment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccompaniment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccompaniment")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "accompaniment", m.GetAccompaniment(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accompaniment' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccompaniment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccompaniment")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccompaniment) isBACnetConstructedDataAccompaniment() bool {
	return true
}

func (m *_BACnetConstructedDataAccompaniment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
