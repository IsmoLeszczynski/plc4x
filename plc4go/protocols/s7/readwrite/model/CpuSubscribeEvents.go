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

// CpuSubscribeEvents is an enum
type CpuSubscribeEvents uint8

type ICpuSubscribeEvents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	CpuSubscribeEvents_CPU CpuSubscribeEvents = 0x01
	CpuSubscribeEvents_IM  CpuSubscribeEvents = 0x02
	CpuSubscribeEvents_FM  CpuSubscribeEvents = 0x04
	CpuSubscribeEvents_CP  CpuSubscribeEvents = 0x80
)

var CpuSubscribeEventsValues []CpuSubscribeEvents

func init() {
	_ = errors.New
	CpuSubscribeEventsValues = []CpuSubscribeEvents{
		CpuSubscribeEvents_CPU,
		CpuSubscribeEvents_IM,
		CpuSubscribeEvents_FM,
		CpuSubscribeEvents_CP,
	}
}

func CpuSubscribeEventsByValue(value uint8) (enum CpuSubscribeEvents, ok bool) {
	switch value {
	case 0x01:
		return CpuSubscribeEvents_CPU, true
	case 0x02:
		return CpuSubscribeEvents_IM, true
	case 0x04:
		return CpuSubscribeEvents_FM, true
	case 0x80:
		return CpuSubscribeEvents_CP, true
	}
	return 0, false
}

func CpuSubscribeEventsByName(value string) (enum CpuSubscribeEvents, ok bool) {
	switch value {
	case "CPU":
		return CpuSubscribeEvents_CPU, true
	case "IM":
		return CpuSubscribeEvents_IM, true
	case "FM":
		return CpuSubscribeEvents_FM, true
	case "CP":
		return CpuSubscribeEvents_CP, true
	}
	return 0, false
}

func CpuSubscribeEventsKnows(value uint8) bool {
	for _, typeValue := range CpuSubscribeEventsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCpuSubscribeEvents(structType any) CpuSubscribeEvents {
	castFunc := func(typ any) CpuSubscribeEvents {
		if sCpuSubscribeEvents, ok := typ.(CpuSubscribeEvents); ok {
			return sCpuSubscribeEvents
		}
		return 0
	}
	return castFunc(structType)
}

func (m CpuSubscribeEvents) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m CpuSubscribeEvents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CpuSubscribeEventsParse(ctx context.Context, theBytes []byte) (CpuSubscribeEvents, error) {
	return CpuSubscribeEventsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CpuSubscribeEventsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CpuSubscribeEvents, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("CpuSubscribeEvents", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CpuSubscribeEvents")
	}
	if enum, ok := CpuSubscribeEventsByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for CpuSubscribeEvents")
		return CpuSubscribeEvents(val), nil
	} else {
		return enum, nil
	}
}

func (e CpuSubscribeEvents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CpuSubscribeEvents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("CpuSubscribeEvents", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CpuSubscribeEvents) PLC4XEnumName() string {
	switch e {
	case CpuSubscribeEvents_CPU:
		return "CPU"
	case CpuSubscribeEvents_IM:
		return "IM"
	case CpuSubscribeEvents_FM:
		return "FM"
	case CpuSubscribeEvents_CP:
		return "CP"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e CpuSubscribeEvents) String() string {
	return e.PLC4XEnumName()
}
