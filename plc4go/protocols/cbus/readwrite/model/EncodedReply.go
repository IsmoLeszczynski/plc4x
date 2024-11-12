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

// EncodedReply is the corresponding interface of EncodedReply
type EncodedReply interface {
	EncodedReplyContract
	EncodedReplyRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsEncodedReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEncodedReply()
	// CreateBuilder creates a EncodedReplyBuilder
	CreateEncodedReplyBuilder() EncodedReplyBuilder
}

// EncodedReplyContract provides a set of functions which can be overwritten by a sub struct
type EncodedReplyContract interface {
	// GetPeekedByte returns PeekedByte (property field)
	GetPeekedByte() byte
	// GetIsMonitoredSAL returns IsMonitoredSAL (virtual field)
	GetIsMonitoredSAL() bool
	// GetCBusOptions() returns a parser argument
	GetCBusOptions() CBusOptions
	// GetRequestContext() returns a parser argument
	GetRequestContext() RequestContext
	// IsEncodedReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsEncodedReply()
	// CreateBuilder creates a EncodedReplyBuilder
	CreateEncodedReplyBuilder() EncodedReplyBuilder
}

// EncodedReplyRequirements provides a set of functions which need to be implemented by a sub struct
type EncodedReplyRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetIsMonitoredSAL returns IsMonitoredSAL (discriminator field)
	GetIsMonitoredSAL() bool
}

// _EncodedReply is the data-structure of this message
type _EncodedReply struct {
	_SubType interface {
		EncodedReplyContract
		EncodedReplyRequirements
	}
	PeekedByte byte

	// Arguments.
	CBusOptions    CBusOptions
	RequestContext RequestContext
}

var _ EncodedReplyContract = (*_EncodedReply)(nil)

// NewEncodedReply factory function for _EncodedReply
func NewEncodedReply(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_EncodedReply {
	return &_EncodedReply{PeekedByte: peekedByte, CBusOptions: cBusOptions, RequestContext: requestContext}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// EncodedReplyBuilder is a builder for EncodedReply
type EncodedReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(peekedByte byte) EncodedReplyBuilder
	// WithPeekedByte adds PeekedByte (property field)
	WithPeekedByte(byte) EncodedReplyBuilder
	// AsMonitoredSALReply converts this build to a subType of EncodedReply. It is always possible to return to current builder using Done()
	AsMonitoredSALReply() MonitoredSALReplyBuilder
	// AsEncodedReplyCALReply converts this build to a subType of EncodedReply. It is always possible to return to current builder using Done()
	AsEncodedReplyCALReply() EncodedReplyCALReplyBuilder
	// Build builds the EncodedReply or returns an error if something is wrong
	PartialBuild() (EncodedReplyContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() EncodedReplyContract
	// Build builds the EncodedReply or returns an error if something is wrong
	Build() (EncodedReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() EncodedReply
}

// NewEncodedReplyBuilder() creates a EncodedReplyBuilder
func NewEncodedReplyBuilder() EncodedReplyBuilder {
	return &_EncodedReplyBuilder{_EncodedReply: new(_EncodedReply)}
}

type _EncodedReplyChildBuilder interface {
	utils.Copyable
	setParent(EncodedReplyContract)
	buildForEncodedReply() (EncodedReply, error)
}

type _EncodedReplyBuilder struct {
	*_EncodedReply

	childBuilder _EncodedReplyChildBuilder

	err *utils.MultiError
}

var _ (EncodedReplyBuilder) = (*_EncodedReplyBuilder)(nil)

func (b *_EncodedReplyBuilder) WithMandatoryFields(peekedByte byte) EncodedReplyBuilder {
	return b.WithPeekedByte(peekedByte)
}

func (b *_EncodedReplyBuilder) WithPeekedByte(peekedByte byte) EncodedReplyBuilder {
	b.PeekedByte = peekedByte
	return b
}

func (b *_EncodedReplyBuilder) PartialBuild() (EncodedReplyContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._EncodedReply.deepCopy(), nil
}

func (b *_EncodedReplyBuilder) PartialMustBuild() EncodedReplyContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EncodedReplyBuilder) AsMonitoredSALReply() MonitoredSALReplyBuilder {
	if cb, ok := b.childBuilder.(MonitoredSALReplyBuilder); ok {
		return cb
	}
	cb := NewMonitoredSALReplyBuilder().(*_MonitoredSALReplyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_EncodedReplyBuilder) AsEncodedReplyCALReply() EncodedReplyCALReplyBuilder {
	if cb, ok := b.childBuilder.(EncodedReplyCALReplyBuilder); ok {
		return cb
	}
	cb := NewEncodedReplyCALReplyBuilder().(*_EncodedReplyCALReplyBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_EncodedReplyBuilder) Build() (EncodedReply, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForEncodedReply()
}

func (b *_EncodedReplyBuilder) MustBuild() EncodedReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_EncodedReplyBuilder) DeepCopy() any {
	_copy := b.CreateEncodedReplyBuilder().(*_EncodedReplyBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_EncodedReplyChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateEncodedReplyBuilder creates a EncodedReplyBuilder
func (b *_EncodedReply) CreateEncodedReplyBuilder() EncodedReplyBuilder {
	if b == nil {
		return NewEncodedReplyBuilder()
	}
	return &_EncodedReplyBuilder{_EncodedReply: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_EncodedReply) GetPeekedByte() byte {
	return m.PeekedByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (pm *_EncodedReply) GetIsMonitoredSAL() bool {
	m := pm._SubType
	ctx := context.Background()
	_ = ctx
	return bool(bool((bool(bool(bool((m.GetPeekedByte()&0x3F) == (0x05))) || bool(bool((m.GetPeekedByte()) == (0x00)))) || bool(bool((m.GetPeekedByte()&0xF8) == (0x00))))) && bool(!(m.GetRequestContext().GetSendIdentifyRequestBefore())))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastEncodedReply(structType any) EncodedReply {
	if casted, ok := structType.(EncodedReply); ok {
		return casted
	}
	if casted, ok := structType.(*EncodedReply); ok {
		return *casted
	}
	return nil
}

func (m *_EncodedReply) GetTypeName() string {
	return "EncodedReply"
}

func (m *_EncodedReply) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_EncodedReply) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_EncodedReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func EncodedReplyParse[T EncodedReply](ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	return EncodedReplyParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func EncodedReplyParseWithBufferProducer[T EncodedReply](cBusOptions CBusOptions, requestContext RequestContext) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := EncodedReplyParseWithBuffer[T](ctx, readBuffer, cBusOptions, requestContext)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func EncodedReplyParseWithBuffer[T EncodedReply](ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (T, error) {
	v, err := (&_EncodedReply{CBusOptions: cBusOptions, RequestContext: requestContext}).parse(ctx, readBuffer, cBusOptions, requestContext)
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

func (m *_EncodedReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (__encodedReply EncodedReply, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EncodedReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EncodedReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	peekedByte, err := ReadPeekField[byte](ctx, "peekedByte", ReadByte(readBuffer, 8), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'peekedByte' field"))
	}
	m.PeekedByte = peekedByte

	isMonitoredSAL, err := ReadVirtualField[bool](ctx, "isMonitoredSAL", (*bool)(nil), bool((bool(bool(bool((peekedByte&0x3F) == (0x05))) || bool(bool((peekedByte) == (0x00)))) || bool(bool((peekedByte&0xF8) == (0x00))))) && bool(!(requestContext.GetSendIdentifyRequestBefore())))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isMonitoredSAL' field"))
	}
	_ = isMonitoredSAL

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child EncodedReply
	switch {
	case isMonitoredSAL == bool(true): // MonitoredSALReply
		if _child, err = new(_MonitoredSALReply).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type MonitoredSALReply for type-switch of EncodedReply")
		}
	case 0 == 0: // EncodedReplyCALReply
		if _child, err = new(_EncodedReplyCALReply).parse(ctx, readBuffer, m, cBusOptions, requestContext); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type EncodedReplyCALReply for type-switch of EncodedReply")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [isMonitoredSAL=%v]", isMonitoredSAL)
	}

	if closeErr := readBuffer.CloseContext("EncodedReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EncodedReply")
	}

	return _child, nil
}

func (pm *_EncodedReply) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child EncodedReply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("EncodedReply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EncodedReply")
	}
	// Virtual field
	isMonitoredSAL := m.GetIsMonitoredSAL()
	_ = isMonitoredSAL
	if _isMonitoredSALErr := writeBuffer.WriteVirtual(ctx, "isMonitoredSAL", m.GetIsMonitoredSAL()); _isMonitoredSALErr != nil {
		return errors.Wrap(_isMonitoredSALErr, "Error serializing 'isMonitoredSAL' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("EncodedReply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EncodedReply")
	}
	return nil
}

////
// Arguments Getter

func (m *_EncodedReply) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_EncodedReply) GetRequestContext() RequestContext {
	return m.RequestContext
}

//
////

func (m *_EncodedReply) IsEncodedReply() {}

func (m *_EncodedReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_EncodedReply) deepCopy() *_EncodedReply {
	if m == nil {
		return nil
	}
	_EncodedReplyCopy := &_EncodedReply{
		nil, // will be set by child
		m.PeekedByte,
		m.CBusOptions,
		m.RequestContext,
	}
	return _EncodedReplyCopy
}
