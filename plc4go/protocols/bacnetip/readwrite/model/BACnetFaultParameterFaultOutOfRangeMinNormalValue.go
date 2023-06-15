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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetFaultParameterFaultOutOfRangeMinNormalValue is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValue
type BACnetFaultParameterFaultOutOfRangeMinNormalValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValue.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	isBACnetFaultParameterFaultOutOfRangeMinNormalValue() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValue is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValue struct {
	_BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements
	OpeningTag      BACnetOpeningTag
	PeekedTagHeader BACnetTagHeader
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetFaultParameterFaultOutOfRangeMinNormalValueParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMinNormalValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetFaultParameterFaultOutOfRangeMinNormalValueChild interface {
	utils.Serializable
	InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMinNormalValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag)
	GetParent() *BACnetFaultParameterFaultOutOfRangeMinNormalValue

	GetTypeName() string
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetClosingTag() BACnetClosingTag {
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

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRangeMinNormalValue factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValue
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValue(structType any) BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValue"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValue, error) {
	return BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Validation
	if !(bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS))) {
		return nil, errors.WithStack(utils.ParseValidationError{"only application tags allowed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement interface {
		BACnetFaultParameterFaultOutOfRangeMinNormalValue
		InitializeParent(BACnetFaultParameterFaultOutOfRangeMinNormalValue, BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValue
	}
	var _childTemp any
	var _child BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x4: // BACnetFaultParameterFaultOutOfRangeMinNormalValueReal
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueRealParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x2: // BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsigned
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueUnsignedParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x5: // BACnetFaultParameterFaultOutOfRangeMinNormalValueDouble
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueDoubleParseWithBuffer(ctx, readBuffer, tagNumber)
	case peekedTagNumber == 0x3: // BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParseWithBuffer(ctx, readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	_child = _childTemp.(BACnetFaultParameterFaultOutOfRangeMinNormalValueChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	// Finish initializing
	_child.InitializeParent(_child, openingTag, peekedTagHeader, closingTag)
	return _child, nil
}

func (pm *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMinNormalValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValue")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) isBACnetFaultParameterFaultOutOfRangeMinNormalValue() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
