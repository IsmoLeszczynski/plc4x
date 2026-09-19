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

/**
 * How the registers of a value are ordered on the wire.
 * <p>
 * The layouts below are what a reader ends up seeing: {@link #swap} performs the byte and word
 * swap, the endianness is left to the buffer the swapped bytes are read from (see
 * {@link #isBigEndian()}). Reversing the bytes commutes with both swaps, so it does not matter that
 * it happens last.
 */
public enum ModbusByteOrder {
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

    private static final int REGISTER_LENGTH_BYTES = 2;

    public boolean isBigEndian() {
        return this == BIG_ENDIAN || this == BIG_ENDIAN_BYTE_SWAP || this == BIG_ENDIAN_WORD_SWAP || this == BIG_ENDIAN_WORD_SWAP_BYTE_SWAP;
    }

    /** Swaps a buffer that holds a single value. */
    public byte[] swap(byte[] in) {
        return swap(in, in.length);
    }

    /**
     * Swaps a buffer that holds one or more values of {@code elementLengthBytes} each.
     * <p>
     * A byte swap exchanges the two bytes of a register, whichever value they belong to: its stride
     * of two always falls on a register boundary, and values narrower than a register are packed two
     * to a register, so the swap has to reach across them.
     * <p>
     * A word swap exchanges adjacent registers, which only means something within one value - across
     * values it would reorder the values themselves instead of their registers. It is therefore
     * applied per element, leaving a value of a single register untouched. Values narrower than a
     * register own none of their own, so there the whole run is swapped as one.
     * <p>
     * A trailing group that the value does not fill - an odd byte, or a register without a second
     * one - has nothing to be swapped with and is passed through. The byte order is a property of
     * the device and applies to all of its tags, so a swap that does not apply to a value is a
     * no-op rather than an error.
     */
    public byte[] swap(byte[] in, int elementLengthBytes) {
        if (in.length == 0) {
            return in;
        }
        if (elementLengthBytes <= 0) {
            throw new IllegalArgumentException("element length must be greater than zero. Was " + elementLengthBytes);
        }
        return switch (this) {
            case BIG_ENDIAN_BYTE_SWAP, LITTLE_ENDIAN_BYTE_SWAP -> byteSwap(in);
            case BIG_ENDIAN_WORD_SWAP, LITTLE_ENDIAN_WORD_SWAP -> wordSwap(in, elementLengthBytes);
            case BIG_ENDIAN_WORD_SWAP_BYTE_SWAP, LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP -> wordSwap(byteSwap(in), elementLengthBytes);
            default -> in;
        };
    }

    private static byte[] byteSwap(byte[] in) {
        byte[] out = in.clone();
        for (int i = 0; i + 1 < in.length; i += REGISTER_LENGTH_BYTES) {
            out[i] = in[i + 1];
            out[i + 1] = in[i];
        }
        return out;
    }

    private static byte[] wordSwap(byte[] in, int elementLengthBytes) {
        int elementLength = elementLengthBytes < REGISTER_LENGTH_BYTES ? in.length : elementLengthBytes;
        byte[] out = in.clone();
        for (int element = 0; element < in.length; element += elementLength) {
            int elementEnd = Math.min(element + elementLength, in.length);
            for (int i = element; i + 3 < elementEnd; i += 2 * REGISTER_LENGTH_BYTES) {
                out[i] = in[i + 2];
                out[i + 1] = in[i + 3];
                out[i + 2] = in[i];
                out[i + 3] = in[i + 1];
            }
        }
        return out;
    }
}
