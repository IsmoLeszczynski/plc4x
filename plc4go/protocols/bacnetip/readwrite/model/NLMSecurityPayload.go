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

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NLMSecurityPayload is the corresponding interface of NLMSecurityPayload
type NLMSecurityPayload interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetPayloadLength returns PayloadLength (property field)
	GetPayloadLength() uint16
	// GetPayload returns Payload (property field)
	GetPayload() []byte
}

// NLMSecurityPayloadExactly can be used when we want exactly this type and not a type which fulfills NLMSecurityPayload.
// This is useful for switch cases.
type NLMSecurityPayloadExactly interface {
	NLMSecurityPayload
	isNLMSecurityPayload() bool
}

// _NLMSecurityPayload is the data-structure of this message
type _NLMSecurityPayload struct {
	NLMContract
	PayloadLength uint16
	Payload       []byte
}

var _ NLMSecurityPayload = (*_NLMSecurityPayload)(nil)
var _ NLMRequirements = (*_NLMSecurityPayload)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMSecurityPayload) GetMessageType() uint8 {
	return 0x0B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMSecurityPayload) GetParent() NLMContract {
	return m.NLMContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMSecurityPayload) GetPayloadLength() uint16 {
	return m.PayloadLength
}

func (m *_NLMSecurityPayload) GetPayload() []byte {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMSecurityPayload factory function for _NLMSecurityPayload
func NewNLMSecurityPayload(payloadLength uint16, payload []byte, apduLength uint16) *_NLMSecurityPayload {
	_result := &_NLMSecurityPayload{
		NLMContract:   NewNLM(apduLength),
		PayloadLength: payloadLength,
		Payload:       payload,
	}
	_result.NLMContract.(*_NLM)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMSecurityPayload(structType any) NLMSecurityPayload {
	if casted, ok := structType.(NLMSecurityPayload); ok {
		return casted
	}
	if casted, ok := structType.(*NLMSecurityPayload); ok {
		return *casted
	}
	return nil
}

func (m *_NLMSecurityPayload) GetTypeName() string {
	return "NLMSecurityPayload"
}

func (m *_NLMSecurityPayload) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NLMContract.(*_NLM).getLengthInBits(ctx))

	// Simple field (payloadLength)
	lengthInBits += 16

	// Array field
	if len(m.Payload) > 0 {
		lengthInBits += 8 * uint16(len(m.Payload))
	}

	return lengthInBits
}

func (m *_NLMSecurityPayload) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NLMSecurityPayload) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NLM, apduLength uint16) (__nLMSecurityPayload NLMSecurityPayload, err error) {
	m.NLMContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMSecurityPayload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMSecurityPayload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payloadLength, err := ReadSimpleField(ctx, "payloadLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payloadLength' field"))
	}
	m.PayloadLength = payloadLength

	payload, err := readBuffer.ReadByteArray("payload", int(payloadLength))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	if closeErr := readBuffer.CloseContext("NLMSecurityPayload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMSecurityPayload")
	}

	return m, nil
}

func (m *_NLMSecurityPayload) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMSecurityPayload) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMSecurityPayload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMSecurityPayload")
		}

		if err := WriteSimpleField[uint16](ctx, "payloadLength", m.GetPayloadLength(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'payloadLength' field")
		}

		if err := WriteByteArrayField(ctx, "payload", m.GetPayload(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("NLMSecurityPayload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMSecurityPayload")
		}
		return nil
	}
	return m.NLMContract.(*_NLM).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMSecurityPayload) isNLMSecurityPayload() bool {
	return true
}

func (m *_NLMSecurityPayload) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
