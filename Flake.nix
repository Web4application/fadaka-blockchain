{
  description = "Fadaka Blockchain Development Environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.nim        # The Nim compiler
            pkgs.nimble     # Nim package manager
            pkgs.openssl    # Required for blockchain networking/crypto
          ];

          shellHook = ''
            echo "--- Fadaka Blockchain Dev Environment Loaded ---"
            nim --version
          '';
        };
      });
}
