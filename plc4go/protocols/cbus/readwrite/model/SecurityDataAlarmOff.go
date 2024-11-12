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

// SecurityDataAlarmOff is the corresponding interface of SecurityDataAlarmOff
type SecurityDataAlarmOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataAlarmOff is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataAlarmOff()
	// CreateBuilder creates a SecurityDataAlarmOffBuilder
	CreateSecurityDataAlarmOffBuilder() SecurityDataAlarmOffBuilder
}

// _SecurityDataAlarmOff is the data-structure of this message
type _SecurityDataAlarmOff struct {
	SecurityDataContract
}

var _ SecurityDataAlarmOff = (*_SecurityDataAlarmOff)(nil)
var _ SecurityDataRequirements = (*_SecurityDataAlarmOff)(nil)

// NewSecurityDataAlarmOff factory function for _SecurityDataAlarmOff
func NewSecurityDataAlarmOff(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataAlarmOff {
	_result := &_SecurityDataAlarmOff{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataAlarmOffBuilder is a builder for SecurityDataAlarmOff
type SecurityDataAlarmOffBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataAlarmOffBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataAlarmOff or returns an error if something is wrong
	Build() (SecurityDataAlarmOff, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataAlarmOff
}

// NewSecurityDataAlarmOffBuilder() creates a SecurityDataAlarmOffBuilder
func NewSecurityDataAlarmOffBuilder() SecurityDataAlarmOffBuilder {
	return &_SecurityDataAlarmOffBuilder{_SecurityDataAlarmOff: new(_SecurityDataAlarmOff)}
}

type _SecurityDataAlarmOffBuilder struct {
	*_SecurityDataAlarmOff

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataAlarmOffBuilder) = (*_SecurityDataAlarmOffBuilder)(nil)

func (b *_SecurityDataAlarmOffBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
}

func (b *_SecurityDataAlarmOffBuilder) WithMandatoryFields() SecurityDataAlarmOffBuilder {
	return b
}

func (b *_SecurityDataAlarmOffBuilder) Build() (SecurityDataAlarmOff, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataAlarmOff.deepCopy(), nil
}

func (b *_SecurityDataAlarmOffBuilder) MustBuild() SecurityDataAlarmOff {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataAlarmOffBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataAlarmOffBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataAlarmOffBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataAlarmOffBuilder().(*_SecurityDataAlarmOffBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataAlarmOffBuilder creates a SecurityDataAlarmOffBuilder
func (b *_SecurityDataAlarmOff) CreateSecurityDataAlarmOffBuilder() SecurityDataAlarmOffBuilder {
	if b == nil {
		return NewSecurityDataAlarmOffBuilder()
	}
	return &_SecurityDataAlarmOffBuilder{_SecurityDataAlarmOff: b.deepCopy()}
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

func (m *_SecurityDataAlarmOff) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataAlarmOff(structType any) SecurityDataAlarmOff {
	if casted, ok := structType.(SecurityDataAlarmOff); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataAlarmOff); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataAlarmOff) GetTypeName() string {
	return "SecurityDataAlarmOff"
}

func (m *_SecurityDataAlarmOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataAlarmOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataAlarmOff) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataAlarmOff SecurityDataAlarmOff, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataAlarmOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataAlarmOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataAlarmOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataAlarmOff")
	}

	return m, nil
}

func (m *_SecurityDataAlarmOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataAlarmOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataAlarmOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataAlarmOff")
		}

		if popErr := writeBuffer.PopContext("SecurityDataAlarmOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataAlarmOff")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataAlarmOff) IsSecurityDataAlarmOff() {}

func (m *_SecurityDataAlarmOff) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataAlarmOff) deepCopy() *_SecurityDataAlarmOff {
	if m == nil {
		return nil
	}
	_SecurityDataAlarmOffCopy := &_SecurityDataAlarmOff{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	m.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataAlarmOffCopy
}

func (m *_SecurityDataAlarmOff) String() string {
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
