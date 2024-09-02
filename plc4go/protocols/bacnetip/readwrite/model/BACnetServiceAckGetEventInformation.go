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

// BACnetServiceAckGetEventInformation is the corresponding interface of BACnetServiceAckGetEventInformation
type BACnetServiceAckGetEventInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() BACnetEventSummariesList
	// GetMoreEvents returns MoreEvents (property field)
	GetMoreEvents() BACnetContextTagBoolean
}

// BACnetServiceAckGetEventInformationExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetEventInformation.
// This is useful for switch cases.
type BACnetServiceAckGetEventInformationExactly interface {
	BACnetServiceAckGetEventInformation
	isBACnetServiceAckGetEventInformation() bool
}

// _BACnetServiceAckGetEventInformation is the data-structure of this message
type _BACnetServiceAckGetEventInformation struct {
	BACnetServiceAckContract
	ListOfEventSummaries BACnetEventSummariesList
	MoreEvents           BACnetContextTagBoolean
}

var _ BACnetServiceAckGetEventInformation = (*_BACnetServiceAckGetEventInformation)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckGetEventInformation)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetEventInformation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetEventInformation) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetEventInformation) GetListOfEventSummaries() BACnetEventSummariesList {
	return m.ListOfEventSummaries
}

func (m *_BACnetServiceAckGetEventInformation) GetMoreEvents() BACnetContextTagBoolean {
	return m.MoreEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetEventInformation factory function for _BACnetServiceAckGetEventInformation
func NewBACnetServiceAckGetEventInformation(listOfEventSummaries BACnetEventSummariesList, moreEvents BACnetContextTagBoolean, serviceAckLength uint32) *_BACnetServiceAckGetEventInformation {
	_result := &_BACnetServiceAckGetEventInformation{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ListOfEventSummaries:     listOfEventSummaries,
		MoreEvents:               moreEvents,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetEventInformation(structType any) BACnetServiceAckGetEventInformation {
	if casted, ok := structType.(BACnetServiceAckGetEventInformation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetEventInformation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetEventInformation) GetTypeName() string {
	return "BACnetServiceAckGetEventInformation"
}

func (m *_BACnetServiceAckGetEventInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (listOfEventSummaries)
	lengthInBits += m.ListOfEventSummaries.GetLengthInBits(ctx)

	// Simple field (moreEvents)
	lengthInBits += m.MoreEvents.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckGetEventInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckGetEventInformation) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckGetEventInformation BACnetServiceAckGetEventInformation, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEventInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetEventInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	listOfEventSummaries, err := ReadSimpleField[BACnetEventSummariesList](ctx, "listOfEventSummaries", ReadComplex[BACnetEventSummariesList](BACnetEventSummariesListParseWithBufferProducer((uint8)(uint8(0))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfEventSummaries' field"))
	}
	m.ListOfEventSummaries = listOfEventSummaries

	moreEvents, err := ReadSimpleField[BACnetContextTagBoolean](ctx, "moreEvents", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'moreEvents' field"))
	}
	m.MoreEvents = moreEvents

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEventInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetEventInformation")
	}

	return m, nil
}

func (m *_BACnetServiceAckGetEventInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetEventInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEventInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetEventInformation")
		}

		if err := WriteSimpleField[BACnetEventSummariesList](ctx, "listOfEventSummaries", m.GetListOfEventSummaries(), WriteComplex[BACnetEventSummariesList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfEventSummaries' field")
		}

		if err := WriteSimpleField[BACnetContextTagBoolean](ctx, "moreEvents", m.GetMoreEvents(), WriteComplex[BACnetContextTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'moreEvents' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEventInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetEventInformation")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetEventInformation) isBACnetServiceAckGetEventInformation() bool {
	return true
}

func (m *_BACnetServiceAckGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
