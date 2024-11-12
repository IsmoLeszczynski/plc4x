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

// ErrorReportingSystemCategoryTypeInputUnits is the corresponding interface of ErrorReportingSystemCategoryTypeInputUnits
type ErrorReportingSystemCategoryTypeInputUnits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForInputUnits
	// IsErrorReportingSystemCategoryTypeInputUnits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsErrorReportingSystemCategoryTypeInputUnits()
	// CreateBuilder creates a ErrorReportingSystemCategoryTypeInputUnitsBuilder
	CreateErrorReportingSystemCategoryTypeInputUnitsBuilder() ErrorReportingSystemCategoryTypeInputUnitsBuilder
}

// _ErrorReportingSystemCategoryTypeInputUnits is the data-structure of this message
type _ErrorReportingSystemCategoryTypeInputUnits struct {
	ErrorReportingSystemCategoryTypeContract
	CategoryForType ErrorReportingSystemCategoryTypeForInputUnits
}

var _ ErrorReportingSystemCategoryTypeInputUnits = (*_ErrorReportingSystemCategoryTypeInputUnits)(nil)
var _ ErrorReportingSystemCategoryTypeRequirements = (*_ErrorReportingSystemCategoryTypeInputUnits)(nil)

// NewErrorReportingSystemCategoryTypeInputUnits factory function for _ErrorReportingSystemCategoryTypeInputUnits
func NewErrorReportingSystemCategoryTypeInputUnits(categoryForType ErrorReportingSystemCategoryTypeForInputUnits) *_ErrorReportingSystemCategoryTypeInputUnits {
	_result := &_ErrorReportingSystemCategoryTypeInputUnits{
		ErrorReportingSystemCategoryTypeContract: NewErrorReportingSystemCategoryType(),
		CategoryForType:                          categoryForType,
	}
	_result.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ErrorReportingSystemCategoryTypeInputUnitsBuilder is a builder for ErrorReportingSystemCategoryTypeInputUnits
type ErrorReportingSystemCategoryTypeInputUnitsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(categoryForType ErrorReportingSystemCategoryTypeForInputUnits) ErrorReportingSystemCategoryTypeInputUnitsBuilder
	// WithCategoryForType adds CategoryForType (property field)
	WithCategoryForType(ErrorReportingSystemCategoryTypeForInputUnits) ErrorReportingSystemCategoryTypeInputUnitsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ErrorReportingSystemCategoryTypeBuilder
	// Build builds the ErrorReportingSystemCategoryTypeInputUnits or returns an error if something is wrong
	Build() (ErrorReportingSystemCategoryTypeInputUnits, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ErrorReportingSystemCategoryTypeInputUnits
}

// NewErrorReportingSystemCategoryTypeInputUnitsBuilder() creates a ErrorReportingSystemCategoryTypeInputUnitsBuilder
func NewErrorReportingSystemCategoryTypeInputUnitsBuilder() ErrorReportingSystemCategoryTypeInputUnitsBuilder {
	return &_ErrorReportingSystemCategoryTypeInputUnitsBuilder{_ErrorReportingSystemCategoryTypeInputUnits: new(_ErrorReportingSystemCategoryTypeInputUnits)}
}

type _ErrorReportingSystemCategoryTypeInputUnitsBuilder struct {
	*_ErrorReportingSystemCategoryTypeInputUnits

	parentBuilder *_ErrorReportingSystemCategoryTypeBuilder

	err *utils.MultiError
}

var _ (ErrorReportingSystemCategoryTypeInputUnitsBuilder) = (*_ErrorReportingSystemCategoryTypeInputUnitsBuilder)(nil)

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) setParent(contract ErrorReportingSystemCategoryTypeContract) {
	b.ErrorReportingSystemCategoryTypeContract = contract
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) WithMandatoryFields(categoryForType ErrorReportingSystemCategoryTypeForInputUnits) ErrorReportingSystemCategoryTypeInputUnitsBuilder {
	return b.WithCategoryForType(categoryForType)
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) WithCategoryForType(categoryForType ErrorReportingSystemCategoryTypeForInputUnits) ErrorReportingSystemCategoryTypeInputUnitsBuilder {
	b.CategoryForType = categoryForType
	return b
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) Build() (ErrorReportingSystemCategoryTypeInputUnits, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ErrorReportingSystemCategoryTypeInputUnits.deepCopy(), nil
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) MustBuild() ErrorReportingSystemCategoryTypeInputUnits {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) Done() ErrorReportingSystemCategoryTypeBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewErrorReportingSystemCategoryTypeBuilder().(*_ErrorReportingSystemCategoryTypeBuilder)
	}
	return b.parentBuilder
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) buildForErrorReportingSystemCategoryType() (ErrorReportingSystemCategoryType, error) {
	return b.Build()
}

func (b *_ErrorReportingSystemCategoryTypeInputUnitsBuilder) DeepCopy() any {
	_copy := b.CreateErrorReportingSystemCategoryTypeInputUnitsBuilder().(*_ErrorReportingSystemCategoryTypeInputUnitsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateErrorReportingSystemCategoryTypeInputUnitsBuilder creates a ErrorReportingSystemCategoryTypeInputUnitsBuilder
func (b *_ErrorReportingSystemCategoryTypeInputUnits) CreateErrorReportingSystemCategoryTypeInputUnitsBuilder() ErrorReportingSystemCategoryTypeInputUnitsBuilder {
	if b == nil {
		return NewErrorReportingSystemCategoryTypeInputUnitsBuilder()
	}
	return &_ErrorReportingSystemCategoryTypeInputUnitsBuilder{_ErrorReportingSystemCategoryTypeInputUnits: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_INPUT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetParent() ErrorReportingSystemCategoryTypeContract {
	return m.ErrorReportingSystemCategoryTypeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetCategoryForType() ErrorReportingSystemCategoryTypeForInputUnits {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeInputUnits(structType any) ErrorReportingSystemCategoryTypeInputUnits {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeInputUnits); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeInputUnits); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeInputUnits"
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).getLengthInBits(ctx))

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ErrorReportingSystemCategoryType, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (__errorReportingSystemCategoryTypeInputUnits ErrorReportingSystemCategoryTypeInputUnits, err error) {
	m.ErrorReportingSystemCategoryTypeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeInputUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeInputUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryForType, err := ReadEnumField[ErrorReportingSystemCategoryTypeForInputUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForInputUnits", ReadEnum(ErrorReportingSystemCategoryTypeForInputUnitsByValue, ReadUnsignedByte(readBuffer, uint8(4))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryForType' field"))
	}
	m.CategoryForType = categoryForType

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeInputUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeInputUnits")
	}

	return m, nil
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeInputUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeInputUnits")
		}

		if err := WriteSimpleEnumField[ErrorReportingSystemCategoryTypeForInputUnits](ctx, "categoryForType", "ErrorReportingSystemCategoryTypeForInputUnits", m.GetCategoryForType(), WriteEnum[ErrorReportingSystemCategoryTypeForInputUnits, uint8](ErrorReportingSystemCategoryTypeForInputUnits.GetValue, ErrorReportingSystemCategoryTypeForInputUnits.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 4))); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeInputUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeInputUnits")
		}
		return nil
	}
	return m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) IsErrorReportingSystemCategoryTypeInputUnits() {
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) deepCopy() *_ErrorReportingSystemCategoryTypeInputUnits {
	if m == nil {
		return nil
	}
	_ErrorReportingSystemCategoryTypeInputUnitsCopy := &_ErrorReportingSystemCategoryTypeInputUnits{
		m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType).deepCopy(),
		m.CategoryForType,
	}
	m.ErrorReportingSystemCategoryTypeContract.(*_ErrorReportingSystemCategoryType)._SubType = m
	return _ErrorReportingSystemCategoryTypeInputUnitsCopy
}

func (m *_ErrorReportingSystemCategoryTypeInputUnits) String() string {
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
