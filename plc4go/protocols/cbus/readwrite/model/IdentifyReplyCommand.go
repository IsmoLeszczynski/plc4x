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

// IdentifyReplyCommand is the corresponding interface of IdentifyReplyCommand
type IdentifyReplyCommand interface {
	IdentifyReplyCommandContract
	IdentifyReplyCommandRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// IdentifyReplyCommandContract provides a set of functions which can be overwritten by a sub struct
type IdentifyReplyCommandContract interface {
	// GetNumBytes() returns a parser argument
	GetNumBytes() uint8
	// IsIdentifyReplyCommand() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIdentifyReplyCommand()
}

// IdentifyReplyCommandRequirements provides a set of functions which need to be implemented by a sub struct
type IdentifyReplyCommandRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetAttribute returns Attribute (discriminator field)
	GetAttribute() Attribute
}

// _IdentifyReplyCommand is the data-structure of this message
type _IdentifyReplyCommand struct {
	_SubType IdentifyReplyCommand

	// Arguments.
	NumBytes uint8
}

var _ IdentifyReplyCommandContract = (*_IdentifyReplyCommand)(nil)

// NewIdentifyReplyCommand factory function for _IdentifyReplyCommand
func NewIdentifyReplyCommand(numBytes uint8) *_IdentifyReplyCommand {
	return &_IdentifyReplyCommand{NumBytes: numBytes}
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommand(structType any) IdentifyReplyCommand {
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommand) GetTypeName() string {
	return "IdentifyReplyCommand"
}

func (m *_IdentifyReplyCommand) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_IdentifyReplyCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandParse[T IdentifyReplyCommand](ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (T, error) {
	return IdentifyReplyCommandParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandParseWithBufferProducer[T IdentifyReplyCommand](attribute Attribute, numBytes uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := IdentifyReplyCommandParseWithBuffer[T](ctx, readBuffer, attribute, numBytes)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func IdentifyReplyCommandParseWithBuffer[T IdentifyReplyCommand](ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (T, error) {
	v, err := (&_IdentifyReplyCommand{NumBytes: numBytes}).parse(ctx, readBuffer, attribute, numBytes)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_IdentifyReplyCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (__identifyReplyCommand IdentifyReplyCommand, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child IdentifyReplyCommand
	switch {
	case attribute == Attribute_Manufacturer: // IdentifyReplyCommandManufacturer
		if _child, err = (&_IdentifyReplyCommandManufacturer{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandManufacturer for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_Type: // IdentifyReplyCommandType
		if _child, err = (&_IdentifyReplyCommandType{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandType for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_FirmwareVersion: // IdentifyReplyCommandFirmwareVersion
		if _child, err = (&_IdentifyReplyCommandFirmwareVersion{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandFirmwareVersion for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_Summary: // IdentifyReplyCommandSummary
		if _child, err = (&_IdentifyReplyCommandSummary{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandSummary for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_ExtendedDiagnosticSummary: // IdentifyReplyCommandExtendedDiagnosticSummary
		if _child, err = (&_IdentifyReplyCommandExtendedDiagnosticSummary{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandExtendedDiagnosticSummary for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_NetworkTerminalLevels: // IdentifyReplyCommandNetworkTerminalLevels
		if _child, err = (&_IdentifyReplyCommandNetworkTerminalLevels{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandNetworkTerminalLevels for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_TerminalLevel: // IdentifyReplyCommandTerminalLevels
		if _child, err = (&_IdentifyReplyCommandTerminalLevels{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandTerminalLevels for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_NetworkVoltage: // IdentifyReplyCommandNetworkVoltage
		if _child, err = (&_IdentifyReplyCommandNetworkVoltage{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandNetworkVoltage for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_GAVValuesCurrent: // IdentifyReplyCommandGAVValuesCurrent
		if _child, err = (&_IdentifyReplyCommandGAVValuesCurrent{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandGAVValuesCurrent for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_GAVValuesStored: // IdentifyReplyCommandGAVValuesStored
		if _child, err = (&_IdentifyReplyCommandGAVValuesStored{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandGAVValuesStored for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_GAVPhysicalAddresses: // IdentifyReplyCommandGAVPhysicalAddresses
		if _child, err = (&_IdentifyReplyCommandGAVPhysicalAddresses{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandGAVPhysicalAddresses for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_LogicalAssignment: // IdentifyReplyCommandLogicalAssignment
		if _child, err = (&_IdentifyReplyCommandLogicalAssignment{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandLogicalAssignment for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_Delays: // IdentifyReplyCommandDelays
		if _child, err = (&_IdentifyReplyCommandDelays{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandDelays for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_MinimumLevels: // IdentifyReplyCommandMinimumLevels
		if _child, err = (&_IdentifyReplyCommandMinimumLevels{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandMinimumLevels for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_MaximumLevels: // IdentifyReplyCommandMaximumLevels
		if _child, err = (&_IdentifyReplyCommandMaximumLevels{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandMaximumLevels for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_CurrentSenseLevels: // IdentifyReplyCommandCurrentSenseLevels
		if _child, err = (&_IdentifyReplyCommandCurrentSenseLevels{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandCurrentSenseLevels for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_OutputUnitSummary: // IdentifyReplyCommandOutputUnitSummary
		if _child, err = (&_IdentifyReplyCommandOutputUnitSummary{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandOutputUnitSummary for type-switch of IdentifyReplyCommand")
		}
	case attribute == Attribute_DSIStatus: // IdentifyReplyCommandDSIStatus
		if _child, err = (&_IdentifyReplyCommandDSIStatus{}).parse(ctx, readBuffer, m, attribute, numBytes); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type IdentifyReplyCommandDSIStatus for type-switch of IdentifyReplyCommand")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [attribute=%v]", attribute)
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommand")
	}

	return _child, nil
}

func (pm *_IdentifyReplyCommand) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child IdentifyReplyCommand, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("IdentifyReplyCommand"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommand")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("IdentifyReplyCommand"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IdentifyReplyCommand")
	}
	return nil
}

////
// Arguments Getter

func (m *_IdentifyReplyCommand) GetNumBytes() uint8 {
	return m.NumBytes
}

//
////

func (m *_IdentifyReplyCommand) IsIdentifyReplyCommand() {}
