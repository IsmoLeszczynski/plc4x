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

// SecurityDataZoneOpen is the corresponding interface of SecurityDataZoneOpen
type SecurityDataZoneOpen interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataZoneOpenExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneOpen.
// This is useful for switch cases.
type SecurityDataZoneOpenExactly interface {
	SecurityDataZoneOpen
	isSecurityDataZoneOpen() bool
}

// _SecurityDataZoneOpen is the data-structure of this message
type _SecurityDataZoneOpen struct {
	*_SecurityData
	ZoneNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneOpen) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataZoneOpen) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneOpen) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataZoneOpen factory function for _SecurityDataZoneOpen
func NewSecurityDataZoneOpen(zoneNumber uint8, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataZoneOpen {
	_result := &_SecurityDataZoneOpen{
		ZoneNumber:    zoneNumber,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneOpen(structType any) SecurityDataZoneOpen {
	if casted, ok := structType.(SecurityDataZoneOpen); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneOpen); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneOpen) GetTypeName() string {
	return "SecurityDataZoneOpen"
}

func (m *_SecurityDataZoneOpen) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SecurityDataZoneOpen) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataZoneOpenParse(ctx context.Context, theBytes []byte) (SecurityDataZoneOpen, error) {
	return SecurityDataZoneOpenParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataZoneOpenParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataZoneOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityDataZoneOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
	_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataZoneOpen")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneOpen")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataZoneOpen{
		_SecurityData: &_SecurityData{},
		ZoneNumber:    zoneNumber,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataZoneOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneOpen) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneOpen")
		}

		// Simple Field (zoneNumber)
		zoneNumber := uint8(m.GetZoneNumber())
		_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, uint8((zoneNumber)))
		if _zoneNumberErr != nil {
			return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataZoneOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneOpen")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataZoneOpen) isSecurityDataZoneOpen() bool {
	return true
}

func (m *_SecurityDataZoneOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
