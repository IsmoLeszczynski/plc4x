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

// CipConnectionManagerCloseRequest is the corresponding interface of CipConnectionManagerCloseRequest
type CipConnectionManagerCloseRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// GetRequestPathSize returns RequestPathSize (property field)
	GetRequestPathSize() uint8
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
	// GetPriority returns Priority (property field)
	GetPriority() uint8
	// GetTickTime returns TickTime (property field)
	GetTickTime() uint8
	// GetTimeoutTicks returns TimeoutTicks (property field)
	GetTimeoutTicks() uint8
	// GetConnectionSerialNumber returns ConnectionSerialNumber (property field)
	GetConnectionSerialNumber() uint16
	// GetOriginatorVendorId returns OriginatorVendorId (property field)
	GetOriginatorVendorId() uint16
	// GetOriginatorSerialNumber returns OriginatorSerialNumber (property field)
	GetOriginatorSerialNumber() uint32
	// GetConnectionPathSize returns ConnectionPathSize (property field)
	GetConnectionPathSize() uint8
	// GetConnectionPaths returns ConnectionPaths (property field)
	GetConnectionPaths() []PathSegment
	// IsCipConnectionManagerCloseRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipConnectionManagerCloseRequest()
	// CreateBuilder creates a CipConnectionManagerCloseRequestBuilder
	CreateCipConnectionManagerCloseRequestBuilder() CipConnectionManagerCloseRequestBuilder
}

// _CipConnectionManagerCloseRequest is the data-structure of this message
type _CipConnectionManagerCloseRequest struct {
	CipServiceContract
	RequestPathSize        uint8
	ClassSegment           PathSegment
	InstanceSegment        PathSegment
	Priority               uint8
	TickTime               uint8
	TimeoutTicks           uint8
	ConnectionSerialNumber uint16
	OriginatorVendorId     uint16
	OriginatorSerialNumber uint32
	ConnectionPathSize     uint8
	ConnectionPaths        []PathSegment
	// Reserved Fields
	reservedField0 *byte
}

var _ CipConnectionManagerCloseRequest = (*_CipConnectionManagerCloseRequest)(nil)
var _ CipServiceRequirements = (*_CipConnectionManagerCloseRequest)(nil)

// NewCipConnectionManagerCloseRequest factory function for _CipConnectionManagerCloseRequest
func NewCipConnectionManagerCloseRequest(requestPathSize uint8, classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, connectionPathSize uint8, connectionPaths []PathSegment, serviceLen uint16) *_CipConnectionManagerCloseRequest {
	if classSegment == nil {
		panic("classSegment of type PathSegment for CipConnectionManagerCloseRequest must not be nil")
	}
	if instanceSegment == nil {
		panic("instanceSegment of type PathSegment for CipConnectionManagerCloseRequest must not be nil")
	}
	_result := &_CipConnectionManagerCloseRequest{
		CipServiceContract:     NewCipService(serviceLen),
		RequestPathSize:        requestPathSize,
		ClassSegment:           classSegment,
		InstanceSegment:        instanceSegment,
		Priority:               priority,
		TickTime:               tickTime,
		TimeoutTicks:           timeoutTicks,
		ConnectionSerialNumber: connectionSerialNumber,
		OriginatorVendorId:     originatorVendorId,
		OriginatorSerialNumber: originatorSerialNumber,
		ConnectionPathSize:     connectionPathSize,
		ConnectionPaths:        connectionPaths,
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipConnectionManagerCloseRequestBuilder is a builder for CipConnectionManagerCloseRequest
type CipConnectionManagerCloseRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestPathSize uint8, classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, connectionPathSize uint8, connectionPaths []PathSegment) CipConnectionManagerCloseRequestBuilder
	// WithRequestPathSize adds RequestPathSize (property field)
	WithRequestPathSize(uint8) CipConnectionManagerCloseRequestBuilder
	// WithClassSegment adds ClassSegment (property field)
	WithClassSegment(PathSegment) CipConnectionManagerCloseRequestBuilder
	// WithClassSegmentBuilder adds ClassSegment (property field) which is build by the builder
	WithClassSegmentBuilder(func(PathSegmentBuilder) PathSegmentBuilder) CipConnectionManagerCloseRequestBuilder
	// WithInstanceSegment adds InstanceSegment (property field)
	WithInstanceSegment(PathSegment) CipConnectionManagerCloseRequestBuilder
	// WithInstanceSegmentBuilder adds InstanceSegment (property field) which is build by the builder
	WithInstanceSegmentBuilder(func(PathSegmentBuilder) PathSegmentBuilder) CipConnectionManagerCloseRequestBuilder
	// WithPriority adds Priority (property field)
	WithPriority(uint8) CipConnectionManagerCloseRequestBuilder
	// WithTickTime adds TickTime (property field)
	WithTickTime(uint8) CipConnectionManagerCloseRequestBuilder
	// WithTimeoutTicks adds TimeoutTicks (property field)
	WithTimeoutTicks(uint8) CipConnectionManagerCloseRequestBuilder
	// WithConnectionSerialNumber adds ConnectionSerialNumber (property field)
	WithConnectionSerialNumber(uint16) CipConnectionManagerCloseRequestBuilder
	// WithOriginatorVendorId adds OriginatorVendorId (property field)
	WithOriginatorVendorId(uint16) CipConnectionManagerCloseRequestBuilder
	// WithOriginatorSerialNumber adds OriginatorSerialNumber (property field)
	WithOriginatorSerialNumber(uint32) CipConnectionManagerCloseRequestBuilder
	// WithConnectionPathSize adds ConnectionPathSize (property field)
	WithConnectionPathSize(uint8) CipConnectionManagerCloseRequestBuilder
	// WithConnectionPaths adds ConnectionPaths (property field)
	WithConnectionPaths(...PathSegment) CipConnectionManagerCloseRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the CipConnectionManagerCloseRequest or returns an error if something is wrong
	Build() (CipConnectionManagerCloseRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipConnectionManagerCloseRequest
}

// NewCipConnectionManagerCloseRequestBuilder() creates a CipConnectionManagerCloseRequestBuilder
func NewCipConnectionManagerCloseRequestBuilder() CipConnectionManagerCloseRequestBuilder {
	return &_CipConnectionManagerCloseRequestBuilder{_CipConnectionManagerCloseRequest: new(_CipConnectionManagerCloseRequest)}
}

type _CipConnectionManagerCloseRequestBuilder struct {
	*_CipConnectionManagerCloseRequest

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (CipConnectionManagerCloseRequestBuilder) = (*_CipConnectionManagerCloseRequestBuilder)(nil)

func (b *_CipConnectionManagerCloseRequestBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithMandatoryFields(requestPathSize uint8, classSegment PathSegment, instanceSegment PathSegment, priority uint8, tickTime uint8, timeoutTicks uint8, connectionSerialNumber uint16, originatorVendorId uint16, originatorSerialNumber uint32, connectionPathSize uint8, connectionPaths []PathSegment) CipConnectionManagerCloseRequestBuilder {
	return b.WithRequestPathSize(requestPathSize).WithClassSegment(classSegment).WithInstanceSegment(instanceSegment).WithPriority(priority).WithTickTime(tickTime).WithTimeoutTicks(timeoutTicks).WithConnectionSerialNumber(connectionSerialNumber).WithOriginatorVendorId(originatorVendorId).WithOriginatorSerialNumber(originatorSerialNumber).WithConnectionPathSize(connectionPathSize).WithConnectionPaths(connectionPaths...)
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithRequestPathSize(requestPathSize uint8) CipConnectionManagerCloseRequestBuilder {
	b.RequestPathSize = requestPathSize
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithClassSegment(classSegment PathSegment) CipConnectionManagerCloseRequestBuilder {
	b.ClassSegment = classSegment
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithClassSegmentBuilder(builderSupplier func(PathSegmentBuilder) PathSegmentBuilder) CipConnectionManagerCloseRequestBuilder {
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

func (b *_CipConnectionManagerCloseRequestBuilder) WithInstanceSegment(instanceSegment PathSegment) CipConnectionManagerCloseRequestBuilder {
	b.InstanceSegment = instanceSegment
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithInstanceSegmentBuilder(builderSupplier func(PathSegmentBuilder) PathSegmentBuilder) CipConnectionManagerCloseRequestBuilder {
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

func (b *_CipConnectionManagerCloseRequestBuilder) WithPriority(priority uint8) CipConnectionManagerCloseRequestBuilder {
	b.Priority = priority
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithTickTime(tickTime uint8) CipConnectionManagerCloseRequestBuilder {
	b.TickTime = tickTime
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithTimeoutTicks(timeoutTicks uint8) CipConnectionManagerCloseRequestBuilder {
	b.TimeoutTicks = timeoutTicks
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithConnectionSerialNumber(connectionSerialNumber uint16) CipConnectionManagerCloseRequestBuilder {
	b.ConnectionSerialNumber = connectionSerialNumber
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithOriginatorVendorId(originatorVendorId uint16) CipConnectionManagerCloseRequestBuilder {
	b.OriginatorVendorId = originatorVendorId
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithOriginatorSerialNumber(originatorSerialNumber uint32) CipConnectionManagerCloseRequestBuilder {
	b.OriginatorSerialNumber = originatorSerialNumber
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithConnectionPathSize(connectionPathSize uint8) CipConnectionManagerCloseRequestBuilder {
	b.ConnectionPathSize = connectionPathSize
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) WithConnectionPaths(connectionPaths ...PathSegment) CipConnectionManagerCloseRequestBuilder {
	b.ConnectionPaths = connectionPaths
	return b
}

func (b *_CipConnectionManagerCloseRequestBuilder) Build() (CipConnectionManagerCloseRequest, error) {
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
	return b._CipConnectionManagerCloseRequest.deepCopy(), nil
}

func (b *_CipConnectionManagerCloseRequestBuilder) MustBuild() CipConnectionManagerCloseRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipConnectionManagerCloseRequestBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_CipConnectionManagerCloseRequestBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_CipConnectionManagerCloseRequestBuilder) DeepCopy() any {
	_copy := b.CreateCipConnectionManagerCloseRequestBuilder().(*_CipConnectionManagerCloseRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipConnectionManagerCloseRequestBuilder creates a CipConnectionManagerCloseRequestBuilder
func (b *_CipConnectionManagerCloseRequest) CreateCipConnectionManagerCloseRequestBuilder() CipConnectionManagerCloseRequestBuilder {
	if b == nil {
		return NewCipConnectionManagerCloseRequestBuilder()
	}
	return &_CipConnectionManagerCloseRequestBuilder{_CipConnectionManagerCloseRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetService() uint8 {
	return 0x4E
}

func (m *_CipConnectionManagerCloseRequest) GetResponse() bool {
	return bool(false)
}

func (m *_CipConnectionManagerCloseRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipConnectionManagerCloseRequest) GetParent() CipServiceContract {
	return m.CipServiceContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipConnectionManagerCloseRequest) GetRequestPathSize() uint8 {
	return m.RequestPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_CipConnectionManagerCloseRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

func (m *_CipConnectionManagerCloseRequest) GetPriority() uint8 {
	return m.Priority
}

func (m *_CipConnectionManagerCloseRequest) GetTickTime() uint8 {
	return m.TickTime
}

func (m *_CipConnectionManagerCloseRequest) GetTimeoutTicks() uint8 {
	return m.TimeoutTicks
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionSerialNumber() uint16 {
	return m.ConnectionSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorVendorId() uint16 {
	return m.OriginatorVendorId
}

func (m *_CipConnectionManagerCloseRequest) GetOriginatorSerialNumber() uint32 {
	return m.OriginatorSerialNumber
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPathSize() uint8 {
	return m.ConnectionPathSize
}

func (m *_CipConnectionManagerCloseRequest) GetConnectionPaths() []PathSegment {
	return m.ConnectionPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipConnectionManagerCloseRequest(structType any) CipConnectionManagerCloseRequest {
	if casted, ok := structType.(CipConnectionManagerCloseRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipConnectionManagerCloseRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipConnectionManagerCloseRequest) GetTypeName() string {
	return "CipConnectionManagerCloseRequest"
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	// Simple field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	// Simple field (priority)
	lengthInBits += 4

	// Simple field (tickTime)
	lengthInBits += 4

	// Simple field (timeoutTicks)
	lengthInBits += 8

	// Simple field (connectionSerialNumber)
	lengthInBits += 16

	// Simple field (originatorVendorId)
	lengthInBits += 16

	// Simple field (originatorSerialNumber)
	lengthInBits += 32

	// Simple field (connectionPathSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Array field
	if len(m.ConnectionPaths) > 0 {
		for _, element := range m.ConnectionPaths {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_CipConnectionManagerCloseRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipConnectionManagerCloseRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__cipConnectionManagerCloseRequest CipConnectionManagerCloseRequest, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipConnectionManagerCloseRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipConnectionManagerCloseRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestPathSize, err := ReadSimpleField(ctx, "requestPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestPathSize' field"))
	}
	m.RequestPathSize = requestPathSize

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

	priority, err := ReadSimpleField(ctx, "priority", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	m.Priority = priority

	tickTime, err := ReadSimpleField(ctx, "tickTime", ReadUnsignedByte(readBuffer, uint8(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'tickTime' field"))
	}
	m.TickTime = tickTime

	timeoutTicks, err := ReadSimpleField(ctx, "timeoutTicks", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeoutTicks' field"))
	}
	m.TimeoutTicks = timeoutTicks

	connectionSerialNumber, err := ReadSimpleField(ctx, "connectionSerialNumber", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionSerialNumber' field"))
	}
	m.ConnectionSerialNumber = connectionSerialNumber

	originatorVendorId, err := ReadSimpleField(ctx, "originatorVendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorVendorId' field"))
	}
	m.OriginatorVendorId = originatorVendorId

	originatorSerialNumber, err := ReadSimpleField(ctx, "originatorSerialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'originatorSerialNumber' field"))
	}
	m.OriginatorSerialNumber = originatorSerialNumber

	connectionPathSize, err := ReadSimpleField(ctx, "connectionPathSize", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPathSize' field"))
	}
	m.ConnectionPathSize = connectionPathSize

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	connectionPaths, err := ReadTerminatedArrayField[PathSegment](ctx, "connectionPaths", ReadComplex[PathSegment](PathSegmentParseWithBuffer, readBuffer), NoMorePathSegments(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionPaths' field"))
	}
	m.ConnectionPaths = connectionPaths

	if closeErr := readBuffer.CloseContext("CipConnectionManagerCloseRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipConnectionManagerCloseRequest")
	}

	return m, nil
}

func (m *_CipConnectionManagerCloseRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipConnectionManagerCloseRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipConnectionManagerCloseRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipConnectionManagerCloseRequest")
		}

		if err := WriteSimpleField[uint8](ctx, "requestPathSize", m.GetRequestPathSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestPathSize' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "classSegment", m.GetClassSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'classSegment' field")
		}

		if err := WriteSimpleField[PathSegment](ctx, "instanceSegment", m.GetInstanceSegment(), WriteComplex[PathSegment](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'instanceSegment' field")
		}

		if err := WriteSimpleField[uint8](ctx, "priority", m.GetPriority(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if err := WriteSimpleField[uint8](ctx, "tickTime", m.GetTickTime(), WriteUnsignedByte(writeBuffer, 4)); err != nil {
			return errors.Wrap(err, "Error serializing 'tickTime' field")
		}

		if err := WriteSimpleField[uint8](ctx, "timeoutTicks", m.GetTimeoutTicks(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeoutTicks' field")
		}

		if err := WriteSimpleField[uint16](ctx, "connectionSerialNumber", m.GetConnectionSerialNumber(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionSerialNumber' field")
		}

		if err := WriteSimpleField[uint16](ctx, "originatorVendorId", m.GetOriginatorVendorId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorVendorId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "originatorSerialNumber", m.GetOriginatorSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'originatorSerialNumber' field")
		}

		if err := WriteSimpleField[uint8](ctx, "connectionPathSize", m.GetConnectionPathSize(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionPathSize' field")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x00), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteComplexTypeArrayField(ctx, "connectionPaths", m.GetConnectionPaths(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionPaths' field")
		}

		if popErr := writeBuffer.PopContext("CipConnectionManagerCloseRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipConnectionManagerCloseRequest")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipConnectionManagerCloseRequest) IsCipConnectionManagerCloseRequest() {}

func (m *_CipConnectionManagerCloseRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipConnectionManagerCloseRequest) deepCopy() *_CipConnectionManagerCloseRequest {
	if m == nil {
		return nil
	}
	_CipConnectionManagerCloseRequestCopy := &_CipConnectionManagerCloseRequest{
		m.CipServiceContract.(*_CipService).deepCopy(),
		m.RequestPathSize,
		utils.DeepCopy[PathSegment](m.ClassSegment),
		utils.DeepCopy[PathSegment](m.InstanceSegment),
		m.Priority,
		m.TickTime,
		m.TimeoutTicks,
		m.ConnectionSerialNumber,
		m.OriginatorVendorId,
		m.OriginatorSerialNumber,
		m.ConnectionPathSize,
		utils.DeepCopySlice[PathSegment, PathSegment](m.ConnectionPaths),
		m.reservedField0,
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _CipConnectionManagerCloseRequestCopy
}

func (m *_CipConnectionManagerCloseRequest) String() string {
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
