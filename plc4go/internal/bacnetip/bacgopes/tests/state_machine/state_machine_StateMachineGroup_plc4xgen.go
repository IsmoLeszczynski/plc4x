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

// Code generated by "plc4xGenerator -type=StateMachineGroup -prefix=state_machine_"; DO NOT EDIT.

package state_machine

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *StateMachineGroup) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *StateMachineGroup) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("StateMachineGroup"); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("stateMachines", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.stateMachines {
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString("stateMachines", uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("stateMachines", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("isSuccessState", d.isSuccessState); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("isFailState", d.isFailState); err != nil {
		return err
	}
	if err := writeBuffer.PushContext("events", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _name, elem := range d.events {
		name := _name
		_value := fmt.Sprintf("%v", elem)

		if err := writeBuffer.WriteString(name, uint32(len(_value)*8), _value); err != nil {
			return err
		}
	}
	if err := writeBuffer.PopContext("events", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("startupFlag", d.startupFlag); err != nil {
		return err
	}

	if err := writeBuffer.WriteBit("isRunning", d.isRunning); err != nil {
		return err
	}
	if err := writeBuffer.PopContext("StateMachineGroup"); err != nil {
		return err
	}
	return nil
}

func (d *StateMachineGroup) String() string {
	if alternateStringer, ok := any(d).(utils.AlternateStringer); ok {
		if alternateString, use := alternateStringer.AlternateString(); use {
			return alternateString
		}
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), d); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
