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

// DeleteAtTimeDetails is the corresponding interface of DeleteAtTimeDetails
type DeleteAtTimeDetails interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetReqTimes returns ReqTimes (property field)
	GetReqTimes() []int64
	// IsDeleteAtTimeDetails is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteAtTimeDetails()
	// CreateBuilder creates a DeleteAtTimeDetailsBuilder
	CreateDeleteAtTimeDetailsBuilder() DeleteAtTimeDetailsBuilder
}

// _DeleteAtTimeDetails is the data-structure of this message
type _DeleteAtTimeDetails struct {
	ExtensionObjectDefinitionContract
	NodeId   NodeId
	ReqTimes []int64
}

var _ DeleteAtTimeDetails = (*_DeleteAtTimeDetails)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteAtTimeDetails)(nil)

// NewDeleteAtTimeDetails factory function for _DeleteAtTimeDetails
func NewDeleteAtTimeDetails(nodeId NodeId, reqTimes []int64) *_DeleteAtTimeDetails {
	if nodeId == nil {
		panic("nodeId of type NodeId for DeleteAtTimeDetails must not be nil")
	}
	_result := &_DeleteAtTimeDetails{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		ReqTimes:                          reqTimes,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteAtTimeDetailsBuilder is a builder for DeleteAtTimeDetails
type DeleteAtTimeDetailsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeId NodeId, reqTimes []int64) DeleteAtTimeDetailsBuilder
	// WithNodeId adds NodeId (property field)
	WithNodeId(NodeId) DeleteAtTimeDetailsBuilder
	// WithNodeIdBuilder adds NodeId (property field) which is build by the builder
	WithNodeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) DeleteAtTimeDetailsBuilder
	// WithReqTimes adds ReqTimes (property field)
	WithReqTimes(...int64) DeleteAtTimeDetailsBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the DeleteAtTimeDetails or returns an error if something is wrong
	Build() (DeleteAtTimeDetails, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteAtTimeDetails
}

// NewDeleteAtTimeDetailsBuilder() creates a DeleteAtTimeDetailsBuilder
func NewDeleteAtTimeDetailsBuilder() DeleteAtTimeDetailsBuilder {
	return &_DeleteAtTimeDetailsBuilder{_DeleteAtTimeDetails: new(_DeleteAtTimeDetails)}
}

type _DeleteAtTimeDetailsBuilder struct {
	*_DeleteAtTimeDetails

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteAtTimeDetailsBuilder) = (*_DeleteAtTimeDetailsBuilder)(nil)

func (b *_DeleteAtTimeDetailsBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteAtTimeDetailsBuilder) WithMandatoryFields(nodeId NodeId, reqTimes []int64) DeleteAtTimeDetailsBuilder {
	return b.WithNodeId(nodeId).WithReqTimes(reqTimes...)
}

func (b *_DeleteAtTimeDetailsBuilder) WithNodeId(nodeId NodeId) DeleteAtTimeDetailsBuilder {
	b.NodeId = nodeId
	return b
}

func (b *_DeleteAtTimeDetailsBuilder) WithNodeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) DeleteAtTimeDetailsBuilder {
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

func (b *_DeleteAtTimeDetailsBuilder) WithReqTimes(reqTimes ...int64) DeleteAtTimeDetailsBuilder {
	b.ReqTimes = reqTimes
	return b
}

func (b *_DeleteAtTimeDetailsBuilder) Build() (DeleteAtTimeDetails, error) {
	if b.NodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteAtTimeDetails.deepCopy(), nil
}

func (b *_DeleteAtTimeDetailsBuilder) MustBuild() DeleteAtTimeDetails {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DeleteAtTimeDetailsBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_DeleteAtTimeDetailsBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteAtTimeDetailsBuilder) DeepCopy() any {
	_copy := b.CreateDeleteAtTimeDetailsBuilder().(*_DeleteAtTimeDetailsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteAtTimeDetailsBuilder creates a DeleteAtTimeDetailsBuilder
func (b *_DeleteAtTimeDetails) CreateDeleteAtTimeDetailsBuilder() DeleteAtTimeDetailsBuilder {
	if b == nil {
		return NewDeleteAtTimeDetailsBuilder()
	}
	return &_DeleteAtTimeDetailsBuilder{_DeleteAtTimeDetails: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteAtTimeDetails) GetExtensionId() int32 {
	return int32(691)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteAtTimeDetails) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteAtTimeDetails) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_DeleteAtTimeDetails) GetReqTimes() []int64 {
	return m.ReqTimes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteAtTimeDetails(structType any) DeleteAtTimeDetails {
	if casted, ok := structType.(DeleteAtTimeDetails); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteAtTimeDetails); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteAtTimeDetails) GetTypeName() string {
	return "DeleteAtTimeDetails"
}

func (m *_DeleteAtTimeDetails) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Implicit Field (noOfReqTimes)
	lengthInBits += 32

	// Array field
	if len(m.ReqTimes) > 0 {
		lengthInBits += 64 * uint16(len(m.ReqTimes))
	}

	return lengthInBits
}

func (m *_DeleteAtTimeDetails) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteAtTimeDetails) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__deleteAtTimeDetails DeleteAtTimeDetails, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteAtTimeDetails"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteAtTimeDetails")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[NodeId](ctx, "nodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	noOfReqTimes, err := ReadImplicitField[int32](ctx, "noOfReqTimes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReqTimes' field"))
	}
	_ = noOfReqTimes

	reqTimes, err := ReadCountArrayField[int64](ctx, "reqTimes", ReadSignedLong(readBuffer, uint8(64)), uint64(noOfReqTimes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reqTimes' field"))
	}
	m.ReqTimes = reqTimes

	if closeErr := readBuffer.CloseContext("DeleteAtTimeDetails"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteAtTimeDetails")
	}

	return m, nil
}

func (m *_DeleteAtTimeDetails) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteAtTimeDetails) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteAtTimeDetails"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteAtTimeDetails")
		}

		if err := WriteSimpleField[NodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}
		noOfReqTimes := int32(utils.InlineIf(bool((m.GetReqTimes()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetReqTimes()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfReqTimes", noOfReqTimes, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReqTimes' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "reqTimes", m.GetReqTimes(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'reqTimes' field")
		}

		if popErr := writeBuffer.PopContext("DeleteAtTimeDetails"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteAtTimeDetails")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteAtTimeDetails) IsDeleteAtTimeDetails() {}

func (m *_DeleteAtTimeDetails) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteAtTimeDetails) deepCopy() *_DeleteAtTimeDetails {
	if m == nil {
		return nil
	}
	_DeleteAtTimeDetailsCopy := &_DeleteAtTimeDetails{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.NodeId),
		utils.DeepCopySlice[int64, int64](m.ReqTimes),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteAtTimeDetailsCopy
}

func (m *_DeleteAtTimeDetails) String() string {
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
