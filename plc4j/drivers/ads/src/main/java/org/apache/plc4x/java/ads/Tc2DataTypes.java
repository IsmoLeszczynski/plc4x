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
package org.apache.plc4x.java.ads;

import java.util.ArrayList;
import java.util.Map;

import org.apache.plc4x.java.ads.readwrite.AdsDataTypeTableEntry;
import org.apache.plc4x.java.ads.readwrite.AdsDatatypeId;

/**
 * Built-in primitive data-type definitions for TwinCAT2 (TC2) devices.
 *
 * <p>Unlike TwinCAT3, TC2 PLCs do not expose definitions for the IEC base types in their
 * uploaded data-type table. When connecting to a TC2 device (ADS major version "2") these
 * entries are injected into the data-type table so that tags using the base types can still
 * be resolved and decoded.
 *
 * <p>Only the fields consumed by the tag-resolution layer are populated (entry length, size,
 * data-type id, names). The SPI3 {@code AdsDataTypeTableEntry} flags and the optional
 * guid/method/attribute/extended-info members are left at their defaults — they only affect
 * optional-field presence during wire parsing, which is bypassed here since the entries are
 * built directly in memory.
 */
public final class Tc2DataTypes {

    private Tc2DataTypes() {
    }

    private static AdsDataTypeTableEntry entry(long entryLength, long size, long dataTypeId,
                                               String mainName, String secondaryName) {
        return new AdsDataTypeTableEntry(
            entryLength, 1L, 0L, 0L, size, 0L, AdsDatatypeId.enumForValue(dataTypeId),
            // 23 flags, all unset
            false, false, false, false, false, false, false, false, false, false, false, false,
            false, false, false, false, false, false, false, false, false, false, false,
            0, 0, mainName, secondaryName, "",
            new ArrayList<>(), new ArrayList<>(),
            null, null, null, null, new byte[0]);
    }

    // Defaults for TC2
    public static final AdsDataTypeTableEntry BOOL = entry(112, 1, 33, "BOOL", "BYTE");
    public static final AdsDataTypeTableEntry ULINT = entry(128, 8, 21, "ULINT", "");
    public static final AdsDataTypeTableEntry BYTE = entry(112, 1, 17, "BYTE", "");
    public static final AdsDataTypeTableEntry LINT = entry(144, 8, 20, "LINT", "");
    public static final AdsDataTypeTableEntry WORD = entry(112, 2, 18, "WORD", "");
    public static final AdsDataTypeTableEntry DINT = entry(128, 4, 3, "DINT", "");
    public static final AdsDataTypeTableEntry UDINT = entry(120, 4, 19, "UDINT", "");
    public static final AdsDataTypeTableEntry UINT = entry(112, 2, 18, "UINT", "");
    public static final AdsDataTypeTableEntry USINT = entry(112, 1, 17, "USINT", "");
    public static final AdsDataTypeTableEntry DWORD = entry(120, 4, 19, "DWORD", "");
    public static final AdsDataTypeTableEntry REAL = entry(120, 4, 4, "REAL", "");
    public static final AdsDataTypeTableEntry SINT = entry(120, 1, 16, "SINT", "");
    public static final AdsDataTypeTableEntry LREAL = entry(120, 8, 5, "LREAL", "");
    public static final AdsDataTypeTableEntry INT = entry(120, 2, 2, "INT", "");
    public static final AdsDataTypeTableEntry TIME = entry(72, 4, 19, "TIME", "");
    public static final AdsDataTypeTableEntry DT = entry(64, 4, 19, "DT", "");

    public static final Map<String, AdsDataTypeTableEntry> ALL = Map.ofEntries(
        Map.entry(BOOL.getMainName(), BOOL),
        Map.entry(BYTE.getMainName(), BYTE), Map.entry(WORD.getMainName(), WORD),
        Map.entry(DWORD.getMainName(), DWORD), Map.entry(SINT.getMainName(), SINT), Map.entry("INT8", SINT),
        Map.entry(INT.getMainName(), INT), Map.entry("INT16", INT), Map.entry(DINT.getMainName(), DINT),
        Map.entry("INT32", DINT), Map.entry(LINT.getMainName(), LINT), Map.entry("INT64", LINT),
        Map.entry(USINT.getMainName(), USINT), Map.entry("UINT8", USINT), Map.entry(UINT.getMainName(), UINT),
        Map.entry("UINT16", UINT), Map.entry(UDINT.getMainName(), UDINT), Map.entry("UINT32", UDINT),
        Map.entry(ULINT.getMainName(), ULINT), Map.entry("UINT64", ULINT), Map.entry(REAL.getMainName(), REAL),
        Map.entry(LREAL.getMainName(), LREAL), Map.entry(TIME.getMainName(), TIME),
        Map.entry(DT.getMainName(), DT));

}
