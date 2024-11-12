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

// BACnetConstructedDataLargeAnalogValueFaultHighLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueFaultHighLimit
type BACnetConstructedDataLargeAnalogValueFaultHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
	// IsBACnetConstructedDataLargeAnalogValueFaultHighLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLargeAnalogValueFaultHighLimit()
	// CreateBuilder creates a BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
	CreateBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
}

// _BACnetConstructedDataLargeAnalogValueFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueFaultHighLimit struct {
	BACnetConstructedDataContract
	FaultHighLimit BACnetApplicationTagDouble
}

var _ BACnetConstructedDataLargeAnalogValueFaultHighLimit = (*_BACnetConstructedDataLargeAnalogValueFaultHighLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLargeAnalogValueFaultHighLimit)(nil)

// NewBACnetConstructedDataLargeAnalogValueFaultHighLimit factory function for _BACnetConstructedDataLargeAnalogValueFaultHighLimit
func NewBACnetConstructedDataLargeAnalogValueFaultHighLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultHighLimit BACnetApplicationTagDouble, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	if faultHighLimit == nil {
		panic("faultHighLimit of type BACnetApplicationTagDouble for BACnetConstructedDataLargeAnalogValueFaultHighLimit must not be nil")
	}
	_result := &_BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultHighLimit:                faultHighLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder is a builder for BACnetConstructedDataLargeAnalogValueFaultHighLimit
type BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultHighLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
	// WithFaultHighLimit adds FaultHighLimit (property field)
	WithFaultHighLimit(BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
	// WithFaultHighLimitBuilder adds FaultHighLimit (property field) which is build by the builder
	WithFaultHighLimitBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLargeAnalogValueFaultHighLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataLargeAnalogValueFaultHighLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLargeAnalogValueFaultHighLimit
}

// NewBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder() creates a BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
func NewBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder {
	return &_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder{_BACnetConstructedDataLargeAnalogValueFaultHighLimit: new(_BACnetConstructedDataLargeAnalogValueFaultHighLimit)}
}

type _BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder struct {
	*_BACnetConstructedDataLargeAnalogValueFaultHighLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) = (*_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder)(nil)

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) WithMandatoryFields(faultHighLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder {
	return b.WithFaultHighLimit(faultHighLimit)
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) WithFaultHighLimit(faultHighLimit BACnetApplicationTagDouble) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder {
	b.FaultHighLimit = faultHighLimit
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) WithFaultHighLimitBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder {
	builder := builderSupplier(b.FaultHighLimit.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.FaultHighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) Build() (BACnetConstructedDataLargeAnalogValueFaultHighLimit, error) {
	if b.FaultHighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'faultHighLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLargeAnalogValueFaultHighLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) MustBuild() BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder().(*_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder creates a BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder
func (b *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) CreateBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder() BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder()
	}
	return &_BACnetConstructedDataLargeAnalogValueFaultHighLimitBuilder{_BACnetConstructedDataLargeAnalogValueFaultHighLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LARGE_ANALOG_VALUE
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagDouble {
	return m.FaultHighLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetActualValue() BACnetApplicationTagDouble {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDouble(m.GetFaultHighLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueFaultHighLimit(structType any) BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueFaultHighLimit"
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLargeAnalogValueFaultHighLimit BACnetConstructedDataLargeAnalogValueFaultHighLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultHighLimit, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "faultHighLimit", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultHighLimit' field"))
	}
	m.FaultHighLimit = faultHighLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagDouble](ctx, "actualValue", (*BACnetApplicationTagDouble)(nil), faultHighLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "faultHighLimit", m.GetFaultHighLimit(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultHighLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueFaultHighLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) IsBACnetConstructedDataLargeAnalogValueFaultHighLimit() {
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) deepCopy() *_BACnetConstructedDataLargeAnalogValueFaultHighLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLargeAnalogValueFaultHighLimitCopy := &_BACnetConstructedDataLargeAnalogValueFaultHighLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.FaultHighLimit),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLargeAnalogValueFaultHighLimitCopy
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultHighLimit) String() string {
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
