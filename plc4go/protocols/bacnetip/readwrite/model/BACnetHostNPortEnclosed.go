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

// BACnetHostNPortEnclosed is the corresponding interface of BACnetHostNPortEnclosed
type BACnetHostNPortEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetBacnetHostNPort returns BacnetHostNPort (property field)
	GetBacnetHostNPort() BACnetHostNPort
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// _BACnetHostNPortEnclosed is the data-structure of this message
type _BACnetHostNPortEnclosed struct {
	OpeningTag      BACnetOpeningTag
	BacnetHostNPort BACnetHostNPort
	ClosingTag      BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

var _ BACnetHostNPortEnclosed = (*_BACnetHostNPortEnclosed)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostNPortEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetHostNPortEnclosed) GetBacnetHostNPort() BACnetHostNPort {
	return m.BacnetHostNPort
}

func (m *_BACnetHostNPortEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostNPortEnclosed factory function for _BACnetHostNPortEnclosed
func NewBACnetHostNPortEnclosed(openingTag BACnetOpeningTag, bacnetHostNPort BACnetHostNPort, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetHostNPortEnclosed {
	return &_BACnetHostNPortEnclosed{OpeningTag: openingTag, BacnetHostNPort: bacnetHostNPort, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetHostNPortEnclosed(structType any) BACnetHostNPortEnclosed {
	if casted, ok := structType.(BACnetHostNPortEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostNPortEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostNPortEnclosed) GetTypeName() string {
	return "BACnetHostNPortEnclosed"
}

func (m *_BACnetHostNPortEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (bacnetHostNPort)
	lengthInBits += m.BacnetHostNPort.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostNPortEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetHostNPortEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetHostNPortEnclosed, error) {
	return BACnetHostNPortEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetHostNPortEnclosedParseWithBufferProducer(tagNumber uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPortEnclosed, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetHostNPortEnclosed, error) {
		return BACnetHostNPortEnclosedParseWithBuffer(ctx, readBuffer, tagNumber)
	}
}

func BACnetHostNPortEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetHostNPortEnclosed, error) {
	v, err := (&_BACnetHostNPortEnclosed{TagNumber: tagNumber}).parse(ctx, readBuffer, tagNumber)
	if err != nil {
		return nil, err
	}
	return v, err
}

func (m *_BACnetHostNPortEnclosed) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (__bACnetHostNPortEnclosed BACnetHostNPortEnclosed, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostNPortEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostNPortEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	bacnetHostNPort, err := ReadSimpleField[BACnetHostNPort](ctx, "bacnetHostNPort", ReadComplex[BACnetHostNPort](BACnetHostNPortParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bacnetHostNPort' field"))
	}
	m.BacnetHostNPort = bacnetHostNPort

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(tagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetHostNPortEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostNPortEnclosed")
	}

	return m, nil
}

func (m *_BACnetHostNPortEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostNPortEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetHostNPortEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostNPortEnclosed")
	}

	if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'openingTag' field")
	}

	if err := WriteSimpleField[BACnetHostNPort](ctx, "bacnetHostNPort", m.GetBacnetHostNPort(), WriteComplex[BACnetHostNPort](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'bacnetHostNPort' field")
	}

	if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostNPortEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostNPortEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetHostNPortEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetHostNPortEnclosed) IsBACnetHostNPortEnclosed() {}

func (m *_BACnetHostNPortEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
