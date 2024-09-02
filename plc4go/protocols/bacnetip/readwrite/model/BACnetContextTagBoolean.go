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

// BACnetContextTagBoolean is the corresponding interface of BACnetContextTagBoolean
type BACnetContextTagBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetValue returns Value (property field)
	GetValue() uint8
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() bool
}

// BACnetContextTagBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagBoolean.
// This is useful for switch cases.
type BACnetContextTagBooleanExactly interface {
	BACnetContextTagBoolean
	isBACnetContextTagBoolean() bool
}

// _BACnetContextTagBoolean is the data-structure of this message
type _BACnetContextTagBoolean struct {
	BACnetContextTagContract
	Value   uint8
	Payload BACnetTagPayloadBoolean
}

var _ BACnetContextTagBoolean = (*_BACnetContextTagBoolean)(nil)
var _ BACnetContextTagRequirements = (*_BACnetContextTagBoolean)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagBoolean) GetDataType() BACnetDataType {
	return BACnetDataType_BOOLEAN
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagBoolean) GetParent() BACnetContextTagContract {
	return m.BACnetContextTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagBoolean) GetValue() uint8 {
	return m.Value
}

func (m *_BACnetContextTagBoolean) GetPayload() BACnetTagPayloadBoolean {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagBoolean) GetActualValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagBoolean factory function for _BACnetContextTagBoolean
func NewBACnetContextTagBoolean(value uint8, payload BACnetTagPayloadBoolean, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagBoolean {
	_result := &_BACnetContextTagBoolean{
		BACnetContextTagContract: NewBACnetContextTag(header, tagNumberArgument),
		Value:                    value,
		Payload:                  payload,
	}
	_result.BACnetContextTagContract.(*_BACnetContextTag)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagBoolean(structType any) BACnetContextTagBoolean {
	if casted, ok := structType.(BACnetContextTagBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagBoolean) GetTypeName() string {
	return "BACnetContextTagBoolean"
}

func (m *_BACnetContextTagBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetContextTagContract.(*_BACnetContextTag).getLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 8

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetContextTagBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetContextTagBoolean) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetContextTag, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (__bACnetContextTagBoolean BACnetContextTagBoolean, err error) {
	m.BACnetContextTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((header.GetActualLength()) == (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "length field should be 1"})
	}

	value, err := ReadSimpleField(ctx, "value", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	payload, err := ReadSimpleField[BACnetTagPayloadBoolean](ctx, "payload", ReadComplex[BACnetTagPayloadBoolean](BACnetTagPayloadBooleanParseWithBufferProducer((uint32)(value)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[bool](ctx, "actualValue", (*bool)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagBoolean")
	}

	return m, nil
}

func (m *_BACnetContextTagBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagBoolean")
		}

		if err := WriteSimpleField[uint8](ctx, "value", m.GetValue(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'value' field")
		}

		if err := WriteSimpleField[BACnetTagPayloadBoolean](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagBoolean")
		}
		return nil
	}
	return m.BACnetContextTagContract.(*_BACnetContextTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagBoolean) isBACnetContextTagBoolean() bool {
	return true
}

func (m *_BACnetContextTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
