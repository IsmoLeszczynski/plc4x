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

// PubSubConfigurationRefDataType is the corresponding interface of PubSubConfigurationRefDataType
type PubSubConfigurationRefDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetConfigurationMask returns ConfigurationMask (property field)
	GetConfigurationMask() PubSubConfigurationRefMask
	// GetElementIndex returns ElementIndex (property field)
	GetElementIndex() uint16
	// GetConnectionIndex returns ConnectionIndex (property field)
	GetConnectionIndex() uint16
	// GetGroupIndex returns GroupIndex (property field)
	GetGroupIndex() uint16
	// IsPubSubConfigurationRefDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubConfigurationRefDataType()
	// CreateBuilder creates a PubSubConfigurationRefDataTypeBuilder
	CreatePubSubConfigurationRefDataTypeBuilder() PubSubConfigurationRefDataTypeBuilder
}

// _PubSubConfigurationRefDataType is the data-structure of this message
type _PubSubConfigurationRefDataType struct {
	ExtensionObjectDefinitionContract
	ConfigurationMask PubSubConfigurationRefMask
	ElementIndex      uint16
	ConnectionIndex   uint16
	GroupIndex        uint16
}

var _ PubSubConfigurationRefDataType = (*_PubSubConfigurationRefDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubConfigurationRefDataType)(nil)

// NewPubSubConfigurationRefDataType factory function for _PubSubConfigurationRefDataType
func NewPubSubConfigurationRefDataType(configurationMask PubSubConfigurationRefMask, elementIndex uint16, connectionIndex uint16, groupIndex uint16) *_PubSubConfigurationRefDataType {
	_result := &_PubSubConfigurationRefDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ConfigurationMask:                 configurationMask,
		ElementIndex:                      elementIndex,
		ConnectionIndex:                   connectionIndex,
		GroupIndex:                        groupIndex,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PubSubConfigurationRefDataTypeBuilder is a builder for PubSubConfigurationRefDataType
type PubSubConfigurationRefDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(configurationMask PubSubConfigurationRefMask, elementIndex uint16, connectionIndex uint16, groupIndex uint16) PubSubConfigurationRefDataTypeBuilder
	// WithConfigurationMask adds ConfigurationMask (property field)
	WithConfigurationMask(PubSubConfigurationRefMask) PubSubConfigurationRefDataTypeBuilder
	// WithElementIndex adds ElementIndex (property field)
	WithElementIndex(uint16) PubSubConfigurationRefDataTypeBuilder
	// WithConnectionIndex adds ConnectionIndex (property field)
	WithConnectionIndex(uint16) PubSubConfigurationRefDataTypeBuilder
	// WithGroupIndex adds GroupIndex (property field)
	WithGroupIndex(uint16) PubSubConfigurationRefDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the PubSubConfigurationRefDataType or returns an error if something is wrong
	Build() (PubSubConfigurationRefDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PubSubConfigurationRefDataType
}

// NewPubSubConfigurationRefDataTypeBuilder() creates a PubSubConfigurationRefDataTypeBuilder
func NewPubSubConfigurationRefDataTypeBuilder() PubSubConfigurationRefDataTypeBuilder {
	return &_PubSubConfigurationRefDataTypeBuilder{_PubSubConfigurationRefDataType: new(_PubSubConfigurationRefDataType)}
}

type _PubSubConfigurationRefDataTypeBuilder struct {
	*_PubSubConfigurationRefDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PubSubConfigurationRefDataTypeBuilder) = (*_PubSubConfigurationRefDataTypeBuilder)(nil)

func (b *_PubSubConfigurationRefDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PubSubConfigurationRefDataTypeBuilder) WithMandatoryFields(configurationMask PubSubConfigurationRefMask, elementIndex uint16, connectionIndex uint16, groupIndex uint16) PubSubConfigurationRefDataTypeBuilder {
	return b.WithConfigurationMask(configurationMask).WithElementIndex(elementIndex).WithConnectionIndex(connectionIndex).WithGroupIndex(groupIndex)
}

func (b *_PubSubConfigurationRefDataTypeBuilder) WithConfigurationMask(configurationMask PubSubConfigurationRefMask) PubSubConfigurationRefDataTypeBuilder {
	b.ConfigurationMask = configurationMask
	return b
}

func (b *_PubSubConfigurationRefDataTypeBuilder) WithElementIndex(elementIndex uint16) PubSubConfigurationRefDataTypeBuilder {
	b.ElementIndex = elementIndex
	return b
}

func (b *_PubSubConfigurationRefDataTypeBuilder) WithConnectionIndex(connectionIndex uint16) PubSubConfigurationRefDataTypeBuilder {
	b.ConnectionIndex = connectionIndex
	return b
}

func (b *_PubSubConfigurationRefDataTypeBuilder) WithGroupIndex(groupIndex uint16) PubSubConfigurationRefDataTypeBuilder {
	b.GroupIndex = groupIndex
	return b
}

func (b *_PubSubConfigurationRefDataTypeBuilder) Build() (PubSubConfigurationRefDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PubSubConfigurationRefDataType.deepCopy(), nil
}

func (b *_PubSubConfigurationRefDataTypeBuilder) MustBuild() PubSubConfigurationRefDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PubSubConfigurationRefDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_PubSubConfigurationRefDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PubSubConfigurationRefDataTypeBuilder) DeepCopy() any {
	_copy := b.CreatePubSubConfigurationRefDataTypeBuilder().(*_PubSubConfigurationRefDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePubSubConfigurationRefDataTypeBuilder creates a PubSubConfigurationRefDataTypeBuilder
func (b *_PubSubConfigurationRefDataType) CreatePubSubConfigurationRefDataTypeBuilder() PubSubConfigurationRefDataTypeBuilder {
	if b == nil {
		return NewPubSubConfigurationRefDataTypeBuilder()
	}
	return &_PubSubConfigurationRefDataTypeBuilder{_PubSubConfigurationRefDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfigurationRefDataType) GetExtensionId() int32 {
	return int32(25521)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfigurationRefDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfigurationRefDataType) GetConfigurationMask() PubSubConfigurationRefMask {
	return m.ConfigurationMask
}

func (m *_PubSubConfigurationRefDataType) GetElementIndex() uint16 {
	return m.ElementIndex
}

func (m *_PubSubConfigurationRefDataType) GetConnectionIndex() uint16 {
	return m.ConnectionIndex
}

func (m *_PubSubConfigurationRefDataType) GetGroupIndex() uint16 {
	return m.GroupIndex
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPubSubConfigurationRefDataType(structType any) PubSubConfigurationRefDataType {
	if casted, ok := structType.(PubSubConfigurationRefDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfigurationRefDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfigurationRefDataType) GetTypeName() string {
	return "PubSubConfigurationRefDataType"
}

func (m *_PubSubConfigurationRefDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (configurationMask)
	lengthInBits += 32

	// Simple field (elementIndex)
	lengthInBits += 16

	// Simple field (connectionIndex)
	lengthInBits += 16

	// Simple field (groupIndex)
	lengthInBits += 16

	return lengthInBits
}

func (m *_PubSubConfigurationRefDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubConfigurationRefDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__pubSubConfigurationRefDataType PubSubConfigurationRefDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubConfigurationRefDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfigurationRefDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	configurationMask, err := ReadEnumField[PubSubConfigurationRefMask](ctx, "configurationMask", "PubSubConfigurationRefMask", ReadEnum(PubSubConfigurationRefMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configurationMask' field"))
	}
	m.ConfigurationMask = configurationMask

	elementIndex, err := ReadSimpleField(ctx, "elementIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'elementIndex' field"))
	}
	m.ElementIndex = elementIndex

	connectionIndex, err := ReadSimpleField(ctx, "connectionIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionIndex' field"))
	}
	m.ConnectionIndex = connectionIndex

	groupIndex, err := ReadSimpleField(ctx, "groupIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupIndex' field"))
	}
	m.GroupIndex = groupIndex

	if closeErr := readBuffer.CloseContext("PubSubConfigurationRefDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfigurationRefDataType")
	}

	return m, nil
}

func (m *_PubSubConfigurationRefDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfigurationRefDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfigurationRefDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfigurationRefDataType")
		}

		if err := WriteSimpleEnumField[PubSubConfigurationRefMask](ctx, "configurationMask", "PubSubConfigurationRefMask", m.GetConfigurationMask(), WriteEnum[PubSubConfigurationRefMask, uint32](PubSubConfigurationRefMask.GetValue, PubSubConfigurationRefMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'configurationMask' field")
		}

		if err := WriteSimpleField[uint16](ctx, "elementIndex", m.GetElementIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'elementIndex' field")
		}

		if err := WriteSimpleField[uint16](ctx, "connectionIndex", m.GetConnectionIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionIndex' field")
		}

		if err := WriteSimpleField[uint16](ctx, "groupIndex", m.GetGroupIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'groupIndex' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfigurationRefDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfigurationRefDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfigurationRefDataType) IsPubSubConfigurationRefDataType() {}

func (m *_PubSubConfigurationRefDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PubSubConfigurationRefDataType) deepCopy() *_PubSubConfigurationRefDataType {
	if m == nil {
		return nil
	}
	_PubSubConfigurationRefDataTypeCopy := &_PubSubConfigurationRefDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ConfigurationMask,
		m.ElementIndex,
		m.ConnectionIndex,
		m.GroupIndex,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PubSubConfigurationRefDataTypeCopy
}

func (m *_PubSubConfigurationRefDataType) String() string {
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
