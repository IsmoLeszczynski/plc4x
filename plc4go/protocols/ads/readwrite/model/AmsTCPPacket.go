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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AmsTCPPacket is the corresponding interface of AmsTCPPacket
type AmsTCPPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetUserdata returns Userdata (property field)
	GetUserdata() AmsPacket
}

// AmsTCPPacketExactly can be used when we want exactly this type and not a type which fulfills AmsTCPPacket.
// This is useful for switch cases.
type AmsTCPPacketExactly interface {
	AmsTCPPacket
	isAmsTCPPacket() bool
}

// _AmsTCPPacket is the data-structure of this message
type _AmsTCPPacket struct {
	Userdata AmsPacket
	// Reserved Fields
	reservedField0 *uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AmsTCPPacket) GetUserdata() AmsPacket {
	return m.Userdata
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAmsTCPPacket factory function for _AmsTCPPacket
func NewAmsTCPPacket(userdata AmsPacket) *_AmsTCPPacket {
	return &_AmsTCPPacket{Userdata: userdata}
}

// Deprecated: use the interface for direct cast
func CastAmsTCPPacket(structType any) AmsTCPPacket {
	if casted, ok := structType.(AmsTCPPacket); ok {
		return casted
	}
	if casted, ok := structType.(*AmsTCPPacket); ok {
		return *casted
	}
	return nil
}

func (m *_AmsTCPPacket) GetTypeName() string {
	return "AmsTCPPacket"
}

func (m *_AmsTCPPacket) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (userdata)
	lengthInBits += m.Userdata.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AmsTCPPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AmsTCPPacketParse(ctx context.Context, theBytes []byte) (AmsTCPPacket, error) {
	return AmsTCPPacketParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.LittleEndian)))
}

func AmsTCPPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AmsTCPPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("AmsTCPPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AmsTCPPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint16
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of AmsTCPPacket")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]any{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AmsTCPPacket")
	}

	// Simple Field (userdata)
	if pullErr := readBuffer.PullContext("userdata"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for userdata")
	}
	_userdata, _userdataErr := AmsPacketParseWithBuffer(ctx, readBuffer)
	if _userdataErr != nil {
		return nil, errors.Wrap(_userdataErr, "Error parsing 'userdata' field of AmsTCPPacket")
	}
	userdata := _userdata.(AmsPacket)
	if closeErr := readBuffer.CloseContext("userdata"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for userdata")
	}

	if closeErr := readBuffer.CloseContext("AmsTCPPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AmsTCPPacket")
	}

	// Create the instance
	return &_AmsTCPPacket{
		Userdata:       userdata,
		reservedField0: reservedField0,
	}, nil
}

func (m *_AmsTCPPacket) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.LittleEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AmsTCPPacket) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("AmsTCPPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for AmsTCPPacket")
	}

	// Reserved Field (reserved)
	{
		var reserved uint16 = uint16(0x0000)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(reserved))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.GetUserdata().GetLengthInBytes(ctx))
	_lengthErr := writeBuffer.WriteUint32("length", 32, uint32((length)))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (userdata)
	if pushErr := writeBuffer.PushContext("userdata"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for userdata")
	}
	_userdataErr := writeBuffer.WriteSerializable(ctx, m.GetUserdata())
	if popErr := writeBuffer.PopContext("userdata"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for userdata")
	}
	if _userdataErr != nil {
		return errors.Wrap(_userdataErr, "Error serializing 'userdata' field")
	}

	if popErr := writeBuffer.PopContext("AmsTCPPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for AmsTCPPacket")
	}
	return nil
}

func (m *_AmsTCPPacket) isAmsTCPPacket() bool {
	return true
}

func (m *_AmsTCPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
