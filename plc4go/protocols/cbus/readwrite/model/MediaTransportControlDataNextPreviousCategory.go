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

// MediaTransportControlDataNextPreviousCategory is the corresponding interface of MediaTransportControlDataNextPreviousCategory
type MediaTransportControlDataNextPreviousCategory interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetOperation returns Operation (property field)
	GetOperation() byte
	// GetIsSetThePreviousCategory returns IsSetThePreviousCategory (virtual field)
	GetIsSetThePreviousCategory() bool
	// GetIsSetTheNextCategory returns IsSetTheNextCategory (virtual field)
	GetIsSetTheNextCategory() bool
}

// _MediaTransportControlDataNextPreviousCategory is the data-structure of this message
type _MediaTransportControlDataNextPreviousCategory struct {
	MediaTransportControlDataContract
	Operation byte
}

var _ MediaTransportControlDataNextPreviousCategory = (*_MediaTransportControlDataNextPreviousCategory)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataNextPreviousCategory)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataNextPreviousCategory) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousCategory) GetOperation() byte {
	return m.Operation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataNextPreviousCategory) GetIsSetThePreviousCategory() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) == (0x00)))
}

func (m *_MediaTransportControlDataNextPreviousCategory) GetIsSetTheNextCategory() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetOperation()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataNextPreviousCategory factory function for _MediaTransportControlDataNextPreviousCategory
func NewMediaTransportControlDataNextPreviousCategory(operation byte, commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte) *_MediaTransportControlDataNextPreviousCategory {
	_result := &_MediaTransportControlDataNextPreviousCategory{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		Operation:                         operation,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataNextPreviousCategory(structType any) MediaTransportControlDataNextPreviousCategory {
	if casted, ok := structType.(MediaTransportControlDataNextPreviousCategory); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataNextPreviousCategory); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataNextPreviousCategory) GetTypeName() string {
	return "MediaTransportControlDataNextPreviousCategory"
}

func (m *_MediaTransportControlDataNextPreviousCategory) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (operation)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_MediaTransportControlDataNextPreviousCategory) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataNextPreviousCategory) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataNextPreviousCategory MediaTransportControlDataNextPreviousCategory, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataNextPreviousCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataNextPreviousCategory")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	operation, err := ReadSimpleField(ctx, "operation", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operation' field"))
	}
	m.Operation = operation

	isSetThePreviousCategory, err := ReadVirtualField[bool](ctx, "isSetThePreviousCategory", (*bool)(nil), bool((operation) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetThePreviousCategory' field"))
	}
	_ = isSetThePreviousCategory

	isSetTheNextCategory, err := ReadVirtualField[bool](ctx, "isSetTheNextCategory", (*bool)(nil), bool((operation) != (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isSetTheNextCategory' field"))
	}
	_ = isSetTheNextCategory

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataNextPreviousCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataNextPreviousCategory")
	}

	return m, nil
}

func (m *_MediaTransportControlDataNextPreviousCategory) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataNextPreviousCategory) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataNextPreviousCategory"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataNextPreviousCategory")
		}

		if err := WriteSimpleField[byte](ctx, "operation", m.GetOperation(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'operation' field")
		}
		// Virtual field
		isSetThePreviousCategory := m.GetIsSetThePreviousCategory()
		_ = isSetThePreviousCategory
		if _isSetThePreviousCategoryErr := writeBuffer.WriteVirtual(ctx, "isSetThePreviousCategory", m.GetIsSetThePreviousCategory()); _isSetThePreviousCategoryErr != nil {
			return errors.Wrap(_isSetThePreviousCategoryErr, "Error serializing 'isSetThePreviousCategory' field")
		}
		// Virtual field
		isSetTheNextCategory := m.GetIsSetTheNextCategory()
		_ = isSetTheNextCategory
		if _isSetTheNextCategoryErr := writeBuffer.WriteVirtual(ctx, "isSetTheNextCategory", m.GetIsSetTheNextCategory()); _isSetTheNextCategoryErr != nil {
			return errors.Wrap(_isSetTheNextCategoryErr, "Error serializing 'isSetTheNextCategory' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataNextPreviousCategory"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataNextPreviousCategory")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataNextPreviousCategory) IsMediaTransportControlDataNextPreviousCategory() {
}

func (m *_MediaTransportControlDataNextPreviousCategory) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
