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

// BACnetUnconfirmedServiceRequestIAm is the corresponding interface of BACnetUnconfirmedServiceRequestIAm
type BACnetUnconfirmedServiceRequestIAm interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetUnconfirmedServiceRequest
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier
	// GetMaximumApduLengthAcceptedLength returns MaximumApduLengthAcceptedLength (property field)
	GetMaximumApduLengthAcceptedLength() BACnetApplicationTagUnsignedInteger
	// GetSegmentationSupported returns SegmentationSupported (property field)
	GetSegmentationSupported() BACnetSegmentationTagged
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// IsBACnetUnconfirmedServiceRequestIAm is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetUnconfirmedServiceRequestIAm()
	// CreateBuilder creates a BACnetUnconfirmedServiceRequestIAmBuilder
	CreateBACnetUnconfirmedServiceRequestIAmBuilder() BACnetUnconfirmedServiceRequestIAmBuilder
}

// _BACnetUnconfirmedServiceRequestIAm is the data-structure of this message
type _BACnetUnconfirmedServiceRequestIAm struct {
	BACnetUnconfirmedServiceRequestContract
	DeviceIdentifier                BACnetApplicationTagObjectIdentifier
	MaximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger
	SegmentationSupported           BACnetSegmentationTagged
	VendorId                        BACnetVendorIdTagged
}

var _ BACnetUnconfirmedServiceRequestIAm = (*_BACnetUnconfirmedServiceRequestIAm)(nil)
var _ BACnetUnconfirmedServiceRequestRequirements = (*_BACnetUnconfirmedServiceRequestIAm)(nil)

// NewBACnetUnconfirmedServiceRequestIAm factory function for _BACnetUnconfirmedServiceRequestIAm
func NewBACnetUnconfirmedServiceRequestIAm(deviceIdentifier BACnetApplicationTagObjectIdentifier, maximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger, segmentationSupported BACnetSegmentationTagged, vendorId BACnetVendorIdTagged, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestIAm {
	if deviceIdentifier == nil {
		panic("deviceIdentifier of type BACnetApplicationTagObjectIdentifier for BACnetUnconfirmedServiceRequestIAm must not be nil")
	}
	if maximumApduLengthAcceptedLength == nil {
		panic("maximumApduLengthAcceptedLength of type BACnetApplicationTagUnsignedInteger for BACnetUnconfirmedServiceRequestIAm must not be nil")
	}
	if segmentationSupported == nil {
		panic("segmentationSupported of type BACnetSegmentationTagged for BACnetUnconfirmedServiceRequestIAm must not be nil")
	}
	if vendorId == nil {
		panic("vendorId of type BACnetVendorIdTagged for BACnetUnconfirmedServiceRequestIAm must not be nil")
	}
	_result := &_BACnetUnconfirmedServiceRequestIAm{
		BACnetUnconfirmedServiceRequestContract: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
		DeviceIdentifier:                        deviceIdentifier,
		MaximumApduLengthAcceptedLength:         maximumApduLengthAcceptedLength,
		SegmentationSupported:                   segmentationSupported,
		VendorId:                                vendorId,
	}
	_result.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetUnconfirmedServiceRequestIAmBuilder is a builder for BACnetUnconfirmedServiceRequestIAm
type BACnetUnconfirmedServiceRequestIAmBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(deviceIdentifier BACnetApplicationTagObjectIdentifier, maximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger, segmentationSupported BACnetSegmentationTagged, vendorId BACnetVendorIdTagged) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithDeviceIdentifier adds DeviceIdentifier (property field)
	WithDeviceIdentifier(BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithDeviceIdentifierBuilder adds DeviceIdentifier (property field) which is build by the builder
	WithDeviceIdentifierBuilder(func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithMaximumApduLengthAcceptedLength adds MaximumApduLengthAcceptedLength (property field)
	WithMaximumApduLengthAcceptedLength(BACnetApplicationTagUnsignedInteger) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithMaximumApduLengthAcceptedLengthBuilder adds MaximumApduLengthAcceptedLength (property field) which is build by the builder
	WithMaximumApduLengthAcceptedLengthBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithSegmentationSupported adds SegmentationSupported (property field)
	WithSegmentationSupported(BACnetSegmentationTagged) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithSegmentationSupportedBuilder adds SegmentationSupported (property field) which is build by the builder
	WithSegmentationSupportedBuilder(func(BACnetSegmentationTaggedBuilder) BACnetSegmentationTaggedBuilder) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithVendorId adds VendorId (property field)
	WithVendorId(BACnetVendorIdTagged) BACnetUnconfirmedServiceRequestIAmBuilder
	// WithVendorIdBuilder adds VendorId (property field) which is build by the builder
	WithVendorIdBuilder(func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetUnconfirmedServiceRequestIAmBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetUnconfirmedServiceRequestBuilder
	// Build builds the BACnetUnconfirmedServiceRequestIAm or returns an error if something is wrong
	Build() (BACnetUnconfirmedServiceRequestIAm, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetUnconfirmedServiceRequestIAm
}

// NewBACnetUnconfirmedServiceRequestIAmBuilder() creates a BACnetUnconfirmedServiceRequestIAmBuilder
func NewBACnetUnconfirmedServiceRequestIAmBuilder() BACnetUnconfirmedServiceRequestIAmBuilder {
	return &_BACnetUnconfirmedServiceRequestIAmBuilder{_BACnetUnconfirmedServiceRequestIAm: new(_BACnetUnconfirmedServiceRequestIAm)}
}

type _BACnetUnconfirmedServiceRequestIAmBuilder struct {
	*_BACnetUnconfirmedServiceRequestIAm

	parentBuilder *_BACnetUnconfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetUnconfirmedServiceRequestIAmBuilder) = (*_BACnetUnconfirmedServiceRequestIAmBuilder)(nil)

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) setParent(contract BACnetUnconfirmedServiceRequestContract) {
	b.BACnetUnconfirmedServiceRequestContract = contract
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithMandatoryFields(deviceIdentifier BACnetApplicationTagObjectIdentifier, maximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger, segmentationSupported BACnetSegmentationTagged, vendorId BACnetVendorIdTagged) BACnetUnconfirmedServiceRequestIAmBuilder {
	return b.WithDeviceIdentifier(deviceIdentifier).WithMaximumApduLengthAcceptedLength(maximumApduLengthAcceptedLength).WithSegmentationSupported(segmentationSupported).WithVendorId(vendorId)
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithDeviceIdentifier(deviceIdentifier BACnetApplicationTagObjectIdentifier) BACnetUnconfirmedServiceRequestIAmBuilder {
	b.DeviceIdentifier = deviceIdentifier
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithDeviceIdentifierBuilder(builderSupplier func(BACnetApplicationTagObjectIdentifierBuilder) BACnetApplicationTagObjectIdentifierBuilder) BACnetUnconfirmedServiceRequestIAmBuilder {
	builder := builderSupplier(b.DeviceIdentifier.CreateBACnetApplicationTagObjectIdentifierBuilder())
	var err error
	b.DeviceIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithMaximumApduLengthAcceptedLength(maximumApduLengthAcceptedLength BACnetApplicationTagUnsignedInteger) BACnetUnconfirmedServiceRequestIAmBuilder {
	b.MaximumApduLengthAcceptedLength = maximumApduLengthAcceptedLength
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithMaximumApduLengthAcceptedLengthBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetUnconfirmedServiceRequestIAmBuilder {
	builder := builderSupplier(b.MaximumApduLengthAcceptedLength.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaximumApduLengthAcceptedLength, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithSegmentationSupported(segmentationSupported BACnetSegmentationTagged) BACnetUnconfirmedServiceRequestIAmBuilder {
	b.SegmentationSupported = segmentationSupported
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithSegmentationSupportedBuilder(builderSupplier func(BACnetSegmentationTaggedBuilder) BACnetSegmentationTaggedBuilder) BACnetUnconfirmedServiceRequestIAmBuilder {
	builder := builderSupplier(b.SegmentationSupported.CreateBACnetSegmentationTaggedBuilder())
	var err error
	b.SegmentationSupported, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetSegmentationTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithVendorId(vendorId BACnetVendorIdTagged) BACnetUnconfirmedServiceRequestIAmBuilder {
	b.VendorId = vendorId
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) WithVendorIdBuilder(builderSupplier func(BACnetVendorIdTaggedBuilder) BACnetVendorIdTaggedBuilder) BACnetUnconfirmedServiceRequestIAmBuilder {
	builder := builderSupplier(b.VendorId.CreateBACnetVendorIdTaggedBuilder())
	var err error
	b.VendorId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetVendorIdTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) Build() (BACnetUnconfirmedServiceRequestIAm, error) {
	if b.DeviceIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deviceIdentifier' not set"))
	}
	if b.MaximumApduLengthAcceptedLength == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maximumApduLengthAcceptedLength' not set"))
	}
	if b.SegmentationSupported == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'segmentationSupported' not set"))
	}
	if b.VendorId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'vendorId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetUnconfirmedServiceRequestIAm.deepCopy(), nil
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) MustBuild() BACnetUnconfirmedServiceRequestIAm {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) Done() BACnetUnconfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetUnconfirmedServiceRequestBuilder().(*_BACnetUnconfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) buildForBACnetUnconfirmedServiceRequest() (BACnetUnconfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetUnconfirmedServiceRequestIAmBuilder) DeepCopy() any {
	_copy := b.CreateBACnetUnconfirmedServiceRequestIAmBuilder().(*_BACnetUnconfirmedServiceRequestIAmBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetUnconfirmedServiceRequestIAmBuilder creates a BACnetUnconfirmedServiceRequestIAmBuilder
func (b *_BACnetUnconfirmedServiceRequestIAm) CreateBACnetUnconfirmedServiceRequestIAmBuilder() BACnetUnconfirmedServiceRequestIAmBuilder {
	if b == nil {
		return NewBACnetUnconfirmedServiceRequestIAmBuilder()
	}
	return &_BACnetUnconfirmedServiceRequestIAmBuilder{_BACnetUnconfirmedServiceRequestIAm: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_I_AM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) GetParent() BACnetUnconfirmedServiceRequestContract {
	return m.BACnetUnconfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestIAm) GetDeviceIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetMaximumApduLengthAcceptedLength() BACnetApplicationTagUnsignedInteger {
	return m.MaximumApduLengthAcceptedLength
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetSegmentationSupported() BACnetSegmentationTagged {
	return m.SegmentationSupported
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestIAm(structType any) BACnetUnconfirmedServiceRequestIAm {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestIAm); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestIAm); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestIAm"
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (deviceIdentifier)
	lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (maximumApduLengthAcceptedLength)
	lengthInBits += m.MaximumApduLengthAcceptedLength.GetLengthInBits(ctx)

	// Simple field (segmentationSupported)
	lengthInBits += m.SegmentationSupported.GetLengthInBits(ctx)

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestIAm) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetUnconfirmedServiceRequestIAm) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetUnconfirmedServiceRequest, serviceRequestLength uint16) (__bACnetUnconfirmedServiceRequestIAm BACnetUnconfirmedServiceRequestIAm, err error) {
	m.BACnetUnconfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestIAm"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestIAm")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	deviceIdentifier, err := ReadSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", ReadComplex[BACnetApplicationTagObjectIdentifier](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagObjectIdentifier](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceIdentifier' field"))
	}
	m.DeviceIdentifier = deviceIdentifier

	maximumApduLengthAcceptedLength, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maximumApduLengthAcceptedLength", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maximumApduLengthAcceptedLength' field"))
	}
	m.MaximumApduLengthAcceptedLength = maximumApduLengthAcceptedLength

	segmentationSupported, err := ReadSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", ReadComplex[BACnetSegmentationTagged](BACnetSegmentationTaggedParseWithBufferProducer((uint8)(uint8(9)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'segmentationSupported' field"))
	}
	m.SegmentationSupported = segmentationSupported

	vendorId, err := ReadSimpleField[BACnetVendorIdTagged](ctx, "vendorId", ReadComplex[BACnetVendorIdTagged](BACnetVendorIdTaggedParseWithBufferProducer((uint8)(uint8(2)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestIAm"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestIAm")
	}

	return m, nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestIAm) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestIAm"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestIAm")
		}

		if err := WriteSimpleField[BACnetApplicationTagObjectIdentifier](ctx, "deviceIdentifier", m.GetDeviceIdentifier(), WriteComplex[BACnetApplicationTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maximumApduLengthAcceptedLength", m.GetMaximumApduLengthAcceptedLength(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maximumApduLengthAcceptedLength' field")
		}

		if err := WriteSimpleField[BACnetSegmentationTagged](ctx, "segmentationSupported", m.GetSegmentationSupported(), WriteComplex[BACnetSegmentationTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'segmentationSupported' field")
		}

		if err := WriteSimpleField[BACnetVendorIdTagged](ctx, "vendorId", m.GetVendorId(), WriteComplex[BACnetVendorIdTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestIAm"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestIAm")
		}
		return nil
	}
	return m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestIAm) IsBACnetUnconfirmedServiceRequestIAm() {}

func (m *_BACnetUnconfirmedServiceRequestIAm) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetUnconfirmedServiceRequestIAm) deepCopy() *_BACnetUnconfirmedServiceRequestIAm {
	if m == nil {
		return nil
	}
	_BACnetUnconfirmedServiceRequestIAmCopy := &_BACnetUnconfirmedServiceRequestIAm{
		m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagObjectIdentifier](m.DeviceIdentifier),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MaximumApduLengthAcceptedLength),
		utils.DeepCopy[BACnetSegmentationTagged](m.SegmentationSupported),
		utils.DeepCopy[BACnetVendorIdTagged](m.VendorId),
	}
	m.BACnetUnconfirmedServiceRequestContract.(*_BACnetUnconfirmedServiceRequest)._SubType = m
	return _BACnetUnconfirmedServiceRequestIAmCopy
}

func (m *_BACnetUnconfirmedServiceRequestIAm) String() string {
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
