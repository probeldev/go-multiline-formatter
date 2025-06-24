{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
      gomultilineformatter-package = pkgs.callPackage ./package.nix {};
    in {
      packages = rec {
        gomultilineformatter = gomultilineformatter-package;
        default = gomultilineformatter;
      };

      apps = rec {
        gomultilineformatter = flake-utils.lib.mkApp {
          drv = self.packages.${system}.gomultilineformatter;
        };
        default = gomultilineformatter;
      };

      devShells.default = pkgs.mkShell {
        packages = (with pkgs; [
          go
        ]);
      };
    });
}
