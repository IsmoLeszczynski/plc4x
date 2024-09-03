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

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusRtuADU is the corresponding interface of ModbusRtuADU
type ModbusRtuADU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusADU
	// GetAddress returns Address (property field)
	GetAddress() uint8
	// GetPdu returns Pdu (property field)
	GetPdu() ModbusPDU
	// IsModbusRtuADU is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsModbusRtuADU()
}

// _ModbusRtuADU is the data-structure of this message
type _ModbusRtuADU struct {
	ModbusADUContract
	Address uint8
	Pdu     ModbusPDU
}

var _ ModbusRtuADU = (*_ModbusRtuADU)(nil)
var _ ModbusADURequirements = (*_ModbusRtuADU)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusRtuADU) GetDriverType() DriverType {
	return DriverType_MODBUS_RTU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusRtuADU) GetParent() ModbusADUContract {
	return m.ModbusADUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusRtuADU) GetAddress() uint8 {
	return m.Address
}

func (m *_ModbusRtuADU) GetPdu() ModbusPDU {
	return m.Pdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusRtuADU factory function for _ModbusRtuADU
func NewModbusRtuADU(address uint8, pdu ModbusPDU, response bool) *_ModbusRtuADU {
	if pdu == nil {
		panic("pdu of type ModbusPDU for ModbusRtuADU must not be nil")
	}
	_result := &_ModbusRtuADU{
		ModbusADUContract: NewModbusADU(response),
		Address:           address,
		Pdu:               pdu,
	}
	_result.ModbusADUContract.(*_ModbusADU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusRtuADU(structType any) ModbusRtuADU {
	if casted, ok := structType.(ModbusRtuADU); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusRtuADU); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusRtuADU) GetTypeName() string {
	return "ModbusRtuADU"
}

func (m *_ModbusRtuADU) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusADUContract.(*_ModbusADU).getLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 8

	// Simple field (pdu)
	lengthInBits += m.Pdu.GetLengthInBits(ctx)

	// Checksum Field (checksum)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusRtuADU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusRtuADU) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusADU, driverType DriverType, response bool) (__modbusRtuADU ModbusRtuADU, err error) {
	m.ModbusADUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusRtuADU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusRtuADU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	address, err := ReadSimpleField(ctx, "address", ReadUnsignedByte(readBuffer, uint8(8)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	pdu, err := ReadSimpleField[ModbusPDU](ctx, "pdu", ReadComplex[ModbusPDU](ModbusPDUParseWithBufferProducer[ModbusPDU]((bool)(response)), readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'pdu' field"))
	}
	m.Pdu = pdu

	crc, err := ReadChecksumField[uint16](ctx, "crc", ReadUnsignedShort(readBuffer, uint8(16)), RtuCrcCheck(ctx, address, pdu), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'crc' field"))
	}
	_ = crc

	if closeErr := readBuffer.CloseContext("ModbusRtuADU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusRtuADU")
	}

	return m, nil
}

func (m *_ModbusRtuADU) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusRtuADU) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusRtuADU"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusRtuADU")
		}

		if err := WriteSimpleField[uint8](ctx, "address", m.GetAddress(), WriteUnsignedByte(writeBuffer, 8), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[ModbusPDU](ctx, "pdu", m.GetPdu(), WriteComplex[ModbusPDU](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'pdu' field")
		}

		if err := WriteChecksumField[uint16](ctx, "crc", RtuCrcCheck(ctx, m.GetAddress(), m.GetPdu()), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'crc' field")
		}

		if popErr := writeBuffer.PopContext("ModbusRtuADU"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusRtuADU")
		}
		return nil
	}
	return m.ModbusADUContract.(*_ModbusADU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusRtuADU) IsModbusRtuADU() {}

func (m *_ModbusRtuADU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
