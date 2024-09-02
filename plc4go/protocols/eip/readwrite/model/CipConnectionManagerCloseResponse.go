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

// CipConnectionManagerCloseResponse is the corresponding interface of CipConnectionManagerCloseResponse
type CipConnectionManagerCloseResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetAdditionalStatusWords returns AdditionalStatusWords (property field)
	GetAdditionalStatusWords() uint8
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetApplicationReplySize returns ApplicationReplySize (property field)
	GetApplicationReplySize() uint8
}

// CipConnectionManagerCloseResponseExactly can be used when we want exactly this type and not a type which fulfills CipConnectionManagerCloseResponse.
// This is useful for switch cases.
type CipConnectionManagerCloseResponseExactly interface {
	CipConnectionManagerCloseResponse
	isCipConnectionManagerCloseResponse() bool
}

// _CipConnectionManagerCloseResponse is the data-structure of this message
type _CipConnectionManagerCloseResponse struct {
	CipServiceContract
	Status                 uint8
	AdditionalStatusWords  uint8
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	ApplicationReplySize   uint8
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ CipConnectionManagerCloseResponse = (*_CipConnectionManagerCloseResponse)(nil)
var _ CipServiceRequirements = (*_CipConnectionManagerCloseResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerCloseResponse) GetService() uint8 {
	return 0x4E
}

func (m *_CipConnectionManagerCloseResponse) GetResponse() bool {
	return bool(true)
}

func (m *_CipConnectionManagerCloseResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerCloseResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerCloseResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_CipConnectionManagerCloseResponse) GetAdditionalStatusWords() uint8 {
	return m.AdditionalStatusWords
}

func (m *_CipConnectionManagerCloseResponse) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerCloseResponse) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerCloseResponse) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerCloseResponse) GetApplicationReplySize() uint8 {
	return m.ApplicationReplySize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipConnectionManagerCloseResponse factory function for _CipConnectionManagerCloseResponse
func NewCipConnectionManagerCloseResponse(status uint8, additionalStatusWords uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, applicationReplySize uint8, serviceLen uint16) *_CipConnectionManagerCloseResponse {
	_result := &_CipConnectionManagerCloseResponse{
		CipServiceContract:     NewCipService(serviceLen),
		Status:                 status,
		AdditionalStatusWords:  additionalStatusWords,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ApplicationReplySize:   applicationReplySize,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerCloseResponse(structType any) CipConnectionManagerCloseResponse {
	if casted, ok := structType.(CipConnectionManagerCloseResponse); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerCloseResponse); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerCloseResponse) GetTypeName() string {
	return "CipConnectionManagerCloseResponse"
}

func (m *_CipConnectionManagerCloseResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (additionalStatusWords)
	lengthInBits += 8

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (applicationReplySize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipConnectionManagerCloseResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectionManagerCloseResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectionManagerCloseResponse CipConnectionManagerCloseResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerCloseResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerCloseResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	additionalStatusWords, err := ReadSimpleField(ctx, "additionalStatusWords", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalStatusWords' field"))
	}
	m.AdditionalStatusWords = additionalStatusWords

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}
	m.ConnectionSerialNumber = connectionSerialNumber

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}
	m.OriginatorVendorId = originatorVendorId

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}
	m.OriginatorSerialNumber = originatorSerialNumber

	applicationReplySize, err := ReadSimpleField(ctx, "applicationReplySize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'applicationReplySize' field"))
	}
	m.ApplicationReplySize = applicationReplySize

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(8)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	if closeErr := readBuffer.CloseContext("CipConnectionManagerCloseResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerCloseResponse")
	}

	return m, nil
}

func (m *_CipConnectionManagerCloseResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerCloseResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerCloseResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerCloseResponse")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[uint8](ctx, "status", m.GetStatus(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint8](ctx, "additionalStatusWords", m.GetAdditionalStatusWords(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalStatusWords' field")
		}

		if err := WriteSimpleField[uint16](ctx, "connectionSerialNumber", m.GetConnectionSerialNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionSerialNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "originatorVendorId", m.GetOriginatorVendorId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorVendorId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originatorSerialNumber", m.GetOriginatorSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorSerialNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "applicationReplySize", m.GetApplicationReplySize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'applicationReplySize' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerCloseResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerCloseResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerCloseResponse) isCipConnectionManagerCloseResponse() bool {
	return true
}

func (m *_CipConnectionManagerCloseResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
