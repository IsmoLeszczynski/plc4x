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

// BrowseResultMask is an enum
type BrowseResultMask uint32

type IBrowseResultMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BrowseResultMask_browseResultMaskNone              BrowseResultMask = 0
	BrowseResultMask_browseResultMaskReferenceTypeId   BrowseResultMask = 1
	BrowseResultMask_browseResultMaskIsForward         BrowseResultMask = 2
	BrowseResultMask_browseResultMaskNodeClass         BrowseResultMask = 4
	BrowseResultMask_browseResultMaskBrowseName        BrowseResultMask = 8
	BrowseResultMask_browseResultMaskDisplayName       BrowseResultMask = 16
	BrowseResultMask_browseResultMaskTypeDefinition    BrowseResultMask = 32
	BrowseResultMask_browseResultMaskAll               BrowseResultMask = 63
	BrowseResultMask_browseResultMaskReferenceTypeInfo BrowseResultMask = 3
	BrowseResultMask_browseResultMaskTargetInfo        BrowseResultMask = 60
)

var BrowseResultMaskValues []BrowseResultMask

func init() {
	_ = errors.New
	BrowseResultMaskValues = []BrowseResultMask{
		BrowseResultMask_browseResultMaskNone,
		BrowseResultMask_browseResultMaskReferenceTypeId,
		BrowseResultMask_browseResultMaskIsForward,
		BrowseResultMask_browseResultMaskNodeClass,
		BrowseResultMask_browseResultMaskBrowseName,
		BrowseResultMask_browseResultMaskDisplayName,
		BrowseResultMask_browseResultMaskTypeDefinition,
		BrowseResultMask_browseResultMaskAll,
		BrowseResultMask_browseResultMaskReferenceTypeInfo,
		BrowseResultMask_browseResultMaskTargetInfo,
	}
}

func BrowseResultMaskByValue(value uint32) (enum BrowseResultMask, ok bool) {
	switch value {
	case 0:
		return BrowseResultMask_browseResultMaskNone, true
	case 1:
		return BrowseResultMask_browseResultMaskReferenceTypeId, true
	case 16:
		return BrowseResultMask_browseResultMaskDisplayName, true
	case 2:
		return BrowseResultMask_browseResultMaskIsForward, true
	case 3:
		return BrowseResultMask_browseResultMaskReferenceTypeInfo, true
	case 32:
		return BrowseResultMask_browseResultMaskTypeDefinition, true
	case 4:
		return BrowseResultMask_browseResultMaskNodeClass, true
	case 60:
		return BrowseResultMask_browseResultMaskTargetInfo, true
	case 63:
		return BrowseResultMask_browseResultMaskAll, true
	case 8:
		return BrowseResultMask_browseResultMaskBrowseName, true
	}
	return 0, false
}

func BrowseResultMaskByName(value string) (enum BrowseResultMask, ok bool) {
	switch value {
	case "browseResultMaskNone":
		return BrowseResultMask_browseResultMaskNone, true
	case "browseResultMaskReferenceTypeId":
		return BrowseResultMask_browseResultMaskReferenceTypeId, true
	case "browseResultMaskDisplayName":
		return BrowseResultMask_browseResultMaskDisplayName, true
	case "browseResultMaskIsForward":
		return BrowseResultMask_browseResultMaskIsForward, true
	case "browseResultMaskReferenceTypeInfo":
		return BrowseResultMask_browseResultMaskReferenceTypeInfo, true
	case "browseResultMaskTypeDefinition":
		return BrowseResultMask_browseResultMaskTypeDefinition, true
	case "browseResultMaskNodeClass":
		return BrowseResultMask_browseResultMaskNodeClass, true
	case "browseResultMaskTargetInfo":
		return BrowseResultMask_browseResultMaskTargetInfo, true
	case "browseResultMaskAll":
		return BrowseResultMask_browseResultMaskAll, true
	case "browseResultMaskBrowseName":
		return BrowseResultMask_browseResultMaskBrowseName, true
	}
	return 0, false
}

func BrowseResultMaskKnows(value uint32) bool {
	for _, typeValue := range BrowseResultMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBrowseResultMask(structType any) BrowseResultMask {
	castFunc := func(typ any) BrowseResultMask {
		if sBrowseResultMask, ok := typ.(BrowseResultMask); ok {
			return sBrowseResultMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m BrowseResultMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m BrowseResultMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BrowseResultMaskParse(ctx context.Context, theBytes []byte) (BrowseResultMask, error) {
	return BrowseResultMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BrowseResultMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BrowseResultMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("BrowseResultMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BrowseResultMask")
	}
	if enum, ok := BrowseResultMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for BrowseResultMask")
		return BrowseResultMask(val), nil
	} else {
		return enum, nil
	}
}

func (e BrowseResultMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BrowseResultMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("BrowseResultMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BrowseResultMask) PLC4XEnumName() string {
	switch e {
	case BrowseResultMask_browseResultMaskNone:
		return "browseResultMaskNone"
	case BrowseResultMask_browseResultMaskReferenceTypeId:
		return "browseResultMaskReferenceTypeId"
	case BrowseResultMask_browseResultMaskDisplayName:
		return "browseResultMaskDisplayName"
	case BrowseResultMask_browseResultMaskIsForward:
		return "browseResultMaskIsForward"
	case BrowseResultMask_browseResultMaskReferenceTypeInfo:
		return "browseResultMaskReferenceTypeInfo"
	case BrowseResultMask_browseResultMaskTypeDefinition:
		return "browseResultMaskTypeDefinition"
	case BrowseResultMask_browseResultMaskNodeClass:
		return "browseResultMaskNodeClass"
	case BrowseResultMask_browseResultMaskTargetInfo:
		return "browseResultMaskTargetInfo"
	case BrowseResultMask_browseResultMaskAll:
		return "browseResultMaskAll"
	case BrowseResultMask_browseResultMaskBrowseName:
		return "browseResultMaskBrowseName"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e BrowseResultMask) String() string {
	return e.PLC4XEnumName()
}
