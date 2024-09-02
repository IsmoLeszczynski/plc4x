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

// BACnetLandingCallStatus is the corresponding interface of BACnetLandingCallStatus
type BACnetLandingCallStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetFloorNumber returns FloorNumber (property field)
	GetFloorNumber() BACnetContextTagUnsignedInteger
	// GetCommand returns Command (property field)
	GetCommand() BACnetLandingCallStatusCommand
	// GetFloorText returns FloorText (property field)
	GetFloorText() BACnetContextTagCharacterString
}

// _BACnetLandingCallStatus is the data-structure of this message
type _BACnetLandingCallStatus struct {
	FloorNumber BACnetContextTagUnsignedInteger
	Command     BACnetLandingCallStatusCommand
	FloorText   BACnetContextTagCharacterString
}

var _ BACnetLandingCallStatus = (*_BACnetLandingCallStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingCallStatus) GetFloorNumber() BACnetContextTagUnsignedInteger {
	return m.FloorNumber
}

func (m *_BACnetLandingCallStatus) GetCommand() BACnetLandingCallStatusCommand {
	return m.Command
}

func (m *_BACnetLandingCallStatus) GetFloorText() BACnetContextTagCharacterString {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingCallStatus factory function for _BACnetLandingCallStatus
func NewBACnetLandingCallStatus(floorNumber BACnetContextTagUnsignedInteger, command BACnetLandingCallStatusCommand, floorText BACnetContextTagCharacterString) *_BACnetLandingCallStatus {
	return &_BACnetLandingCallStatus{FloorNumber: floorNumber, Command: command, FloorText: floorText}
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingCallStatus(structType any) BACnetLandingCallStatus {
	if casted, ok := structType.(BACnetLandingCallStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingCallStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingCallStatus) GetTypeName() string {
	return "BACnetLandingCallStatus"
}

func (m *_BACnetLandingCallStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumber)
	lengthInBits += m.FloorNumber.GetLengthInBits(ctx)

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	// Optional Field (floorText)
	if m.FloorText != nil {
		lengthInBits += m.FloorText.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetLandingCallStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLandingCallStatusParse(ctx context.Context, theBytes []byte) (BACnetLandingCallStatus, error) {
	return BACnetLandingCallStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLandingCallStatusParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingCallStatus, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingCallStatus, error) {
		return BACnetLandingCallStatusParseWithBuffer(ctx, readBuffer)
	}
}

func BACnetLandingCallStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLandingCallStatus, error) {
	v, err := (&_BACnetLandingCallStatus{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetLandingCallStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__bACnetLandingCallStatus BACnetLandingCallStatus, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingCallStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingCallStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "floorNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorNumber' field"))
	}
	m.FloorNumber = floorNumber

	command, err := ReadSimpleField[BACnetLandingCallStatusCommand](ctx, "command", ReadComplex[BACnetLandingCallStatusCommand](BACnetLandingCallStatusCommandParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'command' field"))
	}
	m.Command = command

	var floorText BACnetContextTagCharacterString
	_floorText, err := ReadOptionalField[BACnetContextTagCharacterString](ctx, "floorText", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	if _floorText != nil {
		floorText = *_floorText
		m.FloorText = floorText
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingCallStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingCallStatus")
	}

	return m, nil
}

func (m *_BACnetLandingCallStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLandingCallStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetLandingCallStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingCallStatus")
	}

	if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "floorNumber", m.GetFloorNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'floorNumber' field")
	}

	if err := WriteSimpleField[BACnetLandingCallStatusCommand](ctx, "command", m.GetCommand(), WriteComplex[BACnetLandingCallStatusCommand](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'command' field")
	}

	if err := WriteOptionalField[BACnetContextTagCharacterString](ctx, "floorText", GetRef(m.GetFloorText()), WriteComplex[BACnetContextTagCharacterString](writeBuffer), true); err != nil {
		return errors.Wrap(err, "Error serializing 'floorText' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingCallStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingCallStatus")
	}
	return nil
}

func (m *_BACnetLandingCallStatus) IsBACnetLandingCallStatus() {}

func (m *_BACnetLandingCallStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
