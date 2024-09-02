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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckGetEnrollmentSummary is the corresponding interface of BACnetServiceAckGetEnrollmentSummary
type BACnetServiceAckGetEnrollmentSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetEventType returns EventType (property field)
	GetEventType() BACnetEventTypeTagged
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetPriority returns Priority (property field)
	GetPriority() BACnetApplicationTagUnsignedInteger
	// GetNotificationClass returns NotificationClass (property field)
	GetNotificationClass() BACnetApplicationTagUnsignedInteger
}

// _BACnetServiceAckGetEnrollmentSummary is the data-structure of this message
type _BACnetServiceAckGetEnrollmentSummary struct {
	BACnetServiceAckContract
	ObjectIdentifier  BACnetApplicationTagObjectIdentifier
	EventType         BACnetEventTypeTagged
	EventState        BACnetEventStateTagged
	Priority          BACnetApplicationTagUnsignedInteger
	NotificationClass BACnetApplicationTagUnsignedInteger
}

var _ BACnetServiceAckGetEnrollmentSummary = (*_BACnetServiceAckGetEnrollmentSummary)(nil)
var _ BACnetServiceAckRequirements = (*_BACnetServiceAckGetEnrollmentSummary)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ENROLLMENT_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetParent() BACnetServiceAckContract {
	return m.BACnetServiceAckContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetEnrollmentSummary) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventType() BACnetEventTypeTagged {
	return m.EventType
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetPriority() BACnetApplicationTagUnsignedInteger {
	return m.Priority
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetNotificationClass() BACnetApplicationTagUnsignedInteger {
	return m.NotificationClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetEnrollmentSummary factory function for _BACnetServiceAckGetEnrollmentSummary
func NewBACnetServiceAckGetEnrollmentSummary(objectIdentifier BACnetApplicationTagObjectIdentifier, eventType BACnetEventTypeTagged, eventState BACnetEventStateTagged, priority BACnetApplicationTagUnsignedInteger, notificationClass BACnetApplicationTagUnsignedInteger, serviceAckLength uint32) *_BACnetServiceAckGetEnrollmentSummary {
	_result := &_BACnetServiceAckGetEnrollmentSummary{
		BACnetServiceAckContract: NewBACnetServiceAck(serviceAckLength),
		ObjectIdentifier:         objectIdentifier,
		EventType:                eventType,
		EventState:               eventState,
		Priority:                 priority,
		NotificationClass:        notificationClass,
	}
	_result.BACnetServiceAckContract.(*_BACnetServiceAck)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetEnrollmentSummary(structType any) BACnetServiceAckGetEnrollmentSummary {
	if casted, ok := structType.(BACnetServiceAckGetEnrollmentSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetEnrollmentSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetTypeName() string {
	return "BACnetServiceAckGetEnrollmentSummary"
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetServiceAckContract.(*_BACnetServiceAck).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (eventType)
	lengthInBits += m.EventType.GetLengthInBits(ctx)

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += m.Priority.GetLengthInBits(ctx)

	// Optional Field (notificationClass)
	if m.NotificationClass != nil {
		lengthInBits += m.NotificationClass.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetServiceAckGetEnrollmentSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetServiceAckGetEnrollmentSummary) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetServiceAck, serviceAckLength uint32) (__bACnetServiceAckGetEnrollmentSummary BACnetServiceAckGetEnrollmentSummary, err error) {
	m.BACnetServiceAckContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEnrollmentSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetEnrollmentSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	eventType, err := ReadSimpleField[BACnetEventTypeTagged](ctx, "eventType", ReadComplex[BACnetEventTypeTagged](BACnetEventTypeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventType' field"))
	}
	m.EventType = eventType

	eventState, err := ReadSimpleField[BACnetEventStateTagged](ctx, "eventState", ReadComplex[BACnetEventStateTagged](BACnetEventStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'eventState' field"))
	}
	m.EventState = eventState

	priority, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	var notificationClass BACnetApplicationTagUnsignedInteger
	_notificationClass, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "notificationClass", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'notificationClass' field"))
	}
	if _notificationClass != nil {
		notificationClass = *_notificationClass
		m.NotificationClass = notificationClass
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEnrollmentSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetEnrollmentSummary")
	}

	return m, nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetEnrollmentSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEnrollmentSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetEnrollmentSummary")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetEventTypeTagged](ctx, "eventType", m.GetEventType(), WriteComplex[BACnetEventTypeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventType' field")
		}

		if err := WriteSimpleField[BACnetEventStateTagged](ctx, "eventState", m.GetEventState(), WriteComplex[BACnetEventStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'eventState' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "priority", m.GetPriority(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "notificationClass", GetRef(m.GetNotificationClass()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'notificationClass' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEnrollmentSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetEnrollmentSummary")
		}
		return nil
	}
	return m.BACnetServiceAckContract.(*_BACnetServiceAck).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetEnrollmentSummary) IsBACnetServiceAckGetEnrollmentSummary() {}

func (m *_BACnetServiceAckGetEnrollmentSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
