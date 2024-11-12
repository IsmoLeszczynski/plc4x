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

// TransmitQosDataType is the corresponding interface of TransmitQosDataType
type TransmitQosDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// IsTransmitQosDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTransmitQosDataType()
	// CreateBuilder creates a TransmitQosDataTypeBuilder
	CreateTransmitQosDataTypeBuilder() TransmitQosDataTypeBuilder
}

// _TransmitQosDataType is the data-structure of this message
type _TransmitQosDataType struct {
	ExtensionObjectDefinitionContract
}

var _ TransmitQosDataType = (*_TransmitQosDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_TransmitQosDataType)(nil)

// NewTransmitQosDataType factory function for _TransmitQosDataType
func NewTransmitQosDataType() *_TransmitQosDataType {
	_result := &_TransmitQosDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TransmitQosDataTypeBuilder is a builder for TransmitQosDataType
type TransmitQosDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() TransmitQosDataTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the TransmitQosDataType or returns an error if something is wrong
	Build() (TransmitQosDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TransmitQosDataType
}

// NewTransmitQosDataTypeBuilder() creates a TransmitQosDataTypeBuilder
func NewTransmitQosDataTypeBuilder() TransmitQosDataTypeBuilder {
	return &_TransmitQosDataTypeBuilder{_TransmitQosDataType: new(_TransmitQosDataType)}
}

type _TransmitQosDataTypeBuilder struct {
	*_TransmitQosDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (TransmitQosDataTypeBuilder) = (*_TransmitQosDataTypeBuilder)(nil)

func (b *_TransmitQosDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_TransmitQosDataTypeBuilder) WithMandatoryFields() TransmitQosDataTypeBuilder {
	return b
}

func (b *_TransmitQosDataTypeBuilder) Build() (TransmitQosDataType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TransmitQosDataType.deepCopy(), nil
}

func (b *_TransmitQosDataTypeBuilder) MustBuild() TransmitQosDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TransmitQosDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_TransmitQosDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_TransmitQosDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateTransmitQosDataTypeBuilder().(*_TransmitQosDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTransmitQosDataTypeBuilder creates a TransmitQosDataTypeBuilder
func (b *_TransmitQosDataType) CreateTransmitQosDataTypeBuilder() TransmitQosDataTypeBuilder {
	if b == nil {
		return NewTransmitQosDataTypeBuilder()
	}
	return &_TransmitQosDataTypeBuilder{_TransmitQosDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransmitQosDataType) GetExtensionId() int32 {
	return int32(23606)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransmitQosDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

// Deprecated: use the interface for direct cast
func CastTransmitQosDataType(structType any) TransmitQosDataType {
	if casted, ok := structType.(TransmitQosDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TransmitQosDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TransmitQosDataType) GetTypeName() string {
	return "TransmitQosDataType"
}

func (m *_TransmitQosDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_TransmitQosDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TransmitQosDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__transmitQosDataType TransmitQosDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TransmitQosDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransmitQosDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TransmitQosDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransmitQosDataType")
	}

	return m, nil
}

func (m *_TransmitQosDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransmitQosDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransmitQosDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransmitQosDataType")
		}

		if popErr := writeBuffer.PopContext("TransmitQosDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransmitQosDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TransmitQosDataType) IsTransmitQosDataType() {}

func (m *_TransmitQosDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TransmitQosDataType) deepCopy() *_TransmitQosDataType {
	if m == nil {
		return nil
	}
	_TransmitQosDataTypeCopy := &_TransmitQosDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _TransmitQosDataTypeCopy
}

func (m *_TransmitQosDataType) String() string {
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
