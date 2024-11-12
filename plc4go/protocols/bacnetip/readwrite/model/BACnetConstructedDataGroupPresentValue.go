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

// BACnetConstructedDataGroupPresentValue is the corresponding interface of BACnetConstructedDataGroupPresentValue
type BACnetConstructedDataGroupPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() []BACnetReadAccessResult
	// IsBACnetConstructedDataGroupPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupPresentValue()
	// CreateBuilder creates a BACnetConstructedDataGroupPresentValueBuilder
	CreateBACnetConstructedDataGroupPresentValueBuilder() BACnetConstructedDataGroupPresentValueBuilder
}

// _BACnetConstructedDataGroupPresentValue is the data-structure of this message
type _BACnetConstructedDataGroupPresentValue struct {
	BACnetConstructedDataContract
	PresentValue []BACnetReadAccessResult
}

var _ BACnetConstructedDataGroupPresentValue = (*_BACnetConstructedDataGroupPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupPresentValue)(nil)

// NewBACnetConstructedDataGroupPresentValue factory function for _BACnetConstructedDataGroupPresentValue
func NewBACnetConstructedDataGroupPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue []BACnetReadAccessResult, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupPresentValue {
	_result := &_BACnetConstructedDataGroupPresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupPresentValueBuilder is a builder for BACnetConstructedDataGroupPresentValue
type BACnetConstructedDataGroupPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue []BACnetReadAccessResult) BACnetConstructedDataGroupPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(...BACnetReadAccessResult) BACnetConstructedDataGroupPresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataGroupPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupPresentValue
}

// NewBACnetConstructedDataGroupPresentValueBuilder() creates a BACnetConstructedDataGroupPresentValueBuilder
func NewBACnetConstructedDataGroupPresentValueBuilder() BACnetConstructedDataGroupPresentValueBuilder {
	return &_BACnetConstructedDataGroupPresentValueBuilder{_BACnetConstructedDataGroupPresentValue: new(_BACnetConstructedDataGroupPresentValue)}
}

type _BACnetConstructedDataGroupPresentValueBuilder struct {
	*_BACnetConstructedDataGroupPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupPresentValueBuilder) = (*_BACnetConstructedDataGroupPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataGroupPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) WithMandatoryFields(presentValue []BACnetReadAccessResult) BACnetConstructedDataGroupPresentValueBuilder {
	return b.WithPresentValue(presentValue...)
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) WithPresentValue(presentValue ...BACnetReadAccessResult) BACnetConstructedDataGroupPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) Build() (BACnetConstructedDataGroupPresentValue, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataGroupPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) MustBuild() BACnetConstructedDataGroupPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataGroupPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataGroupPresentValueBuilder().(*_BACnetConstructedDataGroupPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataGroupPresentValueBuilder creates a BACnetConstructedDataGroupPresentValueBuilder
func (b *_BACnetConstructedDataGroupPresentValue) CreateBACnetConstructedDataGroupPresentValueBuilder() BACnetConstructedDataGroupPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataGroupPresentValueBuilder()
	}
	return &_BACnetConstructedDataGroupPresentValueBuilder{_BACnetConstructedDataGroupPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_GROUP
}

func (m *_BACnetConstructedDataGroupPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupPresentValue) GetPresentValue() []BACnetReadAccessResult {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupPresentValue(structType any) BACnetConstructedDataGroupPresentValue {
	if casted, ok := structType.(BACnetConstructedDataGroupPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupPresentValue) GetTypeName() string {
	return "BACnetConstructedDataGroupPresentValue"
}

func (m *_BACnetConstructedDataGroupPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.PresentValue) > 0 {
		for _, element := range m.PresentValue {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupPresentValue BACnetConstructedDataGroupPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadTerminatedArrayField[BACnetReadAccessResult](ctx, "presentValue", ReadComplex[BACnetReadAccessResult](BACnetReadAccessResultParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupPresentValue")
		}

		if err := WriteComplexTypeArrayField(ctx, "presentValue", m.GetPresentValue(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupPresentValue) IsBACnetConstructedDataGroupPresentValue() {}

func (m *_BACnetConstructedDataGroupPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupPresentValue) deepCopy() *_BACnetConstructedDataGroupPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupPresentValueCopy := &_BACnetConstructedDataGroupPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetReadAccessResult, BACnetReadAccessResult](m.PresentValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupPresentValueCopy
}

func (m *_BACnetConstructedDataGroupPresentValue) String() string {
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
