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

// BACnetNotificationParametersDoubleOutOfRange is the corresponding interface of BACnetNotificationParametersDoubleOutOfRange
type BACnetNotificationParametersDoubleOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagDouble
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagDouble
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagDouble
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersDoubleOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersDoubleOutOfRange()
	// CreateBuilder creates a BACnetNotificationParametersDoubleOutOfRangeBuilder
	CreateBACnetNotificationParametersDoubleOutOfRangeBuilder() BACnetNotificationParametersDoubleOutOfRangeBuilder
}

// _BACnetNotificationParametersDoubleOutOfRange is the data-structure of this message
type _BACnetNotificationParametersDoubleOutOfRange struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagDouble
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagDouble
	ExceededLimit   BACnetContextTagDouble
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersDoubleOutOfRange = (*_BACnetNotificationParametersDoubleOutOfRange)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersDoubleOutOfRange)(nil)

// NewBACnetNotificationParametersDoubleOutOfRange factory function for _BACnetNotificationParametersDoubleOutOfRange
func NewBACnetNotificationParametersDoubleOutOfRange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagDouble, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagDouble, exceededLimit BACnetContextTagDouble, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersDoubleOutOfRange {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	if exceedingValue == nil {
		panic("exceedingValue of type BACnetContextTagDouble for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagDouble for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	if exceededLimit == nil {
		panic("exceededLimit of type BACnetContextTagDouble for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersDoubleOutOfRange must not be nil")
	}
	_result := &_BACnetNotificationParametersDoubleOutOfRange{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		ExceedingValue:                       exceedingValue,
		StatusFlags:                          statusFlags,
		Deadband:                             deadband,
		ExceededLimit:                        exceededLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersDoubleOutOfRangeBuilder is a builder for BACnetNotificationParametersDoubleOutOfRange
type BACnetNotificationParametersDoubleOutOfRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagDouble, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagDouble, exceededLimit BACnetContextTagDouble, innerClosingTag BACnetClosingTag) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithExceedingValue adds ExceedingValue (property field)
	WithExceedingValue(BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithExceedingValueBuilder adds ExceedingValue (property field) which is build by the builder
	WithExceedingValueBuilder(func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithExceededLimit adds ExceededLimit (property field)
	WithExceededLimit(BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithExceededLimitBuilder adds ExceededLimit (property field) which is build by the builder
	WithExceededLimitBuilder(func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersBuilder
	// Build builds the BACnetNotificationParametersDoubleOutOfRange or returns an error if something is wrong
	Build() (BACnetNotificationParametersDoubleOutOfRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersDoubleOutOfRange
}

// NewBACnetNotificationParametersDoubleOutOfRangeBuilder() creates a BACnetNotificationParametersDoubleOutOfRangeBuilder
func NewBACnetNotificationParametersDoubleOutOfRangeBuilder() BACnetNotificationParametersDoubleOutOfRangeBuilder {
	return &_BACnetNotificationParametersDoubleOutOfRangeBuilder{_BACnetNotificationParametersDoubleOutOfRange: new(_BACnetNotificationParametersDoubleOutOfRange)}
}

type _BACnetNotificationParametersDoubleOutOfRangeBuilder struct {
	*_BACnetNotificationParametersDoubleOutOfRange

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersDoubleOutOfRangeBuilder) = (*_BACnetNotificationParametersDoubleOutOfRangeBuilder)(nil)

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagDouble, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagDouble, exceededLimit BACnetContextTagDouble, innerClosingTag BACnetClosingTag) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithExceedingValue(exceedingValue).WithStatusFlags(statusFlags).WithDeadband(deadband).WithExceededLimit(exceededLimit).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithExceedingValue(exceedingValue BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.ExceedingValue = exceedingValue
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithExceedingValueBuilder(builderSupplier func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.ExceedingValue.CreateBACnetContextTagDoubleBuilder())
	var err error
	b.ExceedingValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithDeadband(deadband BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetContextTagDoubleBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithExceededLimit(exceededLimit BACnetContextTagDouble) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.ExceededLimit = exceededLimit
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithExceededLimitBuilder(builderSupplier func(BACnetContextTagDoubleBuilder) BACnetContextTagDoubleBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.ExceededLimit.CreateBACnetContextTagDoubleBuilder())
	var err error
	b.ExceededLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersDoubleOutOfRangeBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) Build() (BACnetNotificationParametersDoubleOutOfRange, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.ExceedingValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'exceedingValue' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if b.ExceededLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'exceededLimit' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersDoubleOutOfRange.deepCopy(), nil
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) MustBuild() BACnetNotificationParametersDoubleOutOfRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) Done() BACnetNotificationParametersBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersBuilder().(*_BACnetNotificationParametersBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersDoubleOutOfRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersDoubleOutOfRangeBuilder().(*_BACnetNotificationParametersDoubleOutOfRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersDoubleOutOfRangeBuilder creates a BACnetNotificationParametersDoubleOutOfRangeBuilder
func (b *_BACnetNotificationParametersDoubleOutOfRange) CreateBACnetNotificationParametersDoubleOutOfRangeBuilder() BACnetNotificationParametersDoubleOutOfRangeBuilder {
	if b == nil {
		return NewBACnetNotificationParametersDoubleOutOfRangeBuilder()
	}
	return &_BACnetNotificationParametersDoubleOutOfRangeBuilder{_BACnetNotificationParametersDoubleOutOfRange: b.deepCopy()}
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

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetExceedingValue() BACnetContextTagDouble {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetDeadband() BACnetContextTagDouble {
	return m.Deadband
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetExceededLimit() BACnetContextTagDouble {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersDoubleOutOfRange(structType any) BACnetNotificationParametersDoubleOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersDoubleOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersDoubleOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersDoubleOutOfRange"
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersDoubleOutOfRange BACnetNotificationParametersDoubleOutOfRange, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersDoubleOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersDoubleOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	exceedingValue, err := ReadSimpleField[BACnetContextTagDouble](ctx, "exceedingValue", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceedingValue' field"))
	}
	m.ExceedingValue = exceedingValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	deadband, err := ReadSimpleField[BACnetContextTagDouble](ctx, "deadband", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	exceededLimit, err := ReadSimpleField[BACnetContextTagDouble](ctx, "exceededLimit", ReadComplex[BACnetContextTagDouble](BACnetContextTagParseWithBufferProducer[BACnetContextTagDouble]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_DOUBLE)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}
	m.ExceededLimit = exceededLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersDoubleOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersDoubleOutOfRange")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersDoubleOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersDoubleOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "exceedingValue", m.GetExceedingValue(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceedingValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetContextTagDouble](ctx, "exceededLimit", m.GetExceededLimit(), WriteComplex[BACnetContextTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceededLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersDoubleOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersDoubleOutOfRange")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) IsBACnetNotificationParametersDoubleOutOfRange() {
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) deepCopy() *_BACnetNotificationParametersDoubleOutOfRange {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersDoubleOutOfRangeCopy := &_BACnetNotificationParametersDoubleOutOfRange{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopy[BACnetContextTagDouble](m.ExceedingValue),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
		utils.DeepCopy[BACnetContextTagDouble](m.Deadband),
		utils.DeepCopy[BACnetContextTagDouble](m.ExceededLimit),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersDoubleOutOfRangeCopy
}

func (m *_BACnetNotificationParametersDoubleOutOfRange) String() string {
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
