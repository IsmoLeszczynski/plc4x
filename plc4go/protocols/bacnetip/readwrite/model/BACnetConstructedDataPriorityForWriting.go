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

// BACnetConstructedDataPriorityForWriting is the corresponding interface of BACnetConstructedDataPriorityForWriting
type BACnetConstructedDataPriorityForWriting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPriorityForWriting returns PriorityForWriting (property field)
	GetPriorityForWriting() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataPriorityForWriting is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataPriorityForWriting()
	// CreateBuilder creates a BACnetConstructedDataPriorityForWritingBuilder
	CreateBACnetConstructedDataPriorityForWritingBuilder() BACnetConstructedDataPriorityForWritingBuilder
}

// _BACnetConstructedDataPriorityForWriting is the data-structure of this message
type _BACnetConstructedDataPriorityForWriting struct {
	BACnetConstructedDataContract
	PriorityForWriting BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataPriorityForWriting = (*_BACnetConstructedDataPriorityForWriting)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataPriorityForWriting)(nil)

// NewBACnetConstructedDataPriorityForWriting factory function for _BACnetConstructedDataPriorityForWriting
func NewBACnetConstructedDataPriorityForWriting(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, priorityForWriting BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPriorityForWriting {
	if priorityForWriting == nil {
		panic("priorityForWriting of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataPriorityForWriting must not be nil")
	}
	_result := &_BACnetConstructedDataPriorityForWriting{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PriorityForWriting:            priorityForWriting,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataPriorityForWritingBuilder is a builder for BACnetConstructedDataPriorityForWriting
type BACnetConstructedDataPriorityForWritingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(priorityForWriting BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPriorityForWritingBuilder
	// WithPriorityForWriting adds PriorityForWriting (property field)
	WithPriorityForWriting(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPriorityForWritingBuilder
	// WithPriorityForWritingBuilder adds PriorityForWriting (property field) which is build by the builder
	WithPriorityForWritingBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPriorityForWritingBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataPriorityForWriting or returns an error if something is wrong
	Build() (BACnetConstructedDataPriorityForWriting, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataPriorityForWriting
}

// NewBACnetConstructedDataPriorityForWritingBuilder() creates a BACnetConstructedDataPriorityForWritingBuilder
func NewBACnetConstructedDataPriorityForWritingBuilder() BACnetConstructedDataPriorityForWritingBuilder {
	return &_BACnetConstructedDataPriorityForWritingBuilder{_BACnetConstructedDataPriorityForWriting: new(_BACnetConstructedDataPriorityForWriting)}
}

type _BACnetConstructedDataPriorityForWritingBuilder struct {
	*_BACnetConstructedDataPriorityForWriting

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataPriorityForWritingBuilder) = (*_BACnetConstructedDataPriorityForWritingBuilder)(nil)

func (b *_BACnetConstructedDataPriorityForWritingBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) WithMandatoryFields(priorityForWriting BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPriorityForWritingBuilder {
	return b.WithPriorityForWriting(priorityForWriting)
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) WithPriorityForWriting(priorityForWriting BACnetApplicationTagUnsignedInteger) BACnetConstructedDataPriorityForWritingBuilder {
	b.PriorityForWriting = priorityForWriting
	return b
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) WithPriorityForWritingBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataPriorityForWritingBuilder {
	builder := builderSupplier(b.PriorityForWriting.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.PriorityForWriting, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) Build() (BACnetConstructedDataPriorityForWriting, error) {
	if b.PriorityForWriting == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'priorityForWriting' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataPriorityForWriting.deepCopy(), nil
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) MustBuild() BACnetConstructedDataPriorityForWriting {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataPriorityForWritingBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataPriorityForWritingBuilder().(*_BACnetConstructedDataPriorityForWritingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataPriorityForWritingBuilder creates a BACnetConstructedDataPriorityForWritingBuilder
func (b *_BACnetConstructedDataPriorityForWriting) CreateBACnetConstructedDataPriorityForWritingBuilder() BACnetConstructedDataPriorityForWritingBuilder {
	if b == nil {
		return NewBACnetConstructedDataPriorityForWritingBuilder()
	}
	return &_BACnetConstructedDataPriorityForWritingBuilder{_BACnetConstructedDataPriorityForWriting: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPriorityForWriting) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_FOR_WRITING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetPriorityForWriting() BACnetApplicationTagUnsignedInteger {
	return m.PriorityForWriting
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetPriorityForWriting())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPriorityForWriting(structType any) BACnetConstructedDataPriorityForWriting {
	if casted, ok := structType.(BACnetConstructedDataPriorityForWriting); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityForWriting); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPriorityForWriting) GetTypeName() string {
	return "BACnetConstructedDataPriorityForWriting"
}

func (m *_BACnetConstructedDataPriorityForWriting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (priorityForWriting)
	lengthInBits += m.PriorityForWriting.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPriorityForWriting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataPriorityForWriting) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataPriorityForWriting BACnetConstructedDataPriorityForWriting, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityForWriting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityForWriting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	priorityForWriting, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "priorityForWriting", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priorityForWriting' field"))
	}
	m.PriorityForWriting = priorityForWriting

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), priorityForWriting)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityForWriting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityForWriting")
	}

	return m, nil
}

func (m *_BACnetConstructedDataPriorityForWriting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPriorityForWriting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityForWriting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityForWriting")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "priorityForWriting", m.GetPriorityForWriting(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priorityForWriting' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityForWriting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityForWriting")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPriorityForWriting) IsBACnetConstructedDataPriorityForWriting() {}

func (m *_BACnetConstructedDataPriorityForWriting) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataPriorityForWriting) deepCopy() *_BACnetConstructedDataPriorityForWriting {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataPriorityForWritingCopy := &_BACnetConstructedDataPriorityForWriting{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.PriorityForWriting),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataPriorityForWritingCopy
}

func (m *_BACnetConstructedDataPriorityForWriting) String() string {
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
