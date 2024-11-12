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

// BACnetConstructedDataNumberOfStates is the corresponding interface of BACnetConstructedDataNumberOfStates
type BACnetConstructedDataNumberOfStates interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfState returns NumberOfState (property field)
	GetNumberOfState() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataNumberOfStates is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataNumberOfStates()
	// CreateBuilder creates a BACnetConstructedDataNumberOfStatesBuilder
	CreateBACnetConstructedDataNumberOfStatesBuilder() BACnetConstructedDataNumberOfStatesBuilder
}

// _BACnetConstructedDataNumberOfStates is the data-structure of this message
type _BACnetConstructedDataNumberOfStates struct {
	BACnetConstructedDataContract
	NumberOfState BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataNumberOfStates = (*_BACnetConstructedDataNumberOfStates)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataNumberOfStates)(nil)

// NewBACnetConstructedDataNumberOfStates factory function for _BACnetConstructedDataNumberOfStates
func NewBACnetConstructedDataNumberOfStates(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfState BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNumberOfStates {
	if numberOfState == nil {
		panic("numberOfState of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataNumberOfStates must not be nil")
	}
	_result := &_BACnetConstructedDataNumberOfStates{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfState:                 numberOfState,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataNumberOfStatesBuilder is a builder for BACnetConstructedDataNumberOfStates
type BACnetConstructedDataNumberOfStatesBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numberOfState BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfStatesBuilder
	// WithNumberOfState adds NumberOfState (property field)
	WithNumberOfState(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfStatesBuilder
	// WithNumberOfStateBuilder adds NumberOfState (property field) which is build by the builder
	WithNumberOfStateBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNumberOfStatesBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataNumberOfStates or returns an error if something is wrong
	Build() (BACnetConstructedDataNumberOfStates, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataNumberOfStates
}

// NewBACnetConstructedDataNumberOfStatesBuilder() creates a BACnetConstructedDataNumberOfStatesBuilder
func NewBACnetConstructedDataNumberOfStatesBuilder() BACnetConstructedDataNumberOfStatesBuilder {
	return &_BACnetConstructedDataNumberOfStatesBuilder{_BACnetConstructedDataNumberOfStates: new(_BACnetConstructedDataNumberOfStates)}
}

type _BACnetConstructedDataNumberOfStatesBuilder struct {
	*_BACnetConstructedDataNumberOfStates

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataNumberOfStatesBuilder) = (*_BACnetConstructedDataNumberOfStatesBuilder)(nil)

func (b *_BACnetConstructedDataNumberOfStatesBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) WithMandatoryFields(numberOfState BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfStatesBuilder {
	return b.WithNumberOfState(numberOfState)
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) WithNumberOfState(numberOfState BACnetApplicationTagUnsignedInteger) BACnetConstructedDataNumberOfStatesBuilder {
	b.NumberOfState = numberOfState
	return b
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) WithNumberOfStateBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataNumberOfStatesBuilder {
	builder := builderSupplier(b.NumberOfState.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfState, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) Build() (BACnetConstructedDataNumberOfStates, error) {
	if b.NumberOfState == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'numberOfState' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataNumberOfStates.deepCopy(), nil
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) MustBuild() BACnetConstructedDataNumberOfStates {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataNumberOfStatesBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataNumberOfStatesBuilder().(*_BACnetConstructedDataNumberOfStatesBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataNumberOfStatesBuilder creates a BACnetConstructedDataNumberOfStatesBuilder
func (b *_BACnetConstructedDataNumberOfStates) CreateBACnetConstructedDataNumberOfStatesBuilder() BACnetConstructedDataNumberOfStatesBuilder {
	if b == nil {
		return NewBACnetConstructedDataNumberOfStatesBuilder()
	}
	return &_BACnetConstructedDataNumberOfStatesBuilder{_BACnetConstructedDataNumberOfStates: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNumberOfStates) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_STATES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetNumberOfState() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfStates) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNumberOfState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNumberOfStates(structType any) BACnetConstructedDataNumberOfStates {
	if casted, ok := structType.(BACnetConstructedDataNumberOfStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNumberOfStates) GetTypeName() string {
	return "BACnetConstructedDataNumberOfStates"
}

func (m *_BACnetConstructedDataNumberOfStates) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (numberOfState)
	lengthInBits += m.NumberOfState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNumberOfStates) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataNumberOfStates) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataNumberOfStates BACnetConstructedDataNumberOfStates, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNumberOfStates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numberOfState, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfState", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfState' field"))
	}
	m.NumberOfState = numberOfState

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), numberOfState)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNumberOfStates")
	}

	return m, nil
}

func (m *_BACnetConstructedDataNumberOfStates) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNumberOfStates) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfStates"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNumberOfStates")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfState", m.GetNumberOfState(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfStates"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNumberOfStates")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNumberOfStates) IsBACnetConstructedDataNumberOfStates() {}

func (m *_BACnetConstructedDataNumberOfStates) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataNumberOfStates) deepCopy() *_BACnetConstructedDataNumberOfStates {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataNumberOfStatesCopy := &_BACnetConstructedDataNumberOfStates{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfState),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataNumberOfStatesCopy
}

func (m *_BACnetConstructedDataNumberOfStates) String() string {
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
