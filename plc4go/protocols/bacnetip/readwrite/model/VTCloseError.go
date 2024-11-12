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

// VTCloseError is the corresponding interface of VTCloseError
type VTCloseError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetListOfVtSessionIdentifiers returns ListOfVtSessionIdentifiers (property field)
	GetListOfVtSessionIdentifiers() VTCloseErrorListOfVTSessionIdentifiers
	// IsVTCloseError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsVTCloseError()
	// CreateBuilder creates a VTCloseErrorBuilder
	CreateVTCloseErrorBuilder() VTCloseErrorBuilder
}

// _VTCloseError is the data-structure of this message
type _VTCloseError struct {
	BACnetErrorContract
	ErrorType                  ErrorEnclosed
	ListOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers
}

var _ VTCloseError = (*_VTCloseError)(nil)
var _ BACnetErrorRequirements = (*_VTCloseError)(nil)

// NewVTCloseError factory function for _VTCloseError
func NewVTCloseError(errorType ErrorEnclosed, listOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers) *_VTCloseError {
	if errorType == nil {
		panic("errorType of type ErrorEnclosed for VTCloseError must not be nil")
	}
	_result := &_VTCloseError{
		BACnetErrorContract:        NewBACnetError(),
		ErrorType:                  errorType,
		ListOfVtSessionIdentifiers: listOfVtSessionIdentifiers,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// VTCloseErrorBuilder is a builder for VTCloseError
type VTCloseErrorBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(errorType ErrorEnclosed) VTCloseErrorBuilder
	// WithErrorType adds ErrorType (property field)
	WithErrorType(ErrorEnclosed) VTCloseErrorBuilder
	// WithErrorTypeBuilder adds ErrorType (property field) which is build by the builder
	WithErrorTypeBuilder(func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) VTCloseErrorBuilder
	// WithListOfVtSessionIdentifiers adds ListOfVtSessionIdentifiers (property field)
	WithOptionalListOfVtSessionIdentifiers(VTCloseErrorListOfVTSessionIdentifiers) VTCloseErrorBuilder
	// WithOptionalListOfVtSessionIdentifiersBuilder adds ListOfVtSessionIdentifiers (property field) which is build by the builder
	WithOptionalListOfVtSessionIdentifiersBuilder(func(VTCloseErrorListOfVTSessionIdentifiersBuilder) VTCloseErrorListOfVTSessionIdentifiersBuilder) VTCloseErrorBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetErrorBuilder
	// Build builds the VTCloseError or returns an error if something is wrong
	Build() (VTCloseError, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() VTCloseError
}

// NewVTCloseErrorBuilder() creates a VTCloseErrorBuilder
func NewVTCloseErrorBuilder() VTCloseErrorBuilder {
	return &_VTCloseErrorBuilder{_VTCloseError: new(_VTCloseError)}
}

type _VTCloseErrorBuilder struct {
	*_VTCloseError

	parentBuilder *_BACnetErrorBuilder

	err *utils.MultiError
}

var _ (VTCloseErrorBuilder) = (*_VTCloseErrorBuilder)(nil)

func (b *_VTCloseErrorBuilder) setParent(contract BACnetErrorContract) {
	b.BACnetErrorContract = contract
}

func (b *_VTCloseErrorBuilder) WithMandatoryFields(errorType ErrorEnclosed) VTCloseErrorBuilder {
	return b.WithErrorType(errorType)
}

func (b *_VTCloseErrorBuilder) WithErrorType(errorType ErrorEnclosed) VTCloseErrorBuilder {
	b.ErrorType = errorType
	return b
}

func (b *_VTCloseErrorBuilder) WithErrorTypeBuilder(builderSupplier func(ErrorEnclosedBuilder) ErrorEnclosedBuilder) VTCloseErrorBuilder {
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

func (b *_VTCloseErrorBuilder) WithOptionalListOfVtSessionIdentifiers(listOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers) VTCloseErrorBuilder {
	b.ListOfVtSessionIdentifiers = listOfVtSessionIdentifiers
	return b
}

func (b *_VTCloseErrorBuilder) WithOptionalListOfVtSessionIdentifiersBuilder(builderSupplier func(VTCloseErrorListOfVTSessionIdentifiersBuilder) VTCloseErrorListOfVTSessionIdentifiersBuilder) VTCloseErrorBuilder {
	builder := builderSupplier(b.ListOfVtSessionIdentifiers.CreateVTCloseErrorListOfVTSessionIdentifiersBuilder())
	var err error
	b.ListOfVtSessionIdentifiers, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "VTCloseErrorListOfVTSessionIdentifiersBuilder failed"))
	}
	return b
}

func (b *_VTCloseErrorBuilder) Build() (VTCloseError, error) {
	if b.ErrorType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'errorType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._VTCloseError.deepCopy(), nil
}

func (b *_VTCloseErrorBuilder) MustBuild() VTCloseError {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_VTCloseErrorBuilder) Done() BACnetErrorBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetErrorBuilder().(*_BACnetErrorBuilder)
	}
	return b.parentBuilder
}

func (b *_VTCloseErrorBuilder) buildForBACnetError() (BACnetError, error) {
	return b.Build()
}

func (b *_VTCloseErrorBuilder) DeepCopy() any {
	_copy := b.CreateVTCloseErrorBuilder().(*_VTCloseErrorBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateVTCloseErrorBuilder creates a VTCloseErrorBuilder
func (b *_VTCloseError) CreateVTCloseErrorBuilder() VTCloseErrorBuilder {
	if b == nil {
		return NewVTCloseErrorBuilder()
	}
	return &_VTCloseErrorBuilder{_VTCloseError: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VTCloseError) GetErrorChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VTCloseError) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VTCloseError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_VTCloseError) GetListOfVtSessionIdentifiers() VTCloseErrorListOfVTSessionIdentifiers {
	return m.ListOfVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastVTCloseError(structType any) VTCloseError {
	if casted, ok := structType.(VTCloseError); ok {
		return casted
	}
	if casted, ok := structType.(*VTCloseError); ok {
		return *casted
	}
	return nil
}

func (m *_VTCloseError) GetTypeName() string {
	return "VTCloseError"
}

func (m *_VTCloseError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Optional Field (listOfVtSessionIdentifiers)
	if m.ListOfVtSessionIdentifiers != nil {
		lengthInBits += m.ListOfVtSessionIdentifiers.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_VTCloseError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_VTCloseError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__vTCloseError VTCloseError, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("VTCloseError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VTCloseError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorType, err := ReadSimpleField[ErrorEnclosed](ctx, "errorType", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorType' field"))
	}
	m.ErrorType = errorType

	var listOfVtSessionIdentifiers VTCloseErrorListOfVTSessionIdentifiers
	_listOfVtSessionIdentifiers, err := ReadOptionalField[VTCloseErrorListOfVTSessionIdentifiers](ctx, "listOfVtSessionIdentifiers", ReadComplex[VTCloseErrorListOfVTSessionIdentifiers](VTCloseErrorListOfVTSessionIdentifiersParseWithBufferProducer((uint8)(uint8(1))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfVtSessionIdentifiers' field"))
	}
	if _listOfVtSessionIdentifiers != nil {
		listOfVtSessionIdentifiers = *_listOfVtSessionIdentifiers
		m.ListOfVtSessionIdentifiers = listOfVtSessionIdentifiers
	}

	if closeErr := readBuffer.CloseContext("VTCloseError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VTCloseError")
	}

	return m, nil
}

func (m *_VTCloseError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VTCloseError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VTCloseError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VTCloseError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "errorType", m.GetErrorType(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorType' field")
		}

		if err := WriteOptionalField[VTCloseErrorListOfVTSessionIdentifiers](ctx, "listOfVtSessionIdentifiers", GetRef(m.GetListOfVtSessionIdentifiers()), WriteComplex[VTCloseErrorListOfVTSessionIdentifiers](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfVtSessionIdentifiers' field")
		}

		if popErr := writeBuffer.PopContext("VTCloseError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VTCloseError")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_VTCloseError) IsVTCloseError() {}

func (m *_VTCloseError) DeepCopy() any {
	return m.deepCopy()
}

func (m *_VTCloseError) deepCopy() *_VTCloseError {
	if m == nil {
		return nil
	}
	_VTCloseErrorCopy := &_VTCloseError{
		m.BACnetErrorContract.(*_BACnetError).deepCopy(),
		utils.DeepCopy[ErrorEnclosed](m.ErrorType),
		utils.DeepCopy[VTCloseErrorListOfVTSessionIdentifiers](m.ListOfVtSessionIdentifiers),
	}
	m.BACnetErrorContract.(*_BACnetError)._SubType = m
	return _VTCloseErrorCopy
}

func (m *_VTCloseError) String() string {
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
