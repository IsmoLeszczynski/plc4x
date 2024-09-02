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

// UnregisterNodesRequest is the corresponding interface of UnregisterNodesRequest
type UnregisterNodesRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfNodesToUnregister returns NoOfNodesToUnregister (property field)
	GetNoOfNodesToUnregister() int32
	// GetNodesToUnregister returns NodesToUnregister (property field)
	GetNodesToUnregister() []NodeId
}

// UnregisterNodesRequestExactly can be used when we want exactly this type and not a type which fulfills UnregisterNodesRequest.
// This is useful for switch cases.
type UnregisterNodesRequestExactly interface {
	UnregisterNodesRequest
	isUnregisterNodesRequest() bool
}

// _UnregisterNodesRequest is the data-structure of this message
type _UnregisterNodesRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader         ExtensionObjectDefinition
	NoOfNodesToUnregister int32
	NodesToUnregister     []NodeId
}

var _ UnregisterNodesRequest = (*_UnregisterNodesRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UnregisterNodesRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UnregisterNodesRequest) GetIdentifier() string {
	return "566"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UnregisterNodesRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UnregisterNodesRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_UnregisterNodesRequest) GetNoOfNodesToUnregister() int32 {
	return m.NoOfNodesToUnregister
}

func (m *_UnregisterNodesRequest) GetNodesToUnregister() []NodeId {
	return m.NodesToUnregister
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUnregisterNodesRequest factory function for _UnregisterNodesRequest
func NewUnregisterNodesRequest(requestHeader ExtensionObjectDefinition, noOfNodesToUnregister int32, nodesToUnregister []NodeId) *_UnregisterNodesRequest {
	_result := &_UnregisterNodesRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfNodesToUnregister:             noOfNodesToUnregister,
		NodesToUnregister:                 nodesToUnregister,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUnregisterNodesRequest(structType any) UnregisterNodesRequest {
	if casted, ok := structType.(UnregisterNodesRequest); ok {
		return casted
	}
	if casted, ok := structType.(*UnregisterNodesRequest); ok {
		return *casted
	}
	return nil
}

func (m *_UnregisterNodesRequest) GetTypeName() string {
	return "UnregisterNodesRequest"
}

func (m *_UnregisterNodesRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfNodesToUnregister)
	lengthInBits += 32

	// Array field
	if len(m.NodesToUnregister) > 0 {
		for _curItem, element := range m.NodesToUnregister {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodesToUnregister), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_UnregisterNodesRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UnregisterNodesRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__unregisterNodesRequest UnregisterNodesRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UnregisterNodesRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UnregisterNodesRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfNodesToUnregister, err := ReadSimpleField(ctx, "noOfNodesToUnregister", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNodesToUnregister' field"))
	}
	m.NoOfNodesToUnregister = noOfNodesToUnregister

	nodesToUnregister, err := ReadCountArrayField[NodeId](ctx, "nodesToUnregister", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer), uint64(noOfNodesToUnregister))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodesToUnregister' field"))
	}
	m.NodesToUnregister = nodesToUnregister

	if closeErr := readBuffer.CloseContext("UnregisterNodesRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UnregisterNodesRequest")
	}

	return m, nil
}

func (m *_UnregisterNodesRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UnregisterNodesRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UnregisterNodesRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UnregisterNodesRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNodesToUnregister", m.GetNoOfNodesToUnregister(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNodesToUnregister' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "nodesToUnregister", m.GetNodesToUnregister(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'nodesToUnregister' field")
		}

		if popErr := writeBuffer.PopContext("UnregisterNodesRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UnregisterNodesRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UnregisterNodesRequest) isUnregisterNodesRequest() bool {
	return true
}

func (m *_UnregisterNodesRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
