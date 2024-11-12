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

// HistoryEventFieldList is the corresponding interface of HistoryEventFieldList
type HistoryEventFieldList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetEventFields returns EventFields (property field)
	GetEventFields() []Variant
	// IsHistoryEventFieldList is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsHistoryEventFieldList()
	// CreateBuilder creates a HistoryEventFieldListBuilder
	CreateHistoryEventFieldListBuilder() HistoryEventFieldListBuilder
}

// _HistoryEventFieldList is the data-structure of this message
type _HistoryEventFieldList struct {
	ExtensionObjectDefinitionContract
	EventFields []Variant
}

var _ HistoryEventFieldList = (*_HistoryEventFieldList)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_HistoryEventFieldList)(nil)

// NewHistoryEventFieldList factory function for _HistoryEventFieldList
func NewHistoryEventFieldList(eventFields []Variant) *_HistoryEventFieldList {
	_result := &_HistoryEventFieldList{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		EventFields:                       eventFields,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// HistoryEventFieldListBuilder is a builder for HistoryEventFieldList
type HistoryEventFieldListBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventFields []Variant) HistoryEventFieldListBuilder
	// WithEventFields adds EventFields (property field)
	WithEventFields(...Variant) HistoryEventFieldListBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the HistoryEventFieldList or returns an error if something is wrong
	Build() (HistoryEventFieldList, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() HistoryEventFieldList
}

// NewHistoryEventFieldListBuilder() creates a HistoryEventFieldListBuilder
func NewHistoryEventFieldListBuilder() HistoryEventFieldListBuilder {
	return &_HistoryEventFieldListBuilder{_HistoryEventFieldList: new(_HistoryEventFieldList)}
}

type _HistoryEventFieldListBuilder struct {
	*_HistoryEventFieldList

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (HistoryEventFieldListBuilder) = (*_HistoryEventFieldListBuilder)(nil)

func (b *_HistoryEventFieldListBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_HistoryEventFieldListBuilder) WithMandatoryFields(eventFields []Variant) HistoryEventFieldListBuilder {
	return b.WithEventFields(eventFields...)
}

func (b *_HistoryEventFieldListBuilder) WithEventFields(eventFields ...Variant) HistoryEventFieldListBuilder {
	b.EventFields = eventFields
	return b
}

func (b *_HistoryEventFieldListBuilder) Build() (HistoryEventFieldList, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._HistoryEventFieldList.deepCopy(), nil
}

func (b *_HistoryEventFieldListBuilder) MustBuild() HistoryEventFieldList {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_HistoryEventFieldListBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_HistoryEventFieldListBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_HistoryEventFieldListBuilder) DeepCopy() any {
	_copy := b.CreateHistoryEventFieldListBuilder().(*_HistoryEventFieldListBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateHistoryEventFieldListBuilder creates a HistoryEventFieldListBuilder
func (b *_HistoryEventFieldList) CreateHistoryEventFieldListBuilder() HistoryEventFieldListBuilder {
	if b == nil {
		return NewHistoryEventFieldListBuilder()
	}
	return &_HistoryEventFieldListBuilder{_HistoryEventFieldList: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryEventFieldList) GetExtensionId() int32 {
	return int32(922)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryEventFieldList) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryEventFieldList) GetEventFields() []Variant {
	return m.EventFields
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastHistoryEventFieldList(structType any) HistoryEventFieldList {
	if casted, ok := structType.(HistoryEventFieldList); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryEventFieldList); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryEventFieldList) GetTypeName() string {
	return "HistoryEventFieldList"
}

func (m *_HistoryEventFieldList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Implicit Field (noOfEventFields)
	lengthInBits += 32

	// Array field
	if len(m.EventFields) > 0 {
		for _curItem, element := range m.EventFields {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.EventFields), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_HistoryEventFieldList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_HistoryEventFieldList) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__historyEventFieldList HistoryEventFieldList, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("HistoryEventFieldList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryEventFieldList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	noOfEventFields, err := ReadImplicitField[int32](ctx, "noOfEventFields", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfEventFields' field"))
	}
	_ = noOfEventFields

	eventFields, err := ReadCountArrayField[Variant](ctx, "eventFields", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfEventFields))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventFields' field"))
	}
	m.EventFields = eventFields

	if closeErr := readBuffer.CloseContext("HistoryEventFieldList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryEventFieldList")
	}

	return m, nil
}

func (m *_HistoryEventFieldList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryEventFieldList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryEventFieldList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryEventFieldList")
		}
		noOfEventFields := int32(utils.InlineIf(bool((m.GetEventFields()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetEventFields()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfEventFields", noOfEventFields, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfEventFields' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "eventFields", m.GetEventFields(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'eventFields' field")
		}

		if popErr := writeBuffer.PopContext("HistoryEventFieldList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryEventFieldList")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_HistoryEventFieldList) IsHistoryEventFieldList() {}

func (m *_HistoryEventFieldList) DeepCopy() any {
	return m.deepCopy()
}

func (m *_HistoryEventFieldList) deepCopy() *_HistoryEventFieldList {
	if m == nil {
		return nil
	}
	_HistoryEventFieldListCopy := &_HistoryEventFieldList{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopySlice[Variant, Variant](m.EventFields),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _HistoryEventFieldListCopy
}

func (m *_HistoryEventFieldList) String() string {
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
