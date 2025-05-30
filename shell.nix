{ pkgs ? import <nixpkgs> {
  config.allowUnfree = true;
} }:

let
  lib = import <nixpkgs/lib>;

  NPM_CONFIG_PREFIX = toString ./npm_config_prefix;

in pkgs.mkShell {
  packages = with pkgs; [
    nodejs_23
    nodePackages.prettier
    typescript-language-server
    svelte-language-server
    go
    gopls
    mongodb-ce
    mongosh
    playwright-driver.browsers
  ];

  inherit NPM_CONFIG_PREFIX;

  shellHook = ''
    export PATH="${NPM_CONFIG_PREFIX}/bin:$PATH"
    export PLAYWRIGHT_BROWSERS_PATH=${pkgs.playwright-driver.browsers}

    if ! [[ -d ./database ]]; then
      mkdir database
    fi
  '';
}
