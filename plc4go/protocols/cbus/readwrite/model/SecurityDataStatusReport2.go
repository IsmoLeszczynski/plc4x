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

// SecurityDataStatusReport2 is the corresponding interface of SecurityDataStatusReport2
type SecurityDataStatusReport2 interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneStatus returns ZoneStatus (property field)
	GetZoneStatus() []ZoneStatus
}

// SecurityDataStatusReport2Exactly can be used when we want exactly this type and not a type which fulfills SecurityDataStatusReport2.
// This is useful for switch cases.
type SecurityDataStatusReport2Exactly interface {
	SecurityDataStatusReport2
	isSecurityDataStatusReport2() bool
}

// _SecurityDataStatusReport2 is the data-structure of this message
type _SecurityDataStatusReport2 struct {
	*_SecurityData
	ZoneStatus []ZoneStatus
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataStatusReport2) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataStatusReport2) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataStatusReport2) GetZoneStatus() []ZoneStatus {
	return m.ZoneStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataStatusReport2 factory function for _SecurityDataStatusReport2
func NewSecurityDataStatusReport2(zoneStatus []ZoneStatus, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatusReport2 {
	_result := &_SecurityDataStatusReport2{
		ZoneStatus:    zoneStatus,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatusReport2(structType any) SecurityDataStatusReport2 {
	if casted, ok := structType.(SecurityDataStatusReport2); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatusReport2); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatusReport2) GetTypeName() string {
	return "SecurityDataStatusReport2"
}

func (m *_SecurityDataStatusReport2) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.ZoneStatus) > 0 {
		for _curItem, element := range m.ZoneStatus {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ZoneStatus), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_SecurityDataStatusReport2) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataStatusReport2Parse(ctx context.Context, theBytes []byte) (SecurityDataStatusReport2, error) {
	return SecurityDataStatusReport2ParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataStatusReport2ParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataStatusReport2, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityDataStatusReport2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatusReport2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (zoneStatus)
	if pullErr := readBuffer.PullContext("zoneStatus", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneStatus")
	}
	// Count array
	zoneStatus := make([]ZoneStatus, uint16(48))
	// This happens when the size is set conditional to 0
	if len(zoneStatus) == 0 {
		zoneStatus = nil
	}
	{
		_numItems := uint16(uint16(48))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ZoneStatusParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'zoneStatus' field of SecurityDataStatusReport2")
			}
			zoneStatus[_curItem] = _item.(ZoneStatus)
		}
	}
	if closeErr := readBuffer.CloseContext("zoneStatus", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneStatus")
	}

	if closeErr := readBuffer.CloseContext("SecurityDataStatusReport2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatusReport2")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataStatusReport2{
		_SecurityData: &_SecurityData{},
		ZoneStatus:    zoneStatus,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataStatusReport2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatusReport2) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatusReport2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatusReport2")
		}

		// Array Field (zoneStatus)
		if pushErr := writeBuffer.PushContext("zoneStatus", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zoneStatus")
		}
		for _curItem, _element := range m.GetZoneStatus() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetZoneStatus()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'zoneStatus' field")
			}
		}
		if popErr := writeBuffer.PopContext("zoneStatus", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zoneStatus")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatusReport2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatusReport2")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatusReport2) isSecurityDataStatusReport2() bool {
	return true
}

func (m *_SecurityDataStatusReport2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
