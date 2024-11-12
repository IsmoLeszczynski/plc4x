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

// FindServersOnNetworkResponse is the corresponding interface of FindServersOnNetworkResponse
type FindServersOnNetworkResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ResponseHeader
	// GetLastCounterResetTime returns LastCounterResetTime (property field)
	GetLastCounterResetTime() int64
	// GetServers returns Servers (property field)
	GetServers() []ServerOnNetwork
	// IsFindServersOnNetworkResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsFindServersOnNetworkResponse()
	// CreateBuilder creates a FindServersOnNetworkResponseBuilder
	CreateFindServersOnNetworkResponseBuilder() FindServersOnNetworkResponseBuilder
}

// _FindServersOnNetworkResponse is the data-structure of this message
type _FindServersOnNetworkResponse struct {
	ExtensionObjectDefinitionContract
	ResponseHeader       ResponseHeader
	LastCounterResetTime int64
	Servers              []ServerOnNetwork
}

var _ FindServersOnNetworkResponse = (*_FindServersOnNetworkResponse)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_FindServersOnNetworkResponse)(nil)

// NewFindServersOnNetworkResponse factory function for _FindServersOnNetworkResponse
func NewFindServersOnNetworkResponse(responseHeader ResponseHeader, lastCounterResetTime int64, servers []ServerOnNetwork) *_FindServersOnNetworkResponse {
	if responseHeader == nil {
		panic("responseHeader of type ResponseHeader for FindServersOnNetworkResponse must not be nil")
	}
	_result := &_FindServersOnNetworkResponse{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ResponseHeader:                    responseHeader,
		LastCounterResetTime:              lastCounterResetTime,
		Servers:                           servers,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// FindServersOnNetworkResponseBuilder is a builder for FindServersOnNetworkResponse
type FindServersOnNetworkResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(responseHeader ResponseHeader, lastCounterResetTime int64, servers []ServerOnNetwork) FindServersOnNetworkResponseBuilder
	// WithResponseHeader adds ResponseHeader (property field)
	WithResponseHeader(ResponseHeader) FindServersOnNetworkResponseBuilder
	// WithResponseHeaderBuilder adds ResponseHeader (property field) which is build by the builder
	WithResponseHeaderBuilder(func(ResponseHeaderBuilder) ResponseHeaderBuilder) FindServersOnNetworkResponseBuilder
	// WithLastCounterResetTime adds LastCounterResetTime (property field)
	WithLastCounterResetTime(int64) FindServersOnNetworkResponseBuilder
	// WithServers adds Servers (property field)
	WithServers(...ServerOnNetwork) FindServersOnNetworkResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the FindServersOnNetworkResponse or returns an error if something is wrong
	Build() (FindServersOnNetworkResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() FindServersOnNetworkResponse
}

// NewFindServersOnNetworkResponseBuilder() creates a FindServersOnNetworkResponseBuilder
func NewFindServersOnNetworkResponseBuilder() FindServersOnNetworkResponseBuilder {
	return &_FindServersOnNetworkResponseBuilder{_FindServersOnNetworkResponse: new(_FindServersOnNetworkResponse)}
}

type _FindServersOnNetworkResponseBuilder struct {
	*_FindServersOnNetworkResponse

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (FindServersOnNetworkResponseBuilder) = (*_FindServersOnNetworkResponseBuilder)(nil)

func (b *_FindServersOnNetworkResponseBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_FindServersOnNetworkResponseBuilder) WithMandatoryFields(responseHeader ResponseHeader, lastCounterResetTime int64, servers []ServerOnNetwork) FindServersOnNetworkResponseBuilder {
	return b.WithResponseHeader(responseHeader).WithLastCounterResetTime(lastCounterResetTime).WithServers(servers...)
}

func (b *_FindServersOnNetworkResponseBuilder) WithResponseHeader(responseHeader ResponseHeader) FindServersOnNetworkResponseBuilder {
	b.ResponseHeader = responseHeader
	return b
}

func (b *_FindServersOnNetworkResponseBuilder) WithResponseHeaderBuilder(builderSupplier func(ResponseHeaderBuilder) ResponseHeaderBuilder) FindServersOnNetworkResponseBuilder {
	builder := builderSupplier(b.ResponseHeader.CreateResponseHeaderBuilder())
	var err error
	b.ResponseHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ResponseHeaderBuilder failed"))
	}
	return b
}

func (b *_FindServersOnNetworkResponseBuilder) WithLastCounterResetTime(lastCounterResetTime int64) FindServersOnNetworkResponseBuilder {
	b.LastCounterResetTime = lastCounterResetTime
	return b
}

func (b *_FindServersOnNetworkResponseBuilder) WithServers(servers ...ServerOnNetwork) FindServersOnNetworkResponseBuilder {
	b.Servers = servers
	return b
}

func (b *_FindServersOnNetworkResponseBuilder) Build() (FindServersOnNetworkResponse, error) {
	if b.ResponseHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'responseHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._FindServersOnNetworkResponse.deepCopy(), nil
}

func (b *_FindServersOnNetworkResponseBuilder) MustBuild() FindServersOnNetworkResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_FindServersOnNetworkResponseBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_FindServersOnNetworkResponseBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_FindServersOnNetworkResponseBuilder) DeepCopy() any {
	_copy := b.CreateFindServersOnNetworkResponseBuilder().(*_FindServersOnNetworkResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateFindServersOnNetworkResponseBuilder creates a FindServersOnNetworkResponseBuilder
func (b *_FindServersOnNetworkResponse) CreateFindServersOnNetworkResponseBuilder() FindServersOnNetworkResponseBuilder {
	if b == nil {
		return NewFindServersOnNetworkResponseBuilder()
	}
	return &_FindServersOnNetworkResponseBuilder{_FindServersOnNetworkResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FindServersOnNetworkResponse) GetExtensionId() int32 {
	return int32(12193)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FindServersOnNetworkResponse) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FindServersOnNetworkResponse) GetResponseHeader() ResponseHeader {
	return m.ResponseHeader
}

func (m *_FindServersOnNetworkResponse) GetLastCounterResetTime() int64 {
	return m.LastCounterResetTime
}

func (m *_FindServersOnNetworkResponse) GetServers() []ServerOnNetwork {
	return m.Servers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastFindServersOnNetworkResponse(structType any) FindServersOnNetworkResponse {
	if casted, ok := structType.(FindServersOnNetworkResponse); ok {
		return casted
	}
	if casted, ok := structType.(*FindServersOnNetworkResponse); ok {
		return *casted
	}
	return nil
}

func (m *_FindServersOnNetworkResponse) GetTypeName() string {
	return "FindServersOnNetworkResponse"
}

func (m *_FindServersOnNetworkResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (lastCounterResetTime)
	lengthInBits += 64

	// Implicit Field (noOfServers)
	lengthInBits += 32

	// Array field
	if len(m.Servers) > 0 {
		for _curItem, element := range m.Servers {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Servers), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_FindServersOnNetworkResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_FindServersOnNetworkResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__findServersOnNetworkResponse FindServersOnNetworkResponse, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FindServersOnNetworkResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FindServersOnNetworkResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	responseHeader, err := ReadSimpleField[ResponseHeader](ctx, "responseHeader", ReadComplex[ResponseHeader](ExtensionObjectDefinitionParseWithBufferProducer[ResponseHeader]((int32)(int32(394))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'responseHeader' field"))
	}
	m.ResponseHeader = responseHeader

	lastCounterResetTime, err := ReadSimpleField(ctx, "lastCounterResetTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastCounterResetTime' field"))
	}
	m.LastCounterResetTime = lastCounterResetTime

	noOfServers, err := ReadImplicitField[int32](ctx, "noOfServers", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServers' field"))
	}
	_ = noOfServers

	servers, err := ReadCountArrayField[ServerOnNetwork](ctx, "servers", ReadComplex[ServerOnNetwork](ExtensionObjectDefinitionParseWithBufferProducer[ServerOnNetwork]((int32)(int32(12191))), readBuffer), uint64(noOfServers))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'servers' field"))
	}
	m.Servers = servers

	if closeErr := readBuffer.CloseContext("FindServersOnNetworkResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FindServersOnNetworkResponse")
	}

	return m, nil
}

func (m *_FindServersOnNetworkResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FindServersOnNetworkResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FindServersOnNetworkResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FindServersOnNetworkResponse")
		}

		if err := WriteSimpleField[ResponseHeader](ctx, "responseHeader", m.GetResponseHeader(), WriteComplex[ResponseHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'responseHeader' field")
		}

		if err := WriteSimpleField[int64](ctx, "lastCounterResetTime", m.GetLastCounterResetTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'lastCounterResetTime' field")
		}
		noOfServers := int32(utils.InlineIf(bool((m.GetServers()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetServers()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfServers", noOfServers, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServers' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "servers", m.GetServers(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'servers' field")
		}

		if popErr := writeBuffer.PopContext("FindServersOnNetworkResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FindServersOnNetworkResponse")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_FindServersOnNetworkResponse) IsFindServersOnNetworkResponse() {}

func (m *_FindServersOnNetworkResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_FindServersOnNetworkResponse) deepCopy() *_FindServersOnNetworkResponse {
	if m == nil {
		return nil
	}
	_FindServersOnNetworkResponseCopy := &_FindServersOnNetworkResponse{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ResponseHeader](m.ResponseHeader),
		m.LastCounterResetTime,
		utils.DeepCopySlice[ServerOnNetwork, ServerOnNetwork](m.Servers),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _FindServersOnNetworkResponseCopy
}

func (m *_FindServersOnNetworkResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
