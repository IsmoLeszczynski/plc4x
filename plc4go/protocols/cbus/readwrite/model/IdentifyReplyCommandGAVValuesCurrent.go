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
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// IdentifyReplyCommandGAVValuesCurrent is the corresponding interface of IdentifyReplyCommandGAVValuesCurrent
type IdentifyReplyCommandGAVValuesCurrent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
}

// _IdentifyReplyCommandGAVValuesCurrent is the data-structure of this message
type _IdentifyReplyCommandGAVValuesCurrent struct {
	IdentifyReplyCommandContract
	Values []byte
}

var _ IdentifyReplyCommandGAVValuesCurrent = (*_IdentifyReplyCommandGAVValuesCurrent)(nil)
var _ IdentifyReplyCommandRequirements = (*_IdentifyReplyCommandGAVValuesCurrent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetAttribute() Attribute {
	return Attribute_GAVValuesCurrent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetParent() IdentifyReplyCommandContract {
	return m.IdentifyReplyCommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetValues() []byte {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandGAVValuesCurrent factory function for _IdentifyReplyCommandGAVValuesCurrent
func NewIdentifyReplyCommandGAVValuesCurrent(values []byte, numBytes uint8) *_IdentifyReplyCommandGAVValuesCurrent {
	_result := &_IdentifyReplyCommandGAVValuesCurrent{
		IdentifyReplyCommandContract: NewIdentifyReplyCommand(numBytes),
		Values:                       values,
	}
	_result.IdentifyReplyCommandContract.(*_IdentifyReplyCommand)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVValuesCurrent(structType any) IdentifyReplyCommandGAVValuesCurrent {
	if casted, ok := structType.(IdentifyReplyCommandGAVValuesCurrent); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVValuesCurrent); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesCurrent"
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).getLengthInBits(ctx))

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_IdentifyReplyCommand, attribute Attribute, numBytes uint8) (__identifyReplyCommandGAVValuesCurrent IdentifyReplyCommandGAVValuesCurrent, err error) {
	m.IdentifyReplyCommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesCurrent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVValuesCurrent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	values, err := readBuffer.ReadByteArray("values", int(numBytes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'values' field"))
	}
	m.Values = values

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesCurrent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVValuesCurrent")
	}

	return m, nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesCurrent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVValuesCurrent")
		}

		if err := WriteByteArrayField(ctx, "values", m.GetValues(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'values' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesCurrent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVValuesCurrent")
		}
		return nil
	}
	return m.IdentifyReplyCommandContract.(*_IdentifyReplyCommand).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) IsIdentifyReplyCommandGAVValuesCurrent() {}

func (m *_IdentifyReplyCommandGAVValuesCurrent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
