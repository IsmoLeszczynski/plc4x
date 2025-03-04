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

// SecurityDataStatus1Request is the corresponding interface of SecurityDataStatus1Request
type SecurityDataStatus1Request interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SecurityData
	// IsSecurityDataStatus1Request is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataStatus1Request()
	// CreateBuilder creates a SecurityDataStatus1RequestBuilder
	CreateSecurityDataStatus1RequestBuilder() SecurityDataStatus1RequestBuilder
}

// _SecurityDataStatus1Request is the data-structure of this message
type _SecurityDataStatus1Request struct {
	SecurityDataContract
}

var _ SecurityDataStatus1Request = (*_SecurityDataStatus1Request)(nil)
var _ SecurityDataRequirements = (*_SecurityDataStatus1Request)(nil)

// NewSecurityDataStatus1Request factory function for _SecurityDataStatus1Request
func NewSecurityDataStatus1Request(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataStatus1Request {
	_result := &_SecurityDataStatus1Request{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SecurityDataStatus1RequestBuilder is a builder for SecurityDataStatus1Request
type SecurityDataStatus1RequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() SecurityDataStatus1RequestBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SecurityDataBuilder
	// Build builds the SecurityDataStatus1Request or returns an error if something is wrong
	Build() (SecurityDataStatus1Request, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SecurityDataStatus1Request
}

// NewSecurityDataStatus1RequestBuilder() creates a SecurityDataStatus1RequestBuilder
func NewSecurityDataStatus1RequestBuilder() SecurityDataStatus1RequestBuilder {
	return &_SecurityDataStatus1RequestBuilder{_SecurityDataStatus1Request: new(_SecurityDataStatus1Request)}
}

type _SecurityDataStatus1RequestBuilder struct {
	*_SecurityDataStatus1Request

	parentBuilder *_SecurityDataBuilder

	err *utils.MultiError
}

var _ (SecurityDataStatus1RequestBuilder) = (*_SecurityDataStatus1RequestBuilder)(nil)

func (b *_SecurityDataStatus1RequestBuilder) setParent(contract SecurityDataContract) {
	b.SecurityDataContract = contract
	contract.(*_SecurityData)._SubType = b._SecurityDataStatus1Request
}

func (b *_SecurityDataStatus1RequestBuilder) WithMandatoryFields() SecurityDataStatus1RequestBuilder {
	return b
}

func (b *_SecurityDataStatus1RequestBuilder) Build() (SecurityDataStatus1Request, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SecurityDataStatus1Request.deepCopy(), nil
}

func (b *_SecurityDataStatus1RequestBuilder) MustBuild() SecurityDataStatus1Request {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SecurityDataStatus1RequestBuilder) Done() SecurityDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSecurityDataBuilder().(*_SecurityDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SecurityDataStatus1RequestBuilder) buildForSecurityData() (SecurityData, error) {
	return b.Build()
}

func (b *_SecurityDataStatus1RequestBuilder) DeepCopy() any {
	_copy := b.CreateSecurityDataStatus1RequestBuilder().(*_SecurityDataStatus1RequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSecurityDataStatus1RequestBuilder creates a SecurityDataStatus1RequestBuilder
func (b *_SecurityDataStatus1Request) CreateSecurityDataStatus1RequestBuilder() SecurityDataStatus1RequestBuilder {
	if b == nil {
		return NewSecurityDataStatus1RequestBuilder()
	}
	return &_SecurityDataStatus1RequestBuilder{_SecurityDataStatus1Request: b.deepCopy()}
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

func (m *_SecurityDataStatus1Request) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

// Deprecated: use the interface for direct cast
func CastSecurityDataStatus1Request(structType any) SecurityDataStatus1Request {
	if casted, ok := structType.(SecurityDataStatus1Request); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataStatus1Request); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataStatus1Request) GetTypeName() string {
	return "SecurityDataStatus1Request"
}

func (m *_SecurityDataStatus1Request) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataStatus1Request) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataStatus1Request) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData) (__securityDataStatus1Request SecurityDataStatus1Request, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataStatus1Request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataStatus1Request")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataStatus1Request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataStatus1Request")
	}

	return m, nil
}

func (m *_SecurityDataStatus1Request) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataStatus1Request) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataStatus1Request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataStatus1Request")
		}

		if popErr := writeBuffer.PopContext("SecurityDataStatus1Request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataStatus1Request")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataStatus1Request) IsSecurityDataStatus1Request() {}

func (m *_SecurityDataStatus1Request) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SecurityDataStatus1Request) deepCopy() *_SecurityDataStatus1Request {
	if m == nil {
		return nil
	}
	_SecurityDataStatus1RequestCopy := &_SecurityDataStatus1Request{
		m.SecurityDataContract.(*_SecurityData).deepCopy(),
	}
	_SecurityDataStatus1RequestCopy.SecurityDataContract.(*_SecurityData)._SubType = m
	return _SecurityDataStatus1RequestCopy
}

func (m *_SecurityDataStatus1Request) String() string {
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
