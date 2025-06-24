{
  buildGoModule,
  swaybg
}:
buildGoModule {
  name = "goaddtag";
  src = ./.;
  vendorHash = null;

  buildInputs = [ swaybg ];
}
