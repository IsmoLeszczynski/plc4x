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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// TriggerControlData is the corresponding interface of TriggerControlData
type TriggerControlData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommandTypeContainer returns CommandTypeContainer (property field)
	GetCommandTypeContainer() TriggerControlCommandTypeContainer
	// GetTriggerGroup returns TriggerGroup (property field)
	GetTriggerGroup() byte
	// GetCommandType returns CommandType (virtual field)
	GetCommandType() TriggerControlCommandType
	// GetIsUnused returns IsUnused (virtual field)
	GetIsUnused() bool
}

// TriggerControlDataExactly can be used when we want exactly this type and not a type which fulfills TriggerControlData.
// This is useful for switch cases.
type TriggerControlDataExactly interface {
	TriggerControlData
	isTriggerControlData() bool
}

// _TriggerControlData is the data-structure of this message
type _TriggerControlData struct {
	_TriggerControlDataChildRequirements
	CommandTypeContainer TriggerControlCommandTypeContainer
	TriggerGroup         byte
}

type _TriggerControlDataChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCommandType() TriggerControlCommandType
}

type TriggerControlDataParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TriggerControlData, serializeChildFunction func() error) error
	GetTypeName() string
}

type TriggerControlDataChild interface {
	utils.Serializable
	InitializeParent(parent TriggerControlData, commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte)
	GetParent() *TriggerControlData

	GetTypeName() string
	TriggerControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlData) GetCommandTypeContainer() TriggerControlCommandTypeContainer {
	return m.CommandTypeContainer
}

func (m *_TriggerControlData) GetTriggerGroup() byte {
	return m.TriggerGroup
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_TriggerControlData) GetCommandType() TriggerControlCommandType {
	ctx := context.Background()
	_ = ctx
	return CastTriggerControlCommandType(m.GetCommandTypeContainer().CommandType())
}

func (m *_TriggerControlData) GetIsUnused() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetTriggerGroup()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTriggerControlData factory function for _TriggerControlData
func NewTriggerControlData(commandTypeContainer TriggerControlCommandTypeContainer, triggerGroup byte) *_TriggerControlData {
	return &_TriggerControlData{CommandTypeContainer: commandTypeContainer, TriggerGroup: triggerGroup}
}

// Deprecated: use the interface for direct cast
func CastTriggerControlData(structType any) TriggerControlData {
	if casted, ok := structType.(TriggerControlData); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlData); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlData) GetTypeName() string {
	return "TriggerControlData"
}

func (m *_TriggerControlData) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (commandTypeContainer)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (triggerGroup)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_TriggerControlData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TriggerControlDataParse(ctx context.Context, theBytes []byte) (TriggerControlData, error) {
	return TriggerControlDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TriggerControlDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TriggerControlData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TriggerControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(KnowsTriggerControlCommandTypeContainer(ctx, readBuffer)) {
		return nil, errors.WithStack(utils.ParseAssertError{"no command type could be found"})
	}

	// Simple Field (commandTypeContainer)
	if pullErr := readBuffer.PullContext("commandTypeContainer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for commandTypeContainer")
	}
	_commandTypeContainer, _commandTypeContainerErr := TriggerControlCommandTypeContainerParseWithBuffer(ctx, readBuffer)
	if _commandTypeContainerErr != nil {
		return nil, errors.Wrap(_commandTypeContainerErr, "Error parsing 'commandTypeContainer' field of TriggerControlData")
	}
	commandTypeContainer := _commandTypeContainer
	if closeErr := readBuffer.CloseContext("commandTypeContainer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for commandTypeContainer")
	}

	// Virtual field
	_commandType := commandTypeContainer.CommandType()
	commandType := TriggerControlCommandType(_commandType)
	_ = commandType

	// Simple Field (triggerGroup)
	_triggerGroup, _triggerGroupErr := readBuffer.ReadByte("triggerGroup")
	if _triggerGroupErr != nil {
		return nil, errors.Wrap(_triggerGroupErr, "Error parsing 'triggerGroup' field of TriggerControlData")
	}
	triggerGroup := _triggerGroup

	// Virtual field
	_isUnused := bool((triggerGroup) > (0xFE))
	isUnused := bool(_isUnused)
	_ = isUnused

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type TriggerControlDataChildSerializeRequirement interface {
		TriggerControlData
		InitializeParent(TriggerControlData, TriggerControlCommandTypeContainer, byte)
		GetParent() TriggerControlData
	}
	var _childTemp any
	var _child TriggerControlDataChildSerializeRequirement
	var typeSwitchError error
	switch {
	case commandType == TriggerControlCommandType_TRIGGER_EVENT: // TriggerControlDataTriggerEvent
		_childTemp, typeSwitchError = TriggerControlDataTriggerEventParseWithBuffer(ctx, readBuffer)
	case commandType == TriggerControlCommandType_TRIGGER_MIN: // TriggerControlDataTriggerMin
		_childTemp, typeSwitchError = TriggerControlDataTriggerMinParseWithBuffer(ctx, readBuffer)
	case commandType == TriggerControlCommandType_TRIGGER_MAX: // TriggerControlDataTriggerMax
		_childTemp, typeSwitchError = TriggerControlDataTriggerMaxParseWithBuffer(ctx, readBuffer)
	case commandType == TriggerControlCommandType_INDICATOR_KILL: // TriggerControlDataIndicatorKill
		_childTemp, typeSwitchError = TriggerControlDataIndicatorKillParseWithBuffer(ctx, readBuffer)
	case commandType == TriggerControlCommandType_LABEL: // TriggerControlDataLabel
		_childTemp, typeSwitchError = TriggerControlDataLabelParseWithBuffer(ctx, readBuffer, commandTypeContainer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [commandType=%v]", commandType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of TriggerControlData")
	}
	_child = _childTemp.(TriggerControlDataChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("TriggerControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlData")
	}

	// Finish initializing
	_child.InitializeParent(_child, commandTypeContainer, triggerGroup)
	return _child, nil
}

func (pm *_TriggerControlData) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child TriggerControlData, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("TriggerControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TriggerControlData")
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
	commandType := m.GetCommandType()
	_ = commandType
	if _commandTypeErr := writeBuffer.WriteVirtual(ctx, "commandType", m.GetCommandType()); _commandTypeErr != nil {
		return errors.Wrap(_commandTypeErr, "Error serializing 'commandType' field")
	}

	// Simple Field (triggerGroup)
	triggerGroup := byte(m.GetTriggerGroup())
	_triggerGroupErr := writeBuffer.WriteByte("triggerGroup", (triggerGroup))
	if _triggerGroupErr != nil {
		return errors.Wrap(_triggerGroupErr, "Error serializing 'triggerGroup' field")
	}
	// Virtual field
	isUnused := m.GetIsUnused()
	_ = isUnused
	if _isUnusedErr := writeBuffer.WriteVirtual(ctx, "isUnused", m.GetIsUnused()); _isUnusedErr != nil {
		return errors.Wrap(_isUnusedErr, "Error serializing 'isUnused' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("TriggerControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TriggerControlData")
	}
	return nil
}

func (m *_TriggerControlData) isTriggerControlData() bool {
	return true
}

func (m *_TriggerControlData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
