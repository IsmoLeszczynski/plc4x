/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.eip.base;

import org.apache.plc4x.java.eip.base.EipTcpConnection.Correlator;
import org.apache.plc4x.java.eip.readwrite.EipPacket;
import org.junit.jupiter.api.Test;
import org.mockito.Mockito;

import java.nio.charset.StandardCharsets;
import java.util.concurrent.CompletableFuture;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertSame;
import static org.junit.jupiter.api.Assertions.assertTrue;

/**
 * Responses are matched to requests by the sender context the device echoes back, so the
 * interesting cases are the ones send order alone could not settle: an answer that arrives late,
 * and a device that does not echo the context at all.
 */
class EipCorrelationTest {

    private static EipPacket packetWithContext(byte[] senderContext) {
        EipPacket packet = Mockito.mock(EipPacket.class);
        Mockito.when(packet.getSenderContext()).thenReturn(senderContext);
        return packet;
    }

    /** Registers a request and returns the context its answer has to carry. */
    private static byte[] send(Correlator correlator, CompletableFuture<EipPacket> request) {
        byte[] context = correlator.nextContext();
        correlator.register(context, request);
        return context;
    }

    @Test
    void aResponseAnswersTheRequestWhoseContextItCarries() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> first = new CompletableFuture<>();
        CompletableFuture<EipPacket> second = new CompletableFuture<>();
        byte[] firstContext = send(correlator, first);
        byte[] secondContext = send(correlator, second);

        // Out of send order: the context decides, not the queue position.
        EipPacket secondResponse = packetWithContext(secondContext);
        EipPacket firstResponse = packetWithContext(firstContext);
        correlator.deliver(secondResponse);
        correlator.deliver(firstResponse);

        assertSame(firstResponse, first.getNow(null));
        assertSame(secondResponse, second.getNow(null));
    }

    /**
     * The case that used to hand one caller another caller's data: a request gives up, its answer
     * arrives late, and by then somebody else is waiting.
     */
    @Test
    void aLateResponseIsNotGivenToTheNextRequest() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> abandoned = new CompletableFuture<>();
        byte[] abandonedContext = send(correlator, abandoned);
        correlator.timedOut(abandoned);
        CompletableFuture<EipPacket> next = new CompletableFuture<>();
        send(correlator, next);

        correlator.deliver(packetWithContext(abandonedContext));

        assertFalse(next.isDone(), "a late answer to an abandoned request must not complete a later one");
        assertEquals(1, correlator.pendingCount());
    }

    /**
     * The case that used to time every later request out: with matching by send order, the
     * response after a timeout had to be discarded on suspicion, which cost the next request its
     * answer, and so on for as long as the connection lived.
     */
    @Test
    void aRequestThatGaveUpDoesNotCostTheNextOneItsAnswer() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> abandoned = new CompletableFuture<>();
        send(correlator, abandoned);
        correlator.timedOut(abandoned);
        // The abandoned request is never answered at all.

        CompletableFuture<EipPacket> next = new CompletableFuture<>();
        byte[] nextContext = send(correlator, next);
        EipPacket own = packetWithContext(nextContext);
        correlator.deliver(own);

        assertSame(own, next.getNow(null));
    }

    @Test
    void aDeviceThatDoesNotEchoTheContextIsMatchedBySendOrder() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> only = new CompletableFuture<>();
        send(correlator, only);

        EipPacket response = packetWithContext("PLC4X   ".getBytes(StandardCharsets.US_ASCII));
        correlator.deliver(response);

        assertSame(response, only.getNow(null), "with one request outstanding there is nobody else it could be for");
    }

    @Test
    void anUnechoedContextIsNotGuessedAtWhileSeveralRequestsWait() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> a = new CompletableFuture<>();
        CompletableFuture<EipPacket> b = new CompletableFuture<>();
        send(correlator, a);
        send(correlator, b);

        correlator.deliver(packetWithContext(new byte[8]));

        assertFalse(a.isDone());
        assertFalse(b.isDone());
    }

    @Test
    void aRequestThatWasNeverSentIsJustForgotten() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> unsent = new CompletableFuture<>();
        send(correlator, unsent);
        correlator.forget(unsent);

        assertEquals(0, correlator.pendingCount());
    }

    @Test
    void closingFailsEveryoneStillWaiting() {
        Correlator correlator = new Correlator();
        CompletableFuture<EipPacket> first = new CompletableFuture<>();
        CompletableFuture<EipPacket> second = new CompletableFuture<>();
        send(correlator, first);
        send(correlator, second);

        correlator.failAll(new IllegalStateException("closed"));

        assertTrue(first.isCompletedExceptionally());
        assertTrue(second.isCompletedExceptionally());
        assertEquals(0, correlator.pendingCount());
    }
}
