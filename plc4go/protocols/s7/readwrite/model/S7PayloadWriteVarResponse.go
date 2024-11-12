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

// S7PayloadWriteVarResponse is the corresponding interface of S7PayloadWriteVarResponse
type S7PayloadWriteVarResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	S7Payload
	// GetItems returns Items (property field)
	GetItems() []S7VarPayloadStatusItem
	// IsS7PayloadWriteVarResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsS7PayloadWriteVarResponse()
	// CreateBuilder creates a S7PayloadWriteVarResponseBuilder
	CreateS7PayloadWriteVarResponseBuilder() S7PayloadWriteVarResponseBuilder
}

// _S7PayloadWriteVarResponse is the data-structure of this message
type _S7PayloadWriteVarResponse struct {
	S7PayloadContract
	Items []S7VarPayloadStatusItem
}

var _ S7PayloadWriteVarResponse = (*_S7PayloadWriteVarResponse)(nil)
var _ S7PayloadRequirements = (*_S7PayloadWriteVarResponse)(nil)

// NewS7PayloadWriteVarResponse factory function for _S7PayloadWriteVarResponse
func NewS7PayloadWriteVarResponse(items []S7VarPayloadStatusItem, parameter S7Parameter) *_S7PayloadWriteVarResponse {
	_result := &_S7PayloadWriteVarResponse{
		S7PayloadContract: NewS7Payload(parameter),
		Items:             items,
	}
	_result.S7PayloadContract.(*_S7Payload)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// S7PayloadWriteVarResponseBuilder is a builder for S7PayloadWriteVarResponse
type S7PayloadWriteVarResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(items []S7VarPayloadStatusItem) S7PayloadWriteVarResponseBuilder
	// WithItems adds Items (property field)
	WithItems(...S7VarPayloadStatusItem) S7PayloadWriteVarResponseBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() S7PayloadBuilder
	// Build builds the S7PayloadWriteVarResponse or returns an error if something is wrong
	Build() (S7PayloadWriteVarResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() S7PayloadWriteVarResponse
}

// NewS7PayloadWriteVarResponseBuilder() creates a S7PayloadWriteVarResponseBuilder
func NewS7PayloadWriteVarResponseBuilder() S7PayloadWriteVarResponseBuilder {
	return &_S7PayloadWriteVarResponseBuilder{_S7PayloadWriteVarResponse: new(_S7PayloadWriteVarResponse)}
}

type _S7PayloadWriteVarResponseBuilder struct {
	*_S7PayloadWriteVarResponse

	parentBuilder *_S7PayloadBuilder

	err *utils.MultiError
}

var _ (S7PayloadWriteVarResponseBuilder) = (*_S7PayloadWriteVarResponseBuilder)(nil)

func (b *_S7PayloadWriteVarResponseBuilder) setParent(contract S7PayloadContract) {
	b.S7PayloadContract = contract
}

func (b *_S7PayloadWriteVarResponseBuilder) WithMandatoryFields(items []S7VarPayloadStatusItem) S7PayloadWriteVarResponseBuilder {
	return b.WithItems(items...)
}

func (b *_S7PayloadWriteVarResponseBuilder) WithItems(items ...S7VarPayloadStatusItem) S7PayloadWriteVarResponseBuilder {
	b.Items = items
	return b
}

func (b *_S7PayloadWriteVarResponseBuilder) Build() (S7PayloadWriteVarResponse, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._S7PayloadWriteVarResponse.deepCopy(), nil
}

func (b *_S7PayloadWriteVarResponseBuilder) MustBuild() S7PayloadWriteVarResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_S7PayloadWriteVarResponseBuilder) Done() S7PayloadBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewS7PayloadBuilder().(*_S7PayloadBuilder)
	}
	return b.parentBuilder
}

func (b *_S7PayloadWriteVarResponseBuilder) buildForS7Payload() (S7Payload, error) {
	return b.Build()
}

func (b *_S7PayloadWriteVarResponseBuilder) DeepCopy() any {
	_copy := b.CreateS7PayloadWriteVarResponseBuilder().(*_S7PayloadWriteVarResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateS7PayloadWriteVarResponseBuilder creates a S7PayloadWriteVarResponseBuilder
func (b *_S7PayloadWriteVarResponse) CreateS7PayloadWriteVarResponseBuilder() S7PayloadWriteVarResponseBuilder {
	if b == nil {
		return NewS7PayloadWriteVarResponseBuilder()
	}
	return &_S7PayloadWriteVarResponseBuilder{_S7PayloadWriteVarResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetParameterParameterType() uint8 {
	return 0x05
}

func (m *_S7PayloadWriteVarResponse) GetMessageType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadWriteVarResponse) GetParent() S7PayloadContract {
	return m.S7PayloadContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadWriteVarResponse) GetItems() []S7VarPayloadStatusItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastS7PayloadWriteVarResponse(structType any) S7PayloadWriteVarResponse {
	if casted, ok := structType.(S7PayloadWriteVarResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadWriteVarResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadWriteVarResponse) GetTypeName() string {
	return "S7PayloadWriteVarResponse"
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.S7PayloadContract.(*_S7Payload).getLengthInBits(ctx))

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadWriteVarResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_S7PayloadWriteVarResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_S7Payload, messageType uint8, parameter S7Parameter) (__s7PayloadWriteVarResponse S7PayloadWriteVarResponse, err error) {
	m.S7PayloadContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadWriteVarResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadWriteVarResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	items, err := ReadCountArrayField[S7VarPayloadStatusItem](ctx, "items", ReadComplex[S7VarPayloadStatusItem](S7VarPayloadStatusItemParseWithBuffer, readBuffer), uint64(CastS7ParameterWriteVarResponse(parameter).GetNumItems()))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'items' field"))
	}
	m.Items = items

	if closeErr := readBuffer.CloseContext("S7PayloadWriteVarResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadWriteVarResponse")
	}

	return m, nil
}

func (m *_S7PayloadWriteVarResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadWriteVarResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadWriteVarResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadWriteVarResponse")
		}

		if err := WriteComplexTypeArrayField(ctx, "items", m.GetItems(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'items' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadWriteVarResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadWriteVarResponse")
		}
		return nil
	}
	return m.S7PayloadContract.(*_S7Payload).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadWriteVarResponse) IsS7PayloadWriteVarResponse() {}

func (m *_S7PayloadWriteVarResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_S7PayloadWriteVarResponse) deepCopy() *_S7PayloadWriteVarResponse {
	if m == nil {
		return nil
	}
	_S7PayloadWriteVarResponseCopy := &_S7PayloadWriteVarResponse{
		m.S7PayloadContract.(*_S7Payload).deepCopy(),
		utils.DeepCopySlice[S7VarPayloadStatusItem, S7VarPayloadStatusItem](m.Items),
	}
	m.S7PayloadContract.(*_S7Payload)._SubType = m
	return _S7PayloadWriteVarResponseCopy
}

func (m *_S7PayloadWriteVarResponse) String() string {
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
