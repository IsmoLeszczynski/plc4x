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

// BACnetUnconfirmedServiceRequest is the corresponding interface of BACnetUnconfirmedServiceRequest
type BACnetUnconfirmedServiceRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetServiceChoice returns ServiceChoice (discriminator field)
	GetServiceChoice() BACnetUnconfirmedServiceChoice
}

// BACnetUnconfirmedServiceRequestExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequest.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestExactly interface {
	BACnetUnconfirmedServiceRequest
	isBACnetUnconfirmedServiceRequest() bool
}

// _BACnetUnconfirmedServiceRequest is the data-structure of this message
type _BACnetUnconfirmedServiceRequest struct {
	_BACnetUnconfirmedServiceRequestChildRequirements

	// Arguments.
	ServiceRequestLength uint16
}

type _BACnetUnconfirmedServiceRequestChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetServiceChoice() BACnetUnconfirmedServiceChoice
}

type BACnetUnconfirmedServiceRequestParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetUnconfirmedServiceRequestChild interface {
	utils.Serializable
	InitializeParent(parent BACnetUnconfirmedServiceRequest)
	GetParent() *BACnetUnconfirmedServiceRequest

	GetTypeName() string
	BACnetUnconfirmedServiceRequest
}

// NewBACnetUnconfirmedServiceRequest factory function for _BACnetUnconfirmedServiceRequest
func NewBACnetUnconfirmedServiceRequest(serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequest {
	return &_BACnetUnconfirmedServiceRequest{ServiceRequestLength: serviceRequestLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequest(structType any) BACnetUnconfirmedServiceRequest {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequest) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequest"
}

func (m *_BACnetUnconfirmedServiceRequest) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (serviceChoice)
	lengthInBits += 8

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestParse(ctx context.Context, theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequest, error) {
	return BACnetUnconfirmedServiceRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serviceChoice, err := ReadDiscriminatorEnumField[BACnetUnconfirmedServiceChoice](ctx, "serviceChoice", "BACnetUnconfirmedServiceChoice", ReadEnum(BACnetUnconfirmedServiceChoiceByValue, ReadUnsignedByte(readBuffer, 8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serviceChoice' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetUnconfirmedServiceRequestChildSerializeRequirement interface {
		BACnetUnconfirmedServiceRequest
		InitializeParent(BACnetUnconfirmedServiceRequest)
		GetParent() BACnetUnconfirmedServiceRequest
	}
	var _childTemp any
	var _child BACnetUnconfirmedServiceRequestChildSerializeRequirement
	var typeSwitchError error
	switch {
	case serviceChoice == BACnetUnconfirmedServiceChoice_I_AM: // BACnetUnconfirmedServiceRequestIAm
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestIAmParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_I_HAVE: // BACnetUnconfirmedServiceRequestIHave
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestIHaveParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION: // BACnetUnconfirmedServiceRequestUnconfirmedEventNotification
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedEventNotificationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER: // BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE: // BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedTextMessageParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION: // BACnetUnconfirmedServiceRequestTimeSynchronization
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestTimeSynchronizationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_WHO_HAS: // BACnetUnconfirmedServiceRequestWhoHas
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestWhoHasParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_WHO_IS: // BACnetUnconfirmedServiceRequestWhoIs
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestWhoIsParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION: // BACnetUnconfirmedServiceRequestUTCTimeSynchronization
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUTCTimeSynchronizationParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_WRITE_GROUP: // BACnetUnconfirmedServiceRequestWriteGroup
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestWriteGroupParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case serviceChoice == BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE: // BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	case 0 == 0: // BACnetUnconfirmedServiceRequestUnknown
		_childTemp, typeSwitchError = BACnetUnconfirmedServiceRequestUnknownParseWithBuffer(ctx, readBuffer, serviceRequestLength)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [serviceChoice=%v]", serviceChoice)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetUnconfirmedServiceRequest")
	}
	_child = _childTemp.(BACnetUnconfirmedServiceRequestChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequest")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_BACnetUnconfirmedServiceRequest) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetUnconfirmedServiceRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequest")
	}

	// Discriminator Field (serviceChoice) (Used as input to a switch field)
	serviceChoice := BACnetUnconfirmedServiceChoice(child.GetServiceChoice())
	if pushErr := writeBuffer.PushContext("serviceChoice"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for serviceChoice")
	}
	_serviceChoiceErr := writeBuffer.WriteSerializable(ctx, serviceChoice)
	if popErr := writeBuffer.PopContext("serviceChoice"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for serviceChoice")
	}

	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequest")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetUnconfirmedServiceRequest) GetServiceRequestLength() uint16 {
	return m.ServiceRequestLength
}

//
////

func (m *_BACnetUnconfirmedServiceRequest) isBACnetUnconfirmedServiceRequest() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
