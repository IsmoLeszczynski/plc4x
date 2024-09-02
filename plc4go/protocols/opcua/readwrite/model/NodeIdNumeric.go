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

// NodeIdNumeric is the corresponding interface of NodeIdNumeric
type NodeIdNumeric interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() uint32
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// _NodeIdNumeric is the data-structure of this message
type _NodeIdNumeric struct {
	NodeIdTypeDefinitionContract
	NamespaceIndex uint16
	Id             uint32
}

var _ NodeIdNumeric = (*_NodeIdNumeric)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdNumeric)(nil)

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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

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

		if popErr := writeBuffer.PopContext("NodeIdNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdNumeric")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdNumeric) IsNodeIdNumeric() {}

func (m *_NodeIdNumeric) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
