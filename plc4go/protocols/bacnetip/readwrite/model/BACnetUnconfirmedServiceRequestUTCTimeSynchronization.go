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

// BACnetUnconfirmedServiceRequestUTCTimeSynchronization is the corresponding interface of BACnetUnconfirmedServiceRequestUTCTimeSynchronization
type BACnetUnconfirmedServiceRequestUTCTimeSynchronization interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetSynchronizedDate returns SynchronizedDate (property field)
	GetSynchronizedDate() BACnetApplicationTagDate
	// GetSynchronizedTime returns SynchronizedTime (property field)
	GetSynchronizedTime() BACnetApplicationTagTime
	// IsBACnetUnconfirmedServiceRequestUTCTimeSynchronization is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestUTCTimeSynchronization()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	CreateBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
}

// _BACnetUnconfirmedServiceRequestUTCTimeSynchronization is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUTCTimeSynchronization struct {
	BACnetUnconfirmedServiceRequestContract
	SynchronizedDate BACnetApplicationTagDate
	SynchronizedTime BACnetApplicationTagTime
}

var _ BACnetUnconfirmedServiceRequestUTCTimeSynchronization = (*_BACnetUnconfirmedServiceRequestUTCTimeSynchronization)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestUTCTimeSynchronization)(nil)

// NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization factory function for _BACnetUnconfirmedServiceRequestUTCTimeSynchronization
func NewBACnetUnconfirmedServiceRequestUTCTimeSynchronization(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	if synchronizedDate == nil {
		panic("synchronizedDate of type BACnetApplicationTagDate for BACnetUnconfirmedServiceRequestUTCTimeSynchronization must not be nil")
	}
	if synchronizedTime == nil {
		panic("synchronizedTime of type BACnetApplicationTagTime for BACnetUnconfirmedServiceRequestUTCTimeSynchronization must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		SynchronizedDate:                        synchronizedDate,
		SynchronizedTime:                        synchronizedTime,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder is a builder for BACnetUnconfirmedServiceRequestUTCTimeSynchronization
type BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	// WithSynchronizedDate adds SynchronizedDate (property field)
	WithSynchronizedDate(BACnetApplicationTagDate) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	// WithSynchronizedDateBuilder adds SynchronizedDate (property field) which is build by the builder
	WithSynchronizedDateBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	// WithSynchronizedTime adds SynchronizedTime (property field)
	WithSynchronizedTime(BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	// WithSynchronizedTimeBuilder adds SynchronizedTime (property field) which is build by the builder
	WithSynchronizedTimeBuilder(func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetUnconfirmedServiceRequestBuilder
	// Build builds the BACnetUnconfirmedServiceRequestUTCTimeSynchronization or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestUTCTimeSynchronization, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestUTCTimeSynchronization
}

// NewBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder() creates a BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
func NewBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	return &_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder{_BACnetUnconfirmedServiceRequestUTCTimeSynchronization: new(_BACnetUnconfirmedServiceRequestUTCTimeSynchronization)}
}

type _BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder struct {
	*_BACnetUnconfirmedServiceRequestUTCTimeSynchronization

	parentBuilder *_BACnetUnconfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) = (*_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) setParent(contract BACnetUnconfirmedServiceRequestContract) {
	b.BACnetUnconfirmedServiceRequestContract = contract
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) WithMandatoryFields(synchronizedDate BACnetApplicationTagDate, synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	return b.WithSynchronizedDate(synchronizedDate).WithSynchronizedTime(synchronizedTime)
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) WithSynchronizedDate(synchronizedDate BACnetApplicationTagDate) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	b.SynchronizedDate = synchronizedDate
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) WithSynchronizedDateBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	builder := builderSupplier(b.SynchronizedDate.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.SynchronizedDate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) WithSynchronizedTime(synchronizedTime BACnetApplicationTagTime) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	b.SynchronizedTime = synchronizedTime
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) WithSynchronizedTimeBuilder(builderSupplier func(BACnetApplicationTagTimeBuilder) BACnetApplicationTagTimeBuilder) BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	builder := builderSupplier(b.SynchronizedTime.CreateBACnetApplicationTagTimeBuilder())
	var err error
	b.SynchronizedTime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagTimeBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) Build() (BACnetUnconfirmedServiceRequestUTCTimeSynchronization, error) {
	if b.SynchronizedDate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'synchronizedDate' not set"))
	}
	if b.SynchronizedTime == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'synchronizedTime' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestUTCTimeSynchronization.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) MustBuild() BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) Done() BACnetUnconfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetUnconfirmedServiceRequestBuilder().(*_BACnetUnconfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) buildForBACnetUnconfirmedServiceRequest() (BACnetUnconfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder().(*_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder creates a BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder
func (b *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) CreateBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder() BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationBuilder{_BACnetUnconfirmedServiceRequestUTCTimeSynchronization: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedDate() BACnetApplicationTagDate {
	return m.SynchronizedDate
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetSynchronizedTime() BACnetApplicationTagTime {
	return m.SynchronizedTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUTCTimeSynchronization(structType any) BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUTCTimeSynchronization); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUTCTimeSynchronization"
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (synchronizedDate)
	lengthInBits += m.SynchronizedDate.GetLengthInBits(ctx)

	// Simple field (synchronizedTime)
	lengthInBits += m.SynchronizedTime.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestUTCTimeSynchronization BACnetUnconfirmedServiceRequestUTCTimeSynchronization, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	synchronizedDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedDate' field"))
	}
	m.SynchronizedDate = synchronizedDate

	synchronizedTime, err := ReadSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", ReadComplex[BACnetApplicationTagTime](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagTime](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'synchronizedTime' field"))
	}
	m.SynchronizedTime = synchronizedTime

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "synchronizedDate", m.GetSynchronizedDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedDate' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagTime](ctx, "synchronizedTime", m.GetSynchronizedTime(), WriteComplex[BACnetApplicationTagTime](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'synchronizedTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUTCTimeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUTCTimeSynchronization")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) IsBACnetUnconfirmedServiceRequestUTCTimeSynchronization() {
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) deepCopy() *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestUTCTimeSynchronizationCopy := &_BACnetUnconfirmedServiceRequestUTCTimeSynchronization{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.SynchronizedDate),
		utils.DeepCopy[BACnetApplicationTagTime](m.SynchronizedTime),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestUTCTimeSynchronizationCopy
}

func (m *_BACnetUnconfirmedServiceRequestUTCTimeSynchronization) String() string {
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
