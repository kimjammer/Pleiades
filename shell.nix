{ pkgs ? import <nixpkgs> {} }:

let
  lib = import <nixpkgs/lib>;

  NPM_CONFIG_PREFIX = toString ./npm_config_prefix;

in pkgs.mkShell {
  packages = with pkgs; [
    nodejs_23
    nodePackages.prettier
    typescript-language-server
    svelte-language-server
  ];

  inherit NPM_CONFIG_PREFIX;

  shellHook = ''
    export PATH="${NPM_CONFIG_PREFIX}/bin:$PATH"
  '';
}
