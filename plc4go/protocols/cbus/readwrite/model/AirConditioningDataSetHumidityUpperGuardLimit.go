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

// AirConditioningDataSetHumidityUpperGuardLimit is the corresponding interface of AirConditioningDataSetHumidityUpperGuardLimit
type AirConditioningDataSetHumidityUpperGuardLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACHumidity
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACHumidityModeAndFlags
}

// _AirConditioningDataSetHumidityUpperGuardLimit is the data-structure of this message
type _AirConditioningDataSetHumidityUpperGuardLimit struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACHumidity
	HvacModeAndFlags HVACHumidityModeAndFlags
}

var _ AirConditioningDataSetHumidityUpperGuardLimit = (*_AirConditioningDataSetHumidityUpperGuardLimit)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetHumidityUpperGuardLimit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLimit() HVACHumidity {
	return m.Limit
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetHvacModeAndFlags() HVACHumidityModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHumidityUpperGuardLimit factory function for _AirConditioningDataSetHumidityUpperGuardLimit
func NewAirConditioningDataSetHumidityUpperGuardLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHumidityUpperGuardLimit {
	_result := &_AirConditioningDataSetHumidityUpperGuardLimit{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Limit:                       limit,
		HvacModeAndFlags:            hvacModeAndFlags,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHumidityUpperGuardLimit(structType any) AirConditioningDataSetHumidityUpperGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHumidityUpperGuardLimit"
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetHumidityUpperGuardLimit AirConditioningDataSetHumidityUpperGuardLimit, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHumidityUpperGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHumidityUpperGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	limit, err := ReadSimpleField[HVACHumidity](ctx, "limit", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limit' field"))
	}
	m.Limit = limit

	hvacModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHumidityUpperGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHumidityUpperGuardLimit")
	}

	return m, nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHumidityUpperGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHumidityUpperGuardLimit")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidity](ctx, "limit", m.GetLimit(), WriteComplex[HVACHumidity](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limit' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHumidityUpperGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHumidityUpperGuardLimit")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) IsAirConditioningDataSetHumidityUpperGuardLimit() {
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
