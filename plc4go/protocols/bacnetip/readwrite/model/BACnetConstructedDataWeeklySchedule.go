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

// BACnetConstructedDataWeeklySchedule is the corresponding interface of BACnetConstructedDataWeeklySchedule
type BACnetConstructedDataWeeklySchedule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetWeeklySchedule returns WeeklySchedule (property field)
	GetWeeklySchedule() []BACnetDailySchedule
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataWeeklySchedule is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataWeeklySchedule()
	// CreateBuilder creates a BACnetConstructedDataWeeklyScheduleBuilder
	CreateBACnetConstructedDataWeeklyScheduleBuilder() BACnetConstructedDataWeeklyScheduleBuilder
}

// _BACnetConstructedDataWeeklySchedule is the data-structure of this message
type _BACnetConstructedDataWeeklySchedule struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	WeeklySchedule       []BACnetDailySchedule
}

var _ BACnetConstructedDataWeeklySchedule = (*_BACnetConstructedDataWeeklySchedule)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataWeeklySchedule)(nil)

// NewBACnetConstructedDataWeeklySchedule factory function for _BACnetConstructedDataWeeklySchedule
func NewBACnetConstructedDataWeeklySchedule(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, weeklySchedule []BACnetDailySchedule, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataWeeklySchedule {
	_result := &_BACnetConstructedDataWeeklySchedule{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		WeeklySchedule:                weeklySchedule,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataWeeklyScheduleBuilder is a builder for BACnetConstructedDataWeeklySchedule
type BACnetConstructedDataWeeklyScheduleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(weeklySchedule []BACnetDailySchedule) BACnetConstructedDataWeeklyScheduleBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWeeklyScheduleBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataWeeklyScheduleBuilder
	// WithWeeklySchedule adds WeeklySchedule (property field)
	WithWeeklySchedule(...BACnetDailySchedule) BACnetConstructedDataWeeklyScheduleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataWeeklySchedule or returns an error if something is wrong
	Build() (BACnetConstructedDataWeeklySchedule, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataWeeklySchedule
}

// NewBACnetConstructedDataWeeklyScheduleBuilder() creates a BACnetConstructedDataWeeklyScheduleBuilder
func NewBACnetConstructedDataWeeklyScheduleBuilder() BACnetConstructedDataWeeklyScheduleBuilder {
	return &_BACnetConstructedDataWeeklyScheduleBuilder{_BACnetConstructedDataWeeklySchedule: new(_BACnetConstructedDataWeeklySchedule)}
}

type _BACnetConstructedDataWeeklyScheduleBuilder struct {
	*_BACnetConstructedDataWeeklySchedule

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataWeeklyScheduleBuilder) = (*_BACnetConstructedDataWeeklyScheduleBuilder)(nil)

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) WithMandatoryFields(weeklySchedule []BACnetDailySchedule) BACnetConstructedDataWeeklyScheduleBuilder {
	return b.WithWeeklySchedule(weeklySchedule...)
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataWeeklyScheduleBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataWeeklyScheduleBuilder {
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

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) WithWeeklySchedule(weeklySchedule ...BACnetDailySchedule) BACnetConstructedDataWeeklyScheduleBuilder {
	b.WeeklySchedule = weeklySchedule
	return b
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) Build() (BACnetConstructedDataWeeklySchedule, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataWeeklySchedule.deepCopy(), nil
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) MustBuild() BACnetConstructedDataWeeklySchedule {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataWeeklyScheduleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataWeeklyScheduleBuilder().(*_BACnetConstructedDataWeeklyScheduleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataWeeklyScheduleBuilder creates a BACnetConstructedDataWeeklyScheduleBuilder
func (b *_BACnetConstructedDataWeeklySchedule) CreateBACnetConstructedDataWeeklyScheduleBuilder() BACnetConstructedDataWeeklyScheduleBuilder {
	if b == nil {
		return NewBACnetConstructedDataWeeklyScheduleBuilder()
	}
	return &_BACnetConstructedDataWeeklyScheduleBuilder{_BACnetConstructedDataWeeklySchedule: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataWeeklySchedule) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_WEEKLY_SCHEDULE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataWeeklySchedule) GetWeeklySchedule() []BACnetDailySchedule {
	return m.WeeklySchedule
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataWeeklySchedule) GetZero() uint64 {
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
func CastBACnetConstructedDataWeeklySchedule(structType any) BACnetConstructedDataWeeklySchedule {
	if casted, ok := structType.(BACnetConstructedDataWeeklySchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataWeeklySchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataWeeklySchedule) GetTypeName() string {
	return "BACnetConstructedDataWeeklySchedule"
}

func (m *_BACnetConstructedDataWeeklySchedule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.WeeklySchedule) > 0 {
		for _, element := range m.WeeklySchedule {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataWeeklySchedule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataWeeklySchedule) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataWeeklySchedule BACnetConstructedDataWeeklySchedule, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataWeeklySchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataWeeklySchedule")
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

	weeklySchedule, err := ReadTerminatedArrayField[BACnetDailySchedule](ctx, "weeklySchedule", ReadComplex[BACnetDailySchedule](BACnetDailyScheduleParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'weeklySchedule' field"))
	}
	m.WeeklySchedule = weeklySchedule

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(weeklySchedule)) == (7)))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "weeklySchedule should have exactly 7 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataWeeklySchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataWeeklySchedule")
	}

	return m, nil
}

func (m *_BACnetConstructedDataWeeklySchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataWeeklySchedule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataWeeklySchedule"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataWeeklySchedule")
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

		if err := WriteComplexTypeArrayField(ctx, "weeklySchedule", m.GetWeeklySchedule(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'weeklySchedule' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataWeeklySchedule"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataWeeklySchedule")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataWeeklySchedule) IsBACnetConstructedDataWeeklySchedule() {}

func (m *_BACnetConstructedDataWeeklySchedule) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataWeeklySchedule) deepCopy() *_BACnetConstructedDataWeeklySchedule {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataWeeklyScheduleCopy := &_BACnetConstructedDataWeeklySchedule{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetDailySchedule, BACnetDailySchedule](m.WeeklySchedule),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataWeeklyScheduleCopy
}

func (m *_BACnetConstructedDataWeeklySchedule) String() string {
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
