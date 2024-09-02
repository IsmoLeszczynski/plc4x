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

// IdentifyReplyCommandCurrentSenseLevels is the corresponding interface of IdentifyReplyCommandCurrentSenseLevels
type IdentifyReplyCommandCurrentSenseLevels interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetCurrentSenseLevels returns CurrentSenseLevels (property field)
	GetCurrentSenseLevels() []byte
}

// IdentifyReplyCommandCurrentSenseLevelsExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandCurrentSenseLevels.
// This is useful for switch cases.
type IdentifyReplyCommandCurrentSenseLevelsExactly interface {
	IdentifyReplyCommandCurrentSenseLevels
	isIdentifyReplyCommandCurrentSenseLevels() bool
}

// _IdentifyReplyCommandCurrentSenseLevels is the data-structure of this message
type _IdentifyReplyCommandCurrentSenseLevels struct {
	IdentifyReplyCommandContract
	CurrentSenseLevels []byte
}

var _ IdentifyReplyCommandCurrentSenseLevels = (*_IdentifyReplyCommandCurrentSenseLevels)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandCurrentSenseLevels)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetAttribute() Attribute {
	return Attribute_CurrentSenseLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetCurrentSenseLevels() []byte {
	return m.CurrentSenseLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandCurrentSenseLevels factory function for _IdentifyReplyCommandCurrentSenseLevels
func NewIdentifyReplyCommandCurrentSenseLevels(currentSenseLevels []byte, numBytes uint8) *_IdentifyReplyCommandCurrentSenseLevels {
	_result := &_IdentifyReplyCommandCurrentSenseLevels{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		CurrentSenseLevels:           currentSenseLevels,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandCurrentSenseLevels(structType any) IdentifyReplyCommandCurrentSenseLevels {
	if casted, ok := structType.(IdentifyReplyCommandCurrentSenseLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandCurrentSenseLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetTypeName() string {
	return "IdentifyReplyCommandCurrentSenseLevels"
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.CurrentSenseLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.CurrentSenseLevels))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandCurrentSenseLevels IdentifyReplyCommandCurrentSenseLevels, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandCurrentSenseLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandCurrentSenseLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	currentSenseLevels, err := readBuffer.ReadByteArray("currentSenseLevels", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'currentSenseLevels' field"))
	}
	m.CurrentSenseLevels = currentSenseLevels

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandCurrentSenseLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandCurrentSenseLevels")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandCurrentSenseLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandCurrentSenseLevels")
		}

		if err := WriteByteArrayField(ctx, "currentSenseLevels", m.GetCurrentSenseLevels(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'currentSenseLevels' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandCurrentSenseLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandCurrentSenseLevels")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) isIdentifyReplyCommandCurrentSenseLevels() bool {
	return true
}

func (m *_IdentifyReplyCommandCurrentSenseLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
