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

// VariantInt16 is the corresponding interface of VariantInt16
type VariantInt16 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []int16
}

// VariantInt16Exactly can be used when we want exactly this type and not a type which fulfills VariantInt16.
// This is useful for switch cases.
type VariantInt16Exactly interface {
	VariantInt16
	isVariantInt16() bool
}

// _VariantInt16 is the data-structure of this message
type _VariantInt16 struct {
	VariantContract
	ArrayLength *int32
	Value       []int16
}

var _ VariantInt16 = (*_VariantInt16)(nil)
var _ VariantRequirements = (*_VariantInt16)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantInt16) GetVariantType() uint8 {
	return uint8(4)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantInt16) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantInt16) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantInt16) GetValue() []int16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewVariantInt16 factory function for _VariantInt16
func NewVariantInt16(arrayLength *int32, value []int16, arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool) *_VariantInt16 {
	_result := &_VariantInt16{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantInt16(structType any) VariantInt16 {
	if casted, ok := structType.(VariantInt16); ok {
		return casted
	}
	if casted, ok := structType.(*VariantInt16); ok {
		return *casted
	}
	return nil
}

func (m *_VariantInt16) GetTypeName() string {
	return "VariantInt16"
}

func (m *_VariantInt16) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 16 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantInt16) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantInt16) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantInt16 VariantInt16, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantInt16"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantInt16")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[int16](ctx, "value", ReadSignedShort(readBuffer, uint8(16)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantInt16"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantInt16")
	}

	return m, nil
}

func (m *_VariantInt16) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantInt16) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantInt16"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantInt16")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteSignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantInt16"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantInt16")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantInt16) isVariantInt16() bool {
	return true
}

func (m *_VariantInt16) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
