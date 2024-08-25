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

// HistoryUpdateType is an enum
type HistoryUpdateType uint32

type IHistoryUpdateType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	HistoryUpdateType_historyUpdateTypeInsert  HistoryUpdateType = 1
	HistoryUpdateType_historyUpdateTypeReplace HistoryUpdateType = 2
	HistoryUpdateType_historyUpdateTypeUpdate  HistoryUpdateType = 3
	HistoryUpdateType_historyUpdateTypeDelete  HistoryUpdateType = 4
)

var HistoryUpdateTypeValues []HistoryUpdateType

func init() {
	_ = errors.New
	HistoryUpdateTypeValues = []HistoryUpdateType{
		HistoryUpdateType_historyUpdateTypeInsert,
		HistoryUpdateType_historyUpdateTypeReplace,
		HistoryUpdateType_historyUpdateTypeUpdate,
		HistoryUpdateType_historyUpdateTypeDelete,
	}
}

func HistoryUpdateTypeByValue(value uint32) (enum HistoryUpdateType, ok bool) {
	switch value {
	case 1:
		return HistoryUpdateType_historyUpdateTypeInsert, true
	case 2:
		return HistoryUpdateType_historyUpdateTypeReplace, true
	case 3:
		return HistoryUpdateType_historyUpdateTypeUpdate, true
	case 4:
		return HistoryUpdateType_historyUpdateTypeDelete, true
	}
	return 0, false
}

func HistoryUpdateTypeByName(value string) (enum HistoryUpdateType, ok bool) {
	switch value {
	case "historyUpdateTypeInsert":
		return HistoryUpdateType_historyUpdateTypeInsert, true
	case "historyUpdateTypeReplace":
		return HistoryUpdateType_historyUpdateTypeReplace, true
	case "historyUpdateTypeUpdate":
		return HistoryUpdateType_historyUpdateTypeUpdate, true
	case "historyUpdateTypeDelete":
		return HistoryUpdateType_historyUpdateTypeDelete, true
	}
	return 0, false
}

func HistoryUpdateTypeKnows(value uint32) bool {
	for _, typeValue := range HistoryUpdateTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastHistoryUpdateType(structType any) HistoryUpdateType {
	castFunc := func(typ any) HistoryUpdateType {
		if sHistoryUpdateType, ok := typ.(HistoryUpdateType); ok {
			return sHistoryUpdateType
		}
		return 0
	}
	return castFunc(structType)
}

func (m HistoryUpdateType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m HistoryUpdateType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryUpdateTypeParse(ctx context.Context, theBytes []byte) (HistoryUpdateType, error) {
	return HistoryUpdateTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func HistoryUpdateTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (HistoryUpdateType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("HistoryUpdateType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading HistoryUpdateType")
	}
	if enum, ok := HistoryUpdateTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for HistoryUpdateType")
		return HistoryUpdateType(val), nil
	} else {
		return enum, nil
	}
}

func (e HistoryUpdateType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e HistoryUpdateType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("HistoryUpdateType", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e HistoryUpdateType) PLC4XEnumName() string {
	switch e {
	case HistoryUpdateType_historyUpdateTypeInsert:
		return "historyUpdateTypeInsert"
	case HistoryUpdateType_historyUpdateTypeReplace:
		return "historyUpdateTypeReplace"
	case HistoryUpdateType_historyUpdateTypeUpdate:
		return "historyUpdateTypeUpdate"
	case HistoryUpdateType_historyUpdateTypeDelete:
		return "historyUpdateTypeDelete"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e HistoryUpdateType) String() string {
	return e.PLC4XEnumName()
}
