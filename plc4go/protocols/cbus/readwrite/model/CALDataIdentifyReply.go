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

// CALDataIdentifyReply is the corresponding interface of CALDataIdentifyReply
type CALDataIdentifyReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetAttribute returns Attribute (property field)
	GetAttribute() Attribute
	// GetIdentifyReplyCommand returns IdentifyReplyCommand (property field)
	GetIdentifyReplyCommand() IdentifyReplyCommand
}

// CALDataIdentifyReplyExactly can be used when we want exactly this type and not a type which fulfills CALDataIdentifyReply.
// This is useful for switch cases.
type CALDataIdentifyReplyExactly interface {
	CALDataIdentifyReply
	isCALDataIdentifyReply() bool
}

// _CALDataIdentifyReply is the data-structure of this message
type _CALDataIdentifyReply struct {
	CALDataContract
	Attribute            Attribute
	IdentifyReplyCommand IdentifyReplyCommand
}

var _ CALDataIdentifyReply = (*_CALDataIdentifyReply)(nil)
var _ CALDataRequirements = (*_CALDataIdentifyReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataIdentifyReply) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataIdentifyReply) GetAttribute() Attribute {
	return m.Attribute
}

func (m *_CALDataIdentifyReply) GetIdentifyReplyCommand() IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataIdentifyReply factory function for _CALDataIdentifyReply
func NewCALDataIdentifyReply(attribute Attribute, identifyReplyCommand IdentifyReplyCommand, commandTypeContainer CALCommandTypeContainer, additionalData CALData, requestContext RequestContext) *_CALDataIdentifyReply {
	_result := &_CALDataIdentifyReply{
		CALDataContract:      NewCALData(commandTypeContainer, additionalData, requestContext),
		Attribute:            attribute,
		IdentifyReplyCommand: identifyReplyCommand,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataIdentifyReply(structType any) CALDataIdentifyReply {
	if casted, ok := structType.(CALDataIdentifyReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataIdentifyReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataIdentifyReply) GetTypeName() string {
	return "CALDataIdentifyReply"
}

func (m *_CALDataIdentifyReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (attribute)
	lengthInBits += 8

	// Simple field (identifyReplyCommand)
	lengthInBits += m.IdentifyReplyCommand.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALDataIdentifyReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataIdentifyReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, commandTypeContainer CALCommandTypeContainer, requestContext RequestContext) (__cALDataIdentifyReply CALDataIdentifyReply, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataIdentifyReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataIdentifyReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	attribute, err := ReadEnumField[Attribute](ctx, "attribute", "Attribute", ReadEnum(AttributeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attribute' field"))
	}
	m.Attribute = attribute

	identifyReplyCommand, err := ReadSimpleField[IdentifyReplyCommand](ctx, "identifyReplyCommand", ReadComplex[IdentifyReplyCommand](IdentifyReplyCommandParseWithBufferProducer[IdentifyReplyCommand]((Attribute)(attribute), (uint8)(uint8(commandTypeContainer.NumBytes())-uint8(uint8(1)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifyReplyCommand' field"))
	}
	m.IdentifyReplyCommand = identifyReplyCommand

	if closeErr := readBuffer.CloseContext("CALDataIdentifyReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataIdentifyReply")
	}

	return m, nil
}

func (m *_CALDataIdentifyReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataIdentifyReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataIdentifyReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataIdentifyReply")
		}

		if err := WriteSimpleEnumField[Attribute](ctx, "attribute", "Attribute", m.GetAttribute(), WriteEnum[Attribute, uint8](Attribute.GetValue, Attribute.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'attribute' field")
		}

		if err := WriteSimpleField[IdentifyReplyCommand](ctx, "identifyReplyCommand", m.GetIdentifyReplyCommand(), WriteComplex[IdentifyReplyCommand](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'identifyReplyCommand' field")
		}

		if popErr := writeBuffer.PopContext("CALDataIdentifyReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataIdentifyReply")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataIdentifyReply) isCALDataIdentifyReply() bool {
	return true
}

func (m *_CALDataIdentifyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
