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

// PubSubConfigurationDataType is the corresponding interface of PubSubConfigurationDataType
type PubSubConfigurationDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfPublishedDataSets returns NoOfPublishedDataSets (property field)
	GetNoOfPublishedDataSets() int32
	// GetPublishedDataSets returns PublishedDataSets (property field)
	GetPublishedDataSets() []ExtensionObjectDefinition
	// GetNoOfConnections returns NoOfConnections (property field)
	GetNoOfConnections() int32
	// GetConnections returns Connections (property field)
	GetConnections() []ExtensionObjectDefinition
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
}

// PubSubConfigurationDataTypeExactly can be used when we want exactly this type and not a type which fulfills PubSubConfigurationDataType.
// This is useful for switch cases.
type PubSubConfigurationDataTypeExactly interface {
	PubSubConfigurationDataType
	isPubSubConfigurationDataType() bool
}

// _PubSubConfigurationDataType is the data-structure of this message
type _PubSubConfigurationDataType struct {
	*_ExtensionObjectDefinition
	NoOfPublishedDataSets int32
	PublishedDataSets     []ExtensionObjectDefinition
	NoOfConnections       int32
	Connections           []ExtensionObjectDefinition
	Enabled               bool
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConfigurationDataType) GetIdentifier() string {
	return "15532"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConfigurationDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_PubSubConfigurationDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConfigurationDataType) GetNoOfPublishedDataSets() int32 {
	return m.NoOfPublishedDataSets
}

func (m *_PubSubConfigurationDataType) GetPublishedDataSets() []ExtensionObjectDefinition {
	return m.PublishedDataSets
}

func (m *_PubSubConfigurationDataType) GetNoOfConnections() int32 {
	return m.NoOfConnections
}

func (m *_PubSubConfigurationDataType) GetConnections() []ExtensionObjectDefinition {
	return m.Connections
}

func (m *_PubSubConfigurationDataType) GetEnabled() bool {
	return m.Enabled
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPubSubConfigurationDataType factory function for _PubSubConfigurationDataType
func NewPubSubConfigurationDataType(noOfPublishedDataSets int32, publishedDataSets []ExtensionObjectDefinition, noOfConnections int32, connections []ExtensionObjectDefinition, enabled bool) *_PubSubConfigurationDataType {
	_result := &_PubSubConfigurationDataType{
		NoOfPublishedDataSets:      noOfPublishedDataSets,
		PublishedDataSets:          publishedDataSets,
		NoOfConnections:            noOfConnections,
		Connections:                connections,
		Enabled:                    enabled,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPubSubConfigurationDataType(structType any) PubSubConfigurationDataType {
	if casted, ok := structType.(PubSubConfigurationDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConfigurationDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConfigurationDataType) GetTypeName() string {
	return "PubSubConfigurationDataType"
}

func (m *_PubSubConfigurationDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfPublishedDataSets)
	lengthInBits += 32

	// Array field
	if len(m.PublishedDataSets) > 0 {
		for _curItem, element := range m.PublishedDataSets {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.PublishedDataSets), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfConnections)
	lengthInBits += 32

	// Array field
	if len(m.Connections) > 0 {
		for _curItem, element := range m.Connections {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Connections), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	return lengthInBits
}

func (m *_PubSubConfigurationDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func PubSubConfigurationDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (PubSubConfigurationDataType, error) {
	return PubSubConfigurationDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func PubSubConfigurationDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (PubSubConfigurationDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("PubSubConfigurationDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConfigurationDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfPublishedDataSets)
	_noOfPublishedDataSets, _noOfPublishedDataSetsErr := readBuffer.ReadInt32("noOfPublishedDataSets", 32)
	if _noOfPublishedDataSetsErr != nil {
		return nil, errors.Wrap(_noOfPublishedDataSetsErr, "Error parsing 'noOfPublishedDataSets' field of PubSubConfigurationDataType")
	}
	noOfPublishedDataSets := _noOfPublishedDataSets

	// Array field (publishedDataSets)
	if pullErr := readBuffer.PullContext("publishedDataSets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for publishedDataSets")
	}
	// Count array
	publishedDataSets := make([]ExtensionObjectDefinition, max(noOfPublishedDataSets, 0))
	// This happens when the size is set conditional to 0
	if len(publishedDataSets) == 0 {
		publishedDataSets = nil
	}
	{
		_numItems := uint16(max(noOfPublishedDataSets, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "15580")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'publishedDataSets' field of PubSubConfigurationDataType")
			}
			publishedDataSets[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("publishedDataSets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for publishedDataSets")
	}

	// Simple Field (noOfConnections)
	_noOfConnections, _noOfConnectionsErr := readBuffer.ReadInt32("noOfConnections", 32)
	if _noOfConnectionsErr != nil {
		return nil, errors.Wrap(_noOfConnectionsErr, "Error parsing 'noOfConnections' field of PubSubConfigurationDataType")
	}
	noOfConnections := _noOfConnections

	// Array field (connections)
	if pullErr := readBuffer.PullContext("connections", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for connections")
	}
	// Count array
	connections := make([]ExtensionObjectDefinition, max(noOfConnections, 0))
	// This happens when the size is set conditional to 0
	if len(connections) == 0 {
		connections = nil
	}
	{
		_numItems := uint16(max(noOfConnections, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "15619")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'connections' field of PubSubConfigurationDataType")
			}
			connections[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("connections", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for connections")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of PubSubConfigurationDataType")
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

	// Simple Field (enabled)
	_enabled, _enabledErr := readBuffer.ReadBit("enabled")
	if _enabledErr != nil {
		return nil, errors.Wrap(_enabledErr, "Error parsing 'enabled' field of PubSubConfigurationDataType")
	}
	enabled := _enabled

	if closeErr := readBuffer.CloseContext("PubSubConfigurationDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConfigurationDataType")
	}

	// Create a partially initialized instance
	_child := &_PubSubConfigurationDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfPublishedDataSets:      noOfPublishedDataSets,
		PublishedDataSets:          publishedDataSets,
		NoOfConnections:            noOfConnections,
		Connections:                connections,
		Enabled:                    enabled,
		reservedField0:             reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_PubSubConfigurationDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConfigurationDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConfigurationDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConfigurationDataType")
		}

		// Simple Field (noOfPublishedDataSets)
		noOfPublishedDataSets := int32(m.GetNoOfPublishedDataSets())
		_noOfPublishedDataSetsErr := writeBuffer.WriteInt32("noOfPublishedDataSets", 32, int32((noOfPublishedDataSets)))
		if _noOfPublishedDataSetsErr != nil {
			return errors.Wrap(_noOfPublishedDataSetsErr, "Error serializing 'noOfPublishedDataSets' field")
		}

		// Array Field (publishedDataSets)
		if pushErr := writeBuffer.PushContext("publishedDataSets", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for publishedDataSets")
		}
		for _curItem, _element := range m.GetPublishedDataSets() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetPublishedDataSets()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'publishedDataSets' field")
			}
		}
		if popErr := writeBuffer.PopContext("publishedDataSets", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for publishedDataSets")
		}

		// Simple Field (noOfConnections)
		noOfConnections := int32(m.GetNoOfConnections())
		_noOfConnectionsErr := writeBuffer.WriteInt32("noOfConnections", 32, int32((noOfConnections)))
		if _noOfConnectionsErr != nil {
			return errors.Wrap(_noOfConnectionsErr, "Error serializing 'noOfConnections' field")
		}

		// Array Field (connections)
		if pushErr := writeBuffer.PushContext("connections", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for connections")
		}
		for _curItem, _element := range m.GetConnections() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetConnections()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'connections' field")
			}
		}
		if popErr := writeBuffer.PopContext("connections", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for connections")
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
			_err := writeBuffer.WriteUint8("reserved", 7, uint8(reserved))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (enabled)
		enabled := bool(m.GetEnabled())
		_enabledErr := writeBuffer.WriteBit("enabled", (enabled))
		if _enabledErr != nil {
			return errors.Wrap(_enabledErr, "Error serializing 'enabled' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConfigurationDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConfigurationDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConfigurationDataType) isPubSubConfigurationDataType() bool {
	return true
}

func (m *_PubSubConfigurationDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
