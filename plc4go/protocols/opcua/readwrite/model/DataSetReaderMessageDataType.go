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

// DataSetReaderMessageDataType is the corresponding interface of DataSetReaderMessageDataType
type DataSetReaderMessageDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsDataSetReaderMessageDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDataSetReaderMessageDataType()
	// CreateBuilder creates a DataSetReaderMessageDataTypeBuilder
	CreateDataSetReaderMessageDataTypeBuilder() DataSetReaderMessageDataTypeBuilder
}

// _DataSetReaderMessageDataType is the data-structure of this message
type _DataSetReaderMessageDataType struct {
	ExtensionObjectDefinitionContract
}

var _ DataSetReaderMessageDataType = (*_DataSetReaderMessageDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DataSetReaderMessageDataType)(nil)

// NewDataSetReaderMessageDataType factory function for _DataSetReaderMessageDataType
func NewDataSetReaderMessageDataType() *_DataSetReaderMessageDataType {
	_result := &_DataSetReaderMessageDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DataSetReaderMessageDataTypeBuilder is a builder for DataSetReaderMessageDataType
type DataSetReaderMessageDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() DataSetReaderMessageDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DataSetReaderMessageDataType or returns an error if something is wrong
	Build() (DataSetReaderMessageDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DataSetReaderMessageDataType
}

// NewDataSetReaderMessageDataTypeBuilder() creates a DataSetReaderMessageDataTypeBuilder
func NewDataSetReaderMessageDataTypeBuilder() DataSetReaderMessageDataTypeBuilder {
	return &_DataSetReaderMessageDataTypeBuilder{_DataSetReaderMessageDataType: new(_DataSetReaderMessageDataType)}
}

type _DataSetReaderMessageDataTypeBuilder struct {
	*_DataSetReaderMessageDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DataSetReaderMessageDataTypeBuilder) = (*_DataSetReaderMessageDataTypeBuilder)(nil)

func (b *_DataSetReaderMessageDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DataSetReaderMessageDataTypeBuilder) WithMandatoryFields() DataSetReaderMessageDataTypeBuilder {
	return b
}

func (b *_DataSetReaderMessageDataTypeBuilder) Build() (DataSetReaderMessageDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DataSetReaderMessageDataType.deepCopy(), nil
}

func (b *_DataSetReaderMessageDataTypeBuilder) MustBuild() DataSetReaderMessageDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DataSetReaderMessageDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DataSetReaderMessageDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DataSetReaderMessageDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateDataSetReaderMessageDataTypeBuilder().(*_DataSetReaderMessageDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDataSetReaderMessageDataTypeBuilder creates a DataSetReaderMessageDataTypeBuilder
func (b *_DataSetReaderMessageDataType) CreateDataSetReaderMessageDataTypeBuilder() DataSetReaderMessageDataTypeBuilder {
	if b == nil {
		return NewDataSetReaderMessageDataTypeBuilder()
	}
	return &_DataSetReaderMessageDataTypeBuilder{_DataSetReaderMessageDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DataSetReaderMessageDataType) GetExtensionId() int32 {
	return int32(15631)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DataSetReaderMessageDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastDataSetReaderMessageDataType(structType any) DataSetReaderMessageDataType {
	if casted, ok := structType.(DataSetReaderMessageDataType); ok {
		return casted
	}
	if casted, ok := structType.(*DataSetReaderMessageDataType); ok {
		return *casted
	}
	return nil
}

func (m *_DataSetReaderMessageDataType) GetTypeName() string {
	return "DataSetReaderMessageDataType"
}

func (m *_DataSetReaderMessageDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_DataSetReaderMessageDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DataSetReaderMessageDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__dataSetReaderMessageDataType DataSetReaderMessageDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DataSetReaderMessageDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DataSetReaderMessageDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("DataSetReaderMessageDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DataSetReaderMessageDataType")
	}

	return m, nil
}

func (m *_DataSetReaderMessageDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DataSetReaderMessageDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DataSetReaderMessageDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DataSetReaderMessageDataType")
		}

		if popErr := writeBuffer.PopContext("DataSetReaderMessageDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DataSetReaderMessageDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DataSetReaderMessageDataType) IsDataSetReaderMessageDataType() {}

func (m *_DataSetReaderMessageDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DataSetReaderMessageDataType) deepCopy() *_DataSetReaderMessageDataType {
	if m == nil {
		return nil
	}
	_DataSetReaderMessageDataTypeCopy := &_DataSetReaderMessageDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DataSetReaderMessageDataTypeCopy
}

func (m *_DataSetReaderMessageDataType) String() string {
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
