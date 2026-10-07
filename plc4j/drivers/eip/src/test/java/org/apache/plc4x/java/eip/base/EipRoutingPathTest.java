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
package org.apache.plc4x.java.eip.base;

import org.apache.plc4x.java.eip.readwrite.*;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

/**
 * The route a connection builds from {@code communication-path}. The PLC4X docs write the path
 * in brackets ({@code communication-path=[1,4,2,192.168.0.1,1,1]}) and the value reaches the
 * driver as written, so the brackets must not end up in the first and last element - that drops
 * every hop and the Forward_Open goes to the Ethernet module itself.
 */
class EipRoutingPathTest {

    private static final int SLOT = 3;

    @Test
    void bracketedPathIsRouted() {
        assertEquals(List.of("port 1/0", "class 2", "instance 1"), route("[1,0]"));
    }

    @Test
    void barePathIsRouted() {
        assertEquals(List.of("port 1/0", "class 2", "instance 1"), route("1,0"));
    }

    @Test
    void spacesAroundBracketsAndElementsAreIgnored() {
        assertEquals(List.of("port 1/4", "class 2", "instance 1"), route("  [ 1 , 4 ] "));
        assertEquals(List.of("port 1/4", "class 2", "instance 1"), route("1, 4"));
    }

    @Test
    void multiHopPathWithAnIpAddressIsRouted() {
        List<String> expected = List.of("port 1/4", "port 2/192.168.1.20", "port 1/0", "class 2", "instance 1");
        assertEquals(expected, route("1,4,2,192.168.1.20,1,0"));
        assertEquals(expected, route("[1,4,2,192.168.1.20,1,0]"));
    }

    /** The documented example: an odd-length IP address is padded to an even byte count. */
    @Test
    void oddLengthIpAddressIsPadded() {
        List<PathSegment> segments = EipTcpConnection.routingPath("[1,4,2,192.168.0.1,1,1]", SLOT);
        PortSegmentExtended ip = (PortSegmentExtended) ((PortSegment) segments.get(1)).getSegmentType();
        assertEquals(2, ip.getPort());
        assertEquals(11, ip.getLinkAddressSize());
        assertEquals("192.168.0.1\0", ip.getAddress());
        assertEquals(List.of("port 1/4", "port 2/192.168.0.1", "port 1/1", "class 2", "instance 1"),
            route("[1,4,2,192.168.0.1,1,1]"));
    }

    /** No path, an empty one and empty brackets all route to the configured slot on port 1. */
    @Test
    void emptyPathFallsBackToTheSlot() {
        List<String> slotRoute = List.of("port 1/" + SLOT, "class 2", "instance 1");
        assertEquals(slotRoute, route(null));
        assertEquals(slotRoute, route(""));
        assertEquals(slotRoute, route("[]"));
        assertEquals(slotRoute, route(" [ ] "));
    }

    /** Unchanged: an odd number of elements is not a route, and an unknown port is dropped. */
    @Test
    void malformedPathsKeepTheirBehaviour() {
        assertEquals(List.of("class 2", "instance 1"), route("[1,0,1]"));
        assertEquals(List.of("port 1/0", "class 2", "instance 1"), route("[3,7,1,0]"));
    }

    private static List<String> route(String communicationPath) {
        return EipTcpConnection.routingPath(communicationPath, SLOT).stream()
            .map(EipRoutingPathTest::describe)
            .toList();
    }

    private static String describe(PathSegment segment) {
        if (segment instanceof PortSegment port) {
            if (port.getSegmentType() instanceof PortSegmentNormal normal) {
                return "port " + normal.getPort() + "/" + normal.getLinkAddress();
            }
            PortSegmentExtended extended = (PortSegmentExtended) port.getSegmentType();
            return "port " + extended.getPort() + "/"
                + extended.getAddress().substring(0, extended.getLinkAddressSize());
        }
        LogicalSegmentType logical = ((LogicalSegment) segment).getSegmentType();
        if (logical instanceof ClassID classId) {
            return "class " + classId.getSegmentClass();
        }
        if (logical instanceof InstanceID instanceId) {
            return "instance " + instanceId.getInstance();
        }
        return segment.toString();
    }
}
