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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// ClockAndTimekeepingCommandTypeContainer is an enum
type ClockAndTimekeepingCommandTypeContainer uint8

type IClockAndTimekeepingCommandTypeContainer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NumBytes() uint8
	CommandType() ClockAndTimekeepingCommandType
}

const (
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_0Bytes ClockAndTimekeepingCommandTypeContainer = 0x08
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_1Bytes ClockAndTimekeepingCommandTypeContainer = 0x09
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_2Bytes ClockAndTimekeepingCommandTypeContainer = 0x0A
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_3Bytes ClockAndTimekeepingCommandTypeContainer = 0x0B
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_4Bytes ClockAndTimekeepingCommandTypeContainer = 0x0C
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_5Bytes ClockAndTimekeepingCommandTypeContainer = 0x0D
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_6Bytes ClockAndTimekeepingCommandTypeContainer = 0x0E
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_7Bytes ClockAndTimekeepingCommandTypeContainer = 0x0F
	ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandRequestRefresh               ClockAndTimekeepingCommandTypeContainer = 0x11
)

var ClockAndTimekeepingCommandTypeContainerValues []ClockAndTimekeepingCommandTypeContainer

func init() {
	_ = errors.New
	ClockAndTimekeepingCommandTypeContainerValues = []ClockAndTimekeepingCommandTypeContainer{
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_0Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_1Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_2Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_3Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_4Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_5Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_6Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_7Bytes,
		ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandRequestRefresh,
	}
}

func (e ClockAndTimekeepingCommandTypeContainer) NumBytes() uint8 {
	switch e {
	case 0x08:
		{ /* '0x08' */
			return 0
		}
	case 0x09:
		{ /* '0x09' */
			return 1
		}
	case 0x0A:
		{ /* '0x0A' */
			return 2
		}
	case 0x0B:
		{ /* '0x0B' */
			return 3
		}
	case 0x0C:
		{ /* '0x0C' */
			return 4
		}
	case 0x0D:
		{ /* '0x0D' */
			return 5
		}
	case 0x0E:
		{ /* '0x0E' */
			return 6
		}
	case 0x0F:
		{ /* '0x0F' */
			return 7
		}
	case 0x11:
		{ /* '0x11' */
			return 1
		}
	default:
		{
			return 0
		}
	}
}

func ClockAndTimekeepingCommandTypeContainerFirstEnumForFieldNumBytes(value uint8) (ClockAndTimekeepingCommandTypeContainer, error) {
	for _, sizeValue := range ClockAndTimekeepingCommandTypeContainerValues {
		if sizeValue.NumBytes() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumBytes not found", value)
}

func (e ClockAndTimekeepingCommandTypeContainer) CommandType() ClockAndTimekeepingCommandType {
	switch e {
	case 0x08:
		{ /* '0x08' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x09:
		{ /* '0x09' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0A:
		{ /* '0x0A' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0B:
		{ /* '0x0B' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0C:
		{ /* '0x0C' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0D:
		{ /* '0x0D' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0E:
		{ /* '0x0E' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x0F:
		{ /* '0x0F' */
			return ClockAndTimekeepingCommandType_UPDATE_NETWORK_VARIABLE
		}
	case 0x11:
		{ /* '0x11' */
			return ClockAndTimekeepingCommandType_REQUEST_REFRESH
		}
	default:
		{
			return 0
		}
	}
}

func ClockAndTimekeepingCommandTypeContainerFirstEnumForFieldCommandType(value ClockAndTimekeepingCommandType) (ClockAndTimekeepingCommandTypeContainer, error) {
	for _, sizeValue := range ClockAndTimekeepingCommandTypeContainerValues {
		if sizeValue.CommandType() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing CommandType not found", value)
}
func ClockAndTimekeepingCommandTypeContainerByValue(value uint8) (enum ClockAndTimekeepingCommandTypeContainer, ok bool) {
	switch value {
	case 0x08:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_0Bytes, true
	case 0x09:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_1Bytes, true
	case 0x0A:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_2Bytes, true
	case 0x0B:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_3Bytes, true
	case 0x0C:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_4Bytes, true
	case 0x0D:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_5Bytes, true
	case 0x0E:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_6Bytes, true
	case 0x0F:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_7Bytes, true
	case 0x11:
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandRequestRefresh, true
	}
	return 0, false
}

func ClockAndTimekeepingCommandTypeContainerByName(value string) (enum ClockAndTimekeepingCommandTypeContainer, ok bool) {
	switch value {
	case "MediaTransportControlCommandUpdateNetworkVariable_0Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_0Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_1Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_1Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_2Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_2Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_3Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_3Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_4Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_4Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_5Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_5Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_6Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_6Bytes, true
	case "MediaTransportControlCommandUpdateNetworkVariable_7Bytes":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_7Bytes, true
	case "MediaTransportControlCommandRequestRefresh":
		return ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandRequestRefresh, true
	}
	return 0, false
}

func ClockAndTimekeepingCommandTypeContainerKnows(value uint8) bool {
	for _, typeValue := range ClockAndTimekeepingCommandTypeContainerValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastClockAndTimekeepingCommandTypeContainer(structType any) ClockAndTimekeepingCommandTypeContainer {
	castFunc := func(typ any) ClockAndTimekeepingCommandTypeContainer {
		if sClockAndTimekeepingCommandTypeContainer, ok := typ.(ClockAndTimekeepingCommandTypeContainer); ok {
			return sClockAndTimekeepingCommandTypeContainer
		}
		return 0
	}
	return castFunc(structType)
}

func (m ClockAndTimekeepingCommandTypeContainer) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m ClockAndTimekeepingCommandTypeContainer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ClockAndTimekeepingCommandTypeContainerParse(ctx context.Context, theBytes []byte) (ClockAndTimekeepingCommandTypeContainer, error) {
	return ClockAndTimekeepingCommandTypeContainerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ClockAndTimekeepingCommandTypeContainerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ClockAndTimekeepingCommandTypeContainer, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("ClockAndTimekeepingCommandTypeContainer", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ClockAndTimekeepingCommandTypeContainer")
	}
	if enum, ok := ClockAndTimekeepingCommandTypeContainerByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ClockAndTimekeepingCommandTypeContainer")
		return ClockAndTimekeepingCommandTypeContainer(val), nil
	} else {
		return enum, nil
	}
}

func (e ClockAndTimekeepingCommandTypeContainer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ClockAndTimekeepingCommandTypeContainer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("ClockAndTimekeepingCommandTypeContainer", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ClockAndTimekeepingCommandTypeContainer) PLC4XEnumName() string {
	switch e {
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_0Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_0Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_1Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_1Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_2Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_2Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_3Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_3Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_4Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_4Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_5Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_5Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_6Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_6Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandUpdateNetworkVariable_7Bytes:
		return "MediaTransportControlCommandUpdateNetworkVariable_7Bytes"
	case ClockAndTimekeepingCommandTypeContainer_MediaTransportControlCommandRequestRefresh:
		return "MediaTransportControlCommandRequestRefresh"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e ClockAndTimekeepingCommandTypeContainer) String() string {
	return e.PLC4XEnumName()
}
