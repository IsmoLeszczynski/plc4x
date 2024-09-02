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

// Constant values.
const CycServiceItemType_FUNCTIONID uint8 = 0x12

// CycServiceItemType is the corresponding interface of CycServiceItemType
type CycServiceItemType interface {
	CycServiceItemTypeContract
	CycServiceItemTypeRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// CycServiceItemTypeContract provides a set of functions which can be overwritten by a sub struct
type CycServiceItemTypeContract interface {
	// GetByteLength returns ByteLength (property field)
	GetByteLength() uint8
	// GetSyntaxId returns SyntaxId (property field)
	GetSyntaxId() uint8
	// IsCycServiceItemType() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCycServiceItemType()
}

// CycServiceItemTypeRequirements provides a set of functions which need to be implemented by a sub struct
type CycServiceItemTypeRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetSyntaxId returns SyntaxId (discriminator field)
	GetSyntaxId() uint8
}

// _CycServiceItemType is the data-structure of this message
type _CycServiceItemType struct {
	_SubType   CycServiceItemType
	ByteLength uint8
	SyntaxId   uint8
}

var _ CycServiceItemTypeContract = (*_CycServiceItemType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CycServiceItemType) GetByteLength() uint8 {
	return m.ByteLength
}

func (m *_CycServiceItemType) GetSyntaxId() uint8 {
	return m.SyntaxId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CycServiceItemType) GetFunctionId() uint8 {
	return CycServiceItemType_FUNCTIONID
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCycServiceItemType factory function for _CycServiceItemType
func NewCycServiceItemType(byteLength uint8, syntaxId uint8) *_CycServiceItemType {
	return &_CycServiceItemType{ByteLength: byteLength, SyntaxId: syntaxId}
}

// Deprecated: use the interface for direct cast
func CastCycServiceItemType(structType any) CycServiceItemType {
	if casted, ok := structType.(CycServiceItemType); ok {
		return casted
	}
	if casted, ok := structType.(*CycServiceItemType); ok {
		return *casted
	}
	return nil
}

func (m *_CycServiceItemType) GetTypeName() string {
	return "CycServiceItemType"
}

func (m *_CycServiceItemType) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (functionId)
	lengthInBits += 8

	// Simple field (byteLength)
	lengthInBits += 8

	// Simple field (syntaxId)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CycServiceItemType) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func CycServiceItemTypeParse[T CycServiceItemType](ctx context.Context, theBytes []byte) (T, error) {
	return CycServiceItemTypeParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func CycServiceItemTypeParseWithBufferProducer[T CycServiceItemType]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := CycServiceItemTypeParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func CycServiceItemTypeParseWithBuffer[T CycServiceItemType](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_CycServiceItemType{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_CycServiceItemType) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__cycServiceItemType CycServiceItemType, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CycServiceItemType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CycServiceItemType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	functionId, err := ReadConstField[uint8](ctx, "functionId", ReadUnsignedByte(readBuffer, uint8(8)), CycServiceItemType_FUNCTIONID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'functionId' field"))
	}
	_ = functionId

	byteLength, err := ReadSimpleField(ctx, "byteLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteLength' field"))
	}
	m.ByteLength = byteLength

	syntaxId, err := ReadSimpleField(ctx, "syntaxId", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'syntaxId' field"))
	}
	m.SyntaxId = syntaxId

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child CycServiceItemType
	switch {
	case syntaxId == 0x10: // CycServiceItemAnyType
		if _child, err = (&_CycServiceItemAnyType{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CycServiceItemAnyType for type-switch of CycServiceItemType")
		}
	case syntaxId == 0xb0: // CycServiceItemDbReadType
		if _child, err = (&_CycServiceItemDbReadType{}).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type CycServiceItemDbReadType for type-switch of CycServiceItemType")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [syntaxId=%v]", syntaxId)
	}

	if closeErr := readBuffer.CloseContext("CycServiceItemType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CycServiceItemType")
	}

	return _child, nil
}

func (pm *_CycServiceItemType) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CycServiceItemType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CycServiceItemType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CycServiceItemType")
	}

	if err := WriteConstField(ctx, "functionId", CycServiceItemType_FUNCTIONID, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'functionId' field")
	}

	if err := WriteSimpleField[uint8](ctx, "byteLength", m.GetByteLength(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'byteLength' field")
	}

	if err := WriteSimpleField[uint8](ctx, "syntaxId", m.GetSyntaxId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'syntaxId' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CycServiceItemType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CycServiceItemType")
	}
	return nil
}

func (m *_CycServiceItemType) IsCycServiceItemType() {}
