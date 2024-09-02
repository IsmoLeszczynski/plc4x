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

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetLifetime returns Lifetime (property field)
	GetLifetime() BACnetContextTagUnsignedInteger
	// GetMaxNotificationDelay returns MaxNotificationDelay (property field)
	GetMaxNotificationDelay() BACnetContextTagUnsignedInteger
	// GetListOfCovSubscriptionSpecifications returns ListOfCovSubscriptionSpecifications (property field)
	GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
}

// _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier         BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications         BACnetContextTagBoolean
	Lifetime                            BACnetContextTagUnsignedInteger
	MaxNotificationDelay                BACnetContextTagUnsignedInteger
	ListOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
}

var _ BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple = (*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLifetime() BACnetContextTagUnsignedInteger {
	return m.Lifetime
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetMaxNotificationDelay() BACnetContextTagUnsignedInteger {
	return m.MaxNotificationDelay
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList {
	return m.ListOfCovSubscriptionSpecifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple factory function for _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, issueConfirmedNotifications BACnetContextTagBoolean, lifetime BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	_result := &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		IssueConfirmedNotifications:           issueConfirmedNotifications,
		Lifetime:                              lifetime,
		MaxNotificationDelay:                  maxNotificationDelay,
		ListOfCovSubscriptionSpecifications:   listOfCovSubscriptionSpecifications,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(structType any) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)
	}

	// Optional Field (lifetime)
	if m.Lifetime != nil {
		lengthInBits += m.Lifetime.GetLengthInBits(ctx)
	}

	// Optional Field (maxNotificationDelay)
	if m.MaxNotificationDelay != nil {
		lengthInBits += m.MaxNotificationDelay.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovSubscriptionSpecifications)
	lengthInBits += m.ListOfCovSubscriptionSpecifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	var issueConfirmedNotifications BACnetContextTagBoolean
	_issueConfirmedNotifications, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	if _issueConfirmedNotifications != nil {
		issueConfirmedNotifications = *_issueConfirmedNotifications
		m.IssueConfirmedNotifications = issueConfirmedNotifications
	}

	var lifetime BACnetContextTagUnsignedInteger
	_lifetime, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetime' field"))
	}
	if _lifetime != nil {
		lifetime = *_lifetime
		m.Lifetime = lifetime
	}

	var maxNotificationDelay BACnetContextTagUnsignedInteger
	_maxNotificationDelay, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationDelay' field"))
	}
	if _maxNotificationDelay != nil {
		maxNotificationDelay = *_maxNotificationDelay
		m.MaxNotificationDelay = maxNotificationDelay
	}

	listOfCovSubscriptionSpecifications, err := ReadSimpleField[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](ctx, "listOfCovSubscriptionSpecifications", ReadComplex[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecifications' field"))
	}
	m.ListOfCovSubscriptionSpecifications = listOfCovSubscriptionSpecifications

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", GetRef(m.GetIssueConfirmedNotifications()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", GetRef(m.GetLifetime()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetime' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", GetRef(m.GetMaxNotificationDelay()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationDelay' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](ctx, "listOfCovSubscriptionSpecifications", m.GetListOfCovSubscriptionSpecifications(), WriteComplex[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfCovSubscriptionSpecifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) IsBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple() {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
