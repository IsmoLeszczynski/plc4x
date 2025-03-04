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

// BACnetPropertyStatesLightningInProgress is the corresponding interface of BACnetPropertyStatesLightningInProgress
type BACnetPropertyStatesLightningInProgress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetLightningInProgress returns LightningInProgress (property field)
	GetLightningInProgress() BACnetLightingInProgressTagged
	// IsBACnetPropertyStatesLightningInProgress is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLightningInProgress()
	// CreateBuilder creates a BACnetPropertyStatesLightningInProgressBuilder
	CreateBACnetPropertyStatesLightningInProgressBuilder() BACnetPropertyStatesLightningInProgressBuilder
}

// _BACnetPropertyStatesLightningInProgress is the data-structure of this message
type _BACnetPropertyStatesLightningInProgress struct {
	BACnetPropertyStatesContract
	LightningInProgress BACnetLightingInProgressTagged
}

var _ BACnetPropertyStatesLightningInProgress = (*_BACnetPropertyStatesLightningInProgress)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLightningInProgress)(nil)

// NewBACnetPropertyStatesLightningInProgress factory function for _BACnetPropertyStatesLightningInProgress
func NewBACnetPropertyStatesLightningInProgress(peekedTagHeader BACnetTagHeader, lightningInProgress BACnetLightingInProgressTagged) *_BACnetPropertyStatesLightningInProgress {
	if lightningInProgress == nil {
		panic("lightningInProgress of type BACnetLightingInProgressTagged for BACnetPropertyStatesLightningInProgress must not be nil")
	}
	_result := &_BACnetPropertyStatesLightningInProgress{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LightningInProgress:          lightningInProgress,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesLightningInProgressBuilder is a builder for BACnetPropertyStatesLightningInProgress
type BACnetPropertyStatesLightningInProgressBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(lightningInProgress BACnetLightingInProgressTagged) BACnetPropertyStatesLightningInProgressBuilder
	// WithLightningInProgress adds LightningInProgress (property field)
	WithLightningInProgress(BACnetLightingInProgressTagged) BACnetPropertyStatesLightningInProgressBuilder
	// WithLightningInProgressBuilder adds LightningInProgress (property field) which is build by the builder
	WithLightningInProgressBuilder(func(BACnetLightingInProgressTaggedBuilder) BACnetLightingInProgressTaggedBuilder) BACnetPropertyStatesLightningInProgressBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesLightningInProgress or returns an error if something is wrong
	Build() (BACnetPropertyStatesLightningInProgress, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesLightningInProgress
}

// NewBACnetPropertyStatesLightningInProgressBuilder() creates a BACnetPropertyStatesLightningInProgressBuilder
func NewBACnetPropertyStatesLightningInProgressBuilder() BACnetPropertyStatesLightningInProgressBuilder {
	return &_BACnetPropertyStatesLightningInProgressBuilder{_BACnetPropertyStatesLightningInProgress: new(_BACnetPropertyStatesLightningInProgress)}
}

type _BACnetPropertyStatesLightningInProgressBuilder struct {
	*_BACnetPropertyStatesLightningInProgress

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesLightningInProgressBuilder) = (*_BACnetPropertyStatesLightningInProgressBuilder)(nil)

func (b *_BACnetPropertyStatesLightningInProgressBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
	contract.(*_BACnetPropertyStates)._SubType = b._BACnetPropertyStatesLightningInProgress
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) WithMandatoryFields(lightningInProgress BACnetLightingInProgressTagged) BACnetPropertyStatesLightningInProgressBuilder {
	return b.WithLightningInProgress(lightningInProgress)
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) WithLightningInProgress(lightningInProgress BACnetLightingInProgressTagged) BACnetPropertyStatesLightningInProgressBuilder {
	b.LightningInProgress = lightningInProgress
	return b
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) WithLightningInProgressBuilder(builderSupplier func(BACnetLightingInProgressTaggedBuilder) BACnetLightingInProgressTaggedBuilder) BACnetPropertyStatesLightningInProgressBuilder {
	builder := builderSupplier(b.LightningInProgress.CreateBACnetLightingInProgressTaggedBuilder())
	var err error
	b.LightningInProgress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLightingInProgressTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) Build() (BACnetPropertyStatesLightningInProgress, error) {
	if b.LightningInProgress == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lightningInProgress' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesLightningInProgress.deepCopy(), nil
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) MustBuild() BACnetPropertyStatesLightningInProgress {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesLightningInProgressBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesLightningInProgressBuilder().(*_BACnetPropertyStatesLightningInProgressBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesLightningInProgressBuilder creates a BACnetPropertyStatesLightningInProgressBuilder
func (b *_BACnetPropertyStatesLightningInProgress) CreateBACnetPropertyStatesLightningInProgressBuilder() BACnetPropertyStatesLightningInProgressBuilder {
	if b == nil {
		return NewBACnetPropertyStatesLightningInProgressBuilder()
	}
	return &_BACnetPropertyStatesLightningInProgressBuilder{_BACnetPropertyStatesLightningInProgress: b.deepCopy()}
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

func (m *_BACnetPropertyStatesLightningInProgress) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningInProgress) GetLightningInProgress() BACnetLightingInProgressTagged {
	return m.LightningInProgress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningInProgress(structType any) BACnetPropertyStatesLightningInProgress {
	if casted, ok := structType.(BACnetPropertyStatesLightningInProgress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningInProgress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningInProgress) GetTypeName() string {
	return "BACnetPropertyStatesLightningInProgress"
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (lightningInProgress)
	lengthInBits += m.LightningInProgress.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLightningInProgress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLightningInProgress) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLightningInProgress BACnetPropertyStatesLightningInProgress, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningInProgress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningInProgress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	lightningInProgress, err := ReadSimpleField[BACnetLightingInProgressTagged](ctx, "lightningInProgress", ReadComplex[BACnetLightingInProgressTagged](BACnetLightingInProgressTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lightningInProgress' field"))
	}
	m.LightningInProgress = lightningInProgress

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningInProgress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningInProgress")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLightningInProgress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningInProgress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningInProgress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningInProgress")
		}

		if err := WriteSimpleField[BACnetLightingInProgressTagged](ctx, "lightningInProgress", m.GetLightningInProgress(), WriteComplex[BACnetLightingInProgressTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lightningInProgress' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningInProgress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningInProgress")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLightningInProgress) IsBACnetPropertyStatesLightningInProgress() {}

func (m *_BACnetPropertyStatesLightningInProgress) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesLightningInProgress) deepCopy() *_BACnetPropertyStatesLightningInProgress {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesLightningInProgressCopy := &_BACnetPropertyStatesLightningInProgress{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetLightingInProgressTagged](m.LightningInProgress),
	}
	_BACnetPropertyStatesLightningInProgressCopy.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesLightningInProgressCopy
}

func (m *_BACnetPropertyStatesLightningInProgress) String() string {
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
