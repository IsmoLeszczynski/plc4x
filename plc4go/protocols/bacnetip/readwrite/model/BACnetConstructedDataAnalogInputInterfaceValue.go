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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAnalogInputInterfaceValue is the corresponding interface of BACnetConstructedDataAnalogInputInterfaceValue
type BACnetConstructedDataAnalogInputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalREAL
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalREAL
}

// BACnetConstructedDataAnalogInputInterfaceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogInputInterfaceValue.
// This is useful for switch cases.
type BACnetConstructedDataAnalogInputInterfaceValueExactly interface {
	BACnetConstructedDataAnalogInputInterfaceValue
	isBACnetConstructedDataAnalogInputInterfaceValue() bool
}

// _BACnetConstructedDataAnalogInputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataAnalogInputInterfaceValue struct {
	*_BACnetConstructedData
	InterfaceValue BACnetOptionalREAL
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_INPUT
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetInterfaceValue() BACnetOptionalREAL {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetActualValue() BACnetOptionalREAL {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalREAL(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAnalogInputInterfaceValue factory function for _BACnetConstructedDataAnalogInputInterfaceValue
func NewBACnetConstructedDataAnalogInputInterfaceValue(interfaceValue BACnetOptionalREAL, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogInputInterfaceValue {
	_result := &_BACnetConstructedDataAnalogInputInterfaceValue{
		InterfaceValue:         interfaceValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogInputInterfaceValue(structType any) BACnetConstructedDataAnalogInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogInputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogInputInterfaceValue"
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAnalogInputInterfaceValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogInputInterfaceValue, error) {
	return BACnetConstructedDataAnalogInputInterfaceValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAnalogInputInterfaceValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogInputInterfaceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogInputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogInputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceValue)
	if pullErr := readBuffer.PullContext("interfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for interfaceValue")
	}
	_interfaceValue, _interfaceValueErr := BACnetOptionalREALParseWithBuffer(ctx, readBuffer)
	if _interfaceValueErr != nil {
		return nil, errors.Wrap(_interfaceValueErr, "Error parsing 'interfaceValue' field of BACnetConstructedDataAnalogInputInterfaceValue")
	}
	interfaceValue := _interfaceValue.(BACnetOptionalREAL)
	if closeErr := readBuffer.CloseContext("interfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for interfaceValue")
	}

	// Virtual field
	_actualValue := interfaceValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogInputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogInputInterfaceValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAnalogInputInterfaceValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		InterfaceValue: interfaceValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogInputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogInputInterfaceValue")
		}

		// Simple Field (interfaceValue)
		if pushErr := writeBuffer.PushContext("interfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for interfaceValue")
		}
		_interfaceValueErr := writeBuffer.WriteSerializable(ctx, m.GetInterfaceValue())
		if popErr := writeBuffer.PopContext("interfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for interfaceValue")
		}
		if _interfaceValueErr != nil {
			return errors.Wrap(_interfaceValueErr, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogInputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogInputInterfaceValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) isBACnetConstructedDataAnalogInputInterfaceValue() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogInputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
