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

// TelephonyDataDivert is the corresponding interface of TelephonyDataDivert
type TelephonyDataDivert interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetNumber returns Number (property field)
	GetNumber() string
}

// _TelephonyDataDivert is the data-structure of this message
type _TelephonyDataDivert struct {
	TelephonyDataContract
	Number string
}

var _ TelephonyDataDivert = (*_TelephonyDataDivert)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataDivert)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDivert) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDivert) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataDivert factory function for _TelephonyDataDivert
func NewTelephonyDataDivert(number string, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataDivert {
	_result := &_TelephonyDataDivert{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
		Number:                number,
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataDivert(structType any) TelephonyDataDivert {
	if casted, ok := structType.(TelephonyDataDivert); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDivert); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDivert) GetTypeName() string {
	return "TelephonyDataDivert"
}

func (m *_TelephonyDataDivert) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_TelephonyDataDivert) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataDivert) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData, commandTypeContainer TelephonyCommandTypeContainer) (__telephonyDataDivert TelephonyDataDivert, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDivert"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDivert")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	number, err := ReadSimpleField(ctx, "number", ReadString(readBuffer, uint32(int32((int32(commandTypeContainer.NumBytes())-int32(int32(1))))*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'number' field"))
	}
	m.Number = number

	if closeErr := readBuffer.CloseContext("TelephonyDataDivert"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDivert")
	}

	return m, nil
}

func (m *_TelephonyDataDivert) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataDivert) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDivert"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDivert")
		}

		if err := WriteSimpleField[string](ctx, "number", m.GetNumber(), WriteString(writeBuffer, int32(int32((int32(m.GetCommandTypeContainer().NumBytes())-int32(int32(1))))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'number' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDivert"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDivert")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataDivert) IsTelephonyDataDivert() {}

func (m *_TelephonyDataDivert) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
