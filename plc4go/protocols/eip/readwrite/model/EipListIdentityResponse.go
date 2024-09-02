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

// EipListIdentityResponse is the corresponding interface of EipListIdentityResponse
type EipListIdentityResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
	// GetItems returns Items (property field)
	GetItems() []CommandSpecificDataItem
}

// EipListIdentityResponseExactly can be used when we want exactly this type and not a type which fulfills EipListIdentityResponse.
// This is useful for switch cases.
type EipListIdentityResponseExactly interface {
	EipListIdentityResponse
	isEipListIdentityResponse() bool
}

// _EipListIdentityResponse is the data-structure of this message
type _EipListIdentityResponse struct {
	EipPacketContract
	Items []CommandSpecificDataItem
}

var _ EipListIdentityResponse = (*_EipListIdentityResponse)(nil)
var _ EipPacketRequirements = (*_EipListIdentityResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_EipListIdentityResponse) GetCommand() uint16 {
	return 0x0063
}

func (m *_EipListIdentityResponse) GetResponse() bool {
	return bool(true)
}

func (m *_EipListIdentityResponse) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_EipListIdentityResponse) GetParent() EipPacketContract {
	return m.EipPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EipListIdentityResponse) GetItems() []CommandSpecificDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEipListIdentityResponse factory function for _EipListIdentityResponse
func NewEipListIdentityResponse(items []CommandSpecificDataItem, sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_EipListIdentityResponse {
	_result := &_EipListIdentityResponse{
		EipPacketContract: NewEipPacket(sessionHandle, status, senderContext, options),
		Items:             items,
	}
	_result.EipPacketContract.(*_EipPacket)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastEipListIdentityResponse(structType any) EipListIdentityResponse {
	if casted, ok := structType.(EipListIdentityResponse); ok {
		return casted
	}
	if casted, ok := structType.(*EipListIdentityResponse); ok {
		return *casted
	}
	return nil
}

func (m *_EipListIdentityResponse) GetTypeName() string {
	return "EipListIdentityResponse"
}

func (m *_EipListIdentityResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.EipPacketContract.(*_EipPacket).getLengthInBits(ctx))

	// Implicit Field (itemCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_EipListIdentityResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_EipListIdentityResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_EipPacket, response bool) (__eipListIdentityResponse EipListIdentityResponse, err error) {
	m.EipPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipListIdentityResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipListIdentityResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemCount, err := ReadImplicitField[uint16](ctx, "itemCount", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemCount' field"))
	}
	_ = itemCount

	items, err := ReadCountArrayField[CommandSpecificDataItem](ctx, "items", ReadComplex[CommandSpecificDataItem](CommandSpecificDataItemParseWithBuffer, readBuffer), uint64(itemCount))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("EipListIdentityResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipListIdentityResponse")
	}

	return m, nil
}

func (m *_EipListIdentityResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipListIdentityResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("EipListIdentityResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for EipListIdentityResponse")
		}
		itemCount := uint16(uint16(len(m.GetItems())))
		if err := WriteImplicitField(ctx, "itemCount", itemCount, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemCount' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("EipListIdentityResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for EipListIdentityResponse")
		}
		return nil
	}
	return m.EipPacketContract.(*_EipPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_EipListIdentityResponse) isEipListIdentityResponse() bool {
	return true
}

func (m *_EipListIdentityResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
