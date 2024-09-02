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

// AirConditioningDataSetZoneHumidityMode is the corresponding interface of AirConditioningDataSetZoneHumidityMode
type AirConditioningDataSetZoneHumidityMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHumidityModeAndFlags returns HumidityModeAndFlags (property field)
	GetHumidityModeAndFlags() HVACHumidityModeAndFlags
	// GetHumidityType returns HumidityType (property field)
	GetHumidityType() HVACHumidityType
	// GetLevel returns Level (property field)
	GetLevel() HVACHumidity
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
	// GetAuxLevel returns AuxLevel (property field)
	GetAuxLevel() HVACAuxiliaryLevel
}

// _AirConditioningDataSetZoneHumidityMode is the data-structure of this message
type _AirConditioningDataSetZoneHumidityMode struct {
	AirConditioningDataContract
	ZoneGroup            byte
	ZoneList             HVACZoneList
	HumidityModeAndFlags HVACHumidityModeAndFlags
	HumidityType         HVACHumidityType
	Level                HVACHumidity
	RawLevel             HVACRawLevels
	AuxLevel             HVACAuxiliaryLevel
}

var _ AirConditioningDataSetZoneHumidityMode = (*_AirConditioningDataSetZoneHumidityMode)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetZoneHumidityMode)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataSetZoneHumidityMode) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetZoneHumidityMode) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetHumidityModeAndFlags() HVACHumidityModeAndFlags {
	return m.HumidityModeAndFlags
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetHumidityType() HVACHumidityType {
	return m.HumidityType
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLevel() HVACHumidity {
	return m.Level
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetAuxLevel() HVACAuxiliaryLevel {
	return m.AuxLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataSetZoneHumidityMode factory function for _AirConditioningDataSetZoneHumidityMode
func NewAirConditioningDataSetZoneHumidityMode(zoneGroup byte, zoneList HVACZoneList, humidityModeAndFlags HVACHumidityModeAndFlags, humidityType HVACHumidityType, level HVACHumidity, rawLevel HVACRawLevels, auxLevel HVACAuxiliaryLevel, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataSetZoneHumidityMode {
	_result := &_AirConditioningDataSetZoneHumidityMode{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HumidityModeAndFlags:        humidityModeAndFlags,
		HumidityType:                humidityType,
		Level:                       level,
		RawLevel:                    rawLevel,
		AuxLevel:                    auxLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetZoneHumidityMode(structType any) AirConditioningDataSetZoneHumidityMode {
	if casted, ok := structType.(AirConditioningDataSetZoneHumidityMode); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetZoneHumidityMode); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetTypeName() string {
	return "AirConditioningDataSetZoneHumidityMode"
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (humidityModeAndFlags)
	lengthInBits += m.HumidityModeAndFlags.GetLengthInBits(ctx)

	// Simple field (humidityType)
	lengthInBits += 8

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	// Optional Field (auxLevel)
	if m.AuxLevel != nil {
		lengthInBits += m.AuxLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataSetZoneHumidityMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetZoneHumidityMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetZoneHumidityMode AirConditioningDataSetZoneHumidityMode, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetZoneHumidityMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetZoneHumidityMode")
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

	humidityModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityModeAndFlags' field"))
	}
	m.HumidityModeAndFlags = humidityModeAndFlags

	humidityType, err := ReadEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", ReadEnum(HVACHumidityTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'humidityType' field"))
	}
	m.HumidityType = humidityType

	var level HVACHumidity
	_level, err := ReadOptionalField[HVACHumidity](ctx, "level", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelHumidity())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	var auxLevel HVACAuxiliaryLevel
	_auxLevel, err := ReadOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", ReadComplex[HVACAuxiliaryLevel](HVACAuxiliaryLevelParseWithBuffer, readBuffer), humidityModeAndFlags.GetIsAuxLevelUsed())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auxLevel' field"))
	}
	if _auxLevel != nil {
		auxLevel = *_auxLevel
		m.AuxLevel = auxLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetZoneHumidityMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetZoneHumidityMode")
	}

	return m, nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetZoneHumidityMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetZoneHumidityMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetZoneHumidityMode")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "humidityModeAndFlags", m.GetHumidityModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityModeAndFlags' field")
		}

		if err := WriteSimpleEnumField[HVACHumidityType](ctx, "humidityType", "HVACHumidityType", m.GetHumidityType(), WriteEnum[HVACHumidityType, uint8](HVACHumidityType.GetValue, HVACHumidityType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'humidityType' field")
		}

		if err := WriteOptionalField[HVACHumidity](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACHumidity](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if err := WriteOptionalField[HVACAuxiliaryLevel](ctx, "auxLevel", GetRef(m.GetAuxLevel()), WriteComplex[HVACAuxiliaryLevel](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'auxLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetZoneHumidityMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetZoneHumidityMode")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetZoneHumidityMode) IsAirConditioningDataSetZoneHumidityMode() {}

func (m *_AirConditioningDataSetZoneHumidityMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
