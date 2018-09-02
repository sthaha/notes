package metaprogramming

import java.text.NumberFormat;


class CurrencyCategory {
  static String asCurrency(Number n) {
    NumberFormat.currencyInstance.format(n)
  }

  static String asCurrency(Number n, Locale l) {
    NumberFormat.getCurrencyInstance(l).format(n)
  }
}


use(CurrencyCategory) {
  def amount = 1234543.74
  println amount.asCurrency()
  assert amount.asCurrency() == "\$1,234,543.74"

  println amount.asCurrency(Locale.FRANCE)
  assert amount.asCurrency(Locale.FRANCE) == "1 234 543,74 €"
}
