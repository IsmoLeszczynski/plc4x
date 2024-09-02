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

// SecurityDataLowBatteryCharging is the corresponding interface of SecurityDataLowBatteryCharging
type SecurityDataLowBatteryCharging interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetStartStop returns StartStop (property field)
	GetStartStop() byte
	// GetChargeStopped returns ChargeStopped (virtual field)
	GetChargeStopped() bool
	// GetChargeStarted returns ChargeStarted (virtual field)
	GetChargeStarted() bool
}

// _SecurityDataLowBatteryCharging is the data-structure of this message
type _SecurityDataLowBatteryCharging struct {
	SecurityDataContract
	StartStop byte
}

var _ SecurityDataLowBatteryCharging = (*_SecurityDataLowBatteryCharging)(nil)
var _ SecurityDataRequirements = (*_SecurityDataLowBatteryCharging)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataLowBatteryCharging) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataLowBatteryCharging) GetStartStop() byte {
	return m.StartStop
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataLowBatteryCharging) GetChargeStopped() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStartStop()) == (0x00)))
}

func (m *_SecurityDataLowBatteryCharging) GetChargeStarted() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetStartStop()) > (0xFE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataLowBatteryCharging factory function for _SecurityDataLowBatteryCharging
func NewSecurityDataLowBatteryCharging(startStop byte, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataLowBatteryCharging {
	_result := &_SecurityDataLowBatteryCharging{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		StartStop:            startStop,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataLowBatteryCharging(structType any) SecurityDataLowBatteryCharging {
	if casted, ok := structType.(SecurityDataLowBatteryCharging); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataLowBatteryCharging); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataLowBatteryCharging) GetTypeName() string {
	return "SecurityDataLowBatteryCharging"
}

func (m *_SecurityDataLowBatteryCharging) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Simple field (startStop)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_SecurityDataLowBatteryCharging) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataLowBatteryCharging) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataLowBatteryCharging SecurityDataLowBatteryCharging, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataLowBatteryCharging"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataLowBatteryCharging")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	startStop, err := ReadSimpleField(ctx, "startStop", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startStop' field"))
	}
	m.StartStop = startStop

	chargeStopped, err := ReadVirtualField[bool](ctx, "chargeStopped", (*bool)(nil), bool((startStop) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chargeStopped' field"))
	}
	_ = chargeStopped

	chargeStarted, err := ReadVirtualField[bool](ctx, "chargeStarted", (*bool)(nil), bool((startStop) > (0xFE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'chargeStarted' field"))
	}
	_ = chargeStarted

	if closeErr := readBuffer.CloseContext("SecurityDataLowBatteryCharging"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataLowBatteryCharging")
	}

	return m, nil
}

func (m *_SecurityDataLowBatteryCharging) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataLowBatteryCharging) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataLowBatteryCharging"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataLowBatteryCharging")
		}

		if err := WriteSimpleField[byte](ctx, "startStop", m.GetStartStop(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'startStop' field")
		}
		// Virtual field
		chargeStopped := m.GetChargeStopped()
		_ = chargeStopped
		if _chargeStoppedErr := writeBuffer.WriteVirtual(ctx, "chargeStopped", m.GetChargeStopped()); _chargeStoppedErr != nil {
			return errors.Wrap(_chargeStoppedErr, "Error serializing 'chargeStopped' field")
		}
		// Virtual field
		chargeStarted := m.GetChargeStarted()
		_ = chargeStarted
		if _chargeStartedErr := writeBuffer.WriteVirtual(ctx, "chargeStarted", m.GetChargeStarted()); _chargeStartedErr != nil {
			return errors.Wrap(_chargeStartedErr, "Error serializing 'chargeStarted' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataLowBatteryCharging"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataLowBatteryCharging")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataLowBatteryCharging) IsSecurityDataLowBatteryCharging() {}

func (m *_SecurityDataLowBatteryCharging) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
