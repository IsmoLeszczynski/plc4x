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

// BACnetScaleIntegerScale is the corresponding interface of BACnetScaleIntegerScale
type BACnetScaleIntegerScale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetScale
	// GetIntegerScale returns IntegerScale (property field)
	GetIntegerScale() BACnetContextTagSignedInteger
}

// BACnetScaleIntegerScaleExactly can be used when we want exactly this type and not a type which fulfills BACnetScaleIntegerScale.
// This is useful for switch cases.
type BACnetScaleIntegerScaleExactly interface {
	BACnetScaleIntegerScale
	isBACnetScaleIntegerScale() bool
}

// _BACnetScaleIntegerScale is the data-structure of this message
type _BACnetScaleIntegerScale struct {
	BACnetScaleContract
	IntegerScale BACnetContextTagSignedInteger
}

var _ BACnetScaleIntegerScale = (*_BACnetScaleIntegerScale)(nil)
var _ BACnetScaleRequirements = (*_BACnetScaleIntegerScale)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetScaleIntegerScale) GetParent() BACnetScaleContract {
	return m.BACnetScaleContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScaleIntegerScale) GetIntegerScale() BACnetContextTagSignedInteger {
	return m.IntegerScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetScaleIntegerScale factory function for _BACnetScaleIntegerScale
func NewBACnetScaleIntegerScale(integerScale BACnetContextTagSignedInteger, peekedTagHeader BACnetTagHeader) *_BACnetScaleIntegerScale {
	_result := &_BACnetScaleIntegerScale{
		BACnetScaleContract: NewBACnetScale(peekedTagHeader),
		IntegerScale:        integerScale,
	}
	_result.BACnetScaleContract.(*_BACnetScale)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetScaleIntegerScale(structType any) BACnetScaleIntegerScale {
	if casted, ok := structType.(BACnetScaleIntegerScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScaleIntegerScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScaleIntegerScale) GetTypeName() string {
	return "BACnetScaleIntegerScale"
}

func (m *_BACnetScaleIntegerScale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetScaleContract.(*_BACnetScale).getLengthInBits(ctx))

	// Simple field (integerScale)
	lengthInBits += m.IntegerScale.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetScaleIntegerScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetScaleIntegerScale) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetScale) (__bACnetScaleIntegerScale BACnetScaleIntegerScale, err error) {
	m.BACnetScaleContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleIntegerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleIntegerScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerScale, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerScale", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerScale' field"))
	}
	m.IntegerScale = integerScale

	if closeErr := readBuffer.CloseContext("BACnetScaleIntegerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleIntegerScale")
	}

	return m, nil
}

func (m *_BACnetScaleIntegerScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetScaleIntegerScale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleIntegerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleIntegerScale")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerScale", m.GetIntegerScale(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerScale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetScaleIntegerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleIntegerScale")
		}
		return nil
	}
	return m.BACnetScaleContract.(*_BACnetScale).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetScaleIntegerScale) isBACnetScaleIntegerScale() bool {
	return true
}

func (m *_BACnetScaleIntegerScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
