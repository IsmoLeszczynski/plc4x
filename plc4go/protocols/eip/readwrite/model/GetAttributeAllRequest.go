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

// GetAttributeAllRequest is the corresponding interface of GetAttributeAllRequest
type GetAttributeAllRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// IsGetAttributeAllRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetAttributeAllRequest()
	// CreateBuilder creates a GetAttributeAllRequestBuilder
	CreateGetAttributeAllRequestBuilder() GetAttributeAllRequestBuilder
}

// _GetAttributeAllRequest is the data-structure of this message
type _GetAttributeAllRequest struct {
	CipServiceContract
	ClassSegment    PathSegment
	InstanceSegment PathSegment
}

var _ GetAttributeAllRequest = (*_GetAttributeAllRequest)(nil)
var _ CipServiceRequirements = (*_GetAttributeAllRequest)(nil)

// NewGetAttributeAllRequest factory function for _GetAttributeAllRequest
func NewGetAttributeAllRequest(classSegment PathSegment, instanceSegment PathSegment, serviceLen uint16) *_GetAttributeAllRequest {
	if classSegment == nil {
		panic("classSegment of type PathSegment for GetAttributeAllRequest must not be nil")
	}
	if instanceSegment == nil {
		panic("instanceSegment of type PathSegment for GetAttributeAllRequest must not be nil")
	}
	_result := &_GetAttributeAllRequest{
		CipServiceContract: NewCipService(serviceLen),
		ClassSegment:       classSegment,
		InstanceSegment:    instanceSegment,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetAttributeAllRequestBuilder is a builder for GetAttributeAllRequest
type GetAttributeAllRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(classSegment PathSegment, instanceSegment PathSegment) GetAttributeAllRequestBuilder
	// WithClassSegment adds ClassSegment (property field)
	WithClassSegment(PathSegment) GetAttributeAllRequestBuilder
	// WithClassSegmentBuilder adds ClassSegment (property field) which is build by the builder
	WithClassSegmentBuilder(func(PathSegmentBuilder) PathSegmentBuilder) GetAttributeAllRequestBuilder
	// WithInstanceSegment adds InstanceSegment (property field)
	WithInstanceSegment(PathSegment) GetAttributeAllRequestBuilder
	// WithInstanceSegmentBuilder adds InstanceSegment (property field) which is build by the builder
	WithInstanceSegmentBuilder(func(PathSegmentBuilder) PathSegmentBuilder) GetAttributeAllRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the GetAttributeAllRequest or returns an error if something is wrong
	Build() (GetAttributeAllRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetAttributeAllRequest
}

// NewGetAttributeAllRequestBuilder() creates a GetAttributeAllRequestBuilder
func NewGetAttributeAllRequestBuilder() GetAttributeAllRequestBuilder {
	return &_GetAttributeAllRequestBuilder{_GetAttributeAllRequest: new(_GetAttributeAllRequest)}
}

type _GetAttributeAllRequestBuilder struct {
	*_GetAttributeAllRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (GetAttributeAllRequestBuilder) = (*_GetAttributeAllRequestBuilder)(nil)

func (b *_GetAttributeAllRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
	contract.(*_CipService)._SubType = b._GetAttributeAllRequest
}

func (b *_GetAttributeAllRequestBuilder) WithMandatoryFields(classSegment PathSegment, instanceSegment PathSegment) GetAttributeAllRequestBuilder {
	return b.WithClassSegment(classSegment).WithInstanceSegment(instanceSegment)
}

func (b *_GetAttributeAllRequestBuilder) WithClassSegment(classSegment PathSegment) GetAttributeAllRequestBuilder {
	b.ClassSegment = classSegment
	return b
}

func (b *_GetAttributeAllRequestBuilder) WithClassSegmentBuilder(builderSupplier func(PathSegmentBuilder) PathSegmentBuilder) GetAttributeAllRequestBuilder {
	builder := builderSupplier(b.ClassSegment.CreatePathSegmentBuilder())
	var err error
	b.ClassSegment, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PathSegmentBuilder failed"))
	}
	return b
}

func (b *_GetAttributeAllRequestBuilder) WithInstanceSegment(instanceSegment PathSegment) GetAttributeAllRequestBuilder {
	b.InstanceSegment = instanceSegment
	return b
}

func (b *_GetAttributeAllRequestBuilder) WithInstanceSegmentBuilder(builderSupplier func(PathSegmentBuilder) PathSegmentBuilder) GetAttributeAllRequestBuilder {
	builder := builderSupplier(b.InstanceSegment.CreatePathSegmentBuilder())
	var err error
	b.InstanceSegment, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PathSegmentBuilder failed"))
	}
	return b
}

func (b *_GetAttributeAllRequestBuilder) Build() (GetAttributeAllRequest, error) {
	if b.ClassSegment == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'classSegment' not set"))
	}
	if b.InstanceSegment == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'instanceSegment' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GetAttributeAllRequest.deepCopy(), nil
}

func (b *_GetAttributeAllRequestBuilder) MustBuild() GetAttributeAllRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_GetAttributeAllRequestBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_GetAttributeAllRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_GetAttributeAllRequestBuilder) DeepCopy() any {
	_copy := b.CreateGetAttributeAllRequestBuilder().(*_GetAttributeAllRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGetAttributeAllRequestBuilder creates a GetAttributeAllRequestBuilder
func (b *_GetAttributeAllRequest) CreateGetAttributeAllRequestBuilder() GetAttributeAllRequestBuilder {
	if b == nil {
		return NewGetAttributeAllRequestBuilder()
	}
	return &_GetAttributeAllRequestBuilder{_GetAttributeAllRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllRequest) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeAllRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_GetAttributeAllRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGetAttributeAllRequest(structType any) GetAttributeAllRequest {
	if casted, ok := structType.(GetAttributeAllRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllRequest) GetTypeName() string {
	return "GetAttributeAllRequest"
}

func (m *_GetAttributeAllRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GetAttributeAllRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetAttributeAllRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__getAttributeAllRequest GetAttributeAllRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeAllRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadImplicitField[uint8](ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	_ = requestPathSize

	classSegment, err := ReadSimpleField[PathSegment](ctx, "classSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'classSegment' field"))
	}
	m.ClassSegment = classSegment

	instanceSegment, err := ReadSimpleField[PathSegment](ctx, "instanceSegment", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'instanceSegment' field"))
	}
	m.InstanceSegment = instanceSegment

	if closeErr := readBuffer.CloseContext("GetAttributeAllRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllRequest")
	}

	return m, nil
}

func (m *_GetAttributeAllRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllRequest")
		}
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		if err := WriteImplicitField(ctx, "requestPathSize", requestPathSize, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "classSegment", m.GetClassSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'classSegment' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "instanceSegment", m.GetInstanceSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceSegment' field")
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllRequest) IsGetAttributeAllRequest() {}

func (m *_GetAttributeAllRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetAttributeAllRequest) deepCopy() *_GetAttributeAllRequest {
	if m == nil {
		return nil
	}
	_GetAttributeAllRequestCopy := &_GetAttributeAllRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
		utils.DeepCopy[PathSegment](m.ClassSegment),
		utils.DeepCopy[PathSegment](m.InstanceSegment),
	}
	_GetAttributeAllRequestCopy.CipServiceContract.(*_CipService)._SubType = m
	return _GetAttributeAllRequestCopy
}

func (m *_GetAttributeAllRequest) String() string {
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
