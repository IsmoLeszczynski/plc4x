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

// UserNameIdentityToken is the corresponding interface of UserNameIdentityToken
type UserNameIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	UserIdentityTokenDefinition
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// GetPassword returns Password (property field)
	GetPassword() PascalByteString
	// GetEncryptionAlgorithm returns EncryptionAlgorithm (property field)
	GetEncryptionAlgorithm() PascalString
}

// _UserNameIdentityToken is the data-structure of this message
type _UserNameIdentityToken struct {
	UserIdentityTokenDefinitionContract
	UserName            PascalString
	Password            PascalByteString
	EncryptionAlgorithm PascalString
}

var _ UserNameIdentityToken = (*_UserNameIdentityToken)(nil)
var _ UserIdentityTokenDefinitionRequirements = (*_UserNameIdentityToken)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserNameIdentityToken) GetIdentifier() string {
	return "username"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserNameIdentityToken) GetParent() UserIdentityTokenDefinitionContract {
	return m.UserIdentityTokenDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserNameIdentityToken) GetUserName() PascalString {
	return m.UserName
}

func (m *_UserNameIdentityToken) GetPassword() PascalByteString {
	return m.Password
}

func (m *_UserNameIdentityToken) GetEncryptionAlgorithm() PascalString {
	return m.EncryptionAlgorithm
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewUserNameIdentityToken factory function for _UserNameIdentityToken
func NewUserNameIdentityToken(userName PascalString, password PascalByteString, encryptionAlgorithm PascalString) *_UserNameIdentityToken {
	_result := &_UserNameIdentityToken{
		UserIdentityTokenDefinitionContract: NewUserIdentityTokenDefinition(),
		UserName:                            userName,
		Password:                            password,
		EncryptionAlgorithm:                 encryptionAlgorithm,
	}
	_result.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastUserNameIdentityToken(structType any) UserNameIdentityToken {
	if casted, ok := structType.(UserNameIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*UserNameIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_UserNameIdentityToken) GetTypeName() string {
	return "UserNameIdentityToken"
}

func (m *_UserNameIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).getLengthInBits(ctx))

	// Simple field (userName)
	lengthInBits += m.UserName.GetLengthInBits(ctx)

	// Simple field (password)
	lengthInBits += m.Password.GetLengthInBits(ctx)

	// Simple field (encryptionAlgorithm)
	lengthInBits += m.EncryptionAlgorithm.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_UserNameIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_UserNameIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_UserIdentityTokenDefinition, identifier string) (__userNameIdentityToken UserNameIdentityToken, err error) {
	m.UserIdentityTokenDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserNameIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserNameIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	userName, err := ReadSimpleField[PascalString](ctx, "userName", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'userName' field"))
	}
	m.UserName = userName

	password, err := ReadSimpleField[PascalByteString](ctx, "password", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'password' field"))
	}
	m.Password = password

	encryptionAlgorithm, err := ReadSimpleField[PascalString](ctx, "encryptionAlgorithm", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encryptionAlgorithm' field"))
	}
	m.EncryptionAlgorithm = encryptionAlgorithm

	if closeErr := readBuffer.CloseContext("UserNameIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for UserNameIdentityToken")
	}

	return m, nil
}

func (m *_UserNameIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_UserNameIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("UserNameIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for UserNameIdentityToken")
		}

		if err := WriteSimpleField[PascalString](ctx, "userName", m.GetUserName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'userName' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "password", m.GetPassword(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'password' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "encryptionAlgorithm", m.GetEncryptionAlgorithm(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'encryptionAlgorithm' field")
		}

		if popErr := writeBuffer.PopContext("UserNameIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for UserNameIdentityToken")
		}
		return nil
	}
	return m.UserIdentityTokenDefinitionContract.(*_UserIdentityTokenDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserNameIdentityToken) IsUserNameIdentityToken() {}

func (m *_UserNameIdentityToken) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
