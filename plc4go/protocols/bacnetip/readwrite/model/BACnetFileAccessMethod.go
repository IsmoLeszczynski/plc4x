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

// BACnetFileAccessMethod is an enum
type BACnetFileAccessMethod uint8

type IBACnetFileAccessMethod interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetFileAccessMethod_RECORD_ACCESS BACnetFileAccessMethod = 0
	BACnetFileAccessMethod_STREAM_ACCESS BACnetFileAccessMethod = 1
)

var BACnetFileAccessMethodValues []BACnetFileAccessMethod

func init() {
	_ = errors.New
	BACnetFileAccessMethodValues = []BACnetFileAccessMethod{
		BACnetFileAccessMethod_RECORD_ACCESS,
		BACnetFileAccessMethod_STREAM_ACCESS,
	}
}

func BACnetFileAccessMethodByValue(value uint8) (enum BACnetFileAccessMethod, ok bool) {
	switch value {
	case 0:
		return BACnetFileAccessMethod_RECORD_ACCESS, true
	case 1:
		return BACnetFileAccessMethod_STREAM_ACCESS, true
	}
	return 0, false
}

func BACnetFileAccessMethodByName(value string) (enum BACnetFileAccessMethod, ok bool) {
	switch value {
	case "RECORD_ACCESS":
		return BACnetFileAccessMethod_RECORD_ACCESS, true
	case "STREAM_ACCESS":
		return BACnetFileAccessMethod_STREAM_ACCESS, true
	}
	return 0, false
}

func BACnetFileAccessMethodKnows(value uint8) bool {
	for _, typeValue := range BACnetFileAccessMethodValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetFileAccessMethod(structType any) BACnetFileAccessMethod {
	castFunc := func(typ any) BACnetFileAccessMethod {
		if sBACnetFileAccessMethod, ok := typ.(BACnetFileAccessMethod); ok {
			return sBACnetFileAccessMethod
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetFileAccessMethod) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetFileAccessMethod) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFileAccessMethodParse(ctx context.Context, theBytes []byte) (BACnetFileAccessMethod, error) {
	return BACnetFileAccessMethodParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFileAccessMethodParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFileAccessMethod, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetFileAccessMethod", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetFileAccessMethod")
	}
	if enum, ok := BACnetFileAccessMethodByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BACnetFileAccessMethod")
		return BACnetFileAccessMethod(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetFileAccessMethod) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetFileAccessMethod) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetFileAccessMethod", 8, uint8(uint8(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetFileAccessMethod) PLC4XEnumName() string {
	switch e {
	case BACnetFileAccessMethod_RECORD_ACCESS:
		return "RECORD_ACCESS"
	case BACnetFileAccessMethod_STREAM_ACCESS:
		return "STREAM_ACCESS"
	}
	return fmt.Sprintf("Unknown(%v)", uint8(e))
}

func (e BACnetFileAccessMethod) String() string {
	return e.PLC4XEnumName()
}
