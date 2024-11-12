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

// CALReplyShort is the corresponding interface of CALReplyShort
type CALReplyShort interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CALReply
	// IsCALReplyShort is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALReplyShort()
	// CreateBuilder creates a CALReplyShortBuilder
	CreateCALReplyShortBuilder() CALReplyShortBuilder
}

// _CALReplyShort is the data-structure of this message
type _CALReplyShort struct {
	CALReplyContract
}

var _ CALReplyShort = (*_CALReplyShort)(nil)
var _ CALReplyRequirements = (*_CALReplyShort)(nil)

// NewCALReplyShort factory function for _CALReplyShort
func NewCALReplyShort(calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReplyShort {
	_result := &_CALReplyShort{
		CALReplyContract: NewCALReply(calType, calData, cBusOptions, requestContext),
	}
	_result.CALReplyContract.(*_CALReply)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CALReplyShortBuilder is a builder for CALReplyShort
type CALReplyShortBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() CALReplyShortBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CALReplyBuilder
	// Build builds the CALReplyShort or returns an error if something is wrong
	Build() (CALReplyShort, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CALReplyShort
}

// NewCALReplyShortBuilder() creates a CALReplyShortBuilder
func NewCALReplyShortBuilder() CALReplyShortBuilder {
	return &_CALReplyShortBuilder{_CALReplyShort: new(_CALReplyShort)}
}

type _CALReplyShortBuilder struct {
	*_CALReplyShort

	parentBuilder *_CALReplyBuilder

	err *utils.MultiError
}

var _ (CALReplyShortBuilder) = (*_CALReplyShortBuilder)(nil)

func (b *_CALReplyShortBuilder) setParent(contract CALReplyContract) {
	b.CALReplyContract = contract
}

func (b *_CALReplyShortBuilder) WithMandatoryFields() CALReplyShortBuilder {
	return b
}

func (b *_CALReplyShortBuilder) Build() (CALReplyShort, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CALReplyShort.deepCopy(), nil
}

func (b *_CALReplyShortBuilder) MustBuild() CALReplyShort {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CALReplyShortBuilder) Done() CALReplyBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCALReplyBuilder().(*_CALReplyBuilder)
	}
	return b.parentBuilder
}

func (b *_CALReplyShortBuilder) buildForCALReply() (CALReply, error) {
	return b.Build()
}

func (b *_CALReplyShortBuilder) DeepCopy() any {
	_copy := b.CreateCALReplyShortBuilder().(*_CALReplyShortBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCALReplyShortBuilder creates a CALReplyShortBuilder
func (b *_CALReplyShort) CreateCALReplyShortBuilder() CALReplyShortBuilder {
	if b == nil {
		return NewCALReplyShortBuilder()
	}
	return &_CALReplyShortBuilder{_CALReplyShort: b.deepCopy()}
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

func (m *_CALReplyShort) GetParent() CALReplyContract {
	return m.CALReplyContract
}

// Deprecated: use the interface for direct cast
func CastCALReplyShort(structType any) CALReplyShort {
	if casted, ok := structType.(CALReplyShort); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyShort); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyShort) GetTypeName() string {
	return "CALReplyShort"
}

func (m *_CALReplyShort) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALReplyContract.(*_CALReply).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_CALReplyShort) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALReplyShort) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALReply, cBusOptions CBusOptions, requestContext RequestContext) (__cALReplyShort CALReplyShort, err error) {
	m.CALReplyContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyShort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyShort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALReplyShort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyShort")
	}

	return m, nil
}

func (m *_CALReplyShort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALReplyShort) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyShort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyShort")
		}

		if popErr := writeBuffer.PopContext("CALReplyShort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyShort")
		}
		return nil
	}
	return m.CALReplyContract.(*_CALReply).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALReplyShort) IsCALReplyShort() {}

func (m *_CALReplyShort) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CALReplyShort) deepCopy() *_CALReplyShort {
	if m == nil {
		return nil
	}
	_CALReplyShortCopy := &_CALReplyShort{
		m.CALReplyContract.(*_CALReply).deepCopy(),
	}
	m.CALReplyContract.(*_CALReply)._SubType = m
	return _CALReplyShortCopy
}

func (m *_CALReplyShort) String() string {
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
