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

// PowerUpReply is the corresponding interface of PowerUpReply
type PowerUpReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Reply
	// GetPowerUpIndicator returns PowerUpIndicator (property field)
	GetPowerUpIndicator() PowerUp
}

// PowerUpReplyExactly can be used when we want exactly this type and not a type which fulfills PowerUpReply.
// This is useful for switch cases.
type PowerUpReplyExactly interface {
	PowerUpReply
	isPowerUpReply() bool
}

// _PowerUpReply is the data-structure of this message
type _PowerUpReply struct {
	ReplyContract
	PowerUpIndicator PowerUp
}

var _ PowerUpReply = (*_PowerUpReply)(nil)
var _ ReplyRequirements = (*_PowerUpReply)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PowerUpReply) GetParent() ReplyContract {
	return m.ReplyContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PowerUpReply) GetPowerUpIndicator() PowerUp {
	return m.PowerUpIndicator
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPowerUpReply factory function for _PowerUpReply
func NewPowerUpReply(powerUpIndicator PowerUp, peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_PowerUpReply {
	_result := &_PowerUpReply{
		ReplyContract:    NewReply(peekedByte, cBusOptions, requestContext),
		PowerUpIndicator: powerUpIndicator,
	}
	_result.ReplyContract.(*_Reply)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPowerUpReply(structType any) PowerUpReply {
	if casted, ok := structType.(PowerUpReply); ok {
		return casted
	}
	if casted, ok := structType.(*PowerUpReply); ok {
		return *casted
	}
	return nil
}

func (m *_PowerUpReply) GetTypeName() string {
	return "PowerUpReply"
}

func (m *_PowerUpReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyContract.(*_Reply).getLengthInBits(ctx))

	// Simple field (powerUpIndicator)
	lengthInBits += m.PowerUpIndicator.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_PowerUpReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PowerUpReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_Reply, cBusOptions CBusOptions, requestContext RequestContext) (__powerUpReply PowerUpReply, err error) {
	m.ReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PowerUpReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PowerUpReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	powerUpIndicator, err := ReadSimpleField[PowerUp](ctx, "powerUpIndicator", ReadComplex[PowerUp](PowerUpParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'powerUpIndicator' field"))
	}
	m.PowerUpIndicator = powerUpIndicator

	if closeErr := readBuffer.CloseContext("PowerUpReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PowerUpReply")
	}

	return m, nil
}

func (m *_PowerUpReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PowerUpReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PowerUpReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PowerUpReply")
		}

		if err := WriteSimpleField[PowerUp](ctx, "powerUpIndicator", m.GetPowerUpIndicator(), WriteComplex[PowerUp](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'powerUpIndicator' field")
		}

		if popErr := writeBuffer.PopContext("PowerUpReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PowerUpReply")
		}
		return nil
	}
	return m.ReplyContract.(*_Reply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PowerUpReply) isPowerUpReply() bool {
	return true
}

func (m *_PowerUpReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
