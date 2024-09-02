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

// ModbusPDUReadFileRecordRequest is the corresponding interface of ModbusPDUReadFileRecordRequest
type ModbusPDUReadFileRecordRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetItems returns Items (property field)
	GetItems() []ModbusPDUReadFileRecordRequestItem
}

// ModbusPDUReadFileRecordRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadFileRecordRequest.
// This is useful for switch cases.
type ModbusPDUReadFileRecordRequestExactly interface {
	ModbusPDUReadFileRecordRequest
	isModbusPDUReadFileRecordRequest() bool
}

// _ModbusPDUReadFileRecordRequest is the data-structure of this message
type _ModbusPDUReadFileRecordRequest struct {
	ModbusPDUContract
	Items []ModbusPDUReadFileRecordRequestItem
}

var _ ModbusPDUReadFileRecordRequest = (*_ModbusPDUReadFileRecordRequest)(nil)
var _ ModbusPDURequirements = (*_ModbusPDUReadFileRecordRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadFileRecordRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadFileRecordRequest) GetFunctionFlag() uint8 {
	return 0x14
}

func (m *_ModbusPDUReadFileRecordRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadFileRecordRequest) GetParent() ModbusPDUContract {
	return m.ModbusPDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadFileRecordRequest) GetItems() []ModbusPDUReadFileRecordRequestItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordRequest factory function for _ModbusPDUReadFileRecordRequest
func NewModbusPDUReadFileRecordRequest(items []ModbusPDUReadFileRecordRequestItem) *_ModbusPDUReadFileRecordRequest {
	_result := &_ModbusPDUReadFileRecordRequest{
		ModbusPDUContract: NewModbusPDU(),
		Items:             items,
	}
	_result.ModbusPDUContract.(*_ModbusPDU)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadFileRecordRequest(structType any) ModbusPDUReadFileRecordRequest {
	if casted, ok := structType.(ModbusPDUReadFileRecordRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadFileRecordRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadFileRecordRequest) GetTypeName() string {
	return "ModbusPDUReadFileRecordRequest"
}

func (m *_ModbusPDUReadFileRecordRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ModbusPDUContract.(*_ModbusPDU).getLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUReadFileRecordRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ModbusPDUReadFileRecordRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ModbusPDU, response bool) (__modbusPDUReadFileRecordRequest ModbusPDUReadFileRecordRequest, err error) {
	m.ModbusPDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadFileRecordRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	byteCount, err := ReadImplicitField[uint8](ctx, "byteCount", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'byteCount' field"))
	}
	_ = byteCount

	items, err := ReadLengthArrayField[ModbusPDUReadFileRecordRequestItem](ctx, "items", ReadComplex[ModbusPDUReadFileRecordRequestItem](ModbusPDUReadFileRecordRequestItemParseWithBuffer, readBuffer), int(byteCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadFileRecordRequest")
	}

	return m, nil
}

func (m *_ModbusPDUReadFileRecordRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadFileRecordRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	itemsArraySizeInBytes := func(items []ModbusPDUReadFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes(ctx))
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadFileRecordRequest")
		}
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		if err := WriteImplicitField(ctx, "byteCount", byteCount, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'byteCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadFileRecordRequest")
		}
		return nil
	}
	return m.ModbusPDUContract.(*_ModbusPDU).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadFileRecordRequest) isModbusPDUReadFileRecordRequest() bool {
	return true
}

func (m *_ModbusPDUReadFileRecordRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
