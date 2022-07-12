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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusDeviceInformationConformityLevel is an enum
type ModbusDeviceInformationConformityLevel uint8

type IModbusDeviceInformationConformityLevel interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	ModbusDeviceInformationConformityLevel_BASIC_STREAM_ONLY    ModbusDeviceInformationConformityLevel = 0x01
	ModbusDeviceInformationConformityLevel_REGULAR_STREAM_ONLY  ModbusDeviceInformationConformityLevel = 0x02
	ModbusDeviceInformationConformityLevel_EXTENDED_STREAM_ONLY ModbusDeviceInformationConformityLevel = 0x03
)

var ModbusDeviceInformationConformityLevelValues []ModbusDeviceInformationConformityLevel

func init() {
	_ = errors.New
	ModbusDeviceInformationConformityLevelValues = []ModbusDeviceInformationConformityLevel{
		ModbusDeviceInformationConformityLevel_BASIC_STREAM_ONLY,
		ModbusDeviceInformationConformityLevel_REGULAR_STREAM_ONLY,
		ModbusDeviceInformationConformityLevel_EXTENDED_STREAM_ONLY,
	}
}

func ModbusDeviceInformationConformityLevelByValue(value uint8) (enum ModbusDeviceInformationConformityLevel, ok bool) {
	switch value {
	case 0x01:
		return ModbusDeviceInformationConformityLevel_BASIC_STREAM_ONLY, true
	case 0x02:
		return ModbusDeviceInformationConformityLevel_REGULAR_STREAM_ONLY, true
	case 0x03:
		return ModbusDeviceInformationConformityLevel_EXTENDED_STREAM_ONLY, true
	}
	return 0, false
}

func ModbusDeviceInformationConformityLevelByName(value string) (enum ModbusDeviceInformationConformityLevel, ok bool) {
	switch value {
	case "BASIC_STREAM_ONLY":
		return ModbusDeviceInformationConformityLevel_BASIC_STREAM_ONLY, true
	case "REGULAR_STREAM_ONLY":
		return ModbusDeviceInformationConformityLevel_REGULAR_STREAM_ONLY, true
	case "EXTENDED_STREAM_ONLY":
		return ModbusDeviceInformationConformityLevel_EXTENDED_STREAM_ONLY, true
	}
	return 0, false
}

func ModbusDeviceInformationConformityLevelKnows(value uint8) bool {
	for _, typeValue := range ModbusDeviceInformationConformityLevelValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastModbusDeviceInformationConformityLevel(structType interface{}) ModbusDeviceInformationConformityLevel {
	castFunc := func(typ interface{}) ModbusDeviceInformationConformityLevel {
		if sModbusDeviceInformationConformityLevel, ok := typ.(ModbusDeviceInformationConformityLevel); ok {
			return sModbusDeviceInformationConformityLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModbusDeviceInformationConformityLevel) GetLengthInBits() uint16 {
	return 7
}

func (m ModbusDeviceInformationConformityLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusDeviceInformationConformityLevelParse(readBuffer utils.ReadBuffer) (ModbusDeviceInformationConformityLevel, error) {
	val, err := readBuffer.ReadUint8("ModbusDeviceInformationConformityLevel", 7)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ModbusDeviceInformationConformityLevel")
	}
	if enum, ok := ModbusDeviceInformationConformityLevelByValue(val); !ok {
		return 0, errors.Errorf("no value %v found for ModbusDeviceInformationConformityLevel", val)
	} else {
		return enum, nil
	}
}

func (e ModbusDeviceInformationConformityLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ModbusDeviceInformationConformityLevel", 7, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ModbusDeviceInformationConformityLevel) PLC4XEnumName() string {
	switch e {
	case ModbusDeviceInformationConformityLevel_BASIC_STREAM_ONLY:
		return "BASIC_STREAM_ONLY"
	case ModbusDeviceInformationConformityLevel_REGULAR_STREAM_ONLY:
		return "REGULAR_STREAM_ONLY"
	case ModbusDeviceInformationConformityLevel_EXTENDED_STREAM_ONLY:
		return "EXTENDED_STREAM_ONLY"
	}
	return ""
}

func (e ModbusDeviceInformationConformityLevel) String() string {
	return e.PLC4XEnumName()
}
