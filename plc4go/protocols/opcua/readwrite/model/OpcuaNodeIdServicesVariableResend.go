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

// OpcuaNodeIdServicesVariableResend is an enum
type OpcuaNodeIdServicesVariableResend int32

type IOpcuaNodeIdServicesVariableResend interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableResend_ResendDataMethodType_InputArguments OpcuaNodeIdServicesVariableResend = 12876
)

var OpcuaNodeIdServicesVariableResendValues []OpcuaNodeIdServicesVariableResend

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableResendValues = []OpcuaNodeIdServicesVariableResend{
		OpcuaNodeIdServicesVariableResend_ResendDataMethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableResendByValue(value int32) (enum OpcuaNodeIdServicesVariableResend, ok bool) {
	switch value {
	case 12876:
		return OpcuaNodeIdServicesVariableResend_ResendDataMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableResendByName(value string) (enum OpcuaNodeIdServicesVariableResend, ok bool) {
	switch value {
	case "ResendDataMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableResend_ResendDataMethodType_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableResendKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableResendValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableResend(structType any) OpcuaNodeIdServicesVariableResend {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableResend {
		if sOpcuaNodeIdServicesVariableResend, ok := typ.(OpcuaNodeIdServicesVariableResend); ok {
			return sOpcuaNodeIdServicesVariableResend
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableResend) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableResend) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableResendParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableResend, error) {
	return OpcuaNodeIdServicesVariableResendParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableResendParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableResend, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableResend", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableResend")
	}
	if enum, ok := OpcuaNodeIdServicesVariableResendByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableResend")
		return OpcuaNodeIdServicesVariableResend(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableResend) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableResend) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableResend", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableResend) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableResend) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableResend_ResendDataMethodType_InputArguments:
		return "ResendDataMethodType_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableResend) String() string {
	return e.PLC4XEnumName()
}
