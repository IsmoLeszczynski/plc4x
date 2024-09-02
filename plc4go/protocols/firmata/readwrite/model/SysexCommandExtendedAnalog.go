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

// SysexCommandExtendedAnalog is the corresponding interface of SysexCommandExtendedAnalog
type SysexCommandExtendedAnalog interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// _SysexCommandExtendedAnalog is the data-structure of this message
type _SysexCommandExtendedAnalog struct {
	SysexCommandContract
}

var _ SysexCommandExtendedAnalog = (*_SysexCommandExtendedAnalog)(nil)
var _ SysexCommandRequirements = (*_SysexCommandExtendedAnalog)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandExtendedAnalog) GetCommandType() uint8 {
	return 0x6F
}

func (m *_SysexCommandExtendedAnalog) GetResponse() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandExtendedAnalog) GetParent() SysexCommandContract {
	return m.SysexCommandContract
}

// NewSysexCommandExtendedAnalog factory function for _SysexCommandExtendedAnalog
func NewSysexCommandExtendedAnalog() *_SysexCommandExtendedAnalog {
	_result := &_SysexCommandExtendedAnalog{
		SysexCommandContract: NewSysexCommand(),
	}
	_result.SysexCommandContract.(*_SysexCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandExtendedAnalog(structType any) SysexCommandExtendedAnalog {
	if casted, ok := structType.(SysexCommandExtendedAnalog); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandExtendedAnalog); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandExtendedAnalog) GetTypeName() string {
	return "SysexCommandExtendedAnalog"
}

func (m *_SysexCommandExtendedAnalog) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SysexCommandContract.(*_SysexCommand).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SysexCommandExtendedAnalog) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SysexCommandExtendedAnalog) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SysexCommand, response bool) (__sysexCommandExtendedAnalog SysexCommandExtendedAnalog, err error) {
	m.SysexCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandExtendedAnalog"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandExtendedAnalog")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandExtendedAnalog"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandExtendedAnalog")
	}

	return m, nil
}

func (m *_SysexCommandExtendedAnalog) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandExtendedAnalog) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandExtendedAnalog"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandExtendedAnalog")
		}

		if popErr := writeBuffer.PopContext("SysexCommandExtendedAnalog"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandExtendedAnalog")
		}
		return nil
	}
	return m.SysexCommandContract.(*_SysexCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SysexCommandExtendedAnalog) IsSysexCommandExtendedAnalog() {}

func (m *_SysexCommandExtendedAnalog) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
