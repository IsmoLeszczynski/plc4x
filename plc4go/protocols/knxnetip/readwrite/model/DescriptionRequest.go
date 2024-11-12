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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DescriptionRequest is the corresponding interface of DescriptionRequest
type DescriptionRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	KnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() HPAIControlEndpoint
	// IsDescriptionRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDescriptionRequest()
	// CreateBuilder creates a DescriptionRequestBuilder
	CreateDescriptionRequestBuilder() DescriptionRequestBuilder
}

// _DescriptionRequest is the data-structure of this message
type _DescriptionRequest struct {
	KnxNetIpMessageContract
	HpaiControlEndpoint HPAIControlEndpoint
}

var _ DescriptionRequest = (*_DescriptionRequest)(nil)
var _ KnxNetIpMessageRequirements = (*_DescriptionRequest)(nil)

// NewDescriptionRequest factory function for _DescriptionRequest
func NewDescriptionRequest(hpaiControlEndpoint HPAIControlEndpoint) *_DescriptionRequest {
	if hpaiControlEndpoint == nil {
		panic("hpaiControlEndpoint of type HPAIControlEndpoint for DescriptionRequest must not be nil")
	}
	_result := &_DescriptionRequest{
		KnxNetIpMessageContract: NewKnxNetIpMessage(),
		HpaiControlEndpoint:     hpaiControlEndpoint,
	}
	_result.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DescriptionRequestBuilder is a builder for DescriptionRequest
type DescriptionRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(hpaiControlEndpoint HPAIControlEndpoint) DescriptionRequestBuilder
	// WithHpaiControlEndpoint adds HpaiControlEndpoint (property field)
	WithHpaiControlEndpoint(HPAIControlEndpoint) DescriptionRequestBuilder
	// WithHpaiControlEndpointBuilder adds HpaiControlEndpoint (property field) which is build by the builder
	WithHpaiControlEndpointBuilder(func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) DescriptionRequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() KnxNetIpMessageBuilder
	// Build builds the DescriptionRequest or returns an error if something is wrong
	Build() (DescriptionRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DescriptionRequest
}

// NewDescriptionRequestBuilder() creates a DescriptionRequestBuilder
func NewDescriptionRequestBuilder() DescriptionRequestBuilder {
	return &_DescriptionRequestBuilder{_DescriptionRequest: new(_DescriptionRequest)}
}

type _DescriptionRequestBuilder struct {
	*_DescriptionRequest

	parentBuilder *_KnxNetIpMessageBuilder

	err *utils.MultiError
}

var _ (DescriptionRequestBuilder) = (*_DescriptionRequestBuilder)(nil)

func (b *_DescriptionRequestBuilder) setParent(contract KnxNetIpMessageContract) {
	b.KnxNetIpMessageContract = contract
}

func (b *_DescriptionRequestBuilder) WithMandatoryFields(hpaiControlEndpoint HPAIControlEndpoint) DescriptionRequestBuilder {
	return b.WithHpaiControlEndpoint(hpaiControlEndpoint)
}

func (b *_DescriptionRequestBuilder) WithHpaiControlEndpoint(hpaiControlEndpoint HPAIControlEndpoint) DescriptionRequestBuilder {
	b.HpaiControlEndpoint = hpaiControlEndpoint
	return b
}

func (b *_DescriptionRequestBuilder) WithHpaiControlEndpointBuilder(builderSupplier func(HPAIControlEndpointBuilder) HPAIControlEndpointBuilder) DescriptionRequestBuilder {
	builder := builderSupplier(b.HpaiControlEndpoint.CreateHPAIControlEndpointBuilder())
	var err error
	b.HpaiControlEndpoint, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HPAIControlEndpointBuilder failed"))
	}
	return b
}

func (b *_DescriptionRequestBuilder) Build() (DescriptionRequest, error) {
	if b.HpaiControlEndpoint == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hpaiControlEndpoint' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DescriptionRequest.deepCopy(), nil
}

func (b *_DescriptionRequestBuilder) MustBuild() DescriptionRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DescriptionRequestBuilder) Done() KnxNetIpMessageBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewKnxNetIpMessageBuilder().(*_KnxNetIpMessageBuilder)
	}
	return b.parentBuilder
}

func (b *_DescriptionRequestBuilder) buildForKnxNetIpMessage() (KnxNetIpMessage, error) {
	return b.Build()
}

func (b *_DescriptionRequestBuilder) DeepCopy() any {
	_copy := b.CreateDescriptionRequestBuilder().(*_DescriptionRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDescriptionRequestBuilder creates a DescriptionRequestBuilder
func (b *_DescriptionRequest) CreateDescriptionRequestBuilder() DescriptionRequestBuilder {
	if b == nil {
		return NewDescriptionRequestBuilder()
	}
	return &_DescriptionRequestBuilder{_DescriptionRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DescriptionRequest) GetMsgType() uint16 {
	return 0x0203
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DescriptionRequest) GetParent() KnxNetIpMessageContract {
	return m.KnxNetIpMessageContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DescriptionRequest) GetHpaiControlEndpoint() HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDescriptionRequest(structType any) DescriptionRequest {
	if casted, ok := structType.(DescriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DescriptionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DescriptionRequest) GetTypeName() string {
	return "DescriptionRequest"
}

func (m *_DescriptionRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.KnxNetIpMessageContract.(*_KnxNetIpMessage).getLengthInBits(ctx))

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_DescriptionRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DescriptionRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_KnxNetIpMessage) (__descriptionRequest DescriptionRequest, err error) {
	m.KnxNetIpMessageContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DescriptionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DescriptionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	hpaiControlEndpoint, err := ReadSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", ReadComplex[HPAIControlEndpoint](HPAIControlEndpointParseWithBuffer, readBuffer), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hpaiControlEndpoint' field"))
	}
	m.HpaiControlEndpoint = hpaiControlEndpoint

	if closeErr := readBuffer.CloseContext("DescriptionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DescriptionRequest")
	}

	return m, nil
}

func (m *_DescriptionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DescriptionRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DescriptionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DescriptionRequest")
		}

		if err := WriteSimpleField[HPAIControlEndpoint](ctx, "hpaiControlEndpoint", m.GetHpaiControlEndpoint(), WriteComplex[HPAIControlEndpoint](writeBuffer), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("DescriptionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DescriptionRequest")
		}
		return nil
	}
	return m.KnxNetIpMessageContract.(*_KnxNetIpMessage).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DescriptionRequest) IsDescriptionRequest() {}

func (m *_DescriptionRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DescriptionRequest) deepCopy() *_DescriptionRequest {
	if m == nil {
		return nil
	}
	_DescriptionRequestCopy := &_DescriptionRequest{
		m.KnxNetIpMessageContract.(*_KnxNetIpMessage).deepCopy(),
		utils.DeepCopy[HPAIControlEndpoint](m.HpaiControlEndpoint),
	}
	m.KnxNetIpMessageContract.(*_KnxNetIpMessage)._SubType = m
	return _DescriptionRequestCopy
}

func (m *_DescriptionRequest) String() string {
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
