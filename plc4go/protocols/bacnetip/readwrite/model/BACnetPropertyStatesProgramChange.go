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

// BACnetPropertyStatesProgramChange is the corresponding interface of BACnetPropertyStatesProgramChange
type BACnetPropertyStatesProgramChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetProgramChange returns ProgramChange (property field)
	GetProgramChange() BACnetProgramRequestTagged
	// IsBACnetPropertyStatesProgramChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesProgramChange()
	// CreateBuilder creates a BACnetPropertyStatesProgramChangeBuilder
	CreateBACnetPropertyStatesProgramChangeBuilder() BACnetPropertyStatesProgramChangeBuilder
}

// _BACnetPropertyStatesProgramChange is the data-structure of this message
type _BACnetPropertyStatesProgramChange struct {
	BACnetPropertyStatesContract
	ProgramChange BACnetProgramRequestTagged
}

var _ BACnetPropertyStatesProgramChange = (*_BACnetPropertyStatesProgramChange)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesProgramChange)(nil)

// NewBACnetPropertyStatesProgramChange factory function for _BACnetPropertyStatesProgramChange
func NewBACnetPropertyStatesProgramChange(peekedTagHeader BACnetTagHeader, programChange BACnetProgramRequestTagged) *_BACnetPropertyStatesProgramChange {
	if programChange == nil {
		panic("programChange of type BACnetProgramRequestTagged for BACnetPropertyStatesProgramChange must not be nil")
	}
	_result := &_BACnetPropertyStatesProgramChange{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ProgramChange:                programChange,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesProgramChangeBuilder is a builder for BACnetPropertyStatesProgramChange
type BACnetPropertyStatesProgramChangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(programChange BACnetProgramRequestTagged) BACnetPropertyStatesProgramChangeBuilder
	// WithProgramChange adds ProgramChange (property field)
	WithProgramChange(BACnetProgramRequestTagged) BACnetPropertyStatesProgramChangeBuilder
	// WithProgramChangeBuilder adds ProgramChange (property field) which is build by the builder
	WithProgramChangeBuilder(func(BACnetProgramRequestTaggedBuilder) BACnetProgramRequestTaggedBuilder) BACnetPropertyStatesProgramChangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetPropertyStatesBuilder
	// Build builds the BACnetPropertyStatesProgramChange or returns an error if something is wrong
	Build() (BACnetPropertyStatesProgramChange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesProgramChange
}

// NewBACnetPropertyStatesProgramChangeBuilder() creates a BACnetPropertyStatesProgramChangeBuilder
func NewBACnetPropertyStatesProgramChangeBuilder() BACnetPropertyStatesProgramChangeBuilder {
	return &_BACnetPropertyStatesProgramChangeBuilder{_BACnetPropertyStatesProgramChange: new(_BACnetPropertyStatesProgramChange)}
}

type _BACnetPropertyStatesProgramChangeBuilder struct {
	*_BACnetPropertyStatesProgramChange

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesProgramChangeBuilder) = (*_BACnetPropertyStatesProgramChangeBuilder)(nil)

func (b *_BACnetPropertyStatesProgramChangeBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) WithMandatoryFields(programChange BACnetProgramRequestTagged) BACnetPropertyStatesProgramChangeBuilder {
	return b.WithProgramChange(programChange)
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) WithProgramChange(programChange BACnetProgramRequestTagged) BACnetPropertyStatesProgramChangeBuilder {
	b.ProgramChange = programChange
	return b
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) WithProgramChangeBuilder(builderSupplier func(BACnetProgramRequestTaggedBuilder) BACnetProgramRequestTaggedBuilder) BACnetPropertyStatesProgramChangeBuilder {
	builder := builderSupplier(b.ProgramChange.CreateBACnetProgramRequestTaggedBuilder())
	var err error
	b.ProgramChange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetProgramRequestTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) Build() (BACnetPropertyStatesProgramChange, error) {
	if b.ProgramChange == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'programChange' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesProgramChange.deepCopy(), nil
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) MustBuild() BACnetPropertyStatesProgramChange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) Done() BACnetPropertyStatesBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetPropertyStatesBuilder().(*_BACnetPropertyStatesBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesProgramChangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesProgramChangeBuilder().(*_BACnetPropertyStatesProgramChangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesProgramChangeBuilder creates a BACnetPropertyStatesProgramChangeBuilder
func (b *_BACnetPropertyStatesProgramChange) CreateBACnetPropertyStatesProgramChangeBuilder() BACnetPropertyStatesProgramChangeBuilder {
	if b == nil {
		return NewBACnetPropertyStatesProgramChangeBuilder()
	}
	return &_BACnetPropertyStatesProgramChangeBuilder{_BACnetPropertyStatesProgramChange: b.deepCopy()}
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

func (m *_BACnetPropertyStatesProgramChange) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProgramChange) GetProgramChange() BACnetProgramRequestTagged {
	return m.ProgramChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProgramChange(structType any) BACnetPropertyStatesProgramChange {
	if casted, ok := structType.(BACnetPropertyStatesProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProgramChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProgramChange) GetTypeName() string {
	return "BACnetPropertyStatesProgramChange"
}

func (m *_BACnetPropertyStatesProgramChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (programChange)
	lengthInBits += m.ProgramChange.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesProgramChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesProgramChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesProgramChange BACnetPropertyStatesProgramChange, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	programChange, err := ReadSimpleField[BACnetProgramRequestTagged](ctx, "programChange", ReadComplex[BACnetProgramRequestTagged](BACnetProgramRequestTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'programChange' field"))
	}
	m.ProgramChange = programChange

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProgramChange")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesProgramChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProgramChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProgramChange")
		}

		if err := WriteSimpleField[BACnetProgramRequestTagged](ctx, "programChange", m.GetProgramChange(), WriteComplex[BACnetProgramRequestTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'programChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProgramChange")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesProgramChange) IsBACnetPropertyStatesProgramChange() {}

func (m *_BACnetPropertyStatesProgramChange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesProgramChange) deepCopy() *_BACnetPropertyStatesProgramChange {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesProgramChangeCopy := &_BACnetPropertyStatesProgramChange{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		utils.DeepCopy[BACnetProgramRequestTagged](m.ProgramChange),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesProgramChangeCopy
}

func (m *_BACnetPropertyStatesProgramChange) String() string {
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
