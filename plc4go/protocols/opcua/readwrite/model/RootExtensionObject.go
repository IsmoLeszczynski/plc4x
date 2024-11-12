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

// RootExtensionObject is the corresponding interface of RootExtensionObject
type RootExtensionObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObject
	// GetBody returns Body (property field)
	GetBody() ExtensionObjectDefinition
	// IsRootExtensionObject is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRootExtensionObject()
	// CreateBuilder creates a RootExtensionObjectBuilder
	CreateRootExtensionObjectBuilder() RootExtensionObjectBuilder
}

// _RootExtensionObject is the data-structure of this message
type _RootExtensionObject struct {
	ExtensionObjectContract
	Body ExtensionObjectDefinition

	// Arguments.
	ExtensionId int32
}

var _ RootExtensionObject = (*_RootExtensionObject)(nil)
var _ ExtensionObjectRequirements = (*_RootExtensionObject)(nil)

// NewRootExtensionObject factory function for _RootExtensionObject
func NewRootExtensionObject(typeId ExpandedNodeId, body ExtensionObjectDefinition, extensionId int32) *_RootExtensionObject {
	if body == nil {
		panic("body of type ExtensionObjectDefinition for RootExtensionObject must not be nil")
	}
	_result := &_RootExtensionObject{
		ExtensionObjectContract: NewExtensionObject(typeId),
		Body:                    body,
	}
	_result.ExtensionObjectContract.(*_ExtensionObject)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RootExtensionObjectBuilder is a builder for RootExtensionObject
type RootExtensionObjectBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(body ExtensionObjectDefinition) RootExtensionObjectBuilder
	// WithBody adds Body (property field)
	WithBody(ExtensionObjectDefinition) RootExtensionObjectBuilder
	// WithBodyBuilder adds Body (property field) which is build by the builder
	WithBodyBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) RootExtensionObjectBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ExtensionObjectBuilder
	// Build builds the RootExtensionObject or returns an error if something is wrong
	Build() (RootExtensionObject, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RootExtensionObject
}

// NewRootExtensionObjectBuilder() creates a RootExtensionObjectBuilder
func NewRootExtensionObjectBuilder() RootExtensionObjectBuilder {
	return &_RootExtensionObjectBuilder{_RootExtensionObject: new(_RootExtensionObject)}
}

type _RootExtensionObjectBuilder struct {
	*_RootExtensionObject

	parentBuilder *_ExtensionObjectBuilder

	err *utils.MultiError
}

var _ (RootExtensionObjectBuilder) = (*_RootExtensionObjectBuilder)(nil)

func (b *_RootExtensionObjectBuilder) setParent(contract ExtensionObjectContract) {
	b.ExtensionObjectContract = contract
}

func (b *_RootExtensionObjectBuilder) WithMandatoryFields(body ExtensionObjectDefinition) RootExtensionObjectBuilder {
	return b.WithBody(body)
}

func (b *_RootExtensionObjectBuilder) WithBody(body ExtensionObjectDefinition) RootExtensionObjectBuilder {
	b.Body = body
	return b
}

func (b *_RootExtensionObjectBuilder) WithBodyBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) RootExtensionObjectBuilder {
	builder := builderSupplier(b.Body.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.Body, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_RootExtensionObjectBuilder) Build() (RootExtensionObject, error) {
	if b.Body == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'body' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RootExtensionObject.deepCopy(), nil
}

func (b *_RootExtensionObjectBuilder) MustBuild() RootExtensionObject {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_RootExtensionObjectBuilder) Done() ExtensionObjectBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewExtensionObjectBuilder().(*_ExtensionObjectBuilder)
	}
	return b.parentBuilder
}

func (b *_RootExtensionObjectBuilder) buildForExtensionObject() (ExtensionObject, error) {
	return b.Build()
}

func (b *_RootExtensionObjectBuilder) DeepCopy() any {
	_copy := b.CreateRootExtensionObjectBuilder().(*_RootExtensionObjectBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRootExtensionObjectBuilder creates a RootExtensionObjectBuilder
func (b *_RootExtensionObject) CreateRootExtensionObjectBuilder() RootExtensionObjectBuilder {
	if b == nil {
		return NewRootExtensionObjectBuilder()
	}
	return &_RootExtensionObjectBuilder{_RootExtensionObject: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RootExtensionObject) GetIncludeEncodingMask() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RootExtensionObject) GetParent() ExtensionObjectContract {
	return m.ExtensionObjectContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RootExtensionObject) GetBody() ExtensionObjectDefinition {
	return m.Body
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRootExtensionObject(structType any) RootExtensionObject {
	if casted, ok := structType.(RootExtensionObject); ok {
		return casted
	}
	if casted, ok := structType.(*RootExtensionObject); ok {
		return *casted
	}
	return nil
}

func (m *_RootExtensionObject) GetTypeName() string {
	return "RootExtensionObject"
}

func (m *_RootExtensionObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectContract.(*_ExtensionObject).getLengthInBits(ctx))

	// Simple field (body)
	lengthInBits += m.Body.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RootExtensionObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RootExtensionObject) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObject, extensionId int32, includeEncodingMask bool) (__rootExtensionObject RootExtensionObject, err error) {
	m.ExtensionObjectContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RootExtensionObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RootExtensionObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	body, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "body", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((int32)(extensionId)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'body' field"))
	}
	m.Body = body

	if closeErr := readBuffer.CloseContext("RootExtensionObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RootExtensionObject")
	}

	return m, nil
}

func (m *_RootExtensionObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RootExtensionObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RootExtensionObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RootExtensionObject")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "body", m.GetBody(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'body' field")
		}

		if popErr := writeBuffer.PopContext("RootExtensionObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RootExtensionObject")
		}
		return nil
	}
	return m.ExtensionObjectContract.(*_ExtensionObject).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_RootExtensionObject) GetExtensionId() int32 {
	return m.ExtensionId
}

//
////

func (m *_RootExtensionObject) IsRootExtensionObject() {}

func (m *_RootExtensionObject) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RootExtensionObject) deepCopy() *_RootExtensionObject {
	if m == nil {
		return nil
	}
	_RootExtensionObjectCopy := &_RootExtensionObject{
		m.ExtensionObjectContract.(*_ExtensionObject).deepCopy(),
		utils.DeepCopy[ExtensionObjectDefinition](m.Body),
		m.ExtensionId,
	}
	m.ExtensionObjectContract.(*_ExtensionObject)._SubType = m
	return _RootExtensionObjectCopy
}

func (m *_RootExtensionObject) String() string {
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
