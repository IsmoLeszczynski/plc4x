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

// BACnetConstructedDataSerialNumber is the corresponding interface of BACnetConstructedDataSerialNumber
type BACnetConstructedDataSerialNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataSerialNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataSerialNumber()
	// CreateBuilder creates a BACnetConstructedDataSerialNumberBuilder
	CreateBACnetConstructedDataSerialNumberBuilder() BACnetConstructedDataSerialNumberBuilder
}

// _BACnetConstructedDataSerialNumber is the data-structure of this message
type _BACnetConstructedDataSerialNumber struct {
	BACnetConstructedDataContract
	SerialNumber BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataSerialNumber = (*_BACnetConstructedDataSerialNumber)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataSerialNumber)(nil)

// NewBACnetConstructedDataSerialNumber factory function for _BACnetConstructedDataSerialNumber
func NewBACnetConstructedDataSerialNumber(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, serialNumber BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSerialNumber {
	if serialNumber == nil {
		panic("serialNumber of type BACnetApplicationTagCharacterString for BACnetConstructedDataSerialNumber must not be nil")
	}
	_result := &_BACnetConstructedDataSerialNumber{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		SerialNumber:                  serialNumber,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataSerialNumberBuilder is a builder for BACnetConstructedDataSerialNumber
type BACnetConstructedDataSerialNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(serialNumber BACnetApplicationTagCharacterString) BACnetConstructedDataSerialNumberBuilder
	// WithSerialNumber adds SerialNumber (property field)
	WithSerialNumber(BACnetApplicationTagCharacterString) BACnetConstructedDataSerialNumberBuilder
	// WithSerialNumberBuilder adds SerialNumber (property field) which is build by the builder
	WithSerialNumberBuilder(func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataSerialNumberBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataSerialNumber or returns an error if something is wrong
	Build() (BACnetConstructedDataSerialNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataSerialNumber
}

// NewBACnetConstructedDataSerialNumberBuilder() creates a BACnetConstructedDataSerialNumberBuilder
func NewBACnetConstructedDataSerialNumberBuilder() BACnetConstructedDataSerialNumberBuilder {
	return &_BACnetConstructedDataSerialNumberBuilder{_BACnetConstructedDataSerialNumber: new(_BACnetConstructedDataSerialNumber)}
}

type _BACnetConstructedDataSerialNumberBuilder struct {
	*_BACnetConstructedDataSerialNumber

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataSerialNumberBuilder) = (*_BACnetConstructedDataSerialNumberBuilder)(nil)

func (b *_BACnetConstructedDataSerialNumberBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataSerialNumberBuilder) WithMandatoryFields(serialNumber BACnetApplicationTagCharacterString) BACnetConstructedDataSerialNumberBuilder {
	return b.WithSerialNumber(serialNumber)
}

func (b *_BACnetConstructedDataSerialNumberBuilder) WithSerialNumber(serialNumber BACnetApplicationTagCharacterString) BACnetConstructedDataSerialNumberBuilder {
	b.SerialNumber = serialNumber
	return b
}

func (b *_BACnetConstructedDataSerialNumberBuilder) WithSerialNumberBuilder(builderSupplier func(BACnetApplicationTagCharacterStringBuilder) BACnetApplicationTagCharacterStringBuilder) BACnetConstructedDataSerialNumberBuilder {
	builder := builderSupplier(b.SerialNumber.CreateBACnetApplicationTagCharacterStringBuilder())
	var err error
	b.SerialNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagCharacterStringBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataSerialNumberBuilder) Build() (BACnetConstructedDataSerialNumber, error) {
	if b.SerialNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'serialNumber' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataSerialNumber.deepCopy(), nil
}

func (b *_BACnetConstructedDataSerialNumberBuilder) MustBuild() BACnetConstructedDataSerialNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataSerialNumberBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataSerialNumberBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataSerialNumberBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataSerialNumberBuilder().(*_BACnetConstructedDataSerialNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataSerialNumberBuilder creates a BACnetConstructedDataSerialNumberBuilder
func (b *_BACnetConstructedDataSerialNumber) CreateBACnetConstructedDataSerialNumberBuilder() BACnetConstructedDataSerialNumberBuilder {
	if b == nil {
		return NewBACnetConstructedDataSerialNumberBuilder()
	}
	return &_BACnetConstructedDataSerialNumberBuilder{_BACnetConstructedDataSerialNumber: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSerialNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SERIAL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetSerialNumber() BACnetApplicationTagCharacterString {
	return m.SerialNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetSerialNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSerialNumber(structType any) BACnetConstructedDataSerialNumber {
	if casted, ok := structType.(BACnetConstructedDataSerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSerialNumber) GetTypeName() string {
	return "BACnetConstructedDataSerialNumber"
}

func (m *_BACnetConstructedDataSerialNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (serialNumber)
	lengthInBits += m.SerialNumber.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSerialNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataSerialNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataSerialNumber BACnetConstructedDataSerialNumber, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serialNumber, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "serialNumber", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serialNumber' field"))
	}
	m.SerialNumber = serialNumber

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), serialNumber)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSerialNumber")
	}

	return m, nil
}

func (m *_BACnetConstructedDataSerialNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSerialNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSerialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSerialNumber")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "serialNumber", m.GetSerialNumber(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serialNumber' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSerialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSerialNumber")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSerialNumber) IsBACnetConstructedDataSerialNumber() {}

func (m *_BACnetConstructedDataSerialNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataSerialNumber) deepCopy() *_BACnetConstructedDataSerialNumber {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataSerialNumberCopy := &_BACnetConstructedDataSerialNumber{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagCharacterString](m.SerialNumber),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataSerialNumberCopy
}

func (m *_BACnetConstructedDataSerialNumber) String() string {
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
