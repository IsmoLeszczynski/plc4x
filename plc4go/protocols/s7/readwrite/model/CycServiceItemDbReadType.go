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

// CycServiceItemDbReadType is the corresponding interface of CycServiceItemDbReadType
type CycServiceItemDbReadType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CycServiceItemType
	// GetNumberOfAreas returns NumberOfAreas (property field)
	GetNumberOfAreas() uint8
	// GetItems returns Items (property field)
	GetItems() []SubItem
	// IsCycServiceItemDbReadType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCycServiceItemDbReadType()
	// CreateBuilder creates a CycServiceItemDbReadTypeBuilder
	CreateCycServiceItemDbReadTypeBuilder() CycServiceItemDbReadTypeBuilder
}

// _CycServiceItemDbReadType is the data-structure of this message
type _CycServiceItemDbReadType struct {
	CycServiceItemTypeContract
	NumberOfAreas uint8
	Items         []SubItem
}

var _ CycServiceItemDbReadType = (*_CycServiceItemDbReadType)(nil)
var _ CycServiceItemTypeRequirements = (*_CycServiceItemDbReadType)(nil)

// NewCycServiceItemDbReadType factory function for _CycServiceItemDbReadType
func NewCycServiceItemDbReadType(byteLength uint8, syntaxId uint8, numberOfAreas uint8, items []SubItem) *_CycServiceItemDbReadType {
	_result := &_CycServiceItemDbReadType{
		CycServiceItemTypeContract: NewCycServiceItemType(byteLength, syntaxId),
		NumberOfAreas:              numberOfAreas,
		Items:                      items,
	}
	_result.CycServiceItemTypeContract.(*_CycServiceItemType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CycServiceItemDbReadTypeBuilder is a builder for CycServiceItemDbReadType
type CycServiceItemDbReadTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numberOfAreas uint8, items []SubItem) CycServiceItemDbReadTypeBuilder
	// WithNumberOfAreas adds NumberOfAreas (property field)
	WithNumberOfAreas(uint8) CycServiceItemDbReadTypeBuilder
	// WithItems adds Items (property field)
	WithItems(...SubItem) CycServiceItemDbReadTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CycServiceItemTypeBuilder
	// Build builds the CycServiceItemDbReadType or returns an error if something is wrong
	Build() (CycServiceItemDbReadType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CycServiceItemDbReadType
}

// NewCycServiceItemDbReadTypeBuilder() creates a CycServiceItemDbReadTypeBuilder
func NewCycServiceItemDbReadTypeBuilder() CycServiceItemDbReadTypeBuilder {
	return &_CycServiceItemDbReadTypeBuilder{_CycServiceItemDbReadType: new(_CycServiceItemDbReadType)}
}

type _CycServiceItemDbReadTypeBuilder struct {
	*_CycServiceItemDbReadType

	parentBuilder *_CycServiceItemTypeBuilder

	err *utils.MultiError
}

var _ (CycServiceItemDbReadTypeBuilder) = (*_CycServiceItemDbReadTypeBuilder)(nil)

func (b *_CycServiceItemDbReadTypeBuilder) setParent(contract CycServiceItemTypeContract) {
	b.CycServiceItemTypeContract = contract
}

func (b *_CycServiceItemDbReadTypeBuilder) WithMandatoryFields(numberOfAreas uint8, items []SubItem) CycServiceItemDbReadTypeBuilder {
	return b.WithNumberOfAreas(numberOfAreas).WithItems(items...)
}

func (b *_CycServiceItemDbReadTypeBuilder) WithNumberOfAreas(numberOfAreas uint8) CycServiceItemDbReadTypeBuilder {
	b.NumberOfAreas = numberOfAreas
	return b
}

func (b *_CycServiceItemDbReadTypeBuilder) WithItems(items ...SubItem) CycServiceItemDbReadTypeBuilder {
	b.Items = items
	return b
}

func (b *_CycServiceItemDbReadTypeBuilder) Build() (CycServiceItemDbReadType, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CycServiceItemDbReadType.deepCopy(), nil
}

func (b *_CycServiceItemDbReadTypeBuilder) MustBuild() CycServiceItemDbReadType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CycServiceItemDbReadTypeBuilder) Done() CycServiceItemTypeBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCycServiceItemTypeBuilder().(*_CycServiceItemTypeBuilder)
	}
	return b.parentBuilder
}

func (b *_CycServiceItemDbReadTypeBuilder) buildForCycServiceItemType() (CycServiceItemType, error) {
	return b.Build()
}

func (b *_CycServiceItemDbReadTypeBuilder) DeepCopy() any {
	_copy := b.CreateCycServiceItemDbReadTypeBuilder().(*_CycServiceItemDbReadTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCycServiceItemDbReadTypeBuilder creates a CycServiceItemDbReadTypeBuilder
func (b *_CycServiceItemDbReadType) CreateCycServiceItemDbReadTypeBuilder() CycServiceItemDbReadTypeBuilder {
	if b == nil {
		return NewCycServiceItemDbReadTypeBuilder()
	}
	return &_CycServiceItemDbReadTypeBuilder{_CycServiceItemDbReadType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CycServiceItemDbReadType) GetParent() CycServiceItemTypeContract {
	return m.CycServiceItemTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemDbReadType) GetNumberOfAreas() uint8 {
	return m.NumberOfAreas
}

func (m *_CycServiceItemDbReadType) GetItems() []SubItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCycServiceItemDbReadType(structType any) CycServiceItemDbReadType {
	if casted, ok := structType.(CycServiceItemDbReadType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemDbReadType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemDbReadType) GetTypeName() string {
	return "CycServiceItemDbReadType"
}

func (m *_CycServiceItemDbReadType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CycServiceItemTypeContract.(*_CycServiceItemType).getLengthInBits(ctx))

	// Simple field (numberOfAreas)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CycServiceItemDbReadType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CycServiceItemDbReadType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CycServiceItemType) (__cycServiceItemDbReadType CycServiceItemDbReadType, err error) {
	m.CycServiceItemTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CycServiceItemDbReadType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemDbReadType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numberOfAreas, err := ReadSimpleField(ctx, "numberOfAreas", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfAreas' field"))
	}
	m.NumberOfAreas = numberOfAreas

	items, err := ReadCountArrayField[SubItem](ctx, "items", ReadComplex[SubItem](SubItemParseWithBuffer, readBuffer), uint64(numberOfAreas))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("CycServiceItemDbReadType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemDbReadType")
	}

	return m, nil
}

func (m *_CycServiceItemDbReadType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CycServiceItemDbReadType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CycServiceItemDbReadType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CycServiceItemDbReadType")
		}

		if err := WriteSimpleField[uint8](ctx, "numberOfAreas", m.GetNumberOfAreas(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfAreas' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("CycServiceItemDbReadType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CycServiceItemDbReadType")
		}
		return nil
	}
	return m.CycServiceItemTypeContract.(*_CycServiceItemType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CycServiceItemDbReadType) IsCycServiceItemDbReadType() {}

func (m *_CycServiceItemDbReadType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CycServiceItemDbReadType) deepCopy() *_CycServiceItemDbReadType {
	if m == nil {
		return nil
	}
	_CycServiceItemDbReadTypeCopy := &_CycServiceItemDbReadType{
		m.CycServiceItemTypeContract.(*_CycServiceItemType).deepCopy(),
		m.NumberOfAreas,
		utils.DeepCopySlice[SubItem, SubItem](m.Items),
	}
	m.CycServiceItemTypeContract.(*_CycServiceItemType)._SubType = m
	return _CycServiceItemDbReadTypeCopy
}

func (m *_CycServiceItemDbReadType) String() string {
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
