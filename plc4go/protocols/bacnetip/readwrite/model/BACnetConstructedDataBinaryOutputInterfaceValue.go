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

// BACnetConstructedDataBinaryOutputInterfaceValue is the corresponding interface of BACnetConstructedDataBinaryOutputInterfaceValue
type BACnetConstructedDataBinaryOutputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
	// IsBACnetConstructedDataBinaryOutputInterfaceValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataBinaryOutputInterfaceValue()
	// CreateBuilder creates a BACnetConstructedDataBinaryOutputInterfaceValueBuilder
	CreateBACnetConstructedDataBinaryOutputInterfaceValueBuilder() BACnetConstructedDataBinaryOutputInterfaceValueBuilder
}

// _BACnetConstructedDataBinaryOutputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataBinaryOutputInterfaceValue struct {
	BACnetConstructedDataContract
	InterfaceValue BACnetOptionalBinaryPV
}

var _ BACnetConstructedDataBinaryOutputInterfaceValue = (*_BACnetConstructedDataBinaryOutputInterfaceValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataBinaryOutputInterfaceValue)(nil)

// NewBACnetConstructedDataBinaryOutputInterfaceValue factory function for _BACnetConstructedDataBinaryOutputInterfaceValue
func NewBACnetConstructedDataBinaryOutputInterfaceValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, interfaceValue BACnetOptionalBinaryPV, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputInterfaceValue {
	if interfaceValue == nil {
		panic("interfaceValue of type BACnetOptionalBinaryPV for BACnetConstructedDataBinaryOutputInterfaceValue must not be nil")
	}
	_result := &_BACnetConstructedDataBinaryOutputInterfaceValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InterfaceValue:                interfaceValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataBinaryOutputInterfaceValueBuilder is a builder for BACnetConstructedDataBinaryOutputInterfaceValue
type BACnetConstructedDataBinaryOutputInterfaceValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataBinaryOutputInterfaceValueBuilder
	// WithInterfaceValue adds InterfaceValue (property field)
	WithInterfaceValue(BACnetOptionalBinaryPV) BACnetConstructedDataBinaryOutputInterfaceValueBuilder
	// WithInterfaceValueBuilder adds InterfaceValue (property field) which is build by the builder
	WithInterfaceValueBuilder(func(BACnetOptionalBinaryPVBuilder) BACnetOptionalBinaryPVBuilder) BACnetConstructedDataBinaryOutputInterfaceValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataBinaryOutputInterfaceValue or returns an error if something is wrong
	Build() (BACnetConstructedDataBinaryOutputInterfaceValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataBinaryOutputInterfaceValue
}

// NewBACnetConstructedDataBinaryOutputInterfaceValueBuilder() creates a BACnetConstructedDataBinaryOutputInterfaceValueBuilder
func NewBACnetConstructedDataBinaryOutputInterfaceValueBuilder() BACnetConstructedDataBinaryOutputInterfaceValueBuilder {
	return &_BACnetConstructedDataBinaryOutputInterfaceValueBuilder{_BACnetConstructedDataBinaryOutputInterfaceValue: new(_BACnetConstructedDataBinaryOutputInterfaceValue)}
}

type _BACnetConstructedDataBinaryOutputInterfaceValueBuilder struct {
	*_BACnetConstructedDataBinaryOutputInterfaceValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataBinaryOutputInterfaceValueBuilder) = (*_BACnetConstructedDataBinaryOutputInterfaceValueBuilder)(nil)

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) WithMandatoryFields(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataBinaryOutputInterfaceValueBuilder {
	return b.WithInterfaceValue(interfaceValue)
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) WithInterfaceValue(interfaceValue BACnetOptionalBinaryPV) BACnetConstructedDataBinaryOutputInterfaceValueBuilder {
	b.InterfaceValue = interfaceValue
	return b
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) WithInterfaceValueBuilder(builderSupplier func(BACnetOptionalBinaryPVBuilder) BACnetOptionalBinaryPVBuilder) BACnetConstructedDataBinaryOutputInterfaceValueBuilder {
	builder := builderSupplier(b.InterfaceValue.CreateBACnetOptionalBinaryPVBuilder())
	var err error
	b.InterfaceValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOptionalBinaryPVBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) Build() (BACnetConstructedDataBinaryOutputInterfaceValue, error) {
	if b.InterfaceValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'interfaceValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataBinaryOutputInterfaceValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) MustBuild() BACnetConstructedDataBinaryOutputInterfaceValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataBinaryOutputInterfaceValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataBinaryOutputInterfaceValueBuilder().(*_BACnetConstructedDataBinaryOutputInterfaceValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataBinaryOutputInterfaceValueBuilder creates a BACnetConstructedDataBinaryOutputInterfaceValueBuilder
func (b *_BACnetConstructedDataBinaryOutputInterfaceValue) CreateBACnetConstructedDataBinaryOutputInterfaceValueBuilder() BACnetConstructedDataBinaryOutputInterfaceValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataBinaryOutputInterfaceValueBuilder()
	}
	return &_BACnetConstructedDataBinaryOutputInterfaceValueBuilder{_BACnetConstructedDataBinaryOutputInterfaceValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputInterfaceValue(structType any) BACnetConstructedDataBinaryOutputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputInterfaceValue"
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataBinaryOutputInterfaceValue BACnetConstructedDataBinaryOutputInterfaceValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	interfaceValue, err := ReadSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", ReadComplex[BACnetOptionalBinaryPV](BACnetOptionalBinaryPVParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'interfaceValue' field"))
	}
	m.InterfaceValue = interfaceValue

	actualValue, err := ReadVirtualField[BACnetOptionalBinaryPV](ctx, "actualValue", (*BACnetOptionalBinaryPV)(nil), interfaceValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputInterfaceValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputInterfaceValue")
		}

		if err := WriteSimpleField[BACnetOptionalBinaryPV](ctx, "interfaceValue", m.GetInterfaceValue(), WriteComplex[BACnetOptionalBinaryPV](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputInterfaceValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) IsBACnetConstructedDataBinaryOutputInterfaceValue() {
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) deepCopy() *_BACnetConstructedDataBinaryOutputInterfaceValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataBinaryOutputInterfaceValueCopy := &_BACnetConstructedDataBinaryOutputInterfaceValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetOptionalBinaryPV](m.InterfaceValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataBinaryOutputInterfaceValueCopy
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) String() string {
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
