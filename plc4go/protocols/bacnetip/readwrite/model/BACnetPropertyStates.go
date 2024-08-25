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

// BACnetPropertyStates is the corresponding interface of BACnetPropertyStates
type BACnetPropertyStates interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetPropertyStatesExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStates.
// This is useful for switch cases.
type BACnetPropertyStatesExactly interface {
	BACnetPropertyStates
	isBACnetPropertyStates() bool
}

// _BACnetPropertyStates is the data-structure of this message
type _BACnetPropertyStates struct {
	_BACnetPropertyStatesChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetPropertyStatesChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetPropertyStatesParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetPropertyStatesChild interface {
	utils.Serializable
	InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetPropertyStates

	GetTypeName() string
	BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetPropertyStates) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStates factory function for _BACnetPropertyStates
func NewBACnetPropertyStates(peekedTagHeader BACnetTagHeader) *_BACnetPropertyStates {
	return &_BACnetPropertyStates{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStates(structType any) BACnetPropertyStates {
	if casted, ok := structType.(BACnetPropertyStates); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStates); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStates) GetTypeName() string {
	return "BACnetPropertyStates"
}

func (m *_BACnetPropertyStates) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPropertyStates) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesParse(ctx context.Context, theBytes []byte) (BACnetPropertyStates, error) {
	return BACnetPropertyStatesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetPropertyStatesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetPropertyStates, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStates"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStates")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetPropertyStatesChildSerializeRequirement interface {
		BACnetPropertyStates
		InitializeParent(BACnetPropertyStates, BACnetTagHeader)
		GetParent() BACnetPropertyStates
	}
	var _childTemp any
	var _child BACnetPropertyStatesChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetPropertyStatesBoolean
		_childTemp, typeSwitchError = BACnetPropertyStatesBooleanParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetPropertyStatesBinaryValue
		_childTemp, typeSwitchError = BACnetPropertyStatesBinaryValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(2): // BACnetPropertyStatesEventType
		_childTemp, typeSwitchError = BACnetPropertyStatesEventTypeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(3): // BACnetPropertyStatesPolarity
		_childTemp, typeSwitchError = BACnetPropertyStatesPolarityParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(4): // BACnetPropertyStatesProgramChange
		_childTemp, typeSwitchError = BACnetPropertyStatesProgramChangeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(5): // BACnetPropertyStatesProgramChange
		_childTemp, typeSwitchError = BACnetPropertyStatesProgramChangeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(6): // BACnetPropertyStatesReasonForHalt
		_childTemp, typeSwitchError = BACnetPropertyStatesReasonForHaltParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(7): // BACnetPropertyStatesReliability
		_childTemp, typeSwitchError = BACnetPropertyStatesReliabilityParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(8): // BACnetPropertyStatesState
		_childTemp, typeSwitchError = BACnetPropertyStatesStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(9): // BACnetPropertyStatesSystemStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesSystemStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(10): // BACnetPropertyStatesUnits
		_childTemp, typeSwitchError = BACnetPropertyStatesUnitsParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(11): // BACnetPropertyStatesExtendedValue
		_childTemp, typeSwitchError = BACnetPropertyStatesExtendedValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(12): // BACnetPropertyStatesLifeSafetyMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(13): // BACnetPropertyStatesLifeSafetyState
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(14): // BACnetPropertyStatesRestartReason
		_childTemp, typeSwitchError = BACnetPropertyStatesRestartReasonParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(15): // BACnetPropertyStatesDoorAlarmState
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorAlarmStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(16): // BACnetPropertyStatesAction
		_childTemp, typeSwitchError = BACnetPropertyStatesActionParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(17): // BACnetPropertyStatesDoorSecuredStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorSecuredStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(18): // BACnetPropertyStatesDoorStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(19): // BACnetPropertyStatesDoorValue
		_childTemp, typeSwitchError = BACnetPropertyStatesDoorValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(20): // BACnetPropertyStatesFileAccessMethod
		_childTemp, typeSwitchError = BACnetPropertyStatesFileAccessMethodParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(21): // BACnetPropertyStatesLockStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesLockStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(22): // BACnetPropertyStatesLifeSafetyOperations
		_childTemp, typeSwitchError = BACnetPropertyStatesLifeSafetyOperationsParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(23): // BACnetPropertyStatesMaintenance
		_childTemp, typeSwitchError = BACnetPropertyStatesMaintenanceParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(24): // BACnetPropertyStatesNodeType
		_childTemp, typeSwitchError = BACnetPropertyStatesNodeTypeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(25): // BACnetPropertyStatesNotifyType
		_childTemp, typeSwitchError = BACnetPropertyStatesNotifyTypeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(26): // BACnetPropertyStatesSecurityLevel
		_childTemp, typeSwitchError = BACnetPropertyStatesSecurityLevelParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(27): // BACnetPropertyStatesShedState
		_childTemp, typeSwitchError = BACnetPropertyStatesShedStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(28): // BACnetPropertyStatesSilencedState
		_childTemp, typeSwitchError = BACnetPropertyStatesSilencedStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(30): // BACnetPropertyStatesAccessEvent
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessEventParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(31): // BACnetPropertyStatesZoneOccupanyState
		_childTemp, typeSwitchError = BACnetPropertyStatesZoneOccupanyStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(32): // BACnetPropertyStatesAccessCredentialDisableReason
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessCredentialDisableReasonParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(33): // BACnetPropertyStatesAccessCredentialDisable
		_childTemp, typeSwitchError = BACnetPropertyStatesAccessCredentialDisableParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(34): // BACnetPropertyStatesAuthenticationStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesAuthenticationStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(36): // BACnetPropertyStatesBackupState
		_childTemp, typeSwitchError = BACnetPropertyStatesBackupStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(37): // BACnetPropertyStatesWriteStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesWriteStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(38): // BACnetPropertyStatesLightningInProgress
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningInProgressParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(39): // BACnetPropertyStatesLightningOperation
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningOperationParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(40): // BACnetPropertyStatesLightningTransition
		_childTemp, typeSwitchError = BACnetPropertyStatesLightningTransitionParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(41): // BACnetPropertyStatesIntegerValue
		_childTemp, typeSwitchError = BACnetPropertyStatesIntegerValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(42): // BACnetPropertyStatesBinaryLightningValue
		_childTemp, typeSwitchError = BACnetPropertyStatesBinaryLightningValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(43): // BACnetPropertyStatesTimerState
		_childTemp, typeSwitchError = BACnetPropertyStatesTimerStateParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(44): // BACnetPropertyStatesTimerTransition
		_childTemp, typeSwitchError = BACnetPropertyStatesTimerTransitionParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(45): // BACnetPropertyStatesBacnetIpMode
		_childTemp, typeSwitchError = BACnetPropertyStatesBacnetIpModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(46): // BACnetPropertyStatesNetworkPortCommand
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkPortCommandParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(47): // BACnetPropertyStatesNetworkType
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkTypeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(48): // BACnetPropertyStatesNetworkNumberQuality
		_childTemp, typeSwitchError = BACnetPropertyStatesNetworkNumberQualityParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(49): // BACnetPropertyStatesEscalatorOperationDirection
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorOperationDirectionParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(50): // BACnetPropertyStatesEscalatorFault
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorFaultParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(51): // BACnetPropertyStatesEscalatorMode
		_childTemp, typeSwitchError = BACnetPropertyStatesEscalatorModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(52): // BACnetPropertyStatesLiftCarDirection
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDirectionParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(53): // BACnetPropertyStatesLiftCarDoorCommand
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDoorCommandParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(54): // BACnetPropertyStatesLiftCarDriveStatus
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarDriveStatusParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(55): // BACnetPropertyStatesLiftCarMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftCarModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(56): // BACnetPropertyStatesLiftGroupMode
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftGroupModeParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(57): // BACnetPropertyStatesLiftFault
		_childTemp, typeSwitchError = BACnetPropertyStatesLiftFaultParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(58): // BACnetPropertyStatesProtocolLevel
		_childTemp, typeSwitchError = BACnetPropertyStatesProtocolLevelParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case peekedTagNumber == uint8(63): // BACnetPropertyStatesExtendedValue
		_childTemp, typeSwitchError = BACnetPropertyStatesExtendedValueParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	case true: // BACnetPropertyStateActionUnknown
		_childTemp, typeSwitchError = BACnetPropertyStateActionUnknownParseWithBuffer(ctx, readBuffer, peekedTagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetPropertyStates")
	}
	_child = _childTemp.(BACnetPropertyStatesChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetPropertyStates"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStates")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetPropertyStates) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPropertyStates, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetPropertyStates"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStates")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyStates"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyStates")
	}
	return nil
}

func (m *_BACnetPropertyStates) isBACnetPropertyStates() bool {
	return true
}

func (m *_BACnetPropertyStates) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
