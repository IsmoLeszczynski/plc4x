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

// BACnetConstructedDataAccessAlarmEvents is the corresponding interface of BACnetConstructedDataAccessAlarmEvents
type BACnetConstructedDataAccessAlarmEvents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAccessAlarmEvents returns AccessAlarmEvents (property field)
	GetAccessAlarmEvents() []BACnetAccessEventTagged
	// IsBACnetConstructedDataAccessAlarmEvents is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessAlarmEvents()
	// CreateBuilder creates a BACnetConstructedDataAccessAlarmEventsBuilder
	CreateBACnetConstructedDataAccessAlarmEventsBuilder() BACnetConstructedDataAccessAlarmEventsBuilder
}

// _BACnetConstructedDataAccessAlarmEvents is the data-structure of this message
type _BACnetConstructedDataAccessAlarmEvents struct {
	BACnetConstructedDataContract
	AccessAlarmEvents []BACnetAccessEventTagged
}

var _ BACnetConstructedDataAccessAlarmEvents = (*_BACnetConstructedDataAccessAlarmEvents)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessAlarmEvents)(nil)

// NewBACnetConstructedDataAccessAlarmEvents factory function for _BACnetConstructedDataAccessAlarmEvents
func NewBACnetConstructedDataAccessAlarmEvents(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, accessAlarmEvents []BACnetAccessEventTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessAlarmEvents {
	_result := &_BACnetConstructedDataAccessAlarmEvents{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessAlarmEvents:             accessAlarmEvents,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAccessAlarmEventsBuilder is a builder for BACnetConstructedDataAccessAlarmEvents
type BACnetConstructedDataAccessAlarmEventsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(accessAlarmEvents []BACnetAccessEventTagged) BACnetConstructedDataAccessAlarmEventsBuilder
	// WithAccessAlarmEvents adds AccessAlarmEvents (property field)
	WithAccessAlarmEvents(...BACnetAccessEventTagged) BACnetConstructedDataAccessAlarmEventsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataAccessAlarmEvents or returns an error if something is wrong
	Build() (BACnetConstructedDataAccessAlarmEvents, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAccessAlarmEvents
}

// NewBACnetConstructedDataAccessAlarmEventsBuilder() creates a BACnetConstructedDataAccessAlarmEventsBuilder
func NewBACnetConstructedDataAccessAlarmEventsBuilder() BACnetConstructedDataAccessAlarmEventsBuilder {
	return &_BACnetConstructedDataAccessAlarmEventsBuilder{_BACnetConstructedDataAccessAlarmEvents: new(_BACnetConstructedDataAccessAlarmEvents)}
}

type _BACnetConstructedDataAccessAlarmEventsBuilder struct {
	*_BACnetConstructedDataAccessAlarmEvents

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAccessAlarmEventsBuilder) = (*_BACnetConstructedDataAccessAlarmEventsBuilder)(nil)

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataAccessAlarmEvents
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) WithMandatoryFields(accessAlarmEvents []BACnetAccessEventTagged) BACnetConstructedDataAccessAlarmEventsBuilder {
	return b.WithAccessAlarmEvents(accessAlarmEvents...)
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) WithAccessAlarmEvents(accessAlarmEvents ...BACnetAccessEventTagged) BACnetConstructedDataAccessAlarmEventsBuilder {
	b.AccessAlarmEvents = accessAlarmEvents
	return b
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) Build() (BACnetConstructedDataAccessAlarmEvents, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAccessAlarmEvents.deepCopy(), nil
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) MustBuild() BACnetConstructedDataAccessAlarmEvents {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAccessAlarmEventsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAccessAlarmEventsBuilder().(*_BACnetConstructedDataAccessAlarmEventsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAccessAlarmEventsBuilder creates a BACnetConstructedDataAccessAlarmEventsBuilder
func (b *_BACnetConstructedDataAccessAlarmEvents) CreateBACnetConstructedDataAccessAlarmEventsBuilder() BACnetConstructedDataAccessAlarmEventsBuilder {
	if b == nil {
		return NewBACnetConstructedDataAccessAlarmEventsBuilder()
	}
	return &_BACnetConstructedDataAccessAlarmEventsBuilder{_BACnetConstructedDataAccessAlarmEvents: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_ALARM_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) GetAccessAlarmEvents() []BACnetAccessEventTagged {
	return m.AccessAlarmEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessAlarmEvents(structType any) BACnetConstructedDataAccessAlarmEvents {
	if casted, ok := structType.(BACnetConstructedDataAccessAlarmEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessAlarmEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetTypeName() string {
	return "BACnetConstructedDataAccessAlarmEvents"
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AccessAlarmEvents) > 0 {
		for _, element := range m.AccessAlarmEvents {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessAlarmEvents) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessAlarmEvents BACnetConstructedDataAccessAlarmEvents, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessAlarmEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessAlarmEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessAlarmEvents, err := ReadTerminatedArrayField[BACnetAccessEventTagged](ctx, "accessAlarmEvents", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessAlarmEvents' field"))
	}
	m.AccessAlarmEvents = accessAlarmEvents

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessAlarmEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessAlarmEvents")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessAlarmEvents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessAlarmEvents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessAlarmEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessAlarmEvents")
		}

		if err := WriteComplexTypeArrayField(ctx, "accessAlarmEvents", m.GetAccessAlarmEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'accessAlarmEvents' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessAlarmEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessAlarmEvents")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessAlarmEvents) IsBACnetConstructedDataAccessAlarmEvents() {}

func (m *_BACnetConstructedDataAccessAlarmEvents) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAccessAlarmEvents) deepCopy() *_BACnetConstructedDataAccessAlarmEvents {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAccessAlarmEventsCopy := &_BACnetConstructedDataAccessAlarmEvents{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopySlice[BACnetAccessEventTagged, BACnetAccessEventTagged](m.AccessAlarmEvents),
	}
	_BACnetConstructedDataAccessAlarmEventsCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAccessAlarmEventsCopy
}

func (m *_BACnetConstructedDataAccessAlarmEvents) String() string {
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
