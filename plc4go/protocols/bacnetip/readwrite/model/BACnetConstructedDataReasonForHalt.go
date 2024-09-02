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

// BACnetConstructedDataReasonForHalt is the corresponding interface of BACnetConstructedDataReasonForHalt
type BACnetConstructedDataReasonForHalt interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramError returns ProgramError (property field)
	GetProgramError() BACnetProgramErrorTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProgramErrorTagged
}

// BACnetConstructedDataReasonForHaltExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataReasonForHalt.
// This is useful for switch cases.
type BACnetConstructedDataReasonForHaltExactly interface {
	BACnetConstructedDataReasonForHalt
	isBACnetConstructedDataReasonForHalt() bool
}

// _BACnetConstructedDataReasonForHalt is the data-structure of this message
type _BACnetConstructedDataReasonForHalt struct {
	BACnetConstructedDataContract
	ProgramError BACnetProgramErrorTagged
}

var _ BACnetConstructedDataReasonForHalt = (*_BACnetConstructedDataReasonForHalt)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataReasonForHalt)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataReasonForHalt) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_REASON_FOR_HALT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetProgramError() BACnetProgramErrorTagged {
	return m.ProgramError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetActualValue() BACnetProgramErrorTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetProgramErrorTagged(m.GetProgramError())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReasonForHalt factory function for _BACnetConstructedDataReasonForHalt
func NewBACnetConstructedDataReasonForHalt(programError BACnetProgramErrorTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataReasonForHalt {
	_result := &_BACnetConstructedDataReasonForHalt{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ProgramError:                  programError,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReasonForHalt(structType any) BACnetConstructedDataReasonForHalt {
	if casted, ok := structType.(BACnetConstructedDataReasonForHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReasonForHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReasonForHalt) GetTypeName() string {
	return "BACnetConstructedDataReasonForHalt"
}

func (m *_BACnetConstructedDataReasonForHalt) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (programError)
	lengthInBits += m.ProgramError.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataReasonForHalt) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataReasonForHalt) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataReasonForHalt BACnetConstructedDataReasonForHalt, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReasonForHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	programError, err := ReadSimpleField[BACnetProgramErrorTagged](ctx, "programError", ReadComplex[BACnetProgramErrorTagged](BACnetProgramErrorTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'programError' field"))
	}
	m.ProgramError = programError

	actualValue, err := ReadVirtualField[BACnetProgramErrorTagged](ctx, "actualValue", (*BACnetProgramErrorTagged)(nil), programError)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReasonForHalt")
	}

	return m, nil
}

func (m *_BACnetConstructedDataReasonForHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataReasonForHalt) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReasonForHalt")
		}

		if err := WriteSimpleField[BACnetProgramErrorTagged](ctx, "programError", m.GetProgramError(), WriteComplex[BACnetProgramErrorTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'programError' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReasonForHalt")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataReasonForHalt) isBACnetConstructedDataReasonForHalt() bool {
	return true
}

func (m *_BACnetConstructedDataReasonForHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
