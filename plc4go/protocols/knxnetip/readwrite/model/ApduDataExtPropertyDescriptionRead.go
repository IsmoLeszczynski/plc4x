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

// ApduDataExtPropertyDescriptionRead is the corresponding interface of ApduDataExtPropertyDescriptionRead
type ApduDataExtPropertyDescriptionRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
	// GetObjectIndex returns ObjectIndex (property field)
	GetObjectIndex() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetIndex returns Index (property field)
	GetIndex() uint8
}

// ApduDataExtPropertyDescriptionReadExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtPropertyDescriptionRead.
// This is useful for switch cases.
type ApduDataExtPropertyDescriptionReadExactly interface {
	ApduDataExtPropertyDescriptionRead
	isApduDataExtPropertyDescriptionRead() bool
}

// _ApduDataExtPropertyDescriptionRead is the data-structure of this message
type _ApduDataExtPropertyDescriptionRead struct {
	ApduDataExtContract
	ObjectIndex uint8
	PropertyId  uint8
	Index       uint8
}

var _ ApduDataExtPropertyDescriptionRead = (*_ApduDataExtPropertyDescriptionRead)(nil)
var _ ApduDataExtRequirements = (*_ApduDataExtPropertyDescriptionRead)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) GetExtApciType() uint8 {
	return 0x18
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) GetParent() ApduDataExtContract {
	return m.ApduDataExtContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ApduDataExtPropertyDescriptionRead) GetObjectIndex() uint8 {
	return m.ObjectIndex
}

func (m *_ApduDataExtPropertyDescriptionRead) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_ApduDataExtPropertyDescriptionRead) GetIndex() uint8 {
	return m.Index
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewApduDataExtPropertyDescriptionRead factory function for _ApduDataExtPropertyDescriptionRead
func NewApduDataExtPropertyDescriptionRead(objectIndex uint8, propertyId uint8, index uint8, length uint8) *_ApduDataExtPropertyDescriptionRead {
	_result := &_ApduDataExtPropertyDescriptionRead{
		ApduDataExtContract: NewApduDataExt(length),
		ObjectIndex:         objectIndex,
		PropertyId:          propertyId,
		Index:               index,
	}
	_result.ApduDataExtContract.(*_ApduDataExt)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtPropertyDescriptionRead(structType any) ApduDataExtPropertyDescriptionRead {
	if casted, ok := structType.(ApduDataExtPropertyDescriptionRead); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtPropertyDescriptionRead); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtPropertyDescriptionRead) GetTypeName() string {
	return "ApduDataExtPropertyDescriptionRead"
}

func (m *_ApduDataExtPropertyDescriptionRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduDataExtContract.(*_ApduDataExt).getLengthInBits(ctx))

	// Simple field (objectIndex)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (index)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ApduDataExtPropertyDescriptionRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduDataExtPropertyDescriptionRead) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduDataExt, length uint8) (__apduDataExtPropertyDescriptionRead ApduDataExtPropertyDescriptionRead, err error) {
	m.ApduDataExtContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtPropertyDescriptionRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtPropertyDescriptionRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIndex, err := ReadSimpleField(ctx, "objectIndex", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIndex' field"))
	}
	m.ObjectIndex = objectIndex

	propertyId, err := ReadSimpleField(ctx, "propertyId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyId' field"))
	}
	m.PropertyId = propertyId

	index, err := ReadSimpleField(ctx, "index", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'index' field"))
	}
	m.Index = index

	if closeErr := readBuffer.CloseContext("ApduDataExtPropertyDescriptionRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtPropertyDescriptionRead")
	}

	return m, nil
}

func (m *_ApduDataExtPropertyDescriptionRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtPropertyDescriptionRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtPropertyDescriptionRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtPropertyDescriptionRead")
		}

		if err := WriteSimpleField[uint8](ctx, "objectIndex", m.GetObjectIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIndex' field")
		}

		if err := WriteSimpleField[uint8](ctx, "propertyId", m.GetPropertyId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyId' field")
		}

		if err := WriteSimpleField[uint8](ctx, "index", m.GetIndex(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'index' field")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtPropertyDescriptionRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtPropertyDescriptionRead")
		}
		return nil
	}
	return m.ApduDataExtContract.(*_ApduDataExt).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduDataExtPropertyDescriptionRead) isApduDataExtPropertyDescriptionRead() bool {
	return true
}

func (m *_ApduDataExtPropertyDescriptionRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
