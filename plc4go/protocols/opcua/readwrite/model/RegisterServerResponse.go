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

// RegisterServerResponse is the corresponding interface of RegisterServerResponse
type RegisterServerResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
}

// RegisterServerResponseExactly can be used when we want exactly this type and not a type which fulfills RegisterServerResponse.
// This is useful for switch cases.
type RegisterServerResponseExactly interface {
	RegisterServerResponse
	isRegisterServerResponse() bool
}

// _RegisterServerResponse is the data-structure of this message
type _RegisterServerResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader ExtensionObjectDefinition
}

var _ RegisterServerResponse = (*_RegisterServerResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RegisterServerResponse)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RegisterServerResponse) GetIdentifier() string {
	return "440"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RegisterServerResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RegisterServerResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRegisterServerResponse factory function for _RegisterServerResponse
func NewRegisterServerResponse(responseHeader ExtensionObjectDefinition) *_RegisterServerResponse {
	_result := &_RegisterServerResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRegisterServerResponse(structType any) RegisterServerResponse {
	if casted, ok := structType.(RegisterServerResponse); ok {
		return casted
	}
	if casted, ok := structType.(*RegisterServerResponse); ok {
		return *casted
	}
	return nil
}

func (m *_RegisterServerResponse) GetTypeName() string {
	return "RegisterServerResponse"
}

func (m *_RegisterServerResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RegisterServerResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RegisterServerResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__registerServerResponse RegisterServerResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RegisterServerResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RegisterServerResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("394")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	if closeErr := readBuffer.CloseContext("RegisterServerResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RegisterServerResponse")
	}

	return m, nil
}

func (m *_RegisterServerResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RegisterServerResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RegisterServerResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RegisterServerResponse")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if popErr := writeBuffer.PopContext("RegisterServerResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RegisterServerResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RegisterServerResponse) isRegisterServerResponse() bool {
	return true
}

func (m *_RegisterServerResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
