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

// BACnetConstructedDataInitialTimeout is the corresponding interface of BACnetConstructedDataInitialTimeout
type BACnetConstructedDataInitialTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetInitialTimeout returns InitialTimeout (property field)
	GetInitialTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataInitialTimeout is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataInitialTimeout()
	// CreateBuilder creates a BACnetConstructedDataInitialTimeoutBuilder
	CreateBACnetConstructedDataInitialTimeoutBuilder() BACnetConstructedDataInitialTimeoutBuilder
}

// _BACnetConstructedDataInitialTimeout is the data-structure of this message
type _BACnetConstructedDataInitialTimeout struct {
	BACnetConstructedDataContract
	InitialTimeout BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataInitialTimeout = (*_BACnetConstructedDataInitialTimeout)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataInitialTimeout)(nil)

// NewBACnetConstructedDataInitialTimeout factory function for _BACnetConstructedDataInitialTimeout
func NewBACnetConstructedDataInitialTimeout(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, initialTimeout BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInitialTimeout {
	if initialTimeout == nil {
		panic("initialTimeout of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataInitialTimeout must not be nil")
	}
	_result := &_BACnetConstructedDataInitialTimeout{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		InitialTimeout:                initialTimeout,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataInitialTimeoutBuilder is a builder for BACnetConstructedDataInitialTimeout
type BACnetConstructedDataInitialTimeoutBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(initialTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataInitialTimeoutBuilder
	// WithInitialTimeout adds InitialTimeout (property field)
	WithInitialTimeout(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataInitialTimeoutBuilder
	// WithInitialTimeoutBuilder adds InitialTimeout (property field) which is build by the builder
	WithInitialTimeoutBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataInitialTimeoutBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataInitialTimeout or returns an error if something is wrong
	Build() (BACnetConstructedDataInitialTimeout, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataInitialTimeout
}

// NewBACnetConstructedDataInitialTimeoutBuilder() creates a BACnetConstructedDataInitialTimeoutBuilder
func NewBACnetConstructedDataInitialTimeoutBuilder() BACnetConstructedDataInitialTimeoutBuilder {
	return &_BACnetConstructedDataInitialTimeoutBuilder{_BACnetConstructedDataInitialTimeout: new(_BACnetConstructedDataInitialTimeout)}
}

type _BACnetConstructedDataInitialTimeoutBuilder struct {
	*_BACnetConstructedDataInitialTimeout

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataInitialTimeoutBuilder) = (*_BACnetConstructedDataInitialTimeoutBuilder)(nil)

func (b *_BACnetConstructedDataInitialTimeoutBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) WithMandatoryFields(initialTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataInitialTimeoutBuilder {
	return b.WithInitialTimeout(initialTimeout)
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) WithInitialTimeout(initialTimeout BACnetApplicationTagUnsignedInteger) BACnetConstructedDataInitialTimeoutBuilder {
	b.InitialTimeout = initialTimeout
	return b
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) WithInitialTimeoutBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataInitialTimeoutBuilder {
	builder := builderSupplier(b.InitialTimeout.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.InitialTimeout, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) Build() (BACnetConstructedDataInitialTimeout, error) {
	if b.InitialTimeout == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'initialTimeout' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataInitialTimeout.deepCopy(), nil
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) MustBuild() BACnetConstructedDataInitialTimeout {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataInitialTimeoutBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataInitialTimeoutBuilder().(*_BACnetConstructedDataInitialTimeoutBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataInitialTimeoutBuilder creates a BACnetConstructedDataInitialTimeoutBuilder
func (b *_BACnetConstructedDataInitialTimeout) CreateBACnetConstructedDataInitialTimeoutBuilder() BACnetConstructedDataInitialTimeoutBuilder {
	if b == nil {
		return NewBACnetConstructedDataInitialTimeoutBuilder()
	}
	return &_BACnetConstructedDataInitialTimeoutBuilder{_BACnetConstructedDataInitialTimeout: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInitialTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INITIAL_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetInitialTimeout() BACnetApplicationTagUnsignedInteger {
	return m.InitialTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetInitialTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInitialTimeout(structType any) BACnetConstructedDataInitialTimeout {
	if casted, ok := structType.(BACnetConstructedDataInitialTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInitialTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInitialTimeout) GetTypeName() string {
	return "BACnetConstructedDataInitialTimeout"
}

func (m *_BACnetConstructedDataInitialTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (initialTimeout)
	lengthInBits += m.InitialTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInitialTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataInitialTimeout) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataInitialTimeout BACnetConstructedDataInitialTimeout, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInitialTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInitialTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	initialTimeout, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "initialTimeout", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initialTimeout' field"))
	}
	m.InitialTimeout = initialTimeout

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), initialTimeout)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInitialTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInitialTimeout")
	}

	return m, nil
}

func (m *_BACnetConstructedDataInitialTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInitialTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInitialTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInitialTimeout")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "initialTimeout", m.GetInitialTimeout(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initialTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInitialTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInitialTimeout")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInitialTimeout) IsBACnetConstructedDataInitialTimeout() {}

func (m *_BACnetConstructedDataInitialTimeout) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataInitialTimeout) deepCopy() *_BACnetConstructedDataInitialTimeout {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataInitialTimeoutCopy := &_BACnetConstructedDataInitialTimeout{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.InitialTimeout),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataInitialTimeoutCopy
}

func (m *_BACnetConstructedDataInitialTimeout) String() string {
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
