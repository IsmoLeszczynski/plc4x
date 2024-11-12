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

// COTPParameter is the corresponding interface of COTPParameter
type COTPParameter interface {
	COTPParameterContract
	COTPParameterRequirements
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// IsCOTPParameter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameter()
	// CreateBuilder creates a COTPParameterBuilder
	CreateCOTPParameterBuilder() COTPParameterBuilder
}

// COTPParameterContract provides a set of functions which can be overwritten by a sub struct
type COTPParameterContract interface {
	// GetRest() returns a parser argument
	GetRest() uint8
	// IsCOTPParameter is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPParameter()
	// CreateBuilder creates a COTPParameterBuilder
	CreateCOTPParameterBuilder() COTPParameterBuilder
}

// COTPParameterRequirements provides a set of functions which need to be implemented by a sub struct
type COTPParameterRequirements interface {
	GetLengthInBits(ctx context.Context) uint16
	GetLengthInBytes(ctx context.Context) uint16
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
}

// _COTPParameter is the data-structure of this message
type _COTPParameter struct {
	_SubType interface {
		COTPParameterContract
		COTPParameterRequirements
	}

	// Arguments.
	Rest uint8
}

var _ COTPParameterContract = (*_COTPParameter)(nil)

// NewCOTPParameter factory function for _COTPParameter
func NewCOTPParameter(rest uint8) *_COTPParameter {
	return &_COTPParameter{Rest: rest}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPParameterBuilder is a builder for COTPParameter
type COTPParameterBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() COTPParameterBuilder
	// AsCOTPParameterTpduSize converts this build to a subType of COTPParameter. It is always possible to return to current builder using Done()
	AsCOTPParameterTpduSize() COTPParameterTpduSizeBuilder
	// AsCOTPParameterCallingTsap converts this build to a subType of COTPParameter. It is always possible to return to current builder using Done()
	AsCOTPParameterCallingTsap() COTPParameterCallingTsapBuilder
	// AsCOTPParameterCalledTsap converts this build to a subType of COTPParameter. It is always possible to return to current builder using Done()
	AsCOTPParameterCalledTsap() COTPParameterCalledTsapBuilder
	// AsCOTPParameterChecksum converts this build to a subType of COTPParameter. It is always possible to return to current builder using Done()
	AsCOTPParameterChecksum() COTPParameterChecksumBuilder
	// AsCOTPParameterDisconnectAdditionalInformation converts this build to a subType of COTPParameter. It is always possible to return to current builder using Done()
	AsCOTPParameterDisconnectAdditionalInformation() COTPParameterDisconnectAdditionalInformationBuilder
	// Build builds the COTPParameter or returns an error if something is wrong
	PartialBuild() (COTPParameterContract, error)
	// MustBuild does the same as Build but panics on error
	PartialMustBuild() COTPParameterContract
	// Build builds the COTPParameter or returns an error if something is wrong
	Build() (COTPParameter, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPParameter
}

// NewCOTPParameterBuilder() creates a COTPParameterBuilder
func NewCOTPParameterBuilder() COTPParameterBuilder {
	return &_COTPParameterBuilder{_COTPParameter: new(_COTPParameter)}
}

type _COTPParameterChildBuilder interface {
	utils.Copyable
	setParent(COTPParameterContract)
	buildForCOTPParameter() (COTPParameter, error)
}

type _COTPParameterBuilder struct {
	*_COTPParameter

	childBuilder _COTPParameterChildBuilder

	err *utils.MultiError
}

var _ (COTPParameterBuilder) = (*_COTPParameterBuilder)(nil)

func (b *_COTPParameterBuilder) WithMandatoryFields() COTPParameterBuilder {
	return b
}

func (b *_COTPParameterBuilder) PartialBuild() (COTPParameterContract, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPParameter.deepCopy(), nil
}

func (b *_COTPParameterBuilder) PartialMustBuild() COTPParameterContract {
	build, err := b.PartialBuild()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPParameterBuilder) AsCOTPParameterTpduSize() COTPParameterTpduSizeBuilder {
	if cb, ok := b.childBuilder.(COTPParameterTpduSizeBuilder); ok {
		return cb
	}
	cb := NewCOTPParameterTpduSizeBuilder().(*_COTPParameterTpduSizeBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_COTPParameterBuilder) AsCOTPParameterCallingTsap() COTPParameterCallingTsapBuilder {
	if cb, ok := b.childBuilder.(COTPParameterCallingTsapBuilder); ok {
		return cb
	}
	cb := NewCOTPParameterCallingTsapBuilder().(*_COTPParameterCallingTsapBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_COTPParameterBuilder) AsCOTPParameterCalledTsap() COTPParameterCalledTsapBuilder {
	if cb, ok := b.childBuilder.(COTPParameterCalledTsapBuilder); ok {
		return cb
	}
	cb := NewCOTPParameterCalledTsapBuilder().(*_COTPParameterCalledTsapBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_COTPParameterBuilder) AsCOTPParameterChecksum() COTPParameterChecksumBuilder {
	if cb, ok := b.childBuilder.(COTPParameterChecksumBuilder); ok {
		return cb
	}
	cb := NewCOTPParameterChecksumBuilder().(*_COTPParameterChecksumBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_COTPParameterBuilder) AsCOTPParameterDisconnectAdditionalInformation() COTPParameterDisconnectAdditionalInformationBuilder {
	if cb, ok := b.childBuilder.(COTPParameterDisconnectAdditionalInformationBuilder); ok {
		return cb
	}
	cb := NewCOTPParameterDisconnectAdditionalInformationBuilder().(*_COTPParameterDisconnectAdditionalInformationBuilder)
	cb.parentBuilder = b
	b.childBuilder = cb
	return cb
}

func (b *_COTPParameterBuilder) Build() (COTPParameter, error) {
	v, err := b.PartialBuild()
	if err != nil {
		return nil, errors.Wrap(err, "error occurred during partial build")
	}
	if b.childBuilder == nil {
		return nil, errors.New("no child builder present")
	}
	b.childBuilder.setParent(v)
	return b.childBuilder.buildForCOTPParameter()
}

func (b *_COTPParameterBuilder) MustBuild() COTPParameter {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_COTPParameterBuilder) DeepCopy() any {
	_copy := b.CreateCOTPParameterBuilder().(*_COTPParameterBuilder)
	_copy.childBuilder = b.childBuilder.DeepCopy().(_COTPParameterChildBuilder)
	_copy.childBuilder.setParent(_copy)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPParameterBuilder creates a COTPParameterBuilder
func (b *_COTPParameter) CreateCOTPParameterBuilder() COTPParameterBuilder {
	if b == nil {
		return NewCOTPParameterBuilder()
	}
	return &_COTPParameterBuilder{_COTPParameter: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPParameter(structType any) COTPParameter {
	if casted, ok := structType.(COTPParameter); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameter) GetTypeName() string {
	return "COTPParameter"
}

func (m *_COTPParameter) getLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8

	// Implicit Field (parameterLength)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameter) GetLengthInBits(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx)
}

func (m *_COTPParameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m._SubType.GetLengthInBits(ctx) / 8
}

func COTPParameterParse[T COTPParameter](ctx context.Context, theBytes []byte, rest uint8) (T, error) {
	return COTPParameterParseWithBuffer[T](ctx, utils.NewReadBufferByteBased(theBytes), rest)
}

func COTPParameterParseWithBufferProducer[T COTPParameter](rest uint8) func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (T, error) {
		v, err := COTPParameterParseWithBuffer[T](ctx, readBuffer, rest)
		if err != nil {
			var zero T
			return zero, err
		}
		return v, nil
	}
}

func COTPParameterParseWithBuffer[T COTPParameter](ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (T, error) {
	v, err := (&_COTPParameter{Rest: rest}).parse(ctx, readBuffer, rest)
	if err != nil {
		var zero T
		return zero, err
	}
	vc, ok := v.(T)
	if !ok {
		var zero T
		return zero, errors.Errorf("Unexpected type %T. Expected type %T", v, *new(T))
	}
	return vc, nil
}

func (m *_COTPParameter) parse(ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (__cOTPParameter COTPParameter, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	parameterType, err := ReadDiscriminatorField[uint8](ctx, "parameterType", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterType' field"))
	}

	parameterLength, err := ReadImplicitField[uint8](ctx, "parameterLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'parameterLength' field"))
	}
	_ = parameterLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _child COTPParameter
	switch {
	case parameterType == 0xC0: // COTPParameterTpduSize
		if _child, err = new(_COTPParameterTpduSize).parse(ctx, readBuffer, m, rest); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPParameterTpduSize for type-switch of COTPParameter")
		}
	case parameterType == 0xC1: // COTPParameterCallingTsap
		if _child, err = new(_COTPParameterCallingTsap).parse(ctx, readBuffer, m, rest); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPParameterCallingTsap for type-switch of COTPParameter")
		}
	case parameterType == 0xC2: // COTPParameterCalledTsap
		if _child, err = new(_COTPParameterCalledTsap).parse(ctx, readBuffer, m, rest); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPParameterCalledTsap for type-switch of COTPParameter")
		}
	case parameterType == 0xC3: // COTPParameterChecksum
		if _child, err = new(_COTPParameterChecksum).parse(ctx, readBuffer, m, rest); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPParameterChecksum for type-switch of COTPParameter")
		}
	case parameterType == 0xE0: // COTPParameterDisconnectAdditionalInformation
		if _child, err = new(_COTPParameterDisconnectAdditionalInformation).parse(ctx, readBuffer, m, rest); err != nil {
			return nil, errors.Wrap(err, "Error parsing sub-type COTPParameterDisconnectAdditionalInformation for type-switch of COTPParameter")
		}
	default:
		return nil, errors.Errorf("Unmapped type for parameters [parameterType=%v]", parameterType)
	}

	if closeErr := readBuffer.CloseContext("COTPParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameter")
	}

	return _child, nil
}

func (pm *_COTPParameter) serializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPParameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("COTPParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for COTPParameter")
	}

	if err := WriteDiscriminatorField(ctx, "parameterType", m.GetParameterType(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'parameterType' field")
	}
	parameterLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
	if err := WriteImplicitField(ctx, "parameterLength", parameterLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
		return errors.Wrap(err, "Error serializing 'parameterLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("COTPParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for COTPParameter")
	}
	return nil
}

////
// Arguments Getter

func (m *_COTPParameter) GetRest() uint8 {
	return m.Rest
}

//
////

func (m *_COTPParameter) IsCOTPParameter() {}

func (m *_COTPParameter) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPParameter) deepCopy() *_COTPParameter {
	if m == nil {
		return nil
	}
	_COTPParameterCopy := &_COTPParameter{
		nil, // will be set by child
		m.Rest,
	}
	return _COTPParameterCopy
}
