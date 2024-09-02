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

// MonitoredItemNotification is the corresponding interface of MonitoredItemNotification
type MonitoredItemNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetClientHandle returns ClientHandle (property field)
	GetClientHandle() uint32
	// GetValue returns Value (property field)
	GetValue() DataValue
}

// MonitoredItemNotificationExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemNotification.
// This is useful for switch cases.
type MonitoredItemNotificationExactly interface {
	MonitoredItemNotification
	isMonitoredItemNotification() bool
}

// _MonitoredItemNotification is the data-structure of this message
type _MonitoredItemNotification struct {
	ExtensionObjectDefinitionContract
	ClientHandle uint32
	Value        DataValue
}

var _ MonitoredItemNotification = (*_MonitoredItemNotification)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_MonitoredItemNotification)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemNotification) GetIdentifier() string {
	return "808"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemNotification) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemNotification) GetClientHandle() uint32 {
	return m.ClientHandle
}

func (m *_MonitoredItemNotification) GetValue() DataValue {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemNotification factory function for _MonitoredItemNotification
func NewMonitoredItemNotification(clientHandle uint32, value DataValue) *_MonitoredItemNotification {
	_result := &_MonitoredItemNotification{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ClientHandle:                      clientHandle,
		Value:                             value,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemNotification(structType any) MonitoredItemNotification {
	if casted, ok := structType.(MonitoredItemNotification); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemNotification); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemNotification) GetTypeName() string {
	return "MonitoredItemNotification"
}

func (m *_MonitoredItemNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (clientHandle)
	lengthInBits += 32

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredItemNotification) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__monitoredItemNotification MonitoredItemNotification, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredItemNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	clientHandle, err := ReadSimpleField(ctx, "clientHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientHandle' field"))
	}
	m.ClientHandle = clientHandle

	value, err := ReadSimpleField[DataValue](ctx, "value", ReadComplex[DataValue](DataValueParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("MonitoredItemNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemNotification")
	}

	return m, nil
}

func (m *_MonitoredItemNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemNotification")
		}

		if err := WriteSimpleField[uint32](ctx, "clientHandle", m.GetClientHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientHandle' field")
		}

		if err := WriteSimpleField[DataValue](ctx, "value", m.GetValue(), WriteComplex[DataValue](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemNotification")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemNotification) isMonitoredItemNotification() bool {
	return true
}

func (m *_MonitoredItemNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
