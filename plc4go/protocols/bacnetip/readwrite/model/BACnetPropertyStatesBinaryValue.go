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

// BACnetPropertyStatesBinaryValue is the corresponding interface of BACnetPropertyStatesBinaryValue
type BACnetPropertyStatesBinaryValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBinaryValue returns BinaryValue (property field)
	GetBinaryValue() BACnetBinaryPVTagged
}

// BACnetPropertyStatesBinaryValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesBinaryValue.
// This is useful for switch cases.
type BACnetPropertyStatesBinaryValueExactly interface {
	BACnetPropertyStatesBinaryValue
	isBACnetPropertyStatesBinaryValue() bool
}

// _BACnetPropertyStatesBinaryValue is the data-structure of this message
type _BACnetPropertyStatesBinaryValue struct {
	BACnetPropertyStatesContract
	BinaryValue BACnetBinaryPVTagged
}

var _ BACnetPropertyStatesBinaryValue = (*_BACnetPropertyStatesBinaryValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesBinaryValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesBinaryValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBinaryValue) GetBinaryValue() BACnetBinaryPVTagged {
	return m.BinaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesBinaryValue factory function for _BACnetPropertyStatesBinaryValue
func NewBACnetPropertyStatesBinaryValue(binaryValue BACnetBinaryPVTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesBinaryValue {
	_result := &_BACnetPropertyStatesBinaryValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		BinaryValue:                  binaryValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBinaryValue(structType any) BACnetPropertyStatesBinaryValue {
	if casted, ok := structType.(BACnetPropertyStatesBinaryValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBinaryValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryValue"
}

func (m *_BACnetPropertyStatesBinaryValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (binaryValue)
	lengthInBits += m.BinaryValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesBinaryValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesBinaryValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesBinaryValue BACnetPropertyStatesBinaryValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBinaryValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	binaryValue, err := ReadSimpleField[BACnetBinaryPVTagged](ctx, "binaryValue", ReadComplex[BACnetBinaryPVTagged](BACnetBinaryPVTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'binaryValue' field"))
	}
	m.BinaryValue = binaryValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBinaryValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesBinaryValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBinaryValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBinaryValue")
		}

		if err := WriteSimpleField[BACnetBinaryPVTagged](ctx, "binaryValue", m.GetBinaryValue(), WriteComplex[BACnetBinaryPVTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'binaryValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBinaryValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesBinaryValue) isBACnetPropertyStatesBinaryValue() bool {
	return true
}

func (m *_BACnetPropertyStatesBinaryValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
