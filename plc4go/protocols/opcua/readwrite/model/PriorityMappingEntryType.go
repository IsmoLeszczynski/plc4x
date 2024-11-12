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

// PriorityMappingEntryType is the corresponding interface of PriorityMappingEntryType
type PriorityMappingEntryType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetMappingUri returns MappingUri (property field)
	GetMappingUri() PascalString
	// GetPriorityLabel returns PriorityLabel (property field)
	GetPriorityLabel() PascalString
	// GetPriorityValue_PCP returns PriorityValue_PCP (property field)
	GetPriorityValue_PCP() uint8
	// GetPriorityValue_DSCP returns PriorityValue_DSCP (property field)
	GetPriorityValue_DSCP() uint32
	// IsPriorityMappingEntryType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPriorityMappingEntryType()
	// CreateBuilder creates a PriorityMappingEntryTypeBuilder
	CreatePriorityMappingEntryTypeBuilder() PriorityMappingEntryTypeBuilder
}

// _PriorityMappingEntryType is the data-structure of this message
type _PriorityMappingEntryType struct {
	ExtensionObjectDefinitionContract
	MappingUri         PascalString
	PriorityLabel      PascalString
	PriorityValue_PCP  uint8
	PriorityValue_DSCP uint32
}

var _ PriorityMappingEntryType = (*_PriorityMappingEntryType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PriorityMappingEntryType)(nil)

// NewPriorityMappingEntryType factory function for _PriorityMappingEntryType
func NewPriorityMappingEntryType(mappingUri PascalString, priorityLabel PascalString, priorityValue_PCP uint8, priorityValue_DSCP uint32) *_PriorityMappingEntryType {
	if mappingUri == nil {
		panic("mappingUri of type PascalString for PriorityMappingEntryType must not be nil")
	}
	if priorityLabel == nil {
		panic("priorityLabel of type PascalString for PriorityMappingEntryType must not be nil")
	}
	_result := &_PriorityMappingEntryType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		MappingUri:                        mappingUri,
		PriorityLabel:                     priorityLabel,
		PriorityValue_PCP:                 priorityValue_PCP,
		PriorityValue_DSCP:                priorityValue_DSCP,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// PriorityMappingEntryTypeBuilder is a builder for PriorityMappingEntryType
type PriorityMappingEntryTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(mappingUri PascalString, priorityLabel PascalString, priorityValue_PCP uint8, priorityValue_DSCP uint32) PriorityMappingEntryTypeBuilder
	// WithMappingUri adds MappingUri (property field)
	WithMappingUri(PascalString) PriorityMappingEntryTypeBuilder
	// WithMappingUriBuilder adds MappingUri (property field) which is build by the builder
	WithMappingUriBuilder(func(PascalStringBuilder) PascalStringBuilder) PriorityMappingEntryTypeBuilder
	// WithPriorityLabel adds PriorityLabel (property field)
	WithPriorityLabel(PascalString) PriorityMappingEntryTypeBuilder
	// WithPriorityLabelBuilder adds PriorityLabel (property field) which is build by the builder
	WithPriorityLabelBuilder(func(PascalStringBuilder) PascalStringBuilder) PriorityMappingEntryTypeBuilder
	// WithPriorityValue_PCP adds PriorityValue_PCP (property field)
	WithPriorityValue_PCP(uint8) PriorityMappingEntryTypeBuilder
	// WithPriorityValue_DSCP adds PriorityValue_DSCP (property field)
	WithPriorityValue_DSCP(uint32) PriorityMappingEntryTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the PriorityMappingEntryType or returns an error if something is wrong
	Build() (PriorityMappingEntryType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() PriorityMappingEntryType
}

// NewPriorityMappingEntryTypeBuilder() creates a PriorityMappingEntryTypeBuilder
func NewPriorityMappingEntryTypeBuilder() PriorityMappingEntryTypeBuilder {
	return &_PriorityMappingEntryTypeBuilder{_PriorityMappingEntryType: new(_PriorityMappingEntryType)}
}

type _PriorityMappingEntryTypeBuilder struct {
	*_PriorityMappingEntryType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (PriorityMappingEntryTypeBuilder) = (*_PriorityMappingEntryTypeBuilder)(nil)

func (b *_PriorityMappingEntryTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_PriorityMappingEntryTypeBuilder) WithMandatoryFields(mappingUri PascalString, priorityLabel PascalString, priorityValue_PCP uint8, priorityValue_DSCP uint32) PriorityMappingEntryTypeBuilder {
	return b.WithMappingUri(mappingUri).WithPriorityLabel(priorityLabel).WithPriorityValue_PCP(priorityValue_PCP).WithPriorityValue_DSCP(priorityValue_DSCP)
}

func (b *_PriorityMappingEntryTypeBuilder) WithMappingUri(mappingUri PascalString) PriorityMappingEntryTypeBuilder {
	b.MappingUri = mappingUri
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) WithMappingUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) PriorityMappingEntryTypeBuilder {
	builder := builderSupplier(b.MappingUri.CreatePascalStringBuilder())
	var err error
	b.MappingUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) WithPriorityLabel(priorityLabel PascalString) PriorityMappingEntryTypeBuilder {
	b.PriorityLabel = priorityLabel
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) WithPriorityLabelBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) PriorityMappingEntryTypeBuilder {
	builder := builderSupplier(b.PriorityLabel.CreatePascalStringBuilder())
	var err error
	b.PriorityLabel, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) WithPriorityValue_PCP(priorityValue_PCP uint8) PriorityMappingEntryTypeBuilder {
	b.PriorityValue_PCP = priorityValue_PCP
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) WithPriorityValue_DSCP(priorityValue_DSCP uint32) PriorityMappingEntryTypeBuilder {
	b.PriorityValue_DSCP = priorityValue_DSCP
	return b
}

func (b *_PriorityMappingEntryTypeBuilder) Build() (PriorityMappingEntryType, error) {
	if b.MappingUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'mappingUri' not set"))
	}
	if b.PriorityLabel == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'priorityLabel' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._PriorityMappingEntryType.deepCopy(), nil
}

func (b *_PriorityMappingEntryTypeBuilder) MustBuild() PriorityMappingEntryType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_PriorityMappingEntryTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_PriorityMappingEntryTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_PriorityMappingEntryTypeBuilder) DeepCopy() any {
	_copy := b.CreatePriorityMappingEntryTypeBuilder().(*_PriorityMappingEntryTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreatePriorityMappingEntryTypeBuilder creates a PriorityMappingEntryTypeBuilder
func (b *_PriorityMappingEntryType) CreatePriorityMappingEntryTypeBuilder() PriorityMappingEntryTypeBuilder {
	if b == nil {
		return NewPriorityMappingEntryTypeBuilder()
	}
	return &_PriorityMappingEntryTypeBuilder{_PriorityMappingEntryType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PriorityMappingEntryType) GetExtensionId() int32 {
	return int32(25222)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PriorityMappingEntryType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PriorityMappingEntryType) GetMappingUri() PascalString {
	return m.MappingUri
}

func (m *_PriorityMappingEntryType) GetPriorityLabel() PascalString {
	return m.PriorityLabel
}

func (m *_PriorityMappingEntryType) GetPriorityValue_PCP() uint8 {
	return m.PriorityValue_PCP
}

func (m *_PriorityMappingEntryType) GetPriorityValue_DSCP() uint32 {
	return m.PriorityValue_DSCP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastPriorityMappingEntryType(structType any) PriorityMappingEntryType {
	if casted, ok := structType.(PriorityMappingEntryType); ok {
		return casted
	}
	if casted, ok := structType.(*PriorityMappingEntryType); ok {
		return *casted
	}
	return nil
}

func (m *_PriorityMappingEntryType) GetTypeName() string {
	return "PriorityMappingEntryType"
}

func (m *_PriorityMappingEntryType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (mappingUri)
	lengthInBits += m.MappingUri.GetLengthInBits(ctx)

	// Simple field (priorityLabel)
	lengthInBits += m.PriorityLabel.GetLengthInBits(ctx)

	// Simple field (priorityValue_PCP)
	lengthInBits += 8

	// Simple field (priorityValue_DSCP)
	lengthInBits += 32

	return lengthInBits
}

func (m *_PriorityMappingEntryType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PriorityMappingEntryType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__priorityMappingEntryType PriorityMappingEntryType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PriorityMappingEntryType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PriorityMappingEntryType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	mappingUri, err := ReadSimpleField[PascalString](ctx, "mappingUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'mappingUri' field"))
	}
	m.MappingUri = mappingUri

	priorityLabel, err := ReadSimpleField[PascalString](ctx, "priorityLabel", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityLabel' field"))
	}
	m.PriorityLabel = priorityLabel

	priorityValue_PCP, err := ReadSimpleField(ctx, "priorityValue_PCP", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityValue_PCP' field"))
	}
	m.PriorityValue_PCP = priorityValue_PCP

	priorityValue_DSCP, err := ReadSimpleField(ctx, "priorityValue_DSCP", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityValue_DSCP' field"))
	}
	m.PriorityValue_DSCP = priorityValue_DSCP

	if closeErr := readBuffer.CloseContext("PriorityMappingEntryType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PriorityMappingEntryType")
	}

	return m, nil
}

func (m *_PriorityMappingEntryType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PriorityMappingEntryType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PriorityMappingEntryType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PriorityMappingEntryType")
		}

		if err := WriteSimpleField[PascalString](ctx, "mappingUri", m.GetMappingUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'mappingUri' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "priorityLabel", m.GetPriorityLabel(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityLabel' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priorityValue_PCP", m.GetPriorityValue_PCP(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityValue_PCP' field")
		}

		if err := WriteSimpleField[uint32](ctx, "priorityValue_DSCP", m.GetPriorityValue_DSCP(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityValue_DSCP' field")
		}

		if popErr := writeBuffer.PopContext("PriorityMappingEntryType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PriorityMappingEntryType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PriorityMappingEntryType) IsPriorityMappingEntryType() {}

func (m *_PriorityMappingEntryType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_PriorityMappingEntryType) deepCopy() *_PriorityMappingEntryType {
	if m == nil {
		return nil
	}
	_PriorityMappingEntryTypeCopy := &_PriorityMappingEntryType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.MappingUri),
		utils.DeepCopy[PascalString](m.PriorityLabel),
		m.PriorityValue_PCP,
		m.PriorityValue_DSCP,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _PriorityMappingEntryTypeCopy
}

func (m *_PriorityMappingEntryType) String() string {
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
