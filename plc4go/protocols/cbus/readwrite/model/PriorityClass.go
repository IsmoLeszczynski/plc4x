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

// PriorityClass is an enum
type PriorityClass uint8

type IPriorityClass interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	PriorityClass_Class4 PriorityClass = 0x00
	PriorityClass_Class3 PriorityClass = 0x01
	PriorityClass_Class2 PriorityClass = 0x02
	PriorityClass_Class1 PriorityClass = 0x03
)

var PriorityClassValues []PriorityClass

func init() {
	_ = errors.New
	PriorityClassValues = []PriorityClass{
		PriorityClass_Class4,
		PriorityClass_Class3,
		PriorityClass_Class2,
		PriorityClass_Class1,
	}
}

func PriorityClassByValue(value uint8) (enum PriorityClass, ok bool) {
	switch value {
	case 0x00:
		return PriorityClass_Class4, true
	case 0x01:
		return PriorityClass_Class3, true
	case 0x02:
		return PriorityClass_Class2, true
	case 0x03:
		return PriorityClass_Class1, true
	}
	return 0, false
}

func PriorityClassByName(value string) (enum PriorityClass, ok bool) {
	switch value {
	case "Class4":
		return PriorityClass_Class4, true
	case "Class3":
		return PriorityClass_Class3, true
	case "Class2":
		return PriorityClass_Class2, true
	case "Class1":
		return PriorityClass_Class1, true
	}
	return 0, false
}

func PriorityClassKnows(value uint8) bool {
	for _, typeValue := range PriorityClassValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPriorityClass(structType any) PriorityClass {
	castFunc := func(typ any) PriorityClass {
		if sPriorityClass, ok := typ.(PriorityClass); ok {
			return sPriorityClass
		}
		return 0
	}
	return castFunc(structType)
}

func (m PriorityClass) GetLengthInBits(ctx context.Context) uint16 {
	return 2
}

func (m PriorityClass) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PriorityClassParse(ctx context.Context, theBytes []byte) (PriorityClass, error) {
	return PriorityClassParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func PriorityClassParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (PriorityClass, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("PriorityClass", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading PriorityClass")
	}
	if enum, ok := PriorityClassByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for PriorityClass")
		return PriorityClass(val), nil
	} else {
		return enum, nil
	}
}

func (e PriorityClass) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e PriorityClass) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("PriorityClass", 2, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PriorityClass) PLC4XEnumName() string {
	switch e {
	case PriorityClass_Class4:
		return "Class4"
	case PriorityClass_Class3:
		return "Class3"
	case PriorityClass_Class2:
		return "Class2"
	case PriorityClass_Class1:
		return "Class1"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e PriorityClass) String() string {
	return e.PLC4XEnumName()
}
