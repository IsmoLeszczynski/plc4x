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

// VariantByte is the corresponding interface of VariantByte
type VariantByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []uint8
	// IsVariantByte is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVariantByte()
	// CreateBuilder creates a VariantByteBuilder
	CreateVariantByteBuilder() VariantByteBuilder
}

// _VariantByte is the data-structure of this message
type _VariantByte struct {
	VariantContract
	ArrayLength *int32
	Value       []uint8
}

var _ VariantByte = (*_VariantByte)(nil)
var _ VariantRequirements = (*_VariantByte)(nil)

// NewVariantByte factory function for _VariantByte
func NewVariantByte(arrayLengthSpecified bool, arrayDimensionsSpecified bool, noOfArrayDimensions *int32, arrayDimensions []bool, arrayLength *int32, value []uint8) *_VariantByte {
	_result := &_VariantByte{
		VariantContract: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
		ArrayLength:     arrayLength,
		Value:           value,
	}
	_result.VariantContract.(*_Variant)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VariantByteBuilder is a builder for VariantByte
type VariantByteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value []uint8) VariantByteBuilder
	// WithArrayLength adds ArrayLength (property field)
	WithOptionalArrayLength(int32) VariantByteBuilder
	// WithValue adds Value (property field)
	WithValue(...uint8) VariantByteBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() VariantBuilder
	// Build builds the VariantByte or returns an error if something is wrong
	Build() (VariantByte, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VariantByte
}

// NewVariantByteBuilder() creates a VariantByteBuilder
func NewVariantByteBuilder() VariantByteBuilder {
	return &_VariantByteBuilder{_VariantByte: new(_VariantByte)}
}

type _VariantByteBuilder struct {
	*_VariantByte

	parentBuilder *_VariantBuilder

	err *utils.MultiError
}

var _ (VariantByteBuilder) = (*_VariantByteBuilder)(nil)

func (b *_VariantByteBuilder) setParent(contract VariantContract) {
	b.VariantContract = contract
}

func (b *_VariantByteBuilder) WithMandatoryFields(value []uint8) VariantByteBuilder {
	return b.WithValue(value...)
}

func (b *_VariantByteBuilder) WithOptionalArrayLength(arrayLength int32) VariantByteBuilder {
	b.ArrayLength = &arrayLength
	return b
}

func (b *_VariantByteBuilder) WithValue(value ...uint8) VariantByteBuilder {
	b.Value = value
	return b
}

func (b *_VariantByteBuilder) Build() (VariantByte, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VariantByte.deepCopy(), nil
}

func (b *_VariantByteBuilder) MustBuild() VariantByte {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VariantByteBuilder) Done() VariantBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewVariantBuilder().(*_VariantBuilder)
	}
	return b.parentBuilder
}

func (b *_VariantByteBuilder) buildForVariant() (Variant, error) {
	return b.Build()
}

func (b *_VariantByteBuilder) DeepCopy() any {
	_copy := b.CreateVariantByteBuilder().(*_VariantByteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVariantByteBuilder creates a VariantByteBuilder
func (b *_VariantByte) CreateVariantByteBuilder() VariantByteBuilder {
	if b == nil {
		return NewVariantByteBuilder()
	}
	return &_VariantByteBuilder{_VariantByte: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantByte) GetVariantType() uint8 {
	return uint8(3)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantByte) GetParent() VariantContract {
	return m.VariantContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantByte) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantByte) GetValue() []uint8 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVariantByte(structType any) VariantByte {
	if casted, ok := structType.(VariantByte); ok {
		return casted
	}
	if casted, ok := structType.(*VariantByte); ok {
		return *casted
	}
	return nil
}

func (m *_VariantByte) GetTypeName() string {
	return "VariantByte"
}

func (m *_VariantByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.VariantContract.(*_Variant).getLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}

func (m *_VariantByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VariantByte) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Variant, arrayLengthSpecified bool) (__variantByte VariantByte, err error) {
	m.VariantContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VariantByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var arrayLength *int32
	arrayLength, err = ReadOptionalField[int32](ctx, "arrayLength", ReadSignedInt(readBuffer, uint8(32)), arrayLengthSpecified)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayLength' field"))
	}
	m.ArrayLength = arrayLength

	value, err := ReadCountArrayField[uint8](ctx, "value", ReadUnsignedByte(readBuffer, uint8(8)), uint64(utils.InlineIf(bool((arrayLength) == (nil)), func() any { return int32(int32(1)) }, func() any { return int32((*arrayLength)) }).(int32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("VariantByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantByte")
	}

	return m, nil
}

func (m *_VariantByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantByte")
		}

		if err := WriteOptionalField[int32](ctx, "arrayLength", m.GetArrayLength(), WriteSignedInt(writeBuffer, 32), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayLength' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "value", m.GetValue(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("VariantByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantByte")
		}
		return nil
	}
	return m.VariantContract.(*_Variant).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VariantByte) IsVariantByte() {}

func (m *_VariantByte) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VariantByte) deepCopy() *_VariantByte {
	if m == nil {
		return nil
	}
	_VariantByteCopy := &_VariantByte{
		m.VariantContract.(*_Variant).deepCopy(),
		utils.CopyPtr[int32](m.ArrayLength),
		utils.DeepCopySlice[uint8, uint8](m.Value),
	}
	m.VariantContract.(*_Variant)._SubType = m
	return _VariantByteCopy
}

func (m *_VariantByte) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
