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

// OpcuaNodeIdServicesVariableExpression is an enum
type OpcuaNodeIdServicesVariableExpression int32

type IOpcuaNodeIdServicesVariableExpression interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableExpression_ExpressionGuardVariableType_Expression OpcuaNodeIdServicesVariableExpression = 15129
)

var OpcuaNodeIdServicesVariableExpressionValues []OpcuaNodeIdServicesVariableExpression

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableExpressionValues = []OpcuaNodeIdServicesVariableExpression{
		OpcuaNodeIdServicesVariableExpression_ExpressionGuardVariableType_Expression,
	}
}

func OpcuaNodeIdServicesVariableExpressionByValue(value int32) (enum OpcuaNodeIdServicesVariableExpression, ok bool) {
	switch value {
	case 15129:
		return OpcuaNodeIdServicesVariableExpression_ExpressionGuardVariableType_Expression, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExpressionByName(value string) (enum OpcuaNodeIdServicesVariableExpression, ok bool) {
	switch value {
	case "ExpressionGuardVariableType_Expression":
		return OpcuaNodeIdServicesVariableExpression_ExpressionGuardVariableType_Expression, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableExpressionKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableExpressionValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableExpression(structType any) OpcuaNodeIdServicesVariableExpression {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableExpression {
		if sOpcuaNodeIdServicesVariableExpression, ok := typ.(OpcuaNodeIdServicesVariableExpression); ok {
			return sOpcuaNodeIdServicesVariableExpression
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableExpression) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableExpression) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableExpressionParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableExpression, error) {
	return OpcuaNodeIdServicesVariableExpressionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableExpressionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableExpression, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableExpression", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableExpression")
	}
	if enum, ok := OpcuaNodeIdServicesVariableExpressionByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableExpression")
		return OpcuaNodeIdServicesVariableExpression(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableExpression) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableExpression) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableExpression", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableExpression) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableExpression_ExpressionGuardVariableType_Expression:
		return "ExpressionGuardVariableType_Expression"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableExpression) String() string {
	return e.PLC4XEnumName()
}
