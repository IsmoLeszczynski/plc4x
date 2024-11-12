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

// BACnetConstructedDataEventType is the corresponding interface of BACnetConstructedDataEventType
type BACnetConstructedDataEventType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTypeTagged
	// IsBACnetConstructedDataEventType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventType()
	// CreateBuilder creates a BACnetConstructedDataEventTypeBuilder
	CreateBACnetConstructedDataEventTypeBuilder() BACnetConstructedDataEventTypeBuilder
}

// _BACnetConstructedDataEventType is the data-structure of this message
type _BACnetConstructedDataEventType struct {
	BACnetConstructedDataContract
	EventType BACnetEventTypeTagged
}

var _ BACnetConstructedDataEventType = (*_BACnetConstructedDataEventType)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventType)(nil)

// NewBACnetConstructedDataEventType factory function for _BACnetConstructedDataEventType
func NewBACnetConstructedDataEventType(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventType BACnetEventTypeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventType {
	if eventType == nil {
		panic("eventType of type BACnetEventTypeTagged for BACnetConstructedDataEventType must not be nil")
	}
	_result := &_BACnetConstructedDataEventType{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventType:                     eventType,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventTypeBuilder is a builder for BACnetConstructedDataEventType
type BACnetConstructedDataEventTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventType BACnetEventTypeTagged) BACnetConstructedDataEventTypeBuilder
	// WithEventType adds EventType (property field)
	WithEventType(BACnetEventTypeTagged) BACnetConstructedDataEventTypeBuilder
	// WithEventTypeBuilder adds EventType (property field) which is build by the builder
	WithEventTypeBuilder(func(BACnetEventTypeTaggedBuilder) BACnetEventTypeTaggedBuilder) BACnetConstructedDataEventTypeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEventType or returns an error if something is wrong
	Build() (BACnetConstructedDataEventType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventType
}

// NewBACnetConstructedDataEventTypeBuilder() creates a BACnetConstructedDataEventTypeBuilder
func NewBACnetConstructedDataEventTypeBuilder() BACnetConstructedDataEventTypeBuilder {
	return &_BACnetConstructedDataEventTypeBuilder{_BACnetConstructedDataEventType: new(_BACnetConstructedDataEventType)}
}

type _BACnetConstructedDataEventTypeBuilder struct {
	*_BACnetConstructedDataEventType

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventTypeBuilder) = (*_BACnetConstructedDataEventTypeBuilder)(nil)

func (b *_BACnetConstructedDataEventTypeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataEventTypeBuilder) WithMandatoryFields(eventType BACnetEventTypeTagged) BACnetConstructedDataEventTypeBuilder {
	return b.WithEventType(eventType)
}

func (b *_BACnetConstructedDataEventTypeBuilder) WithEventType(eventType BACnetEventTypeTagged) BACnetConstructedDataEventTypeBuilder {
	b.EventType = eventType
	return b
}

func (b *_BACnetConstructedDataEventTypeBuilder) WithEventTypeBuilder(builderSupplier func(BACnetEventTypeTaggedBuilder) BACnetEventTypeTaggedBuilder) BACnetConstructedDataEventTypeBuilder {
	builder := builderSupplier(b.EventType.CreateBACnetEventTypeTaggedBuilder())
	var err error
	b.EventType, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventTypeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEventTypeBuilder) Build() (BACnetConstructedDataEventType, error) {
	if b.EventType == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventType' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEventType.deepCopy(), nil
}

func (b *_BACnetConstructedDataEventTypeBuilder) MustBuild() BACnetConstructedDataEventType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEventTypeBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEventTypeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEventTypeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEventTypeBuilder().(*_BACnetConstructedDataEventTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEventTypeBuilder creates a BACnetConstructedDataEventTypeBuilder
func (b *_BACnetConstructedDataEventType) CreateBACnetConstructedDataEventTypeBuilder() BACnetConstructedDataEventTypeBuilder {
	if b == nil {
		return NewBACnetConstructedDataEventTypeBuilder()
	}
	return &_BACnetConstructedDataEventTypeBuilder{_BACnetConstructedDataEventType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventType) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventType) GetActualValue() BACnetEventTypeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventTypeTagged(m.GetEventType())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventType(structType any) BACnetConstructedDataEventType {
	if casted, ok := structType.(BACnetConstructedDataEventType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventType) GetTypeName() string {
	return "BACnetConstructedDataEventType"
}

func (m *_BACnetConstructedDataEventType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventType BACnetConstructedDataEventType, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}
	m.EventType = eventType

	actualValue, err := ReadVirtualField[BACnetEventTypeTagged](ctx, "actualValue", (*BACnetEventTypeTagged)(nil), eventType)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventType")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventType")
		}

		if err := WriteSimpleField[BACnetEventTypeTagged](ctx, "eventType", m.GetEventType(), WriteComplex[BACnetEventTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventType' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventType")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventType) IsBACnetConstructedDataEventType() {}

func (m *_BACnetConstructedDataEventType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventType) deepCopy() *_BACnetConstructedDataEventType {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventTypeCopy := &_BACnetConstructedDataEventType{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEventTypeTagged](m.EventType),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventTypeCopy
}

func (m *_BACnetConstructedDataEventType) String() string {
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
