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

// BACnetConfirmedServiceRequestReadRangeRangeByPosition is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeByPosition
type BACnetConfirmedServiceRequestReadRangeRangeByPosition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceIndex returns ReferenceIndex (property field)
	GetReferenceIndex() BACnetApplicationTagUnsignedInteger
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
}

// BACnetConfirmedServiceRequestReadRangeRangeByPositionExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadRangeRangeByPosition.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadRangeRangeByPositionExactly interface {
	BACnetConfirmedServiceRequestReadRangeRangeByPosition
	isBACnetConfirmedServiceRequestReadRangeRangeByPosition() bool
}

// _BACnetConfirmedServiceRequestReadRangeRangeByPosition is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeByPosition struct {
	BACnetConfirmedServiceRequestReadRangeRangeContract
	ReferenceIndex BACnetApplicationTagUnsignedInteger
	Count          BACnetApplicationTagSignedInteger
}

var _ BACnetConfirmedServiceRequestReadRangeRangeByPosition = (*_BACnetConfirmedServiceRequestReadRangeRangeByPosition)(nil)
var _ BACnetConfirmedServiceRequestReadRangeRangeRequirements = (*_BACnetConfirmedServiceRequestReadRangeRangeByPosition)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetParent() BACnetConfirmedServiceRequestReadRangeRangeContract {
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetReferenceIndex() BACnetApplicationTagUnsignedInteger {
	return m.ReferenceIndex
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRangeByPosition factory function for _BACnetConfirmedServiceRequestReadRangeRangeByPosition
func NewBACnetConfirmedServiceRequestReadRangeRangeByPosition(referenceIndex BACnetApplicationTagUnsignedInteger, count BACnetApplicationTagSignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestReadRangeRangeByPosition {
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeByPosition{
		BACnetConfirmedServiceRequestReadRangeRangeContract: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
		ReferenceIndex: referenceIndex,
		Count:          count,
	}
	_result.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeByPosition(structType any) BACnetConfirmedServiceRequestReadRangeRangeByPosition {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeByPosition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeByPosition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeByPosition"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).getLengthInBits(ctx))

	// Simple field (referenceIndex)
	lengthInBits += m.ReferenceIndex.GetLengthInBits(ctx)

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestReadRangeRange) (__bACnetConfirmedServiceRequestReadRangeRangeByPosition BACnetConfirmedServiceRequestReadRangeRangeByPosition, err error) {
	m.BACnetConfirmedServiceRequestReadRangeRangeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceIndex, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "referenceIndex", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceIndex' field"))
	}
	m.ReferenceIndex = referenceIndex

	count, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "referenceIndex", m.GetReferenceIndex(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceIndex' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", m.GetCount(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) isBACnetConfirmedServiceRequestReadRangeRangeByPosition() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
