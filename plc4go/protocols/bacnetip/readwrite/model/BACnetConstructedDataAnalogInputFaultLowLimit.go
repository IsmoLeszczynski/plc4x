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

// BACnetConstructedDataAnalogInputFaultLowLimit is the corresponding interface of BACnetConstructedDataAnalogInputFaultLowLimit
type BACnetConstructedDataAnalogInputFaultLowLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataAnalogInputFaultLowLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogInputFaultLowLimit()
	// CreateBuilder creates a BACnetConstructedDataAnalogInputFaultLowLimitBuilder
	CreateBACnetConstructedDataAnalogInputFaultLowLimitBuilder() BACnetConstructedDataAnalogInputFaultLowLimitBuilder
}

// _BACnetConstructedDataAnalogInputFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataAnalogInputFaultLowLimit struct {
	BACnetConstructedDataContract
	FaultLowLimit BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogInputFaultLowLimit = (*_BACnetConstructedDataAnalogInputFaultLowLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogInputFaultLowLimit)(nil)

// NewBACnetConstructedDataAnalogInputFaultLowLimit factory function for _BACnetConstructedDataAnalogInputFaultLowLimit
func NewBACnetConstructedDataAnalogInputFaultLowLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, faultLowLimit BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogInputFaultLowLimit {
	if faultLowLimit == nil {
		panic("faultLowLimit of type BACnetApplicationTagReal for BACnetConstructedDataAnalogInputFaultLowLimit must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogInputFaultLowLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FaultLowLimit:                 faultLowLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogInputFaultLowLimitBuilder is a builder for BACnetConstructedDataAnalogInputFaultLowLimit
type BACnetConstructedDataAnalogInputFaultLowLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogInputFaultLowLimitBuilder
	// WithFaultLowLimit adds FaultLowLimit (property field)
	WithFaultLowLimit(BACnetApplicationTagReal) BACnetConstructedDataAnalogInputFaultLowLimitBuilder
	// WithFaultLowLimitBuilder adds FaultLowLimit (property field) which is build by the builder
	WithFaultLowLimitBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogInputFaultLowLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAnalogInputFaultLowLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogInputFaultLowLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogInputFaultLowLimit
}

// NewBACnetConstructedDataAnalogInputFaultLowLimitBuilder() creates a BACnetConstructedDataAnalogInputFaultLowLimitBuilder
func NewBACnetConstructedDataAnalogInputFaultLowLimitBuilder() BACnetConstructedDataAnalogInputFaultLowLimitBuilder {
	return &_BACnetConstructedDataAnalogInputFaultLowLimitBuilder{_BACnetConstructedDataAnalogInputFaultLowLimit: new(_BACnetConstructedDataAnalogInputFaultLowLimit)}
}

type _BACnetConstructedDataAnalogInputFaultLowLimitBuilder struct {
	*_BACnetConstructedDataAnalogInputFaultLowLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogInputFaultLowLimitBuilder) = (*_BACnetConstructedDataAnalogInputFaultLowLimitBuilder)(nil)

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) WithMandatoryFields(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogInputFaultLowLimitBuilder {
	return b.WithFaultLowLimit(faultLowLimit)
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) WithFaultLowLimit(faultLowLimit BACnetApplicationTagReal) BACnetConstructedDataAnalogInputFaultLowLimitBuilder {
	b.FaultLowLimit = faultLowLimit
	return b
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) WithFaultLowLimitBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogInputFaultLowLimitBuilder {
	builder := builderSupplier(b.FaultLowLimit.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.FaultLowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) Build() (BACnetConstructedDataAnalogInputFaultLowLimit, error) {
	if b.FaultLowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'faultLowLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogInputFaultLowLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) MustBuild() BACnetConstructedDataAnalogInputFaultLowLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogInputFaultLowLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogInputFaultLowLimitBuilder().(*_BACnetConstructedDataAnalogInputFaultLowLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogInputFaultLowLimitBuilder creates a BACnetConstructedDataAnalogInputFaultLowLimitBuilder
func (b *_BACnetConstructedDataAnalogInputFaultLowLimit) CreateBACnetConstructedDataAnalogInputFaultLowLimitBuilder() BACnetConstructedDataAnalogInputFaultLowLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogInputFaultLowLimitBuilder()
	}
	return &_BACnetConstructedDataAnalogInputFaultLowLimitBuilder{_BACnetConstructedDataAnalogInputFaultLowLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_INPUT
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAULT_LOW_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagReal {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogInputFaultLowLimit(structType any) BACnetConstructedDataAnalogInputFaultLowLimit {
	if casted, ok := structType.(BACnetConstructedDataAnalogInputFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogInputFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataAnalogInputFaultLowLimit"
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogInputFaultLowLimit BACnetConstructedDataAnalogInputFaultLowLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogInputFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogInputFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	faultLowLimit, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "faultLowLimit", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'faultLowLimit' field"))
	}
	m.FaultLowLimit = faultLowLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), faultLowLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogInputFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogInputFaultLowLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogInputFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogInputFaultLowLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "faultLowLimit", m.GetFaultLowLimit(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'faultLowLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogInputFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogInputFaultLowLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) IsBACnetConstructedDataAnalogInputFaultLowLimit() {
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) deepCopy() *_BACnetConstructedDataAnalogInputFaultLowLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogInputFaultLowLimitCopy := &_BACnetConstructedDataAnalogInputFaultLowLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.FaultLowLimit),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogInputFaultLowLimitCopy
}

func (m *_BACnetConstructedDataAnalogInputFaultLowLimit) String() string {
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
