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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableKey is an enum
type OpcuaNodeIdServicesVariableKey int32

type IOpcuaNodeIdServicesVariableKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri                      OpcuaNodeIdServicesVariableKey = 17512
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri                       OpcuaNodeIdServicesVariableKey = 17513
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls                     OpcuaNodeIdServicesVariableKey = 17514
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus                    OpcuaNodeIdServicesVariableKey = 17515
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments  OpcuaNodeIdServicesVariableKey = 17517
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments OpcuaNodeIdServicesVariableKey = 17518
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments  OpcuaNodeIdServicesVariableKey = 17520
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_InputArguments                          OpcuaNodeIdServicesVariableKey = 17523
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments                         OpcuaNodeIdServicesVariableKey = 17524
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_InputArguments                                    OpcuaNodeIdServicesVariableKey = 17529
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_OutputArguments                                   OpcuaNodeIdServicesVariableKey = 17530
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_InputArguments                                OpcuaNodeIdServicesVariableKey = 17535
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments                               OpcuaNodeIdServicesVariableKey = 17536
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_EndpointUrls                                                   OpcuaNodeIdServicesVariableKey = 18004
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ServiceStatus                                                  OpcuaNodeIdServicesVariableKey = 18005
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_UpdateCredential_InputArguments                                OpcuaNodeIdServicesVariableKey = 18007
	OpcuaNodeIdServicesVariableKey_KeyCredentialUpdateMethodType_InputArguments                                                  OpcuaNodeIdServicesVariableKey = 18010
	OpcuaNodeIdServicesVariableKey_KeyCredentialAuditEventType_ResourceUri                                                       OpcuaNodeIdServicesVariableKey = 18028
	OpcuaNodeIdServicesVariableKey_KeyCredentialDeletedAuditEventType_ResourceUri                                                OpcuaNodeIdServicesVariableKey = 18064
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ResourceUri                                                    OpcuaNodeIdServicesVariableKey = 18069
	OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ProfileUri                                                     OpcuaNodeIdServicesVariableKey = 18165
)

var OpcuaNodeIdServicesVariableKeyValues []OpcuaNodeIdServicesVariableKey

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableKeyValues = []OpcuaNodeIdServicesVariableKey{
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_OutputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_EndpointUrls,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ServiceStatus,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_UpdateCredential_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialUpdateMethodType_InputArguments,
		OpcuaNodeIdServicesVariableKey_KeyCredentialAuditEventType_ResourceUri,
		OpcuaNodeIdServicesVariableKey_KeyCredentialDeletedAuditEventType_ResourceUri,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ResourceUri,
		OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ProfileUri,
	}
}

func OpcuaNodeIdServicesVariableKeyByValue(value int32) (enum OpcuaNodeIdServicesVariableKey, ok bool) {
	switch value {
	case 17512:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri, true
	case 17513:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri, true
	case 17514:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls, true
	case 17515:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus, true
	case 17517:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments, true
	case 17518:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments, true
	case 17520:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments, true
	case 17523:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_InputArguments, true
	case 17524:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments, true
	case 17529:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_InputArguments, true
	case 17530:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_OutputArguments, true
	case 17535:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_InputArguments, true
	case 17536:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments, true
	case 18004:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_EndpointUrls, true
	case 18005:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ServiceStatus, true
	case 18007:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_UpdateCredential_InputArguments, true
	case 18010:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialUpdateMethodType_InputArguments, true
	case 18028:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialAuditEventType_ResourceUri, true
	case 18064:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialDeletedAuditEventType_ResourceUri, true
	case 18069:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ResourceUri, true
	case 18165:
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ProfileUri, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableKeyByName(value string) (enum OpcuaNodeIdServicesVariableKey, ok bool) {
	switch value {
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments, true
	case "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments, true
	case "KeyCredentialConfigurationFolderType_CreateCredential_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_InputArguments, true
	case "KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments, true
	case "KeyCredentialConfiguration_CreateCredential_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_InputArguments, true
	case "KeyCredentialConfiguration_CreateCredential_OutputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_OutputArguments, true
	case "KeyCredentialConfigurationType_GetEncryptingKey_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_InputArguments, true
	case "KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments, true
	case "KeyCredentialConfigurationType_EndpointUrls":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_EndpointUrls, true
	case "KeyCredentialConfigurationType_ServiceStatus":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ServiceStatus, true
	case "KeyCredentialConfigurationType_UpdateCredential_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_UpdateCredential_InputArguments, true
	case "KeyCredentialUpdateMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialUpdateMethodType_InputArguments, true
	case "KeyCredentialAuditEventType_ResourceUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialAuditEventType_ResourceUri, true
	case "KeyCredentialDeletedAuditEventType_ResourceUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialDeletedAuditEventType_ResourceUri, true
	case "KeyCredentialConfigurationType_ResourceUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ResourceUri, true
	case "KeyCredentialConfigurationType_ProfileUri":
		return OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ProfileUri, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableKeyKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableKeyValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableKey(structType any) OpcuaNodeIdServicesVariableKey {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableKey {
		if sOpcuaNodeIdServicesVariableKey, ok := typ.(OpcuaNodeIdServicesVariableKey); ok {
			return sOpcuaNodeIdServicesVariableKey
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableKey) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableKeyParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableKey, error) {
	return OpcuaNodeIdServicesVariableKeyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableKeyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableKey, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableKey", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableKey")
	}
	if enum, ok := OpcuaNodeIdServicesVariableKeyByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableKey")
		return OpcuaNodeIdServicesVariableKey(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableKey", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableKey) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ResourceUri"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ProfileUri"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_EndpointUrls"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_ServiceStatus"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_GetEncryptingKey_OutputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments:
		return "KeyCredentialConfigurationFolderType_ServiceName_Placeholder_UpdateCredential_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_InputArguments:
		return "KeyCredentialConfigurationFolderType_CreateCredential_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments:
		return "KeyCredentialConfigurationFolderType_CreateCredential_OutputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_InputArguments:
		return "KeyCredentialConfiguration_CreateCredential_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfiguration_CreateCredential_OutputArguments:
		return "KeyCredentialConfiguration_CreateCredential_OutputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_InputArguments:
		return "KeyCredentialConfigurationType_GetEncryptingKey_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments:
		return "KeyCredentialConfigurationType_GetEncryptingKey_OutputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_EndpointUrls:
		return "KeyCredentialConfigurationType_EndpointUrls"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ServiceStatus:
		return "KeyCredentialConfigurationType_ServiceStatus"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_UpdateCredential_InputArguments:
		return "KeyCredentialConfigurationType_UpdateCredential_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialUpdateMethodType_InputArguments:
		return "KeyCredentialUpdateMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialAuditEventType_ResourceUri:
		return "KeyCredentialAuditEventType_ResourceUri"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialDeletedAuditEventType_ResourceUri:
		return "KeyCredentialDeletedAuditEventType_ResourceUri"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ResourceUri:
		return "KeyCredentialConfigurationType_ResourceUri"
	case OpcuaNodeIdServicesVariableKey_KeyCredentialConfigurationType_ProfileUri:
		return "KeyCredentialConfigurationType_ProfileUri"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableKey) String() string {
	return e.PLC4XEnumName()
}
