{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs?ref=nixos-unstable";
    # Pinned to gotk4's nixos-26.05 commit before WebKitGTK 2.54 and libsoup 3.8
    # were backported to nixos-26.05.
    nixpkgs-gotk4.url = "github:NixOS/nixpkgs/21ea275a7c46aef9d4d6ddc962e6d562e9d94183";
    gotk4-nix.url = "github:diamondburned/gotk4-nix/main";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      nixpkgs-gotk4,
      gotk4-nix,
      flake-utils,
    }:

    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        pkgs-gotk4 = import nixpkgs-gotk4 {
          inherit system;
          overlays = [ gotk4-nix.overlays.patchelf ];
        };
      in
      {
        devShells.default = gotk4-nix.lib.mkShell {
          base.pname = "gotk4-webkitgtk";
          pkgs = pkgs-gotk4;

          buildInputs = with pkgs-gotk4; [
            gobject-introspection
            glib
            graphene
            gdk-pixbuf
            gtk4
            gtk3
            vulkan-headers
            webkitgtk_4_1
            webkitgtk_6_0
            libsoup_3
          ];

          packages = [ self.formatter.${system} ];
        };

        formatter = pkgs.nixfmt;
      }
    );
}
