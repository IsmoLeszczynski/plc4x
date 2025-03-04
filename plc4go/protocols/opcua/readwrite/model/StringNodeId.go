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

// StringNodeId is the corresponding interface of StringNodeId
type StringNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() PascalString
	// IsStringNodeId is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsStringNodeId()
	// CreateBuilder creates a StringNodeIdBuilder
	CreateStringNodeIdBuilder() StringNodeIdBuilder
}

// _StringNodeId is the data-structure of this message
type _StringNodeId struct {
	NamespaceIndex uint16
	Identifier     PascalString
}

var _ StringNodeId = (*_StringNodeId)(nil)

// NewStringNodeId factory function for _StringNodeId
func NewStringNodeId(namespaceIndex uint16, identifier PascalString) *_StringNodeId {
	if identifier == nil {
		panic("identifier of type PascalString for StringNodeId must not be nil")
	}
	return &_StringNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// StringNodeIdBuilder is a builder for StringNodeId
type StringNodeIdBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceIndex uint16, identifier PascalString) StringNodeIdBuilder
	// WithNamespaceIndex adds NamespaceIndex (property field)
	WithNamespaceIndex(uint16) StringNodeIdBuilder
	// WithIdentifier adds Identifier (property field)
	WithIdentifier(PascalString) StringNodeIdBuilder
	// WithIdentifierBuilder adds Identifier (property field) which is build by the builder
	WithIdentifierBuilder(func(PascalStringBuilder) PascalStringBuilder) StringNodeIdBuilder
	// Build builds the StringNodeId or returns an error if something is wrong
	Build() (StringNodeId, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() StringNodeId
}

// NewStringNodeIdBuilder() creates a StringNodeIdBuilder
func NewStringNodeIdBuilder() StringNodeIdBuilder {
	return &_StringNodeIdBuilder{_StringNodeId: new(_StringNodeId)}
}

type _StringNodeIdBuilder struct {
	*_StringNodeId

	err *utils.MultiError
}

var _ (StringNodeIdBuilder) = (*_StringNodeIdBuilder)(nil)

func (b *_StringNodeIdBuilder) WithMandatoryFields(namespaceIndex uint16, identifier PascalString) StringNodeIdBuilder {
	return b.WithNamespaceIndex(namespaceIndex).WithIdentifier(identifier)
}

func (b *_StringNodeIdBuilder) WithNamespaceIndex(namespaceIndex uint16) StringNodeIdBuilder {
	b.NamespaceIndex = namespaceIndex
	return b
}

func (b *_StringNodeIdBuilder) WithIdentifier(identifier PascalString) StringNodeIdBuilder {
	b.Identifier = identifier
	return b
}

func (b *_StringNodeIdBuilder) WithIdentifierBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) StringNodeIdBuilder {
	builder := builderSupplier(b.Identifier.CreatePascalStringBuilder())
	var err error
	b.Identifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_StringNodeIdBuilder) Build() (StringNodeId, error) {
	if b.Identifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'identifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._StringNodeId.deepCopy(), nil
}

func (b *_StringNodeIdBuilder) MustBuild() StringNodeId {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_StringNodeIdBuilder) DeepCopy() any {
	_copy := b.CreateStringNodeIdBuilder().(*_StringNodeIdBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateStringNodeIdBuilder creates a StringNodeIdBuilder
func (b *_StringNodeId) CreateStringNodeIdBuilder() StringNodeIdBuilder {
	if b == nil {
		return NewStringNodeIdBuilder()
	}
	return &_StringNodeIdBuilder{_StringNodeId: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StringNodeId) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_StringNodeId) GetIdentifier() PascalString {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastStringNodeId(structType any) StringNodeId {
	if casted, ok := structType.(StringNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*StringNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_StringNodeId) GetTypeName() string {
	return "StringNodeId"
}

func (m *_StringNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (identifier)
	lengthInBits += m.Identifier.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StringNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StringNodeIdParse(ctx context.Context, theBytes []byte) (StringNodeId, error) {
	return StringNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func StringNodeIdParseWithBufferProducer() func(ctx context.Context, readBuffer utils.ReadBuffer) (StringNodeId, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (StringNodeId, error) {
		return StringNodeIdParseWithBuffer(ctx, readBuffer)
	}
}

func StringNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (StringNodeId, error) {
	v, err := (&_StringNodeId{}).parse(ctx, readBuffer)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_StringNodeId) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__stringNodeId StringNodeId, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StringNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StringNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	identifier, err := ReadSimpleField[PascalString](ctx, "identifier", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	m.Identifier = identifier

	if closeErr := readBuffer.CloseContext("StringNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StringNodeId")
	}

	return m, nil
}

func (m *_StringNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StringNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("StringNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StringNodeId")
	}

	if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
		return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
	}

	if err := WriteSimpleField[PascalString](ctx, "identifier", m.GetIdentifier(), WriteComplex[PascalString](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("StringNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StringNodeId")
	}
	return nil
}

func (m *_StringNodeId) IsStringNodeId() {}

func (m *_StringNodeId) DeepCopy() any {
	return m.deepCopy()
}

func (m *_StringNodeId) deepCopy() *_StringNodeId {
	if m == nil {
		return nil
	}
	_StringNodeIdCopy := &_StringNodeId{
		m.NamespaceIndex,
		utils.DeepCopy[PascalString](m.Identifier),
	}
	return _StringNodeIdCopy
}

func (m *_StringNodeId) String() string {
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
