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

// PublishedDataSetCustomSourceDataType is the corresponding interface of PublishedDataSetCustomSourceDataType
type PublishedDataSetCustomSourceDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetCyclicDataSet returns CyclicDataSet (property field)
	GetCyclicDataSet() bool
	// IsPublishedDataSetCustomSourceDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPublishedDataSetCustomSourceDataType()
	// CreateBuilder creates a PublishedDataSetCustomSourceDataTypeBuilder
	CreatePublishedDataSetCustomSourceDataTypeBuilder() PublishedDataSetCustomSourceDataTypeBuilder
}

// _PublishedDataSetCustomSourceDataType is the data-structure of this message
type _PublishedDataSetCustomSourceDataType struct {
	ExtensionObjectDefinitionContract
	CyclicDataSet bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ PublishedDataSetCustomSourceDataType = (*_PublishedDataSetCustomSourceDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PublishedDataSetCustomSourceDataType)(nil)

// NewPublishedDataSetCustomSourceDataType factory function for _PublishedDataSetCustomSourceDataType
func NewPublishedDataSetCustomSourceDataType(cyclicDataSet bool) *_PublishedDataSetCustomSourceDataType {
	_result := &_PublishedDataSetCustomSourceDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		CyclicDataSet:                     cyclicDataSet,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PublishedDataSetCustomSourceDataTypeBuilder is a builder for PublishedDataSetCustomSourceDataType
type PublishedDataSetCustomSourceDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(cyclicDataSet bool) PublishedDataSetCustomSourceDataTypeBuilder
	// WithCyclicDataSet adds CyclicDataSet (property field)
	WithCyclicDataSet(bool) PublishedDataSetCustomSourceDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the PublishedDataSetCustomSourceDataType or returns an error if something is wrong
	Build() (PublishedDataSetCustomSourceDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PublishedDataSetCustomSourceDataType
}

// NewPublishedDataSetCustomSourceDataTypeBuilder() creates a PublishedDataSetCustomSourceDataTypeBuilder
func NewPublishedDataSetCustomSourceDataTypeBuilder() PublishedDataSetCustomSourceDataTypeBuilder {
	return &_PublishedDataSetCustomSourceDataTypeBuilder{_PublishedDataSetCustomSourceDataType: new(_PublishedDataSetCustomSourceDataType)}
}

type _PublishedDataSetCustomSourceDataTypeBuilder struct {
	*_PublishedDataSetCustomSourceDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PublishedDataSetCustomSourceDataTypeBuilder) = (*_PublishedDataSetCustomSourceDataTypeBuilder)(nil)

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) WithMandatoryFields(cyclicDataSet bool) PublishedDataSetCustomSourceDataTypeBuilder {
	return b.WithCyclicDataSet(cyclicDataSet)
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) WithCyclicDataSet(cyclicDataSet bool) PublishedDataSetCustomSourceDataTypeBuilder {
	b.CyclicDataSet = cyclicDataSet
	return b
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) Build() (PublishedDataSetCustomSourceDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PublishedDataSetCustomSourceDataType.deepCopy(), nil
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) MustBuild() PublishedDataSetCustomSourceDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PublishedDataSetCustomSourceDataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePublishedDataSetCustomSourceDataTypeBuilder().(*_PublishedDataSetCustomSourceDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePublishedDataSetCustomSourceDataTypeBuilder creates a PublishedDataSetCustomSourceDataTypeBuilder
func (b *_PublishedDataSetCustomSourceDataType) CreatePublishedDataSetCustomSourceDataTypeBuilder() PublishedDataSetCustomSourceDataTypeBuilder {
	if b == nil {
		return NewPublishedDataSetCustomSourceDataTypeBuilder()
	}
	return &_PublishedDataSetCustomSourceDataTypeBuilder{_PublishedDataSetCustomSourceDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PublishedDataSetCustomSourceDataType) GetExtensionId() int32 {
	return int32(25271)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PublishedDataSetCustomSourceDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PublishedDataSetCustomSourceDataType) GetCyclicDataSet() bool {
	return m.CyclicDataSet
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPublishedDataSetCustomSourceDataType(structType any) PublishedDataSetCustomSourceDataType {
	if casted, ok := structType.(PublishedDataSetCustomSourceDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PublishedDataSetCustomSourceDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PublishedDataSetCustomSourceDataType) GetTypeName() string {
	return "PublishedDataSetCustomSourceDataType"
}

func (m *_PublishedDataSetCustomSourceDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (cyclicDataSet)
	lengthInBits += 1

	return lengthInBits
}

func (m *_PublishedDataSetCustomSourceDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PublishedDataSetCustomSourceDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__publishedDataSetCustomSourceDataType PublishedDataSetCustomSourceDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PublishedDataSetCustomSourceDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PublishedDataSetCustomSourceDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	cyclicDataSet, err := ReadSimpleField(ctx, "cyclicDataSet", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'cyclicDataSet' field"))
	}
	m.CyclicDataSet = cyclicDataSet

	if closeErr := readBuffer.CloseContext("PublishedDataSetCustomSourceDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PublishedDataSetCustomSourceDataType")
	}

	return m, nil
}

func (m *_PublishedDataSetCustomSourceDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PublishedDataSetCustomSourceDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PublishedDataSetCustomSourceDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PublishedDataSetCustomSourceDataType")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "cyclicDataSet", m.GetCyclicDataSet(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'cyclicDataSet' field")
		}

		if popErr := writeBuffer.PopContext("PublishedDataSetCustomSourceDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PublishedDataSetCustomSourceDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PublishedDataSetCustomSourceDataType) IsPublishedDataSetCustomSourceDataType() {}

func (m *_PublishedDataSetCustomSourceDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PublishedDataSetCustomSourceDataType) deepCopy() *_PublishedDataSetCustomSourceDataType {
	if m == nil {
		return nil
	}
	_PublishedDataSetCustomSourceDataTypeCopy := &_PublishedDataSetCustomSourceDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.CyclicDataSet,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PublishedDataSetCustomSourceDataTypeCopy
}

func (m *_PublishedDataSetCustomSourceDataType) String() string {
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
