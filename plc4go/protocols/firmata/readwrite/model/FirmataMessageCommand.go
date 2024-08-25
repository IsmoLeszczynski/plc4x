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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FirmataMessageCommand is the corresponding interface of FirmataMessageCommand
type FirmataMessageCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	FirmataMessage
	// GetCommand returns Command (property field)
	GetCommand() FirmataCommand
}

// FirmataMessageCommandExactly can be used when we want exactly this type and not a type which fulfills FirmataMessageCommand.
// This is useful for switch cases.
type FirmataMessageCommandExactly interface {
	FirmataMessageCommand
	isFirmataMessageCommand() bool
}

// _FirmataMessageCommand is the data-structure of this message
type _FirmataMessageCommand struct {
	*_FirmataMessage
	Command FirmataCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataMessageCommand) GetMessageType() uint8 {
	return 0xF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataMessageCommand) InitializeParent(parent FirmataMessage) {}

func (m *_FirmataMessageCommand) GetParent() FirmataMessage {
	return m._FirmataMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataMessageCommand) GetCommand() FirmataCommand {
	return m.Command
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFirmataMessageCommand factory function for _FirmataMessageCommand
func NewFirmataMessageCommand(command FirmataCommand, response bool) *_FirmataMessageCommand {
	_result := &_FirmataMessageCommand{
		Command:         command,
		_FirmataMessage: NewFirmataMessage(response),
	}
	_result._FirmataMessage._FirmataMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataMessageCommand(structType any) FirmataMessageCommand {
	if casted, ok := structType.(FirmataMessageCommand); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataMessageCommand); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataMessageCommand) GetTypeName() string {
	return "FirmataMessageCommand"
}

func (m *_FirmataMessageCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_FirmataMessageCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataMessageCommandParse(ctx context.Context, theBytes []byte, response bool) (FirmataMessageCommand, error) {
	return FirmataMessageCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)), response)
}

func FirmataMessageCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataMessageCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("FirmataMessageCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataMessageCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := FirmataCommandParseWithBuffer(ctx, readBuffer, bool(response))
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of FirmataMessageCommand")
	}
	command := _command.(FirmataCommand)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	if closeErr := readBuffer.CloseContext("FirmataMessageCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataMessageCommand")
	}

	// Create a partially initialized instance
	_child := &_FirmataMessageCommand{
		_FirmataMessage: &_FirmataMessage{
			Response: response,
		},
		Command: command,
	}
	_child._FirmataMessage._FirmataMessageChildRequirements = _child
	return _child, nil
}

func (m *_FirmataMessageCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataMessageCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataMessageCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataMessageCommand")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(ctx, m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("FirmataMessageCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataMessageCommand")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FirmataMessageCommand) isFirmataMessageCommand() bool {
	return true
}

func (m *_FirmataMessageCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
