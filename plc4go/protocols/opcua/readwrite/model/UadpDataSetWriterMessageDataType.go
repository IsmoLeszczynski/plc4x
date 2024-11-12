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

// UadpDataSetWriterMessageDataType is the corresponding interface of UadpDataSetWriterMessageDataType
type UadpDataSetWriterMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetDataSetMessageContentMask returns DataSetMessageContentMask (property field)
	GetDataSetMessageContentMask() UadpDataSetMessageContentMask
	// GetConfiguredSize returns ConfiguredSize (property field)
	GetConfiguredSize() uint16
	// GetNetworkMessageNumber returns NetworkMessageNumber (property field)
	GetNetworkMessageNumber() uint16
	// GetDataSetOffset returns DataSetOffset (property field)
	GetDataSetOffset() uint16
	// IsUadpDataSetWriterMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUadpDataSetWriterMessageDataType()
	// CreateBuilder creates a UadpDataSetWriterMessageDataTypeBuilder
	CreateUadpDataSetWriterMessageDataTypeBuilder() UadpDataSetWriterMessageDataTypeBuilder
}

// _UadpDataSetWriterMessageDataType is the data-structure of this message
type _UadpDataSetWriterMessageDataType struct {
	ExtensionObjectDefinitionContract
	DataSetMessageContentMask UadpDataSetMessageContentMask
	ConfiguredSize            uint16
	NetworkMessageNumber      uint16
	DataSetOffset             uint16
}

var _ UadpDataSetWriterMessageDataType = (*_UadpDataSetWriterMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UadpDataSetWriterMessageDataType)(nil)

// NewUadpDataSetWriterMessageDataType factory function for _UadpDataSetWriterMessageDataType
func NewUadpDataSetWriterMessageDataType(dataSetMessageContentMask UadpDataSetMessageContentMask, configuredSize uint16, networkMessageNumber uint16, dataSetOffset uint16) *_UadpDataSetWriterMessageDataType {
	_result := &_UadpDataSetWriterMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		DataSetMessageContentMask:         dataSetMessageContentMask,
		ConfiguredSize:                    configuredSize,
		NetworkMessageNumber:              networkMessageNumber,
		DataSetOffset:                     dataSetOffset,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UadpDataSetWriterMessageDataTypeBuilder is a builder for UadpDataSetWriterMessageDataType
type UadpDataSetWriterMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(dataSetMessageContentMask UadpDataSetMessageContentMask, configuredSize uint16, networkMessageNumber uint16, dataSetOffset uint16) UadpDataSetWriterMessageDataTypeBuilder
	// WithDataSetMessageContentMask adds DataSetMessageContentMask (property field)
	WithDataSetMessageContentMask(UadpDataSetMessageContentMask) UadpDataSetWriterMessageDataTypeBuilder
	// WithConfiguredSize adds ConfiguredSize (property field)
	WithConfiguredSize(uint16) UadpDataSetWriterMessageDataTypeBuilder
	// WithNetworkMessageNumber adds NetworkMessageNumber (property field)
	WithNetworkMessageNumber(uint16) UadpDataSetWriterMessageDataTypeBuilder
	// WithDataSetOffset adds DataSetOffset (property field)
	WithDataSetOffset(uint16) UadpDataSetWriterMessageDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the UadpDataSetWriterMessageDataType or returns an error if something is wrong
	Build() (UadpDataSetWriterMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UadpDataSetWriterMessageDataType
}

// NewUadpDataSetWriterMessageDataTypeBuilder() creates a UadpDataSetWriterMessageDataTypeBuilder
func NewUadpDataSetWriterMessageDataTypeBuilder() UadpDataSetWriterMessageDataTypeBuilder {
	return &_UadpDataSetWriterMessageDataTypeBuilder{_UadpDataSetWriterMessageDataType: new(_UadpDataSetWriterMessageDataType)}
}

type _UadpDataSetWriterMessageDataTypeBuilder struct {
	*_UadpDataSetWriterMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UadpDataSetWriterMessageDataTypeBuilder) = (*_UadpDataSetWriterMessageDataTypeBuilder)(nil)

func (b *_UadpDataSetWriterMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) WithMandatoryFields(dataSetMessageContentMask UadpDataSetMessageContentMask, configuredSize uint16, networkMessageNumber uint16, dataSetOffset uint16) UadpDataSetWriterMessageDataTypeBuilder {
	return b.WithDataSetMessageContentMask(dataSetMessageContentMask).WithConfiguredSize(configuredSize).WithNetworkMessageNumber(networkMessageNumber).WithDataSetOffset(dataSetOffset)
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) WithDataSetMessageContentMask(dataSetMessageContentMask UadpDataSetMessageContentMask) UadpDataSetWriterMessageDataTypeBuilder {
	b.DataSetMessageContentMask = dataSetMessageContentMask
	return b
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) WithConfiguredSize(configuredSize uint16) UadpDataSetWriterMessageDataTypeBuilder {
	b.ConfiguredSize = configuredSize
	return b
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) WithNetworkMessageNumber(networkMessageNumber uint16) UadpDataSetWriterMessageDataTypeBuilder {
	b.NetworkMessageNumber = networkMessageNumber
	return b
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) WithDataSetOffset(dataSetOffset uint16) UadpDataSetWriterMessageDataTypeBuilder {
	b.DataSetOffset = dataSetOffset
	return b
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) Build() (UadpDataSetWriterMessageDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UadpDataSetWriterMessageDataType.deepCopy(), nil
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) MustBuild() UadpDataSetWriterMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UadpDataSetWriterMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateUadpDataSetWriterMessageDataTypeBuilder().(*_UadpDataSetWriterMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUadpDataSetWriterMessageDataTypeBuilder creates a UadpDataSetWriterMessageDataTypeBuilder
func (b *_UadpDataSetWriterMessageDataType) CreateUadpDataSetWriterMessageDataTypeBuilder() UadpDataSetWriterMessageDataTypeBuilder {
	if b == nil {
		return NewUadpDataSetWriterMessageDataTypeBuilder()
	}
	return &_UadpDataSetWriterMessageDataTypeBuilder{_UadpDataSetWriterMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UadpDataSetWriterMessageDataType) GetExtensionId() int32 {
	return int32(15654)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UadpDataSetWriterMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UadpDataSetWriterMessageDataType) GetDataSetMessageContentMask() UadpDataSetMessageContentMask {
	return m.DataSetMessageContentMask
}

func (m *_UadpDataSetWriterMessageDataType) GetConfiguredSize() uint16 {
	return m.ConfiguredSize
}

func (m *_UadpDataSetWriterMessageDataType) GetNetworkMessageNumber() uint16 {
	return m.NetworkMessageNumber
}

func (m *_UadpDataSetWriterMessageDataType) GetDataSetOffset() uint16 {
	return m.DataSetOffset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastUadpDataSetWriterMessageDataType(structType any) UadpDataSetWriterMessageDataType {
	if casted, ok := structType.(UadpDataSetWriterMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*UadpDataSetWriterMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_UadpDataSetWriterMessageDataType) GetTypeName() string {
	return "UadpDataSetWriterMessageDataType"
}

func (m *_UadpDataSetWriterMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (dataSetMessageContentMask)
	lengthInBits += 32

	// Simple field (configuredSize)
	lengthInBits += 16

	// Simple field (networkMessageNumber)
	lengthInBits += 16

	// Simple field (dataSetOffset)
	lengthInBits += 16

	return lengthInBits
}

func (m *_UadpDataSetWriterMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UadpDataSetWriterMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__uadpDataSetWriterMessageDataType UadpDataSetWriterMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UadpDataSetWriterMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UadpDataSetWriterMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	dataSetMessageContentMask, err := ReadEnumField[UadpDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "UadpDataSetMessageContentMask", ReadEnum(UadpDataSetMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetMessageContentMask' field"))
	}
	m.DataSetMessageContentMask = dataSetMessageContentMask

	configuredSize, err := ReadSimpleField(ctx, "configuredSize", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'configuredSize' field"))
	}
	m.ConfiguredSize = configuredSize

	networkMessageNumber, err := ReadSimpleField(ctx, "networkMessageNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkMessageNumber' field"))
	}
	m.NetworkMessageNumber = networkMessageNumber

	dataSetOffset, err := ReadSimpleField(ctx, "dataSetOffset", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'dataSetOffset' field"))
	}
	m.DataSetOffset = dataSetOffset

	if closeErr := readBuffer.CloseContext("UadpDataSetWriterMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UadpDataSetWriterMessageDataType")
	}

	return m, nil
}

func (m *_UadpDataSetWriterMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UadpDataSetWriterMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UadpDataSetWriterMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UadpDataSetWriterMessageDataType")
		}

		if err := WriteSimpleEnumField[UadpDataSetMessageContentMask](ctx, "dataSetMessageContentMask", "UadpDataSetMessageContentMask", m.GetDataSetMessageContentMask(), WriteEnum[UadpDataSetMessageContentMask, uint32](UadpDataSetMessageContentMask.GetValue, UadpDataSetMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetMessageContentMask' field")
		}

		if err := WriteSimpleField[uint16](ctx, "configuredSize", m.GetConfiguredSize(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'configuredSize' field")
		}

		if err := WriteSimpleField[uint16](ctx, "networkMessageNumber", m.GetNetworkMessageNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'networkMessageNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "dataSetOffset", m.GetDataSetOffset(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'dataSetOffset' field")
		}

		if popErr := writeBuffer.PopContext("UadpDataSetWriterMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UadpDataSetWriterMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UadpDataSetWriterMessageDataType) IsUadpDataSetWriterMessageDataType() {}

func (m *_UadpDataSetWriterMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UadpDataSetWriterMessageDataType) deepCopy() *_UadpDataSetWriterMessageDataType {
	if m == nil {
		return nil
	}
	_UadpDataSetWriterMessageDataTypeCopy := &_UadpDataSetWriterMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.DataSetMessageContentMask,
		m.ConfiguredSize,
		m.NetworkMessageNumber,
		m.DataSetOffset,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UadpDataSetWriterMessageDataTypeCopy
}

func (m *_UadpDataSetWriterMessageDataType) String() string {
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
