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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// CALData is the corresponding interface of CALData
type CALData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() CALCommandTypeContainer
	// GetAdditionalData returns AdditionalData (property field)
	GetAdditionalData() CALData
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() CALCommandType
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (virtual field)
	GetSendIdentifyRequestBefore() bool
}

// CALDataExactly can be used when we want exactly this type and not a type which fulfills CALData.
// This is useful for switch cases.
type CALDataExactly interface {
	CALData
	isCALData() bool
}

// _CALData is the data-structure of this message
type _CALData struct {
	_CALDataChildRequirements
	CommandTypeContainer CALCommandTypeContainer
	AdditionalData       CALData

	// Arguments.
	RequestContext RequestContext
}

type _CALDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() CALCommandType
	GetSendIdentifyRequestBefore() bool
}

type CALDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALData, serializeChildFunction func() error) error
	GetTypeName() string
}

type CALDataChild interface {
	utils.Serializable
	InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer, additionalData CALData)
	GetParent() *CALData

	GetTypeName() string
	CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALData) GetCommandTypeContainer() CALCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_CALData) GetAdditionalData() CALData {
	return m.AdditionalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_CALData) GetCommandType() CALCommandType {
	ctx := context.Background()
	_ = ctx
	additionalData := m.AdditionalData
	_ = additionalData
	return CastCALCommandType(m.GetCommandTypeContainer().CommandType())
}

func (m *_CALData) GetSendIdentifyRequestBefore() bool {
	ctx := context.Background()
	_ = ctx
	additionalData := m.AdditionalData
	_ = additionalData
	return bool(utils.InlineIf(bool((m.RequestContext) != (nil)), func() any { return bool(m.RequestContext.GetSendIdentifyRequestBefore()) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALData factory function for _CALData
func NewCALData(commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALData {
	return &_CALData{CommandTypeContainer: commandTypeContainer, AdditionalData: additionalData, RequestContext: requestContext}
}

// Deprecated: use the interface for direct cast
func CastCALData(structType any) CALData {
	if casted, ok := structType.(CALData); ok {
		return casted
	}
	if casted, ok := structType.(*CALData); ok {
		return *casted
	}
	return nil
}

func (m *_CALData) GetTypeName() string {
	return "CALData"
}

func (m *_CALData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (additionalData)
	if m.AdditionalData != nil {
		lengthInBits += m.AdditionalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_CALData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALDataParse(ctx context.Context, theBytes []byte, requestContext RequestContext) (CALData, error) {
	return CALDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), requestContext)
}

func CALDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, requestContext RequestContext) (CALData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CALData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsCALCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := CALCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of CALData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := CALCommandType(_commandType)
	_ = commandType

	// Virtual field
	_sendIdentifyRequestBefore := utils.InlineIf(bool((requestContext) != (nil)), func() any { return bool(requestContext.GetSendIdentifyRequestBefore()) }, func() any { return bool(bool(false)) }).(bool)
	sendIdentifyRequestBefore := bool(_sendIdentifyRequestBefore)
	_ = sendIdentifyRequestBefore

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CALDataChildSerializeRequirement interface {
		CALData
		InitializeParent(CALData, CALCommandTypeContainer, CALData)
		GetParent() CALData
	}
	var _childTemp any
	var _child CALDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == CALCommandType_RESET: // CALDataReset
		_childTemp, typeSwitchError = CALDataResetParseWithBuffer(ctx, readBuffer, requestContext)
	case commandType == CALCommandType_RECALL: // CALDataRecall
		_childTemp, typeSwitchError = CALDataRecallParseWithBuffer(ctx, readBuffer, requestContext)
	case commandType == CALCommandType_IDENTIFY: // CALDataIdentify
		_childTemp, typeSwitchError = CALDataIdentifyParseWithBuffer(ctx, readBuffer, requestContext)
	case commandType == CALCommandType_GET_STATUS: // CALDataGetStatus
		_childTemp, typeSwitchError = CALDataGetStatusParseWithBuffer(ctx, readBuffer, requestContext)
	case commandType == CALCommandType_WRITE: // CALDataWrite
		_childTemp, typeSwitchError = CALDataWriteParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	case commandType == CALCommandType_REPLY && sendIdentifyRequestBefore == bool(true): // CALDataIdentifyReply
		_childTemp, typeSwitchError = CALDataIdentifyReplyParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	case commandType == CALCommandType_REPLY: // CALDataReply
		_childTemp, typeSwitchError = CALDataReplyParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	case commandType == CALCommandType_ACKNOWLEDGE: // CALDataAcknowledge
		_childTemp, typeSwitchError = CALDataAcknowledgeParseWithBuffer(ctx, readBuffer, requestContext)
	case commandType == CALCommandType_STATUS: // CALDataStatus
		_childTemp, typeSwitchError = CALDataStatusParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	case commandType == CALCommandType_STATUS_EXTENDED: // CALDataStatusExtended
		_childTemp, typeSwitchError = CALDataStatusExtendedParseWithBuffer(ctx, readBuffer, commandTypeContainer, requestContext)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v, sendIdentifyRequestBefore=%v]", commandType, sendIdentifyRequestBefore)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CALData")
	}
	_child = _childTemp.(CALDataChildSerializeRequirement)

	// Optional Field (additionalData) (Can be skipped, if a given expression evaluates to false)
	var additionalData CALData = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("additionalData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for additionalData")
		}
		_val, _err := CALDataParseWithBuffer(ctx, readBuffer, nil)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'additionalData' field of CALData")
		default:
			additionalData = _val.(CALData)
			if closeErr := readBuffer.CloseContext("additionalData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for additionalData")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("CALData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, additionalData)
	return _child, nil
}

func (pm *_CALData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CALData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CALData")
	}

	// Simple Field (commandTypeContainer)
	if pushErr := writeBuffer.PushContext("commandTypeContainer"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for commandTypeContainer")
	}
	_commandTypeContainerErr := writeBuffer.WriteSerializable(ctx, m.GetCommandTypeContainer())
	if popErr := writeBuffer.PopContext("commandTypeContainer"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for commandTypeContainer")
	}
	if _commandTypeContainerErr != nil {
		return errors.Wrap(_commandTypeContainerErr, "Error serializing 'commandTypeContainer' field")
	}
	// Virtual field
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}
	// Virtual field
	if _sendIdentifyRequestBeforeErr := writeBuffer.WriteVirtual(ctx, "sendIdentifyRequestBefore", m.GetSendIdentifyRequestBefore()); _sendIdentifyRequestBeforeErr != nil {
		return errors.Wrap(_sendIdentifyRequestBeforeErr, "Error serializing 'sendIdentifyRequestBefore' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Optional Field (additionalData) (Can be skipped, if the value is null)
	var additionalData CALData = nil
	if m.GetAdditionalData() != nil {
		if pushErr := writeBuffer.PushContext("additionalData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for additionalData")
		}
		additionalData = m.GetAdditionalData()
		_additionalDataErr := writeBuffer.WriteSerializable(ctx, additionalData)
		if popErr := writeBuffer.PopContext("additionalData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for additionalData")
		}
		if _additionalDataErr != nil {
			return errors.Wrap(_additionalDataErr, "Error serializing 'additionalData' field")
		}
	}

	if popErr := writeBuffer.PopContext("CALData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CALData")
	}
	return nil
}

////
// Arguments Getter

func (m *_CALData) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_CALData) isCALData() bool {
	return true
}

func (m *_CALData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
