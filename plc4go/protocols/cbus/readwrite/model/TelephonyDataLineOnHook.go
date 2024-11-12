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

// TelephonyDataLineOnHook is the corresponding interface of TelephonyDataLineOnHook
type TelephonyDataLineOnHook interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	TelephonyData
	// IsTelephonyDataLineOnHook is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsTelephonyDataLineOnHook()
	// CreateBuilder creates a TelephonyDataLineOnHookBuilder
	CreateTelephonyDataLineOnHookBuilder() TelephonyDataLineOnHookBuilder
}

// _TelephonyDataLineOnHook is the data-structure of this message
type _TelephonyDataLineOnHook struct {
	TelephonyDataContract
}

var _ TelephonyDataLineOnHook = (*_TelephonyDataLineOnHook)(nil)
var _ TelephonyDataRequirements = (*_TelephonyDataLineOnHook)(nil)

// NewTelephonyDataLineOnHook factory function for _TelephonyDataLineOnHook
func NewTelephonyDataLineOnHook(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataLineOnHook {
	_result := &_TelephonyDataLineOnHook{
		TelephonyDataContract: NewTelephonyData(commandTypeContainer, argument),
	}
	_result.TelephonyDataContract.(*_TelephonyData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// TelephonyDataLineOnHookBuilder is a builder for TelephonyDataLineOnHook
type TelephonyDataLineOnHookBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() TelephonyDataLineOnHookBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() TelephonyDataBuilder
	// Build builds the TelephonyDataLineOnHook or returns an error if something is wrong
	Build() (TelephonyDataLineOnHook, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() TelephonyDataLineOnHook
}

// NewTelephonyDataLineOnHookBuilder() creates a TelephonyDataLineOnHookBuilder
func NewTelephonyDataLineOnHookBuilder() TelephonyDataLineOnHookBuilder {
	return &_TelephonyDataLineOnHookBuilder{_TelephonyDataLineOnHook: new(_TelephonyDataLineOnHook)}
}

type _TelephonyDataLineOnHookBuilder struct {
	*_TelephonyDataLineOnHook

	parentBuilder *_TelephonyDataBuilder

	err *utils.MultiError
}

var _ (TelephonyDataLineOnHookBuilder) = (*_TelephonyDataLineOnHookBuilder)(nil)

func (b *_TelephonyDataLineOnHookBuilder) setParent(contract TelephonyDataContract) {
	b.TelephonyDataContract = contract
}

func (b *_TelephonyDataLineOnHookBuilder) WithMandatoryFields() TelephonyDataLineOnHookBuilder {
	return b
}

func (b *_TelephonyDataLineOnHookBuilder) Build() (TelephonyDataLineOnHook, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._TelephonyDataLineOnHook.deepCopy(), nil
}

func (b *_TelephonyDataLineOnHookBuilder) MustBuild() TelephonyDataLineOnHook {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_TelephonyDataLineOnHookBuilder) Done() TelephonyDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewTelephonyDataBuilder().(*_TelephonyDataBuilder)
	}
	return b.parentBuilder
}

func (b *_TelephonyDataLineOnHookBuilder) buildForTelephonyData() (TelephonyData, error) {
	return b.Build()
}

func (b *_TelephonyDataLineOnHookBuilder) DeepCopy() any {
	_copy := b.CreateTelephonyDataLineOnHookBuilder().(*_TelephonyDataLineOnHookBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateTelephonyDataLineOnHookBuilder creates a TelephonyDataLineOnHookBuilder
func (b *_TelephonyDataLineOnHook) CreateTelephonyDataLineOnHookBuilder() TelephonyDataLineOnHookBuilder {
	if b == nil {
		return NewTelephonyDataLineOnHookBuilder()
	}
	return &_TelephonyDataLineOnHookBuilder{_TelephonyDataLineOnHook: b.deepCopy()}
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

func (m *_TelephonyDataLineOnHook) GetParent() TelephonyDataContract {
	return m.TelephonyDataContract
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataLineOnHook(structType any) TelephonyDataLineOnHook {
	if casted, ok := structType.(TelephonyDataLineOnHook); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataLineOnHook); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataLineOnHook) GetTypeName() string {
	return "TelephonyDataLineOnHook"
}

func (m *_TelephonyDataLineOnHook) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.TelephonyDataContract.(*_TelephonyData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_TelephonyDataLineOnHook) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_TelephonyDataLineOnHook) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_TelephonyData) (__telephonyDataLineOnHook TelephonyDataLineOnHook, err error) {
	m.TelephonyDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataLineOnHook"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataLineOnHook")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataLineOnHook"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataLineOnHook")
	}

	return m, nil
}

func (m *_TelephonyDataLineOnHook) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataLineOnHook) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataLineOnHook"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataLineOnHook")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataLineOnHook"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataLineOnHook")
		}
		return nil
	}
	return m.TelephonyDataContract.(*_TelephonyData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataLineOnHook) IsTelephonyDataLineOnHook() {}

func (m *_TelephonyDataLineOnHook) DeepCopy() any {
	return m.deepCopy()
}

func (m *_TelephonyDataLineOnHook) deepCopy() *_TelephonyDataLineOnHook {
	if m == nil {
		return nil
	}
	_TelephonyDataLineOnHookCopy := &_TelephonyDataLineOnHook{
		m.TelephonyDataContract.(*_TelephonyData).deepCopy(),
	}
	m.TelephonyDataContract.(*_TelephonyData)._SubType = m
	return _TelephonyDataLineOnHookCopy
}

func (m *_TelephonyDataLineOnHook) String() string {
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
