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

// AirConditioningDataSetHumidityLowerGuardLimit is the corresponding interface of AirConditioningDataSetHumidityLowerGuardLimit
type AirConditioningDataSetHumidityLowerGuardLimit interface {
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

// _AirConditioningDataSetHumidityLowerGuardLimit is the data-structure of this message
type _AirConditioningDataSetHumidityLowerGuardLimit struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACHumidity
	HvacModeAndFlags HVACHumidityModeAndFlags
}

var _ AirConditioningDataSetHumidityLowerGuardLimit = (*_AirConditioningDataSetHumidityLowerGuardLimit)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetHumidityLowerGuardLimit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetLimit() HVACHumidity {
	return m.Limit
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetHvacModeAndFlags() HVACHumidityModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetHumidityLowerGuardLimit factory function for _AirConditioningDataSetHumidityLowerGuardLimit
func NewAirConditioningDataSetHumidityLowerGuardLimit(zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetHumidityLowerGuardLimit {
	_result := &_AirConditioningDataSetHumidityLowerGuardLimit{
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
func CastAirConditioningDataSetHumidityLowerGuardLimit(structType any) AirConditioningDataSetHumidityLowerGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHumidityLowerGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHumidityLowerGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHumidityLowerGuardLimit"
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetLengthInBits(ctx context.Context) uint16 {
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

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetHumidityLowerGuardLimit AirConditioningDataSetHumidityLowerGuardLimit, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHumidityLowerGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHumidityLowerGuardLimit")
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

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHumidityLowerGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHumidityLowerGuardLimit")
	}

	return m, nil
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHumidityLowerGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHumidityLowerGuardLimit")
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

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHumidityLowerGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHumidityLowerGuardLimit")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) IsAirConditioningDataSetHumidityLowerGuardLimit() {
}

func (m *_AirConditioningDataSetHumidityLowerGuardLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
