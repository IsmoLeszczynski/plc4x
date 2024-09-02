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

// BACnetConstructedDataLastRestartReason is the corresponding interface of BACnetConstructedDataLastRestartReason
type BACnetConstructedDataLastRestartReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastRestartReason returns LastRestartReason (property field)
	GetLastRestartReason() BACnetRestartReasonTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetRestartReasonTagged
}

// BACnetConstructedDataLastRestartReasonExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastRestartReason.
// This is useful for switch cases.
type BACnetConstructedDataLastRestartReasonExactly interface {
	BACnetConstructedDataLastRestartReason
	isBACnetConstructedDataLastRestartReason() bool
}

// _BACnetConstructedDataLastRestartReason is the data-structure of this message
type _BACnetConstructedDataLastRestartReason struct {
	BACnetConstructedDataContract
	LastRestartReason BACnetRestartReasonTagged
}

var _ BACnetConstructedDataLastRestartReason = (*_BACnetConstructedDataLastRestartReason)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLastRestartReason)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastRestartReason) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_RESTART_REASON
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetLastRestartReason() BACnetRestartReasonTagged {
	return m.LastRestartReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestartReason) GetActualValue() BACnetRestartReasonTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetRestartReasonTagged(m.GetLastRestartReason())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastRestartReason factory function for _BACnetConstructedDataLastRestartReason
func NewBACnetConstructedDataLastRestartReason(lastRestartReason BACnetRestartReasonTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastRestartReason {
	_result := &_BACnetConstructedDataLastRestartReason{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LastRestartReason:             lastRestartReason,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastRestartReason(structType any) BACnetConstructedDataLastRestartReason {
	if casted, ok := structType.(BACnetConstructedDataLastRestartReason); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestartReason); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastRestartReason) GetTypeName() string {
	return "BACnetConstructedDataLastRestartReason"
}

func (m *_BACnetConstructedDataLastRestartReason) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (lastRestartReason)
	lengthInBits += m.LastRestartReason.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastRestartReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLastRestartReason) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLastRestartReason BACnetConstructedDataLastRestartReason, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestartReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastRestartReason")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lastRestartReason, err := ReadSimpleField[BACnetRestartReasonTagged](ctx, "lastRestartReason", ReadComplex[BACnetRestartReasonTagged](BACnetRestartReasonTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastRestartReason' field"))
	}
	m.LastRestartReason = lastRestartReason

	actualValue, err := ReadVirtualField[BACnetRestartReasonTagged](ctx, "actualValue", (*BACnetRestartReasonTagged)(nil), lastRestartReason)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestartReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastRestartReason")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLastRestartReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastRestartReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestartReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastRestartReason")
		}

		if err := WriteSimpleField[BACnetRestartReasonTagged](ctx, "lastRestartReason", m.GetLastRestartReason(), WriteComplex[BACnetRestartReasonTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastRestartReason' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestartReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastRestartReason")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastRestartReason) isBACnetConstructedDataLastRestartReason() bool {
	return true
}

func (m *_BACnetConstructedDataLastRestartReason) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
