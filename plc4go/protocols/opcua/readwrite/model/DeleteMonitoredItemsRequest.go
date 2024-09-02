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

// DeleteMonitoredItemsRequest is the corresponding interface of DeleteMonitoredItemsRequest
type DeleteMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetNoOfMonitoredItemIds returns NoOfMonitoredItemIds (property field)
	GetNoOfMonitoredItemIds() int32
	// GetMonitoredItemIds returns MonitoredItemIds (property field)
	GetMonitoredItemIds() []uint32
}

// DeleteMonitoredItemsRequestExactly can be used when we want exactly this type and not a type which fulfills DeleteMonitoredItemsRequest.
// This is useful for switch cases.
type DeleteMonitoredItemsRequestExactly interface {
	DeleteMonitoredItemsRequest
	isDeleteMonitoredItemsRequest() bool
}

// _DeleteMonitoredItemsRequest is the data-structure of this message
type _DeleteMonitoredItemsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader        ExtensionObjectDefinition
	SubscriptionId       uint32
	NoOfMonitoredItemIds int32
	MonitoredItemIds     []uint32
}

var _ DeleteMonitoredItemsRequest = (*_DeleteMonitoredItemsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteMonitoredItemsRequest)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetIdentifier() string {
	return "781"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteMonitoredItemsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteMonitoredItemsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_DeleteMonitoredItemsRequest) GetNoOfMonitoredItemIds() int32 {
	return m.NoOfMonitoredItemIds
}

func (m *_DeleteMonitoredItemsRequest) GetMonitoredItemIds() []uint32 {
	return m.MonitoredItemIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDeleteMonitoredItemsRequest factory function for _DeleteMonitoredItemsRequest
func NewDeleteMonitoredItemsRequest(requestHeader ExtensionObjectDefinition, subscriptionId uint32, noOfMonitoredItemIds int32, monitoredItemIds []uint32) *_DeleteMonitoredItemsRequest {
	_result := &_DeleteMonitoredItemsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		NoOfMonitoredItemIds:              noOfMonitoredItemIds,
		MonitoredItemIds:                  monitoredItemIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDeleteMonitoredItemsRequest(structType any) DeleteMonitoredItemsRequest {
	if casted, ok := structType.(DeleteMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteMonitoredItemsRequest) GetTypeName() string {
	return "DeleteMonitoredItemsRequest"
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (noOfMonitoredItemIds)
	lengthInBits += 32

	// Array field
	if len(m.MonitoredItemIds) > 0 {
		lengthInBits += 32 * uint16(len(m.MonitoredItemIds))
	}

	return lengthInBits
}

func (m *_DeleteMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteMonitoredItemsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteMonitoredItemsRequest DeleteMonitoredItemsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	noOfMonitoredItemIds, err := ReadSimpleField(ctx, "noOfMonitoredItemIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfMonitoredItemIds' field"))
	}
	m.NoOfMonitoredItemIds = noOfMonitoredItemIds

	monitoredItemIds, err := ReadCountArrayField[uint32](ctx, "monitoredItemIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfMonitoredItemIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'monitoredItemIds' field"))
	}
	m.MonitoredItemIds = monitoredItemIds

	if closeErr := readBuffer.CloseContext("DeleteMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteMonitoredItemsRequest")
	}

	return m, nil
}

func (m *_DeleteMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteMonitoredItemsRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfMonitoredItemIds", m.GetNoOfMonitoredItemIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfMonitoredItemIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "monitoredItemIds", m.GetMonitoredItemIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'monitoredItemIds' field")
		}

		if popErr := writeBuffer.PopContext("DeleteMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteMonitoredItemsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteMonitoredItemsRequest) isDeleteMonitoredItemsRequest() bool {
	return true
}

func (m *_DeleteMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
