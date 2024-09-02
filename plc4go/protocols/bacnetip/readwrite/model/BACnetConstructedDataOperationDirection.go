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

// BACnetConstructedDataOperationDirection is the corresponding interface of BACnetConstructedDataOperationDirection
type BACnetConstructedDataOperationDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOperationDirection returns OperationDirection (property field)
	GetOperationDirection() BACnetEscalatorOperationDirectionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEscalatorOperationDirectionTagged
}

// BACnetConstructedDataOperationDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOperationDirection.
// This is useful for switch cases.
type BACnetConstructedDataOperationDirectionExactly interface {
	BACnetConstructedDataOperationDirection
	isBACnetConstructedDataOperationDirection() bool
}

// _BACnetConstructedDataOperationDirection is the data-structure of this message
type _BACnetConstructedDataOperationDirection struct {
	BACnetConstructedDataContract
	OperationDirection BACnetEscalatorOperationDirectionTagged
}

var _ BACnetConstructedDataOperationDirection = (*_BACnetConstructedDataOperationDirection)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOperationDirection)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOperationDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OPERATION_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetOperationDirection() BACnetEscalatorOperationDirectionTagged {
	return m.OperationDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOperationDirection) GetActualValue() BACnetEscalatorOperationDirectionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEscalatorOperationDirectionTagged(m.GetOperationDirection())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOperationDirection factory function for _BACnetConstructedDataOperationDirection
func NewBACnetConstructedDataOperationDirection(operationDirection BACnetEscalatorOperationDirectionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOperationDirection {
	_result := &_BACnetConstructedDataOperationDirection{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OperationDirection:            operationDirection,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOperationDirection(structType any) BACnetConstructedDataOperationDirection {
	if casted, ok := structType.(BACnetConstructedDataOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOperationDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOperationDirection) GetTypeName() string {
	return "BACnetConstructedDataOperationDirection"
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (operationDirection)
	lengthInBits += m.OperationDirection.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOperationDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOperationDirection) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOperationDirection BACnetConstructedDataOperationDirection, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operationDirection, err := ReadSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "operationDirection", ReadComplex[BACnetEscalatorOperationDirectionTagged](BACnetEscalatorOperationDirectionTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operationDirection' field"))
	}
	m.OperationDirection = operationDirection

	actualValue, err := ReadVirtualField[BACnetEscalatorOperationDirectionTagged](ctx, "actualValue", (*BACnetEscalatorOperationDirectionTagged)(nil), operationDirection)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOperationDirection")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOperationDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOperationDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOperationDirection")
		}

		if err := WriteSimpleField[BACnetEscalatorOperationDirectionTagged](ctx, "operationDirection", m.GetOperationDirection(), WriteComplex[BACnetEscalatorOperationDirectionTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'operationDirection' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOperationDirection")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOperationDirection) isBACnetConstructedDataOperationDirection() bool {
	return true
}

func (m *_BACnetConstructedDataOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
