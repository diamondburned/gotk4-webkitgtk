{
	fetchFromGitHub ? (import <nixpkgs> {}).fetchFromGitHub,

	gotk4-nix ? fetchFromGitHub {
		owner  = "diamondburned";
		repo   = "gotk4-nix";
		rev    = "ad91dabf706946c4380d0a105f0937e4e8ffd75f";
		sha256 = "0rkw9k98qy7ifwypkh2fqhdn7y2qphy2f8xjisj0cyp5pjja62im";
	},

	nixpkgs ? fetchFromGitHub {
		owner = "NixOS";
		repo  = "nixpkgs";
		rev   = "c1be43e8e837b8dbee2b3665a007e761680f0c3d"; # nixos-23.11
		hash  = "sha256-C36QmoJd5tdQ5R9MC1jM7fBkZW9zBUqbUCsgwS6j4QU=";
	},
}:

import "${gotk4-nix}/shell.nix" {
	base = {
		pname = "gotk4-webkit";
		version = "dev";
		buildInputs = pkgs: with pkgs; [
			webkitgtk_4_1
			webkitgtk_6_0
			libsoup
			libsoup_3
		];
	};
	pkgs = import "${gotk4-nix}/pkgs.nix" {
		sourceNixpkgs = nixpkgs;
		useFetched = true;
	};
}
