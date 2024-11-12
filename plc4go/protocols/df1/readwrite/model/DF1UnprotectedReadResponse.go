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
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// DF1UnprotectedReadResponse is the corresponding interface of DF1UnprotectedReadResponse
type DF1UnprotectedReadResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	DF1Command
	// GetData returns Data (property field)
	GetData() []byte
	// IsDF1UnprotectedReadResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDF1UnprotectedReadResponse()
	// CreateBuilder creates a DF1UnprotectedReadResponseBuilder
	CreateDF1UnprotectedReadResponseBuilder() DF1UnprotectedReadResponseBuilder
}

// _DF1UnprotectedReadResponse is the data-structure of this message
type _DF1UnprotectedReadResponse struct {
	DF1CommandContract
	Data []byte
}

var _ DF1UnprotectedReadResponse = (*_DF1UnprotectedReadResponse)(nil)
var _ DF1CommandRequirements = (*_DF1UnprotectedReadResponse)(nil)

// NewDF1UnprotectedReadResponse factory function for _DF1UnprotectedReadResponse
func NewDF1UnprotectedReadResponse(status uint8, transactionCounter uint16, data []byte) *_DF1UnprotectedReadResponse {
	_result := &_DF1UnprotectedReadResponse{
		DF1CommandContract: NewDF1Command(status, transactionCounter),
		Data:               data,
	}
	_result.DF1CommandContract.(*_DF1Command)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DF1UnprotectedReadResponseBuilder is a builder for DF1UnprotectedReadResponse
type DF1UnprotectedReadResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(data []byte) DF1UnprotectedReadResponseBuilder
	// WithData adds Data (property field)
	WithData(...byte) DF1UnprotectedReadResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() DF1CommandBuilder
	// Build builds the DF1UnprotectedReadResponse or returns an error if something is wrong
	Build() (DF1UnprotectedReadResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DF1UnprotectedReadResponse
}

// NewDF1UnprotectedReadResponseBuilder() creates a DF1UnprotectedReadResponseBuilder
func NewDF1UnprotectedReadResponseBuilder() DF1UnprotectedReadResponseBuilder {
	return &_DF1UnprotectedReadResponseBuilder{_DF1UnprotectedReadResponse: new(_DF1UnprotectedReadResponse)}
}

type _DF1UnprotectedReadResponseBuilder struct {
	*_DF1UnprotectedReadResponse

	parentBuilder *_DF1CommandBuilder

	err *utils.MultiError
}

var _ (DF1UnprotectedReadResponseBuilder) = (*_DF1UnprotectedReadResponseBuilder)(nil)

func (b *_DF1UnprotectedReadResponseBuilder) setParent(contract DF1CommandContract) {
	b.DF1CommandContract = contract
}

func (b *_DF1UnprotectedReadResponseBuilder) WithMandatoryFields(data []byte) DF1UnprotectedReadResponseBuilder {
	return b.WithData(data...)
}

func (b *_DF1UnprotectedReadResponseBuilder) WithData(data ...byte) DF1UnprotectedReadResponseBuilder {
	b.Data = data
	return b
}

func (b *_DF1UnprotectedReadResponseBuilder) Build() (DF1UnprotectedReadResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DF1UnprotectedReadResponse.deepCopy(), nil
}

func (b *_DF1UnprotectedReadResponseBuilder) MustBuild() DF1UnprotectedReadResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_DF1UnprotectedReadResponseBuilder) Done() DF1CommandBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewDF1CommandBuilder().(*_DF1CommandBuilder)
	}
	return b.parentBuilder
}

func (b *_DF1UnprotectedReadResponseBuilder) buildForDF1Command() (DF1Command, error) {
	return b.Build()
}

func (b *_DF1UnprotectedReadResponseBuilder) DeepCopy() any {
	_copy := b.CreateDF1UnprotectedReadResponseBuilder().(*_DF1UnprotectedReadResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDF1UnprotectedReadResponseBuilder creates a DF1UnprotectedReadResponseBuilder
func (b *_DF1UnprotectedReadResponse) CreateDF1UnprotectedReadResponseBuilder() DF1UnprotectedReadResponseBuilder {
	if b == nil {
		return NewDF1UnprotectedReadResponseBuilder()
	}
	return &_DF1UnprotectedReadResponseBuilder{_DF1UnprotectedReadResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1UnprotectedReadResponse) GetCommandCode() uint8 {
	return 0x41
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1UnprotectedReadResponse) GetParent() DF1CommandContract {
	return m.DF1CommandContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1UnprotectedReadResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDF1UnprotectedReadResponse(structType any) DF1UnprotectedReadResponse {
	if casted, ok := structType.(DF1UnprotectedReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DF1UnprotectedReadResponse) GetTypeName() string {
	return "DF1UnprotectedReadResponse"
}

func (m *_DF1UnprotectedReadResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.DF1CommandContract.(*_DF1Command).getLengthInBits(ctx))

	// Manual Array Field (data)
	lengthInBits += uint16(DataLength(ctx, m.GetData()))

	return lengthInBits
}

func (m *_DF1UnprotectedReadResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DF1UnprotectedReadResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_DF1Command) (__dF1UnprotectedReadResponse DF1UnprotectedReadResponse, err error) {
	m.DF1CommandContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1UnprotectedReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := ReadManualByteArrayField(ctx, "data", readBuffer, DataTerminate(ctx, readBuffer), ReadData(ctx, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1UnprotectedReadResponse")
	}

	return m, nil
}

func (m *_DF1UnprotectedReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1UnprotectedReadResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1UnprotectedReadResponse")
		}

		if err := WriteManualArrayField[byte](ctx, "data", m.GetData(), func(ctx context.Context, writeBuffer utils.WriteBuffer, m byte) error {
			return WriteData(ctx, writeBuffer, m)
		}, writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1UnprotectedReadResponse")
		}
		return nil
	}
	return m.DF1CommandContract.(*_DF1Command).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1UnprotectedReadResponse) IsDF1UnprotectedReadResponse() {}

func (m *_DF1UnprotectedReadResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DF1UnprotectedReadResponse) deepCopy() *_DF1UnprotectedReadResponse {
	if m == nil {
		return nil
	}
	_DF1UnprotectedReadResponseCopy := &_DF1UnprotectedReadResponse{
		m.DF1CommandContract.(*_DF1Command).deepCopy(),
		utils.DeepCopySlice[byte, byte](m.Data),
	}
	m.DF1CommandContract.(*_DF1Command)._SubType = m
	return _DF1UnprotectedReadResponseCopy
}

func (m *_DF1UnprotectedReadResponse) String() string {
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
