{
  description = "finnhub-cli — Go CLI client for Finnhub API";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "finnhub-cli";
          version = "0.1.0";
          src = ./.;
          vendorHash = null; # update after first build
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            goreleaser
          ];
        };
      });
}
