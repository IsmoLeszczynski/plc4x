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

// SetAttributeAllResponse is the corresponding interface of SetAttributeAllResponse
type SetAttributeAllResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CipService
	// IsSetAttributeAllResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSetAttributeAllResponse()
	// CreateBuilder creates a SetAttributeAllResponseBuilder
	CreateSetAttributeAllResponseBuilder() SetAttributeAllResponseBuilder
}

// _SetAttributeAllResponse is the data-structure of this message
type _SetAttributeAllResponse struct {
	CipServiceContract
}

var _ SetAttributeAllResponse = (*_SetAttributeAllResponse)(nil)
var _ CipServiceRequirements = (*_SetAttributeAllResponse)(nil)

// NewSetAttributeAllResponse factory function for _SetAttributeAllResponse
func NewSetAttributeAllResponse(serviceLen uint16) *_SetAttributeAllResponse {
	_result := &_SetAttributeAllResponse{
		CipServiceContract: NewCipService(serviceLen),
	}
	_result.CipServiceContract.(*_CipService)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SetAttributeAllResponseBuilder is a builder for SetAttributeAllResponse
type SetAttributeAllResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SetAttributeAllResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CipServiceBuilder
	// Build builds the SetAttributeAllResponse or returns an error if something is wrong
	Build() (SetAttributeAllResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SetAttributeAllResponse
}

// NewSetAttributeAllResponseBuilder() creates a SetAttributeAllResponseBuilder
func NewSetAttributeAllResponseBuilder() SetAttributeAllResponseBuilder {
	return &_SetAttributeAllResponseBuilder{_SetAttributeAllResponse: new(_SetAttributeAllResponse)}
}

type _SetAttributeAllResponseBuilder struct {
	*_SetAttributeAllResponse

	parentBuilder *_CipServiceBuilder

	err *utils.MultiError
}

var _ (SetAttributeAllResponseBuilder) = (*_SetAttributeAllResponseBuilder)(nil)

func (b *_SetAttributeAllResponseBuilder) setParent(contract CipServiceContract) {
	b.CipServiceContract = contract
}

func (b *_SetAttributeAllResponseBuilder) WithMandatoryFields() SetAttributeAllResponseBuilder {
	return b
}

func (b *_SetAttributeAllResponseBuilder) Build() (SetAttributeAllResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SetAttributeAllResponse.deepCopy(), nil
}

func (b *_SetAttributeAllResponseBuilder) MustBuild() SetAttributeAllResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SetAttributeAllResponseBuilder) Done() CipServiceBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCipServiceBuilder().(*_CipServiceBuilder)
	}
	return b.parentBuilder
}

func (b *_SetAttributeAllResponseBuilder) buildForCipService() (CipService, error) {
	return b.Build()
}

func (b *_SetAttributeAllResponseBuilder) DeepCopy() any {
	_copy := b.CreateSetAttributeAllResponseBuilder().(*_SetAttributeAllResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSetAttributeAllResponseBuilder creates a SetAttributeAllResponseBuilder
func (b *_SetAttributeAllResponse) CreateSetAttributeAllResponseBuilder() SetAttributeAllResponseBuilder {
	if b == nil {
		return NewSetAttributeAllResponseBuilder()
	}
	return &_SetAttributeAllResponseBuilder{_SetAttributeAllResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeAllResponse) GetService() uint8 {
	return 0x02
}

func (m *_SetAttributeAllResponse) GetResponse() bool {
	return bool(true)
}

func (m *_SetAttributeAllResponse) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeAllResponse) GetParent() CipServiceContract {
	return m.CipServiceContract
}

// Deprecated: use the interface for direct cast
func CastSetAttributeAllResponse(structType any) SetAttributeAllResponse {
	if casted, ok := structType.(SetAttributeAllResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeAllResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeAllResponse) GetTypeName() string {
	return "SetAttributeAllResponse"
}

func (m *_SetAttributeAllResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CipServiceContract.(*_CipService).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeAllResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SetAttributeAllResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CipService, connected bool, serviceLen uint16) (__setAttributeAllResponse SetAttributeAllResponse, err error) {
	m.CipServiceContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeAllResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeAllResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeAllResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeAllResponse")
	}

	return m, nil
}

func (m *_SetAttributeAllResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeAllResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeAllResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeAllResponse")
		}

		if popErr := writeBuffer.PopContext("SetAttributeAllResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeAllResponse")
		}
		return nil
	}
	return m.CipServiceContract.(*_CipService).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeAllResponse) IsSetAttributeAllResponse() {}

func (m *_SetAttributeAllResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SetAttributeAllResponse) deepCopy() *_SetAttributeAllResponse {
	if m == nil {
		return nil
	}
	_SetAttributeAllResponseCopy := &_SetAttributeAllResponse{
		m.CipServiceContract.(*_CipService).deepCopy(),
	}
	m.CipServiceContract.(*_CipService)._SubType = m
	return _SetAttributeAllResponseCopy
}

func (m *_SetAttributeAllResponse) String() string {
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
