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

// PubSubConnectionDataType is the corresponding interface of PubSubConnectionDataType
type PubSubConnectionDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetName returns Name (property field)
	GetName() PascalString
	// GetEnabled returns Enabled (property field)
	GetEnabled() bool
	// GetPublisherId returns PublisherId (property field)
	GetPublisherId() Variant
	// GetTransportProfileUri returns TransportProfileUri (property field)
	GetTransportProfileUri() PascalString
	// GetAddress returns Address (property field)
	GetAddress() ExtensionObject
	// GetNoOfConnectionProperties returns NoOfConnectionProperties (property field)
	GetNoOfConnectionProperties() int32
	// GetConnectionProperties returns ConnectionProperties (property field)
	GetConnectionProperties() []ExtensionObjectDefinition
	// GetTransportSettings returns TransportSettings (property field)
	GetTransportSettings() ExtensionObject
	// GetNoOfWriterGroups returns NoOfWriterGroups (property field)
	GetNoOfWriterGroups() int32
	// GetWriterGroups returns WriterGroups (property field)
	GetWriterGroups() []PubSubGroupDataType
	// GetNoOfReaderGroups returns NoOfReaderGroups (property field)
	GetNoOfReaderGroups() int32
	// GetReaderGroups returns ReaderGroups (property field)
	GetReaderGroups() []PubSubGroupDataType
	// IsPubSubConnectionDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsPubSubConnectionDataType()
}

// _PubSubConnectionDataType is the data-structure of this message
type _PubSubConnectionDataType struct {
	ExtensionObjectDefinitionContract
	Name                     PascalString
	Enabled                  bool
	PublisherId              Variant
	TransportProfileUri      PascalString
	Address                  ExtensionObject
	NoOfConnectionProperties int32
	ConnectionProperties     []ExtensionObjectDefinition
	TransportSettings        ExtensionObject
	NoOfWriterGroups         int32
	WriterGroups             []PubSubGroupDataType
	NoOfReaderGroups         int32
	ReaderGroups             []PubSubGroupDataType
	// Reserved Fields
	reservedField0 *uint8
}

var _ PubSubConnectionDataType = (*_PubSubConnectionDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_PubSubConnectionDataType)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_PubSubConnectionDataType) GetIdentifier() string {
	return "15619"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_PubSubConnectionDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_PubSubConnectionDataType) GetName() PascalString {
	return m.Name
}

func (m *_PubSubConnectionDataType) GetEnabled() bool {
	return m.Enabled
}

func (m *_PubSubConnectionDataType) GetPublisherId() Variant {
	return m.PublisherId
}

func (m *_PubSubConnectionDataType) GetTransportProfileUri() PascalString {
	return m.TransportProfileUri
}

func (m *_PubSubConnectionDataType) GetAddress() ExtensionObject {
	return m.Address
}

func (m *_PubSubConnectionDataType) GetNoOfConnectionProperties() int32 {
	return m.NoOfConnectionProperties
}

func (m *_PubSubConnectionDataType) GetConnectionProperties() []ExtensionObjectDefinition {
	return m.ConnectionProperties
}

func (m *_PubSubConnectionDataType) GetTransportSettings() ExtensionObject {
	return m.TransportSettings
}

func (m *_PubSubConnectionDataType) GetNoOfWriterGroups() int32 {
	return m.NoOfWriterGroups
}

func (m *_PubSubConnectionDataType) GetWriterGroups() []PubSubGroupDataType {
	return m.WriterGroups
}

func (m *_PubSubConnectionDataType) GetNoOfReaderGroups() int32 {
	return m.NoOfReaderGroups
}

func (m *_PubSubConnectionDataType) GetReaderGroups() []PubSubGroupDataType {
	return m.ReaderGroups
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewPubSubConnectionDataType factory function for _PubSubConnectionDataType
func NewPubSubConnectionDataType(name PascalString, enabled bool, publisherId Variant, transportProfileUri PascalString, address ExtensionObject, noOfConnectionProperties int32, connectionProperties []ExtensionObjectDefinition, transportSettings ExtensionObject, noOfWriterGroups int32, writerGroups []PubSubGroupDataType, noOfReaderGroups int32, readerGroups []PubSubGroupDataType) *_PubSubConnectionDataType {
	if name == nil {
		panic("name of type PascalString for PubSubConnectionDataType must not be nil")
	}
	if publisherId == nil {
		panic("publisherId of type Variant for PubSubConnectionDataType must not be nil")
	}
	if transportProfileUri == nil {
		panic("transportProfileUri of type PascalString for PubSubConnectionDataType must not be nil")
	}
	if address == nil {
		panic("address of type ExtensionObject for PubSubConnectionDataType must not be nil")
	}
	if transportSettings == nil {
		panic("transportSettings of type ExtensionObject for PubSubConnectionDataType must not be nil")
	}
	_result := &_PubSubConnectionDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		Name:                              name,
		Enabled:                           enabled,
		PublisherId:                       publisherId,
		TransportProfileUri:               transportProfileUri,
		Address:                           address,
		NoOfConnectionProperties:          noOfConnectionProperties,
		ConnectionProperties:              connectionProperties,
		TransportSettings:                 transportSettings,
		NoOfWriterGroups:                  noOfWriterGroups,
		WriterGroups:                      writerGroups,
		NoOfReaderGroups:                  noOfReaderGroups,
		ReaderGroups:                      readerGroups,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastPubSubConnectionDataType(structType any) PubSubConnectionDataType {
	if casted, ok := structType.(PubSubConnectionDataType); ok {
		return casted
	}
	if casted, ok := structType.(*PubSubConnectionDataType); ok {
		return *casted
	}
	return nil
}

func (m *_PubSubConnectionDataType) GetTypeName() string {
	return "PubSubConnectionDataType"
}

func (m *_PubSubConnectionDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (enabled)
	lengthInBits += 1

	// Simple field (publisherId)
	lengthInBits += m.PublisherId.GetLengthInBits(ctx)

	// Simple field (transportProfileUri)
	lengthInBits += m.TransportProfileUri.GetLengthInBits(ctx)

	// Simple field (address)
	lengthInBits += m.Address.GetLengthInBits(ctx)

	// Simple field (noOfConnectionProperties)
	lengthInBits += 32

	// Array field
	if len(m.ConnectionProperties) > 0 {
		for _curItem, element := range m.ConnectionProperties {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ConnectionProperties), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (transportSettings)
	lengthInBits += m.TransportSettings.GetLengthInBits(ctx)

	// Simple field (noOfWriterGroups)
	lengthInBits += 32

	// Array field
	if len(m.WriterGroups) > 0 {
		for _curItem, element := range m.WriterGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.WriterGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfReaderGroups)
	lengthInBits += 32

	// Array field
	if len(m.ReaderGroups) > 0 {
		for _curItem, element := range m.ReaderGroups {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReaderGroups), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_PubSubConnectionDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_PubSubConnectionDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__pubSubConnectionDataType PubSubConnectionDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("PubSubConnectionDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for PubSubConnectionDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[PascalString](ctx, "name", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	enabled, err := ReadSimpleField(ctx, "enabled", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'enabled' field"))
	}
	m.Enabled = enabled

	publisherId, err := ReadSimpleField[Variant](ctx, "publisherId", ReadComplex[Variant](VariantParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'publisherId' field"))
	}
	m.PublisherId = publisherId

	transportProfileUri, err := ReadSimpleField[PascalString](ctx, "transportProfileUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportProfileUri' field"))
	}
	m.TransportProfileUri = transportProfileUri

	address, err := ReadSimpleField[ExtensionObject](ctx, "address", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'address' field"))
	}
	m.Address = address

	noOfConnectionProperties, err := ReadSimpleField(ctx, "noOfConnectionProperties", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfConnectionProperties' field"))
	}
	m.NoOfConnectionProperties = noOfConnectionProperties

	connectionProperties, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "connectionProperties", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("14535")), readBuffer), uint64(noOfConnectionProperties))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'connectionProperties' field"))
	}
	m.ConnectionProperties = connectionProperties

	transportSettings, err := ReadSimpleField[ExtensionObject](ctx, "transportSettings", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'transportSettings' field"))
	}
	m.TransportSettings = transportSettings

	noOfWriterGroups, err := ReadSimpleField(ctx, "noOfWriterGroups", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfWriterGroups' field"))
	}
	m.NoOfWriterGroups = noOfWriterGroups

	writerGroups, err := ReadCountArrayField[PubSubGroupDataType](ctx, "writerGroups", ReadComplex[PubSubGroupDataType](ExtensionObjectDefinitionParseWithBufferProducer[PubSubGroupDataType]((string)("15609")), readBuffer), uint64(noOfWriterGroups))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'writerGroups' field"))
	}
	m.WriterGroups = writerGroups

	noOfReaderGroups, err := ReadSimpleField(ctx, "noOfReaderGroups", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfReaderGroups' field"))
	}
	m.NoOfReaderGroups = noOfReaderGroups

	readerGroups, err := ReadCountArrayField[PubSubGroupDataType](ctx, "readerGroups", ReadComplex[PubSubGroupDataType](ExtensionObjectDefinitionParseWithBufferProducer[PubSubGroupDataType]((string)("15609")), readBuffer), uint64(noOfReaderGroups))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readerGroups' field"))
	}
	m.ReaderGroups = readerGroups

	if closeErr := readBuffer.CloseContext("PubSubConnectionDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for PubSubConnectionDataType")
	}

	return m, nil
}

func (m *_PubSubConnectionDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_PubSubConnectionDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("PubSubConnectionDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for PubSubConnectionDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "name", m.GetName(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "enabled", m.GetEnabled(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'enabled' field")
		}

		if err := WriteSimpleField[Variant](ctx, "publisherId", m.GetPublisherId(), WriteComplex[Variant](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'publisherId' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "transportProfileUri", m.GetTransportProfileUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportProfileUri' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "address", m.GetAddress(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'address' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfConnectionProperties", m.GetNoOfConnectionProperties(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfConnectionProperties' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "connectionProperties", m.GetConnectionProperties(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'connectionProperties' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "transportSettings", m.GetTransportSettings(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'transportSettings' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfWriterGroups", m.GetNoOfWriterGroups(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfWriterGroups' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "writerGroups", m.GetWriterGroups(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'writerGroups' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfReaderGroups", m.GetNoOfReaderGroups(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfReaderGroups' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "readerGroups", m.GetReaderGroups(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'readerGroups' field")
		}

		if popErr := writeBuffer.PopContext("PubSubConnectionDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for PubSubConnectionDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_PubSubConnectionDataType) IsPubSubConnectionDataType() {}

func (m *_PubSubConnectionDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
