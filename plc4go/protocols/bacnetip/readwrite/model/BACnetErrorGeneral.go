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

// BACnetErrorGeneral is the corresponding interface of BACnetErrorGeneral
type BACnetErrorGeneral interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetError returns Error (property field)
	GetError() Error
}

// _BACnetErrorGeneral is the data-structure of this message
type _BACnetErrorGeneral struct {
	BACnetErrorContract
	Error Error
}

var _ BACnetErrorGeneral = (*_BACnetErrorGeneral)(nil)
var _ BACnetErrorRequirements = (*_BACnetErrorGeneral)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetErrorGeneral) GetErrorChoice() BACnetConfirmedServiceChoice {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetErrorGeneral) GetParent() BACnetErrorContract {
	return m.BACnetErrorContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetErrorGeneral) GetError() Error {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetErrorGeneral factory function for _BACnetErrorGeneral
func NewBACnetErrorGeneral(error Error) *_BACnetErrorGeneral {
	_result := &_BACnetErrorGeneral{
		BACnetErrorContract: NewBACnetError(),
		Error:               error,
	}
	_result.BACnetErrorContract.(*_BACnetError)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetErrorGeneral(structType any) BACnetErrorGeneral {
	if casted, ok := structType.(BACnetErrorGeneral); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetErrorGeneral); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetErrorGeneral) GetTypeName() string {
	return "BACnetErrorGeneral"
}

func (m *_BACnetErrorGeneral) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetErrorContract.(*_BACnetError).getLengthInBits(ctx))

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetErrorGeneral) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetErrorGeneral) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetError, errorChoice BACnetConfirmedServiceChoice) (__bACnetErrorGeneral BACnetErrorGeneral, err error) {
	m.BACnetErrorContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetErrorGeneral"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetErrorGeneral")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	error, err := ReadSimpleField[Error](ctx, "error", ReadComplex[Error](ErrorParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'error' field"))
	}
	m.Error = error

	if closeErr := readBuffer.CloseContext("BACnetErrorGeneral"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetErrorGeneral")
	}

	return m, nil
}

func (m *_BACnetErrorGeneral) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetErrorGeneral) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorGeneral"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetErrorGeneral")
		}

		if err := WriteSimpleField[Error](ctx, "error", m.GetError(), WriteComplex[Error](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("BACnetErrorGeneral"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetErrorGeneral")
		}
		return nil
	}
	return m.BACnetErrorContract.(*_BACnetError).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetErrorGeneral) IsBACnetErrorGeneral() {}

func (m *_BACnetErrorGeneral) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
