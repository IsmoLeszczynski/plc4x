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

// ApduDataExtLinkWrite is the corresponding interface of ApduDataExtLinkWrite
type ApduDataExtLinkWrite interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtLinkWriteExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtLinkWrite.
// This is useful for switch cases.
type ApduDataExtLinkWriteExactly interface {
	ApduDataExtLinkWrite
	isApduDataExtLinkWrite() bool
}

// _ApduDataExtLinkWrite is the data-structure of this message
type _ApduDataExtLinkWrite struct {
	ApduDataExtContract
}

var _ ApduDataExtLinkWrite = (*_ApduDataExtLinkWrite)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtLinkWrite)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtLinkWrite) GetExtApciType() uint8 {
	return 0x27
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtLinkWrite) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

// NewApduDataExtLinkWrite factory function for _ApduDataExtLinkWrite
func NewApduDataExtLinkWrite(length uint8) *_ApduDataExtLinkWrite {
	_result := &_ApduDataExtLinkWrite{
		ApduDataExtContract: NewApduDataExt(length),
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtLinkWrite(structType any) ApduDataExtLinkWrite {
	if casted, ok := structType.(ApduDataExtLinkWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtLinkWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtLinkWrite) GetTypeName() string {
	return "ApduDataExtLinkWrite"
}

func (m *_ApduDataExtLinkWrite) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtLinkWrite) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtLinkWrite) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtLinkWrite ApduDataExtLinkWrite, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtLinkWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtLinkWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtLinkWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtLinkWrite")
	}

	return m, nil
}

func (m *_ApduDataExtLinkWrite) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtLinkWrite) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtLinkWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtLinkWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtLinkWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtLinkWrite")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtLinkWrite) isApduDataExtLinkWrite() bool {
	return true
}

func (m *_ApduDataExtLinkWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
