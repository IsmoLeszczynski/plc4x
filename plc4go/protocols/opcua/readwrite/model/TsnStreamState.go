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

// TsnStreamState is an enum
type TsnStreamState uint32

type ITsnStreamState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	TsnStreamState_tsnStreamStateDisabled    TsnStreamState = 0
	TsnStreamState_tsnStreamStateConfiguring TsnStreamState = 1
	TsnStreamState_tsnStreamStateReady       TsnStreamState = 2
	TsnStreamState_tsnStreamStateOperational TsnStreamState = 3
	TsnStreamState_tsnStreamStateError       TsnStreamState = 4
)

var TsnStreamStateValues []TsnStreamState

func init() {
	_ = errors.New
	TsnStreamStateValues = []TsnStreamState{
		TsnStreamState_tsnStreamStateDisabled,
		TsnStreamState_tsnStreamStateConfiguring,
		TsnStreamState_tsnStreamStateReady,
		TsnStreamState_tsnStreamStateOperational,
		TsnStreamState_tsnStreamStateError,
	}
}

func TsnStreamStateByValue(value uint32) (enum TsnStreamState, ok bool) {
	switch value {
	case 0:
		return TsnStreamState_tsnStreamStateDisabled, true
	case 1:
		return TsnStreamState_tsnStreamStateConfiguring, true
	case 2:
		return TsnStreamState_tsnStreamStateReady, true
	case 3:
		return TsnStreamState_tsnStreamStateOperational, true
	case 4:
		return TsnStreamState_tsnStreamStateError, true
	}
	return 0, false
}

func TsnStreamStateByName(value string) (enum TsnStreamState, ok bool) {
	switch value {
	case "tsnStreamStateDisabled":
		return TsnStreamState_tsnStreamStateDisabled, true
	case "tsnStreamStateConfiguring":
		return TsnStreamState_tsnStreamStateConfiguring, true
	case "tsnStreamStateReady":
		return TsnStreamState_tsnStreamStateReady, true
	case "tsnStreamStateOperational":
		return TsnStreamState_tsnStreamStateOperational, true
	case "tsnStreamStateError":
		return TsnStreamState_tsnStreamStateError, true
	}
	return 0, false
}

func TsnStreamStateKnows(value uint32) bool {
	for _, typeValue := range TsnStreamStateValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastTsnStreamState(structType any) TsnStreamState {
	castFunc := func(typ any) TsnStreamState {
		if sTsnStreamState, ok := typ.(TsnStreamState); ok {
			return sTsnStreamState
		}
		return 0
	}
	return castFunc(structType)
}

func (m TsnStreamState) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m TsnStreamState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TsnStreamStateParse(ctx context.Context, theBytes []byte) (TsnStreamState, error) {
	return TsnStreamStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TsnStreamStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TsnStreamState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("TsnStreamState", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading TsnStreamState")
	}
	if enum, ok := TsnStreamStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for TsnStreamState")
		return TsnStreamState(val), nil
	} else {
		return enum, nil
	}
}

func (e TsnStreamState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e TsnStreamState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteUint32("TsnStreamState", 32, uint32(uint32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e TsnStreamState) GetValue() uint32 {
	return uint32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e TsnStreamState) PLC4XEnumName() string {
	switch e {
	case TsnStreamState_tsnStreamStateDisabled:
		return "tsnStreamStateDisabled"
	case TsnStreamState_tsnStreamStateConfiguring:
		return "tsnStreamStateConfiguring"
	case TsnStreamState_tsnStreamStateReady:
		return "tsnStreamStateReady"
	case TsnStreamState_tsnStreamStateOperational:
		return "tsnStreamStateOperational"
	case TsnStreamState_tsnStreamStateError:
		return "tsnStreamStateError"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e TsnStreamState) String() string {
	return e.PLC4XEnumName()
}
