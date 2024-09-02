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

// BACnetConstructedDataProgramChange is the corresponding interface of BACnetConstructedDataProgramChange
type BACnetConstructedDataProgramChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramChange returns ProgramChange (property field)
	GetProgramChange() BACnetProgramRequestTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProgramRequestTagged
}

// _BACnetConstructedDataProgramChange is the data-structure of this message
type _BACnetConstructedDataProgramChange struct {
	BACnetConstructedDataContract
	ProgramChange BACnetProgramRequestTagged
}

var _ BACnetConstructedDataProgramChange = (*_BACnetConstructedDataProgramChange)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataProgramChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProgramChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProgramChange) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetProgramChange() BACnetProgramRequestTagged {
	return m.ProgramChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetActualValue() BACnetProgramRequestTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProgramRequestTagged(m.GetProgramChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramChange factory function for _BACnetConstructedDataProgramChange
func NewBACnetConstructedDataProgramChange(programChange BACnetProgramRequestTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProgramChange {
	_result := &_BACnetConstructedDataProgramChange{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProgramChange:                 programChange,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProgramChange(structType any) BACnetConstructedDataProgramChange {
	if casted, ok := structType.(BACnetConstructedDataProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProgramChange) GetTypeName() string {
	return "BACnetConstructedDataProgramChange"
}

func (m *_BACnetConstructedDataProgramChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (programChange)
	lengthInBits += m.ProgramChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProgramChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataProgramChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataProgramChange BACnetConstructedDataProgramChange, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	programChange, err := ReadSimpleField[BACnetProgramRequestTagged](ctx, "programChange", ReadComplex[BACnetProgramRequestTagged](BACnetProgramRequestTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'programChange' field"))
	}
	m.ProgramChange = programChange

	actualValue, err := ReadVirtualField[BACnetProgramRequestTagged](ctx, "actualValue", (*BACnetProgramRequestTagged)(nil), programChange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramChange")
	}

	return m, nil
}

func (m *_BACnetConstructedDataProgramChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProgramChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramChange")
		}

		if err := WriteSimpleField[BACnetProgramRequestTagged](ctx, "programChange", m.GetProgramChange(), WriteComplex[BACnetProgramRequestTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'programChange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramChange")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProgramChange) IsBACnetConstructedDataProgramChange() {}

func (m *_BACnetConstructedDataProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
