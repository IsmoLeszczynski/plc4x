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

// SerialNumber is the corresponding interface of SerialNumber
type SerialNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOctet1 returns Octet1 (property field)
	GetOctet1() byte
	// GetOctet2 returns Octet2 (property field)
	GetOctet2() byte
	// GetOctet3 returns Octet3 (property field)
	GetOctet3() byte
	// GetOctet4 returns Octet4 (property field)
	GetOctet4() byte
}

// _SerialNumber is the data-structure of this message
type _SerialNumber struct {
	Octet1 byte
	Octet2 byte
	Octet3 byte
	Octet4 byte
}

var _ SerialNumber = (*_SerialNumber)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SerialNumber) GetOctet1() byte {
	return m.Octet1
}

func (m *_SerialNumber) GetOctet2() byte {
	return m.Octet2
}

func (m *_SerialNumber) GetOctet3() byte {
	return m.Octet3
}

func (m *_SerialNumber) GetOctet4() byte {
	return m.Octet4
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSerialNumber factory function for _SerialNumber
func NewSerialNumber(octet1 byte, octet2 byte, octet3 byte, octet4 byte) *_SerialNumber {
	return &_SerialNumber{Octet1: octet1, Octet2: octet2, Octet3: octet3, Octet4: octet4}
}

// Deprecated: use the interface for direct cast
func CastSerialNumber(structType any) SerialNumber {
	if casted, ok := structType.(SerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*SerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_SerialNumber) GetTypeName() string {
	return "SerialNumber"
}

func (m *_SerialNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (octet1)
	lengthInBits += 8

	// Simple field (octet2)
	lengthInBits += 8

	// Simple field (octet3)
	lengthInBits += 8

	// Simple field (octet4)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SerialNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SerialNumberParse(ctx context.Context, theBytes []byte) (SerialNumber, error) {
	return SerialNumberParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SerialNumberParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (SerialNumber, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (SerialNumber, error) {
		return SerialNumberParseWithBuffer(ctx, readBuffer)
	}
}

func SerialNumberParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SerialNumber, error) {
	v, err := (&_SerialNumber{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_SerialNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__serialNumber SerialNumber, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octet1, err := ReadSimpleField(ctx, "octet1", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet1' field"))
	}
	m.Octet1 = octet1

	octet2, err := ReadSimpleField(ctx, "octet2", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet2' field"))
	}
	m.Octet2 = octet2

	octet3, err := ReadSimpleField(ctx, "octet3", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet3' field"))
	}
	m.Octet3 = octet3

	octet4, err := ReadSimpleField(ctx, "octet4", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octet4' field"))
	}
	m.Octet4 = octet4

	if closeErr := readBuffer.CloseContext("SerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SerialNumber")
	}

	return m, nil
}

func (m *_SerialNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SerialNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("SerialNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for SerialNumber")
	}

	if err := WriteSimpleField[byte](ctx, "octet1", m.GetOctet1(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet1' field")
	}

	if err := WriteSimpleField[byte](ctx, "octet2", m.GetOctet2(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet2' field")
	}

	if err := WriteSimpleField[byte](ctx, "octet3", m.GetOctet3(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet3' field")
	}

	if err := WriteSimpleField[byte](ctx, "octet4", m.GetOctet4(), WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'octet4' field")
	}

	if popErr := writeBuffer.PopContext("SerialNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for SerialNumber")
	}
	return nil
}

func (m *_SerialNumber) IsSerialNumber() {}

func (m *_SerialNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
