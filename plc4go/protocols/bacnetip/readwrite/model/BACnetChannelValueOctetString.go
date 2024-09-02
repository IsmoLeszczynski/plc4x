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

// BACnetChannelValueOctetString is the corresponding interface of BACnetChannelValueOctetString
type BACnetChannelValueOctetString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() BACnetApplicationTagOctetString
}

// _BACnetChannelValueOctetString is the data-structure of this message
type _BACnetChannelValueOctetString struct {
	BACnetChannelValueContract
	OctetStringValue BACnetApplicationTagOctetString
}

var _ BACnetChannelValueOctetString = (*_BACnetChannelValueOctetString)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueOctetString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueOctetString) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueOctetString) GetOctetStringValue() BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetChannelValueOctetString factory function for _BACnetChannelValueOctetString
func NewBACnetChannelValueOctetString(octetStringValue BACnetApplicationTagOctetString, peekedTagHeader BACnetTagHeader) *_BACnetChannelValueOctetString {
	_result := &_BACnetChannelValueOctetString{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		OctetStringValue:           octetStringValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueOctetString(structType any) BACnetChannelValueOctetString {
	if casted, ok := structType.(BACnetChannelValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueOctetString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueOctetString) GetTypeName() string {
	return "BACnetChannelValueOctetString"
}

func (m *_BACnetChannelValueOctetString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueOctetString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueOctetString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueOctetString BACnetChannelValueOctetString, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	octetStringValue, err := ReadSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", ReadComplex[BACnetApplicationTagOctetString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagOctetString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'octetStringValue' field"))
	}
	m.OctetStringValue = octetStringValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueOctetString")
	}

	return m, nil
}

func (m *_BACnetChannelValueOctetString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueOctetString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueOctetString")
		}

		if err := WriteSimpleField[BACnetApplicationTagOctetString](ctx, "octetStringValue", m.GetOctetStringValue(), WriteComplex[BACnetApplicationTagOctetString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueOctetString")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueOctetString) IsBACnetChannelValueOctetString() {}

func (m *_BACnetChannelValueOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
