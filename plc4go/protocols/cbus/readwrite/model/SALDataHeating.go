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

// SALDataHeating is the corresponding interface of SALDataHeating
type SALDataHeating interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	SALData
	// GetHeatingData returns HeatingData (property field)
	GetHeatingData() LightingData
	// IsSALDataHeating is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSALDataHeating()
	// CreateBuilder creates a SALDataHeatingBuilder
	CreateSALDataHeatingBuilder() SALDataHeatingBuilder
}

// _SALDataHeating is the data-structure of this message
type _SALDataHeating struct {
	SALDataContract
	HeatingData LightingData
}

var _ SALDataHeating = (*_SALDataHeating)(nil)
var _ SALDataRequirements = (*_SALDataHeating)(nil)

// NewSALDataHeating factory function for _SALDataHeating
func NewSALDataHeating(salData SALData, heatingData LightingData) *_SALDataHeating {
	if heatingData == nil {
		panic("heatingData of type LightingData for SALDataHeating must not be nil")
	}
	_result := &_SALDataHeating{
		SALDataContract: NewSALData(salData),
		HeatingData:     heatingData,
	}
	_result.SALDataContract.(*_SALData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// SALDataHeatingBuilder is a builder for SALDataHeating
type SALDataHeatingBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(heatingData LightingData) SALDataHeatingBuilder
	// WithHeatingData adds HeatingData (property field)
	WithHeatingData(LightingData) SALDataHeatingBuilder
	// WithHeatingDataBuilder adds HeatingData (property field) which is build by the builder
	WithHeatingDataBuilder(func(LightingDataBuilder) LightingDataBuilder) SALDataHeatingBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() SALDataBuilder
	// Build builds the SALDataHeating or returns an error if something is wrong
	Build() (SALDataHeating, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() SALDataHeating
}

// NewSALDataHeatingBuilder() creates a SALDataHeatingBuilder
func NewSALDataHeatingBuilder() SALDataHeatingBuilder {
	return &_SALDataHeatingBuilder{_SALDataHeating: new(_SALDataHeating)}
}

type _SALDataHeatingBuilder struct {
	*_SALDataHeating

	parentBuilder *_SALDataBuilder

	err *utils.MultiError
}

var _ (SALDataHeatingBuilder) = (*_SALDataHeatingBuilder)(nil)

func (b *_SALDataHeatingBuilder) setParent(contract SALDataContract) {
	b.SALDataContract = contract
}

func (b *_SALDataHeatingBuilder) WithMandatoryFields(heatingData LightingData) SALDataHeatingBuilder {
	return b.WithHeatingData(heatingData)
}

func (b *_SALDataHeatingBuilder) WithHeatingData(heatingData LightingData) SALDataHeatingBuilder {
	b.HeatingData = heatingData
	return b
}

func (b *_SALDataHeatingBuilder) WithHeatingDataBuilder(builderSupplier func(LightingDataBuilder) LightingDataBuilder) SALDataHeatingBuilder {
	builder := builderSupplier(b.HeatingData.CreateLightingDataBuilder())
	var err error
	b.HeatingData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "LightingDataBuilder failed"))
	}
	return b
}

func (b *_SALDataHeatingBuilder) Build() (SALDataHeating, error) {
	if b.HeatingData == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'heatingData' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._SALDataHeating.deepCopy(), nil
}

func (b *_SALDataHeatingBuilder) MustBuild() SALDataHeating {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_SALDataHeatingBuilder) Done() SALDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewSALDataBuilder().(*_SALDataBuilder)
	}
	return b.parentBuilder
}

func (b *_SALDataHeatingBuilder) buildForSALData() (SALData, error) {
	return b.Build()
}

func (b *_SALDataHeatingBuilder) DeepCopy() any {
	_copy := b.CreateSALDataHeatingBuilder().(*_SALDataHeatingBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateSALDataHeatingBuilder creates a SALDataHeatingBuilder
func (b *_SALDataHeating) CreateSALDataHeatingBuilder() SALDataHeatingBuilder {
	if b == nil {
		return NewSALDataHeatingBuilder()
	}
	return &_SALDataHeatingBuilder{_SALDataHeating: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataHeating) GetApplicationId() ApplicationId {
	return ApplicationId_HEATING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataHeating) GetParent() SALDataContract {
	return m.SALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataHeating) GetHeatingData() LightingData {
	return m.HeatingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastSALDataHeating(structType any) SALDataHeating {
	if casted, ok := structType.(SALDataHeating); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataHeating); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataHeating) GetTypeName() string {
	return "SALDataHeating"
}

func (m *_SALDataHeating) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SALDataContract.(*_SALData).getLengthInBits(ctx))

	// Simple field (heatingData)
	lengthInBits += m.HeatingData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataHeating) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SALDataHeating) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SALData, applicationId ApplicationId) (__sALDataHeating SALDataHeating, err error) {
	m.SALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataHeating"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataHeating")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	heatingData, err := ReadSimpleField[LightingData](ctx, "heatingData", ReadComplex[LightingData](LightingDataParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'heatingData' field"))
	}
	m.HeatingData = heatingData

	if closeErr := readBuffer.CloseContext("SALDataHeating"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataHeating")
	}

	return m, nil
}

func (m *_SALDataHeating) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataHeating) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataHeating"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataHeating")
		}

		if err := WriteSimpleField[LightingData](ctx, "heatingData", m.GetHeatingData(), WriteComplex[LightingData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'heatingData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataHeating"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataHeating")
		}
		return nil
	}
	return m.SALDataContract.(*_SALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataHeating) IsSALDataHeating() {}

func (m *_SALDataHeating) DeepCopy() any {
	return m.deepCopy()
}

func (m *_SALDataHeating) deepCopy() *_SALDataHeating {
	if m == nil {
		return nil
	}
	_SALDataHeatingCopy := &_SALDataHeating{
		m.SALDataContract.(*_SALData).deepCopy(),
		utils.DeepCopy[LightingData](m.HeatingData),
	}
	m.SALDataContract.(*_SALData)._SubType = m
	return _SALDataHeatingCopy
}

func (m *_SALDataHeating) String() string {
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
