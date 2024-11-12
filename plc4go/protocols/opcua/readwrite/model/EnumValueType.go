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

// EnumValueType is the corresponding interface of EnumValueType
type EnumValueType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetValue returns Value (property field)
	GetValue() int64
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// IsEnumValueType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEnumValueType()
	// CreateBuilder creates a EnumValueTypeBuilder
	CreateEnumValueTypeBuilder() EnumValueTypeBuilder
}

// _EnumValueType is the data-structure of this message
type _EnumValueType struct {
	ExtensionObjectDefinitionContract
	Value       int64
	DisplayName LocalizedText
	Description LocalizedText
}

var _ EnumValueType = (*_EnumValueType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_EnumValueType)(nil)

// NewEnumValueType factory function for _EnumValueType
func NewEnumValueType(value int64, displayName LocalizedText, description LocalizedText) *_EnumValueType {
	if displayName == nil {
		panic("displayName of type LocalizedText for EnumValueType must not be nil")
	}
	if description == nil {
		panic("description of type LocalizedText for EnumValueType must not be nil")
	}
	_result := &_EnumValueType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Value:                             value,
		DisplayName:                       displayName,
		Description:                       description,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EnumValueTypeBuilder is a builder for EnumValueType
type EnumValueTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(value int64, displayName LocalizedText, description LocalizedText) EnumValueTypeBuilder
	// WithValue adds Value (property field)
	WithValue(int64) EnumValueTypeBuilder
	// WithDisplayName adds DisplayName (property field)
	WithDisplayName(LocalizedText) EnumValueTypeBuilder
	// WithDisplayNameBuilder adds DisplayName (property field) which is build by the builder
	WithDisplayNameBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) EnumValueTypeBuilder
	// WithDescription adds Description (property field)
	WithDescription(LocalizedText) EnumValueTypeBuilder
	// WithDescriptionBuilder adds Description (property field) which is build by the builder
	WithDescriptionBuilder(func(LocalizedTextBuilder) LocalizedTextBuilder) EnumValueTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the EnumValueType or returns an error if something is wrong
	Build() (EnumValueType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EnumValueType
}

// NewEnumValueTypeBuilder() creates a EnumValueTypeBuilder
func NewEnumValueTypeBuilder() EnumValueTypeBuilder {
	return &_EnumValueTypeBuilder{_EnumValueType: new(_EnumValueType)}
}

type _EnumValueTypeBuilder struct {
	*_EnumValueType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (EnumValueTypeBuilder) = (*_EnumValueTypeBuilder)(nil)

func (b *_EnumValueTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_EnumValueTypeBuilder) WithMandatoryFields(value int64, displayName LocalizedText, description LocalizedText) EnumValueTypeBuilder {
	return b.WithValue(value).WithDisplayName(displayName).WithDescription(description)
}

func (b *_EnumValueTypeBuilder) WithValue(value int64) EnumValueTypeBuilder {
	b.Value = value
	return b
}

func (b *_EnumValueTypeBuilder) WithDisplayName(displayName LocalizedText) EnumValueTypeBuilder {
	b.DisplayName = displayName
	return b
}

func (b *_EnumValueTypeBuilder) WithDisplayNameBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) EnumValueTypeBuilder {
	builder := builderSupplier(b.DisplayName.CreateLocalizedTextBuilder())
	var err error
	b.DisplayName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_EnumValueTypeBuilder) WithDescription(description LocalizedText) EnumValueTypeBuilder {
	b.Description = description
	return b
}

func (b *_EnumValueTypeBuilder) WithDescriptionBuilder(builderSupplier func(LocalizedTextBuilder) LocalizedTextBuilder) EnumValueTypeBuilder {
	builder := builderSupplier(b.Description.CreateLocalizedTextBuilder())
	var err error
	b.Description, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LocalizedTextBuilder failed"))
	}
	return b
}

func (b *_EnumValueTypeBuilder) Build() (EnumValueType, error) {
	if b.DisplayName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'displayName' not set"))
	}
	if b.Description == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'description' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EnumValueType.deepCopy(), nil
}

func (b *_EnumValueTypeBuilder) MustBuild() EnumValueType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EnumValueTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_EnumValueTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_EnumValueTypeBuilder) DeepCopy() any {
	_copy := b.CreateEnumValueTypeBuilder().(*_EnumValueTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEnumValueTypeBuilder creates a EnumValueTypeBuilder
func (b *_EnumValueType) CreateEnumValueTypeBuilder() EnumValueTypeBuilder {
	if b == nil {
		return NewEnumValueTypeBuilder()
	}
	return &_EnumValueTypeBuilder{_EnumValueType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EnumValueType) GetExtensionId() int32 {
	return int32(7596)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EnumValueType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EnumValueType) GetValue() int64 {
	return m.Value
}

func (m *_EnumValueType) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_EnumValueType) GetDescription() LocalizedText {
	return m.Description
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEnumValueType(structType any) EnumValueType {
	if casted, ok := structType.(EnumValueType); ok {
		return casted
	}
	if casted, ok := structType.(*EnumValueType); ok {
		return *casted
	}
	return nil
}

func (m *_EnumValueType) GetTypeName() string {
	return "EnumValueType"
}

func (m *_EnumValueType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 64

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_EnumValueType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EnumValueType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__enumValueType EnumValueType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EnumValueType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EnumValueType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadSimpleField(ctx, "value", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	displayName, err := ReadSimpleField[LocalizedText](ctx, "displayName", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'displayName' field"))
	}
	m.DisplayName = displayName

	description, err := ReadSimpleField[LocalizedText](ctx, "description", ReadComplex[LocalizedText](LocalizedTextParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'description' field"))
	}
	m.Description = description

	if closeErr := readBuffer.CloseContext("EnumValueType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EnumValueType")
	}

	return m, nil
}

func (m *_EnumValueType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EnumValueType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EnumValueType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EnumValueType")
		}

		if err := WriteSimpleField[int64](ctx, "value", m.GetValue(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "displayName", m.GetDisplayName(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'displayName' field")
		}

		if err := WriteSimpleField[LocalizedText](ctx, "description", m.GetDescription(), WriteComplex[LocalizedText](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'description' field")
		}

		if popErr := writeBuffer.PopContext("EnumValueType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EnumValueType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EnumValueType) IsEnumValueType() {}

func (m *_EnumValueType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EnumValueType) deepCopy() *_EnumValueType {
	if m == nil {
		return nil
	}
	_EnumValueTypeCopy := &_EnumValueType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.Value,
		utils.DeepCopy[LocalizedText](m.DisplayName),
		utils.DeepCopy[LocalizedText](m.Description),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _EnumValueTypeCopy
}

func (m *_EnumValueType) String() string {
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
