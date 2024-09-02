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

// BACnetSpecialEventListOfTimeValues is the corresponding interface of BACnetSpecialEventListOfTimeValues
type BACnetSpecialEventListOfTimeValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfTimeValues returns ListOfTimeValues (property field)
	GetListOfTimeValues() []BACnetTimeValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// _BACnetSpecialEventListOfTimeValues is the data-structure of this message
type _BACnetSpecialEventListOfTimeValues struct {
	OpeningTag       BACnetOpeningTag
	ListOfTimeValues []BACnetTimeValue
	ClosingTag       BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetSpecialEventListOfTimeValues = (*_BACnetSpecialEventListOfTimeValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSpecialEventListOfTimeValues) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetSpecialEventListOfTimeValues) GetListOfTimeValues() []BACnetTimeValue {
	return m.ListOfTimeValues
}

func (m *_BACnetSpecialEventListOfTimeValues) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSpecialEventListOfTimeValues factory function for _BACnetSpecialEventListOfTimeValues
func NewBACnetSpecialEventListOfTimeValues(openingTag BACnetOpeningTag, listOfTimeValues []BACnetTimeValue, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetSpecialEventListOfTimeValues {
	return &_BACnetSpecialEventListOfTimeValues{OpeningTag: openingTag, ListOfTimeValues: listOfTimeValues, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetSpecialEventListOfTimeValues(structType any) BACnetSpecialEventListOfTimeValues {
	if casted, ok := structType.(BACnetSpecialEventListOfTimeValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSpecialEventListOfTimeValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSpecialEventListOfTimeValues) GetTypeName() string {
	return "BACnetSpecialEventListOfTimeValues"
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.ListOfTimeValues) > 0 {
		for _, element := range m.ListOfTimeValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSpecialEventListOfTimeValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSpecialEventListOfTimeValuesParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	return BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetSpecialEventListOfTimeValuesParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventListOfTimeValues, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetSpecialEventListOfTimeValues, error) {
		return BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetSpecialEventListOfTimeValuesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetSpecialEventListOfTimeValues, error) {
	v, err := (&_BACnetSpecialEventListOfTimeValues{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetSpecialEventListOfTimeValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetSpecialEventListOfTimeValues BACnetSpecialEventListOfTimeValues, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetSpecialEventListOfTimeValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSpecialEventListOfTimeValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	listOfTimeValues, err := ReadTerminatedArrayField[BACnetTimeValue](ctx, "listOfTimeValues", ReadComplex[BACnetTimeValue](BACnetTimeValueParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfTimeValues' field"))
	}
	m.ListOfTimeValues = listOfTimeValues

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetSpecialEventListOfTimeValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSpecialEventListOfTimeValues")
	}

	return m, nil
}

func (m *_BACnetSpecialEventListOfTimeValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSpecialEventListOfTimeValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSpecialEventListOfTimeValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSpecialEventListOfTimeValues")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteComplexTypeArrayField(ctx, "listOfTimeValues", m.GetListOfTimeValues(), writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'listOfTimeValues' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetSpecialEventListOfTimeValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSpecialEventListOfTimeValues")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSpecialEventListOfTimeValues) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetSpecialEventListOfTimeValues) IsBACnetSpecialEventListOfTimeValues() {}

func (m *_BACnetSpecialEventListOfTimeValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
