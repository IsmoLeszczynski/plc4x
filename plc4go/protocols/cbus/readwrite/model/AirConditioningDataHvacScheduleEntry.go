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

// AirConditioningDataHvacScheduleEntry is the corresponding interface of AirConditioningDataHvacScheduleEntry
type AirConditioningDataHvacScheduleEntry interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetEntry returns Entry (property field)
	GetEntry() uint8
	// GetFormat returns Format (property field)
	GetFormat() byte
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACModeAndFlags
	// GetStartTime returns StartTime (property field)
	GetStartTime() HVACStartTime
	// GetLevel returns Level (property field)
	GetLevel() HVACTemperature
	// GetRawLevel returns RawLevel (property field)
	GetRawLevel() HVACRawLevels
}

// AirConditioningDataHvacScheduleEntryExactly can be used when we want exactly this type and not a type which fulfills AirConditioningDataHvacScheduleEntry.
// This is useful for switch cases.
type AirConditioningDataHvacScheduleEntryExactly interface {
	AirConditioningDataHvacScheduleEntry
	isAirConditioningDataHvacScheduleEntry() bool
}

// _AirConditioningDataHvacScheduleEntry is the data-structure of this message
type _AirConditioningDataHvacScheduleEntry struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Entry            uint8
	Format           byte
	HvacModeAndFlags HVACModeAndFlags
	StartTime        HVACStartTime
	Level            HVACTemperature
	RawLevel         HVACRawLevels
}

var _ AirConditioningDataHvacScheduleEntry = (*_AirConditioningDataHvacScheduleEntry)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataHvacScheduleEntry)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataHvacScheduleEntry) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataHvacScheduleEntry) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataHvacScheduleEntry) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataHvacScheduleEntry) GetEntry() uint8 {
	return m.Entry
}

func (m *_AirConditioningDataHvacScheduleEntry) GetFormat() byte {
	return m.Format
}

func (m *_AirConditioningDataHvacScheduleEntry) GetHvacModeAndFlags() HVACModeAndFlags {
	return m.HvacModeAndFlags
}

func (m *_AirConditioningDataHvacScheduleEntry) GetStartTime() HVACStartTime {
	return m.StartTime
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLevel() HVACTemperature {
	return m.Level
}

func (m *_AirConditioningDataHvacScheduleEntry) GetRawLevel() HVACRawLevels {
	return m.RawLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataHvacScheduleEntry factory function for _AirConditioningDataHvacScheduleEntry
func NewAirConditioningDataHvacScheduleEntry(zoneGroup byte, zoneList HVACZoneList, entry uint8, format byte, hvacModeAndFlags HVACModeAndFlags, startTime HVACStartTime, level HVACTemperature, rawLevel HVACRawLevels, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataHvacScheduleEntry {
	_result := &_AirConditioningDataHvacScheduleEntry{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Entry:                       entry,
		Format:                      format,
		HvacModeAndFlags:            hvacModeAndFlags,
		StartTime:                   startTime,
		Level:                       level,
		RawLevel:                    rawLevel,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataHvacScheduleEntry(structType any) AirConditioningDataHvacScheduleEntry {
	if casted, ok := structType.(AirConditioningDataHvacScheduleEntry); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataHvacScheduleEntry); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataHvacScheduleEntry) GetTypeName() string {
	return "AirConditioningDataHvacScheduleEntry"
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (entry)
	lengthInBits += 8

	// Simple field (format)
	lengthInBits += 8

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	// Simple field (startTime)
	lengthInBits += m.StartTime.GetLengthInBits(ctx)

	// Optional Field (level)
	if m.Level != nil {
		lengthInBits += m.Level.GetLengthInBits(ctx)
	}

	// Optional Field (rawLevel)
	if m.RawLevel != nil {
		lengthInBits += m.RawLevel.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_AirConditioningDataHvacScheduleEntry) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataHvacScheduleEntry) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataHvacScheduleEntry AirConditioningDataHvacScheduleEntry, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataHvacScheduleEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataHvacScheduleEntry")
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

	entry, err := ReadSimpleField(ctx, "entry", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'entry' field"))
	}
	m.Entry = entry

	format, err := ReadSimpleField(ctx, "format", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'format' field"))
	}
	m.Format = format

	hvacModeAndFlags, err := ReadSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACModeAndFlags](HVACModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	startTime, err := ReadSimpleField[HVACStartTime](ctx, "startTime", ReadComplex[HVACStartTime](HVACStartTimeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	var level HVACTemperature
	_level, err := ReadOptionalField[HVACTemperature](ctx, "level", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelTemperature())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'level' field"))
	}
	if _level != nil {
		level = *_level
		m.Level = level
	}

	var rawLevel HVACRawLevels
	_rawLevel, err := ReadOptionalField[HVACRawLevels](ctx, "rawLevel", ReadComplex[HVACRawLevels](HVACRawLevelsParseWithBuffer, readBuffer), hvacModeAndFlags.GetIsLevelRaw())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'rawLevel' field"))
	}
	if _rawLevel != nil {
		rawLevel = *_rawLevel
		m.RawLevel = rawLevel
	}

	if closeErr := readBuffer.CloseContext("AirConditioningDataHvacScheduleEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataHvacScheduleEntry")
	}

	return m, nil
}

func (m *_AirConditioningDataHvacScheduleEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataHvacScheduleEntry) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataHvacScheduleEntry"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataHvacScheduleEntry")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[uint8](ctx, "entry", m.GetEntry(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'entry' field")
		}

		if err := WriteSimpleField[byte](ctx, "format", m.GetFormat(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'format' field")
		}

		if err := WriteSimpleField[HVACModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if err := WriteSimpleField[HVACStartTime](ctx, "startTime", m.GetStartTime(), WriteComplex[HVACStartTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteOptionalField[HVACTemperature](ctx, "level", GetRef(m.GetLevel()), WriteComplex[HVACTemperature](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'level' field")
		}

		if err := WriteOptionalField[HVACRawLevels](ctx, "rawLevel", GetRef(m.GetRawLevel()), WriteComplex[HVACRawLevels](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'rawLevel' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataHvacScheduleEntry"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataHvacScheduleEntry")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataHvacScheduleEntry) isAirConditioningDataHvacScheduleEntry() bool {
	return true
}

func (m *_AirConditioningDataHvacScheduleEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
