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

// TsnTalkerStatus is an enum
type TsnTalkerStatus uint32

type ITsnTalkerStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TsnTalkerStatus_tsnTalkerStatusNone   TsnTalkerStatus = 0
	TsnTalkerStatus_tsnTalkerStatusReady  TsnTalkerStatus = 1
	TsnTalkerStatus_tsnTalkerStatusFailed TsnTalkerStatus = 2
)

var TsnTalkerStatusValues []TsnTalkerStatus

func init() {
	_ = errors.New
	TsnTalkerStatusValues = []TsnTalkerStatus{
		TsnTalkerStatus_tsnTalkerStatusNone,
		TsnTalkerStatus_tsnTalkerStatusReady,
		TsnTalkerStatus_tsnTalkerStatusFailed,
	}
}

func TsnTalkerStatusByValue(value uint32) (enum TsnTalkerStatus, ok bool) {
	switch value {
	case 0:
		return TsnTalkerStatus_tsnTalkerStatusNone, true
	case 1:
		return TsnTalkerStatus_tsnTalkerStatusReady, true
	case 2:
		return TsnTalkerStatus_tsnTalkerStatusFailed, true
	}
	return 0, false
}

func TsnTalkerStatusByName(value string) (enum TsnTalkerStatus, ok bool) {
	switch value {
	case "tsnTalkerStatusNone":
		return TsnTalkerStatus_tsnTalkerStatusNone, true
	case "tsnTalkerStatusReady":
		return TsnTalkerStatus_tsnTalkerStatusReady, true
	case "tsnTalkerStatusFailed":
		return TsnTalkerStatus_tsnTalkerStatusFailed, true
	}
	return 0, false
}

func TsnTalkerStatusKnows(value uint32) bool {
	for _, typeValue := range TsnTalkerStatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTsnTalkerStatus(structType any) TsnTalkerStatus {
	castFunc := func(typ any) TsnTalkerStatus {
		if sTsnTalkerStatus, ok := typ.(TsnTalkerStatus); ok {
			return sTsnTalkerStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m TsnTalkerStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TsnTalkerStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TsnTalkerStatusParse(ctx context.Context, theBytes []byte) (TsnTalkerStatus, error) {
	return TsnTalkerStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TsnTalkerStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TsnTalkerStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("TsnTalkerStatus", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TsnTalkerStatus")
	}
	if enum, ok := TsnTalkerStatusByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TsnTalkerStatus")
		return TsnTalkerStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e TsnTalkerStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TsnTalkerStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("TsnTalkerStatus", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TsnTalkerStatus) PLC4XEnumName() string {
	switch e {
	case TsnTalkerStatus_tsnTalkerStatusNone:
		return "tsnTalkerStatusNone"
	case TsnTalkerStatus_tsnTalkerStatusReady:
		return "tsnTalkerStatusReady"
	case TsnTalkerStatus_tsnTalkerStatusFailed:
		return "tsnTalkerStatusFailed"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e TsnTalkerStatus) String() string {
	return e.PLC4XEnumName()
}
