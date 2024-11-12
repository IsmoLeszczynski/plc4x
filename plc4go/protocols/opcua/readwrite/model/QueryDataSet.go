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

// QueryDataSet is the corresponding interface of QueryDataSet
type QueryDataSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() ExpandedNodeId
	// GetTypeDefinitionNode returns TypeDefinitionNode (property field)
	GetTypeDefinitionNode() ExpandedNodeId
	// GetValues returns Values (property field)
	GetValues() []Variant
	// IsQueryDataSet is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsQueryDataSet()
	// CreateBuilder creates a QueryDataSetBuilder
	CreateQueryDataSetBuilder() QueryDataSetBuilder
}

// _QueryDataSet is the data-structure of this message
type _QueryDataSet struct {
	ExtensionObjectDefinitionContract
	NodeId             ExpandedNodeId
	TypeDefinitionNode ExpandedNodeId
	Values             []Variant
}

var _ QueryDataSet = (*_QueryDataSet)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_QueryDataSet)(nil)

// NewQueryDataSet factory function for _QueryDataSet
func NewQueryDataSet(nodeId ExpandedNodeId, typeDefinitionNode ExpandedNodeId, values []Variant) *_QueryDataSet {
	if nodeId == nil {
		panic("nodeId of type ExpandedNodeId for QueryDataSet must not be nil")
	}
	if typeDefinitionNode == nil {
		panic("typeDefinitionNode of type ExpandedNodeId for QueryDataSet must not be nil")
	}
	_result := &_QueryDataSet{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		NodeId:                            nodeId,
		TypeDefinitionNode:                typeDefinitionNode,
		Values:                            values,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// QueryDataSetBuilder is a builder for QueryDataSet
type QueryDataSetBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nodeId ExpandedNodeId, typeDefinitionNode ExpandedNodeId, values []Variant) QueryDataSetBuilder
	// WithNodeId adds NodeId (property field)
	WithNodeId(ExpandedNodeId) QueryDataSetBuilder
	// WithNodeIdBuilder adds NodeId (property field) which is build by the builder
	WithNodeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) QueryDataSetBuilder
	// WithTypeDefinitionNode adds TypeDefinitionNode (property field)
	WithTypeDefinitionNode(ExpandedNodeId) QueryDataSetBuilder
	// WithTypeDefinitionNodeBuilder adds TypeDefinitionNode (property field) which is build by the builder
	WithTypeDefinitionNodeBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) QueryDataSetBuilder
	// WithValues adds Values (property field)
	WithValues(...Variant) QueryDataSetBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectDefinitionBuilder
	// Build builds the QueryDataSet or returns an error if something is wrong
	Build() (QueryDataSet, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() QueryDataSet
}

// NewQueryDataSetBuilder() creates a QueryDataSetBuilder
func NewQueryDataSetBuilder() QueryDataSetBuilder {
	return &_QueryDataSetBuilder{_QueryDataSet: new(_QueryDataSet)}
}

type _QueryDataSetBuilder struct {
	*_QueryDataSet

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (QueryDataSetBuilder) = (*_QueryDataSetBuilder)(nil)

func (b *_QueryDataSetBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_QueryDataSetBuilder) WithMandatoryFields(nodeId ExpandedNodeId, typeDefinitionNode ExpandedNodeId, values []Variant) QueryDataSetBuilder {
	return b.WithNodeId(nodeId).WithTypeDefinitionNode(typeDefinitionNode).WithValues(values...)
}

func (b *_QueryDataSetBuilder) WithNodeId(nodeId ExpandedNodeId) QueryDataSetBuilder {
	b.NodeId = nodeId
	return b
}

func (b *_QueryDataSetBuilder) WithNodeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) QueryDataSetBuilder {
	builder := builderSupplier(b.NodeId.CreateExpandedNodeIdBuilder())
	var err error
	b.NodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_QueryDataSetBuilder) WithTypeDefinitionNode(typeDefinitionNode ExpandedNodeId) QueryDataSetBuilder {
	b.TypeDefinitionNode = typeDefinitionNode
	return b
}

func (b *_QueryDataSetBuilder) WithTypeDefinitionNodeBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) QueryDataSetBuilder {
	builder := builderSupplier(b.TypeDefinitionNode.CreateExpandedNodeIdBuilder())
	var err error
	b.TypeDefinitionNode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_QueryDataSetBuilder) WithValues(values ...Variant) QueryDataSetBuilder {
	b.Values = values
	return b
}

func (b *_QueryDataSetBuilder) Build() (QueryDataSet, error) {
	if b.NodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nodeId' not set"))
	}
	if b.TypeDefinitionNode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'typeDefinitionNode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._QueryDataSet.deepCopy(), nil
}

func (b *_QueryDataSetBuilder) MustBuild() QueryDataSet {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_QueryDataSetBuilder) Done() ExtensionObjectDefinitionBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectDefinitionBuilder().(*_ExtensionObjectDefinitionBuilder)
	}
	return b.parentBuilder
}

func (b *_QueryDataSetBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_QueryDataSetBuilder) DeepCopy() any {
	_copy := b.CreateQueryDataSetBuilder().(*_QueryDataSetBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateQueryDataSetBuilder creates a QueryDataSetBuilder
func (b *_QueryDataSet) CreateQueryDataSetBuilder() QueryDataSetBuilder {
	if b == nil {
		return NewQueryDataSetBuilder()
	}
	return &_QueryDataSetBuilder{_QueryDataSet: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryDataSet) GetExtensionId() int32 {
	return int32(579)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryDataSet) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryDataSet) GetNodeId() ExpandedNodeId {
	return m.NodeId
}

func (m *_QueryDataSet) GetTypeDefinitionNode() ExpandedNodeId {
	return m.TypeDefinitionNode
}

func (m *_QueryDataSet) GetValues() []Variant {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastQueryDataSet(structType any) QueryDataSet {
	if casted, ok := structType.(QueryDataSet); ok {
		return casted
	}
	if casted, ok := structType.(*QueryDataSet); ok {
		return *casted
	}
	return nil
}

func (m *_QueryDataSet) GetTypeName() string {
	return "QueryDataSet"
}

func (m *_QueryDataSet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (typeDefinitionNode)
	lengthInBits += m.TypeDefinitionNode.GetLengthInBits(ctx)

	// Implicit Field (noOfValues)
	lengthInBits += 32

	// Array field
	if len(m.Values) > 0 {
		for _curItem, element := range m.Values {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Values), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_QueryDataSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_QueryDataSet) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__queryDataSet QueryDataSet, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("QueryDataSet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryDataSet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "nodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeId' field"))
	}
	m.NodeId = nodeId

	typeDefinitionNode, err := ReadSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'typeDefinitionNode' field"))
	}
	m.TypeDefinitionNode = typeDefinitionNode

	noOfValues, err := ReadImplicitField[int32](ctx, "noOfValues", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfValues' field"))
	}
	_ = noOfValues

	values, err := ReadCountArrayField[Variant](ctx, "values", ReadComplex[Variant](VariantParseWithBuffer, readBuffer), uint64(noOfValues))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'values' field"))
	}
	m.Values = values

	if closeErr := readBuffer.CloseContext("QueryDataSet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryDataSet")
	}

	return m, nil
}

func (m *_QueryDataSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryDataSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryDataSet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryDataSet")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "nodeId", m.GetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nodeId' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "typeDefinitionNode", m.GetTypeDefinitionNode(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'typeDefinitionNode' field")
		}
		noOfValues := int32(utils.InlineIf(bool((m.GetValues()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetValues()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfValues", noOfValues, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfValues' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "values", m.GetValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'values' field")
		}

		if popErr := writeBuffer.PopContext("QueryDataSet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryDataSet")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryDataSet) IsQueryDataSet() {}

func (m *_QueryDataSet) DeepCopy() any {
	return m.deepCopy()
}

func (m *_QueryDataSet) deepCopy() *_QueryDataSet {
	if m == nil {
		return nil
	}
	_QueryDataSetCopy := &_QueryDataSet{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[ExpandedNodeId](m.NodeId),
		utils.DeepCopy[ExpandedNodeId](m.TypeDefinitionNode),
		utils.DeepCopySlice[Variant, Variant](m.Values),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _QueryDataSetCopy
}

func (m *_QueryDataSet) String() string {
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
