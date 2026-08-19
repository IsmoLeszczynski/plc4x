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
package org.apache.plc4x.java.modbus.types;

import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;public enum ModbusByteOrder {
    // [1, 2, 3, 4]
    // [1, 2, 3, 4, 5, 6, 7, 8]
    BIG_ENDIAN,
    // [4, 3, 2, 1]
    // [8, 7, 6, 5, 4, 3, 2, 1]
    LITTLE_ENDIAN,
    // [2, 1, 4, 3]
    // [2, 1, 4, 3, 6, 5, 8, 7]
    BIG_ENDIAN_BYTE_SWAP,
    // [3, 4, 1, 2]
    // [7, 8, 5, 6, 3, 4, 1, 2]
    LITTLE_ENDIAN_BYTE_SWAP,
    // [3, 4, 1, 2]
    // [3, 4, 1, 2, 7, 8, 5, 6]
    BIG_ENDIAN_WORD_SWAP,
    // [2, 1, 4, 3]
    // [6, 5, 8, 7, 2, 1, 4, 3]
    LITTLE_ENDIAN_WORD_SWAP,
    // [4, 3, 2, 1]
    // [4, 3, 2, 1, 8, 7, 6, 5]
    BIG_ENDIAN_WORD_SWAP_BYTE_SWAP,
    // [1, 2, 3, 4]
    // [5, 6, 7, 8, 1, 2, 3, 4]
    LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP;

    public boolean isBigEndian() {
        return this == BIG_ENDIAN || this == BIG_ENDIAN_BYTE_SWAP || this == BIG_ENDIAN_WORD_SWAP || this == BIG_ENDIAN_WORD_SWAP_BYTE_SWAP;
    }

    public byte[] swap(byte[] in) {
        return switch (this) {
            case BIG_ENDIAN_BYTE_SWAP, LITTLE_ENDIAN_BYTE_SWAP -> byteSwap(in);
            case BIG_ENDIAN_WORD_SWAP, LITTLE_ENDIAN_WORD_SWAP -> wordSwap(in);
            case BIG_ENDIAN_WORD_SWAP_BYTE_SWAP, LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP -> wordSwap(byteSwap(in));
            default -> in;
        };
    }

    private static byte[] byteSwap(byte[] in) {
        byte[] out = new byte[in.length];
        for (int i = 0; i < out.length - 1; i += 2) {
            out[i] = in[i + 1];
            out[i + 1] = in[i];
        }
        // Handle odd-length arrays
        if (in.length % 2 != 0) {
            out[in.length - 1] = in[in.length - 1];
        }
        return out;
    }

    private static byte[] wordSwap(byte[] in) {
        if (in.length % 2 != 0) {
            throw new PlcRuntimeException("Input byte array length must be a multiple of 2 for word swapping.");
        }
        byte[] out = new byte[in.length];
        for (int i = 0; i < in.length; i += 4) {
            out[i] = in[i + 2];
            out[i + 1] = in[i + 3];
            out[i + 2] = in[i];
            out[i + 3] = in[i + 1];
        }
        return out;
    }
}
