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

// BACnetConstructedDataIsUTC is the corresponding interface of BACnetConstructedDataIsUTC
type BACnetConstructedDataIsUTC interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetIsUtc returns IsUtc (property field)
	GetIsUtc() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataIsUTC is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataIsUTC()
	// CreateBuilder creates a BACnetConstructedDataIsUTCBuilder
	CreateBACnetConstructedDataIsUTCBuilder() BACnetConstructedDataIsUTCBuilder
}

// _BACnetConstructedDataIsUTC is the data-structure of this message
type _BACnetConstructedDataIsUTC struct {
	BACnetConstructedDataContract
	IsUtc BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataIsUTC = (*_BACnetConstructedDataIsUTC)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataIsUTC)(nil)

// NewBACnetConstructedDataIsUTC factory function for _BACnetConstructedDataIsUTC
func NewBACnetConstructedDataIsUTC(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, isUtc BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIsUTC {
	if isUtc == nil {
		panic("isUtc of type BACnetApplicationTagBoolean for BACnetConstructedDataIsUTC must not be nil")
	}
	_result := &_BACnetConstructedDataIsUTC{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		IsUtc:                         isUtc,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataIsUTCBuilder is a builder for BACnetConstructedDataIsUTC
type BACnetConstructedDataIsUTCBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(isUtc BACnetApplicationTagBoolean) BACnetConstructedDataIsUTCBuilder
	// WithIsUtc adds IsUtc (property field)
	WithIsUtc(BACnetApplicationTagBoolean) BACnetConstructedDataIsUTCBuilder
	// WithIsUtcBuilder adds IsUtc (property field) which is build by the builder
	WithIsUtcBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIsUTCBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataIsUTC or returns an error if something is wrong
	Build() (BACnetConstructedDataIsUTC, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataIsUTC
}

// NewBACnetConstructedDataIsUTCBuilder() creates a BACnetConstructedDataIsUTCBuilder
func NewBACnetConstructedDataIsUTCBuilder() BACnetConstructedDataIsUTCBuilder {
	return &_BACnetConstructedDataIsUTCBuilder{_BACnetConstructedDataIsUTC: new(_BACnetConstructedDataIsUTC)}
}

type _BACnetConstructedDataIsUTCBuilder struct {
	*_BACnetConstructedDataIsUTC

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataIsUTCBuilder) = (*_BACnetConstructedDataIsUTCBuilder)(nil)

func (b *_BACnetConstructedDataIsUTCBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataIsUTCBuilder) WithMandatoryFields(isUtc BACnetApplicationTagBoolean) BACnetConstructedDataIsUTCBuilder {
	return b.WithIsUtc(isUtc)
}

func (b *_BACnetConstructedDataIsUTCBuilder) WithIsUtc(isUtc BACnetApplicationTagBoolean) BACnetConstructedDataIsUTCBuilder {
	b.IsUtc = isUtc
	return b
}

func (b *_BACnetConstructedDataIsUTCBuilder) WithIsUtcBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataIsUTCBuilder {
	builder := builderSupplier(b.IsUtc.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.IsUtc, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataIsUTCBuilder) Build() (BACnetConstructedDataIsUTC, error) {
	if b.IsUtc == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'isUtc' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataIsUTC.deepCopy(), nil
}

func (b *_BACnetConstructedDataIsUTCBuilder) MustBuild() BACnetConstructedDataIsUTC {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataIsUTCBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataIsUTCBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataIsUTCBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataIsUTCBuilder().(*_BACnetConstructedDataIsUTCBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataIsUTCBuilder creates a BACnetConstructedDataIsUTCBuilder
func (b *_BACnetConstructedDataIsUTC) CreateBACnetConstructedDataIsUTCBuilder() BACnetConstructedDataIsUTCBuilder {
	if b == nil {
		return NewBACnetConstructedDataIsUTCBuilder()
	}
	return &_BACnetConstructedDataIsUTCBuilder{_BACnetConstructedDataIsUTC: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIsUTC) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIsUTC) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IS_UTC
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIsUTC) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIsUTC) GetIsUtc() BACnetApplicationTagBoolean {
	return m.IsUtc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIsUTC) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetIsUtc())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIsUTC(structType any) BACnetConstructedDataIsUTC {
	if casted, ok := structType.(BACnetConstructedDataIsUTC); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIsUTC); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIsUTC) GetTypeName() string {
	return "BACnetConstructedDataIsUTC"
}

func (m *_BACnetConstructedDataIsUTC) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (isUtc)
	lengthInBits += m.IsUtc.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIsUTC) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataIsUTC) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataIsUTC BACnetConstructedDataIsUTC, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIsUTC"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIsUTC")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	isUtc, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "isUtc", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUtc' field"))
	}
	m.IsUtc = isUtc

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), isUtc)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIsUTC"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIsUTC")
	}

	return m, nil
}

func (m *_BACnetConstructedDataIsUTC) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIsUTC) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIsUTC"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIsUTC")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "isUtc", m.GetIsUtc(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isUtc' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIsUTC"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIsUTC")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIsUTC) IsBACnetConstructedDataIsUTC() {}

func (m *_BACnetConstructedDataIsUTC) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataIsUTC) deepCopy() *_BACnetConstructedDataIsUTC {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataIsUTCCopy := &_BACnetConstructedDataIsUTC{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.IsUtc),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataIsUTCCopy
}

func (m *_BACnetConstructedDataIsUTC) String() string {
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
