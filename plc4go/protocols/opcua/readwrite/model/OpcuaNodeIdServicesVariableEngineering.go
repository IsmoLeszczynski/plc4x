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

// OpcuaNodeIdServicesVariableEngineering is an enum
type OpcuaNodeIdServicesVariableEngineering int32

type IOpcuaNodeIdServicesVariableEngineering interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableEngineering_EngineeringUnits OpcuaNodeIdServicesVariableEngineering = 11513
)

var OpcuaNodeIdServicesVariableEngineeringValues []OpcuaNodeIdServicesVariableEngineering

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableEngineeringValues = []OpcuaNodeIdServicesVariableEngineering{
		OpcuaNodeIdServicesVariableEngineering_EngineeringUnits,
	}
}

func OpcuaNodeIdServicesVariableEngineeringByValue(value int32) (enum OpcuaNodeIdServicesVariableEngineering, ok bool) {
	switch value {
	case 11513:
		return OpcuaNodeIdServicesVariableEngineering_EngineeringUnits, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableEngineeringByName(value string) (enum OpcuaNodeIdServicesVariableEngineering, ok bool) {
	switch value {
	case "EngineeringUnits":
		return OpcuaNodeIdServicesVariableEngineering_EngineeringUnits, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableEngineeringKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableEngineeringValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableEngineering(structType any) OpcuaNodeIdServicesVariableEngineering {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableEngineering {
		if sOpcuaNodeIdServicesVariableEngineering, ok := typ.(OpcuaNodeIdServicesVariableEngineering); ok {
			return sOpcuaNodeIdServicesVariableEngineering
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableEngineering) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableEngineering) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableEngineeringParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableEngineering, error) {
	return OpcuaNodeIdServicesVariableEngineeringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableEngineeringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableEngineering, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableEngineering", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableEngineering")
	}
	if enum, ok := OpcuaNodeIdServicesVariableEngineeringByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableEngineering")
		return OpcuaNodeIdServicesVariableEngineering(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableEngineering) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableEngineering) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableEngineering", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableEngineering) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableEngineering) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableEngineering_EngineeringUnits:
		return "EngineeringUnits"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableEngineering) String() string {
	return e.PLC4XEnumName()
}
