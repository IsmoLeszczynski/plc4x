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

// AxisScaleEnumeration is an enum
type AxisScaleEnumeration uint32

type IAxisScaleEnumeration interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AxisScaleEnumeration_axisScaleEnumerationLinear AxisScaleEnumeration = 0
	AxisScaleEnumeration_axisScaleEnumerationLog    AxisScaleEnumeration = 1
	AxisScaleEnumeration_axisScaleEnumerationLn     AxisScaleEnumeration = 2
)

var AxisScaleEnumerationValues []AxisScaleEnumeration

func init() {
	_ = errors.New
	AxisScaleEnumerationValues = []AxisScaleEnumeration{
		AxisScaleEnumeration_axisScaleEnumerationLinear,
		AxisScaleEnumeration_axisScaleEnumerationLog,
		AxisScaleEnumeration_axisScaleEnumerationLn,
	}
}

func AxisScaleEnumerationByValue(value uint32) (enum AxisScaleEnumeration, ok bool) {
	switch value {
	case 0:
		return AxisScaleEnumeration_axisScaleEnumerationLinear, true
	case 1:
		return AxisScaleEnumeration_axisScaleEnumerationLog, true
	case 2:
		return AxisScaleEnumeration_axisScaleEnumerationLn, true
	}
	return 0, false
}

func AxisScaleEnumerationByName(value string) (enum AxisScaleEnumeration, ok bool) {
	switch value {
	case "axisScaleEnumerationLinear":
		return AxisScaleEnumeration_axisScaleEnumerationLinear, true
	case "axisScaleEnumerationLog":
		return AxisScaleEnumeration_axisScaleEnumerationLog, true
	case "axisScaleEnumerationLn":
		return AxisScaleEnumeration_axisScaleEnumerationLn, true
	}
	return 0, false
}

func AxisScaleEnumerationKnows(value uint32) bool {
	for _, typeValue := range AxisScaleEnumerationValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAxisScaleEnumeration(structType any) AxisScaleEnumeration {
	castFunc := func(typ any) AxisScaleEnumeration {
		if sAxisScaleEnumeration, ok := typ.(AxisScaleEnumeration); ok {
			return sAxisScaleEnumeration
		}
		return 0
	}
	return castFunc(structType)
}

func (m AxisScaleEnumeration) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m AxisScaleEnumeration) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AxisScaleEnumerationParse(ctx context.Context, theBytes []byte) (AxisScaleEnumeration, error) {
	return AxisScaleEnumerationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AxisScaleEnumerationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AxisScaleEnumeration, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("AxisScaleEnumeration", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AxisScaleEnumeration")
	}
	if enum, ok := AxisScaleEnumerationByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AxisScaleEnumeration")
		return AxisScaleEnumeration(val), nil
	} else {
		return enum, nil
	}
}

func (e AxisScaleEnumeration) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AxisScaleEnumeration) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("AxisScaleEnumeration", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AxisScaleEnumeration) PLC4XEnumName() string {
	switch e {
	case AxisScaleEnumeration_axisScaleEnumerationLinear:
		return "axisScaleEnumerationLinear"
	case AxisScaleEnumeration_axisScaleEnumerationLog:
		return "axisScaleEnumerationLog"
	case AxisScaleEnumeration_axisScaleEnumerationLn:
		return "axisScaleEnumerationLn"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e AxisScaleEnumeration) String() string {
	return e.PLC4XEnumName()
}
