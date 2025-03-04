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

// NodeIdTypeDefinition is the corresponding interface of NodeIdTypeDefinition
type NodeIdTypeDefinition interface {
	NodeIdTypeDefinitionContract
	NodeIdTypeDefinitionRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsNodeIdTypeDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdTypeDefinition()
	// CreateBuilder creates a NodeIdTypeDefinitionBuilder
	CreateNodeIdTypeDefinitionBuilder() NodeIdTypeDefinitionBuilder
}

// NodeIdTypeDefinitionContract provides a set of functions which can be overwritten by a sub struct
type NodeIdTypeDefinitionContract interface {
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
	// GetNamespace returns Namespace (abstract field)
	GetNamespace() int16
	// IsNodeIdTypeDefinition is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdTypeDefinition()
	// CreateBuilder creates a NodeIdTypeDefinitionBuilder
	CreateNodeIdTypeDefinitionBuilder() NodeIdTypeDefinitionBuilder
}

// NodeIdTypeDefinitionRequirements provides a set of functions which need to be implemented by a sub struct
type NodeIdTypeDefinitionRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetNodeType returns NodeType (discriminator field)
	GetNodeType() NodeIdType
	// GetIdentifier returns Identifier (abstract field)
	GetIdentifier() string
	// GetNamespace returns Namespace (abstract field)
	GetNamespace() int16
}

// _NodeIdTypeDefinition is the data-structure of this message
type _NodeIdTypeDefinition struct {
	_SubType interface {
		NodeIdTypeDefinitionContract
		NodeIdTypeDefinitionRequirements
	}
}

var _ NodeIdTypeDefinitionContract = (*_NodeIdTypeDefinition)(nil)

// NewNodeIdTypeDefinition factory function for _NodeIdTypeDefinition
func NewNodeIdTypeDefinition() *_NodeIdTypeDefinition {
	return &_NodeIdTypeDefinition{}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdTypeDefinitionBuilder is a builder for NodeIdTypeDefinition
type NodeIdTypeDefinitionBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() NodeIdTypeDefinitionBuilder
	// AsNodeIdTwoByte converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdTwoByte() NodeIdTwoByteBuilder
	// AsNodeIdFourByte converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdFourByte() NodeIdFourByteBuilder
	// AsNodeIdNumeric converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdNumeric() NodeIdNumericBuilder
	// AsNodeIdString converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdString() NodeIdStringBuilder
	// AsNodeIdGuid converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdGuid() NodeIdGuidBuilder
	// AsNodeIdByteString converts this build to a subType of NodeIdTypeDefinition. It is always possible to return to current builder using Done()
	AsNodeIdByteString() NodeIdByteStringBuilder
	// Build builds the NodeIdTypeDefinition or returns an error if something is wrong
	PartialBuild() (NodeIdTypeDefinitionContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() NodeIdTypeDefinitionContract
	// Build builds the NodeIdTypeDefinition or returns an error if something is wrong
	Build() (NodeIdTypeDefinition, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdTypeDefinition
}

// NewNodeIdTypeDefinitionBuilder() creates a NodeIdTypeDefinitionBuilder
func NewNodeIdTypeDefinitionBuilder() NodeIdTypeDefinitionBuilder {
	return &_NodeIdTypeDefinitionBuilder{_NodeIdTypeDefinition: new(_NodeIdTypeDefinition)}
}

type _NodeIdTypeDefinitionChildBuilder interface {
	utils.Copyable
	setParent(NodeIdTypeDefinitionContract)
	buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error)
}

type _NodeIdTypeDefinitionBuilder struct {
	*_NodeIdTypeDefinition

	childBuilder _NodeIdTypeDefinitionChildBuilder

	err *utils.MultiError
}

var _ (NodeIdTypeDefinitionBuilder) = (*_NodeIdTypeDefinitionBuilder)(nil)

func (b *_NodeIdTypeDefinitionBuilder) WithMandatoryFields() NodeIdTypeDefinitionBuilder {
	return b
}

func (b *_NodeIdTypeDefinitionBuilder) PartialBuild() (NodeIdTypeDefinitionContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdTypeDefinition.deepCopy(), nil
}

func (b *_NodeIdTypeDefinitionBuilder) PartialMustBuild() NodeIdTypeDefinitionContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdTwoByte() NodeIdTwoByteBuilder {
	if cb, ok := b.childBuilder.(NodeIdTwoByteBuilder); ok {
		return cb
	}
	cb := NewNodeIdTwoByteBuilder().(*_NodeIdTwoByteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdFourByte() NodeIdFourByteBuilder {
	if cb, ok := b.childBuilder.(NodeIdFourByteBuilder); ok {
		return cb
	}
	cb := NewNodeIdFourByteBuilder().(*_NodeIdFourByteBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdNumeric() NodeIdNumericBuilder {
	if cb, ok := b.childBuilder.(NodeIdNumericBuilder); ok {
		return cb
	}
	cb := NewNodeIdNumericBuilder().(*_NodeIdNumericBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdString() NodeIdStringBuilder {
	if cb, ok := b.childBuilder.(NodeIdStringBuilder); ok {
		return cb
	}
	cb := NewNodeIdStringBuilder().(*_NodeIdStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdGuid() NodeIdGuidBuilder {
	if cb, ok := b.childBuilder.(NodeIdGuidBuilder); ok {
		return cb
	}
	cb := NewNodeIdGuidBuilder().(*_NodeIdGuidBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) AsNodeIdByteString() NodeIdByteStringBuilder {
	if cb, ok := b.childBuilder.(NodeIdByteStringBuilder); ok {
		return cb
	}
	cb := NewNodeIdByteStringBuilder().(*_NodeIdByteStringBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_NodeIdTypeDefinitionBuilder) Build() (NodeIdTypeDefinition, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForNodeIdTypeDefinition()
}

func (b *_NodeIdTypeDefinitionBuilder) MustBuild() NodeIdTypeDefinition {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_NodeIdTypeDefinitionBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdTypeDefinitionBuilder().(*_NodeIdTypeDefinitionBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_NodeIdTypeDefinitionChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdTypeDefinitionBuilder creates a NodeIdTypeDefinitionBuilder
func (b *_NodeIdTypeDefinition) CreateNodeIdTypeDefinitionBuilder() NodeIdTypeDefinitionBuilder {
	if b == nil {
		return NewNodeIdTypeDefinitionBuilder()
	}
	return &_NodeIdTypeDefinitionBuilder{_NodeIdTypeDefinition: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for abstract fields.
///////////////////////

func (m *_NodeIdTypeDefinition) GetIdentifier() string {
	return m._SubType.GetIdentifier()
}

func (m *_NodeIdTypeDefinition) GetNamespace() int16 {
	return m._SubType.GetNamespace()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdTypeDefinition(structType any) NodeIdTypeDefinition {
	if casted, ok := structType.(NodeIdTypeDefinition); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdTypeDefinition); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdTypeDefinition) GetTypeName() string {
	return "NodeIdTypeDefinition"
}

func (m *_NodeIdTypeDefinition) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (nodeType)
	lengthInBits += 6

	return lengthInBits
}

func (m *_NodeIdTypeDefinition) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_NodeIdTypeDefinition) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func NodeIdTypeDefinitionParse[T NodeIdTypeDefinition](ctx context.Context, theBytes []byte) (T, error) {
	return NodeIdTypeDefinitionParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdTypeDefinitionParseWithBufferProducer[T NodeIdTypeDefinition]() func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := NodeIdTypeDefinitionParseWithBuffer[T](ctx, readBuffer)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func NodeIdTypeDefinitionParseWithBuffer[T NodeIdTypeDefinition](ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	v, err := (&_NodeIdTypeDefinition{}).parse(ctx, readBuffer)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_NodeIdTypeDefinition) parse(ctx context.Context, readBuffer utils.ReadBuffer) (__nodeIdTypeDefinition NodeIdTypeDefinition, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdTypeDefinition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdTypeDefinition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nodeType, err := ReadDiscriminatorEnumField[NodeIdType](ctx, "nodeType", "NodeIdType", ReadEnum(NodeIdTypeByValue, ReadUnsignedByte(readBuffer, uint8(6))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nodeType' field"))
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child NodeIdTypeDefinition
	switch {
	case nodeType == NodeIdType_nodeIdTypeTwoByte: // NodeIdTwoByte
		if _child, err = new(_NodeIdTwoByte).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdTwoByte for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeFourByte: // NodeIdFourByte
		if _child, err = new(_NodeIdFourByte).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdFourByte for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeNumeric: // NodeIdNumeric
		if _child, err = new(_NodeIdNumeric).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdNumeric for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeString: // NodeIdString
		if _child, err = new(_NodeIdString).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdString for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeGuid: // NodeIdGuid
		if _child, err = new(_NodeIdGuid).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdGuid for type-switch of NodeIdTypeDefinition")
		}
	case nodeType == NodeIdType_nodeIdTypeByteString: // NodeIdByteString
		if _child, err = new(_NodeIdByteString).parse(ctx, readBuffer, m); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type NodeIdByteString for type-switch of NodeIdTypeDefinition")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [nodeType=%v]", nodeType)
	}

	if closeErr := readBuffer.CloseContext("NodeIdTypeDefinition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdTypeDefinition")
	}

	return _child, nil
}

func (pm *_NodeIdTypeDefinition) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child NodeIdTypeDefinition, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NodeIdTypeDefinition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NodeIdTypeDefinition")
	}

	if err := WriteDiscriminatorEnumField(ctx, "nodeType", "NodeIdType", m.GetNodeType(), WriteEnum[NodeIdType, uint8](NodeIdType.GetValue, NodeIdType.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 6))); err != nil {
		return errors.Wrap(err, "Error serializing 'nodeType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("NodeIdTypeDefinition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NodeIdTypeDefinition")
	}
	return nil
}

func (m *_NodeIdTypeDefinition) IsNodeIdTypeDefinition() {}

func (m *_NodeIdTypeDefinition) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdTypeDefinition) deepCopy() *_NodeIdTypeDefinition {
	if m == nil {
		return nil
	}
	_NodeIdTypeDefinitionCopy := &_NodeIdTypeDefinition{
		nil, // will be set by child
	}
	return _NodeIdTypeDefinitionCopy
}
