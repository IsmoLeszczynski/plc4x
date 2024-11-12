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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// AdsInvalidResponse is the corresponding interface of AdsInvalidResponse
type AdsInvalidResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AmsPacket
	// IsAdsInvalidResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAdsInvalidResponse()
	// CreateBuilder creates a AdsInvalidResponseBuilder
	CreateAdsInvalidResponseBuilder() AdsInvalidResponseBuilder
}

// _AdsInvalidResponse is the data-structure of this message
type _AdsInvalidResponse struct {
	AmsPacketContract
}

var _ AdsInvalidResponse = (*_AdsInvalidResponse)(nil)
var _ AmsPacketRequirements = (*_AdsInvalidResponse)(nil)

// NewAdsInvalidResponse factory function for _AdsInvalidResponse
func NewAdsInvalidResponse(targetAmsNetId AmsNetId, targetAmsPort uint16, sourceAmsNetId AmsNetId, sourceAmsPort uint16, errorCode uint32, invokeId uint32) *_AdsInvalidResponse {
	_result := &_AdsInvalidResponse{
		AmsPacketContract: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result.AmsPacketContract.(*_AmsPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AdsInvalidResponseBuilder is a builder for AdsInvalidResponse
type AdsInvalidResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() AdsInvalidResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AmsPacketBuilder
	// Build builds the AdsInvalidResponse or returns an error if something is wrong
	Build() (AdsInvalidResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AdsInvalidResponse
}

// NewAdsInvalidResponseBuilder() creates a AdsInvalidResponseBuilder
func NewAdsInvalidResponseBuilder() AdsInvalidResponseBuilder {
	return &_AdsInvalidResponseBuilder{_AdsInvalidResponse: new(_AdsInvalidResponse)}
}

type _AdsInvalidResponseBuilder struct {
	*_AdsInvalidResponse

	parentBuilder *_AmsPacketBuilder

	err *utils.MultiError
}

var _ (AdsInvalidResponseBuilder) = (*_AdsInvalidResponseBuilder)(nil)

func (b *_AdsInvalidResponseBuilder) setParent(contract AmsPacketContract) {
	b.AmsPacketContract = contract
}

func (b *_AdsInvalidResponseBuilder) WithMandatoryFields() AdsInvalidResponseBuilder {
	return b
}

func (b *_AdsInvalidResponseBuilder) Build() (AdsInvalidResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AdsInvalidResponse.deepCopy(), nil
}

func (b *_AdsInvalidResponseBuilder) MustBuild() AdsInvalidResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AdsInvalidResponseBuilder) Done() AmsPacketBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAmsPacketBuilder().(*_AmsPacketBuilder)
	}
	return b.parentBuilder
}

func (b *_AdsInvalidResponseBuilder) buildForAmsPacket() (AmsPacket, error) {
	return b.Build()
}

func (b *_AdsInvalidResponseBuilder) DeepCopy() any {
	_copy := b.CreateAdsInvalidResponseBuilder().(*_AdsInvalidResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAdsInvalidResponseBuilder creates a AdsInvalidResponseBuilder
func (b *_AdsInvalidResponse) CreateAdsInvalidResponseBuilder() AdsInvalidResponseBuilder {
	if b == nil {
		return NewAdsInvalidResponseBuilder()
	}
	return &_AdsInvalidResponseBuilder{_AdsInvalidResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsInvalidResponse) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *_AdsInvalidResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsInvalidResponse) GetParent() AmsPacketContract {
	return m.AmsPacketContract
}

// Deprecated: use the interface for direct cast
func CastAdsInvalidResponse(structType any) AdsInvalidResponse {
	if casted, ok := structType.(AdsInvalidResponse); ok {
		return casted
	}
	if casted, ok := structType.(*AdsInvalidResponse); ok {
		return *casted
	}
	return nil
}

func (m *_AdsInvalidResponse) GetTypeName() string {
	return "AdsInvalidResponse"
}

func (m *_AdsInvalidResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AmsPacketContract.(*_AmsPacket).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_AdsInvalidResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AdsInvalidResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AmsPacket) (__adsInvalidResponse AdsInvalidResponse, err error) {
	m.AmsPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsInvalidResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsInvalidResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsInvalidResponse")
	}

	return m, nil
}

func (m *_AdsInvalidResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsInvalidResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsInvalidResponse")
		}

		if popErr := writeBuffer.PopContext("AdsInvalidResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsInvalidResponse")
		}
		return nil
	}
	return m.AmsPacketContract.(*_AmsPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AdsInvalidResponse) IsAdsInvalidResponse() {}

func (m *_AdsInvalidResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AdsInvalidResponse) deepCopy() *_AdsInvalidResponse {
	if m == nil {
		return nil
	}
	_AdsInvalidResponseCopy := &_AdsInvalidResponse{
		m.AmsPacketContract.(*_AmsPacket).deepCopy(),
	}
	m.AmsPacketContract.(*_AmsPacket)._SubType = m
	return _AdsInvalidResponseCopy
}

func (m *_AdsInvalidResponse) String() string {
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
