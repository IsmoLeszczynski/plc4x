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

// BrowseDescription is the corresponding interface of BrowseDescription
type BrowseDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetBrowseDirection returns BrowseDirection (property field)
	GetBrowseDirection() BrowseDirection
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIncludeSubtypes returns IncludeSubtypes (property field)
	GetIncludeSubtypes() bool
	// GetNodeClassMask returns NodeClassMask (property field)
	GetNodeClassMask() uint32
	// GetResultMask returns ResultMask (property field)
	GetResultMask() uint32
}

// BrowseDescriptionExactly can be used when we want exactly this type and not a type which fulfills BrowseDescription.
// This is useful for switch cases.
type BrowseDescriptionExactly interface {
	BrowseDescription
	isBrowseDescription() bool
}

// _BrowseDescription is the data-structure of this message
type _BrowseDescription struct {
	ExtensionObjectDefinitionContract
	NodeId          NodeId
	BrowseDirection BrowseDirection
	ReferenceTypeId NodeId
	IncludeSubtypes bool
	NodeClassMask   uint32
	ResultMask      uint32
	// Reserved Fields
	reservedField0 *uint8
}

var _ BrowseDescription = (*_BrowseDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_BrowseDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BrowseDescription) GetIdentifier() string {
	return "516"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BrowseDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BrowseDescription) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_BrowseDescription) GetBrowseDirection() BrowseDirection {
	return m.BrowseDirection
}

func (m *_BrowseDescription) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_BrowseDescription) GetIncludeSubtypes() bool {
	return m.IncludeSubtypes
}

func (m *_BrowseDescription) GetNodeClassMask() uint32 {
	return m.NodeClassMask
}

func (m *_BrowseDescription) GetResultMask() uint32 {
	return m.ResultMask
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBrowseDescription factory function for _BrowseDescription
func NewBrowseDescription(nodeId NodeId, browseDirection BrowseDirection, referenceTypeId NodeId, includeSubtypes bool, nodeClassMask uint32, resultMask uint32) *_BrowseDescription {
	_result := &_BrowseDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		BrowseDirection:                   browseDirection,
		ReferenceTypeId:                   referenceTypeId,
		IncludeSubtypes:                   includeSubtypes,
		NodeClassMask:                     nodeClassMask,
		ResultMask:                        resultMask,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBrowseDescription(structType any) BrowseDescription {
	if casted, ok := structType.(BrowseDescription); ok {
		return casted
	}
	if casted, ok := structType.(*BrowseDescription); ok {
		return *casted
	}
	return nil
}

func (m *_BrowseDescription) GetTypeName() string {
	return "BrowseDescription"
}

func (m *_BrowseDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (browseDirection)
	lengthInBits += 32

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (includeSubtypes)
	lengthInBits += 1

	// Simple field (nodeClassMask)
	lengthInBits += 32

	// Simple field (resultMask)
	lengthInBits += 32

	return lengthInBits
}

func (m *_BrowseDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BrowseDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__browseDescription BrowseDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BrowseDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BrowseDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	browseDirection, err := ReadEnumField[BrowseDirection](ctx, "browseDirection", "BrowseDirection", ReadEnum(BrowseDirectionByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'browseDirection' field"))
	}
	m.BrowseDirection = browseDirection

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

	includeSubtypes, err := ReadSimpleField(ctx, "includeSubtypes", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'includeSubtypes' field"))
	}
	m.IncludeSubtypes = includeSubtypes

	nodeClassMask, err := ReadSimpleField(ctx, "nodeClassMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeClassMask' field"))
	}
	m.NodeClassMask = nodeClassMask

	resultMask, err := ReadSimpleField(ctx, "resultMask", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'resultMask' field"))
	}
	m.ResultMask = resultMask

	if closeErr := readBuffer.CloseContext("BrowseDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BrowseDescription")
	}

	return m, nil
}

func (m *_BrowseDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BrowseDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BrowseDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BrowseDescription")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleEnumField[BrowseDirection](ctx, "browseDirection", "BrowseDirection", m.GetBrowseDirection(), WriteEnum[BrowseDirection, uint32](BrowseDirection.GetValue, BrowseDirection.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'browseDirection' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "includeSubtypes", m.GetIncludeSubtypes(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'includeSubtypes' field")
		}

		if err := WriteSimpleField[uint32](ctx, "nodeClassMask", m.GetNodeClassMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeClassMask' field")
		}

		if err := WriteSimpleField[uint32](ctx, "resultMask", m.GetResultMask(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'resultMask' field")
		}

		if popErr := writeBuffer.PopContext("BrowseDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BrowseDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BrowseDescription) isBrowseDescription() bool {
	return true
}

func (m *_BrowseDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
