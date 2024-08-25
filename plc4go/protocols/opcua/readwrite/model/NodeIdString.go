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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// NodeIdString is the corresponding interface of NodeIdString
type NodeIdString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetId returns Id (property field)
	GetId() PascalString
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdStringExactly can be used when we want exactly this type and not a type which fulfills NodeIdString.
// This is useful for switch cases.
type NodeIdStringExactly interface {
	NodeIdString
	isNodeIdString() bool
}

// _NodeIdString is the data-structure of this message
type _NodeIdString struct {
	*_NodeIdTypeDefinition
	NamespaceIndex uint16
	Id             PascalString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdString) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeString
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdString) InitializeParent(parent NodeIdTypeDefinition) {}

func (m *_NodeIdString) GetParent() NodeIdTypeDefinition {
	return m._NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdString) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NodeIdString) GetId() PascalString {
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

func (m *_NodeIdString) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId().GetStringValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdString factory function for _NodeIdString
func NewNodeIdString(namespaceIndex uint16, id PascalString) *_NodeIdString {
	_result := &_NodeIdString{
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
		_NodeIdTypeDefinition: NewNodeIdTypeDefinition(),
	}
	_result._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdString(structType any) NodeIdString {
	if casted, ok := structType.(NodeIdString); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdString); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdString) GetTypeName() string {
	return "NodeIdString"
}

func (m *_NodeIdString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (id)
	lengthInBits += m.Id.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdStringParse(ctx context.Context, theBytes []byte) (NodeIdString, error) {
	return NodeIdStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeIdString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of NodeIdString")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (id)
	if pullErr := readBuffer.PullContext("id"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for id")
	}
	_id, _idErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _idErr != nil {
		return nil, errors.Wrap(_idErr, "Error parsing 'id' field of NodeIdString")
	}
	id := _id.(PascalString)
	if closeErr := readBuffer.CloseContext("id"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for id")
	}

	// Virtual field
	_identifier := id.GetStringValue()
	identifier := fmt.Sprintf("%v", _identifier)
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdString")
	}

	// Create a partially initialized instance
	_child := &_NodeIdString{
		_NodeIdTypeDefinition: &_NodeIdTypeDefinition{},
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
	}
	_child._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeIdString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdString")
		}

		// Simple Field (namespaceIndex)
		namespaceIndex := uint16(m.GetNamespaceIndex())
		_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, uint16((namespaceIndex)))
		if _namespaceIndexErr != nil {
			return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
		}

		// Simple Field (id)
		if pushErr := writeBuffer.PushContext("id"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for id")
		}
		_idErr := writeBuffer.WriteSerializable(ctx, m.GetId())
		if popErr := writeBuffer.PopContext("id"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for id")
		}
		if _idErr != nil {
			return errors.Wrap(_idErr, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdString")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdString) isNodeIdString() bool {
	return true
}

func (m *_NodeIdString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
