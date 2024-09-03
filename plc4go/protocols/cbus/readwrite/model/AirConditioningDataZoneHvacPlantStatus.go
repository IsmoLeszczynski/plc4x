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

// AirConditioningDataZoneHvacPlantStatus is the corresponding interface of AirConditioningDataZoneHvacPlantStatus
type AirConditioningDataZoneHvacPlantStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetHvacType returns HvacType (property field)
	GetHvacType() HVACType
	// GetHvacStatus returns HvacStatus (property field)
	GetHvacStatus() HVACStatusFlags
	// GetHvacErrorCode returns HvacErrorCode (property field)
	GetHvacErrorCode() HVACError
	// IsAirConditioningDataZoneHvacPlantStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataZoneHvacPlantStatus()
}

// _AirConditioningDataZoneHvacPlantStatus is the data-structure of this message
type _AirConditioningDataZoneHvacPlantStatus struct {
	AirConditioningDataContract
	ZoneGroup     byte
	ZoneList      HVACZoneList
	HvacType      HVACType
	HvacStatus    HVACStatusFlags
	HvacErrorCode HVACError
}

var _ AirConditioningDataZoneHvacPlantStatus = (*_AirConditioningDataZoneHvacPlantStatus)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataZoneHvacPlantStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacType() HVACType {
	return m.HvacType
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacStatus() HVACStatusFlags {
	return m.HvacStatus
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetHvacErrorCode() HVACError {
	return m.HvacErrorCode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAirConditioningDataZoneHvacPlantStatus factory function for _AirConditioningDataZoneHvacPlantStatus
func NewAirConditioningDataZoneHvacPlantStatus(zoneGroup byte, zoneList HVACZoneList, hvacType HVACType, hvacStatus HVACStatusFlags, hvacErrorCode HVACError, commandTypeContainer AirConditioningCommandTypeContainer) *_AirConditioningDataZoneHvacPlantStatus {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataZoneHvacPlantStatus must not be nil")
	}
	if hvacStatus == nil {
		panic("hvacStatus of type HVACStatusFlags for AirConditioningDataZoneHvacPlantStatus must not be nil")
	}
	_result := &_AirConditioningDataZoneHvacPlantStatus{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		HvacType:                    hvacType,
		HvacStatus:                  hvacStatus,
		HvacErrorCode:               hvacErrorCode,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAirConditioningDataZoneHvacPlantStatus(structType any) AirConditioningDataZoneHvacPlantStatus {
	if casted, ok := structType.(AirConditioningDataZoneHvacPlantStatus); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataZoneHvacPlantStatus); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetTypeName() string {
	return "AirConditioningDataZoneHvacPlantStatus"
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (hvacType)
	lengthInBits += 8

	// Simple field (hvacStatus)
	lengthInBits += m.HvacStatus.GetLengthInBits(ctx)

	// Simple field (hvacErrorCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *_AirConditioningDataZoneHvacPlantStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataZoneHvacPlantStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataZoneHvacPlantStatus AirConditioningDataZoneHvacPlantStatus, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataZoneHvacPlantStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataZoneHvacPlantStatus")
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

	hvacType, err := ReadEnumField[HVACType](ctx, "hvacType", "HVACType", ReadEnum(HVACTypeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacType' field"))
	}
	m.HvacType = hvacType

	hvacStatus, err := ReadSimpleField[HVACStatusFlags](ctx, "hvacStatus", ReadComplex[HVACStatusFlags](HVACStatusFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacStatus' field"))
	}
	m.HvacStatus = hvacStatus

	hvacErrorCode, err := ReadEnumField[HVACError](ctx, "hvacErrorCode", "HVACError", ReadEnum(HVACErrorByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacErrorCode' field"))
	}
	m.HvacErrorCode = hvacErrorCode

	if closeErr := readBuffer.CloseContext("AirConditioningDataZoneHvacPlantStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataZoneHvacPlantStatus")
	}

	return m, nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataZoneHvacPlantStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataZoneHvacPlantStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataZoneHvacPlantStatus")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleEnumField[HVACType](ctx, "hvacType", "HVACType", m.GetHvacType(), WriteEnum[HVACType, uint8](HVACType.GetValue, HVACType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacType' field")
		}

		if err := WriteSimpleField[HVACStatusFlags](ctx, "hvacStatus", m.GetHvacStatus(), WriteComplex[HVACStatusFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacStatus' field")
		}

		if err := WriteSimpleEnumField[HVACError](ctx, "hvacErrorCode", "HVACError", m.GetHvacErrorCode(), WriteEnum[HVACError, uint8](HVACError.GetValue, HVACError.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacErrorCode' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataZoneHvacPlantStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataZoneHvacPlantStatus")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataZoneHvacPlantStatus) IsAirConditioningDataZoneHvacPlantStatus() {}

func (m *_AirConditioningDataZoneHvacPlantStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
