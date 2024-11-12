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

// BACnetConstructedDataLinkSpeeds is the corresponding interface of BACnetConstructedDataLinkSpeeds
type BACnetConstructedDataLinkSpeeds interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetLinkSpeeds returns LinkSpeeds (property field)
	GetLinkSpeeds() []BACnetApplicationTagReal
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataLinkSpeeds is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLinkSpeeds()
	// CreateBuilder creates a BACnetConstructedDataLinkSpeedsBuilder
	CreateBACnetConstructedDataLinkSpeedsBuilder() BACnetConstructedDataLinkSpeedsBuilder
}

// _BACnetConstructedDataLinkSpeeds is the data-structure of this message
type _BACnetConstructedDataLinkSpeeds struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	LinkSpeeds           []BACnetApplicationTagReal
}

var _ BACnetConstructedDataLinkSpeeds = (*_BACnetConstructedDataLinkSpeeds)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLinkSpeeds)(nil)

// NewBACnetConstructedDataLinkSpeeds factory function for _BACnetConstructedDataLinkSpeeds
func NewBACnetConstructedDataLinkSpeeds(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, linkSpeeds []BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLinkSpeeds {
	_result := &_BACnetConstructedDataLinkSpeeds{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		LinkSpeeds:                    linkSpeeds,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLinkSpeedsBuilder is a builder for BACnetConstructedDataLinkSpeeds
type BACnetConstructedDataLinkSpeedsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(linkSpeeds []BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedsBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLinkSpeedsBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLinkSpeedsBuilder
	// WithLinkSpeeds adds LinkSpeeds (property field)
	WithLinkSpeeds(...BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLinkSpeeds or returns an error if something is wrong
	Build() (BACnetConstructedDataLinkSpeeds, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLinkSpeeds
}

// NewBACnetConstructedDataLinkSpeedsBuilder() creates a BACnetConstructedDataLinkSpeedsBuilder
func NewBACnetConstructedDataLinkSpeedsBuilder() BACnetConstructedDataLinkSpeedsBuilder {
	return &_BACnetConstructedDataLinkSpeedsBuilder{_BACnetConstructedDataLinkSpeeds: new(_BACnetConstructedDataLinkSpeeds)}
}

type _BACnetConstructedDataLinkSpeedsBuilder struct {
	*_BACnetConstructedDataLinkSpeeds

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLinkSpeedsBuilder) = (*_BACnetConstructedDataLinkSpeedsBuilder)(nil)

func (b *_BACnetConstructedDataLinkSpeedsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) WithMandatoryFields(linkSpeeds []BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedsBuilder {
	return b.WithLinkSpeeds(linkSpeeds...)
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLinkSpeedsBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLinkSpeedsBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) WithLinkSpeeds(linkSpeeds ...BACnetApplicationTagReal) BACnetConstructedDataLinkSpeedsBuilder {
	b.LinkSpeeds = linkSpeeds
	return b
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) Build() (BACnetConstructedDataLinkSpeeds, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLinkSpeeds.deepCopy(), nil
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) MustBuild() BACnetConstructedDataLinkSpeeds {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLinkSpeedsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLinkSpeedsBuilder().(*_BACnetConstructedDataLinkSpeedsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLinkSpeedsBuilder creates a BACnetConstructedDataLinkSpeedsBuilder
func (b *_BACnetConstructedDataLinkSpeeds) CreateBACnetConstructedDataLinkSpeedsBuilder() BACnetConstructedDataLinkSpeedsBuilder {
	if b == nil {
		return NewBACnetConstructedDataLinkSpeedsBuilder()
	}
	return &_BACnetConstructedDataLinkSpeedsBuilder{_BACnetConstructedDataLinkSpeeds: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeeds) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLinkSpeeds) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LINK_SPEEDS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLinkSpeeds) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeeds) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataLinkSpeeds) GetLinkSpeeds() []BACnetApplicationTagReal {
	return m.LinkSpeeds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeeds) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLinkSpeeds(structType any) BACnetConstructedDataLinkSpeeds {
	if casted, ok := structType.(BACnetConstructedDataLinkSpeeds); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeeds); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLinkSpeeds) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeeds"
}

func (m *_BACnetConstructedDataLinkSpeeds) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.LinkSpeeds) > 0 {
		for _, element := range m.LinkSpeeds {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLinkSpeeds) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLinkSpeeds) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLinkSpeeds BACnetConstructedDataLinkSpeeds, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeeds"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeeds")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	linkSpeeds, err := ReadTerminatedArrayField[BACnetApplicationTagReal](ctx, "linkSpeeds", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'linkSpeeds' field"))
	}
	m.LinkSpeeds = linkSpeeds

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeeds"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeeds")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLinkSpeeds) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLinkSpeeds) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeeds"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeeds")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "linkSpeeds", m.GetLinkSpeeds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'linkSpeeds' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeeds"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeeds")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLinkSpeeds) IsBACnetConstructedDataLinkSpeeds() {}

func (m *_BACnetConstructedDataLinkSpeeds) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLinkSpeeds) deepCopy() *_BACnetConstructedDataLinkSpeeds {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLinkSpeedsCopy := &_BACnetConstructedDataLinkSpeeds{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetApplicationTagReal, BACnetApplicationTagReal](m.LinkSpeeds),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLinkSpeedsCopy
}

func (m *_BACnetConstructedDataLinkSpeeds) String() string {
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
