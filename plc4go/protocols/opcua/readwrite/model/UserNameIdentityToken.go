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

// UserNameIdentityToken is the corresponding interface of UserNameIdentityToken
type UserNameIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetUserName returns UserName (property field)
	GetUserName() PascalString
	// GetPassword returns Password (property field)
	GetPassword() PascalByteString
	// GetEncryptionAlgorithm returns EncryptionAlgorithm (property field)
	GetEncryptionAlgorithm() PascalString
	// IsUserNameIdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsUserNameIdentityToken()
	// CreateBuilder creates a UserNameIdentityTokenBuilder
	CreateUserNameIdentityTokenBuilder() UserNameIdentityTokenBuilder
}

// _UserNameIdentityToken is the data-structure of this message
type _UserNameIdentityToken struct {
	ExtensionObjectDefinitionContract
	PolicyId            PascalString
	UserName            PascalString
	Password            PascalByteString
	EncryptionAlgorithm PascalString
}

var _ UserNameIdentityToken = (*_UserNameIdentityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_UserNameIdentityToken)(nil)

// NewUserNameIdentityToken factory function for _UserNameIdentityToken
func NewUserNameIdentityToken(policyId PascalString, userName PascalString, password PascalByteString, encryptionAlgorithm PascalString) *_UserNameIdentityToken {
	if policyId == nil {
		panic("policyId of type PascalString for UserNameIdentityToken must not be nil")
	}
	if userName == nil {
		panic("userName of type PascalString for UserNameIdentityToken must not be nil")
	}
	if password == nil {
		panic("password of type PascalByteString for UserNameIdentityToken must not be nil")
	}
	if encryptionAlgorithm == nil {
		panic("encryptionAlgorithm of type PascalString for UserNameIdentityToken must not be nil")
	}
	_result := &_UserNameIdentityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
		UserName:                          userName,
		Password:                          password,
		EncryptionAlgorithm:               encryptionAlgorithm,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// UserNameIdentityTokenBuilder is a builder for UserNameIdentityToken
type UserNameIdentityTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString, userName PascalString, password PascalByteString, encryptionAlgorithm PascalString) UserNameIdentityTokenBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) UserNameIdentityTokenBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder
	// WithUserName adds UserName (property field)
	WithUserName(PascalString) UserNameIdentityTokenBuilder
	// WithUserNameBuilder adds UserName (property field) which is build by the builder
	WithUserNameBuilder(func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder
	// WithPassword adds Password (property field)
	WithPassword(PascalByteString) UserNameIdentityTokenBuilder
	// WithPasswordBuilder adds Password (property field) which is build by the builder
	WithPasswordBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) UserNameIdentityTokenBuilder
	// WithEncryptionAlgorithm adds EncryptionAlgorithm (property field)
	WithEncryptionAlgorithm(PascalString) UserNameIdentityTokenBuilder
	// WithEncryptionAlgorithmBuilder adds EncryptionAlgorithm (property field) which is build by the builder
	WithEncryptionAlgorithmBuilder(func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the UserNameIdentityToken or returns an error if something is wrong
	Build() (UserNameIdentityToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() UserNameIdentityToken
}

// NewUserNameIdentityTokenBuilder() creates a UserNameIdentityTokenBuilder
func NewUserNameIdentityTokenBuilder() UserNameIdentityTokenBuilder {
	return &_UserNameIdentityTokenBuilder{_UserNameIdentityToken: new(_UserNameIdentityToken)}
}

type _UserNameIdentityTokenBuilder struct {
	*_UserNameIdentityToken

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (UserNameIdentityTokenBuilder) = (*_UserNameIdentityTokenBuilder)(nil)

func (b *_UserNameIdentityTokenBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_UserNameIdentityTokenBuilder) WithMandatoryFields(policyId PascalString, userName PascalString, password PascalByteString, encryptionAlgorithm PascalString) UserNameIdentityTokenBuilder {
	return b.WithPolicyId(policyId).WithUserName(userName).WithPassword(password).WithEncryptionAlgorithm(encryptionAlgorithm)
}

func (b *_UserNameIdentityTokenBuilder) WithPolicyId(policyId PascalString) UserNameIdentityTokenBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder {
	builder := builderSupplier(b.PolicyId.CreatePascalStringBuilder())
	var err error
	b.PolicyId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithUserName(userName PascalString) UserNameIdentityTokenBuilder {
	b.UserName = userName
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithUserNameBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder {
	builder := builderSupplier(b.UserName.CreatePascalStringBuilder())
	var err error
	b.UserName, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithPassword(password PascalByteString) UserNameIdentityTokenBuilder {
	b.Password = password
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithPasswordBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) UserNameIdentityTokenBuilder {
	builder := builderSupplier(b.Password.CreatePascalByteStringBuilder())
	var err error
	b.Password, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithEncryptionAlgorithm(encryptionAlgorithm PascalString) UserNameIdentityTokenBuilder {
	b.EncryptionAlgorithm = encryptionAlgorithm
	return b
}

func (b *_UserNameIdentityTokenBuilder) WithEncryptionAlgorithmBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) UserNameIdentityTokenBuilder {
	builder := builderSupplier(b.EncryptionAlgorithm.CreatePascalStringBuilder())
	var err error
	b.EncryptionAlgorithm, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_UserNameIdentityTokenBuilder) Build() (UserNameIdentityToken, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.UserName == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'userName' not set"))
	}
	if b.Password == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'password' not set"))
	}
	if b.EncryptionAlgorithm == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'encryptionAlgorithm' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._UserNameIdentityToken.deepCopy(), nil
}

func (b *_UserNameIdentityTokenBuilder) MustBuild() UserNameIdentityToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_UserNameIdentityTokenBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_UserNameIdentityTokenBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_UserNameIdentityTokenBuilder) DeepCopy() any {
	_copy := b.CreateUserNameIdentityTokenBuilder().(*_UserNameIdentityTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateUserNameIdentityTokenBuilder creates a UserNameIdentityTokenBuilder
func (b *_UserNameIdentityToken) CreateUserNameIdentityTokenBuilder() UserNameIdentityTokenBuilder {
	if b == nil {
		return NewUserNameIdentityTokenBuilder()
	}
	return &_UserNameIdentityTokenBuilder{_UserNameIdentityToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_UserNameIdentityToken) GetExtensionId() int32 {
	return int32(324)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_UserNameIdentityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_UserNameIdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

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
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

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

func (m *_UserNameIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__userNameIdentityToken UserNameIdentityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("UserNameIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for UserNameIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

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

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
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
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_UserNameIdentityToken) IsUserNameIdentityToken() {}

func (m *_UserNameIdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_UserNameIdentityToken) deepCopy() *_UserNameIdentityToken {
	if m == nil {
		return nil
	}
	_UserNameIdentityTokenCopy := &_UserNameIdentityToken{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.PolicyId),
		utils.DeepCopy[PascalString](m.UserName),
		utils.DeepCopy[PascalByteString](m.Password),
		utils.DeepCopy[PascalString](m.EncryptionAlgorithm),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _UserNameIdentityTokenCopy
}

func (m *_UserNameIdentityToken) String() string {
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
