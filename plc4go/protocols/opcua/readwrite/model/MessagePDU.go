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

// MessagePDU is the corresponding interface of MessagePDU
type MessagePDU interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageType returns MessageType (discriminator field)
	GetMessageType() string
	// GetResponse returns Response (discriminator field)
	GetResponse() bool
	// GetChunk returns Chunk (property field)
	GetChunk() ChunkType
}

// MessagePDUExactly can be used when we want exactly this type and not a type which fulfills MessagePDU.
// This is useful for switch cases.
type MessagePDUExactly interface {
	MessagePDU
	isMessagePDU() bool
}

// _MessagePDU is the data-structure of this message
type _MessagePDU struct {
	_MessagePDUChildRequirements
	Chunk ChunkType
}

type _MessagePDUChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageType() string
	GetResponse() bool
}

type MessagePDUParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MessagePDU, serializeChildFunction func() error) error
	GetTypeName() string
}

type MessagePDUChild interface {
	utils.Serializable
	InitializeParent(parent MessagePDU, chunk ChunkType)
	GetParent() *MessagePDU

	GetTypeName() string
	MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MessagePDU) GetChunk() ChunkType {
	return m.Chunk
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMessagePDU factory function for _MessagePDU
func NewMessagePDU(chunk ChunkType) *_MessagePDU {
	return &_MessagePDU{Chunk: chunk}
}

// Deprecated: use the interface for direct cast
func CastMessagePDU(structType any) MessagePDU {
	if casted, ok := structType.(MessagePDU); ok {
		return casted
	}
	if casted, ok := structType.(*MessagePDU); ok {
		return *casted
	}
	return nil
}

func (m *_MessagePDU) GetTypeName() string {
	return "MessagePDU"
}

func (m *_MessagePDU) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageType)
	lengthInBits += 24

	// Simple field (chunk)
	lengthInBits += 8

	// Implicit Field (totalLength)
	lengthInBits += 32

	return lengthInBits
}

func (m *_MessagePDU) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MessagePDUParse(ctx context.Context, theBytes []byte, response bool) (MessagePDU, error) {
	return MessagePDUParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func MessagePDUParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (MessagePDU, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MessagePDU"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MessagePDU")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	messageType, err := ReadDiscriminatorField[string](ctx, "messageType", ReadString(readBuffer, 24))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'messageType' field"))
	}

	// Simple Field (chunk)
	if pullErr := readBuffer.PullContext("chunk"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for chunk")
	}
	_chunk, _chunkErr := ChunkTypeParseWithBuffer(ctx, readBuffer)
	if _chunkErr != nil {
		return nil, errors.Wrap(_chunkErr, "Error parsing 'chunk' field of MessagePDU")
	}
	chunk := _chunk
	if closeErr := readBuffer.CloseContext("chunk"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for chunk")
	}

	// Implicit Field (totalLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	totalLength, _totalLengthErr := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadUint32("totalLength", 32)
	_ = totalLength
	if _totalLengthErr != nil {
		return nil, errors.Wrap(_totalLengthErr, "Error parsing 'totalLength' field of MessagePDU")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type MessagePDUChildSerializeRequirement interface {
		MessagePDU
		InitializeParent(MessagePDU, ChunkType)
		GetParent() MessagePDU
	}
	var _childTemp any
	var _child MessagePDUChildSerializeRequirement
	var typeSwitchError error
	switch {
	case messageType == "HEL" && response == bool(false): // OpcuaHelloRequest
		_childTemp, typeSwitchError = OpcuaHelloRequestParseWithBuffer(ctx, readBuffer, response)
	case messageType == "ACK" && response == bool(true): // OpcuaAcknowledgeResponse
		_childTemp, typeSwitchError = OpcuaAcknowledgeResponseParseWithBuffer(ctx, readBuffer, response)
	case messageType == "OPN" && response == bool(false): // OpcuaOpenRequest
		_childTemp, typeSwitchError = OpcuaOpenRequestParseWithBuffer(ctx, readBuffer, totalLength, response)
	case messageType == "OPN" && response == bool(true): // OpcuaOpenResponse
		_childTemp, typeSwitchError = OpcuaOpenResponseParseWithBuffer(ctx, readBuffer, totalLength, response)
	case messageType == "CLO" && response == bool(false): // OpcuaCloseRequest
		_childTemp, typeSwitchError = OpcuaCloseRequestParseWithBuffer(ctx, readBuffer, response)
	case messageType == "MSG" && response == bool(false): // OpcuaMessageRequest
		_childTemp, typeSwitchError = OpcuaMessageRequestParseWithBuffer(ctx, readBuffer, totalLength, response)
	case messageType == "MSG" && response == bool(true): // OpcuaMessageResponse
		_childTemp, typeSwitchError = OpcuaMessageResponseParseWithBuffer(ctx, readBuffer, totalLength, response)
	case messageType == "ERR" && response == bool(true): // OpcuaMessageError
		_childTemp, typeSwitchError = OpcuaMessageErrorParseWithBuffer(ctx, readBuffer, response)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageType=%v, response=%v]", messageType, response)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of MessagePDU")
	}
	_child = _childTemp.(MessagePDUChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("MessagePDU"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MessagePDU")
	}

	// Finish initializing
	_child.InitializeParent(_child, chunk)
	return _child, nil
}

func (pm *_MessagePDU) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child MessagePDU, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("MessagePDU"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for MessagePDU")
	}

	// Discriminator Field (messageType) (Used as input to a switch field)
	messageType := string(child.GetMessageType())
	_messageTypeErr := /*TODO: migrate me*/ writeBuffer.WriteString("messageType", uint32(24), (messageType), utils.WithEncoding("UTF-8)"))

	if _messageTypeErr != nil {
		return errors.Wrap(_messageTypeErr, "Error serializing 'messageType' field")
	}

	// Simple Field (chunk)
	if pushErr := writeBuffer.PushContext("chunk"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for chunk")
	}
	_chunkErr := writeBuffer.WriteSerializable(ctx, m.GetChunk())
	if popErr := writeBuffer.PopContext("chunk"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for chunk")
	}
	if _chunkErr != nil {
		return errors.Wrap(_chunkErr, "Error serializing 'chunk' field")
	}

	// Implicit Field (totalLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	totalLength := uint32(uint32(m.GetLengthInBytes(ctx)))
	_totalLengthErr := /*TODO: migrate me*/ writeBuffer.WriteUint32("totalLength", 32, uint32((totalLength)))
	if _totalLengthErr != nil {
		return errors.Wrap(_totalLengthErr, "Error serializing 'totalLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("MessagePDU"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for MessagePDU")
	}
	return nil
}

func (m *_MessagePDU) isMessagePDU() bool {
	return true
}

func (m *_MessagePDU) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
