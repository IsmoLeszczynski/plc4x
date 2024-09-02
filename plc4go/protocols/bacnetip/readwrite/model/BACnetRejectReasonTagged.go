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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetRejectReasonTagged is the corresponding interface of BACnetRejectReasonTagged
type BACnetRejectReasonTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (property field)
	GetValue() BACnetRejectReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// _BACnetRejectReasonTagged is the data-structure of this message
type _BACnetRejectReasonTagged struct {
	Value            BACnetRejectReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

var _ BACnetRejectReasonTagged = (*_BACnetRejectReasonTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRejectReasonTagged) GetValue() BACnetRejectReason {
	return m.Value
}

func (m *_BACnetRejectReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetRejectReasonTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetRejectReasonTagged factory function for _BACnetRejectReasonTagged
func NewBACnetRejectReasonTagged(value BACnetRejectReason, proprietaryValue uint32, actualLength uint32) *_BACnetRejectReasonTagged {
	return &_BACnetRejectReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetRejectReasonTagged(structType any) BACnetRejectReasonTagged {
	if casted, ok := structType.(BACnetRejectReasonTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRejectReasonTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRejectReasonTagged) GetTypeName() string {
	return "BACnetRejectReasonTagged"
}

func (m *_BACnetRejectReasonTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32(int32(0)) }, func() any { return int32((int32(m.GetActualLength()) * int32(int32(8)))) }).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any { return int32((int32(m.GetActualLength()) * int32(int32(8)))) }, func() any { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *_BACnetRejectReasonTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRejectReasonTaggedParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetRejectReasonTagged, error) {
	return BACnetRejectReasonTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetRejectReasonTaggedParseWithBufferProducer(actualLength uint32) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRejectReasonTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRejectReasonTagged, error) {
		return BACnetRejectReasonTaggedParseWithBuffer(ctx, readBuffer, actualLength)
	}
}

func BACnetRejectReasonTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetRejectReasonTagged, error) {
	v, err := (&_BACnetRejectReasonTagged{ActualLength: actualLength}).parse(ctx, readBuffer, actualLength)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetRejectReasonTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (__bACnetRejectReasonTagged BACnetRejectReasonTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetRejectReasonTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRejectReasonTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	value, err := ReadManualField[BACnetRejectReason](ctx, "value", readBuffer, EnsureType[BACnetRejectReason](ReadEnumGeneric(ctx, readBuffer, actualLength, BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	isProprietary, err := ReadVirtualField[bool](ctx, "isProprietary", (*bool)(nil), bool((value) == (BACnetRejectReason_VENDOR_PROPRIETARY_VALUE)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isProprietary' field"))
	}
	_ = isProprietary

	proprietaryValue, err := ReadManualField[uint32](ctx, "proprietaryValue", readBuffer, EnsureType[uint32](ReadProprietaryEnumGeneric(ctx, readBuffer, actualLength, isProprietary)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'proprietaryValue' field"))
	}
	m.ProprietaryValue = proprietaryValue

	if closeErr := readBuffer.CloseContext("BACnetRejectReasonTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRejectReasonTagged")
	}

	return m, nil
}

func (m *_BACnetRejectReasonTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRejectReasonTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetRejectReasonTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRejectReasonTagged")
	}

	if err := WriteManualField[BACnetRejectReason](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	if err := WriteManualField[uint32](ctx, "proprietaryValue", func(ctx context.Context) error {
		return WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	}, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRejectReasonTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRejectReasonTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetRejectReasonTagged) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetRejectReasonTagged) IsBACnetRejectReasonTagged() {}

func (m *_BACnetRejectReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
