class MainPage {
  late String text;
  late String logo; // todo should be image
  late String? largeLogo; // todo should be image
  late List<String>? features;

  MainPage({
    required this.text,
    required this.logo,
    this.largeLogo,
    this.features,
  });
}