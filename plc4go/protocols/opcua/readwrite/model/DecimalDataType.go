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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DecimalDataType is the corresponding interface of DecimalDataType
type DecimalDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetScale returns Scale (property field)
	GetScale() int16
	// GetValue returns Value (property field)
	GetValue() PascalByteString
}

// DecimalDataTypeExactly can be used when we want exactly this type and not a type which fulfills DecimalDataType.
// This is useful for switch cases.
type DecimalDataTypeExactly interface {
	DecimalDataType
	isDecimalDataType() bool
}

// _DecimalDataType is the data-structure of this message
type _DecimalDataType struct {
	*_ExtensionObjectDefinition
	Scale int16
	Value PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DecimalDataType) GetIdentifier() string {
	return "17863"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DecimalDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_DecimalDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DecimalDataType) GetScale() int16 {
	return m.Scale
}

func (m *_DecimalDataType) GetValue() PascalByteString {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDecimalDataType factory function for _DecimalDataType
func NewDecimalDataType(scale int16, value PascalByteString) *_DecimalDataType {
	_result := &_DecimalDataType{
		Scale:                      scale,
		Value:                      value,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDecimalDataType(structType any) DecimalDataType {
	if casted, ok := structType.(DecimalDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DecimalDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DecimalDataType) GetTypeName() string {
	return "DecimalDataType"
}

func (m *_DecimalDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (scale)
	lengthInBits += 16

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DecimalDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DecimalDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (DecimalDataType, error) {
	return DecimalDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func DecimalDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (DecimalDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DecimalDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DecimalDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (scale)
	_scale, _scaleErr := readBuffer.ReadInt16("scale", 16)
	if _scaleErr != nil {
		return nil, errors.Wrap(_scaleErr, "Error parsing 'scale' field of DecimalDataType")
	}
	scale := _scale

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of DecimalDataType")
	}
	value := _value.(PascalByteString)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("DecimalDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DecimalDataType")
	}

	// Create a partially initialized instance
	_child := &_DecimalDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Scale:                      scale,
		Value:                      value,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_DecimalDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DecimalDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DecimalDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DecimalDataType")
		}

		// Simple Field (scale)
		scale := int16(m.GetScale())
		_scaleErr := writeBuffer.WriteInt16("scale", 16, (scale))
		if _scaleErr != nil {
			return errors.Wrap(_scaleErr, "Error serializing 'scale' field")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("DecimalDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DecimalDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DecimalDataType) isDecimalDataType() bool {
	return true
}

func (m *_DecimalDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
