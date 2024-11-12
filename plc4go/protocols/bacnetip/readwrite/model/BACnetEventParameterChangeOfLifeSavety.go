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

// BACnetEventParameterChangeOfLifeSavety is the corresponding interface of BACnetEventParameterChangeOfLifeSavety
type BACnetEventParameterChangeOfLifeSavety interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterChangeOfLifeSavety is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterChangeOfLifeSavety()
	// CreateBuilder creates a BACnetEventParameterChangeOfLifeSavetyBuilder
	CreateBACnetEventParameterChangeOfLifeSavetyBuilder() BACnetEventParameterChangeOfLifeSavetyBuilder
}

// _BACnetEventParameterChangeOfLifeSavety is the data-structure of this message
type _BACnetEventParameterChangeOfLifeSavety struct {
	BACnetEventParameterContract
	OpeningTag                  BACnetOpeningTag
	TimeDelay                   BACnetContextTagUnsignedInteger
	ListOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	ListOfAlarmValues           BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	ModePropertyReference       BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag                  BACnetClosingTag
}

var _ BACnetEventParameterChangeOfLifeSavety = (*_BACnetEventParameterChangeOfLifeSavety)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterChangeOfLifeSavety)(nil)

// NewBACnetEventParameterChangeOfLifeSavety factory function for _BACnetEventParameterChangeOfLifeSavety
func NewBACnetEventParameterChangeOfLifeSavety(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) *_BACnetEventParameterChangeOfLifeSavety {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	if listOfLifeSavetyAlarmValues == nil {
		panic("listOfLifeSavetyAlarmValues of type BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	if listOfAlarmValues == nil {
		panic("listOfAlarmValues of type BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	if modePropertyReference == nil {
		panic("modePropertyReference of type BACnetDeviceObjectPropertyReferenceEnclosed for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterChangeOfLifeSavety must not be nil")
	}
	_result := &_BACnetEventParameterChangeOfLifeSavety{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		ListOfLifeSavetyAlarmValues:  listOfLifeSavetyAlarmValues,
		ListOfAlarmValues:            listOfAlarmValues,
		ModePropertyReference:        modePropertyReference,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterChangeOfLifeSavetyBuilder is a builder for BACnetEventParameterChangeOfLifeSavety
type BACnetEventParameterChangeOfLifeSavetyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithListOfLifeSavetyAlarmValues adds ListOfLifeSavetyAlarmValues (property field)
	WithListOfLifeSavetyAlarmValues(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithListOfLifeSavetyAlarmValuesBuilder adds ListOfLifeSavetyAlarmValues (property field) which is build by the builder
	WithListOfLifeSavetyAlarmValuesBuilder(func(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithListOfAlarmValues adds ListOfAlarmValues (property field)
	WithListOfAlarmValues(BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithListOfAlarmValuesBuilder adds ListOfAlarmValues (property field) which is build by the builder
	WithListOfAlarmValuesBuilder(func(BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithModePropertyReference adds ModePropertyReference (property field)
	WithModePropertyReference(BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithModePropertyReferenceBuilder adds ModePropertyReference (property field) which is build by the builder
	WithModePropertyReferenceBuilder(func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetEventParameterBuilder
	// Build builds the BACnetEventParameterChangeOfLifeSavety or returns an error if something is wrong
	Build() (BACnetEventParameterChangeOfLifeSavety, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterChangeOfLifeSavety
}

// NewBACnetEventParameterChangeOfLifeSavetyBuilder() creates a BACnetEventParameterChangeOfLifeSavetyBuilder
func NewBACnetEventParameterChangeOfLifeSavetyBuilder() BACnetEventParameterChangeOfLifeSavetyBuilder {
	return &_BACnetEventParameterChangeOfLifeSavetyBuilder{_BACnetEventParameterChangeOfLifeSavety: new(_BACnetEventParameterChangeOfLifeSavety)}
}

type _BACnetEventParameterChangeOfLifeSavetyBuilder struct {
	*_BACnetEventParameterChangeOfLifeSavety

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterChangeOfLifeSavetyBuilder) = (*_BACnetEventParameterChangeOfLifeSavetyBuilder)(nil)

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed, closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithListOfLifeSavetyAlarmValues(listOfLifeSavetyAlarmValues).WithListOfAlarmValues(listOfAlarmValues).WithModePropertyReference(modePropertyReference).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithListOfLifeSavetyAlarmValues(listOfLifeSavetyAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.ListOfLifeSavetyAlarmValues = listOfLifeSavetyAlarmValues
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithListOfLifeSavetyAlarmValuesBuilder(builderSupplier func(BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.ListOfLifeSavetyAlarmValues.CreateBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder())
	var err error
	b.ListOfLifeSavetyAlarmValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithListOfAlarmValues(listOfAlarmValues BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.ListOfAlarmValues = listOfAlarmValues
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithListOfAlarmValuesBuilder(builderSupplier func(BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.ListOfAlarmValues.CreateBACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder())
	var err error
	b.ListOfAlarmValues, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithModePropertyReference(modePropertyReference BACnetDeviceObjectPropertyReferenceEnclosed) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.ModePropertyReference = modePropertyReference
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithModePropertyReferenceBuilder(builderSupplier func(BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetDeviceObjectPropertyReferenceEnclosedBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.ModePropertyReference.CreateBACnetDeviceObjectPropertyReferenceEnclosedBuilder())
	var err error
	b.ModePropertyReference, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectPropertyReferenceEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterChangeOfLifeSavetyBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterChangeOfLifeSavetyBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) Build() (BACnetEventParameterChangeOfLifeSavety, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.ListOfLifeSavetyAlarmValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfLifeSavetyAlarmValues' not set"))
	}
	if b.ListOfAlarmValues == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfAlarmValues' not set"))
	}
	if b.ModePropertyReference == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'modePropertyReference' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterChangeOfLifeSavety.deepCopy(), nil
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) MustBuild() BACnetEventParameterChangeOfLifeSavety {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) Done() BACnetEventParameterBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetEventParameterBuilder().(*_BACnetEventParameterBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterChangeOfLifeSavetyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterChangeOfLifeSavetyBuilder().(*_BACnetEventParameterChangeOfLifeSavetyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterChangeOfLifeSavetyBuilder creates a BACnetEventParameterChangeOfLifeSavetyBuilder
func (b *_BACnetEventParameterChangeOfLifeSavety) CreateBACnetEventParameterChangeOfLifeSavetyBuilder() BACnetEventParameterChangeOfLifeSavetyBuilder {
	if b == nil {
		return NewBACnetEventParameterChangeOfLifeSavetyBuilder()
	}
	return &_BACnetEventParameterChangeOfLifeSavetyBuilder{_BACnetEventParameterChangeOfLifeSavety: b.deepCopy()}
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

func (m *_BACnetEventParameterChangeOfLifeSavety) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfLifeSavety) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfLifeSavetyAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetListOfAlarmValues() BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetModePropertyReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfLifeSavety(structType any) BACnetEventParameterChangeOfLifeSavety {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavety); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavety); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavety"
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (listOfLifeSavetyAlarmValues)
	lengthInBits += m.ListOfLifeSavetyAlarmValues.GetLengthInBits(ctx)

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits(ctx)

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterChangeOfLifeSavety) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterChangeOfLifeSavety) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterChangeOfLifeSavety BACnetEventParameterChangeOfLifeSavety, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavety"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfLifeSavety")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	listOfLifeSavetyAlarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](ctx, "listOfLifeSavetyAlarmValues", ReadComplex[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParseWithBufferProducer((uint8)(uint8(1))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfLifeSavetyAlarmValues' field"))
	}
	m.ListOfLifeSavetyAlarmValues = listOfLifeSavetyAlarmValues

	listOfAlarmValues, err := ReadSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](ctx, "listOfAlarmValues", ReadComplex[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfAlarmValues' field"))
	}
	m.ListOfAlarmValues = listOfAlarmValues

	modePropertyReference, err := ReadSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", ReadComplex[BACnetDeviceObjectPropertyReferenceEnclosed](BACnetDeviceObjectPropertyReferenceEnclosedParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'modePropertyReference' field"))
	}
	m.ModePropertyReference = modePropertyReference

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(8))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavety"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfLifeSavety")
	}

	return m, nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfLifeSavety) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavety"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfLifeSavety")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](ctx, "listOfLifeSavetyAlarmValues", m.GetListOfLifeSavetyAlarmValues(), WriteComplex[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfLifeSavetyAlarmValues' field")
		}

		if err := WriteSimpleField[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](ctx, "listOfAlarmValues", m.GetListOfAlarmValues(), WriteComplex[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfAlarmValues' field")
		}

		if err := WriteSimpleField[BACnetDeviceObjectPropertyReferenceEnclosed](ctx, "modePropertyReference", m.GetModePropertyReference(), WriteComplex[BACnetDeviceObjectPropertyReferenceEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'modePropertyReference' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavety"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfLifeSavety")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterChangeOfLifeSavety) IsBACnetEventParameterChangeOfLifeSavety() {}

func (m *_BACnetEventParameterChangeOfLifeSavety) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterChangeOfLifeSavety) deepCopy() *_BACnetEventParameterChangeOfLifeSavety {
	if m == nil {
		return nil
	}
	_BACnetEventParameterChangeOfLifeSavetyCopy := &_BACnetEventParameterChangeOfLifeSavety{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDelay),
		utils.DeepCopy[BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues](m.ListOfLifeSavetyAlarmValues),
		utils.DeepCopy[BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues](m.ListOfAlarmValues),
		utils.DeepCopy[BACnetDeviceObjectPropertyReferenceEnclosed](m.ModePropertyReference),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterChangeOfLifeSavetyCopy
}

func (m *_BACnetEventParameterChangeOfLifeSavety) String() string {
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
