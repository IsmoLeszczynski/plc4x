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

// NullCommandRequest is the corresponding interface of NullCommandRequest
type NullCommandRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// NullCommandRequestExactly can be used when we want exactly this type and not a type which fulfills NullCommandRequest.
// This is useful for switch cases.
type NullCommandRequestExactly interface {
	NullCommandRequest
	isNullCommandRequest() bool
}

// _NullCommandRequest is the data-structure of this message
type _NullCommandRequest struct {
	EipPacketContract
}

var _ NullCommandRequest = (*_NullCommandRequest)(nil)
var _ EipPacketRequirements = (*_NullCommandRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullCommandRequest) GetCommand() uint16 {
	return 0x0001
}

func (m *_NullCommandRequest) GetResponse() bool {
	return bool(false)
}

func (m *_NullCommandRequest) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullCommandRequest) GetParent() EipPacketContract {
	return m.EipPacketContract
}

// NewNullCommandRequest factory function for _NullCommandRequest
func NewNullCommandRequest(sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_NullCommandRequest {
	_result := &_NullCommandRequest{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullCommandRequest(structType any) NullCommandRequest {
	if casted, ok := structType.(NullCommandRequest); ok {
		return casted
	}
	if casted, ok := structType.(*NullCommandRequest); ok {
		return *casted
	}
	return nil
}

func (m *_NullCommandRequest) GetTypeName() string {
	return "NullCommandRequest"
}

func (m *_NullCommandRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullCommandRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NullCommandRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__nullCommandRequest NullCommandRequest, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullCommandRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullCommandRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullCommandRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullCommandRequest")
	}

	return m, nil
}

func (m *_NullCommandRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullCommandRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullCommandRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullCommandRequest")
		}

		if popErr := writeBuffer.PopContext("NullCommandRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullCommandRequest")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullCommandRequest) isNullCommandRequest() bool {
	return true
}

func (m *_NullCommandRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
