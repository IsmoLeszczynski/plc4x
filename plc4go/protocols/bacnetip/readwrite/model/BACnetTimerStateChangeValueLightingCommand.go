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

// BACnetTimerStateChangeValueLightingCommand is the corresponding interface of BACnetTimerStateChangeValueLightingCommand
type BACnetTimerStateChangeValueLightingCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetLigthingCommandValue returns LigthingCommandValue (property field)
	GetLigthingCommandValue() BACnetLightingCommandEnclosed
}

// BACnetTimerStateChangeValueLightingCommandExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueLightingCommand.
// This is useful for switch cases.
type BACnetTimerStateChangeValueLightingCommandExactly interface {
	BACnetTimerStateChangeValueLightingCommand
	isBACnetTimerStateChangeValueLightingCommand() bool
}

// _BACnetTimerStateChangeValueLightingCommand is the data-structure of this message
type _BACnetTimerStateChangeValueLightingCommand struct {
	BACnetTimerStateChangeValueContract
	LigthingCommandValue BACnetLightingCommandEnclosed
}

var _ BACnetTimerStateChangeValueLightingCommand = (*_BACnetTimerStateChangeValueLightingCommand)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueLightingCommand)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLigthingCommandValue() BACnetLightingCommandEnclosed {
	return m.LigthingCommandValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerStateChangeValueLightingCommand factory function for _BACnetTimerStateChangeValueLightingCommand
func NewBACnetTimerStateChangeValueLightingCommand(ligthingCommandValue BACnetLightingCommandEnclosed, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueLightingCommand {
	_result := &_BACnetTimerStateChangeValueLightingCommand{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		LigthingCommandValue:                ligthingCommandValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueLightingCommand(structType any) BACnetTimerStateChangeValueLightingCommand {
	if casted, ok := structType.(BACnetTimerStateChangeValueLightingCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueLightingCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetTypeName() string {
	return "BACnetTimerStateChangeValueLightingCommand"
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (ligthingCommandValue)
	lengthInBits += m.LigthingCommandValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueLightingCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueLightingCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueLightingCommand BACnetTimerStateChangeValueLightingCommand, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueLightingCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueLightingCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	ligthingCommandValue, err := ReadSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", ReadComplex[BACnetLightingCommandEnclosed](BACnetLightingCommandEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'ligthingCommandValue' field"))
	}
	m.LigthingCommandValue = ligthingCommandValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueLightingCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueLightingCommand")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueLightingCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueLightingCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueLightingCommand")
		}

		if err := WriteSimpleField[BACnetLightingCommandEnclosed](ctx, "ligthingCommandValue", m.GetLigthingCommandValue(), WriteComplex[BACnetLightingCommandEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'ligthingCommandValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueLightingCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueLightingCommand")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueLightingCommand) isBACnetTimerStateChangeValueLightingCommand() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueLightingCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
