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

// BACnetConstructedDataIntegerValueDeadband is the corresponding interface of BACnetConstructedDataIntegerValueDeadband
type BACnetConstructedDataIntegerValueDeadband interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataIntegerValueDeadband is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIntegerValueDeadband()
	// CreateBuilder creates a BACnetConstructedDataIntegerValueDeadbandBuilder
	CreateBACnetConstructedDataIntegerValueDeadbandBuilder() BACnetConstructedDataIntegerValueDeadbandBuilder
}

// _BACnetConstructedDataIntegerValueDeadband is the data-structure of this message
type _BACnetConstructedDataIntegerValueDeadband struct {
	BACnetConstructedDataContract
	Deadband BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataIntegerValueDeadband = (*_BACnetConstructedDataIntegerValueDeadband)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIntegerValueDeadband)(nil)

// NewBACnetConstructedDataIntegerValueDeadband factory function for _BACnetConstructedDataIntegerValueDeadband
func NewBACnetConstructedDataIntegerValueDeadband(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, deadband BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIntegerValueDeadband {
	if deadband == nil {
		panic("deadband of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataIntegerValueDeadband must not be nil")
	}
	_result := &_BACnetConstructedDataIntegerValueDeadband{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		Deadband:                      deadband,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIntegerValueDeadbandBuilder is a builder for BACnetConstructedDataIntegerValueDeadband
type BACnetConstructedDataIntegerValueDeadbandBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deadband BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntegerValueDeadbandBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntegerValueDeadbandBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIntegerValueDeadbandBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIntegerValueDeadband or returns an error if something is wrong
	Build() (BACnetConstructedDataIntegerValueDeadband, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIntegerValueDeadband
}

// NewBACnetConstructedDataIntegerValueDeadbandBuilder() creates a BACnetConstructedDataIntegerValueDeadbandBuilder
func NewBACnetConstructedDataIntegerValueDeadbandBuilder() BACnetConstructedDataIntegerValueDeadbandBuilder {
	return &_BACnetConstructedDataIntegerValueDeadbandBuilder{_BACnetConstructedDataIntegerValueDeadband: new(_BACnetConstructedDataIntegerValueDeadband)}
}

type _BACnetConstructedDataIntegerValueDeadbandBuilder struct {
	*_BACnetConstructedDataIntegerValueDeadband

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIntegerValueDeadbandBuilder) = (*_BACnetConstructedDataIntegerValueDeadbandBuilder)(nil)

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) WithMandatoryFields(deadband BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntegerValueDeadbandBuilder {
	return b.WithDeadband(deadband)
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) WithDeadband(deadband BACnetApplicationTagUnsignedInteger) BACnetConstructedDataIntegerValueDeadbandBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) WithDeadbandBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataIntegerValueDeadbandBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) Build() (BACnetConstructedDataIntegerValueDeadband, error) {
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIntegerValueDeadband.deepCopy(), nil
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) MustBuild() BACnetConstructedDataIntegerValueDeadband {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIntegerValueDeadbandBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIntegerValueDeadbandBuilder().(*_BACnetConstructedDataIntegerValueDeadbandBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIntegerValueDeadbandBuilder creates a BACnetConstructedDataIntegerValueDeadbandBuilder
func (b *_BACnetConstructedDataIntegerValueDeadband) CreateBACnetConstructedDataIntegerValueDeadbandBuilder() BACnetConstructedDataIntegerValueDeadbandBuilder {
	if b == nil {
		return NewBACnetConstructedDataIntegerValueDeadbandBuilder()
	}
	return &_BACnetConstructedDataIntegerValueDeadbandBuilder{_BACnetConstructedDataIntegerValueDeadband: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueDeadband) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_INTEGER_VALUE
}

func (m *_BACnetConstructedDataIntegerValueDeadband) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEADBAND
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegerValueDeadband) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueDeadband) GetDeadband() BACnetApplicationTagUnsignedInteger {
	return m.Deadband
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegerValueDeadband) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetDeadband())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegerValueDeadband(structType any) BACnetConstructedDataIntegerValueDeadband {
	if casted, ok := structType.(BACnetConstructedDataIntegerValueDeadband); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegerValueDeadband); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegerValueDeadband) GetTypeName() string {
	return "BACnetConstructedDataIntegerValueDeadband"
}

func (m *_BACnetConstructedDataIntegerValueDeadband) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIntegerValueDeadband) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIntegerValueDeadband) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIntegerValueDeadband BACnetConstructedDataIntegerValueDeadband, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegerValueDeadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegerValueDeadband")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deadband, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), deadband)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegerValueDeadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegerValueDeadband")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIntegerValueDeadband) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIntegerValueDeadband) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegerValueDeadband"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegerValueDeadband")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegerValueDeadband"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegerValueDeadband")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIntegerValueDeadband) IsBACnetConstructedDataIntegerValueDeadband() {}

func (m *_BACnetConstructedDataIntegerValueDeadband) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIntegerValueDeadband) deepCopy() *_BACnetConstructedDataIntegerValueDeadband {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIntegerValueDeadbandCopy := &_BACnetConstructedDataIntegerValueDeadband{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.Deadband),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIntegerValueDeadbandCopy
}

func (m *_BACnetConstructedDataIntegerValueDeadband) String() string {
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
