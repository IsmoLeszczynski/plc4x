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

// AddReferencesItem is the corresponding interface of AddReferencesItem
type AddReferencesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSourceNodeId returns SourceNodeId (property field)
	GetSourceNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetServerUri returns TargetServerUri (property field)
	GetTargetServerUri() PascalString
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() ExpandedNodeId
	// GetTargetNodeClass returns TargetNodeClass (property field)
	GetTargetNodeClass() NodeClass
	// IsAddReferencesItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAddReferencesItem()
}

// _AddReferencesItem is the data-structure of this message
type _AddReferencesItem struct {
	ExtensionObjectDefinitionContract
	SourceNodeId    NodeId
	ReferenceTypeId NodeId
	IsForward       bool
	TargetServerUri PascalString
	TargetNodeId    ExpandedNodeId
	TargetNodeClass NodeClass
	// Reserved Fields
	reservedField0 *uint8
}

var _ AddReferencesItem = (*_AddReferencesItem)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_AddReferencesItem)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AddReferencesItem) GetIdentifier() string {
	return "381"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AddReferencesItem) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AddReferencesItem) GetSourceNodeId() NodeId {
	return m.SourceNodeId
}

func (m *_AddReferencesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_AddReferencesItem) GetIsForward() bool {
	return m.IsForward
}

func (m *_AddReferencesItem) GetTargetServerUri() PascalString {
	return m.TargetServerUri
}

func (m *_AddReferencesItem) GetTargetNodeId() ExpandedNodeId {
	return m.TargetNodeId
}

func (m *_AddReferencesItem) GetTargetNodeClass() NodeClass {
	return m.TargetNodeClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAddReferencesItem factory function for _AddReferencesItem
func NewAddReferencesItem(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetServerUri PascalString, targetNodeId ExpandedNodeId, targetNodeClass NodeClass) *_AddReferencesItem {
	if sourceNodeId == nil {
		panic("sourceNodeId of type NodeId for AddReferencesItem must not be nil")
	}
	if referenceTypeId == nil {
		panic("referenceTypeId of type NodeId for AddReferencesItem must not be nil")
	}
	if targetServerUri == nil {
		panic("targetServerUri of type PascalString for AddReferencesItem must not be nil")
	}
	if targetNodeId == nil {
		panic("targetNodeId of type ExpandedNodeId for AddReferencesItem must not be nil")
	}
	_result := &_AddReferencesItem{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SourceNodeId:                      sourceNodeId,
		ReferenceTypeId:                   referenceTypeId,
		IsForward:                         isForward,
		TargetServerUri:                   targetServerUri,
		TargetNodeId:                      targetNodeId,
		TargetNodeClass:                   targetNodeClass,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAddReferencesItem(structType any) AddReferencesItem {
	if casted, ok := structType.(AddReferencesItem); ok {
		return casted
	}
	if casted, ok := structType.(*AddReferencesItem); ok {
		return *casted
	}
	return nil
}

func (m *_AddReferencesItem) GetTypeName() string {
	return "AddReferencesItem"
}

func (m *_AddReferencesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (sourceNodeId)
	lengthInBits += m.SourceNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetServerUri)
	lengthInBits += m.TargetServerUri.GetLengthInBits(ctx)

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Simple field (targetNodeClass)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AddReferencesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AddReferencesItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__addReferencesItem AddReferencesItem, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AddReferencesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AddReferencesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sourceNodeId, err := ReadSimpleField[NodeId](ctx, "sourceNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceNodeId' field"))
	}
	m.SourceNodeId = sourceNodeId

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}
	m.ReferenceTypeId = referenceTypeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isForward, err := ReadSimpleField(ctx, "isForward", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isForward' field"))
	}
	m.IsForward = isForward

	targetServerUri, err := ReadSimpleField[PascalString](ctx, "targetServerUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetServerUri' field"))
	}
	m.TargetServerUri = targetServerUri

	targetNodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "targetNodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeId' field"))
	}
	m.TargetNodeId = targetNodeId

	targetNodeClass, err := ReadEnumField[NodeClass](ctx, "targetNodeClass", "NodeClass", ReadEnum(NodeClassByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeClass' field"))
	}
	m.TargetNodeClass = targetNodeClass

	if closeErr := readBuffer.CloseContext("AddReferencesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AddReferencesItem")
	}

	return m, nil
}

func (m *_AddReferencesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AddReferencesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AddReferencesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AddReferencesItem")
		}

		if err := WriteSimpleField[NodeId](ctx, "sourceNodeId", m.GetSourceNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceNodeId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isForward", m.GetIsForward(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isForward' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "targetServerUri", m.GetTargetServerUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetServerUri' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetNodeId", m.GetTargetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNodeId' field")
		}

		if err := WriteSimpleEnumField[NodeClass](ctx, "targetNodeClass", "NodeClass", m.GetTargetNodeClass(), WriteEnum[NodeClass, uint32](NodeClass.GetValue, NodeClass.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNodeClass' field")
		}

		if popErr := writeBuffer.PopContext("AddReferencesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AddReferencesItem")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AddReferencesItem) IsAddReferencesItem() {}

func (m *_AddReferencesItem) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
