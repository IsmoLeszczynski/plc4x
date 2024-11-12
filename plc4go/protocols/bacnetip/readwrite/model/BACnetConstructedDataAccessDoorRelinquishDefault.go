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

// BACnetConstructedDataAccessDoorRelinquishDefault is the corresponding interface of BACnetConstructedDataAccessDoorRelinquishDefault
type BACnetConstructedDataAccessDoorRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetDoorValueTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorValueTagged
	// IsBACnetConstructedDataAccessDoorRelinquishDefault is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessDoorRelinquishDefault()
	// CreateBuilder creates a BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
	CreateBACnetConstructedDataAccessDoorRelinquishDefaultBuilder() BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
}

// _BACnetConstructedDataAccessDoorRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataAccessDoorRelinquishDefault struct {
	BACnetConstructedDataContract
	RelinquishDefault BACnetDoorValueTagged
}

var _ BACnetConstructedDataAccessDoorRelinquishDefault = (*_BACnetConstructedDataAccessDoorRelinquishDefault)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessDoorRelinquishDefault)(nil)

// NewBACnetConstructedDataAccessDoorRelinquishDefault factory function for _BACnetConstructedDataAccessDoorRelinquishDefault
func NewBACnetConstructedDataAccessDoorRelinquishDefault(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, relinquishDefault BACnetDoorValueTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessDoorRelinquishDefault {
	if relinquishDefault == nil {
		panic("relinquishDefault of type BACnetDoorValueTagged for BACnetConstructedDataAccessDoorRelinquishDefault must not be nil")
	}
	_result := &_BACnetConstructedDataAccessDoorRelinquishDefault{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		RelinquishDefault:             relinquishDefault,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessDoorRelinquishDefaultBuilder is a builder for BACnetConstructedDataAccessDoorRelinquishDefault
type BACnetConstructedDataAccessDoorRelinquishDefaultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(relinquishDefault BACnetDoorValueTagged) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
	// WithRelinquishDefault adds RelinquishDefault (property field)
	WithRelinquishDefault(BACnetDoorValueTagged) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
	// WithRelinquishDefaultBuilder adds RelinquishDefault (property field) which is build by the builder
	WithRelinquishDefaultBuilder(func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessDoorRelinquishDefault or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessDoorRelinquishDefault, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessDoorRelinquishDefault
}

// NewBACnetConstructedDataAccessDoorRelinquishDefaultBuilder() creates a BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
func NewBACnetConstructedDataAccessDoorRelinquishDefaultBuilder() BACnetConstructedDataAccessDoorRelinquishDefaultBuilder {
	return &_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder{_BACnetConstructedDataAccessDoorRelinquishDefault: new(_BACnetConstructedDataAccessDoorRelinquishDefault)}
}

type _BACnetConstructedDataAccessDoorRelinquishDefaultBuilder struct {
	*_BACnetConstructedDataAccessDoorRelinquishDefault

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) = (*_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder)(nil)

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) WithMandatoryFields(relinquishDefault BACnetDoorValueTagged) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder {
	return b.WithRelinquishDefault(relinquishDefault)
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) WithRelinquishDefault(relinquishDefault BACnetDoorValueTagged) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder {
	b.RelinquishDefault = relinquishDefault
	return b
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) WithRelinquishDefaultBuilder(builderSupplier func(BACnetDoorValueTaggedBuilder) BACnetDoorValueTaggedBuilder) BACnetConstructedDataAccessDoorRelinquishDefaultBuilder {
	builder := builderSupplier(b.RelinquishDefault.CreateBACnetDoorValueTaggedBuilder())
	var err error
	b.RelinquishDefault, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDoorValueTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) Build() (BACnetConstructedDataAccessDoorRelinquishDefault, error) {
	if b.RelinquishDefault == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'relinquishDefault' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessDoorRelinquishDefault.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) MustBuild() BACnetConstructedDataAccessDoorRelinquishDefault {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessDoorRelinquishDefaultBuilder().(*_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessDoorRelinquishDefaultBuilder creates a BACnetConstructedDataAccessDoorRelinquishDefaultBuilder
func (b *_BACnetConstructedDataAccessDoorRelinquishDefault) CreateBACnetConstructedDataAccessDoorRelinquishDefaultBuilder() BACnetConstructedDataAccessDoorRelinquishDefaultBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessDoorRelinquishDefaultBuilder()
	}
	return &_BACnetConstructedDataAccessDoorRelinquishDefaultBuilder{_BACnetConstructedDataAccessDoorRelinquishDefault: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ACCESS_DOOR
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetRelinquishDefault() BACnetDoorValueTagged {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetActualValue() BACnetDoorValueTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorValueTagged(m.GetRelinquishDefault())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessDoorRelinquishDefault(structType any) BACnetConstructedDataAccessDoorRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataAccessDoorRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessDoorRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataAccessDoorRelinquishDefault"
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessDoorRelinquishDefault BACnetConstructedDataAccessDoorRelinquishDefault, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessDoorRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessDoorRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relinquishDefault, err := ReadSimpleField[BACnetDoorValueTagged](ctx, "relinquishDefault", ReadComplex[BACnetDoorValueTagged](BACnetDoorValueTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relinquishDefault' field"))
	}
	m.RelinquishDefault = relinquishDefault

	actualValue, err := ReadVirtualField[BACnetDoorValueTagged](ctx, "actualValue", (*BACnetDoorValueTagged)(nil), relinquishDefault)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessDoorRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessDoorRelinquishDefault")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessDoorRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessDoorRelinquishDefault")
		}

		if err := WriteSimpleField[BACnetDoorValueTagged](ctx, "relinquishDefault", m.GetRelinquishDefault(), WriteComplex[BACnetDoorValueTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relinquishDefault' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessDoorRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessDoorRelinquishDefault")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) IsBACnetConstructedDataAccessDoorRelinquishDefault() {
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) deepCopy() *_BACnetConstructedDataAccessDoorRelinquishDefault {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessDoorRelinquishDefaultCopy := &_BACnetConstructedDataAccessDoorRelinquishDefault{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDoorValueTagged](m.RelinquishDefault),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessDoorRelinquishDefaultCopy
}

func (m *_BACnetConstructedDataAccessDoorRelinquishDefault) String() string {
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
