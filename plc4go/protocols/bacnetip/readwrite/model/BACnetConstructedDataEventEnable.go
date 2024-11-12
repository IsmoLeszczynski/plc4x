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

// BACnetConstructedDataEventEnable is the corresponding interface of BACnetConstructedDataEventEnable
type BACnetConstructedDataEventEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEventEnable returns EventEnable (property field)
	GetEventEnable() BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTransitionBitsTagged
	// IsBACnetConstructedDataEventEnable is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventEnable()
	// CreateBuilder creates a BACnetConstructedDataEventEnableBuilder
	CreateBACnetConstructedDataEventEnableBuilder() BACnetConstructedDataEventEnableBuilder
}

// _BACnetConstructedDataEventEnable is the data-structure of this message
type _BACnetConstructedDataEventEnable struct {
	BACnetConstructedDataContract
	EventEnable BACnetEventTransitionBitsTagged
}

var _ BACnetConstructedDataEventEnable = (*_BACnetConstructedDataEventEnable)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventEnable)(nil)

// NewBACnetConstructedDataEventEnable factory function for _BACnetConstructedDataEventEnable
func NewBACnetConstructedDataEventEnable(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, eventEnable BACnetEventTransitionBitsTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventEnable {
	if eventEnable == nil {
		panic("eventEnable of type BACnetEventTransitionBitsTagged for BACnetConstructedDataEventEnable must not be nil")
	}
	_result := &_BACnetConstructedDataEventEnable{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EventEnable:                   eventEnable,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEventEnableBuilder is a builder for BACnetConstructedDataEventEnable
type BACnetConstructedDataEventEnableBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(eventEnable BACnetEventTransitionBitsTagged) BACnetConstructedDataEventEnableBuilder
	// WithEventEnable adds EventEnable (property field)
	WithEventEnable(BACnetEventTransitionBitsTagged) BACnetConstructedDataEventEnableBuilder
	// WithEventEnableBuilder adds EventEnable (property field) which is build by the builder
	WithEventEnableBuilder(func(BACnetEventTransitionBitsTaggedBuilder) BACnetEventTransitionBitsTaggedBuilder) BACnetConstructedDataEventEnableBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEventEnable or returns an error if something is wrong
	Build() (BACnetConstructedDataEventEnable, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEventEnable
}

// NewBACnetConstructedDataEventEnableBuilder() creates a BACnetConstructedDataEventEnableBuilder
func NewBACnetConstructedDataEventEnableBuilder() BACnetConstructedDataEventEnableBuilder {
	return &_BACnetConstructedDataEventEnableBuilder{_BACnetConstructedDataEventEnable: new(_BACnetConstructedDataEventEnable)}
}

type _BACnetConstructedDataEventEnableBuilder struct {
	*_BACnetConstructedDataEventEnable

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEventEnableBuilder) = (*_BACnetConstructedDataEventEnableBuilder)(nil)

func (b *_BACnetConstructedDataEventEnableBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataEventEnableBuilder) WithMandatoryFields(eventEnable BACnetEventTransitionBitsTagged) BACnetConstructedDataEventEnableBuilder {
	return b.WithEventEnable(eventEnable)
}

func (b *_BACnetConstructedDataEventEnableBuilder) WithEventEnable(eventEnable BACnetEventTransitionBitsTagged) BACnetConstructedDataEventEnableBuilder {
	b.EventEnable = eventEnable
	return b
}

func (b *_BACnetConstructedDataEventEnableBuilder) WithEventEnableBuilder(builderSupplier func(BACnetEventTransitionBitsTaggedBuilder) BACnetEventTransitionBitsTaggedBuilder) BACnetConstructedDataEventEnableBuilder {
	builder := builderSupplier(b.EventEnable.CreateBACnetEventTransitionBitsTaggedBuilder())
	var err error
	b.EventEnable, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventTransitionBitsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEventEnableBuilder) Build() (BACnetConstructedDataEventEnable, error) {
	if b.EventEnable == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'eventEnable' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEventEnable.deepCopy(), nil
}

func (b *_BACnetConstructedDataEventEnableBuilder) MustBuild() BACnetConstructedDataEventEnable {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEventEnableBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEventEnableBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEventEnableBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEventEnableBuilder().(*_BACnetConstructedDataEventEnableBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEventEnableBuilder creates a BACnetConstructedDataEventEnableBuilder
func (b *_BACnetConstructedDataEventEnable) CreateBACnetConstructedDataEventEnableBuilder() BACnetConstructedDataEventEnableBuilder {
	if b == nil {
		return NewBACnetConstructedDataEventEnableBuilder()
	}
	return &_BACnetConstructedDataEventEnableBuilder{_BACnetConstructedDataEventEnable: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventEnable) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventEnable) GetEventEnable() BACnetEventTransitionBitsTagged {
	return m.EventEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventEnable) GetActualValue() BACnetEventTransitionBitsTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetEventTransitionBitsTagged(m.GetEventEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventEnable(structType any) BACnetConstructedDataEventEnable {
	if casted, ok := structType.(BACnetConstructedDataEventEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventEnable) GetTypeName() string {
	return "BACnetConstructedDataEventEnable"
}

func (m *_BACnetConstructedDataEventEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (eventEnable)
	lengthInBits += m.EventEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventEnable) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventEnable BACnetConstructedDataEventEnable, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	eventEnable, err := ReadSimpleField[BACnetEventTransitionBitsTagged](ctx, "eventEnable", ReadComplex[BACnetEventTransitionBitsTagged](BACnetEventTransitionBitsTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventEnable' field"))
	}
	m.EventEnable = eventEnable

	actualValue, err := ReadVirtualField[BACnetEventTransitionBitsTagged](ctx, "actualValue", (*BACnetEventTransitionBitsTagged)(nil), eventEnable)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventEnable")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventEnable")
		}

		if err := WriteSimpleField[BACnetEventTransitionBitsTagged](ctx, "eventEnable", m.GetEventEnable(), WriteComplex[BACnetEventTransitionBitsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventEnable' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventEnable")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventEnable) IsBACnetConstructedDataEventEnable() {}

func (m *_BACnetConstructedDataEventEnable) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEventEnable) deepCopy() *_BACnetConstructedDataEventEnable {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEventEnableCopy := &_BACnetConstructedDataEventEnable{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetEventTransitionBitsTagged](m.EventEnable),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEventEnableCopy
}

func (m *_BACnetConstructedDataEventEnable) String() string {
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
