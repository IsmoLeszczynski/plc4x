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
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/codegen"
	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CipIdentity_ZEROES1 uint32 = 0x00000000
const CipIdentity_ZEROES2 uint32 = 0x00000000

// CipIdentity is the corresponding interface of CipIdentity
type CipIdentity interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CommandSpecificDataItem
	// GetEncapsulationProtocolVersion returns EncapsulationProtocolVersion (property field)
	GetEncapsulationProtocolVersion() uint16
	// GetSocketAddressFamily returns SocketAddressFamily (property field)
	GetSocketAddressFamily() uint16
	// GetSocketAddressPort returns SocketAddressPort (property field)
	GetSocketAddressPort() uint16
	// GetSocketAddressAddress returns SocketAddressAddress (property field)
	GetSocketAddressAddress() []uint8
	// GetVendorId returns VendorId (property field)
	GetVendorId() uint16
	// GetDeviceType returns DeviceType (property field)
	GetDeviceType() uint16
	// GetProductCode returns ProductCode (property field)
	GetProductCode() uint16
	// GetRevisionMajor returns RevisionMajor (property field)
	GetRevisionMajor() uint8
	// GetRevisionMinor returns RevisionMinor (property field)
	GetRevisionMinor() uint8
	// GetStatus returns Status (property field)
	GetStatus() uint16
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() uint32
	// GetProductName returns ProductName (property field)
	GetProductName() string
	// GetState returns State (property field)
	GetState() uint8
	// IsCipIdentity is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCipIdentity()
	// CreateBuilder creates a CipIdentityBuilder
	CreateCipIdentityBuilder() CipIdentityBuilder
}

// _CipIdentity is the data-structure of this message
type _CipIdentity struct {
	CommandSpecificDataItemContract
	EncapsulationProtocolVersion uint16
	SocketAddressFamily          uint16
	SocketAddressPort            uint16
	SocketAddressAddress         []uint8
	VendorId                     uint16
	DeviceType                   uint16
	ProductCode                  uint16
	RevisionMajor                uint8
	RevisionMinor                uint8
	Status                       uint16
	SerialNumber                 uint32
	ProductName                  string
	State                        uint8
}

var _ CipIdentity = (*_CipIdentity)(nil)
var _ CommandSpecificDataItemRequirements = (*_CipIdentity)(nil)

// NewCipIdentity factory function for _CipIdentity
func NewCipIdentity(encapsulationProtocolVersion uint16, socketAddressFamily uint16, socketAddressPort uint16, socketAddressAddress []uint8, vendorId uint16, deviceType uint16, productCode uint16, revisionMajor uint8, revisionMinor uint8, status uint16, serialNumber uint32, productName string, state uint8) *_CipIdentity {
	_result := &_CipIdentity{
		CommandSpecificDataItemContract: NewCommandSpecificDataItem(),
		EncapsulationProtocolVersion:    encapsulationProtocolVersion,
		SocketAddressFamily:             socketAddressFamily,
		SocketAddressPort:               socketAddressPort,
		SocketAddressAddress:            socketAddressAddress,
		VendorId:                        vendorId,
		DeviceType:                      deviceType,
		ProductCode:                     productCode,
		RevisionMajor:                   revisionMajor,
		RevisionMinor:                   revisionMinor,
		Status:                          status,
		SerialNumber:                    serialNumber,
		ProductName:                     productName,
		State:                           state,
	}
	_result.CommandSpecificDataItemContract.(*_CommandSpecificDataItem)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CipIdentityBuilder is a builder for CipIdentity
type CipIdentityBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(encapsulationProtocolVersion uint16, socketAddressFamily uint16, socketAddressPort uint16, socketAddressAddress []uint8, vendorId uint16, deviceType uint16, productCode uint16, revisionMajor uint8, revisionMinor uint8, status uint16, serialNumber uint32, productName string, state uint8) CipIdentityBuilder
	// WithEncapsulationProtocolVersion adds EncapsulationProtocolVersion (property field)
	WithEncapsulationProtocolVersion(uint16) CipIdentityBuilder
	// WithSocketAddressFamily adds SocketAddressFamily (property field)
	WithSocketAddressFamily(uint16) CipIdentityBuilder
	// WithSocketAddressPort adds SocketAddressPort (property field)
	WithSocketAddressPort(uint16) CipIdentityBuilder
	// WithSocketAddressAddress adds SocketAddressAddress (property field)
	WithSocketAddressAddress(...uint8) CipIdentityBuilder
	// WithVendorId adds VendorId (property field)
	WithVendorId(uint16) CipIdentityBuilder
	// WithDeviceType adds DeviceType (property field)
	WithDeviceType(uint16) CipIdentityBuilder
	// WithProductCode adds ProductCode (property field)
	WithProductCode(uint16) CipIdentityBuilder
	// WithRevisionMajor adds RevisionMajor (property field)
	WithRevisionMajor(uint8) CipIdentityBuilder
	// WithRevisionMinor adds RevisionMinor (property field)
	WithRevisionMinor(uint8) CipIdentityBuilder
	// WithStatus adds Status (property field)
	WithStatus(uint16) CipIdentityBuilder
	// WithSerialNumber adds SerialNumber (property field)
	WithSerialNumber(uint32) CipIdentityBuilder
	// WithProductName adds ProductName (property field)
	WithProductName(string) CipIdentityBuilder
	// WithState adds State (property field)
	WithState(uint8) CipIdentityBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() CommandSpecificDataItemBuilder
	// Build builds the CipIdentity or returns an error if something is wrong
	Build() (CipIdentity, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CipIdentity
}

// NewCipIdentityBuilder() creates a CipIdentityBuilder
func NewCipIdentityBuilder() CipIdentityBuilder {
	return &_CipIdentityBuilder{_CipIdentity: new(_CipIdentity)}
}

type _CipIdentityBuilder struct {
	*_CipIdentity

	parentBuilder *_CommandSpecificDataItemBuilder

	err *utils.MultiError
}

var _ (CipIdentityBuilder) = (*_CipIdentityBuilder)(nil)

func (b *_CipIdentityBuilder) setParent(contract CommandSpecificDataItemContract) {
	b.CommandSpecificDataItemContract = contract
}

func (b *_CipIdentityBuilder) WithMandatoryFields(encapsulationProtocolVersion uint16, socketAddressFamily uint16, socketAddressPort uint16, socketAddressAddress []uint8, vendorId uint16, deviceType uint16, productCode uint16, revisionMajor uint8, revisionMinor uint8, status uint16, serialNumber uint32, productName string, state uint8) CipIdentityBuilder {
	return b.WithEncapsulationProtocolVersion(encapsulationProtocolVersion).WithSocketAddressFamily(socketAddressFamily).WithSocketAddressPort(socketAddressPort).WithSocketAddressAddress(socketAddressAddress...).WithVendorId(vendorId).WithDeviceType(deviceType).WithProductCode(productCode).WithRevisionMajor(revisionMajor).WithRevisionMinor(revisionMinor).WithStatus(status).WithSerialNumber(serialNumber).WithProductName(productName).WithState(state)
}

func (b *_CipIdentityBuilder) WithEncapsulationProtocolVersion(encapsulationProtocolVersion uint16) CipIdentityBuilder {
	b.EncapsulationProtocolVersion = encapsulationProtocolVersion
	return b
}

func (b *_CipIdentityBuilder) WithSocketAddressFamily(socketAddressFamily uint16) CipIdentityBuilder {
	b.SocketAddressFamily = socketAddressFamily
	return b
}

func (b *_CipIdentityBuilder) WithSocketAddressPort(socketAddressPort uint16) CipIdentityBuilder {
	b.SocketAddressPort = socketAddressPort
	return b
}

func (b *_CipIdentityBuilder) WithSocketAddressAddress(socketAddressAddress ...uint8) CipIdentityBuilder {
	b.SocketAddressAddress = socketAddressAddress
	return b
}

func (b *_CipIdentityBuilder) WithVendorId(vendorId uint16) CipIdentityBuilder {
	b.VendorId = vendorId
	return b
}

func (b *_CipIdentityBuilder) WithDeviceType(deviceType uint16) CipIdentityBuilder {
	b.DeviceType = deviceType
	return b
}

func (b *_CipIdentityBuilder) WithProductCode(productCode uint16) CipIdentityBuilder {
	b.ProductCode = productCode
	return b
}

func (b *_CipIdentityBuilder) WithRevisionMajor(revisionMajor uint8) CipIdentityBuilder {
	b.RevisionMajor = revisionMajor
	return b
}

func (b *_CipIdentityBuilder) WithRevisionMinor(revisionMinor uint8) CipIdentityBuilder {
	b.RevisionMinor = revisionMinor
	return b
}

func (b *_CipIdentityBuilder) WithStatus(status uint16) CipIdentityBuilder {
	b.Status = status
	return b
}

func (b *_CipIdentityBuilder) WithSerialNumber(serialNumber uint32) CipIdentityBuilder {
	b.SerialNumber = serialNumber
	return b
}

func (b *_CipIdentityBuilder) WithProductName(productName string) CipIdentityBuilder {
	b.ProductName = productName
	return b
}

func (b *_CipIdentityBuilder) WithState(state uint8) CipIdentityBuilder {
	b.State = state
	return b
}

func (b *_CipIdentityBuilder) Build() (CipIdentity, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CipIdentity.deepCopy(), nil
}

func (b *_CipIdentityBuilder) MustBuild() CipIdentity {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_CipIdentityBuilder) Done() CommandSpecificDataItemBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewCommandSpecificDataItemBuilder().(*_CommandSpecificDataItemBuilder)
	}
	return b.parentBuilder
}

func (b *_CipIdentityBuilder) buildForCommandSpecificDataItem() (CommandSpecificDataItem, error) {
	return b.Build()
}

func (b *_CipIdentityBuilder) DeepCopy() any {
	_copy := b.CreateCipIdentityBuilder().(*_CipIdentityBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCipIdentityBuilder creates a CipIdentityBuilder
func (b *_CipIdentity) CreateCipIdentityBuilder() CipIdentityBuilder {
	if b == nil {
		return NewCipIdentityBuilder()
	}
	return &_CipIdentityBuilder{_CipIdentity: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipIdentity) GetItemType() uint16 {
	return 0x000C
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipIdentity) GetParent() CommandSpecificDataItemContract {
	return m.CommandSpecificDataItemContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipIdentity) GetEncapsulationProtocolVersion() uint16 {
	return m.EncapsulationProtocolVersion
}

func (m *_CipIdentity) GetSocketAddressFamily() uint16 {
	return m.SocketAddressFamily
}

func (m *_CipIdentity) GetSocketAddressPort() uint16 {
	return m.SocketAddressPort
}

func (m *_CipIdentity) GetSocketAddressAddress() []uint8 {
	return m.SocketAddressAddress
}

func (m *_CipIdentity) GetVendorId() uint16 {
	return m.VendorId
}

func (m *_CipIdentity) GetDeviceType() uint16 {
	return m.DeviceType
}

func (m *_CipIdentity) GetProductCode() uint16 {
	return m.ProductCode
}

func (m *_CipIdentity) GetRevisionMajor() uint8 {
	return m.RevisionMajor
}

func (m *_CipIdentity) GetRevisionMinor() uint8 {
	return m.RevisionMinor
}

func (m *_CipIdentity) GetStatus() uint16 {
	return m.Status
}

func (m *_CipIdentity) GetSerialNumber() uint32 {
	return m.SerialNumber
}

func (m *_CipIdentity) GetProductName() string {
	return m.ProductName
}

func (m *_CipIdentity) GetState() uint8 {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CipIdentity) GetZeroes1() uint32 {
	return CipIdentity_ZEROES1
}

func (m *_CipIdentity) GetZeroes2() uint32 {
	return CipIdentity_ZEROES2
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCipIdentity(structType any) CipIdentity {
	if casted, ok := structType.(CipIdentity); ok {
		return casted
	}
	if casted, ok := structType.(*CipIdentity); ok {
		return *casted
	}
	return nil
}

func (m *_CipIdentity) GetTypeName() string {
	return "CipIdentity"
}

func (m *_CipIdentity) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem).getLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 16

	// Simple field (encapsulationProtocolVersion)
	lengthInBits += 16

	// Simple field (socketAddressFamily)
	lengthInBits += 16

	// Simple field (socketAddressPort)
	lengthInBits += 16

	// Array field
	if len(m.SocketAddressAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.SocketAddressAddress))
	}

	// Const Field (zeroes1)
	lengthInBits += 32

	// Const Field (zeroes2)
	lengthInBits += 32

	// Simple field (vendorId)
	lengthInBits += 16

	// Simple field (deviceType)
	lengthInBits += 16

	// Simple field (productCode)
	lengthInBits += 16

	// Simple field (revisionMajor)
	lengthInBits += 8

	// Simple field (revisionMinor)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 16

	// Simple field (serialNumber)
	lengthInBits += 32

	// Implicit Field (productNameLength)
	lengthInBits += 8

	// Simple field (productName)
	lengthInBits += uint16(int32(uint8(len(m.GetProductName()))) * int32(int32(8)))

	// Simple field (state)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CipIdentity) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CipIdentity) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CommandSpecificDataItem) (__cipIdentity CipIdentity, err error) {
	m.CommandSpecificDataItemContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipIdentity"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipIdentity")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	itemLength, err := ReadImplicitField[uint16](ctx, "itemLength", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemLength' field"))
	}
	_ = itemLength

	encapsulationProtocolVersion, err := ReadSimpleField(ctx, "encapsulationProtocolVersion", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'encapsulationProtocolVersion' field"))
	}
	m.EncapsulationProtocolVersion = encapsulationProtocolVersion

	socketAddressFamily, err := ReadSimpleField(ctx, "socketAddressFamily", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressFamily' field"))
	}
	m.SocketAddressFamily = socketAddressFamily

	socketAddressPort, err := ReadSimpleField(ctx, "socketAddressPort", ReadUnsignedShort(readBuffer, uint8(16)), codegen.WithByteOrder(binary.BigEndian))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressPort' field"))
	}
	m.SocketAddressPort = socketAddressPort

	socketAddressAddress, err := ReadCountArrayField[uint8](ctx, "socketAddressAddress", ReadUnsignedByte(readBuffer, uint8(8)), uint64(int32(4)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'socketAddressAddress' field"))
	}
	m.SocketAddressAddress = socketAddressAddress

	zeroes1, err := ReadConstField[uint32](ctx, "zeroes1", ReadUnsignedInt(readBuffer, uint8(32)), CipIdentity_ZEROES1)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zeroes1' field"))
	}
	_ = zeroes1

	zeroes2, err := ReadConstField[uint32](ctx, "zeroes2", ReadUnsignedInt(readBuffer, uint8(32)), CipIdentity_ZEROES2)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zeroes2' field"))
	}
	_ = zeroes2

	vendorId, err := ReadSimpleField(ctx, "vendorId", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'vendorId' field"))
	}
	m.VendorId = vendorId

	deviceType, err := ReadSimpleField(ctx, "deviceType", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deviceType' field"))
	}
	m.DeviceType = deviceType

	productCode, err := ReadSimpleField(ctx, "productCode", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productCode' field"))
	}
	m.ProductCode = productCode

	revisionMajor, err := ReadSimpleField(ctx, "revisionMajor", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisionMajor' field"))
	}
	m.RevisionMajor = revisionMajor

	revisionMinor, err := ReadSimpleField(ctx, "revisionMinor", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'revisionMinor' field"))
	}
	m.RevisionMinor = revisionMinor

	status, err := ReadSimpleField(ctx, "status", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'status' field"))
	}
	m.Status = status

	serialNumber, err := ReadSimpleField(ctx, "serialNumber", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serialNumber' field"))
	}
	m.SerialNumber = serialNumber

	productNameLength, err := ReadImplicitField[uint8](ctx, "productNameLength", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productNameLength' field"))
	}
	_ = productNameLength

	productName, err := ReadSimpleField(ctx, "productName", ReadString(readBuffer, uint32(int32(productNameLength)*int32(int32(8)))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'productName' field"))
	}
	m.ProductName = productName

	state, err := ReadSimpleField(ctx, "state", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'state' field"))
	}
	m.State = state

	if closeErr := readBuffer.CloseContext("CipIdentity"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipIdentity")
	}

	return m, nil
}

func (m *_CipIdentity) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipIdentity) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipIdentity"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipIdentity")
		}
		itemLength := uint16(uint16(uint16(34)) + uint16(uint8(len(m.GetProductName()))))
		if err := WriteImplicitField(ctx, "itemLength", itemLength, WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'itemLength' field")
		}

		if err := WriteSimpleField[uint16](ctx, "encapsulationProtocolVersion", m.GetEncapsulationProtocolVersion(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'encapsulationProtocolVersion' field")
		}

		if err := WriteSimpleField[uint16](ctx, "socketAddressFamily", m.GetSocketAddressFamily(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'socketAddressFamily' field")
		}

		if err := WriteSimpleField[uint16](ctx, "socketAddressPort", m.GetSocketAddressPort(), WriteUnsignedShort(writeBuffer, 16), codegen.WithByteOrder(binary.BigEndian)); err != nil {
			return errors.Wrap(err, "Error serializing 'socketAddressPort' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "socketAddressAddress", m.GetSocketAddressAddress(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'socketAddressAddress' field")
		}

		if err := WriteConstField(ctx, "zeroes1", CipIdentity_ZEROES1, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'zeroes1' field")
		}

		if err := WriteConstField(ctx, "zeroes2", CipIdentity_ZEROES2, WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'zeroes2' field")
		}

		if err := WriteSimpleField[uint16](ctx, "vendorId", m.GetVendorId(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'vendorId' field")
		}

		if err := WriteSimpleField[uint16](ctx, "deviceType", m.GetDeviceType(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'deviceType' field")
		}

		if err := WriteSimpleField[uint16](ctx, "productCode", m.GetProductCode(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'productCode' field")
		}

		if err := WriteSimpleField[uint8](ctx, "revisionMajor", m.GetRevisionMajor(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisionMajor' field")
		}

		if err := WriteSimpleField[uint8](ctx, "revisionMinor", m.GetRevisionMinor(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'revisionMinor' field")
		}

		if err := WriteSimpleField[uint16](ctx, "status", m.GetStatus(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'status' field")
		}

		if err := WriteSimpleField[uint32](ctx, "serialNumber", m.GetSerialNumber(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'serialNumber' field")
		}
		productNameLength := uint8(uint8(len(m.GetProductName())))
		if err := WriteImplicitField(ctx, "productNameLength", productNameLength, WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'productNameLength' field")
		}

		if err := WriteSimpleField[string](ctx, "productName", m.GetProductName(), WriteString(writeBuffer, int32(int32(uint8(len(m.GetProductName())))*int32(int32(8))))); err != nil {
			return errors.Wrap(err, "Error serializing 'productName' field")
		}

		if err := WriteSimpleField[uint8](ctx, "state", m.GetState(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'state' field")
		}

		if popErr := writeBuffer.PopContext("CipIdentity"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipIdentity")
		}
		return nil
	}
	return m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipIdentity) IsCipIdentity() {}

func (m *_CipIdentity) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CipIdentity) deepCopy() *_CipIdentity {
	if m == nil {
		return nil
	}
	_CipIdentityCopy := &_CipIdentity{
		m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem).deepCopy(),
		m.EncapsulationProtocolVersion,
		m.SocketAddressFamily,
		m.SocketAddressPort,
		utils.DeepCopySlice[uint8, uint8](m.SocketAddressAddress),
		m.VendorId,
		m.DeviceType,
		m.ProductCode,
		m.RevisionMajor,
		m.RevisionMinor,
		m.Status,
		m.SerialNumber,
		m.ProductName,
		m.State,
	}
	m.CommandSpecificDataItemContract.(*_CommandSpecificDataItem)._SubType = m
	return _CipIdentityCopy
}

func (m *_CipIdentity) String() string {
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
