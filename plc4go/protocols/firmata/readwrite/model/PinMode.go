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

// PinMode is an enum
type PinMode uint8

type IPinMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PinMode_PinModeInput   PinMode = 0x0
	PinMode_PinModeOutput  PinMode = 0x1
	PinMode_PinModeAnalog  PinMode = 0x2
	PinMode_PinModePwm     PinMode = 0x3
	PinMode_PinModeServo   PinMode = 0x4
	PinMode_PinModeShift   PinMode = 0x5
	PinMode_PinModeI2C     PinMode = 0x6
	PinMode_PinModeOneWire PinMode = 0x7
	PinMode_PinModeStepper PinMode = 0x8
	PinMode_PinModeEncoder PinMode = 0x9
	PinMode_PinModeSerial  PinMode = 0xA
	PinMode_PinModePullup  PinMode = 0xB
)

var PinModeValues []PinMode

func init() {
	_ = errors.New
	PinModeValues = []PinMode{
		PinMode_PinModeInput,
		PinMode_PinModeOutput,
		PinMode_PinModeAnalog,
		PinMode_PinModePwm,
		PinMode_PinModeServo,
		PinMode_PinModeShift,
		PinMode_PinModeI2C,
		PinMode_PinModeOneWire,
		PinMode_PinModeStepper,
		PinMode_PinModeEncoder,
		PinMode_PinModeSerial,
		PinMode_PinModePullup,
	}
}

func PinModeByValue(value uint8) (enum PinMode, ok bool) {
	switch value {
	case 0x0:
		return PinMode_PinModeInput, true
	case 0x1:
		return PinMode_PinModeOutput, true
	case 0x2:
		return PinMode_PinModeAnalog, true
	case 0x3:
		return PinMode_PinModePwm, true
	case 0x4:
		return PinMode_PinModeServo, true
	case 0x5:
		return PinMode_PinModeShift, true
	case 0x6:
		return PinMode_PinModeI2C, true
	case 0x7:
		return PinMode_PinModeOneWire, true
	case 0x8:
		return PinMode_PinModeStepper, true
	case 0x9:
		return PinMode_PinModeEncoder, true
	case 0xA:
		return PinMode_PinModeSerial, true
	case 0xB:
		return PinMode_PinModePullup, true
	}
	return 0, false
}

func PinModeByName(value string) (enum PinMode, ok bool) {
	switch value {
	case "PinModeInput":
		return PinMode_PinModeInput, true
	case "PinModeOutput":
		return PinMode_PinModeOutput, true
	case "PinModeAnalog":
		return PinMode_PinModeAnalog, true
	case "PinModePwm":
		return PinMode_PinModePwm, true
	case "PinModeServo":
		return PinMode_PinModeServo, true
	case "PinModeShift":
		return PinMode_PinModeShift, true
	case "PinModeI2C":
		return PinMode_PinModeI2C, true
	case "PinModeOneWire":
		return PinMode_PinModeOneWire, true
	case "PinModeStepper":
		return PinMode_PinModeStepper, true
	case "PinModeEncoder":
		return PinMode_PinModeEncoder, true
	case "PinModeSerial":
		return PinMode_PinModeSerial, true
	case "PinModePullup":
		return PinMode_PinModePullup, true
	}
	return 0, false
}

func PinModeKnows(value uint8) bool {
	for _, typeValue := range PinModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPinMode(structType any) PinMode {
	castFunc := func(typ any) PinMode {
		if sPinMode, ok := typ.(PinMode); ok {
			return sPinMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m PinMode) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m PinMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PinModeParse(ctx context.Context, theBytes []byte) (PinMode, error) {
	return PinModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PinModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PinMode, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint8("PinMode", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PinMode")
	}
	if enum, ok := PinModeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PinMode")
		return PinMode(val), nil
	} else {
		return enum, nil
	}
}

func (e PinMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PinMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint8("PinMode", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e PinMode) GetValue() uint8 {
	return uint8(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PinMode) PLC4XEnumName() string {
	switch e {
	case PinMode_PinModeInput:
		return "PinModeInput"
	case PinMode_PinModeOutput:
		return "PinModeOutput"
	case PinMode_PinModeAnalog:
		return "PinModeAnalog"
	case PinMode_PinModePwm:
		return "PinModePwm"
	case PinMode_PinModeServo:
		return "PinModeServo"
	case PinMode_PinModeShift:
		return "PinModeShift"
	case PinMode_PinModeI2C:
		return "PinModeI2C"
	case PinMode_PinModeOneWire:
		return "PinModeOneWire"
	case PinMode_PinModeStepper:
		return "PinModeStepper"
	case PinMode_PinModeEncoder:
		return "PinModeEncoder"
	case PinMode_PinModeSerial:
		return "PinModeSerial"
	case PinMode_PinModePullup:
		return "PinModePullup"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e PinMode) String() string {
	return e.PLC4XEnumName()
}
