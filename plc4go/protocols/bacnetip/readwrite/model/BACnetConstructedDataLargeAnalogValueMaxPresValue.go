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

// BACnetConstructedDataLargeAnalogValueMaxPresValue is the corresponding interface of BACnetConstructedDataLargeAnalogValueMaxPresValue
type BACnetConstructedDataLargeAnalogValueMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueMaxPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueMaxPresValue()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
	CreateBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder() BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
}

// _BACnetConstructedDataLargeAnalogValueMaxPresValue is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueMaxPresValue = (*_BACnetConstructedDataLargeAnalogValueMaxPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueMaxPresValue)(nil)

// NewBACnetConstructedDataLargeAnalogValueMaxPresValue factory function for _BACnetConstructedDataLargeAnalogValueMaxPresValue
func NewBACnetConstructedDataLargeAnalogValueMaxPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxPresValue BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueMaxPresValue {
	if maxPresValue == nil {
		panic("maxPresValue of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueMaxPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder is a builder for BACnetConstructedDataLargeAnalogValueMaxPresValue
type BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
	// WithMaxPresValue adds MaxPresValue (property field)
	WithMaxPresValue(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
	// WithMaxPresValueBuilder adds MaxPresValue (property field) which is build by the builder
	WithMaxPresValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueMaxPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueMaxPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueMaxPresValue
}

// NewBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder() creates a BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
func NewBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder() BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder {
	return &_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder{_BACnetConstructedDataLargeAnalogValueMaxPresValue: new(_BACnetConstructedDataLargeAnalogValueMaxPresValue)}
}

type _BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueMaxPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) = (*_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) WithMandatoryFields(maxPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder {
	return b.WithMaxPresValue(maxPresValue)
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) WithMaxPresValue(maxPresValue BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder {
	b.MaxPresValue = maxPresValue
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) WithMaxPresValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder {
	builder := builderSupplier(b.MaxPresValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.MaxPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) Build() (BACnetConstructedDataLargeAnalogValueMaxPresValue, error) {
	if b.MaxPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueMaxPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueMaxPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder().(*_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder creates a BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder
func (b *_BACnetConstructedDataLargeAnalogValueMaxPresValue) CreateBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder() BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueMaxPresValueBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueMaxPresValueBuilder{_BACnetConstructedDataLargeAnalogValueMaxPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetMaxPresValue() BACnetApplicationTagDouble {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueMaxPresValue(structType any) BACnetConstructedDataLargeAnalogValueMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueMaxPresValue"
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueMaxPresValue BACnetConstructedDataLargeAnalogValueMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) IsBACnetConstructedDataLargeAnalogValueMaxPresValue() {
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) deepCopy() *_BACnetConstructedDataLargeAnalogValueMaxPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueMaxPresValueCopy := &_BACnetConstructedDataLargeAnalogValueMaxPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.MaxPresValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueMaxPresValueCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueMaxPresValue) String() string {
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
