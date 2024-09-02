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

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableNamespace is an enum
type OpcuaNodeIdServicesVariableNamespace int32

type IOpcuaNodeIdServicesVariableNamespace interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceUri                              OpcuaNodeIdServicesVariableNamespace = 11617
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceVersion                          OpcuaNodeIdServicesVariableNamespace = 11618
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespacePublicationDate                  OpcuaNodeIdServicesVariableNamespace = 11619
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_IsNamespaceSubset                         OpcuaNodeIdServicesVariableNamespace = 11620
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNodeIdTypes                         OpcuaNodeIdServicesVariableNamespace = 11621
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNumericNodeIdRange                  OpcuaNodeIdServicesVariableNamespace = 11622
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticStringNodeIdPattern                 OpcuaNodeIdServicesVariableNamespace = 11623
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Size                        OpcuaNodeIdServicesVariableNamespace = 11625
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_OpenCount                   OpcuaNodeIdServicesVariableNamespace = 11628
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_InputArguments         OpcuaNodeIdServicesVariableNamespace = 11630
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_OutputArguments        OpcuaNodeIdServicesVariableNamespace = 11631
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Close_InputArguments        OpcuaNodeIdServicesVariableNamespace = 11633
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_InputArguments         OpcuaNodeIdServicesVariableNamespace = 11635
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_OutputArguments        OpcuaNodeIdServicesVariableNamespace = 11636
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Write_InputArguments        OpcuaNodeIdServicesVariableNamespace = 11638
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments  OpcuaNodeIdServicesVariableNamespace = 11640
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments OpcuaNodeIdServicesVariableNamespace = 11641
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments  OpcuaNodeIdServicesVariableNamespace = 11643
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Writable                    OpcuaNodeIdServicesVariableNamespace = 12690
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_UserWritable                OpcuaNodeIdServicesVariableNamespace = 12691
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MimeType                    OpcuaNodeIdServicesVariableNamespace = 13399
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultRolePermissions                    OpcuaNodeIdServicesVariableNamespace = 16137
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultUserRolePermissions                OpcuaNodeIdServicesVariableNamespace = 16138
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultAccessRestrictions                 OpcuaNodeIdServicesVariableNamespace = 16139
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MaxByteStringLength         OpcuaNodeIdServicesVariableNamespace = 24246
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_LastModifiedTime            OpcuaNodeIdServicesVariableNamespace = 25202
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ConfigurationVersion                      OpcuaNodeIdServicesVariableNamespace = 25267
	OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ModelVersion                              OpcuaNodeIdServicesVariableNamespace = 32419
)

var OpcuaNodeIdServicesVariableNamespaceValues []OpcuaNodeIdServicesVariableNamespace

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableNamespaceValues = []OpcuaNodeIdServicesVariableNamespace{
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceUri,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceVersion,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespacePublicationDate,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_IsNamespaceSubset,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNodeIdTypes,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNumericNodeIdRange,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticStringNodeIdPattern,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Size,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_OpenCount,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_OutputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Close_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_OutputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Write_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Writable,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_UserWritable,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MimeType,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultRolePermissions,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultUserRolePermissions,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultAccessRestrictions,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MaxByteStringLength,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_LastModifiedTime,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ConfigurationVersion,
		OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ModelVersion,
	}
}

func OpcuaNodeIdServicesVariableNamespaceByValue(value int32) (enum OpcuaNodeIdServicesVariableNamespace, ok bool) {
	switch value {
	case 11617:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceUri, true
	case 11618:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceVersion, true
	case 11619:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespacePublicationDate, true
	case 11620:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_IsNamespaceSubset, true
	case 11621:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNodeIdTypes, true
	case 11622:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNumericNodeIdRange, true
	case 11623:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticStringNodeIdPattern, true
	case 11625:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Size, true
	case 11628:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_OpenCount, true
	case 11630:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_InputArguments, true
	case 11631:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_OutputArguments, true
	case 11633:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Close_InputArguments, true
	case 11635:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_InputArguments, true
	case 11636:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_OutputArguments, true
	case 11638:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Write_InputArguments, true
	case 11640:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments, true
	case 11641:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments, true
	case 11643:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments, true
	case 12690:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Writable, true
	case 12691:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_UserWritable, true
	case 13399:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MimeType, true
	case 16137:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultRolePermissions, true
	case 16138:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultUserRolePermissions, true
	case 16139:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultAccessRestrictions, true
	case 24246:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MaxByteStringLength, true
	case 25202:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_LastModifiedTime, true
	case 25267:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ConfigurationVersion, true
	case 32419:
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ModelVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNamespaceByName(value string) (enum OpcuaNodeIdServicesVariableNamespace, ok bool) {
	switch value {
	case "NamespaceMetadataType_NamespaceUri":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceUri, true
	case "NamespaceMetadataType_NamespaceVersion":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceVersion, true
	case "NamespaceMetadataType_NamespacePublicationDate":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespacePublicationDate, true
	case "NamespaceMetadataType_IsNamespaceSubset":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_IsNamespaceSubset, true
	case "NamespaceMetadataType_StaticNodeIdTypes":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNodeIdTypes, true
	case "NamespaceMetadataType_StaticNumericNodeIdRange":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNumericNodeIdRange, true
	case "NamespaceMetadataType_StaticStringNodeIdPattern":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticStringNodeIdPattern, true
	case "NamespaceMetadataType_NamespaceFile_Size":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Size, true
	case "NamespaceMetadataType_NamespaceFile_OpenCount":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_OpenCount, true
	case "NamespaceMetadataType_NamespaceFile_Open_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Open_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_OutputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Close_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Close_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Read_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Read_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_OutputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Write_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Write_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments, true
	case "NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments, true
	case "NamespaceMetadataType_NamespaceFile_Writable":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Writable, true
	case "NamespaceMetadataType_NamespaceFile_UserWritable":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_UserWritable, true
	case "NamespaceMetadataType_NamespaceFile_MimeType":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MimeType, true
	case "NamespaceMetadataType_DefaultRolePermissions":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultRolePermissions, true
	case "NamespaceMetadataType_DefaultUserRolePermissions":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultUserRolePermissions, true
	case "NamespaceMetadataType_DefaultAccessRestrictions":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultAccessRestrictions, true
	case "NamespaceMetadataType_NamespaceFile_MaxByteStringLength":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MaxByteStringLength, true
	case "NamespaceMetadataType_NamespaceFile_LastModifiedTime":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_LastModifiedTime, true
	case "NamespaceMetadataType_ConfigurationVersion":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ConfigurationVersion, true
	case "NamespaceMetadataType_ModelVersion":
		return OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ModelVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableNamespaceKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableNamespaceValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableNamespace(structType any) OpcuaNodeIdServicesVariableNamespace {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableNamespace {
		if sOpcuaNodeIdServicesVariableNamespace, ok := typ.(OpcuaNodeIdServicesVariableNamespace); ok {
			return sOpcuaNodeIdServicesVariableNamespace
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableNamespace) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableNamespace) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableNamespaceParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableNamespace, error) {
	return OpcuaNodeIdServicesVariableNamespaceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableNamespaceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableNamespace, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.ReadInt32("OpcuaNodeIdServicesVariableNamespace", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableNamespace")
	}
	if enum, ok := OpcuaNodeIdServicesVariableNamespaceByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableNamespace")
		return OpcuaNodeIdServicesVariableNamespace(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableNamespace) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableNamespace) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return /*TODO: migrate me*/ writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableNamespace", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

func (e OpcuaNodeIdServicesVariableNamespace) GetValue() int32 {
	return int32(e)
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableNamespace) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceUri:
		return "NamespaceMetadataType_NamespaceUri"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceVersion:
		return "NamespaceMetadataType_NamespaceVersion"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespacePublicationDate:
		return "NamespaceMetadataType_NamespacePublicationDate"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_IsNamespaceSubset:
		return "NamespaceMetadataType_IsNamespaceSubset"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNodeIdTypes:
		return "NamespaceMetadataType_StaticNodeIdTypes"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticNumericNodeIdRange:
		return "NamespaceMetadataType_StaticNumericNodeIdRange"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_StaticStringNodeIdPattern:
		return "NamespaceMetadataType_StaticStringNodeIdPattern"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Size:
		return "NamespaceMetadataType_NamespaceFile_Size"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_OpenCount:
		return "NamespaceMetadataType_NamespaceFile_OpenCount"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_Open_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Open_OutputArguments:
		return "NamespaceMetadataType_NamespaceFile_Open_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Close_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_Close_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_Read_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Read_OutputArguments:
		return "NamespaceMetadataType_NamespaceFile_Read_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Write_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_Write_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_GetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments:
		return "NamespaceMetadataType_NamespaceFile_GetPosition_OutputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments:
		return "NamespaceMetadataType_NamespaceFile_SetPosition_InputArguments"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_Writable:
		return "NamespaceMetadataType_NamespaceFile_Writable"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_UserWritable:
		return "NamespaceMetadataType_NamespaceFile_UserWritable"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MimeType:
		return "NamespaceMetadataType_NamespaceFile_MimeType"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultRolePermissions:
		return "NamespaceMetadataType_DefaultRolePermissions"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultUserRolePermissions:
		return "NamespaceMetadataType_DefaultUserRolePermissions"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_DefaultAccessRestrictions:
		return "NamespaceMetadataType_DefaultAccessRestrictions"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_MaxByteStringLength:
		return "NamespaceMetadataType_NamespaceFile_MaxByteStringLength"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_NamespaceFile_LastModifiedTime:
		return "NamespaceMetadataType_NamespaceFile_LastModifiedTime"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ConfigurationVersion:
		return "NamespaceMetadataType_ConfigurationVersion"
	case OpcuaNodeIdServicesVariableNamespace_NamespaceMetadataType_ModelVersion:
		return "NamespaceMetadataType_ModelVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableNamespace) String() string {
	return e.PLC4XEnumName()
}
