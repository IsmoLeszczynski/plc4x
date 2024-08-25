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

// IdType is an enum
type IdType uint32

type IIdType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	IdType_idTypeNumeric IdType = 0
	IdType_idTypeString  IdType = 1
	IdType_idTypeGuid    IdType = 2
	IdType_idTypeOpaque  IdType = 3
)

var IdTypeValues []IdType

func init() {
	_ = errors.New
	IdTypeValues = []IdType{
		IdType_idTypeNumeric,
		IdType_idTypeString,
		IdType_idTypeGuid,
		IdType_idTypeOpaque,
	}
}

func IdTypeByValue(value uint32) (enum IdType, ok bool) {
	switch value {
	case 0:
		return IdType_idTypeNumeric, true
	case 1:
		return IdType_idTypeString, true
	case 2:
		return IdType_idTypeGuid, true
	case 3:
		return IdType_idTypeOpaque, true
	}
	return 0, false
}

func IdTypeByName(value string) (enum IdType, ok bool) {
	switch value {
	case "idTypeNumeric":
		return IdType_idTypeNumeric, true
	case "idTypeString":
		return IdType_idTypeString, true
	case "idTypeGuid":
		return IdType_idTypeGuid, true
	case "idTypeOpaque":
		return IdType_idTypeOpaque, true
	}
	return 0, false
}

func IdTypeKnows(value uint32) bool {
	for _, typeValue := range IdTypeValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastIdType(structType any) IdType {
	castFunc := func(typ any) IdType {
		if sIdType, ok := typ.(IdType); ok {
			return sIdType
		}
		return 0
	}
	return castFunc(structType)
}

func (m IdType) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m IdType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdTypeParse(ctx context.Context, theBytes []byte) (IdType, error) {
	return IdTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IdTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IdType, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("IdType", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading IdType")
	}
	if enum, ok := IdTypeByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for IdType")
		return IdType(val), nil
	} else {
		return enum, nil
	}
}

func (e IdType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e IdType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("IdType", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e IdType) PLC4XEnumName() string {
	switch e {
	case IdType_idTypeNumeric:
		return "idTypeNumeric"
	case IdType_idTypeString:
		return "idTypeString"
	case IdType_idTypeGuid:
		return "idTypeGuid"
	case IdType_idTypeOpaque:
		return "idTypeOpaque"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e IdType) String() string {
	return e.PLC4XEnumName()
}
