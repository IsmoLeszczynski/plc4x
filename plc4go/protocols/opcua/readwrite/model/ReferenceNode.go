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

// ReferenceNode is the corresponding interface of ReferenceNode
type ReferenceNode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsInverse returns IsInverse (property field)
	GetIsInverse() bool
	// GetTargetId returns TargetId (property field)
	GetTargetId() ExpandedNodeId
}

// ReferenceNodeExactly can be used when we want exactly this type and not a type which fulfills ReferenceNode.
// This is useful for switch cases.
type ReferenceNodeExactly interface {
	ReferenceNode
	isReferenceNode() bool
}

// _ReferenceNode is the data-structure of this message
type _ReferenceNode struct {
	*_ExtensionObjectDefinition
	ReferenceTypeId NodeId
	IsInverse       bool
	TargetId        ExpandedNodeId
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ReferenceNode) GetIdentifier() string {
	return "287"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ReferenceNode) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ReferenceNode) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ReferenceNode) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_ReferenceNode) GetIsInverse() bool {
	return m.IsInverse
}

func (m *_ReferenceNode) GetTargetId() ExpandedNodeId {
	return m.TargetId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReferenceNode factory function for _ReferenceNode
func NewReferenceNode(referenceTypeId NodeId, isInverse bool, targetId ExpandedNodeId) *_ReferenceNode {
	_result := &_ReferenceNode{
		ReferenceTypeId:            referenceTypeId,
		IsInverse:                  isInverse,
		TargetId:                   targetId,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastReferenceNode(structType any) ReferenceNode {
	if casted, ok := structType.(ReferenceNode); ok {
		return casted
	}
	if casted, ok := structType.(*ReferenceNode); ok {
		return *casted
	}
	return nil
}

func (m *_ReferenceNode) GetTypeName() string {
	return "ReferenceNode"
}

func (m *_ReferenceNode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isInverse)
	lengthInBits += 1

	// Simple field (targetId)
	lengthInBits += m.TargetId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_ReferenceNode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ReferenceNodeParse(ctx context.Context, theBytes []byte, identifier string) (ReferenceNode, error) {
	return ReferenceNodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ReferenceNodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ReferenceNode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ReferenceNode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ReferenceNode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceTypeId)
	if pullErr := readBuffer.PullContext("referenceTypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTypeId")
	}
	_referenceTypeId, _referenceTypeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeIdErr != nil {
		return nil, errors.Wrap(_referenceTypeIdErr, "Error parsing 'referenceTypeId' field of ReferenceNode")
	}
	referenceTypeId := _referenceTypeId.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceTypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTypeId")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of ReferenceNode")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (isInverse)
	_isInverse, _isInverseErr := readBuffer.ReadBit("isInverse")
	if _isInverseErr != nil {
		return nil, errors.Wrap(_isInverseErr, "Error parsing 'isInverse' field of ReferenceNode")
	}
	isInverse := _isInverse

	// Simple Field (targetId)
	if pullErr := readBuffer.PullContext("targetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetId")
	}
	_targetId, _targetIdErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _targetIdErr != nil {
		return nil, errors.Wrap(_targetIdErr, "Error parsing 'targetId' field of ReferenceNode")
	}
	targetId := _targetId.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("targetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetId")
	}

	if closeErr := readBuffer.CloseContext("ReferenceNode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ReferenceNode")
	}

	// Create a partially initialized instance
	_child := &_ReferenceNode{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ReferenceTypeId:            referenceTypeId,
		IsInverse:                  isInverse,
		TargetId:                   targetId,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ReferenceNode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ReferenceNode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ReferenceNode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ReferenceNode")
		}

		// Simple Field (referenceTypeId)
		if pushErr := writeBuffer.PushContext("referenceTypeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceTypeId")
		}
		_referenceTypeIdErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceTypeId())
		if popErr := writeBuffer.PopContext("referenceTypeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceTypeId")
		}
		if _referenceTypeIdErr != nil {
			return errors.Wrap(_referenceTypeIdErr, "Error serializing 'referenceTypeId' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				log.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 7, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (isInverse)
		isInverse := bool(m.GetIsInverse())
		_isInverseErr := writeBuffer.WriteBit("isInverse", (isInverse))
		if _isInverseErr != nil {
			return errors.Wrap(_isInverseErr, "Error serializing 'isInverse' field")
		}

		// Simple Field (targetId)
		if pushErr := writeBuffer.PushContext("targetId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for targetId")
		}
		_targetIdErr := writeBuffer.WriteSerializable(ctx, m.GetTargetId())
		if popErr := writeBuffer.PopContext("targetId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for targetId")
		}
		if _targetIdErr != nil {
			return errors.Wrap(_targetIdErr, "Error serializing 'targetId' field")
		}

		if popErr := writeBuffer.PopContext("ReferenceNode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ReferenceNode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ReferenceNode) isReferenceNode() bool {
	return true
}

func (m *_ReferenceNode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
