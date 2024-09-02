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

// BACnetEventTransitionBitsTagged is the corresponding interface of BACnetEventTransitionBitsTagged
type BACnetEventTransitionBitsTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetToOffnormal returns ToOffnormal (virtual field)
	GetToOffnormal() bool
	// GetToFault returns ToFault (virtual field)
	GetToFault() bool
	// GetToNormal returns ToNormal (virtual field)
	GetToNormal() bool
}

// _BACnetEventTransitionBitsTagged is the data-structure of this message
type _BACnetEventTransitionBitsTagged struct {
	Header  BACnetTagHeader
	Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetEventTransitionBitsTagged = (*_BACnetEventTransitionBitsTagged)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTransitionBitsTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetEventTransitionBitsTagged) GetPayload() BACnetTagPayloadBitString {
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

func (m *_BACnetEventTransitionBitsTagged) GetToOffnormal() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (0))), func() any { return bool(m.GetPayload().GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetEventTransitionBitsTagged) GetToFault() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (1))), func() any { return bool(m.GetPayload().GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
}

func (m *_BACnetEventTransitionBitsTagged) GetToNormal() bool {
	ctx := context.Background()
	_ = ctx
	return bool(utils.InlineIf((bool((len(m.GetPayload().GetData())) > (2))), func() any { return bool(m.GetPayload().GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventTransitionBitsTagged factory function for _BACnetEventTransitionBitsTagged
func NewBACnetEventTransitionBitsTagged(header BACnetTagHeader, payload BACnetTagPayloadBitString, tagNumber uint8, tagClass TagClass) *_BACnetEventTransitionBitsTagged {
	return &_BACnetEventTransitionBitsTagged{Header: header, Payload: payload, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventTransitionBitsTagged(structType any) BACnetEventTransitionBitsTagged {
	if casted, ok := structType.(BACnetEventTransitionBitsTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTransitionBitsTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTransitionBitsTagged) GetTypeName() string {
	return "BACnetEventTransitionBitsTagged"
}

func (m *_BACnetEventTransitionBitsTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetEventTransitionBitsTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventTransitionBitsTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetEventTransitionBitsTagged, error) {
	return BACnetEventTransitionBitsTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetEventTransitionBitsTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTransitionBitsTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventTransitionBitsTagged, error) {
		return BACnetEventTransitionBitsTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetEventTransitionBitsTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetEventTransitionBitsTagged, error) {
	v, err := (&_BACnetEventTransitionBitsTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetEventTransitionBitsTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetEventTransitionBitsTagged BACnetEventTransitionBitsTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTransitionBitsTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTransitionBitsTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	payload, err := ReadSimpleField[BACnetTagPayloadBitString](ctx, "payload", ReadComplex[BACnetTagPayloadBitString](BACnetTagPayloadBitStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	toOffnormal, err := ReadVirtualField[bool](ctx, "toOffnormal", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (0))), func() any { return bool(payload.GetData()[0]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toOffnormal' field"))
	}
	_ = toOffnormal

	toFault, err := ReadVirtualField[bool](ctx, "toFault", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (1))), func() any { return bool(payload.GetData()[1]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toFault' field"))
	}
	_ = toFault

	toNormal, err := ReadVirtualField[bool](ctx, "toNormal", (*bool)(nil), utils.InlineIf((bool((len(payload.GetData())) > (2))), func() any { return bool(payload.GetData()[2]) }, func() any { return bool(bool(false)) }).(bool))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'toNormal' field"))
	}
	_ = toNormal

	if closeErr := readBuffer.CloseContext("BACnetEventTransitionBitsTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTransitionBitsTagged")
	}

	return m, nil
}

func (m *_BACnetEventTransitionBitsTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTransitionBitsTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetEventTransitionBitsTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTransitionBitsTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteSimpleField[BACnetTagPayloadBitString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadBitString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'payload' field")
	}
	// Virtual field
	toOffnormal := m.GetToOffnormal()
	_ = toOffnormal
	if _toOffnormalErr := writeBuffer.WriteVirtual(ctx, "toOffnormal", m.GetToOffnormal()); _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}
	// Virtual field
	toFault := m.GetToFault()
	_ = toFault
	if _toFaultErr := writeBuffer.WriteVirtual(ctx, "toFault", m.GetToFault()); _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}
	// Virtual field
	toNormal := m.GetToNormal()
	_ = toNormal
	if _toNormalErr := writeBuffer.WriteVirtual(ctx, "toNormal", m.GetToNormal()); _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTransitionBitsTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTransitionBitsTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventTransitionBitsTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetEventTransitionBitsTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetEventTransitionBitsTagged) IsBACnetEventTransitionBitsTagged() {}

func (m *_BACnetEventTransitionBitsTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
