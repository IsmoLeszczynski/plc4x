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

// ServerOnNetwork is the corresponding interface of ServerOnNetwork
type ServerOnNetwork interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRecordId returns RecordId (property field)
	GetRecordId() uint32
	// GetServerName returns ServerName (property field)
	GetServerName() PascalString
	// GetDiscoveryUrl returns DiscoveryUrl (property field)
	GetDiscoveryUrl() PascalString
	// GetNoOfServerCapabilities returns NoOfServerCapabilities (property field)
	GetNoOfServerCapabilities() int32
	// GetServerCapabilities returns ServerCapabilities (property field)
	GetServerCapabilities() []PascalString
	// IsServerOnNetwork is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServerOnNetwork()
}

// _ServerOnNetwork is the data-structure of this message
type _ServerOnNetwork struct {
	ExtensionObjectDefinitionContract
	RecordId               uint32
	ServerName             PascalString
	DiscoveryUrl           PascalString
	NoOfServerCapabilities int32
	ServerCapabilities     []PascalString
}

var _ ServerOnNetwork = (*_ServerOnNetwork)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ServerOnNetwork)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ServerOnNetwork) GetIdentifier() string {
	return "12191"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerOnNetwork) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ServerOnNetwork) GetRecordId() uint32 {
	return m.RecordId
}

func (m *_ServerOnNetwork) GetServerName() PascalString {
	return m.ServerName
}

func (m *_ServerOnNetwork) GetDiscoveryUrl() PascalString {
	return m.DiscoveryUrl
}

func (m *_ServerOnNetwork) GetNoOfServerCapabilities() int32 {
	return m.NoOfServerCapabilities
}

func (m *_ServerOnNetwork) GetServerCapabilities() []PascalString {
	return m.ServerCapabilities
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewServerOnNetwork factory function for _ServerOnNetwork
func NewServerOnNetwork(recordId uint32, serverName PascalString, discoveryUrl PascalString, noOfServerCapabilities int32, serverCapabilities []PascalString) *_ServerOnNetwork {
	if serverName == nil {
		panic("serverName of type PascalString for ServerOnNetwork must not be nil")
	}
	if discoveryUrl == nil {
		panic("discoveryUrl of type PascalString for ServerOnNetwork must not be nil")
	}
	_result := &_ServerOnNetwork{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RecordId:                          recordId,
		ServerName:                        serverName,
		DiscoveryUrl:                      discoveryUrl,
		NoOfServerCapabilities:            noOfServerCapabilities,
		ServerCapabilities:                serverCapabilities,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastServerOnNetwork(structType any) ServerOnNetwork {
	if casted, ok := structType.(ServerOnNetwork); ok {
		return casted
	}
	if casted, ok := structType.(*ServerOnNetwork); ok {
		return *casted
	}
	return nil
}

func (m *_ServerOnNetwork) GetTypeName() string {
	return "ServerOnNetwork"
}

func (m *_ServerOnNetwork) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (recordId)
	lengthInBits += 32

	// Simple field (serverName)
	lengthInBits += m.ServerName.GetLengthInBits(ctx)

	// Simple field (discoveryUrl)
	lengthInBits += m.DiscoveryUrl.GetLengthInBits(ctx)

	// Simple field (noOfServerCapabilities)
	lengthInBits += 32

	// Array field
	if len(m.ServerCapabilities) > 0 {
		for _curItem, element := range m.ServerCapabilities {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ServerCapabilities), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ServerOnNetwork) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServerOnNetwork) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__serverOnNetwork ServerOnNetwork, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServerOnNetwork"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerOnNetwork")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	recordId, err := ReadSimpleField(ctx, "recordId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'recordId' field"))
	}
	m.RecordId = recordId

	serverName, err := ReadSimpleField[PascalString](ctx, "serverName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverName' field"))
	}
	m.ServerName = serverName

	discoveryUrl, err := ReadSimpleField[PascalString](ctx, "discoveryUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'discoveryUrl' field"))
	}
	m.DiscoveryUrl = discoveryUrl

	noOfServerCapabilities, err := ReadSimpleField(ctx, "noOfServerCapabilities", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfServerCapabilities' field"))
	}
	m.NoOfServerCapabilities = noOfServerCapabilities

	serverCapabilities, err := ReadCountArrayField[PascalString](ctx, "serverCapabilities", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfServerCapabilities))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverCapabilities' field"))
	}
	m.ServerCapabilities = serverCapabilities

	if closeErr := readBuffer.CloseContext("ServerOnNetwork"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerOnNetwork")
	}

	return m, nil
}

func (m *_ServerOnNetwork) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerOnNetwork) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerOnNetwork"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerOnNetwork")
		}

		if err := WriteSimpleField[uint32](ctx, "recordId", m.GetRecordId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'recordId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "serverName", m.GetServerName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverName' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "discoveryUrl", m.GetDiscoveryUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'discoveryUrl' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfServerCapabilities", m.GetNoOfServerCapabilities(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfServerCapabilities' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "serverCapabilities", m.GetServerCapabilities(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'serverCapabilities' field")
		}

		if popErr := writeBuffer.PopContext("ServerOnNetwork"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerOnNetwork")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerOnNetwork) IsServerOnNetwork() {}

func (m *_ServerOnNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
