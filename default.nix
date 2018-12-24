with import <nixpkgs> {};

buildGoPackage {
  name = "procon";

  goPackagePath = "github.com/geistesk/procon";

  src = lib.cleanSource ./.;
  goDeps = ./deps.nix;
}
