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

// ReaderGroupTransportDataType is the corresponding interface of ReaderGroupTransportDataType
type ReaderGroupTransportDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsReaderGroupTransportDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsReaderGroupTransportDataType()
	// CreateBuilder creates a ReaderGroupTransportDataTypeBuilder
	CreateReaderGroupTransportDataTypeBuilder() ReaderGroupTransportDataTypeBuilder
}

// _ReaderGroupTransportDataType is the data-structure of this message
type _ReaderGroupTransportDataType struct {
	ExtensionObjectDefinitionContract
}

var _ ReaderGroupTransportDataType = (*_ReaderGroupTransportDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ReaderGroupTransportDataType)(nil)

// NewReaderGroupTransportDataType factory function for _ReaderGroupTransportDataType
func NewReaderGroupTransportDataType() *_ReaderGroupTransportDataType {
	_result := &_ReaderGroupTransportDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ReaderGroupTransportDataTypeBuilder is a builder for ReaderGroupTransportDataType
type ReaderGroupTransportDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ReaderGroupTransportDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the ReaderGroupTransportDataType or returns an error if something is wrong
	Build() (ReaderGroupTransportDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ReaderGroupTransportDataType
}

// NewReaderGroupTransportDataTypeBuilder() creates a ReaderGroupTransportDataTypeBuilder
func NewReaderGroupTransportDataTypeBuilder() ReaderGroupTransportDataTypeBuilder {
	return &_ReaderGroupTransportDataTypeBuilder{_ReaderGroupTransportDataType: new(_ReaderGroupTransportDataType)}
}

type _ReaderGroupTransportDataTypeBuilder struct {
	*_ReaderGroupTransportDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ReaderGroupTransportDataTypeBuilder) = (*_ReaderGroupTransportDataTypeBuilder)(nil)

func (b *_ReaderGroupTransportDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
	contract.(*_ExtensionObjectDefinition)._SubType = b._ReaderGroupTransportDataType
}

func (b *_ReaderGroupTransportDataTypeBuilder) WithMandatoryFields() ReaderGroupTransportDataTypeBuilder {
	return b
}

func (b *_ReaderGroupTransportDataTypeBuilder) Build() (ReaderGroupTransportDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ReaderGroupTransportDataType.deepCopy(), nil
}

func (b *_ReaderGroupTransportDataTypeBuilder) MustBuild() ReaderGroupTransportDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ReaderGroupTransportDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_ReaderGroupTransportDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ReaderGroupTransportDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateReaderGroupTransportDataTypeBuilder().(*_ReaderGroupTransportDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateReaderGroupTransportDataTypeBuilder creates a ReaderGroupTransportDataTypeBuilder
func (b *_ReaderGroupTransportDataType) CreateReaderGroupTransportDataTypeBuilder() ReaderGroupTransportDataTypeBuilder {
	if b == nil {
		return NewReaderGroupTransportDataTypeBuilder()
	}
	return &_ReaderGroupTransportDataTypeBuilder{_ReaderGroupTransportDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReaderGroupTransportDataType) GetExtensionId() int32 {
	return int32(15623)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReaderGroupTransportDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastReaderGroupTransportDataType(structType any) ReaderGroupTransportDataType {
	if casted, ok := structType.(ReaderGroupTransportDataType); ok {
		return casted
	}
	if casted, ok := structType.(*ReaderGroupTransportDataType); ok {
		return *casted
	}
	return nil
}

func (m *_ReaderGroupTransportDataType) GetTypeName() string {
	return "ReaderGroupTransportDataType"
}

func (m *_ReaderGroupTransportDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ReaderGroupTransportDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ReaderGroupTransportDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__readerGroupTransportDataType ReaderGroupTransportDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ReaderGroupTransportDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReaderGroupTransportDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ReaderGroupTransportDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReaderGroupTransportDataType")
	}

	return m, nil
}

func (m *_ReaderGroupTransportDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReaderGroupTransportDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReaderGroupTransportDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReaderGroupTransportDataType")
		}

		if popErr := writeBuffer.PopContext("ReaderGroupTransportDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReaderGroupTransportDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReaderGroupTransportDataType) IsReaderGroupTransportDataType() {}

func (m *_ReaderGroupTransportDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ReaderGroupTransportDataType) deepCopy() *_ReaderGroupTransportDataType {
	if m == nil {
		return nil
	}
	_ReaderGroupTransportDataTypeCopy := &_ReaderGroupTransportDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	_ReaderGroupTransportDataTypeCopy.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ReaderGroupTransportDataTypeCopy
}

func (m *_ReaderGroupTransportDataType) String() string {
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
