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

// LiteralOperand is the corresponding interface of LiteralOperand
type LiteralOperand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetValue returns Value (property field)
	GetValue() Variant
	// IsLiteralOperand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLiteralOperand()
	// CreateBuilder creates a LiteralOperandBuilder
	CreateLiteralOperandBuilder() LiteralOperandBuilder
}

// _LiteralOperand is the data-structure of this message
type _LiteralOperand struct {
	ExtensionObjectDefinitionContract
	Value Variant
}

var _ LiteralOperand = (*_LiteralOperand)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_LiteralOperand)(nil)

// NewLiteralOperand factory function for _LiteralOperand
func NewLiteralOperand(value Variant) *_LiteralOperand {
	if value == nil {
		panic("value of type Variant for LiteralOperand must not be nil")
	}
	_result := &_LiteralOperand{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Value:                             value,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LiteralOperandBuilder is a builder for LiteralOperand
type LiteralOperandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value Variant) LiteralOperandBuilder
	// WithValue adds Value (property field)
	WithValue(Variant) LiteralOperandBuilder
	// WithValueBuilder adds Value (property field) which is build by the builder
	WithValueBuilder(func(VariantBuilder) VariantBuilder) LiteralOperandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the LiteralOperand or returns an error if something is wrong
	Build() (LiteralOperand, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LiteralOperand
}

// NewLiteralOperandBuilder() creates a LiteralOperandBuilder
func NewLiteralOperandBuilder() LiteralOperandBuilder {
	return &_LiteralOperandBuilder{_LiteralOperand: new(_LiteralOperand)}
}

type _LiteralOperandBuilder struct {
	*_LiteralOperand

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (LiteralOperandBuilder) = (*_LiteralOperandBuilder)(nil)

func (b *_LiteralOperandBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_LiteralOperandBuilder) WithMandatoryFields(value Variant) LiteralOperandBuilder {
	return b.WithValue(value)
}

func (b *_LiteralOperandBuilder) WithValue(value Variant) LiteralOperandBuilder {
	b.Value = value
	return b
}

func (b *_LiteralOperandBuilder) WithValueBuilder(builderSupplier func(VariantBuilder) VariantBuilder) LiteralOperandBuilder {
	builder := builderSupplier(b.Value.CreateVariantBuilder())
	var err error
	b.Value, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "VariantBuilder failed"))
	}
	return b
}

func (b *_LiteralOperandBuilder) Build() (LiteralOperand, error) {
	if b.Value == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'value' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LiteralOperand.deepCopy(), nil
}

func (b *_LiteralOperandBuilder) MustBuild() LiteralOperand {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_LiteralOperandBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_LiteralOperandBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_LiteralOperandBuilder) DeepCopy() any {
	_copy := b.CreateLiteralOperandBuilder().(*_LiteralOperandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLiteralOperandBuilder creates a LiteralOperandBuilder
func (b *_LiteralOperand) CreateLiteralOperandBuilder() LiteralOperandBuilder {
	if b == nil {
		return NewLiteralOperandBuilder()
	}
	return &_LiteralOperandBuilder{_LiteralOperand: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LiteralOperand) GetExtensionId() int32 {
	return int32(597)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LiteralOperand) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LiteralOperand) GetValue() Variant {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastLiteralOperand(structType any) LiteralOperand {
	if casted, ok := structType.(LiteralOperand); ok {
		return casted
	}
	if casted, ok := structType.(*LiteralOperand); ok {
		return *casted
	}
	return nil
}

func (m *_LiteralOperand) GetTypeName() string {
	return "LiteralOperand"
}

func (m *_LiteralOperand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_LiteralOperand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LiteralOperand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__literalOperand LiteralOperand, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LiteralOperand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LiteralOperand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField[Variant](ctx, "value", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("LiteralOperand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LiteralOperand")
	}

	return m, nil
}

func (m *_LiteralOperand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LiteralOperand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LiteralOperand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LiteralOperand")
		}

		if err := WriteSimpleField[Variant](ctx, "value", m.GetValue(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("LiteralOperand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LiteralOperand")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LiteralOperand) IsLiteralOperand() {}

func (m *_LiteralOperand) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LiteralOperand) deepCopy() *_LiteralOperand {
	if m == nil {
		return nil
	}
	_LiteralOperandCopy := &_LiteralOperand{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[Variant](m.Value),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _LiteralOperandCopy
}

func (m *_LiteralOperand) String() string {
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
