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

// Constant values.
const ServerErrorReply_ERRORMARKER byte = 0x21

// ServerErrorReply is the corresponding interface of ServerErrorReply
type ServerErrorReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ReplyOrConfirmation
	// IsServerErrorReply is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsServerErrorReply()
	// CreateBuilder creates a ServerErrorReplyBuilder
	CreateServerErrorReplyBuilder() ServerErrorReplyBuilder
}

// _ServerErrorReply is the data-structure of this message
type _ServerErrorReply struct {
	ReplyOrConfirmationContract
}

var _ ServerErrorReply = (*_ServerErrorReply)(nil)
var _ ReplyOrConfirmationRequirements = (*_ServerErrorReply)(nil)

// NewServerErrorReply factory function for _ServerErrorReply
func NewServerErrorReply(peekedByte byte, cBusOptions CBusOptions, requestContext RequestContext) *_ServerErrorReply {
	_result := &_ServerErrorReply{
		ReplyOrConfirmationContract: NewReplyOrConfirmation(peekedByte, cBusOptions, requestContext),
	}
	_result.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ServerErrorReplyBuilder is a builder for ServerErrorReply
type ServerErrorReplyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ServerErrorReplyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ReplyOrConfirmationBuilder
	// Build builds the ServerErrorReply or returns an error if something is wrong
	Build() (ServerErrorReply, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ServerErrorReply
}

// NewServerErrorReplyBuilder() creates a ServerErrorReplyBuilder
func NewServerErrorReplyBuilder() ServerErrorReplyBuilder {
	return &_ServerErrorReplyBuilder{_ServerErrorReply: new(_ServerErrorReply)}
}

type _ServerErrorReplyBuilder struct {
	*_ServerErrorReply

	parentBuilder *_ReplyOrConfirmationBuilder

	err *utils.MultiError
}

var _ (ServerErrorReplyBuilder) = (*_ServerErrorReplyBuilder)(nil)

func (b *_ServerErrorReplyBuilder) setParent(contract ReplyOrConfirmationContract) {
	b.ReplyOrConfirmationContract = contract
}

func (b *_ServerErrorReplyBuilder) WithMandatoryFields() ServerErrorReplyBuilder {
	return b
}

func (b *_ServerErrorReplyBuilder) Build() (ServerErrorReply, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ServerErrorReply.deepCopy(), nil
}

func (b *_ServerErrorReplyBuilder) MustBuild() ServerErrorReply {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ServerErrorReplyBuilder) Done() ReplyOrConfirmationBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewReplyOrConfirmationBuilder().(*_ReplyOrConfirmationBuilder)
	}
	return b.parentBuilder
}

func (b *_ServerErrorReplyBuilder) buildForReplyOrConfirmation() (ReplyOrConfirmation, error) {
	return b.Build()
}

func (b *_ServerErrorReplyBuilder) DeepCopy() any {
	_copy := b.CreateServerErrorReplyBuilder().(*_ServerErrorReplyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateServerErrorReplyBuilder creates a ServerErrorReplyBuilder
func (b *_ServerErrorReply) CreateServerErrorReplyBuilder() ServerErrorReplyBuilder {
	if b == nil {
		return NewServerErrorReplyBuilder()
	}
	return &_ServerErrorReplyBuilder{_ServerErrorReply: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ServerErrorReply) GetParent() ReplyOrConfirmationContract {
	return m.ReplyOrConfirmationContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ServerErrorReply) GetErrorMarker() byte {
	return ServerErrorReply_ERRORMARKER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastServerErrorReply(structType any) ServerErrorReply {
	if casted, ok := structType.(ServerErrorReply); ok {
		return casted
	}
	if casted, ok := structType.(*ServerErrorReply); ok {
		return *casted
	}
	return nil
}

func (m *_ServerErrorReply) GetTypeName() string {
	return "ServerErrorReply"
}

func (m *_ServerErrorReply) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).getLengthInBits(ctx))

	// Const Field (errorMarker)
	lengthInBits += 8

	return lengthInBits
}

func (m *_ServerErrorReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ServerErrorReply) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ReplyOrConfirmation, cBusOptions CBusOptions, requestContext RequestContext) (__serverErrorReply ServerErrorReply, err error) {
	m.ReplyOrConfirmationContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ServerErrorReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ServerErrorReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	errorMarker, err := ReadConstField[byte](ctx, "errorMarker", ReadByte(readBuffer, 8), ServerErrorReply_ERRORMARKER)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'errorMarker' field"))
	}
	_ = errorMarker

	if closeErr := readBuffer.CloseContext("ServerErrorReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ServerErrorReply")
	}

	return m, nil
}

func (m *_ServerErrorReply) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ServerErrorReply) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ServerErrorReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ServerErrorReply")
		}

		if err := WriteConstField(ctx, "errorMarker", ServerErrorReply_ERRORMARKER, WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'errorMarker' field")
		}

		if popErr := writeBuffer.PopContext("ServerErrorReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ServerErrorReply")
		}
		return nil
	}
	return m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ServerErrorReply) IsServerErrorReply() {}

func (m *_ServerErrorReply) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ServerErrorReply) deepCopy() *_ServerErrorReply {
	if m == nil {
		return nil
	}
	_ServerErrorReplyCopy := &_ServerErrorReply{
		m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation).deepCopy(),
	}
	m.ReplyOrConfirmationContract.(*_ReplyOrConfirmation)._SubType = m
	return _ServerErrorReplyCopy
}

func (m *_ServerErrorReply) String() string {
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
