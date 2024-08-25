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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// FindServersOnNetworkRequest is the corresponding interface of FindServersOnNetworkRequest
type FindServersOnNetworkRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetStartingRecordId returns StartingRecordId (property field)
	GetStartingRecordId() uint32
	// GetMaxRecordsToReturn returns MaxRecordsToReturn (property field)
	GetMaxRecordsToReturn() uint32
	// GetNoOfServerCapabilityFilter returns NoOfServerCapabilityFilter (property field)
	GetNoOfServerCapabilityFilter() int32
	// GetServerCapabilityFilter returns ServerCapabilityFilter (property field)
	GetServerCapabilityFilter() []PascalString
}

// FindServersOnNetworkRequestExactly can be used when we want exactly this type and not a type which fulfills FindServersOnNetworkRequest.
// This is useful for switch cases.
type FindServersOnNetworkRequestExactly interface {
	FindServersOnNetworkRequest
	isFindServersOnNetworkRequest() bool
}

// _FindServersOnNetworkRequest is the data-structure of this message
type _FindServersOnNetworkRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader              ExtensionObjectDefinition
	StartingRecordId           uint32
	MaxRecordsToReturn         uint32
	NoOfServerCapabilityFilter int32
	ServerCapabilityFilter     []PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetIdentifier() string {
	return "12192"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersOnNetworkRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_FindServersOnNetworkRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersOnNetworkRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_FindServersOnNetworkRequest) GetStartingRecordId() uint32 {
	return m.StartingRecordId
}

func (m *_FindServersOnNetworkRequest) GetMaxRecordsToReturn() uint32 {
	return m.MaxRecordsToReturn
}

func (m *_FindServersOnNetworkRequest) GetNoOfServerCapabilityFilter() int32 {
	return m.NoOfServerCapabilityFilter
}

func (m *_FindServersOnNetworkRequest) GetServerCapabilityFilter() []PascalString {
	return m.ServerCapabilityFilter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewFindServersOnNetworkRequest factory function for _FindServersOnNetworkRequest
func NewFindServersOnNetworkRequest(requestHeader ExtensionObjectDefinition, startingRecordId uint32, maxRecordsToReturn uint32, noOfServerCapabilityFilter int32, serverCapabilityFilter []PascalString) *_FindServersOnNetworkRequest {
	_result := &_FindServersOnNetworkRequest{
		RequestHeader:              requestHeader,
		StartingRecordId:           startingRecordId,
		MaxRecordsToReturn:         maxRecordsToReturn,
		NoOfServerCapabilityFilter: noOfServerCapabilityFilter,
		ServerCapabilityFilter:     serverCapabilityFilter,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFindServersOnNetworkRequest(structType any) FindServersOnNetworkRequest {
	if casted, ok := structType.(FindServersOnNetworkRequest); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersOnNetworkRequest); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersOnNetworkRequest) GetTypeName() string {
	return "FindServersOnNetworkRequest"
}

func (m *_FindServersOnNetworkRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (startingRecordId)
	lengthInBits += 32

	// Simple field (maxRecordsToReturn)
	lengthInBits += 32

	// Simple field (noOfServerCapabilityFilter)
	lengthInBits += 32

	// Array field
	if len(m.ServerCapabilityFilter) > 0 {
		for _curItem, element := range m.ServerCapabilityFilter {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerCapabilityFilter), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersOnNetworkRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FindServersOnNetworkRequestParse(ctx context.Context, theBytes []byte, identifier string) (FindServersOnNetworkRequest, error) {
	return FindServersOnNetworkRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func FindServersOnNetworkRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (FindServersOnNetworkRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("FindServersOnNetworkRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersOnNetworkRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of FindServersOnNetworkRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (startingRecordId)
	_startingRecordId, _startingRecordIdErr := readBuffer.ReadUint32("startingRecordId", 32)
	if _startingRecordIdErr != nil {
		return nil, errors.Wrap(_startingRecordIdErr, "Error parsing 'startingRecordId' field of FindServersOnNetworkRequest")
	}
	startingRecordId := _startingRecordId

	// Simple Field (maxRecordsToReturn)
	_maxRecordsToReturn, _maxRecordsToReturnErr := readBuffer.ReadUint32("maxRecordsToReturn", 32)
	if _maxRecordsToReturnErr != nil {
		return nil, errors.Wrap(_maxRecordsToReturnErr, "Error parsing 'maxRecordsToReturn' field of FindServersOnNetworkRequest")
	}
	maxRecordsToReturn := _maxRecordsToReturn

	// Simple Field (noOfServerCapabilityFilter)
	_noOfServerCapabilityFilter, _noOfServerCapabilityFilterErr := readBuffer.ReadInt32("noOfServerCapabilityFilter", 32)
	if _noOfServerCapabilityFilterErr != nil {
		return nil, errors.Wrap(_noOfServerCapabilityFilterErr, "Error parsing 'noOfServerCapabilityFilter' field of FindServersOnNetworkRequest")
	}
	noOfServerCapabilityFilter := _noOfServerCapabilityFilter

	// Array field (serverCapabilityFilter)
	if pullErr := readBuffer.PullContext("serverCapabilityFilter", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serverCapabilityFilter")
	}
	// Count array
	serverCapabilityFilter := make([]PascalString, max(noOfServerCapabilityFilter, 0))
	// This happens when the size is set conditional to 0
	if len(serverCapabilityFilter) == 0 {
		serverCapabilityFilter = nil
	}
	{
		_numItems := uint16(max(noOfServerCapabilityFilter, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'serverCapabilityFilter' field of FindServersOnNetworkRequest")
			}
			serverCapabilityFilter[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("serverCapabilityFilter", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serverCapabilityFilter")
	}

	if closeErr := readBuffer.CloseContext("FindServersOnNetworkRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersOnNetworkRequest")
	}

	// Create a partially initialized instance
	_child := &_FindServersOnNetworkRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		StartingRecordId:           startingRecordId,
		MaxRecordsToReturn:         maxRecordsToReturn,
		NoOfServerCapabilityFilter: noOfServerCapabilityFilter,
		ServerCapabilityFilter:     serverCapabilityFilter,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_FindServersOnNetworkRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersOnNetworkRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersOnNetworkRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersOnNetworkRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (startingRecordId)
		startingRecordId := uint32(m.GetStartingRecordId())
		_startingRecordIdErr := writeBuffer.WriteUint32("startingRecordId", 32, uint32((startingRecordId)))
		if _startingRecordIdErr != nil {
			return errors.Wrap(_startingRecordIdErr, "Error serializing 'startingRecordId' field")
		}

		// Simple Field (maxRecordsToReturn)
		maxRecordsToReturn := uint32(m.GetMaxRecordsToReturn())
		_maxRecordsToReturnErr := writeBuffer.WriteUint32("maxRecordsToReturn", 32, uint32((maxRecordsToReturn)))
		if _maxRecordsToReturnErr != nil {
			return errors.Wrap(_maxRecordsToReturnErr, "Error serializing 'maxRecordsToReturn' field")
		}

		// Simple Field (noOfServerCapabilityFilter)
		noOfServerCapabilityFilter := int32(m.GetNoOfServerCapabilityFilter())
		_noOfServerCapabilityFilterErr := writeBuffer.WriteInt32("noOfServerCapabilityFilter", 32, int32((noOfServerCapabilityFilter)))
		if _noOfServerCapabilityFilterErr != nil {
			return errors.Wrap(_noOfServerCapabilityFilterErr, "Error serializing 'noOfServerCapabilityFilter' field")
		}

		// Array Field (serverCapabilityFilter)
		if pushErr := writeBuffer.PushContext("serverCapabilityFilter", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serverCapabilityFilter")
		}
		for _curItem, _element := range m.GetServerCapabilityFilter() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetServerCapabilityFilter()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'serverCapabilityFilter' field")
			}
		}
		if popErr := writeBuffer.PopContext("serverCapabilityFilter", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serverCapabilityFilter")
		}

		if popErr := writeBuffer.PopContext("FindServersOnNetworkRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersOnNetworkRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersOnNetworkRequest) isFindServersOnNetworkRequest() bool {
	return true
}

func (m *_FindServersOnNetworkRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
