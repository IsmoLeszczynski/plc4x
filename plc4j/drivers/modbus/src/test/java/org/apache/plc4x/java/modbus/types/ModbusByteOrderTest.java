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

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import java.nio.charset.StandardCharsets;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;
import org.junit.jupiter.params.provider.EnumSource;

/**
 * Verifies {@link ModbusByteOrder} against the Modbus endianness/swap specification.
 * <p>
 * Bytes are spelled as letters so the register layout stays readable:
 * <pre>
 * Bytes:      A B C D E F G H
 * Registers:  AB | CD | EF | GH
 * </pre>
 * Note that {@link ModbusByteOrder#swap(byte[])} implements only steps 2 and 3 of the
 * specification (byte swap, word swap). Step 1 (endianness) is applied by the caller, which
 * hands the swapped bytes to a read buffer configured from {@link ModbusByteOrder#isBigEndian()};
 * a little-endian read buffer reverses the whole sequence. Tests that assert the specification's
 * end-to-end examples therefore go through {@link #pipeline(ModbusByteOrder, String)}.
 */
class ModbusByteOrderTest {

    @Nested
    @DisplayName("swap() - byte swap and word swap only, no endianness")
    class Swap {

        @ParameterizedTest(name = "{0}: ABCDEFGH -> {1}")
        @CsvSource({
            "BIG_ENDIAN,                          ABCDEFGH",
            "LITTLE_ENDIAN,                       ABCDEFGH",
            "BIG_ENDIAN_BYTE_SWAP,                BADCFEHG",
            "LITTLE_ENDIAN_BYTE_SWAP,             BADCFEHG",
            "BIG_ENDIAN_WORD_SWAP,                CDABGHEF",
            "LITTLE_ENDIAN_WORD_SWAP,             CDABGHEF",
            "BIG_ENDIAN_WORD_SWAP_BYTE_SWAP,      DCBAHGFE",
            "LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP,   DCBAHGFE"})
        void sixtyFourBit(ModbusByteOrder order, String expected) {
            assertSwap(order, "ABCDEFGH", expected);
        }

        @ParameterizedTest(name = "{0}: ABCD -> {1}")
        @CsvSource({
            "BIG_ENDIAN,                          ABCD",
            "LITTLE_ENDIAN,                       ABCD",
            "BIG_ENDIAN_BYTE_SWAP,                BADC",
            "LITTLE_ENDIAN_BYTE_SWAP,             BADC",
            "BIG_ENDIAN_WORD_SWAP,                CDAB",
            "LITTLE_ENDIAN_WORD_SWAP,             CDAB",
            "BIG_ENDIAN_WORD_SWAP_BYTE_SWAP,      DCBA",
            "LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP,   DCBA"})
        void thirtyTwoBit(ModbusByteOrder order, String expected) {
            assertSwap(order, "ABCD", expected);
        }

        @Test
        @DisplayName("byte swap swaps the two bytes inside every register")
        void byteSwapPermutation() {
            assertSwap(ModbusByteOrder.BIG_ENDIAN_BYTE_SWAP, "ABCDEFGHIJKLMNOP", "BADCFEHGJILKNMPO");
        }

        @Test
        @DisplayName("word swap swaps adjacent register pairs, it does not reverse the register order")
        void wordSwapIsAdjacentPairsOnly() {
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP, "ABCDEFGHIJKLMNOP", "CDABGHEFKLIJOPMN");
            // Full register reversal would be the wrong reading of the spec.
            assertNotEquals("OPMNKLIJGHEFCDAB",
                str(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP.swap(bytes("ABCDEFGHIJKLMNOP"))));
        }

        @Test
        @DisplayName("byte swap and word swap commute")
        void swapsCommute() {
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP_BYTE_SWAP, "ABCDEFGH", "DCBAHGFE");
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP_BYTE_SWAP, "ABCD", "DCBA");
        }

        @ParameterizedTest
        @EnumSource(ModbusByteOrder.class)
        void doesNotMutateTheInput(ModbusByteOrder order) {
            byte[] in = bytes("ABCDEFGH");
            order.swap(in);
            assertArrayEquals(bytes("ABCDEFGH"), in);
        }
    }

    @Nested
    @DisplayName("endianness + swap - the specification's 64-bit example table")
    class SpecificationExamples {

        @ParameterizedTest(name = "{0}: ABCDEFGH -> {1}")
        @CsvSource({
            "BIG_ENDIAN,                          ABCDEFGH",
            "BIG_ENDIAN_BYTE_SWAP,                BADCFEHG",
            "BIG_ENDIAN_WORD_SWAP,                CDABGHEF",
            "BIG_ENDIAN_WORD_SWAP_BYTE_SWAP,      DCBAHGFE",
            "LITTLE_ENDIAN,                       HGFEDCBA",
            "LITTLE_ENDIAN_BYTE_SWAP,             GHEFCDAB",
            "LITTLE_ENDIAN_WORD_SWAP,             FEHGBADC",
            "LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP,   EFGHABCD"})
        void sixtyFourBit(ModbusByteOrder order, String expected) {
            assertEquals(expected, pipeline(order, "ABCDEFGH"));
        }

        @ParameterizedTest(name = "{0}: ABCD -> {1}")
        @CsvSource({
            "BIG_ENDIAN,                          ABCD",
            "BIG_ENDIAN_BYTE_SWAP,                BADC",
            "BIG_ENDIAN_WORD_SWAP,                CDAB",
            "BIG_ENDIAN_WORD_SWAP_BYTE_SWAP,      DCBA",
            "LITTLE_ENDIAN,                       DCBA",
            "LITTLE_ENDIAN_BYTE_SWAP,             CDAB",
            "LITTLE_ENDIAN_WORD_SWAP,             BADC",
            "LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP,   ABCD"})
        void thirtyTwoBit(ModbusByteOrder order, String expected) {
            assertEquals(expected, pipeline(order, "ABCD"));
        }

        @Test
        @DisplayName("little endian + byte swap yields reversed register order for 64-bit values")
        void littleEndianByteSwapReversesRegisters() {
            assertEquals("GHEFCDAB", pipeline(ModbusByteOrder.LITTLE_ENDIAN_BYTE_SWAP, "ABCDEFGH"));
        }

        @ParameterizedTest
        @EnumSource(ModbusByteOrder.class)
        @DisplayName("swapping before or after the endianness reversal gives the same result")
        void endiannessCommutesWithTheSwaps(ModbusByteOrder order) {
            byte[] endiannessFirst = order.isBigEndian() ? bytes("ABCDEFGH") : reverse(bytes("ABCDEFGH"));
            assertEquals(pipeline(order, "ABCDEFGH"), str(order.swap(endiannessFirst)));
        }
    }

    @Nested
    @DisplayName("ragged lengths - an incomplete trailing group is passed through")
    class RaggedLengths {

        @ParameterizedTest
        @EnumSource(value = ModbusByteOrder.class, names = {"BIG_ENDIAN_BYTE_SWAP", "LITTLE_ENDIAN_BYTE_SWAP"})
        @DisplayName("byte swap passes a trailing odd byte through")
        void byteSwapPassesTheOddByteThrough(ModbusByteOrder order) {
            assertSwap(order, "ABC", "BAC");
            assertSwap(order, "ABCDE", "BADCE");
        }

        @ParameterizedTest
        @EnumSource(value = ModbusByteOrder.class, names = {"BIG_ENDIAN_WORD_SWAP", "LITTLE_ENDIAN_WORD_SWAP"})
        @DisplayName("word swap leaves a value of a single register alone")
        void wordSwapLeavesASingleRegisterAlone(ModbusByteOrder order) {
            assertSwap(order, "AB", "AB");
        }

        @ParameterizedTest
        @EnumSource(value = ModbusByteOrder.class, names = {"BIG_ENDIAN_WORD_SWAP", "LITTLE_ENDIAN_WORD_SWAP"})
        @DisplayName("word swap passes a register without a second one through")
        void wordSwapPassesTheUnpairedRegisterThrough(ModbusByteOrder order) {
            assertSwap(order, "ABCDEF", "CDABEF");
            assertSwap(order, "ABC", "ABC");
        }

        @ParameterizedTest
        @EnumSource(value = ModbusByteOrder.class, names = {"BIG_ENDIAN_WORD_SWAP_BYTE_SWAP", "LITTLE_ENDIAN_WORD_SWAP_BYTE_SWAP"})
        @DisplayName("both swaps together pass their own remainder through")
        void bothSwapsPassTheRemainderThrough(ModbusByteOrder order) {
            assertSwap(order, "AB", "BA");
            assertSwap(order, "ABCDEF", "DCBAFE");
        }

        @ParameterizedTest
        @EnumSource(ModbusByteOrder.class)
        @DisplayName("no length throws")
        void noLengthThrows(ModbusByteOrder order) {
            for (int length = 0; length < 10; length++) {
                byte[] in = new byte[length];
                assertDoesNotThrow(() -> order.swap(in), order.name() + " length " + in.length);
            }
        }
    }

    @Nested
    @DisplayName("arrays - a word swap stays inside one value")
    class ValueArrays {

        @Test
        @DisplayName("a word swap does not reorder the elements of a 16-bit array")
        void wordSwapIsANoOpForSingleRegisterElements() {
            // ABCDEFGH is four INTs. Swapping registers across them would hand back B, A, D, C.
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP, "ABCDEFGH", 2, "ABCDEFGH");
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP_BYTE_SWAP, "ABCDEFGH", 2, "BADCFEHG");
        }

        @Test
        @DisplayName("a word swap of a 32-bit array swaps the registers of each element")
        void wordSwapOfThirtyTwoBitElements() {
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP, "ABCDEFGH", 4, "CDABGHEF");
        }

        @Test
        @DisplayName("a word swap of a 64-bit array swaps the register pairs of each element")
        void wordSwapOfSixtyFourBitElements() {
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP, "ABCDEFGHIJKLMNOP", 8, "CDABGHEFKLIJOPMN");
        }

        @Test
        @DisplayName("values narrower than a register are swapped as one run")
        void subRegisterElementsAreSwappedAsOneRun() {
            // Eight CHARs are packed two to a register, so the registers belong to the run.
            assertSwap(ModbusByteOrder.BIG_ENDIAN_BYTE_SWAP, "ABCDEFGH", 1, "BADCFEHG");
            assertSwap(ModbusByteOrder.BIG_ENDIAN_WORD_SWAP, "ABCDEFGH", 1, "CDABGHEF");
        }

        @Test
        @DisplayName("a byte swap needs no element length, its stride is a register")
        void byteSwapIsElementAgnostic() {
            for (int elementLength = 1; elementLength <= 8; elementLength++) {
                assertSwap(ModbusByteOrder.BIG_ENDIAN_BYTE_SWAP, "ABCDEFGH", elementLength, "BADCFEHG");
            }
        }

        @ParameterizedTest
        @EnumSource(ModbusByteOrder.class)
        void rejectsANonPositiveElementLength(ModbusByteOrder order) {
            assertThrows(IllegalArgumentException.class, () -> order.swap(bytes("ABCD"), 0));
        }
    }

    @Nested
    @DisplayName("isBigEndian()")
    class Endianness {

        @ParameterizedTest
        @EnumSource(ModbusByteOrder.class)
        void namesTheEndianness(ModbusByteOrder order) {
            assertEquals(order.name().startsWith("BIG_ENDIAN"), order.isBigEndian(), order.name());
        }
    }

    private static void assertSwap(ModbusByteOrder order, String in, String expected) {
        assertEquals(expected, str(order.swap(bytes(in))), order.name());
    }

    private static void assertSwap(ModbusByteOrder order, String in, int elementLength, String expected) {
        assertEquals(expected, str(order.swap(bytes(in), elementLength)),
            order.name() + ", element length " + elementLength);
    }

    /** Byte order as the Modbus connections apply it: swap first, then endianness in the read buffer. */
    private static String pipeline(ModbusByteOrder order, String in) {
        byte[] swapped = order.swap(bytes(in));
        return str(order.isBigEndian() ? swapped : reverse(swapped));
    }

    private static byte[] reverse(byte[] in) {
        byte[] out = new byte[in.length];
        for (int i = 0; i < in.length; i++) {
            out[i] = in[in.length - 1 - i];
        }
        return out;
    }

    private static byte[] bytes(String letters) {
        return letters.getBytes(StandardCharsets.US_ASCII);
    }

    private static String str(byte[] in) {
        return new String(in, StandardCharsets.US_ASCII);
    }
}
