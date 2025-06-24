{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
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
      goaddtag-package = pkgs.callPackage ./package.nix {};
    in {
      packages = rec {
        goaddtag = goaddtag-package;
        default = goaddtag;
      };

      apps = rec {
        goaddtag = flake-utils.lib.mkApp {
          drv = self.packages.${system}.goaddtag;
        };
        default = goaddtag;
      };

      devShells.default = pkgs.mkShell {
        packages = (with pkgs; [
          go
        ]);
      };
    });
}
