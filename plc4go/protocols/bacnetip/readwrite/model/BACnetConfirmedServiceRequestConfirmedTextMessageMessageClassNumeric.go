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

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetNumericValue returns NumericValue (property field)
	GetNumericValue() BACnetContextTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
	CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric struct {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
	NumericValue BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)(nil)
var _ BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassRequirements = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)(nil)

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numericValue BACnetContextTagUnsignedInteger, tagNumber uint8) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if numericValue == nil {
		panic("numericValue of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
		NumericValue: numericValue,
	}
	_result.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder is a builder for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(numericValue BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
	// WithNumericValue adds NumericValue (property field)
	WithNumericValue(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
	// WithNumericValueBuilder adds NumericValue (property field) which is build by the builder
	WithNumericValueBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
}

// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder() creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric: new(_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)}
}

type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric

	parentBuilder *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) = (*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) setParent(contract BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract) {
	b.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = contract
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) WithMandatoryFields(numericValue BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	return b.WithNumericValue(numericValue)
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) WithNumericValue(numericValue BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	b.NumericValue = numericValue
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) WithNumericValueBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	builder := builderSupplier(b.NumericValue.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.NumericValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) Build() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric, error) {
	if b.NumericValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'numericValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) Done() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) buildForBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass() (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder().(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder creates a BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder
func (b *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) CreateBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericBuilder{_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric: b.deepCopy()}
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

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract {
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetNumericValue() BACnetContextTagUnsignedInteger {
	return m.NumericValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric(structType any) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).getLengthInBits(ctx))

	// Simple field (numericValue)
	lengthInBits += m.NumericValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, tagNumber uint8) (__bACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric, err error) {
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	numericValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "numericValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numericValue' field"))
	}
	m.NumericValue = numericValue

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "numericValue", m.GetNumericValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'numericValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) IsBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) deepCopy() *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericCopy := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.NumericValue),
	}
	m.BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassContract.(*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) String() string {
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
