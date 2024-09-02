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
const ParameterChange_SPECIALCHAR1 byte = 0x3D
const ParameterChange_SPECIALCHAR2 byte = 0x3D

// ParameterChange is the corresponding interface of ParameterChange
type ParameterChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// _ParameterChange is the data-structure of this message
type _ParameterChange struct {
}

var _ ParameterChange = (*_ParameterChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ParameterChange) GetSpecialChar1() byte {
	return ParameterChange_SPECIALCHAR1
}

func (m *_ParameterChange) GetSpecialChar2() byte {
	return ParameterChange_SPECIALCHAR2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterChange factory function for _ParameterChange
func NewParameterChange() *_ParameterChange {
	return &_ParameterChange{}
}

// Deprecated: use the interface for direct cast
func CastParameterChange(structType any) ParameterChange {
	if casted, ok := structType.(ParameterChange); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterChange); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterChange) GetTypeName() string {
	return "ParameterChange"
}

func (m *_ParameterChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (specialChar1)
	lengthInBits += 8

	// Const Field (specialChar2)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ParameterChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterChangeParse(ctx context.Context, theBytes []byte) (ParameterChange, error) {
	return ParameterChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func ParameterChangeParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (ParameterChange, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (ParameterChange, error) {
		return ParameterChangeParseWithBuffer(ctx, readBuffer)
	}
}

func ParameterChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ParameterChange, error) {
	v, err := (&_ParameterChange{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_ParameterChange) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__parameterChange ParameterChange, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	specialChar1, err := ReadConstField[byte](ctx, "specialChar1", ReadByte(readBuffer, 8), ParameterChange_SPECIALCHAR1)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specialChar1' field"))
	}
	_ = specialChar1

	specialChar2, err := ReadConstField[byte](ctx, "specialChar2", ReadByte(readBuffer, 8), ParameterChange_SPECIALCHAR2)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'specialChar2' field"))
	}
	_ = specialChar2

	if closeErr := readBuffer.CloseContext("ParameterChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterChange")
	}

	return m, nil
}

func (m *_ParameterChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("ParameterChange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterChange")
	}

	if err := WriteConstField(ctx, "specialChar1", ParameterChange_SPECIALCHAR1, WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'specialChar1' field")
	}

	if err := WriteConstField(ctx, "specialChar2", ParameterChange_SPECIALCHAR2, WriteByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'specialChar2' field")
	}

	if popErr := writeBuffer.PopContext("ParameterChange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterChange")
	}
	return nil
}

func (m *_ParameterChange) IsParameterChange() {}

func (m *_ParameterChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
