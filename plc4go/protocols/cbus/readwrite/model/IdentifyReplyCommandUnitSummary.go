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
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandUnitSummary is the corresponding interface of IdentifyReplyCommandUnitSummary
type IdentifyReplyCommandUnitSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetAssertingNetworkBurden returns AssertingNetworkBurden (property field)
	GetAssertingNetworkBurden() bool
	// GetRestrikeTimingActive returns RestrikeTimingActive (property field)
	GetRestrikeTimingActive() bool
	// GetRemoteOFFInputAsserted returns RemoteOFFInputAsserted (property field)
	GetRemoteOFFInputAsserted() bool
	// GetRemoteONInputAsserted returns RemoteONInputAsserted (property field)
	GetRemoteONInputAsserted() bool
	// GetLocalToggleEnabled returns LocalToggleEnabled (property field)
	GetLocalToggleEnabled() bool
	// GetLocalToggleActiveState returns LocalToggleActiveState (property field)
	GetLocalToggleActiveState() bool
	// GetClockGenerationEnabled returns ClockGenerationEnabled (property field)
	GetClockGenerationEnabled() bool
	// GetUnitGeneratingClock returns UnitGeneratingClock (property field)
	GetUnitGeneratingClock() bool
}

// IdentifyReplyCommandUnitSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandUnitSummary.
// This is useful for switch cases.
type IdentifyReplyCommandUnitSummaryExactly interface {
	IdentifyReplyCommandUnitSummary
	isIdentifyReplyCommandUnitSummary() bool
}

// _IdentifyReplyCommandUnitSummary is the data-structure of this message
type _IdentifyReplyCommandUnitSummary struct {
	AssertingNetworkBurden bool
	RestrikeTimingActive   bool
	RemoteOFFInputAsserted bool
	RemoteONInputAsserted  bool
	LocalToggleEnabled     bool
	LocalToggleActiveState bool
	ClockGenerationEnabled bool
	UnitGeneratingClock    bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandUnitSummary) GetAssertingNetworkBurden() bool {
	return m.AssertingNetworkBurden
}

func (m *_IdentifyReplyCommandUnitSummary) GetRestrikeTimingActive() bool {
	return m.RestrikeTimingActive
}

func (m *_IdentifyReplyCommandUnitSummary) GetRemoteOFFInputAsserted() bool {
	return m.RemoteOFFInputAsserted
}

func (m *_IdentifyReplyCommandUnitSummary) GetRemoteONInputAsserted() bool {
	return m.RemoteONInputAsserted
}

func (m *_IdentifyReplyCommandUnitSummary) GetLocalToggleEnabled() bool {
	return m.LocalToggleEnabled
}

func (m *_IdentifyReplyCommandUnitSummary) GetLocalToggleActiveState() bool {
	return m.LocalToggleActiveState
}

func (m *_IdentifyReplyCommandUnitSummary) GetClockGenerationEnabled() bool {
	return m.ClockGenerationEnabled
}

func (m *_IdentifyReplyCommandUnitSummary) GetUnitGeneratingClock() bool {
	return m.UnitGeneratingClock
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandUnitSummary factory function for _IdentifyReplyCommandUnitSummary
func NewIdentifyReplyCommandUnitSummary(assertingNetworkBurden bool, restrikeTimingActive bool, remoteOFFInputAsserted bool, remoteONInputAsserted bool, localToggleEnabled bool, localToggleActiveState bool, clockGenerationEnabled bool, unitGeneratingClock bool) *_IdentifyReplyCommandUnitSummary {
	return &_IdentifyReplyCommandUnitSummary{AssertingNetworkBurden: assertingNetworkBurden, RestrikeTimingActive: restrikeTimingActive, RemoteOFFInputAsserted: remoteOFFInputAsserted, RemoteONInputAsserted: remoteONInputAsserted, LocalToggleEnabled: localToggleEnabled, LocalToggleActiveState: localToggleActiveState, ClockGenerationEnabled: clockGenerationEnabled, UnitGeneratingClock: unitGeneratingClock}
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandUnitSummary(structType any) IdentifyReplyCommandUnitSummary {
	if casted, ok := structType.(IdentifyReplyCommandUnitSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandUnitSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandUnitSummary) GetTypeName() string {
	return "IdentifyReplyCommandUnitSummary"
}

func (m *_IdentifyReplyCommandUnitSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (assertingNetworkBurden)
	lengthInBits += 1

	// Simple field (restrikeTimingActive)
	lengthInBits += 1

	// Simple field (remoteOFFInputAsserted)
	lengthInBits += 1

	// Simple field (remoteONInputAsserted)
	lengthInBits += 1

	// Simple field (localToggleEnabled)
	lengthInBits += 1

	// Simple field (localToggleActiveState)
	lengthInBits += 1

	// Simple field (clockGenerationEnabled)
	lengthInBits += 1

	// Simple field (unitGeneratingClock)
	lengthInBits += 1

	return lengthInBits
}

func (m *_IdentifyReplyCommandUnitSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandUnitSummaryParse(ctx context.Context, theBytes []byte) (IdentifyReplyCommandUnitSummary, error) {
	return IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IdentifyReplyCommandUnitSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IdentifyReplyCommandUnitSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandUnitSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandUnitSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (assertingNetworkBurden)
	_assertingNetworkBurden, _assertingNetworkBurdenErr := readBuffer.ReadBit("assertingNetworkBurden")
	if _assertingNetworkBurdenErr != nil {
		return nil, errors.Wrap(_assertingNetworkBurdenErr, "Error parsing 'assertingNetworkBurden' field of IdentifyReplyCommandUnitSummary")
	}
	assertingNetworkBurden := _assertingNetworkBurden

	// Simple Field (restrikeTimingActive)
	_restrikeTimingActive, _restrikeTimingActiveErr := readBuffer.ReadBit("restrikeTimingActive")
	if _restrikeTimingActiveErr != nil {
		return nil, errors.Wrap(_restrikeTimingActiveErr, "Error parsing 'restrikeTimingActive' field of IdentifyReplyCommandUnitSummary")
	}
	restrikeTimingActive := _restrikeTimingActive

	// Simple Field (remoteOFFInputAsserted)
	_remoteOFFInputAsserted, _remoteOFFInputAssertedErr := readBuffer.ReadBit("remoteOFFInputAsserted")
	if _remoteOFFInputAssertedErr != nil {
		return nil, errors.Wrap(_remoteOFFInputAssertedErr, "Error parsing 'remoteOFFInputAsserted' field of IdentifyReplyCommandUnitSummary")
	}
	remoteOFFInputAsserted := _remoteOFFInputAsserted

	// Simple Field (remoteONInputAsserted)
	_remoteONInputAsserted, _remoteONInputAssertedErr := readBuffer.ReadBit("remoteONInputAsserted")
	if _remoteONInputAssertedErr != nil {
		return nil, errors.Wrap(_remoteONInputAssertedErr, "Error parsing 'remoteONInputAsserted' field of IdentifyReplyCommandUnitSummary")
	}
	remoteONInputAsserted := _remoteONInputAsserted

	// Simple Field (localToggleEnabled)
	_localToggleEnabled, _localToggleEnabledErr := readBuffer.ReadBit("localToggleEnabled")
	if _localToggleEnabledErr != nil {
		return nil, errors.Wrap(_localToggleEnabledErr, "Error parsing 'localToggleEnabled' field of IdentifyReplyCommandUnitSummary")
	}
	localToggleEnabled := _localToggleEnabled

	// Simple Field (localToggleActiveState)
	_localToggleActiveState, _localToggleActiveStateErr := readBuffer.ReadBit("localToggleActiveState")
	if _localToggleActiveStateErr != nil {
		return nil, errors.Wrap(_localToggleActiveStateErr, "Error parsing 'localToggleActiveState' field of IdentifyReplyCommandUnitSummary")
	}
	localToggleActiveState := _localToggleActiveState

	// Simple Field (clockGenerationEnabled)
	_clockGenerationEnabled, _clockGenerationEnabledErr := readBuffer.ReadBit("clockGenerationEnabled")
	if _clockGenerationEnabledErr != nil {
		return nil, errors.Wrap(_clockGenerationEnabledErr, "Error parsing 'clockGenerationEnabled' field of IdentifyReplyCommandUnitSummary")
	}
	clockGenerationEnabled := _clockGenerationEnabled

	// Simple Field (unitGeneratingClock)
	_unitGeneratingClock, _unitGeneratingClockErr := readBuffer.ReadBit("unitGeneratingClock")
	if _unitGeneratingClockErr != nil {
		return nil, errors.Wrap(_unitGeneratingClockErr, "Error parsing 'unitGeneratingClock' field of IdentifyReplyCommandUnitSummary")
	}
	unitGeneratingClock := _unitGeneratingClock

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandUnitSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandUnitSummary")
	}

	// Create the instance
	return &_IdentifyReplyCommandUnitSummary{
		AssertingNetworkBurden: assertingNetworkBurden,
		RestrikeTimingActive:   restrikeTimingActive,
		RemoteOFFInputAsserted: remoteOFFInputAsserted,
		RemoteONInputAsserted:  remoteONInputAsserted,
		LocalToggleEnabled:     localToggleEnabled,
		LocalToggleActiveState: localToggleActiveState,
		ClockGenerationEnabled: clockGenerationEnabled,
		UnitGeneratingClock:    unitGeneratingClock,
	}, nil
}

func (m *_IdentifyReplyCommandUnitSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandUnitSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("IdentifyReplyCommandUnitSummary"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandUnitSummary")
	}

	// Simple Field (assertingNetworkBurden)
	assertingNetworkBurden := bool(m.GetAssertingNetworkBurden())
	_assertingNetworkBurdenErr := writeBuffer.WriteBit("assertingNetworkBurden", (assertingNetworkBurden))
	if _assertingNetworkBurdenErr != nil {
		return errors.Wrap(_assertingNetworkBurdenErr, "Error serializing 'assertingNetworkBurden' field")
	}

	// Simple Field (restrikeTimingActive)
	restrikeTimingActive := bool(m.GetRestrikeTimingActive())
	_restrikeTimingActiveErr := writeBuffer.WriteBit("restrikeTimingActive", (restrikeTimingActive))
	if _restrikeTimingActiveErr != nil {
		return errors.Wrap(_restrikeTimingActiveErr, "Error serializing 'restrikeTimingActive' field")
	}

	// Simple Field (remoteOFFInputAsserted)
	remoteOFFInputAsserted := bool(m.GetRemoteOFFInputAsserted())
	_remoteOFFInputAssertedErr := writeBuffer.WriteBit("remoteOFFInputAsserted", (remoteOFFInputAsserted))
	if _remoteOFFInputAssertedErr != nil {
		return errors.Wrap(_remoteOFFInputAssertedErr, "Error serializing 'remoteOFFInputAsserted' field")
	}

	// Simple Field (remoteONInputAsserted)
	remoteONInputAsserted := bool(m.GetRemoteONInputAsserted())
	_remoteONInputAssertedErr := writeBuffer.WriteBit("remoteONInputAsserted", (remoteONInputAsserted))
	if _remoteONInputAssertedErr != nil {
		return errors.Wrap(_remoteONInputAssertedErr, "Error serializing 'remoteONInputAsserted' field")
	}

	// Simple Field (localToggleEnabled)
	localToggleEnabled := bool(m.GetLocalToggleEnabled())
	_localToggleEnabledErr := writeBuffer.WriteBit("localToggleEnabled", (localToggleEnabled))
	if _localToggleEnabledErr != nil {
		return errors.Wrap(_localToggleEnabledErr, "Error serializing 'localToggleEnabled' field")
	}

	// Simple Field (localToggleActiveState)
	localToggleActiveState := bool(m.GetLocalToggleActiveState())
	_localToggleActiveStateErr := writeBuffer.WriteBit("localToggleActiveState", (localToggleActiveState))
	if _localToggleActiveStateErr != nil {
		return errors.Wrap(_localToggleActiveStateErr, "Error serializing 'localToggleActiveState' field")
	}

	// Simple Field (clockGenerationEnabled)
	clockGenerationEnabled := bool(m.GetClockGenerationEnabled())
	_clockGenerationEnabledErr := writeBuffer.WriteBit("clockGenerationEnabled", (clockGenerationEnabled))
	if _clockGenerationEnabledErr != nil {
		return errors.Wrap(_clockGenerationEnabledErr, "Error serializing 'clockGenerationEnabled' field")
	}

	// Simple Field (unitGeneratingClock)
	unitGeneratingClock := bool(m.GetUnitGeneratingClock())
	_unitGeneratingClockErr := writeBuffer.WriteBit("unitGeneratingClock", (unitGeneratingClock))
	if _unitGeneratingClockErr != nil {
		return errors.Wrap(_unitGeneratingClockErr, "Error serializing 'unitGeneratingClock' field")
	}

	if popErr := writeBuffer.PopContext("IdentifyReplyCommandUnitSummary"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandUnitSummary")
	}
	return nil
}

func (m *_IdentifyReplyCommandUnitSummary) isIdentifyReplyCommandUnitSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandUnitSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
