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

// CloseSessionRequest is the corresponding interface of CloseSessionRequest
type CloseSessionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetDeleteSubscriptions returns DeleteSubscriptions (property field)
	GetDeleteSubscriptions() bool
}

// _CloseSessionRequest is the data-structure of this message
type _CloseSessionRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       ExtensionObjectDefinition
	DeleteSubscriptions bool
	// Reserved Fields
	reservedField0 *uint8
}

var _ CloseSessionRequest = (*_CloseSessionRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSessionRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSessionRequest) GetIdentifier() string {
	return "473"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSessionRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSessionRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_CloseSessionRequest) GetDeleteSubscriptions() bool {
	return m.DeleteSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCloseSessionRequest factory function for _CloseSessionRequest
func NewCloseSessionRequest(requestHeader ExtensionObjectDefinition, deleteSubscriptions bool) *_CloseSessionRequest {
	_result := &_CloseSessionRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		DeleteSubscriptions:               deleteSubscriptions,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCloseSessionRequest(structType any) CloseSessionRequest {
	if casted, ok := structType.(CloseSessionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSessionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSessionRequest) GetTypeName() string {
	return "CloseSessionRequest"
}

func (m *_CloseSessionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteSubscriptions)
	lengthInBits += 1

	return lengthInBits
}

func (m *_CloseSessionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSessionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__closeSessionRequest CloseSessionRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSessionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSessionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	deleteSubscriptions, err := ReadSimpleField(ctx, "deleteSubscriptions", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deleteSubscriptions' field"))
	}
	m.DeleteSubscriptions = deleteSubscriptions

	if closeErr := readBuffer.CloseContext("CloseSessionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSessionRequest")
	}

	return m, nil
}

func (m *_CloseSessionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSessionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSessionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSessionRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "deleteSubscriptions", m.GetDeleteSubscriptions(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deleteSubscriptions' field")
		}

		if popErr := writeBuffer.PopContext("CloseSessionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSessionRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSessionRequest) IsCloseSessionRequest() {}

func (m *_CloseSessionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
