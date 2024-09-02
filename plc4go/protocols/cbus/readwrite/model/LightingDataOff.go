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

// LightingDataOff is the corresponding interface of LightingDataOff
type LightingDataOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LightingData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// LightingDataOffExactly can be used when we want exactly this type and not a type which fulfills LightingDataOff.
// This is useful for switch cases.
type LightingDataOffExactly interface {
	LightingDataOff
	isLightingDataOff() bool
}

// _LightingDataOff is the data-structure of this message
type _LightingDataOff struct {
	LightingDataContract
	Group byte
}

var _ LightingDataOff = (*_LightingDataOff)(nil)
var _ LightingDataRequirements = (*_LightingDataOff)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LightingDataOff) GetParent() LightingDataContract {
	return m.LightingDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LightingDataOff) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLightingDataOff factory function for _LightingDataOff
func NewLightingDataOff(group byte, commandTypeContainer LightingCommandTypeContainer) *_LightingDataOff {
	_result := &_LightingDataOff{
		LightingDataContract: NewLightingData(commandTypeContainer),
		Group:                group,
	}
	_result.LightingDataContract.(*_LightingData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLightingDataOff(structType any) LightingDataOff {
	if casted, ok := structType.(LightingDataOff); ok {
		return casted
	}
	if casted, ok := structType.(*LightingDataOff); ok {
		return *casted
	}
	return nil
}

func (m *_LightingDataOff) GetTypeName() string {
	return "LightingDataOff"
}

func (m *_LightingDataOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.LightingDataContract.(*_LightingData).getLengthInBits(ctx))

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_LightingDataOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LightingDataOff) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_LightingData) (__lightingDataOff LightingDataOff, err error) {
	m.LightingDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LightingDataOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LightingDataOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	group, err := ReadSimpleField(ctx, "group", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'group' field"))
	}
	m.Group = group

	if closeErr := readBuffer.CloseContext("LightingDataOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LightingDataOff")
	}

	return m, nil
}

func (m *_LightingDataOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LightingDataOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LightingDataOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LightingDataOff")
		}

		if err := WriteSimpleField[byte](ctx, "group", m.GetGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("LightingDataOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LightingDataOff")
		}
		return nil
	}
	return m.LightingDataContract.(*_LightingData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LightingDataOff) isLightingDataOff() bool {
	return true
}

func (m *_LightingDataOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
