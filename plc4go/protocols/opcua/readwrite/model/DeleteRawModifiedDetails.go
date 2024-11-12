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

// DeleteRawModifiedDetails is the corresponding interface of DeleteRawModifiedDetails
type DeleteRawModifiedDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetIsDeleteModified returns IsDeleteModified (property field)
	GetIsDeleteModified() bool
	// GetStartTime returns StartTime (property field)
	GetStartTime() int64
	// GetEndTime returns EndTime (property field)
	GetEndTime() int64
	// IsDeleteRawModifiedDetails is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteRawModifiedDetails()
	// CreateBuilder creates a DeleteRawModifiedDetailsBuilder
	CreateDeleteRawModifiedDetailsBuilder() DeleteRawModifiedDetailsBuilder
}

// _DeleteRawModifiedDetails is the data-structure of this message
type _DeleteRawModifiedDetails struct {
	ExtensionObjectDefinitionContract
	NodeId           NodeId
	IsDeleteModified bool
	StartTime        int64
	EndTime          int64
	// Reserved Fields
	reservedField0 *uint8
}

var _ DeleteRawModifiedDetails = (*_DeleteRawModifiedDetails)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteRawModifiedDetails)(nil)

// NewDeleteRawModifiedDetails factory function for _DeleteRawModifiedDetails
func NewDeleteRawModifiedDetails(nodeId NodeId, isDeleteModified bool, startTime int64, endTime int64) *_DeleteRawModifiedDetails {
	if nodeId == nil {
		panic("nodeId of type NodeId for DeleteRawModifiedDetails must not be nil")
	}
	_result := &_DeleteRawModifiedDetails{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		IsDeleteModified:                  isDeleteModified,
		StartTime:                         startTime,
		EndTime:                           endTime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteRawModifiedDetailsBuilder is a builder for DeleteRawModifiedDetails
type DeleteRawModifiedDetailsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeId NodeId, isDeleteModified bool, startTime int64, endTime int64) DeleteRawModifiedDetailsBuilder
	// WithNodeId adds NodeId (property field)
	WithNodeId(NodeId) DeleteRawModifiedDetailsBuilder
	// WithNodeIdBuilder adds NodeId (property field) which is build by the builder
	WithNodeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) DeleteRawModifiedDetailsBuilder
	// WithIsDeleteModified adds IsDeleteModified (property field)
	WithIsDeleteModified(bool) DeleteRawModifiedDetailsBuilder
	// WithStartTime adds StartTime (property field)
	WithStartTime(int64) DeleteRawModifiedDetailsBuilder
	// WithEndTime adds EndTime (property field)
	WithEndTime(int64) DeleteRawModifiedDetailsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DeleteRawModifiedDetails or returns an error if something is wrong
	Build() (DeleteRawModifiedDetails, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteRawModifiedDetails
}

// NewDeleteRawModifiedDetailsBuilder() creates a DeleteRawModifiedDetailsBuilder
func NewDeleteRawModifiedDetailsBuilder() DeleteRawModifiedDetailsBuilder {
	return &_DeleteRawModifiedDetailsBuilder{_DeleteRawModifiedDetails: new(_DeleteRawModifiedDetails)}
}

type _DeleteRawModifiedDetailsBuilder struct {
	*_DeleteRawModifiedDetails

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteRawModifiedDetailsBuilder) = (*_DeleteRawModifiedDetailsBuilder)(nil)

func (b *_DeleteRawModifiedDetailsBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteRawModifiedDetailsBuilder) WithMandatoryFields(nodeId NodeId, isDeleteModified bool, startTime int64, endTime int64) DeleteRawModifiedDetailsBuilder {
	return b.WithNodeId(nodeId).WithIsDeleteModified(isDeleteModified).WithStartTime(startTime).WithEndTime(endTime)
}

func (b *_DeleteRawModifiedDetailsBuilder) WithNodeId(nodeId NodeId) DeleteRawModifiedDetailsBuilder {
	b.NodeId = nodeId
	return b
}

func (b *_DeleteRawModifiedDetailsBuilder) WithNodeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) DeleteRawModifiedDetailsBuilder {
	builder := builderSupplier(b.NodeId.CreateNodeIdBuilder())
	var err error
	b.NodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_DeleteRawModifiedDetailsBuilder) WithIsDeleteModified(isDeleteModified bool) DeleteRawModifiedDetailsBuilder {
	b.IsDeleteModified = isDeleteModified
	return b
}

func (b *_DeleteRawModifiedDetailsBuilder) WithStartTime(startTime int64) DeleteRawModifiedDetailsBuilder {
	b.StartTime = startTime
	return b
}

func (b *_DeleteRawModifiedDetailsBuilder) WithEndTime(endTime int64) DeleteRawModifiedDetailsBuilder {
	b.EndTime = endTime
	return b
}

func (b *_DeleteRawModifiedDetailsBuilder) Build() (DeleteRawModifiedDetails, error) {
	if b.NodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteRawModifiedDetails.deepCopy(), nil
}

func (b *_DeleteRawModifiedDetailsBuilder) MustBuild() DeleteRawModifiedDetails {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeleteRawModifiedDetailsBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DeleteRawModifiedDetailsBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteRawModifiedDetailsBuilder) DeepCopy() any {
	_copy := b.CreateDeleteRawModifiedDetailsBuilder().(*_DeleteRawModifiedDetailsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteRawModifiedDetailsBuilder creates a DeleteRawModifiedDetailsBuilder
func (b *_DeleteRawModifiedDetails) CreateDeleteRawModifiedDetailsBuilder() DeleteRawModifiedDetailsBuilder {
	if b == nil {
		return NewDeleteRawModifiedDetailsBuilder()
	}
	return &_DeleteRawModifiedDetailsBuilder{_DeleteRawModifiedDetails: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteRawModifiedDetails) GetExtensionId() int32 {
	return int32(688)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteRawModifiedDetails) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteRawModifiedDetails) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_DeleteRawModifiedDetails) GetIsDeleteModified() bool {
	return m.IsDeleteModified
}

func (m *_DeleteRawModifiedDetails) GetStartTime() int64 {
	return m.StartTime
}

func (m *_DeleteRawModifiedDetails) GetEndTime() int64 {
	return m.EndTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteRawModifiedDetails(structType any) DeleteRawModifiedDetails {
	if casted, ok := structType.(DeleteRawModifiedDetails); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteRawModifiedDetails); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteRawModifiedDetails) GetTypeName() string {
	return "DeleteRawModifiedDetails"
}

func (m *_DeleteRawModifiedDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isDeleteModified)
	lengthInBits += 1

	// Simple field (startTime)
	lengthInBits += 64

	// Simple field (endTime)
	lengthInBits += 64

	return lengthInBits
}

func (m *_DeleteRawModifiedDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteRawModifiedDetails) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteRawModifiedDetails DeleteRawModifiedDetails, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteRawModifiedDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteRawModifiedDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isDeleteModified, err := ReadSimpleField(ctx, "isDeleteModified", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isDeleteModified' field"))
	}
	m.IsDeleteModified = isDeleteModified

	startTime, err := ReadSimpleField(ctx, "startTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'startTime' field"))
	}
	m.StartTime = startTime

	endTime, err := ReadSimpleField(ctx, "endTime", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endTime' field"))
	}
	m.EndTime = endTime

	if closeErr := readBuffer.CloseContext("DeleteRawModifiedDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteRawModifiedDetails")
	}

	return m, nil
}

func (m *_DeleteRawModifiedDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteRawModifiedDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteRawModifiedDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteRawModifiedDetails")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isDeleteModified", m.GetIsDeleteModified(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isDeleteModified' field")
		}

		if err := WriteSimpleField[int64](ctx, "startTime", m.GetStartTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'startTime' field")
		}

		if err := WriteSimpleField[int64](ctx, "endTime", m.GetEndTime(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'endTime' field")
		}

		if popErr := writeBuffer.PopContext("DeleteRawModifiedDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteRawModifiedDetails")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteRawModifiedDetails) IsDeleteRawModifiedDetails() {}

func (m *_DeleteRawModifiedDetails) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteRawModifiedDetails) deepCopy() *_DeleteRawModifiedDetails {
	if m == nil {
		return nil
	}
	_DeleteRawModifiedDetailsCopy := &_DeleteRawModifiedDetails{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.NodeId),
		m.IsDeleteModified,
		m.StartTime,
		m.EndTime,
		m.reservedField0,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteRawModifiedDetailsCopy
}

func (m *_DeleteRawModifiedDetails) String() string {
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
