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

// IssuedIdentityToken is the corresponding interface of IssuedIdentityToken
type IssuedIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	UserIdentityTokenDefinition
	// GetTokenData returns TokenData (property field)
	GetTokenData() PascalByteString
	// GetEncryptionAlgorithm returns EncryptionAlgorithm (property field)
	GetEncryptionAlgorithm() PascalString
}

// IssuedIdentityTokenExactly can be used when we want exactly this type and not a type which fulfills IssuedIdentityToken.
// This is useful for switch cases.
type IssuedIdentityTokenExactly interface {
	IssuedIdentityToken
	isIssuedIdentityToken() bool
}

// _IssuedIdentityToken is the data-structure of this message
type _IssuedIdentityToken struct {
	UserIdentityTokenDefinitionContract
	TokenData           PascalByteString
	EncryptionAlgorithm PascalString
}

var _ IssuedIdentityToken = (*_IssuedIdentityToken)(nil)
var _ UserIdentityTokenDefinitionRequirements = (*_IssuedIdentityToken)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IssuedIdentityToken) GetIdentifier() string {
	return "identity"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IssuedIdentityToken) GetParent() UserIdentityTokenDefinitionContract {
	return m.UserIdentityTokenDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IssuedIdentityToken) GetTokenData() PascalByteString {
	return m.TokenData
}

func (m *_IssuedIdentityToken) GetEncryptionAlgorithm() PascalString {
	return m.EncryptionAlgorithm
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIssuedIdentityToken factory function for _IssuedIdentityToken
func NewIssuedIdentityToken(tokenData PascalByteString, encryptionAlgorithm PascalString) *_IssuedIdentityToken {
	_result := &_IssuedIdentityToken{
		UserIdentityTokenDefinitionContract: NewUserIdentityTokenDefinition(),
		TokenData:                           tokenData,
		EncryptionAlgorithm:                 encryptionAlgorithm,
	}
	_result.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIssuedIdentityToken(structType any) IssuedIdentityToken {
	if casted, ok := structType.(IssuedIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*IssuedIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_IssuedIdentityToken) GetTypeName() string {
	return "IssuedIdentityToken"
}

func (m *_IssuedIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).getLengthInBits(ctx))

	// Simple field (tokenData)
	lengthInBits += m.TokenData.GetLengthInBits(ctx)

	// Simple field (encryptionAlgorithm)
	lengthInBits += m.EncryptionAlgorithm.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_IssuedIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IssuedIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_UserIdentityTokenDefinition, identifier string) (__issuedIdentityToken IssuedIdentityToken, err error) {
	m.UserIdentityTokenDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IssuedIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IssuedIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	tokenData, err := ReadSimpleField[PascalByteString](ctx, "tokenData", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tokenData' field"))
	}
	m.TokenData = tokenData

	encryptionAlgorithm, err := ReadSimpleField[PascalString](ctx, "encryptionAlgorithm", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encryptionAlgorithm' field"))
	}
	m.EncryptionAlgorithm = encryptionAlgorithm

	if closeErr := readBuffer.CloseContext("IssuedIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IssuedIdentityToken")
	}

	return m, nil
}

func (m *_IssuedIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IssuedIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IssuedIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IssuedIdentityToken")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "tokenData", m.GetTokenData(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'tokenData' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "encryptionAlgorithm", m.GetEncryptionAlgorithm(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'encryptionAlgorithm' field")
		}

		if popErr := writeBuffer.PopContext("IssuedIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IssuedIdentityToken")
		}
		return nil
	}
	return m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IssuedIdentityToken) isIssuedIdentityToken() bool {
	return true
}

func (m *_IssuedIdentityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
