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

// SecurityDataArmSystem is the corresponding interface of SecurityDataArmSystem
type SecurityDataArmSystem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetArmMode returns ArmMode (property field)
	GetArmMode() byte
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
	// GetIsArmToAwayMode returns IsArmToAwayMode (virtual field)
	GetIsArmToAwayMode() bool
	// GetIsArmToNightMode returns IsArmToNightMode (virtual field)
	GetIsArmToNightMode() bool
	// GetIsArmToDayMode returns IsArmToDayMode (virtual field)
	GetIsArmToDayMode() bool
	// GetIsArmToVacationMode returns IsArmToVacationMode (virtual field)
	GetIsArmToVacationMode() bool
	// GetIsArmToHighestLevelOfProtection returns IsArmToHighestLevelOfProtection (virtual field)
	GetIsArmToHighestLevelOfProtection() bool
}

// SecurityDataArmSystemExactly can be used when we want exactly this type and not a type which fulfills SecurityDataArmSystem.
// This is useful for switch cases.
type SecurityDataArmSystemExactly interface {
	SecurityDataArmSystem
	isSecurityDataArmSystem() bool
}

// _SecurityDataArmSystem is the data-structure of this message
type _SecurityDataArmSystem struct {
	SecurityDataContract
	ArmMode byte
}

var _ SecurityDataArmSystem = (*_SecurityDataArmSystem)(nil)
var _ SecurityDataRequirements = (*_SecurityDataArmSystem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataArmSystem) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataArmSystem) GetArmMode() byte {
	return m.ArmMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataArmSystem) GetIsReserved() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool(bool((m.GetArmMode()) == (0x00))) || bool((bool(bool((m.GetArmMode()) >= (0x05))) && bool(bool((m.GetArmMode()) <= (0xFE))))))
}

func (m *_SecurityDataArmSystem) GetIsArmToAwayMode() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetArmMode()) == (0x01)))
}

func (m *_SecurityDataArmSystem) GetIsArmToNightMode() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetArmMode()) == (0x02)))
}

func (m *_SecurityDataArmSystem) GetIsArmToDayMode() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetArmMode()) == (0x03)))
}

func (m *_SecurityDataArmSystem) GetIsArmToVacationMode() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetArmMode()) == (0x04)))
}

func (m *_SecurityDataArmSystem) GetIsArmToHighestLevelOfProtection() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetArmMode()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataArmSystem factory function for _SecurityDataArmSystem
func NewSecurityDataArmSystem(armMode byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmSystem {
	_result := &_SecurityDataArmSystem{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		ArmMode:              armMode,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmSystem(structType any) SecurityDataArmSystem {
	if casted, ok := structType.(SecurityDataArmSystem); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmSystem); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmSystem) GetTypeName() string {
	return "SecurityDataArmSystem"
}

func (m *_SecurityDataArmSystem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (armMode)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataArmSystem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataArmSystem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataArmSystem SecurityDataArmSystem, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmSystem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmSystem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	armMode, err := ReadSimpleField(ctx, "armMode", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'armMode' field"))
	}
	m.ArmMode = armMode

	isReserved, err := ReadVirtualField[bool](ctx, "isReserved", (*bool)(nil), bool(bool((armMode) == (0x00))) || bool((bool(bool((armMode) >= (0x05))) && bool(bool((armMode) <= (0xFE))))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isReserved' field"))
	}
	_ = isReserved

	isArmToAwayMode, err := ReadVirtualField[bool](ctx, "isArmToAwayMode", (*bool)(nil), bool((armMode) == (0x01)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmToAwayMode' field"))
	}
	_ = isArmToAwayMode

	isArmToNightMode, err := ReadVirtualField[bool](ctx, "isArmToNightMode", (*bool)(nil), bool((armMode) == (0x02)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmToNightMode' field"))
	}
	_ = isArmToNightMode

	isArmToDayMode, err := ReadVirtualField[bool](ctx, "isArmToDayMode", (*bool)(nil), bool((armMode) == (0x03)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmToDayMode' field"))
	}
	_ = isArmToDayMode

	isArmToVacationMode, err := ReadVirtualField[bool](ctx, "isArmToVacationMode", (*bool)(nil), bool((armMode) == (0x04)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmToVacationMode' field"))
	}
	_ = isArmToVacationMode

	isArmToHighestLevelOfProtection, err := ReadVirtualField[bool](ctx, "isArmToHighestLevelOfProtection", (*bool)(nil), bool((armMode) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isArmToHighestLevelOfProtection' field"))
	}
	_ = isArmToHighestLevelOfProtection

	if closeErr := readBuffer.CloseContext("SecurityDataArmSystem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmSystem")
	}

	return m, nil
}

func (m *_SecurityDataArmSystem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataArmSystem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmSystem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmSystem")
		}

		if err := WriteSimpleField[byte](ctx, "armMode", m.GetArmMode(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'armMode' field")
		}
		// Virtual field
		isReserved := m.GetIsReserved()
		_ = isReserved
		if _isReservedErr := writeBuffer.WriteVirtual(ctx, "isReserved", m.GetIsReserved()); _isReservedErr != nil {
			return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
		}
		// Virtual field
		isArmToAwayMode := m.GetIsArmToAwayMode()
		_ = isArmToAwayMode
		if _isArmToAwayModeErr := writeBuffer.WriteVirtual(ctx, "isArmToAwayMode", m.GetIsArmToAwayMode()); _isArmToAwayModeErr != nil {
			return errors.Wrap(_isArmToAwayModeErr, "Error serializing 'isArmToAwayMode' field")
		}
		// Virtual field
		isArmToNightMode := m.GetIsArmToNightMode()
		_ = isArmToNightMode
		if _isArmToNightModeErr := writeBuffer.WriteVirtual(ctx, "isArmToNightMode", m.GetIsArmToNightMode()); _isArmToNightModeErr != nil {
			return errors.Wrap(_isArmToNightModeErr, "Error serializing 'isArmToNightMode' field")
		}
		// Virtual field
		isArmToDayMode := m.GetIsArmToDayMode()
		_ = isArmToDayMode
		if _isArmToDayModeErr := writeBuffer.WriteVirtual(ctx, "isArmToDayMode", m.GetIsArmToDayMode()); _isArmToDayModeErr != nil {
			return errors.Wrap(_isArmToDayModeErr, "Error serializing 'isArmToDayMode' field")
		}
		// Virtual field
		isArmToVacationMode := m.GetIsArmToVacationMode()
		_ = isArmToVacationMode
		if _isArmToVacationModeErr := writeBuffer.WriteVirtual(ctx, "isArmToVacationMode", m.GetIsArmToVacationMode()); _isArmToVacationModeErr != nil {
			return errors.Wrap(_isArmToVacationModeErr, "Error serializing 'isArmToVacationMode' field")
		}
		// Virtual field
		isArmToHighestLevelOfProtection := m.GetIsArmToHighestLevelOfProtection()
		_ = isArmToHighestLevelOfProtection
		if _isArmToHighestLevelOfProtectionErr := writeBuffer.WriteVirtual(ctx, "isArmToHighestLevelOfProtection", m.GetIsArmToHighestLevelOfProtection()); _isArmToHighestLevelOfProtectionErr != nil {
			return errors.Wrap(_isArmToHighestLevelOfProtectionErr, "Error serializing 'isArmToHighestLevelOfProtection' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmSystem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmSystem")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataArmSystem) isSecurityDataArmSystem() bool {
	return true
}

func (m *_SecurityDataArmSystem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
