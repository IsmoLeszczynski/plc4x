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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLogDataLogDataEntryEnumeratedValue is the corresponding interface of BACnetLogDataLogDataEntryEnumeratedValue
type BACnetLogDataLogDataEntryEnumeratedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetLogDataLogDataEntry
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetContextTagEnumerated
}

// BACnetLogDataLogDataEntryEnumeratedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogDataEntryEnumeratedValue.
// This is useful for switch cases.
type BACnetLogDataLogDataEntryEnumeratedValueExactly interface {
	BACnetLogDataLogDataEntryEnumeratedValue
	isBACnetLogDataLogDataEntryEnumeratedValue() bool
}

// _BACnetLogDataLogDataEntryEnumeratedValue is the data-structure of this message
type _BACnetLogDataLogDataEntryEnumeratedValue struct {
	*_BACnetLogDataLogDataEntry
	EnumeratedValue BACnetContextTagEnumerated
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) InitializeParent(parent BACnetLogDataLogDataEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetParent() BACnetLogDataLogDataEntry {
	return m._BACnetLogDataLogDataEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetEnumeratedValue() BACnetContextTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogDataLogDataEntryEnumeratedValue factory function for _BACnetLogDataLogDataEntryEnumeratedValue
func NewBACnetLogDataLogDataEntryEnumeratedValue(enumeratedValue BACnetContextTagEnumerated, peekedTagHeader BACnetTagHeader) *_BACnetLogDataLogDataEntryEnumeratedValue {
	_result := &_BACnetLogDataLogDataEntryEnumeratedValue{
		EnumeratedValue:            enumeratedValue,
		_BACnetLogDataLogDataEntry: NewBACnetLogDataLogDataEntry(peekedTagHeader),
	}
	_result._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogDataEntryEnumeratedValue(structType any) BACnetLogDataLogDataEntryEnumeratedValue {
	if casted, ok := structType.(BACnetLogDataLogDataEntryEnumeratedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogDataEntryEnumeratedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetTypeName() string {
	return "BACnetLogDataLogDataEntryEnumeratedValue"
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogDataLogDataEntryEnumeratedValueParse(ctx context.Context, theBytes []byte) (BACnetLogDataLogDataEntryEnumeratedValue, error) {
	return BACnetLogDataLogDataEntryEnumeratedValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogDataLogDataEntryEnumeratedValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogDataLogDataEntryEnumeratedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLogDataLogDataEntryEnumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogDataEntryEnumeratedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enumeratedValue")
	}
	_enumeratedValue, _enumeratedValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_ENUMERATED))
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field of BACnetLogDataLogDataEntryEnumeratedValue")
	}
	enumeratedValue := _enumeratedValue.(BACnetContextTagEnumerated)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enumeratedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogDataEntryEnumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogDataEntryEnumeratedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogDataEntryEnumeratedValue{
		_BACnetLogDataLogDataEntry: &_BACnetLogDataLogDataEntry{},
		EnumeratedValue:            enumeratedValue,
	}
	_child._BACnetLogDataLogDataEntry._BACnetLogDataLogDataEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogDataEntryEnumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogDataEntryEnumeratedValue")
		}

		// Simple Field (enumeratedValue)
		if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
		}
		_enumeratedValueErr := writeBuffer.WriteSerializable(ctx, m.GetEnumeratedValue())
		if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for enumeratedValue")
		}
		if _enumeratedValueErr != nil {
			return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogDataEntryEnumeratedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogDataEntryEnumeratedValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) isBACnetLogDataLogDataEntryEnumeratedValue() bool {
	return true
}

func (m *_BACnetLogDataLogDataEntryEnumeratedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
