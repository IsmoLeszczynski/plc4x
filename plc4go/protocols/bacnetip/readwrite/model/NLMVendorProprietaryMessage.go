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

// NLMVendorProprietaryMessage is the corresponding interface of NLMVendorProprietaryMessage
type NLMVendorProprietaryMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NLM
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorId
	// GetProprietaryMessage returns ProprietaryMessage (property field)
	GetProprietaryMessage() []byte
}

// NLMVendorProprietaryMessageExactly can be used when we want exactly this type and not a type which fulfills NLMVendorProprietaryMessage.
// This is useful for switch cases.
type NLMVendorProprietaryMessageExactly interface {
	NLMVendorProprietaryMessage
	isNLMVendorProprietaryMessage() bool
}

// _NLMVendorProprietaryMessage is the data-structure of this message
type _NLMVendorProprietaryMessage struct {
	*_NLM
	VendorId           BACnetVendorId
	ProprietaryMessage []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMVendorProprietaryMessage) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMVendorProprietaryMessage) InitializeParent(parent NLM) {}

func (m *_NLMVendorProprietaryMessage) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMVendorProprietaryMessage) GetVendorId() BACnetVendorId {
	return m.VendorId
}

func (m *_NLMVendorProprietaryMessage) GetProprietaryMessage() []byte {
	return m.ProprietaryMessage
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMVendorProprietaryMessage factory function for _NLMVendorProprietaryMessage
func NewNLMVendorProprietaryMessage(vendorId BACnetVendorId, proprietaryMessage []byte, apduLength uint16) *_NLMVendorProprietaryMessage {
	_result := &_NLMVendorProprietaryMessage{
		VendorId:           vendorId,
		ProprietaryMessage: proprietaryMessage,
		_NLM:               NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMVendorProprietaryMessage(structType any) NLMVendorProprietaryMessage {
	if casted, ok := structType.(NLMVendorProprietaryMessage); ok {
		return casted
	}
	if casted, ok := structType.(*NLMVendorProprietaryMessage); ok {
		return *casted
	}
	return nil
}

func (m *_NLMVendorProprietaryMessage) GetTypeName() string {
	return "NLMVendorProprietaryMessage"
}

func (m *_NLMVendorProprietaryMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (vendorId)
	lengthInBits += 16

	// Array field
	if len(m.ProprietaryMessage) > 0 {
		lengthInBits += 8 * uint16(len(m.ProprietaryMessage))
	}

	return lengthInBits
}

func (m *_NLMVendorProprietaryMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMVendorProprietaryMessageParse(ctx context.Context, theBytes []byte, apduLength uint16) (NLMVendorProprietaryMessage, error) {
	return NLMVendorProprietaryMessageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMVendorProprietaryMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMVendorProprietaryMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NLMVendorProprietaryMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMVendorProprietaryMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
	}
	_vendorId, _vendorIdErr := BACnetVendorIdParseWithBuffer(ctx, readBuffer)
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field of NLMVendorProprietaryMessage")
	}
	vendorId := _vendorId
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorId")
	}
	// Byte Array field (proprietaryMessage)
	numberOfBytesproprietaryMessage := int(utils.InlineIf((bool((apduLength) > (0))), func() any { return uint16((uint16(apduLength) - uint16(uint16(3)))) }, func() any { return uint16(uint16(0)) }).(uint16))
	proprietaryMessage, _readArrayErr := readBuffer.ReadByteArray("proprietaryMessage", numberOfBytesproprietaryMessage)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'proprietaryMessage' field of NLMVendorProprietaryMessage")
	}

	if closeErr := readBuffer.CloseContext("NLMVendorProprietaryMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMVendorProprietaryMessage")
	}

	// Create a partially initialized instance
	_child := &_NLMVendorProprietaryMessage{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		VendorId:           vendorId,
		ProprietaryMessage: proprietaryMessage,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMVendorProprietaryMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMVendorProprietaryMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMVendorProprietaryMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMVendorProprietaryMessage")
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		_vendorIdErr := writeBuffer.WriteSerializable(ctx, m.GetVendorId())
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Array Field (proprietaryMessage)
		// Byte Array field (proprietaryMessage)
		if err := writeBuffer.WriteByteArray("proprietaryMessage", m.GetProprietaryMessage()); err != nil {
			return errors.Wrap(err, "Error serializing 'proprietaryMessage' field")
		}

		if popErr := writeBuffer.PopContext("NLMVendorProprietaryMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMVendorProprietaryMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMVendorProprietaryMessage) isNLMVendorProprietaryMessage() bool {
	return true
}

func (m *_NLMVendorProprietaryMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
