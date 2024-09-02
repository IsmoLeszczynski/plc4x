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

// QueryDataDescription is the corresponding interface of QueryDataDescription
type QueryDataDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRelativePath returns RelativePath (property field)
	GetRelativePath() ExtensionObjectDefinition
	// GetAttributeId returns AttributeId (property field)
	GetAttributeId() uint32
	// GetIndexRange returns IndexRange (property field)
	GetIndexRange() PascalString
}

// _QueryDataDescription is the data-structure of this message
type _QueryDataDescription struct {
	ExtensionObjectDefinitionContract
	RelativePath ExtensionObjectDefinition
	AttributeId  uint32
	IndexRange   PascalString
}

var _ QueryDataDescription = (*_QueryDataDescription)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_QueryDataDescription)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryDataDescription) GetIdentifier() string {
	return "572"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryDataDescription) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryDataDescription) GetRelativePath() ExtensionObjectDefinition {
	return m.RelativePath
}

func (m *_QueryDataDescription) GetAttributeId() uint32 {
	return m.AttributeId
}

func (m *_QueryDataDescription) GetIndexRange() PascalString {
	return m.IndexRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryDataDescription factory function for _QueryDataDescription
func NewQueryDataDescription(relativePath ExtensionObjectDefinition, attributeId uint32, indexRange PascalString) *_QueryDataDescription {
	_result := &_QueryDataDescription{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RelativePath:                      relativePath,
		AttributeId:                       attributeId,
		IndexRange:                        indexRange,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryDataDescription(structType any) QueryDataDescription {
	if casted, ok := structType.(QueryDataDescription); ok {
		return casted
	}
	if casted, ok := structType.(*QueryDataDescription); ok {
		return *casted
	}
	return nil
}

func (m *_QueryDataDescription) GetTypeName() string {
	return "QueryDataDescription"
}

func (m *_QueryDataDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (relativePath)
	lengthInBits += m.RelativePath.GetLengthInBits(ctx)

	// Simple field (attributeId)
	lengthInBits += 32

	// Simple field (indexRange)
	lengthInBits += m.IndexRange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_QueryDataDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_QueryDataDescription) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__queryDataDescription QueryDataDescription, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QueryDataDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryDataDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	relativePath, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "relativePath", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("542")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'relativePath' field"))
	}
	m.RelativePath = relativePath

	attributeId, err := ReadSimpleField(ctx, "attributeId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attributeId' field"))
	}
	m.AttributeId = attributeId

	indexRange, err := ReadSimpleField[PascalString](ctx, "indexRange", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'indexRange' field"))
	}
	m.IndexRange = indexRange

	if closeErr := readBuffer.CloseContext("QueryDataDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryDataDescription")
	}

	return m, nil
}

func (m *_QueryDataDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryDataDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryDataDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryDataDescription")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "relativePath", m.GetRelativePath(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'relativePath' field")
		}

		if err := WriteSimpleField[uint32](ctx, "attributeId", m.GetAttributeId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'attributeId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "indexRange", m.GetIndexRange(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'indexRange' field")
		}

		if popErr := writeBuffer.PopContext("QueryDataDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryDataDescription")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryDataDescription) IsQueryDataDescription() {}

func (m *_QueryDataDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
