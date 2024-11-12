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

// ApduControlAck is the corresponding interface of ApduControlAck
type ApduControlAck interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ApduControl
	// IsApduControlAck is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsApduControlAck()
	// CreateBuilder creates a ApduControlAckBuilder
	CreateApduControlAckBuilder() ApduControlAckBuilder
}

// _ApduControlAck is the data-structure of this message
type _ApduControlAck struct {
	ApduControlContract
}

var _ ApduControlAck = (*_ApduControlAck)(nil)
var _ ApduControlRequirements = (*_ApduControlAck)(nil)

// NewApduControlAck factory function for _ApduControlAck
func NewApduControlAck() *_ApduControlAck {
	_result := &_ApduControlAck{
		ApduControlContract: NewApduControl(),
	}
	_result.ApduControlContract.(*_ApduControl)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ApduControlAckBuilder is a builder for ApduControlAck
type ApduControlAckBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() ApduControlAckBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() ApduControlBuilder
	// Build builds the ApduControlAck or returns an error if something is wrong
	Build() (ApduControlAck, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ApduControlAck
}

// NewApduControlAckBuilder() creates a ApduControlAckBuilder
func NewApduControlAckBuilder() ApduControlAckBuilder {
	return &_ApduControlAckBuilder{_ApduControlAck: new(_ApduControlAck)}
}

type _ApduControlAckBuilder struct {
	*_ApduControlAck

	parentBuilder *_ApduControlBuilder

	err *utils.MultiError
}

var _ (ApduControlAckBuilder) = (*_ApduControlAckBuilder)(nil)

func (b *_ApduControlAckBuilder) setParent(contract ApduControlContract) {
	b.ApduControlContract = contract
}

func (b *_ApduControlAckBuilder) WithMandatoryFields() ApduControlAckBuilder {
	return b
}

func (b *_ApduControlAckBuilder) Build() (ApduControlAck, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ApduControlAck.deepCopy(), nil
}

func (b *_ApduControlAckBuilder) MustBuild() ApduControlAck {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_ApduControlAckBuilder) Done() ApduControlBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewApduControlBuilder().(*_ApduControlBuilder)
	}
	return b.parentBuilder
}

func (b *_ApduControlAckBuilder) buildForApduControl() (ApduControl, error) {
	return b.Build()
}

func (b *_ApduControlAckBuilder) DeepCopy() any {
	_copy := b.CreateApduControlAckBuilder().(*_ApduControlAckBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateApduControlAckBuilder creates a ApduControlAckBuilder
func (b *_ApduControlAck) CreateApduControlAckBuilder() ApduControlAckBuilder {
	if b == nil {
		return NewApduControlAckBuilder()
	}
	return &_ApduControlAckBuilder{_ApduControlAck: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlAck) GetControlType() uint8 {
	return 0x2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlAck) GetParent() ApduControlContract {
	return m.ApduControlContract
}

// Deprecated: use the interface for direct cast
func CastApduControlAck(structType any) ApduControlAck {
	if casted, ok := structType.(ApduControlAck); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlAck); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlAck) GetTypeName() string {
	return "ApduControlAck"
}

func (m *_ApduControlAck) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ApduControlContract.(*_ApduControl).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_ApduControlAck) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ApduControlAck) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ApduControl) (__apduControlAck ApduControlAck, err error) {
	m.ApduControlContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlAck")
	}

	return m, nil
}

func (m *_ApduControlAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlAck) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlAck")
		}

		if popErr := writeBuffer.PopContext("ApduControlAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlAck")
		}
		return nil
	}
	return m.ApduControlContract.(*_ApduControl).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ApduControlAck) IsApduControlAck() {}

func (m *_ApduControlAck) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ApduControlAck) deepCopy() *_ApduControlAck {
	if m == nil {
		return nil
	}
	_ApduControlAckCopy := &_ApduControlAck{
		m.ApduControlContract.(*_ApduControl).deepCopy(),
	}
	m.ApduControlContract.(*_ApduControl)._SubType = m
	return _ApduControlAckCopy
}

func (m *_ApduControlAck) String() string {
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
