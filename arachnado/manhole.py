# -*- coding: utf-8 -*-
"""
An interactive Python interpreter available through telnet.
"""
from __future__ import absolute_import
from twisted.conch.telnet import TelnetTransport, TelnetBootstrapProtocol
from twisted.internet import protocol


def start(port=None, host=None, telnet_vars=None):
    from twisted.conch.manhole import ColoredManhole
    from twisted.conch.insults import insults
    from twisted.internet import reactor

    port = int(port) if port else 6023
    host = host or '127.0.0.1'
    telnet_vars = telnet_vars or {}

    factory = protocol.ServerFactory()
    factory.protocol = lambda: TelnetTransport(
        TelnetBootstrapProtocol,
        insults.ServerProtocol,
        ColoredManhole,
        telnet_vars
    )
    return reactor.listenTCP(port, factory, interface=host)
