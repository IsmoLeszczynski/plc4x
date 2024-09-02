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

// ChannelSecurityToken is the corresponding interface of ChannelSecurityToken
type ChannelSecurityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetChannelId returns ChannelId (property field)
	GetChannelId() uint32
	// GetTokenId returns TokenId (property field)
	GetTokenId() uint32
	// GetCreatedAt returns CreatedAt (property field)
	GetCreatedAt() int64
	// GetRevisedLifetime returns RevisedLifetime (property field)
	GetRevisedLifetime() uint32
}

// ChannelSecurityTokenExactly can be used when we want exactly this type and not a type which fulfills ChannelSecurityToken.
// This is useful for switch cases.
type ChannelSecurityTokenExactly interface {
	ChannelSecurityToken
	isChannelSecurityToken() bool
}

// _ChannelSecurityToken is the data-structure of this message
type _ChannelSecurityToken struct {
	ExtensionObjectDefinitionContract
	ChannelId       uint32
	TokenId         uint32
	CreatedAt       int64
	RevisedLifetime uint32
}

var _ ChannelSecurityToken = (*_ChannelSecurityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ChannelSecurityToken)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChannelSecurityToken) GetIdentifier() string {
	return "443"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChannelSecurityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChannelSecurityToken) GetChannelId() uint32 {
	return m.ChannelId
}

func (m *_ChannelSecurityToken) GetTokenId() uint32 {
	return m.TokenId
}

func (m *_ChannelSecurityToken) GetCreatedAt() int64 {
	return m.CreatedAt
}

func (m *_ChannelSecurityToken) GetRevisedLifetime() uint32 {
	return m.RevisedLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewChannelSecurityToken factory function for _ChannelSecurityToken
func NewChannelSecurityToken(channelId uint32, tokenId uint32, createdAt int64, revisedLifetime uint32) *_ChannelSecurityToken {
	_result := &_ChannelSecurityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ChannelId:                         channelId,
		TokenId:                           tokenId,
		CreatedAt:                         createdAt,
		RevisedLifetime:                   revisedLifetime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastChannelSecurityToken(structType any) ChannelSecurityToken {
	if casted, ok := structType.(ChannelSecurityToken); ok {
		return casted
	}
	if casted, ok := structType.(*ChannelSecurityToken); ok {
		return *casted
	}
	return nil
}

func (m *_ChannelSecurityToken) GetTypeName() string {
	return "ChannelSecurityToken"
}

func (m *_ChannelSecurityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (channelId)
	lengthInBits += 32

	// Simple field (tokenId)
	lengthInBits += 32

	// Simple field (createdAt)
	lengthInBits += 64

	// Simple field (revisedLifetime)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ChannelSecurityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ChannelSecurityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__channelSecurityToken ChannelSecurityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChannelSecurityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChannelSecurityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	channelId, err := ReadSimpleField(ctx, "channelId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'channelId' field"))
	}
	m.ChannelId = channelId

	tokenId, err := ReadSimpleField(ctx, "tokenId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tokenId' field"))
	}
	m.TokenId = tokenId

	createdAt, err := ReadSimpleField(ctx, "createdAt", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'createdAt' field"))
	}
	m.CreatedAt = createdAt

	revisedLifetime, err := ReadSimpleField(ctx, "revisedLifetime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisedLifetime' field"))
	}
	m.RevisedLifetime = revisedLifetime

	if closeErr := readBuffer.CloseContext("ChannelSecurityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChannelSecurityToken")
	}

	return m, nil
}

func (m *_ChannelSecurityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChannelSecurityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChannelSecurityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChannelSecurityToken")
		}

		if err := WriteSimpleField[uint32](ctx, "channelId", m.GetChannelId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'channelId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "tokenId", m.GetTokenId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'tokenId' field")
		}

		if err := WriteSimpleField[int64](ctx, "createdAt", m.GetCreatedAt(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'createdAt' field")
		}

		if err := WriteSimpleField[uint32](ctx, "revisedLifetime", m.GetRevisedLifetime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisedLifetime' field")
		}

		if popErr := writeBuffer.PopContext("ChannelSecurityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChannelSecurityToken")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ChannelSecurityToken) isChannelSecurityToken() bool {
	return true
}

func (m *_ChannelSecurityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
