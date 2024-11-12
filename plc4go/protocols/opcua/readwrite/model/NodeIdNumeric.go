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

// NodeIdNumeric is the corresponding interface of NodeIdNumeric
type NodeIdNumeric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() uint32
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
	// GetNamespace returns Namespace (virtual field)
	GetNamespace() int16
	// IsNodeIdNumeric is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdNumeric()
	// CreateBuilder creates a NodeIdNumericBuilder
	CreateNodeIdNumericBuilder() NodeIdNumericBuilder
}

// _NodeIdNumeric is the data-structure of this message
type _NodeIdNumeric struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint16
	Id             uint32
}

var _ NodeIdNumeric = (*_NodeIdNumeric)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdNumeric)(nil)

// NewNodeIdNumeric factory function for _NodeIdNumeric
func NewNodeIdNumeric(namespaceIndex uint16, id uint32) *_NodeIdNumeric {
	_result := &_NodeIdNumeric{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		NamespaceIndex:               namespaceIndex,
		Id:                           id,
	}
	_result.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdNumericBuilder is a builder for NodeIdNumeric
type NodeIdNumericBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(namespaceIndex uint16, id uint32) NodeIdNumericBuilder
	// WithNamespaceIndex adds NamespaceIndex (property field)
	WithNamespaceIndex(uint16) NodeIdNumericBuilder
	// WithId adds Id (property field)
	WithId(uint32) NodeIdNumericBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() NodeIdTypeDefinitionBuilder
	// Build builds the NodeIdNumeric or returns an error if something is wrong
	Build() (NodeIdNumeric, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdNumeric
}

// NewNodeIdNumericBuilder() creates a NodeIdNumericBuilder
func NewNodeIdNumericBuilder() NodeIdNumericBuilder {
	return &_NodeIdNumericBuilder{_NodeIdNumeric: new(_NodeIdNumeric)}
}

type _NodeIdNumericBuilder struct {
	*_NodeIdNumeric

	parentBuilder *_NodeIdTypeDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeIdNumericBuilder) = (*_NodeIdNumericBuilder)(nil)

func (b *_NodeIdNumericBuilder) setParent(contract NodeIdTypeDefinitionContract) {
	b.NodeIdTypeDefinitionContract = contract
}

func (b *_NodeIdNumericBuilder) WithMandatoryFields(namespaceIndex uint16, id uint32) NodeIdNumericBuilder {
	return b.WithNamespaceIndex(namespaceIndex).WithId(id)
}

func (b *_NodeIdNumericBuilder) WithNamespaceIndex(namespaceIndex uint16) NodeIdNumericBuilder {
	b.NamespaceIndex = namespaceIndex
	return b
}

func (b *_NodeIdNumericBuilder) WithId(id uint32) NodeIdNumericBuilder {
	b.Id = id
	return b
}

func (b *_NodeIdNumericBuilder) Build() (NodeIdNumeric, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdNumeric.deepCopy(), nil
}

func (b *_NodeIdNumericBuilder) MustBuild() NodeIdNumeric {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NodeIdNumericBuilder) Done() NodeIdTypeDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewNodeIdTypeDefinitionBuilder().(*_NodeIdTypeDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_NodeIdNumericBuilder) buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error) {
	return b.Build()
}

func (b *_NodeIdNumericBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdNumericBuilder().(*_NodeIdNumericBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdNumericBuilder creates a NodeIdNumericBuilder
func (b *_NodeIdNumeric) CreateNodeIdNumericBuilder() NodeIdNumericBuilder {
	if b == nil {
		return NewNodeIdNumericBuilder()
	}
	return &_NodeIdNumericBuilder{_NodeIdNumeric: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdNumeric) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeNumeric
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdNumeric) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdNumeric) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdNumeric) GetId() uint32 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdNumeric) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

func (m *_NodeIdNumeric) GetNamespace() int16 {
	ctx := context.Background()
	_ = ctx
	return int16(m.GetNamespaceIndex())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdNumeric(structType any) NodeIdNumeric {
	if casted, ok := structType.(NodeIdNumeric); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdNumeric); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdNumeric) GetTypeName() string {
	return "NodeIdNumeric"
}

func (m *_NodeIdNumeric) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += 32

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdNumeric) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdNumeric) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdNumeric NodeIdNumeric, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdNumeric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdNumeric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	namespaceIndex, err := ReadSimpleField(ctx, "namespaceIndex", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespaceIndex' field"))
	}
	m.NamespaceIndex = namespaceIndex

	id, err := ReadSimpleField(ctx, "id", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	namespace, err := ReadVirtualField[int16](ctx, "namespace", (*int16)(nil), namespaceIndex)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'namespace' field"))
	}
	_ = namespace

	if closeErr := readBuffer.CloseContext("NodeIdNumeric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdNumeric")
	}

	return m, nil
}

func (m *_NodeIdNumeric) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdNumeric) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdNumeric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdNumeric")
		}

		if err := WriteSimpleField[uint16](ctx, "namespaceIndex", m.GetNamespaceIndex(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'namespaceIndex' field")
		}

		if err := WriteSimpleField[uint32](ctx, "id", m.GetId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}
		// Virtual field
		namespace := m.GetNamespace()
		_ = namespace
		if _namespaceErr := writeBuffer.WriteVirtual(ctx, "namespace", m.GetNamespace()); _namespaceErr != nil {
			return errors.Wrap(_namespaceErr, "Error serializing 'namespace' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdNumeric")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdNumeric) IsNodeIdNumeric() {}

func (m *_NodeIdNumeric) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdNumeric) deepCopy() *_NodeIdNumeric {
	if m == nil {
		return nil
	}
	_NodeIdNumericCopy := &_NodeIdNumeric{
		m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).deepCopy(),
		m.NamespaceIndex,
		m.Id,
	}
	m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = m
	return _NodeIdNumericCopy
}

func (m *_NodeIdNumeric) String() string {
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
