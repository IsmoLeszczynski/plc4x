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

// BACnetConstructedDataIntegerValueHighLimit is the corresponding interface of BACnetConstructedDataIntegerValueHighLimit
type BACnetConstructedDataIntegerValueHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataIntegerValueHighLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueHighLimit()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueHighLimitBuilder
	CreateBACnetConstructedDataIntegerValueHighLimitBuilder() BACnetConstructedDataIntegerValueHighLimitBuilder
}

// _BACnetConstructedDataIntegerValueHighLimit is the data-structure of this message
type _BACnetConstructedDataIntegerValueHighLimit struct {
	BACnetConstructedDataContract
	HighLimit BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataIntegerValueHighLimit = (*_BACnetConstructedDataIntegerValueHighLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueHighLimit)(nil)

// NewBACnetConstructedDataIntegerValueHighLimit factory function for _BACnetConstructedDataIntegerValueHighLimit
func NewBACnetConstructedDataIntegerValueHighLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, highLimit BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueHighLimit {
	if highLimit == nil {
		panic("highLimit of type BACnetApplicationTagSignedInteger for BACnetConstructedDataIntegerValueHighLimit must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueHighLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		HighLimit:                     highLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueHighLimitBuilder is a builder for BACnetConstructedDataIntegerValueHighLimit
type BACnetConstructedDataIntegerValueHighLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(highLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueHighLimitBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueHighLimitBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueHighLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIntegerValueHighLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueHighLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueHighLimit
}

// NewBACnetConstructedDataIntegerValueHighLimitBuilder() creates a BACnetConstructedDataIntegerValueHighLimitBuilder
func NewBACnetConstructedDataIntegerValueHighLimitBuilder() BACnetConstructedDataIntegerValueHighLimitBuilder {
	return &_BACnetConstructedDataIntegerValueHighLimitBuilder{_BACnetConstructedDataIntegerValueHighLimit: new(_BACnetConstructedDataIntegerValueHighLimit)}
}

type _BACnetConstructedDataIntegerValueHighLimitBuilder struct {
	*_BACnetConstructedDataIntegerValueHighLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueHighLimitBuilder) = (*_BACnetConstructedDataIntegerValueHighLimitBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) WithMandatoryFields(highLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueHighLimitBuilder {
	return b.WithHighLimit(highLimit)
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) WithHighLimit(highLimit BACnetApplicationTagSignedInteger) BACnetConstructedDataIntegerValueHighLimitBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) WithHighLimitBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataIntegerValueHighLimitBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) Build() (BACnetConstructedDataIntegerValueHighLimit, error) {
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueHighLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) MustBuild() BACnetConstructedDataIntegerValueHighLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueHighLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueHighLimitBuilder().(*_BACnetConstructedDataIntegerValueHighLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueHighLimitBuilder creates a BACnetConstructedDataIntegerValueHighLimitBuilder
func (b *_BACnetConstructedDataIntegerValueHighLimit) CreateBACnetConstructedDataIntegerValueHighLimitBuilder() BACnetConstructedDataIntegerValueHighLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueHighLimitBuilder()
	}
	return &_BACnetConstructedDataIntegerValueHighLimitBuilder{_BACnetConstructedDataIntegerValueHighLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetHighLimit() BACnetApplicationTagSignedInteger {
	return m.HighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueHighLimit(structType any) BACnetConstructedDataIntegerValueHighLimit {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueHighLimit"
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueHighLimit BACnetConstructedDataIntegerValueHighLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	highLimit, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "highLimit", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), highLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueHighLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueHighLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueHighLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) IsBACnetConstructedDataIntegerValueHighLimit() {
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) deepCopy() *_BACnetConstructedDataIntegerValueHighLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueHighLimitCopy := &_BACnetConstructedDataIntegerValueHighLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.HighLimit),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueHighLimitCopy
}

func (m *_BACnetConstructedDataIntegerValueHighLimit) String() string {
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
