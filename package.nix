{
  buildGoModule,
  swaybg
}:
buildGoModule {
  name = "gomultilineformatter";
  src = ./.;
  vendorHash = null;

  buildInputs = [ swaybg ];
}
