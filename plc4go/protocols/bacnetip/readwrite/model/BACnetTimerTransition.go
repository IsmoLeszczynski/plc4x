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

// BACnetTimerTransition is an enum
type BACnetTimerTransition uint8

type IBACnetTimerTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetTimerTransition_NONE               BACnetTimerTransition = 0
	BACnetTimerTransition_IDLE_TO_RUNNING    BACnetTimerTransition = 1
	BACnetTimerTransition_RUNNING_TO_IDLE    BACnetTimerTransition = 2
	BACnetTimerTransition_RUNNING_TO_RUNNING BACnetTimerTransition = 3
	BACnetTimerTransition_RUNNING_TO_EXPIRED BACnetTimerTransition = 4
	BACnetTimerTransition_FORCED_TO_EXPIRED  BACnetTimerTransition = 5
	BACnetTimerTransition_EXPIRED_TO_IDLE    BACnetTimerTransition = 6
	BACnetTimerTransition_EXPIRED_TO_RUNNING BACnetTimerTransition = 7
)

var BACnetTimerTransitionValues []BACnetTimerTransition

func init() {
	_ = errors.New
	BACnetTimerTransitionValues = []BACnetTimerTransition{
		BACnetTimerTransition_NONE,
		BACnetTimerTransition_IDLE_TO_RUNNING,
		BACnetTimerTransition_RUNNING_TO_IDLE,
		BACnetTimerTransition_RUNNING_TO_RUNNING,
		BACnetTimerTransition_RUNNING_TO_EXPIRED,
		BACnetTimerTransition_FORCED_TO_EXPIRED,
		BACnetTimerTransition_EXPIRED_TO_IDLE,
		BACnetTimerTransition_EXPIRED_TO_RUNNING,
	}
}

func BACnetTimerTransitionByValue(value uint8) (enum BACnetTimerTransition, ok bool) {
	switch value {
	case 0:
		return BACnetTimerTransition_NONE, true
	case 1:
		return BACnetTimerTransition_IDLE_TO_RUNNING, true
	case 2:
		return BACnetTimerTransition_RUNNING_TO_IDLE, true
	case 3:
		return BACnetTimerTransition_RUNNING_TO_RUNNING, true
	case 4:
		return BACnetTimerTransition_RUNNING_TO_EXPIRED, true
	case 5:
		return BACnetTimerTransition_FORCED_TO_EXPIRED, true
	case 6:
		return BACnetTimerTransition_EXPIRED_TO_IDLE, true
	case 7:
		return BACnetTimerTransition_EXPIRED_TO_RUNNING, true
	}
	return 0, false
}

func BACnetTimerTransitionByName(value string) (enum BACnetTimerTransition, ok bool) {
	switch value {
	case "NONE":
		return BACnetTimerTransition_NONE, true
	case "IDLE_TO_RUNNING":
		return BACnetTimerTransition_IDLE_TO_RUNNING, true
	case "RUNNING_TO_IDLE":
		return BACnetTimerTransition_RUNNING_TO_IDLE, true
	case "RUNNING_TO_RUNNING":
		return BACnetTimerTransition_RUNNING_TO_RUNNING, true
	case "RUNNING_TO_EXPIRED":
		return BACnetTimerTransition_RUNNING_TO_EXPIRED, true
	case "FORCED_TO_EXPIRED":
		return BACnetTimerTransition_FORCED_TO_EXPIRED, true
	case "EXPIRED_TO_IDLE":
		return BACnetTimerTransition_EXPIRED_TO_IDLE, true
	case "EXPIRED_TO_RUNNING":
		return BACnetTimerTransition_EXPIRED_TO_RUNNING, true
	}
	return 0, false
}

func BACnetTimerTransitionKnows(value uint8) bool {
	for _, typeValue := range BACnetTimerTransitionValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetTimerTransition(structType any) BACnetTimerTransition {
	castFunc := func(typ any) BACnetTimerTransition {
		if sBACnetTimerTransition, ok := typ.(BACnetTimerTransition); ok {
			return sBACnetTimerTransition
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetTimerTransition) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetTimerTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerTransitionParse(ctx context.Context, theBytes []byte) (BACnetTimerTransition, error) {
	return BACnetTimerTransitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimerTransitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetTimerTransition, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetTimerTransition", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetTimerTransition")
	}
	if enum, ok := BACnetTimerTransitionByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetTimerTransition(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetTimerTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetTimerTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetTimerTransition", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetTimerTransition) PLC4XEnumName() string {
	switch e {
	case BACnetTimerTransition_NONE:
		return "NONE"
	case BACnetTimerTransition_IDLE_TO_RUNNING:
		return "IDLE_TO_RUNNING"
	case BACnetTimerTransition_RUNNING_TO_IDLE:
		return "RUNNING_TO_IDLE"
	case BACnetTimerTransition_RUNNING_TO_RUNNING:
		return "RUNNING_TO_RUNNING"
	case BACnetTimerTransition_RUNNING_TO_EXPIRED:
		return "RUNNING_TO_EXPIRED"
	case BACnetTimerTransition_FORCED_TO_EXPIRED:
		return "FORCED_TO_EXPIRED"
	case BACnetTimerTransition_EXPIRED_TO_IDLE:
		return "EXPIRED_TO_IDLE"
	case BACnetTimerTransition_EXPIRED_TO_RUNNING:
		return "EXPIRED_TO_RUNNING"
	}
	return ""
}

func (e BACnetTimerTransition) String() string {
	return e.PLC4XEnumName()
}
