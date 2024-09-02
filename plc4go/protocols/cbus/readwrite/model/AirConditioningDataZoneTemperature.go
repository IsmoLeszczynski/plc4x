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

// AirConditioningDataZoneTemperature is the corresponding interface of AirConditioningDataZoneTemperature
type AirConditioningDataZoneTemperature interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetTemperature returns Temperature (property field)
	GetTemperature() HVACTemperature
	// GetSensorStatus returns SensorStatus (property field)
	GetSensorStatus() HVACSensorStatus
}

// _AirConditioningDataZoneTemperature is the data-structure of this message
type _AirConditioningDataZoneTemperature struct {
	AirConditioningDataContract
	ZoneGroup    byte
	ZoneList     HVACZoneList
	Temperature  HVACTemperature
	SensorStatus HVACSensorStatus
}

var _ AirConditioningDataZoneTemperature = (*_AirConditioningDataZoneTemperature)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneTemperature)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneTemperature) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneTemperature) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneTemperature) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneTemperature) GetTemperature() HVACTemperature {
	return m.Temperature
}

func (m *_AirConditioningDataZoneTemperature) GetSensorStatus() HVACSensorStatus {
	return m.SensorStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneTemperature factory function for _AirConditioningDataZoneTemperature
func NewAirConditioningDataZoneTemperature(zoneGroup byte, zoneList HVACZoneList, temperature HVACTemperature, sensorStatus HVACSensorStatus, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneTemperature {
	_result := &_AirConditioningDataZoneTemperature{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Temperature:                 temperature,
		SensorStatus:                sensorStatus,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneTemperature(structType any) AirConditioningDataZoneTemperature {
	if casted, ok := structType.(AirConditioningDataZoneTemperature); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneTemperature); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneTemperature) GetTypeName() string {
	return "AirConditioningDataZoneTemperature"
}

func (m *_AirConditioningDataZoneTemperature) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (temperature)
	lengthInBits += m.Temperature.GetLengthInBits(ctx)

	// Simple field (sensorStatus)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneTemperature) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneTemperature) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneTemperature AirConditioningDataZoneTemperature, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneTemperature"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneTemperature")
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

	temperature, err := ReadSimpleField[HVACTemperature](ctx, "temperature", ReadComplex[HVACTemperature](HVACTemperatureParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'temperature' field"))
	}
	m.Temperature = temperature

	sensorStatus, err := ReadEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", ReadEnum(HVACSensorStatusByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sensorStatus' field"))
	}
	m.SensorStatus = sensorStatus

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneTemperature"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneTemperature")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneTemperature) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneTemperature) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneTemperature"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneTemperature")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACTemperature](ctx, "temperature", m.GetTemperature(), WriteComplex[HVACTemperature](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'temperature' field")
		}

		if err := WriteSimpleEnumField[HVACSensorStatus](ctx, "sensorStatus", "HVACSensorStatus", m.GetSensorStatus(), WriteEnum[HVACSensorStatus, uint8](HVACSensorStatus.GetValue, HVACSensorStatus.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'sensorStatus' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneTemperature"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneTemperature")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneTemperature) IsAirConditioningDataZoneTemperature() {}

func (m *_AirConditioningDataZoneTemperature) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
