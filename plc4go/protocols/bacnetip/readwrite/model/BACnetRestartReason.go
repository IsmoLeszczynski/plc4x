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

// BACnetRestartReason is an enum
type BACnetRestartReason uint8

type IBACnetRestartReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetRestartReason_UNKNOWN                  BACnetRestartReason = 0
	BACnetRestartReason_COLDSTART                BACnetRestartReason = 1
	BACnetRestartReason_WARMSTART                BACnetRestartReason = 2
	BACnetRestartReason_DETECTED_POWER_LOST      BACnetRestartReason = 3
	BACnetRestartReason_DETECTED_POWERED_OFF     BACnetRestartReason = 4
	BACnetRestartReason_HARDWARE_WATCHDOG        BACnetRestartReason = 5
	BACnetRestartReason_SOFTWARE_WATCHDOG        BACnetRestartReason = 6
	BACnetRestartReason_SUSPENDED                BACnetRestartReason = 7
	BACnetRestartReason_ACTIVATE_CHANGES         BACnetRestartReason = 8
	BACnetRestartReason_VENDOR_PROPRIETARY_VALUE BACnetRestartReason = 0xFF
)

var BACnetRestartReasonValues []BACnetRestartReason

func init() {
	_ = errors.New
	BACnetRestartReasonValues = []BACnetRestartReason{
		BACnetRestartReason_UNKNOWN,
		BACnetRestartReason_COLDSTART,
		BACnetRestartReason_WARMSTART,
		BACnetRestartReason_DETECTED_POWER_LOST,
		BACnetRestartReason_DETECTED_POWERED_OFF,
		BACnetRestartReason_HARDWARE_WATCHDOG,
		BACnetRestartReason_SOFTWARE_WATCHDOG,
		BACnetRestartReason_SUSPENDED,
		BACnetRestartReason_ACTIVATE_CHANGES,
		BACnetRestartReason_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetRestartReasonByValue(value uint8) (enum BACnetRestartReason, ok bool) {
	switch value {
	case 0:
		return BACnetRestartReason_UNKNOWN, true
	case 0xFF:
		return BACnetRestartReason_VENDOR_PROPRIETARY_VALUE, true
	case 1:
		return BACnetRestartReason_COLDSTART, true
	case 2:
		return BACnetRestartReason_WARMSTART, true
	case 3:
		return BACnetRestartReason_DETECTED_POWER_LOST, true
	case 4:
		return BACnetRestartReason_DETECTED_POWERED_OFF, true
	case 5:
		return BACnetRestartReason_HARDWARE_WATCHDOG, true
	case 6:
		return BACnetRestartReason_SOFTWARE_WATCHDOG, true
	case 7:
		return BACnetRestartReason_SUSPENDED, true
	case 8:
		return BACnetRestartReason_ACTIVATE_CHANGES, true
	}
	return 0, false
}

func BACnetRestartReasonByName(value string) (enum BACnetRestartReason, ok bool) {
	switch value {
	case "UNKNOWN":
		return BACnetRestartReason_UNKNOWN, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetRestartReason_VENDOR_PROPRIETARY_VALUE, true
	case "COLDSTART":
		return BACnetRestartReason_COLDSTART, true
	case "WARMSTART":
		return BACnetRestartReason_WARMSTART, true
	case "DETECTED_POWER_LOST":
		return BACnetRestartReason_DETECTED_POWER_LOST, true
	case "DETECTED_POWERED_OFF":
		return BACnetRestartReason_DETECTED_POWERED_OFF, true
	case "HARDWARE_WATCHDOG":
		return BACnetRestartReason_HARDWARE_WATCHDOG, true
	case "SOFTWARE_WATCHDOG":
		return BACnetRestartReason_SOFTWARE_WATCHDOG, true
	case "SUSPENDED":
		return BACnetRestartReason_SUSPENDED, true
	case "ACTIVATE_CHANGES":
		return BACnetRestartReason_ACTIVATE_CHANGES, true
	}
	return 0, false
}

func BACnetRestartReasonKnows(value uint8) bool {
	for _, typeValue := range BACnetRestartReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRestartReason(structType any) BACnetRestartReason {
	castFunc := func(typ any) BACnetRestartReason {
		if sBACnetRestartReason, ok := typ.(BACnetRestartReason); ok {
			return sBACnetRestartReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRestartReason) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetRestartReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRestartReasonParse(ctx context.Context, theBytes []byte) (BACnetRestartReason, error) {
	return BACnetRestartReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRestartReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRestartReason, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetRestartReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetRestartReason")
	}
	if enum, ok := BACnetRestartReasonByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetRestartReason(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetRestartReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetRestartReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetRestartReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetRestartReason) PLC4XEnumName() string {
	switch e {
	case BACnetRestartReason_UNKNOWN:
		return "UNKNOWN"
	case BACnetRestartReason_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	case BACnetRestartReason_COLDSTART:
		return "COLDSTART"
	case BACnetRestartReason_WARMSTART:
		return "WARMSTART"
	case BACnetRestartReason_DETECTED_POWER_LOST:
		return "DETECTED_POWER_LOST"
	case BACnetRestartReason_DETECTED_POWERED_OFF:
		return "DETECTED_POWERED_OFF"
	case BACnetRestartReason_HARDWARE_WATCHDOG:
		return "HARDWARE_WATCHDOG"
	case BACnetRestartReason_SOFTWARE_WATCHDOG:
		return "SOFTWARE_WATCHDOG"
	case BACnetRestartReason_SUSPENDED:
		return "SUSPENDED"
	case BACnetRestartReason_ACTIVATE_CHANGES:
		return "ACTIVATE_CHANGES"
	}
	return ""
}

func (e BACnetRestartReason) String() string {
	return e.PLC4XEnumName()
}
