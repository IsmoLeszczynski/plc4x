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

// Code generated by "plc4xGenerator -type=Connection"; DO NOT EDIT.

package opcua

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

var _ = fmt.Printf

func (d *Connection) Serialize() ([]byte, error) {
	if d == nil {
		return nil, fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := d.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (d *Connection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	if d == nil {
		return fmt.Errorf("(*DeviceInfoCache)(nil)")
	}
	if err := writeBuffer.PushContext("Connection"); err != nil {
		return err
	}
	if err := d.DefaultConnection.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
		return err
	}
	if d.messageCodec != nil {
		{
			_value := fmt.Sprintf("%v", *d.messageCodec)

			if err := writeBuffer.WriteString("messageCodec", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PushContext("subscribers", utils.WithRenderAsList(true)); err != nil {
		return err
	}
	for _, elem := range d.subscribers {
		var elem any = elem

		if elem != nil {
			if serializableField, ok := any(elem).(utils.Serializable); ok {
				if err := writeBuffer.PushContext("value"); err != nil {
					return err
				}
				if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
					return err
				}
				if err := writeBuffer.PopContext("value"); err != nil {
					return err
				}
			} else {
				stringValue := fmt.Sprintf("%v", elem)
				if err := writeBuffer.WriteString("value", uint32(len(stringValue)*8), stringValue); err != nil {
					return err
				}
			}
		}
	}
	if err := writeBuffer.PopContext("subscribers", utils.WithRenderAsList(true)); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("configuration", uint32(len(d.configuration.String())*8), d.configuration.String()); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("driverContext", uint32(len(d.driverContext.String())*8), d.driverContext.String()); err != nil {
		return err
	}
	if d.channel != nil {
		{
			_value := fmt.Sprintf("%v", *d.channel)

			if err := writeBuffer.WriteString("channel", uint32(len(_value)*8), _value); err != nil {
				return err
			}
		}
	}

	_connectEvent_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.connectEvent))
	if err := writeBuffer.WriteString("connectEvent", uint32(len(_connectEvent_plx4gen_description)*8), _connectEvent_plx4gen_description); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("connectTimeout", uint32(len(fmt.Sprintf("%s", d.connectTimeout))*8), fmt.Sprintf("%s", d.connectTimeout)); err != nil {
		return err
	}

	_disconnectEvent_plx4gen_description := fmt.Sprintf("%d element(s)", len(d.disconnectEvent))
	if err := writeBuffer.WriteString("disconnectEvent", uint32(len(_disconnectEvent_plx4gen_description)*8), _disconnectEvent_plx4gen_description); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("disconnectTimeout", uint32(len(fmt.Sprintf("%s", d.disconnectTimeout))*8), fmt.Sprintf("%s", d.disconnectTimeout)); err != nil {
		return err
	}

	if err := writeBuffer.WriteString("connectionId", uint32(len(d.connectionId)*8), d.connectionId); err != nil {
		return err
	}

	if d.tracer != nil {
		if serializableField, ok := any(d.tracer).(utils.Serializable); ok {
			if err := writeBuffer.PushContext("tracer"); err != nil {
				return err
			}
			if err := serializableField.SerializeWithWriteBuffer(ctx, writeBuffer); err != nil {
				return err
			}
			if err := writeBuffer.PopContext("tracer"); err != nil {
				return err
			}
		} else {
			stringValue := fmt.Sprintf("%v", d.tracer)
			if err := writeBuffer.WriteString("tracer", uint32(len(stringValue)*8), stringValue); err != nil {
				return err
			}
		}
	}
	if err := writeBuffer.PopContext("Connection"); err != nil {
		return err
	}
	return nil
}

func (d *Connection) String() string {
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
