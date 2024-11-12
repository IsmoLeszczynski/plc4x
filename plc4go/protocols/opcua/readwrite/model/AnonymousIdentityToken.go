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

// AnonymousIdentityToken is the corresponding interface of AnonymousIdentityToken
type AnonymousIdentityToken interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetPolicyId returns PolicyId (property field)
	GetPolicyId() PascalString
	// IsAnonymousIdentityToken is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAnonymousIdentityToken()
	// CreateBuilder creates a AnonymousIdentityTokenBuilder
	CreateAnonymousIdentityTokenBuilder() AnonymousIdentityTokenBuilder
}

// _AnonymousIdentityToken is the data-structure of this message
type _AnonymousIdentityToken struct {
	ExtensionObjectDefinitionContract
	PolicyId PascalString
}

var _ AnonymousIdentityToken = (*_AnonymousIdentityToken)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AnonymousIdentityToken)(nil)

// NewAnonymousIdentityToken factory function for _AnonymousIdentityToken
func NewAnonymousIdentityToken(policyId PascalString) *_AnonymousIdentityToken {
	if policyId == nil {
		panic("policyId of type PascalString for AnonymousIdentityToken must not be nil")
	}
	_result := &_AnonymousIdentityToken{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		PolicyId:                          policyId,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AnonymousIdentityTokenBuilder is a builder for AnonymousIdentityToken
type AnonymousIdentityTokenBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(policyId PascalString) AnonymousIdentityTokenBuilder
	// WithPolicyId adds PolicyId (property field)
	WithPolicyId(PascalString) AnonymousIdentityTokenBuilder
	// WithPolicyIdBuilder adds PolicyId (property field) which is build by the builder
	WithPolicyIdBuilder(func(PascalStringBuilder) PascalStringBuilder) AnonymousIdentityTokenBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the AnonymousIdentityToken or returns an error if something is wrong
	Build() (AnonymousIdentityToken, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AnonymousIdentityToken
}

// NewAnonymousIdentityTokenBuilder() creates a AnonymousIdentityTokenBuilder
func NewAnonymousIdentityTokenBuilder() AnonymousIdentityTokenBuilder {
	return &_AnonymousIdentityTokenBuilder{_AnonymousIdentityToken: new(_AnonymousIdentityToken)}
}

type _AnonymousIdentityTokenBuilder struct {
	*_AnonymousIdentityToken

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (AnonymousIdentityTokenBuilder) = (*_AnonymousIdentityTokenBuilder)(nil)

func (b *_AnonymousIdentityTokenBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_AnonymousIdentityTokenBuilder) WithMandatoryFields(policyId PascalString) AnonymousIdentityTokenBuilder {
	return b.WithPolicyId(policyId)
}

func (b *_AnonymousIdentityTokenBuilder) WithPolicyId(policyId PascalString) AnonymousIdentityTokenBuilder {
	b.PolicyId = policyId
	return b
}

func (b *_AnonymousIdentityTokenBuilder) WithPolicyIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) AnonymousIdentityTokenBuilder {
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

func (b *_AnonymousIdentityTokenBuilder) Build() (AnonymousIdentityToken, error) {
	if b.PolicyId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'policyId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AnonymousIdentityToken.deepCopy(), nil
}

func (b *_AnonymousIdentityTokenBuilder) MustBuild() AnonymousIdentityToken {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AnonymousIdentityTokenBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_AnonymousIdentityTokenBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_AnonymousIdentityTokenBuilder) DeepCopy() any {
	_copy := b.CreateAnonymousIdentityTokenBuilder().(*_AnonymousIdentityTokenBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAnonymousIdentityTokenBuilder creates a AnonymousIdentityTokenBuilder
func (b *_AnonymousIdentityToken) CreateAnonymousIdentityTokenBuilder() AnonymousIdentityTokenBuilder {
	if b == nil {
		return NewAnonymousIdentityTokenBuilder()
	}
	return &_AnonymousIdentityTokenBuilder{_AnonymousIdentityToken: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AnonymousIdentityToken) GetExtensionId() int32 {
	return int32(321)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AnonymousIdentityToken) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AnonymousIdentityToken) GetPolicyId() PascalString {
	return m.PolicyId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAnonymousIdentityToken(structType any) AnonymousIdentityToken {
	if casted, ok := structType.(AnonymousIdentityToken); ok {
		return casted
	}
	if casted, ok := structType.(*AnonymousIdentityToken); ok {
		return *casted
	}
	return nil
}

func (m *_AnonymousIdentityToken) GetTypeName() string {
	return "AnonymousIdentityToken"
}

func (m *_AnonymousIdentityToken) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (policyId)
	lengthInBits += m.PolicyId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AnonymousIdentityToken) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AnonymousIdentityToken) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__anonymousIdentityToken AnonymousIdentityToken, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AnonymousIdentityToken"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AnonymousIdentityToken")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	policyId, err := ReadSimpleField[PascalString](ctx, "policyId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'policyId' field"))
	}
	m.PolicyId = policyId

	if closeErr := readBuffer.CloseContext("AnonymousIdentityToken"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AnonymousIdentityToken")
	}

	return m, nil
}

func (m *_AnonymousIdentityToken) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AnonymousIdentityToken) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AnonymousIdentityToken"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AnonymousIdentityToken")
		}

		if err := WriteSimpleField[PascalString](ctx, "policyId", m.GetPolicyId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'policyId' field")
		}

		if popErr := writeBuffer.PopContext("AnonymousIdentityToken"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AnonymousIdentityToken")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AnonymousIdentityToken) IsAnonymousIdentityToken() {}

func (m *_AnonymousIdentityToken) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AnonymousIdentityToken) deepCopy() *_AnonymousIdentityToken {
	if m == nil {
		return nil
	}
	_AnonymousIdentityTokenCopy := &_AnonymousIdentityToken{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[PascalString](m.PolicyId),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _AnonymousIdentityTokenCopy
}

func (m *_AnonymousIdentityToken) String() string {
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
