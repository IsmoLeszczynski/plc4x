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
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// GetTokenData returns TokenData (property field)
	GetTokenData() PascalByteString
	// GetEncryptionAlgorithm returns EncryptionAlgorithm (property field)
	GetEncryptionAlgorithm() PascalString
	// IsIssuedIdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsIssuedIdentityToken()
	// CreateBuilder creates a IssuedIdentityTokenBuilder
	CreateIssuedIdentityTokenBuilder() IssuedIdentityTokenBuilder
}

// _IssuedIdentityToken is the data-structure of this message
type _IssuedIdentityToken struct {
	ExtensionObjectDefinitionContract
	PolicyId            PascalString
	TokenData           PascalByteString
	EncryptionAlgorithm PascalString
}

var _ IssuedIdentityToken = (*_IssuedIdentityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_IssuedIdentityToken)(nil)

// NewIssuedIdentityToken factory function for _IssuedIdentityToken
func NewIssuedIdentityToken(policyId PascalString, tokenData PascalByteString, encryptionAlgorithm PascalString) *_IssuedIdentityToken {
	if policyId == nil {
		panic("policyId of type PascalString for IssuedIdentityToken must not be nil")
	}
	if tokenData == nil {
		panic("tokenData of type PascalByteString for IssuedIdentityToken must not be nil")
	}
	if encryptionAlgorithm == nil {
		panic("encryptionAlgorithm of type PascalString for IssuedIdentityToken must not be nil")
	}
	_result := &_IssuedIdentityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
		TokenData:                         tokenData,
		EncryptionAlgorithm:               encryptionAlgorithm,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// IssuedIdentityTokenBuilder is a builder for IssuedIdentityToken
type IssuedIdentityTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString, tokenData PascalByteString, encryptionAlgorithm PascalString) IssuedIdentityTokenBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) IssuedIdentityTokenBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) IssuedIdentityTokenBuilder
	// WithTokenData adds TokenData (property field)
	WithTokenData(PascalByteString) IssuedIdentityTokenBuilder
	// WithTokenDataBuilder adds TokenData (property field) which is build by the builder
	WithTokenDataBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) IssuedIdentityTokenBuilder
	// WithEncryptionAlgorithm adds EncryptionAlgorithm (property field)
	WithEncryptionAlgorithm(PascalString) IssuedIdentityTokenBuilder
	// WithEncryptionAlgorithmBuilder adds EncryptionAlgorithm (property field) which is build by the builder
	WithEncryptionAlgorithmBuilder(func(PascalStringBuilder) PascalStringBuilder) IssuedIdentityTokenBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the IssuedIdentityToken or returns an error if something is wrong
	Build() (IssuedIdentityToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() IssuedIdentityToken
}

// NewIssuedIdentityTokenBuilder() creates a IssuedIdentityTokenBuilder
func NewIssuedIdentityTokenBuilder() IssuedIdentityTokenBuilder {
	return &_IssuedIdentityTokenBuilder{_IssuedIdentityToken: new(_IssuedIdentityToken)}
}

type _IssuedIdentityTokenBuilder struct {
	*_IssuedIdentityToken

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (IssuedIdentityTokenBuilder) = (*_IssuedIdentityTokenBuilder)(nil)

func (b *_IssuedIdentityTokenBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_IssuedIdentityTokenBuilder) WithMandatoryFields(policyId PascalString, tokenData PascalByteString, encryptionAlgorithm PascalString) IssuedIdentityTokenBuilder {
	return b.WithPolicyId(policyId).WithTokenData(tokenData).WithEncryptionAlgorithm(encryptionAlgorithm)
}

func (b *_IssuedIdentityTokenBuilder) WithPolicyId(policyId PascalString) IssuedIdentityTokenBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_IssuedIdentityTokenBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) IssuedIdentityTokenBuilder {
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

func (b *_IssuedIdentityTokenBuilder) WithTokenData(tokenData PascalByteString) IssuedIdentityTokenBuilder {
	b.TokenData = tokenData
	return b
}

func (b *_IssuedIdentityTokenBuilder) WithTokenDataBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) IssuedIdentityTokenBuilder {
	builder := builderSupplier(b.TokenData.CreatePascalByteStringBuilder())
	var err error
	b.TokenData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_IssuedIdentityTokenBuilder) WithEncryptionAlgorithm(encryptionAlgorithm PascalString) IssuedIdentityTokenBuilder {
	b.EncryptionAlgorithm = encryptionAlgorithm
	return b
}

func (b *_IssuedIdentityTokenBuilder) WithEncryptionAlgorithmBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) IssuedIdentityTokenBuilder {
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

func (b *_IssuedIdentityTokenBuilder) Build() (IssuedIdentityToken, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.TokenData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'tokenData' not set"))
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
	return b._IssuedIdentityToken.deepCopy(), nil
}

func (b *_IssuedIdentityTokenBuilder) MustBuild() IssuedIdentityToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_IssuedIdentityTokenBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_IssuedIdentityTokenBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_IssuedIdentityTokenBuilder) DeepCopy() any {
	_copy := b.CreateIssuedIdentityTokenBuilder().(*_IssuedIdentityTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateIssuedIdentityTokenBuilder creates a IssuedIdentityTokenBuilder
func (b *_IssuedIdentityToken) CreateIssuedIdentityTokenBuilder() IssuedIdentityTokenBuilder {
	if b == nil {
		return NewIssuedIdentityTokenBuilder()
	}
	return &_IssuedIdentityTokenBuilder{_IssuedIdentityToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IssuedIdentityToken) GetExtensionId() int32 {
	return int32(940)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IssuedIdentityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IssuedIdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

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
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	// Simple field (tokenData)
	lengthInBits += m.TokenData.GetLengthInBits(ctx)

	// Simple field (encryptionAlgorithm)
	lengthInBits += m.EncryptionAlgorithm.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_IssuedIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_IssuedIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__issuedIdentityToken IssuedIdentityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IssuedIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IssuedIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

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

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
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
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_IssuedIdentityToken) IsIssuedIdentityToken() {}

func (m *_IssuedIdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_IssuedIdentityToken) deepCopy() *_IssuedIdentityToken {
	if m == nil {
		return nil
	}
	_IssuedIdentityTokenCopy := &_IssuedIdentityToken{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.PolicyId),
		utils.DeepCopy[PascalByteString](m.TokenData),
		utils.DeepCopy[PascalString](m.EncryptionAlgorithm),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _IssuedIdentityTokenCopy
}

func (m *_IssuedIdentityToken) String() string {
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
