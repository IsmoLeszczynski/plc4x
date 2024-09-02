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

// InterfaceOptions1 is the corresponding interface of InterfaceOptions1
type InterfaceOptions1 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetIdmon returns Idmon (property field)
	GetIdmon() bool
	// GetMonitor returns Monitor (property field)
	GetMonitor() bool
	// GetSmart returns Smart (property field)
	GetSmart() bool
	// GetSrchk returns Srchk (property field)
	GetSrchk() bool
	// GetXonXoff returns XonXoff (property field)
	GetXonXoff() bool
	// GetConnect returns Connect (property field)
	GetConnect() bool
}

// _InterfaceOptions1 is the data-structure of this message
type _InterfaceOptions1 struct {
	Idmon   bool
	Monitor bool
	Smart   bool
	Srchk   bool
	XonXoff bool
	Connect bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
}

var _ InterfaceOptions1 = (*_InterfaceOptions1)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions1) GetIdmon() bool {
	return m.Idmon
}

func (m *_InterfaceOptions1) GetMonitor() bool {
	return m.Monitor
}

func (m *_InterfaceOptions1) GetSmart() bool {
	return m.Smart
}

func (m *_InterfaceOptions1) GetSrchk() bool {
	return m.Srchk
}

func (m *_InterfaceOptions1) GetXonXoff() bool {
	return m.XonXoff
}

func (m *_InterfaceOptions1) GetConnect() bool {
	return m.Connect
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInterfaceOptions1 factory function for _InterfaceOptions1
func NewInterfaceOptions1(idmon bool, monitor bool, smart bool, srchk bool, xonXoff bool, connect bool) *_InterfaceOptions1 {
	return &_InterfaceOptions1{Idmon: idmon, Monitor: monitor, Smart: smart, Srchk: srchk, XonXoff: xonXoff, Connect: connect}
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions1(structType any) InterfaceOptions1 {
	if casted, ok := structType.(InterfaceOptions1); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions1); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions1) GetTypeName() string {
	return "InterfaceOptions1"
}

func (m *_InterfaceOptions1) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (idmon)
	lengthInBits += 1

	// Simple field (monitor)
	lengthInBits += 1

	// Simple field (smart)
	lengthInBits += 1

	// Simple field (srchk)
	lengthInBits += 1

	// Simple field (xonXoff)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (connect)
	lengthInBits += 1

	return lengthInBits
}

func (m *_InterfaceOptions1) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceOptions1Parse(ctx context.Context, theBytes []byte) (InterfaceOptions1, error) {
	return InterfaceOptions1ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions1ParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1, error) {
		return InterfaceOptions1ParseWithBuffer(ctx, readBuffer)
	}
}

func InterfaceOptions1ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceOptions1, error) {
	v, err := (&_InterfaceOptions1{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_InterfaceOptions1) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__interfaceOptions1 InterfaceOptions1, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions1")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	idmon, err := ReadSimpleField(ctx, "idmon", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'idmon' field"))
	}
	m.Idmon = idmon

	monitor, err := ReadSimpleField(ctx, "monitor", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitor' field"))
	}
	m.Monitor = monitor

	smart, err := ReadSimpleField(ctx, "smart", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'smart' field"))
	}
	m.Smart = smart

	srchk, err := ReadSimpleField(ctx, "srchk", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'srchk' field"))
	}
	m.Srchk = srchk

	xonXoff, err := ReadSimpleField(ctx, "xonXoff", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'xonXoff' field"))
	}
	m.XonXoff = xonXoff

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadBoolean(readBuffer), bool(false))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	connect, err := ReadSimpleField(ctx, "connect", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connect' field"))
	}
	m.Connect = connect

	if closeErr := readBuffer.CloseContext("InterfaceOptions1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions1")
	}

	return m, nil
}

func (m *_InterfaceOptions1) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions1) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("InterfaceOptions1"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions1")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 1")
	}

	if err := WriteSimpleField[bool](ctx, "idmon", m.GetIdmon(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'idmon' field")
	}

	if err := WriteSimpleField[bool](ctx, "monitor", m.GetMonitor(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'monitor' field")
	}

	if err := WriteSimpleField[bool](ctx, "smart", m.GetSmart(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'smart' field")
	}

	if err := WriteSimpleField[bool](ctx, "srchk", m.GetSrchk(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'srchk' field")
	}

	if err := WriteSimpleField[bool](ctx, "xonXoff", m.GetXonXoff(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'xonXoff' field")
	}

	if err := WriteReservedField[bool](ctx, "reserved", bool(false), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'reserved' field number 2")
	}

	if err := WriteSimpleField[bool](ctx, "connect", m.GetConnect(), WriteBoolean(writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'connect' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions1"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions1")
	}
	return nil
}

func (m *_InterfaceOptions1) IsInterfaceOptions1() {}

func (m *_InterfaceOptions1) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
