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

// AirConditioningDataSetHumidityUpperGuardLimit is the corresponding interface of AirConditioningDataSetHumidityUpperGuardLimit
type AirConditioningDataSetHumidityUpperGuardLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	AirConditioningData
	// GetZoneGroup returns ZoneGroup (property field)
	GetZoneGroup() byte
	// GetZoneList returns ZoneList (property field)
	GetZoneList() HVACZoneList
	// GetLimit returns Limit (property field)
	GetLimit() HVACHumidity
	// GetHvacModeAndFlags returns HvacModeAndFlags (property field)
	GetHvacModeAndFlags() HVACHumidityModeAndFlags
	// IsAirConditioningDataSetHumidityUpperGuardLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsAirConditioningDataSetHumidityUpperGuardLimit()
	// CreateBuilder creates a AirConditioningDataSetHumidityUpperGuardLimitBuilder
	CreateAirConditioningDataSetHumidityUpperGuardLimitBuilder() AirConditioningDataSetHumidityUpperGuardLimitBuilder
}

// _AirConditioningDataSetHumidityUpperGuardLimit is the data-structure of this message
type _AirConditioningDataSetHumidityUpperGuardLimit struct {
	AirConditioningDataContract
	ZoneGroup        byte
	ZoneList         HVACZoneList
	Limit            HVACHumidity
	HvacModeAndFlags HVACHumidityModeAndFlags
}

var _ AirConditioningDataSetHumidityUpperGuardLimit = (*_AirConditioningDataSetHumidityUpperGuardLimit)(nil)
var _ AirConditioningDataRequirements = (*_AirConditioningDataSetHumidityUpperGuardLimit)(nil)

// NewAirConditioningDataSetHumidityUpperGuardLimit factory function for _AirConditioningDataSetHumidityUpperGuardLimit
func NewAirConditioningDataSetHumidityUpperGuardLimit(commandTypeContainer AirConditioningCommandTypeContainer, zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags) *_AirConditioningDataSetHumidityUpperGuardLimit {
	if zoneList == nil {
		panic("zoneList of type HVACZoneList for AirConditioningDataSetHumidityUpperGuardLimit must not be nil")
	}
	if limit == nil {
		panic("limit of type HVACHumidity for AirConditioningDataSetHumidityUpperGuardLimit must not be nil")
	}
	if hvacModeAndFlags == nil {
		panic("hvacModeAndFlags of type HVACHumidityModeAndFlags for AirConditioningDataSetHumidityUpperGuardLimit must not be nil")
	}
	_result := &_AirConditioningDataSetHumidityUpperGuardLimit{
		AirConditioningDataContract: NewAirConditioningData(commandTypeContainer),
		ZoneGroup:                   zoneGroup,
		ZoneList:                    zoneList,
		Limit:                       limit,
		HvacModeAndFlags:            hvacModeAndFlags,
	}
	_result.AirConditioningDataContract.(*_AirConditioningData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// AirConditioningDataSetHumidityUpperGuardLimitBuilder is a builder for AirConditioningDataSetHumidityUpperGuardLimit
type AirConditioningDataSetHumidityUpperGuardLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithZoneGroup adds ZoneGroup (property field)
	WithZoneGroup(byte) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithZoneList adds ZoneList (property field)
	WithZoneList(HVACZoneList) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithZoneListBuilder adds ZoneList (property field) which is build by the builder
	WithZoneListBuilder(func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithLimit adds Limit (property field)
	WithLimit(HVACHumidity) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithLimitBuilder adds Limit (property field) which is build by the builder
	WithLimitBuilder(func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithHvacModeAndFlags adds HvacModeAndFlags (property field)
	WithHvacModeAndFlags(HVACHumidityModeAndFlags) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// WithHvacModeAndFlagsBuilder adds HvacModeAndFlags (property field) which is build by the builder
	WithHvacModeAndFlagsBuilder(func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() AirConditioningDataBuilder
	// Build builds the AirConditioningDataSetHumidityUpperGuardLimit or returns an error if something is wrong
	Build() (AirConditioningDataSetHumidityUpperGuardLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() AirConditioningDataSetHumidityUpperGuardLimit
}

// NewAirConditioningDataSetHumidityUpperGuardLimitBuilder() creates a AirConditioningDataSetHumidityUpperGuardLimitBuilder
func NewAirConditioningDataSetHumidityUpperGuardLimitBuilder() AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	return &_AirConditioningDataSetHumidityUpperGuardLimitBuilder{_AirConditioningDataSetHumidityUpperGuardLimit: new(_AirConditioningDataSetHumidityUpperGuardLimit)}
}

type _AirConditioningDataSetHumidityUpperGuardLimitBuilder struct {
	*_AirConditioningDataSetHumidityUpperGuardLimit

	parentBuilder *_AirConditioningDataBuilder

	err *utils.MultiError
}

var _ (AirConditioningDataSetHumidityUpperGuardLimitBuilder) = (*_AirConditioningDataSetHumidityUpperGuardLimitBuilder)(nil)

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) setParent(contract AirConditioningDataContract) {
	b.AirConditioningDataContract = contract
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithMandatoryFields(zoneGroup byte, zoneList HVACZoneList, limit HVACHumidity, hvacModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	return b.WithZoneGroup(zoneGroup).WithZoneList(zoneList).WithLimit(limit).WithHvacModeAndFlags(hvacModeAndFlags)
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithZoneGroup(zoneGroup byte) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	b.ZoneGroup = zoneGroup
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithZoneList(zoneList HVACZoneList) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	b.ZoneList = zoneList
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithZoneListBuilder(builderSupplier func(HVACZoneListBuilder) HVACZoneListBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	builder := builderSupplier(b.ZoneList.CreateHVACZoneListBuilder())
	var err error
	b.ZoneList, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACZoneListBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithLimit(limit HVACHumidity) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	b.Limit = limit
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithLimitBuilder(builderSupplier func(HVACHumidityBuilder) HVACHumidityBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	builder := builderSupplier(b.Limit.CreateHVACHumidityBuilder())
	var err error
	b.Limit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithHvacModeAndFlags(hvacModeAndFlags HVACHumidityModeAndFlags) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	b.HvacModeAndFlags = hvacModeAndFlags
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) WithHvacModeAndFlagsBuilder(builderSupplier func(HVACHumidityModeAndFlagsBuilder) HVACHumidityModeAndFlagsBuilder) AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	builder := builderSupplier(b.HvacModeAndFlags.CreateHVACHumidityModeAndFlagsBuilder())
	var err error
	b.HvacModeAndFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "HVACHumidityModeAndFlagsBuilder failed"))
	}
	return b
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) Build() (AirConditioningDataSetHumidityUpperGuardLimit, error) {
	if b.ZoneList == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'zoneList' not set"))
	}
	if b.Limit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'limit' not set"))
	}
	if b.HvacModeAndFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'hvacModeAndFlags' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._AirConditioningDataSetHumidityUpperGuardLimit.deepCopy(), nil
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) MustBuild() AirConditioningDataSetHumidityUpperGuardLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) Done() AirConditioningDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewAirConditioningDataBuilder().(*_AirConditioningDataBuilder)
	}
	return b.parentBuilder
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) buildForAirConditioningData() (AirConditioningData, error) {
	return b.Build()
}

func (b *_AirConditioningDataSetHumidityUpperGuardLimitBuilder) DeepCopy() any {
	_copy := b.CreateAirConditioningDataSetHumidityUpperGuardLimitBuilder().(*_AirConditioningDataSetHumidityUpperGuardLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateAirConditioningDataSetHumidityUpperGuardLimitBuilder creates a AirConditioningDataSetHumidityUpperGuardLimitBuilder
func (b *_AirConditioningDataSetHumidityUpperGuardLimit) CreateAirConditioningDataSetHumidityUpperGuardLimitBuilder() AirConditioningDataSetHumidityUpperGuardLimitBuilder {
	if b == nil {
		return NewAirConditioningDataSetHumidityUpperGuardLimitBuilder()
	}
	return &_AirConditioningDataSetHumidityUpperGuardLimitBuilder{_AirConditioningDataSetHumidityUpperGuardLimit: b.deepCopy()}
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

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetParent() AirConditioningDataContract {
	return m.AirConditioningDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneGroup() byte {
	return m.ZoneGroup
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetZoneList() HVACZoneList {
	return m.ZoneList
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLimit() HVACHumidity {
	return m.Limit
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetHvacModeAndFlags() HVACHumidityModeAndFlags {
	return m.HvacModeAndFlags
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastAirConditioningDataSetHumidityUpperGuardLimit(structType any) AirConditioningDataSetHumidityUpperGuardLimit {
	if casted, ok := structType.(AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return casted
	}
	if casted, ok := structType.(*AirConditioningDataSetHumidityUpperGuardLimit); ok {
		return *casted
	}
	return nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetTypeName() string {
	return "AirConditioningDataSetHumidityUpperGuardLimit"
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.AirConditioningDataContract.(*_AirConditioningData).getLengthInBits(ctx))

	// Simple field (zoneGroup)
	lengthInBits += 8

	// Simple field (zoneList)
	lengthInBits += m.ZoneList.GetLengthInBits(ctx)

	// Simple field (limit)
	lengthInBits += m.Limit.GetLengthInBits(ctx)

	// Simple field (hvacModeAndFlags)
	lengthInBits += m.HvacModeAndFlags.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_AirConditioningData) (__airConditioningDataSetHumidityUpperGuardLimit AirConditioningDataSetHumidityUpperGuardLimit, err error) {
	m.AirConditioningDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AirConditioningDataSetHumidityUpperGuardLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AirConditioningDataSetHumidityUpperGuardLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zoneGroup, err := ReadSimpleField(ctx, "zoneGroup", ReadByte(readBuffer, 8))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneGroup' field"))
	}
	m.ZoneGroup = zoneGroup

	zoneList, err := ReadSimpleField[HVACZoneList](ctx, "zoneList", ReadComplex[HVACZoneList](HVACZoneListParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zoneList' field"))
	}
	m.ZoneList = zoneList

	limit, err := ReadSimpleField[HVACHumidity](ctx, "limit", ReadComplex[HVACHumidity](HVACHumidityParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'limit' field"))
	}
	m.Limit = limit

	hvacModeAndFlags, err := ReadSimpleField[HVACHumidityModeAndFlags](ctx, "hvacModeAndFlags", ReadComplex[HVACHumidityModeAndFlags](HVACHumidityModeAndFlagsParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'hvacModeAndFlags' field"))
	}
	m.HvacModeAndFlags = hvacModeAndFlags

	if closeErr := readBuffer.CloseContext("AirConditioningDataSetHumidityUpperGuardLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AirConditioningDataSetHumidityUpperGuardLimit")
	}

	return m, nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AirConditioningDataSetHumidityUpperGuardLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AirConditioningDataSetHumidityUpperGuardLimit")
		}

		if err := WriteSimpleField[byte](ctx, "zoneGroup", m.GetZoneGroup(), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneGroup' field")
		}

		if err := WriteSimpleField[HVACZoneList](ctx, "zoneList", m.GetZoneList(), WriteComplex[HVACZoneList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'zoneList' field")
		}

		if err := WriteSimpleField[HVACHumidity](ctx, "limit", m.GetLimit(), WriteComplex[HVACHumidity](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'limit' field")
		}

		if err := WriteSimpleField[HVACHumidityModeAndFlags](ctx, "hvacModeAndFlags", m.GetHvacModeAndFlags(), WriteComplex[HVACHumidityModeAndFlags](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'hvacModeAndFlags' field")
		}

		if popErr := writeBuffer.PopContext("AirConditioningDataSetHumidityUpperGuardLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AirConditioningDataSetHumidityUpperGuardLimit")
		}
		return nil
	}
	return m.AirConditioningDataContract.(*_AirConditioningData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) IsAirConditioningDataSetHumidityUpperGuardLimit() {
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) deepCopy() *_AirConditioningDataSetHumidityUpperGuardLimit {
	if m == nil {
		return nil
	}
	_AirConditioningDataSetHumidityUpperGuardLimitCopy := &_AirConditioningDataSetHumidityUpperGuardLimit{
		m.AirConditioningDataContract.(*_AirConditioningData).deepCopy(),
		m.ZoneGroup,
		utils.DeepCopy[HVACZoneList](m.ZoneList),
		utils.DeepCopy[HVACHumidity](m.Limit),
		utils.DeepCopy[HVACHumidityModeAndFlags](m.HvacModeAndFlags),
	}
	m.AirConditioningDataContract.(*_AirConditioningData)._SubType = m
	return _AirConditioningDataSetHumidityUpperGuardLimitCopy
}

func (m *_AirConditioningDataSetHumidityUpperGuardLimit) String() string {
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
