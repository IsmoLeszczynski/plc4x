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

// DeleteReferencesRequest is the corresponding interface of DeleteReferencesRequest
type DeleteReferencesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfReferencesToDelete returns NoOfReferencesToDelete (property field)
	GetNoOfReferencesToDelete() int32
	// GetReferencesToDelete returns ReferencesToDelete (property field)
	GetReferencesToDelete() []ExtensionObjectDefinition
}

// _DeleteReferencesRequest is the data-structure of this message
type _DeleteReferencesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader          ExtensionObjectDefinition
	NoOfReferencesToDelete int32
	ReferencesToDelete     []ExtensionObjectDefinition
}

var _ DeleteReferencesRequest = (*_DeleteReferencesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteReferencesRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesRequest) GetIdentifier() string {
	return "506"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteReferencesRequest) GetNoOfReferencesToDelete() int32 {
	return m.NoOfReferencesToDelete
}

func (m *_DeleteReferencesRequest) GetReferencesToDelete() []ExtensionObjectDefinition {
	return m.ReferencesToDelete
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteReferencesRequest factory function for _DeleteReferencesRequest
func NewDeleteReferencesRequest(requestHeader ExtensionObjectDefinition, noOfReferencesToDelete int32, referencesToDelete []ExtensionObjectDefinition) *_DeleteReferencesRequest {
	_result := &_DeleteReferencesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfReferencesToDelete:            noOfReferencesToDelete,
		ReferencesToDelete:                referencesToDelete,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteReferencesRequest(structType any) DeleteReferencesRequest {
	if casted, ok := structType.(DeleteReferencesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesRequest) GetTypeName() string {
	return "DeleteReferencesRequest"
}

func (m *_DeleteReferencesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfReferencesToDelete)
	lengthInBits += 32

	// Array field
	if len(m.ReferencesToDelete) > 0 {
		for _curItem, element := range m.ReferencesToDelete {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencesToDelete), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_DeleteReferencesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteReferencesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteReferencesRequest DeleteReferencesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteReferencesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfReferencesToDelete, err := ReadSimpleField(ctx, "noOfReferencesToDelete", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReferencesToDelete' field"))
	}
	m.NoOfReferencesToDelete = noOfReferencesToDelete

	referencesToDelete, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "referencesToDelete", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("387")), readBuffer), uint64(noOfReferencesToDelete))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referencesToDelete' field"))
	}
	m.ReferencesToDelete = referencesToDelete

	if closeErr := readBuffer.CloseContext("DeleteReferencesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesRequest")
	}

	return m, nil
}

func (m *_DeleteReferencesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfReferencesToDelete", m.GetNoOfReferencesToDelete(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReferencesToDelete' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "referencesToDelete", m.GetReferencesToDelete(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'referencesToDelete' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesRequest) IsDeleteReferencesRequest() {}

func (m *_DeleteReferencesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
