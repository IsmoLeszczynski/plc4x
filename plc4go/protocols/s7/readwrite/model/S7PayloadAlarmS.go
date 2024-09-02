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

// S7PayloadAlarmS is the corresponding interface of S7PayloadAlarmS
type S7PayloadAlarmS interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetAlarmMessage returns AlarmMessage (property field)
	GetAlarmMessage() AlarmMessagePushType
}

// _S7PayloadAlarmS is the data-structure of this message
type _S7PayloadAlarmS struct {
	S7PayloadUserDataItemContract
	AlarmMessage AlarmMessagePushType
}

var _ S7PayloadAlarmS = (*_S7PayloadAlarmS)(nil)
var _ S7PayloadUserDataItemRequirements = (*_S7PayloadAlarmS)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadAlarmS) GetCpuFunctionGroup() uint8 {
	return 0x04
}

func (m *_S7PayloadAlarmS) GetCpuFunctionType() uint8 {
	return 0x00
}

func (m *_S7PayloadAlarmS) GetCpuSubfunction() uint8 {
	return 0x12
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadAlarmS) GetParent() S7PayloadUserDataItemContract {
	return m.S7PayloadUserDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadAlarmS) GetAlarmMessage() AlarmMessagePushType {
	return m.AlarmMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadAlarmS factory function for _S7PayloadAlarmS
func NewS7PayloadAlarmS(alarmMessage AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize, dataLength uint16) *_S7PayloadAlarmS {
	_result := &_S7PayloadAlarmS{
		S7PayloadUserDataItemContract: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
		AlarmMessage:                  alarmMessage,
	}
	_result.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadAlarmS(structType any) S7PayloadAlarmS {
	if casted, ok := structType.(S7PayloadAlarmS); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadAlarmS); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadAlarmS) GetTypeName() string {
	return "S7PayloadAlarmS"
}

func (m *_S7PayloadAlarmS) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).getLengthInBits(ctx))

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_S7PayloadAlarmS) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadAlarmS) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7PayloadUserDataItem, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (__s7PayloadAlarmS S7PayloadAlarmS, err error) {
	m.S7PayloadUserDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadAlarmS"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadAlarmS")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmMessage, err := ReadSimpleField[AlarmMessagePushType](ctx, "alarmMessage", ReadComplex[AlarmMessagePushType](AlarmMessagePushTypeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmMessage' field"))
	}
	m.AlarmMessage = alarmMessage

	if closeErr := readBuffer.CloseContext("S7PayloadAlarmS"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadAlarmS")
	}

	return m, nil
}

func (m *_S7PayloadAlarmS) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadAlarmS) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarmS"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadAlarmS")
		}

		if err := WriteSimpleField[AlarmMessagePushType](ctx, "alarmMessage", m.GetAlarmMessage(), WriteComplex[AlarmMessagePushType](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarmS"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadAlarmS")
		}
		return nil
	}
	return m.S7PayloadUserDataItemContract.(*_S7PayloadUserDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadAlarmS) IsS7PayloadAlarmS() {}

func (m *_S7PayloadAlarmS) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
