# Package
version       = "0.1.0"
author        = "Web4Application"
description   = "Layer 1 Fadaka Blockchain"
license       = "MIT"
srcDir        = "src"
bin           = @["fadaka_node"]

# Dependencies
requires "nim >= 2.0.0"
requires "nimcrypto >= 0.5.4" # High-performance hashing
requires "chronos >= 3.0.0"   # Async P2P networking
