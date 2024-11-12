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

// ChangeListAddError is the corresponding interface of ChangeListAddError
type ChangeListAddError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
	// IsChangeListAddError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsChangeListAddError()
	// CreateBuilder creates a ChangeListAddErrorBuilder
	CreateChangeListAddErrorBuilder() ChangeListAddErrorBuilder
}

// _ChangeListAddError is the data-structure of this message
type _ChangeListAddError struct {
	BACnetErrorContract
	ErrorType                ErrorEnclosed
	FirstFailedElementNumber BACnetContextTagUnsignedInteger
}

var _ ChangeListAddError = (*_ChangeListAddError)(nil)
var _ BACnetErrorRequirements = (*_ChangeListAddError)(nil)

// NewChangeListAddError factory function for _ChangeListAddError
func NewChangeListAddError(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) *_ChangeListAddError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for ChangeListAddError must not be nil")
	}
	if firstFailedElementNumber == nil {
		panic("firstFailedElementNumber of type BACnetContextTagUnsignedInteger for ChangeListAddError must not be nil")
	}
	_result := &_ChangeListAddError{
		BACnetErrorContract:      NewBACnetError(),
		ErrorType:                errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ChangeListAddErrorBuilder is a builder for ChangeListAddError
type ChangeListAddErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListAddErrorBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) ChangeListAddErrorBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) ChangeListAddErrorBuilder
	// WithFirstFailedElementNumber adds FirstFailedElementNumber (property field)
	WithFirstFailedElementNumber(BACnetContextTagUnsignedInteger) ChangeListAddErrorBuilder
	// WithFirstFailedElementNumberBuilder adds FirstFailedElementNumber (property field) which is build by the builder
	WithFirstFailedElementNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) ChangeListAddErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetErrorBuilder
	// Build builds the ChangeListAddError or returns an error if something is wrong
	Build() (ChangeListAddError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ChangeListAddError
}

// NewChangeListAddErrorBuilder() creates a ChangeListAddErrorBuilder
func NewChangeListAddErrorBuilder() ChangeListAddErrorBuilder {
	return &_ChangeListAddErrorBuilder{_ChangeListAddError: new(_ChangeListAddError)}
}

type _ChangeListAddErrorBuilder struct {
	*_ChangeListAddError

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (ChangeListAddErrorBuilder) = (*_ChangeListAddErrorBuilder)(nil)

func (b *_ChangeListAddErrorBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
}

func (b *_ChangeListAddErrorBuilder) WithMandatoryFields(errorType ErrorEnclosed, firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListAddErrorBuilder {
	return b.WithErrorType(errorType).WithFirstFailedElementNumber(firstFailedElementNumber)
}

func (b *_ChangeListAddErrorBuilder) WithErrorType(errorType ErrorEnclosed) ChangeListAddErrorBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_ChangeListAddErrorBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) ChangeListAddErrorBuilder {
	builder := builderSupplier(b.ErrorType.CreateErrorEnclosedBuilder())
	var err error
	b.ErrorType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ErrorEnclosedBuilder failed"))
	}
	return b
}

func (b *_ChangeListAddErrorBuilder) WithFirstFailedElementNumber(firstFailedElementNumber BACnetContextTagUnsignedInteger) ChangeListAddErrorBuilder {
	b.FirstFailedElementNumber = firstFailedElementNumber
	return b
}

func (b *_ChangeListAddErrorBuilder) WithFirstFailedElementNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) ChangeListAddErrorBuilder {
	builder := builderSupplier(b.FirstFailedElementNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.FirstFailedElementNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_ChangeListAddErrorBuilder) Build() (ChangeListAddError, error) {
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.FirstFailedElementNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'firstFailedElementNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ChangeListAddError.deepCopy(), nil
}

func (b *_ChangeListAddErrorBuilder) MustBuild() ChangeListAddError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ChangeListAddErrorBuilder) Done() BACnetErrorBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetErrorBuilder().(*_BACnetErrorBuilder)
	}
	return b.parentBuilder
}

func (b *_ChangeListAddErrorBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_ChangeListAddErrorBuilder) DeepCopy() any {
	_copy := b.CreateChangeListAddErrorBuilder().(*_ChangeListAddErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateChangeListAddErrorBuilder creates a ChangeListAddErrorBuilder
func (b *_ChangeListAddError) CreateChangeListAddErrorBuilder() ChangeListAddErrorBuilder {
	if b == nil {
		return NewChangeListAddErrorBuilder()
	}
	return &_ChangeListAddErrorBuilder{_ChangeListAddError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChangeListAddError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ADD_LIST_ELEMENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChangeListAddError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChangeListAddError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ChangeListAddError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastChangeListAddError(structType any) ChangeListAddError {
	if casted, ok := structType.(ChangeListAddError); ok {
		return casted
	}
	if casted, ok := structType.(*ChangeListAddError); ok {
		return *casted
	}
	return nil
}

func (m *_ChangeListAddError) GetTypeName() string {
	return "ChangeListAddError"
}

func (m *_ChangeListAddError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ChangeListAddError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ChangeListAddError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__changeListAddError ChangeListAddError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChangeListAddError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChangeListAddError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	firstFailedElementNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'firstFailedElementNumber' field"))
	}
	m.FirstFailedElementNumber = firstFailedElementNumber

	if closeErr := readBuffer.CloseContext("ChangeListAddError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChangeListAddError")
	}

	return m, nil
}

func (m *_ChangeListAddError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChangeListAddError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChangeListAddError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChangeListAddError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "firstFailedElementNumber", m.GetFirstFailedElementNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'firstFailedElementNumber' field")
		}

		if popErr := writeBuffer.PopContext("ChangeListAddError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChangeListAddError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ChangeListAddError) IsChangeListAddError() {}

func (m *_ChangeListAddError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ChangeListAddError) deepCopy() *_ChangeListAddError {
	if m == nil {
		return nil
	}
	_ChangeListAddErrorCopy := &_ChangeListAddError{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		utils.DeepCopy[ErrorEnclosed](m.ErrorType),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.FirstFailedElementNumber),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _ChangeListAddErrorCopy
}

func (m *_ChangeListAddError) String() string {
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
