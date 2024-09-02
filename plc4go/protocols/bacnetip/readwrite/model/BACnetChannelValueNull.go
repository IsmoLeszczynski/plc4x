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

// BACnetChannelValueNull is the corresponding interface of BACnetChannelValueNull
type BACnetChannelValueNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetChannelValueNullExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueNull.
// This is useful for switch cases.
type BACnetChannelValueNullExactly interface {
	BACnetChannelValueNull
	isBACnetChannelValueNull() bool
}

// _BACnetChannelValueNull is the data-structure of this message
type _BACnetChannelValueNull struct {
	BACnetChannelValueContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetChannelValueNull = (*_BACnetChannelValueNull)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueNull)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueNull) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueNull factory function for _BACnetChannelValueNull
func NewBACnetChannelValueNull(nullValue BACnetApplicationTagNull, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueNull {
	_result := &_BACnetChannelValueNull{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		NullValue:                  nullValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueNull(structType any) BACnetChannelValueNull {
	if casted, ok := structType.(BACnetChannelValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueNull) GetTypeName() string {
	return "BACnetChannelValueNull"
}

func (m *_BACnetChannelValueNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueNull BACnetChannelValueNull, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueNull")
	}

	return m, nil
}

func (m *_BACnetChannelValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueNull")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueNull) isBACnetChannelValueNull() bool {
	return true
}

func (m *_BACnetChannelValueNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
