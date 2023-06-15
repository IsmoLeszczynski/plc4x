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

// ApduDataExtDomainAddressSerialNumberRead is the corresponding interface of ApduDataExtDomainAddressSerialNumberRead
type ApduDataExtDomainAddressSerialNumberRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtDomainAddressSerialNumberReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtDomainAddressSerialNumberRead.
// This is useful for switch cases.
type ApduDataExtDomainAddressSerialNumberReadExactly interface {
	ApduDataExtDomainAddressSerialNumberRead
	isApduDataExtDomainAddressSerialNumberRead() bool
}

// _ApduDataExtDomainAddressSerialNumberRead is the data-structure of this message
type _ApduDataExtDomainAddressSerialNumberRead struct {
	*_ApduDataExt
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetExtApciType() uint8 {
	return 0x2C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtDomainAddressSerialNumberRead) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtDomainAddressSerialNumberRead factory function for _ApduDataExtDomainAddressSerialNumberRead
func NewApduDataExtDomainAddressSerialNumberRead(length uint8) *_ApduDataExtDomainAddressSerialNumberRead {
	_result := &_ApduDataExtDomainAddressSerialNumberRead{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtDomainAddressSerialNumberRead(structType any) ApduDataExtDomainAddressSerialNumberRead {
	if casted, ok := structType.(ApduDataExtDomainAddressSerialNumberRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtDomainAddressSerialNumberRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetTypeName() string {
	return "ApduDataExtDomainAddressSerialNumberRead"
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtDomainAddressSerialNumberReadParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtDomainAddressSerialNumberRead, error) {
	return ApduDataExtDomainAddressSerialNumberReadParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtDomainAddressSerialNumberReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtDomainAddressSerialNumberRead, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtDomainAddressSerialNumberRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtDomainAddressSerialNumberRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtDomainAddressSerialNumberRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtDomainAddressSerialNumberRead")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtDomainAddressSerialNumberRead{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtDomainAddressSerialNumberRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtDomainAddressSerialNumberRead")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtDomainAddressSerialNumberRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtDomainAddressSerialNumberRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) isApduDataExtDomainAddressSerialNumberRead() bool {
	return true
}

func (m *_ApduDataExtDomainAddressSerialNumberRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
