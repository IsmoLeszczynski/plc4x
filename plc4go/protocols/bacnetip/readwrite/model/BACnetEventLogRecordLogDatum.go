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

// BACnetEventLogRecordLogDatum is the corresponding interface of BACnetEventLogRecordLogDatum
type BACnetEventLogRecordLogDatum interface {
	BACnetEventLogRecordLogDatumContract
	BACnetEventLogRecordLogDatumRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BACnetEventLogRecordLogDatumContract provides a set of functions which can be overwritten by a sub struct
type BACnetEventLogRecordLogDatumContract interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetTagNumber() returns a parser argument
	GetTagNumber() uint8
	// IsBACnetEventLogRecordLogDatum() is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecordLogDatum()
}

// BACnetEventLogRecordLogDatumRequirements provides a set of functions which need to be implemented by a sub struct
type BACnetEventLogRecordLogDatumRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetPeekedTagNumber returns PeekedTagNumber (discriminator field)
	GetPeekedTagNumber() uint8
}

// _BACnetEventLogRecordLogDatum is the data-structure of this message
type _BACnetEventLogRecordLogDatum struct {
	_SubType        BACnetEventLogRecordLogDatum
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetEventLogRecordLogDatumContract = (*_BACnetEventLogRecordLogDatum)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatum) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventLogRecordLogDatum) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetEventLogRecordLogDatum) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_BACnetEventLogRecordLogDatum) GetPeekedTagNumber() uint8 {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatum factory function for _BACnetEventLogRecordLogDatum
func NewBACnetEventLogRecordLogDatum(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventLogRecordLogDatum {
	return &_BACnetEventLogRecordLogDatum{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatum(structType any) BACnetEventLogRecordLogDatum {
	if casted, ok := structType.(BACnetEventLogRecordLogDatum); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatum); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatum) GetTypeName() string {
	return "BACnetEventLogRecordLogDatum"
}

func (m *_BACnetEventLogRecordLogDatum) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatum) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func BACnetEventLogRecordLogDatumParse[T BACnetEventLogRecordLogDatum](ctx context.Context, theBytes []byte, tagNumber uint8) (T, error) {
	return BACnetEventLogRecordLogDatumParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventLogRecordLogDatumParseWithBufferProducer[T BACnetEventLogRecordLogDatum](tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := BACnetEventLogRecordLogDatumParseWithBuffer[T](ctx, readBuffer, tagNumber)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, err
	}
}

func BACnetEventLogRecordLogDatumParseWithBuffer[T BACnetEventLogRecordLogDatum](ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (T, error) {
	v, err := (&_BACnetEventLogRecordLogDatum{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		var zero T
		return zero, err
	}
	return v.(T), err
}

func (m *_BACnetEventLogRecordLogDatum) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetEventLogRecordLogDatum BACnetEventLogRecordLogDatum, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatum")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	peekedTagHeader, err := ReadPeekField[BACnetTagHeader](ctx, "peekedTagHeader", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagHeader' field"))
	}
	m.PeekedTagHeader = peekedTagHeader

	peekedTagNumber, err := ReadVirtualField[uint8](ctx, "peekedTagNumber", (*uint8)(nil), peekedTagHeader.GetActualTagNumber())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedTagNumber' field"))
	}
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child BACnetEventLogRecordLogDatum
	switch {
	case peekedTagNumber == uint8(0): // BACnetEventLogRecordLogDatumLogStatus
		if _child, err = (&_BACnetEventLogRecordLogDatumLogStatus{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumLogStatus for type-switch of BACnetEventLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(1): // BACnetEventLogRecordLogDatumNotification
		if _child, err = (&_BACnetEventLogRecordLogDatumNotification{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumNotification for type-switch of BACnetEventLogRecordLogDatum")
		}
	case peekedTagNumber == uint8(2): // BACnetEventLogRecordLogDatumTimeChange
		if _child, err = (&_BACnetEventLogRecordLogDatumTimeChange{}).parse(ctx, readBuffer, m, tagNumber); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type BACnetEventLogRecordLogDatumTimeChange for type-switch of BACnetEventLogRecordLogDatum")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatum")
	}

	return _child, nil
}

func (pm *_BACnetEventLogRecordLogDatum) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventLogRecordLogDatum, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatum")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatum")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventLogRecordLogDatum) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventLogRecordLogDatum) IsBACnetEventLogRecordLogDatum() {}
