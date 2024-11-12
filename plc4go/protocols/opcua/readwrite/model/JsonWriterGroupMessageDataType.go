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

// JsonWriterGroupMessageDataType is the corresponding interface of JsonWriterGroupMessageDataType
type JsonWriterGroupMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNetworkMessageContentMask returns NetworkMessageContentMask (property field)
	GetNetworkMessageContentMask() JsonNetworkMessageContentMask
	// IsJsonWriterGroupMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsJsonWriterGroupMessageDataType()
	// CreateBuilder creates a JsonWriterGroupMessageDataTypeBuilder
	CreateJsonWriterGroupMessageDataTypeBuilder() JsonWriterGroupMessageDataTypeBuilder
}

// _JsonWriterGroupMessageDataType is the data-structure of this message
type _JsonWriterGroupMessageDataType struct {
	ExtensionObjectDefinitionContract
	NetworkMessageContentMask JsonNetworkMessageContentMask
}

var _ JsonWriterGroupMessageDataType = (*_JsonWriterGroupMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_JsonWriterGroupMessageDataType)(nil)

// NewJsonWriterGroupMessageDataType factory function for _JsonWriterGroupMessageDataType
func NewJsonWriterGroupMessageDataType(networkMessageContentMask JsonNetworkMessageContentMask) *_JsonWriterGroupMessageDataType {
	_result := &_JsonWriterGroupMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NetworkMessageContentMask:         networkMessageContentMask,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// JsonWriterGroupMessageDataTypeBuilder is a builder for JsonWriterGroupMessageDataType
type JsonWriterGroupMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(networkMessageContentMask JsonNetworkMessageContentMask) JsonWriterGroupMessageDataTypeBuilder
	// WithNetworkMessageContentMask adds NetworkMessageContentMask (property field)
	WithNetworkMessageContentMask(JsonNetworkMessageContentMask) JsonWriterGroupMessageDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the JsonWriterGroupMessageDataType or returns an error if something is wrong
	Build() (JsonWriterGroupMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() JsonWriterGroupMessageDataType
}

// NewJsonWriterGroupMessageDataTypeBuilder() creates a JsonWriterGroupMessageDataTypeBuilder
func NewJsonWriterGroupMessageDataTypeBuilder() JsonWriterGroupMessageDataTypeBuilder {
	return &_JsonWriterGroupMessageDataTypeBuilder{_JsonWriterGroupMessageDataType: new(_JsonWriterGroupMessageDataType)}
}

type _JsonWriterGroupMessageDataTypeBuilder struct {
	*_JsonWriterGroupMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (JsonWriterGroupMessageDataTypeBuilder) = (*_JsonWriterGroupMessageDataTypeBuilder)(nil)

func (b *_JsonWriterGroupMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) WithMandatoryFields(networkMessageContentMask JsonNetworkMessageContentMask) JsonWriterGroupMessageDataTypeBuilder {
	return b.WithNetworkMessageContentMask(networkMessageContentMask)
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) WithNetworkMessageContentMask(networkMessageContentMask JsonNetworkMessageContentMask) JsonWriterGroupMessageDataTypeBuilder {
	b.NetworkMessageContentMask = networkMessageContentMask
	return b
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) Build() (JsonWriterGroupMessageDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._JsonWriterGroupMessageDataType.deepCopy(), nil
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) MustBuild() JsonWriterGroupMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_JsonWriterGroupMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateJsonWriterGroupMessageDataTypeBuilder().(*_JsonWriterGroupMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateJsonWriterGroupMessageDataTypeBuilder creates a JsonWriterGroupMessageDataTypeBuilder
func (b *_JsonWriterGroupMessageDataType) CreateJsonWriterGroupMessageDataTypeBuilder() JsonWriterGroupMessageDataTypeBuilder {
	if b == nil {
		return NewJsonWriterGroupMessageDataTypeBuilder()
	}
	return &_JsonWriterGroupMessageDataTypeBuilder{_JsonWriterGroupMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_JsonWriterGroupMessageDataType) GetExtensionId() int32 {
	return int32(15659)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_JsonWriterGroupMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_JsonWriterGroupMessageDataType) GetNetworkMessageContentMask() JsonNetworkMessageContentMask {
	return m.NetworkMessageContentMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastJsonWriterGroupMessageDataType(structType any) JsonWriterGroupMessageDataType {
	if casted, ok := structType.(JsonWriterGroupMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*JsonWriterGroupMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_JsonWriterGroupMessageDataType) GetTypeName() string {
	return "JsonWriterGroupMessageDataType"
}

func (m *_JsonWriterGroupMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (networkMessageContentMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_JsonWriterGroupMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_JsonWriterGroupMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__jsonWriterGroupMessageDataType JsonWriterGroupMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("JsonWriterGroupMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for JsonWriterGroupMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	networkMessageContentMask, err := ReadEnumField[JsonNetworkMessageContentMask](ctx, "networkMessageContentMask", "JsonNetworkMessageContentMask", ReadEnum(JsonNetworkMessageContentMaskByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkMessageContentMask' field"))
	}
	m.NetworkMessageContentMask = networkMessageContentMask

	if closeErr := readBuffer.CloseContext("JsonWriterGroupMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for JsonWriterGroupMessageDataType")
	}

	return m, nil
}

func (m *_JsonWriterGroupMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_JsonWriterGroupMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("JsonWriterGroupMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for JsonWriterGroupMessageDataType")
		}

		if err := WriteSimpleEnumField[JsonNetworkMessageContentMask](ctx, "networkMessageContentMask", "JsonNetworkMessageContentMask", m.GetNetworkMessageContentMask(), WriteEnum[JsonNetworkMessageContentMask, uint32](JsonNetworkMessageContentMask.GetValue, JsonNetworkMessageContentMask.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'networkMessageContentMask' field")
		}

		if popErr := writeBuffer.PopContext("JsonWriterGroupMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for JsonWriterGroupMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_JsonWriterGroupMessageDataType) IsJsonWriterGroupMessageDataType() {}

func (m *_JsonWriterGroupMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_JsonWriterGroupMessageDataType) deepCopy() *_JsonWriterGroupMessageDataType {
	if m == nil {
		return nil
	}
	_JsonWriterGroupMessageDataTypeCopy := &_JsonWriterGroupMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.NetworkMessageContentMask,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _JsonWriterGroupMessageDataTypeCopy
}

func (m *_JsonWriterGroupMessageDataType) String() string {
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
