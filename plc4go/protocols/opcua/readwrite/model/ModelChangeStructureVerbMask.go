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

// ModelChangeStructureVerbMask is an enum
type ModelChangeStructureVerbMask uint32

type IModelChangeStructureVerbMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeAdded        ModelChangeStructureVerbMask = 1
	ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeDeleted      ModelChangeStructureVerbMask = 2
	ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceAdded   ModelChangeStructureVerbMask = 4
	ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceDeleted ModelChangeStructureVerbMask = 8
	ModelChangeStructureVerbMask_modelChangeStructureVerbMaskDataTypeChanged  ModelChangeStructureVerbMask = 16
)

var ModelChangeStructureVerbMaskValues []ModelChangeStructureVerbMask

func init() {
	_ = errors.New
	ModelChangeStructureVerbMaskValues = []ModelChangeStructureVerbMask{
		ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeAdded,
		ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeDeleted,
		ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceAdded,
		ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceDeleted,
		ModelChangeStructureVerbMask_modelChangeStructureVerbMaskDataTypeChanged,
	}
}

func ModelChangeStructureVerbMaskByValue(value uint32) (enum ModelChangeStructureVerbMask, ok bool) {
	switch value {
	case 1:
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeAdded, true
	case 16:
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskDataTypeChanged, true
	case 2:
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeDeleted, true
	case 4:
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceAdded, true
	case 8:
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceDeleted, true
	}
	return 0, false
}

func ModelChangeStructureVerbMaskByName(value string) (enum ModelChangeStructureVerbMask, ok bool) {
	switch value {
	case "modelChangeStructureVerbMaskNodeAdded":
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeAdded, true
	case "modelChangeStructureVerbMaskDataTypeChanged":
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskDataTypeChanged, true
	case "modelChangeStructureVerbMaskNodeDeleted":
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeDeleted, true
	case "modelChangeStructureVerbMaskReferenceAdded":
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceAdded, true
	case "modelChangeStructureVerbMaskReferenceDeleted":
		return ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceDeleted, true
	}
	return 0, false
}

func ModelChangeStructureVerbMaskKnows(value uint32) bool {
	for _, typeValue := range ModelChangeStructureVerbMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastModelChangeStructureVerbMask(structType any) ModelChangeStructureVerbMask {
	castFunc := func(typ any) ModelChangeStructureVerbMask {
		if sModelChangeStructureVerbMask, ok := typ.(ModelChangeStructureVerbMask); ok {
			return sModelChangeStructureVerbMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m ModelChangeStructureVerbMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m ModelChangeStructureVerbMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModelChangeStructureVerbMaskParse(ctx context.Context, theBytes []byte) (ModelChangeStructureVerbMask, error) {
	return ModelChangeStructureVerbMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ModelChangeStructureVerbMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModelChangeStructureVerbMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("ModelChangeStructureVerbMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ModelChangeStructureVerbMask")
	}
	if enum, ok := ModelChangeStructureVerbMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for ModelChangeStructureVerbMask")
		return ModelChangeStructureVerbMask(val), nil
	} else {
		return enum, nil
	}
}

func (e ModelChangeStructureVerbMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ModelChangeStructureVerbMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("ModelChangeStructureVerbMask", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ModelChangeStructureVerbMask) PLC4XEnumName() string {
	switch e {
	case ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeAdded:
		return "modelChangeStructureVerbMaskNodeAdded"
	case ModelChangeStructureVerbMask_modelChangeStructureVerbMaskDataTypeChanged:
		return "modelChangeStructureVerbMaskDataTypeChanged"
	case ModelChangeStructureVerbMask_modelChangeStructureVerbMaskNodeDeleted:
		return "modelChangeStructureVerbMaskNodeDeleted"
	case ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceAdded:
		return "modelChangeStructureVerbMaskReferenceAdded"
	case ModelChangeStructureVerbMask_modelChangeStructureVerbMaskReferenceDeleted:
		return "modelChangeStructureVerbMaskReferenceDeleted"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e ModelChangeStructureVerbMask) String() string {
	return e.PLC4XEnumName()
}
