// Copyright (c) of parts are held by the various contributors (see the CLA)
// Licensed under the MIT License. See LICENSE file in the project root for full license information.
package polling_test

import (
	"testing"
)

// Need an api key to run this
//func TestActualOpenExchangeRatesPeggedAssets(t *testing.T) {
//	ActualDataSourceTest(t, "OpenExchangeRates")
//}

// TestFixedOpenExchangeRatesPeggedAssets tests all the crypto assets are found on OpenExchangeRates from fixed
func TestFixedOpenExchangeRatesPeggedAssets(t *testing.T) {
	FixedDataSourceTest(t, "OpenExchangeRates", []byte(openExchangeRateResponse))
}

var openExchangeRateResponse = `{
  "disclaimer": "Usage subject to terms: https://openexchangerates.org/terms",
  "license": "https://openexchangerates.org/license",
  "timestamp": 1563829200,
  "base": "USD",
  "rates": {
    "AED": 3.672943,
    "AFN": 80.065265,
    "ALL": 108.962898,
    "AMD": 476.478003,
    "ANG": 1.785062,
    "AOA": 346.5555,
    "ARS": 42.453308,
    "AUD": 1.421393,
    "AWG": 1.799748,
    "AZN": 1.7025,
    "BAM": 1.74291,
    "BBD": 2,
    "BDT": 84.503274,
    "BGN": 1.744972,
    "BHD": 0.377012,
    "BIF": 1843.896028,
    "BMD": 1,
    "BND": 1.350649,
    "BOB": 6.913577,
    "BRL": 3.740078,
    "BSD": 1,
    "BTC": 0.000096895686,
    "BTN": 68.903371,
    "BWP": 10.576381,
    "BYN": 2.01739,
    "BZD": 2.015772,
    "CAD": 1.311914,
    "CDF": 1658.54141,
    "CHF": 0.982004,
    "CLF": 0.02448,
    "CLP": 689.201477,
    "CNH": 6.883993,
    "CNY": 6.881111,
    "COP": 3179.574354,
    "CRC": 574.735159,
    "CUC": 1,
    "CUP": 25.75,
    "CVE": 98.2225,
    "CZK": 22.786418,
    "DJF": 178.05,
    "DKK": 6.660568,
    "DOP": 50.952397,
    "DZD": 119.259478,
    "EGP": 16.618853,
    "ERN": 14.999544,
    "ETB": 28.915834,
    "EUR": 0.892115,
    "FJD": 2.125946,
    "FKP": 0.801524,
    "GBP": 0.801524,
    "GEL": 2.875,
    "GGP": 0.801524,
    "GHS": 5.381009,
    "GIP": 0.801524,
    "GMD": 49.9575,
    "GNF": 9205.05413,
    "GTQ": 7.665431,
    "GYD": 208.82472,
    "HKD": 7.810553,
    "HNL": 24.630646,
    "HRK": 6.592365,
    "HTG": 94.060581,
    "HUF": 290.25871,
    "IDR": 13939.664278,
    "ILS": 3.527525,
    "IMP": 0.801524,
    "INR": 68.930029,
    "IQD": 1191.951828,
    "IRR": 42105,
    "ISK": 124.797205,
    "JEP": 0.801524,
    "JMD": 134.725012,
    "JOD": 0.707676,
    "JPY": 107.87131873,
    "KES": 103.595882,
    "KGS": 69.568994,
    "KHR": 4081.938206,
    "KMF": 438.799462,
    "KPW": 900,
    "KRW": 1176.751157,
    "KWD": 0.304254,
    "KYD": 0.833403,
    "KZT": 384.09491,
    "LAK": 8714.870023,
    "LBP": 1510.928748,
    "LKR": 175.927983,
    "LRD": 201.624911,
    "LSL": 13.877692,
    "LYD": 1.400829,
    "MAD": 9.599364,
    "MDL": 17.511561,
    "MGA": 3634.61443,
    "MKD": 54.664077,
    "MMK": 1518.09185,
    "MNT": 2663.205287,
    "MOP": 8.044468,
    "MRO": 357,
    "MRU": 36.75,
    "MUR": 35.900182,
    "MVR": 15.449981,
    "MWK": 760.059801,
    "MXN": 19.05685,
    "MYR": 4.112534,
    "MZN": 61.925697,
    "NAD": 13.877692,
    "NGN": 360.05481,
    "NIO": 33.194843,
    "NOK": 8.607838,
    "NPR": 110.24481,
    "NZD": 1.479409,
    "OMR": 0.38499,
    "PAB": 1,
    "PEN": 3.284944,
    "PGK": 3.394778,
    "PHP": 51.115152,
    "PKR": 160.085,
    "PLN": 3.788393,
    "PYG": 5977.559651,
    "QAR": 3.641584,
    "RON": 4.207965,
    "RSD": 104.9985,
    "RUB": 63.073406,
    "RWF": 909.291013,
    "SAR": 3.750633,
    "SBD": 8.271367,
    "SCR": 13.742548,
    "SDG": 45.114342,
    "SEK": 9.412761,
    "SGD": 1.360854,
    "SHP": 0.801524,
    "SLL": 7114.37678,
    "SOS": 579.119666,
    "SRD": 7.458,
    "SSP": 130.26,
    "STD": 21560.79,
    "STN": 21.89,
    "SVC": 8.750088,
    "SYP": 514.995043,
    "SZL": 13.885269,
    "THB": 30.847337,
    "TJS": 9.430295,
    "TMT": 3.499986,
    "TND": 2.863737,
    "TOP": 2.267915,
    "TRY": 5.679025,
    "TTD": 6.771598,
    "TWD": 31.070084,
    "TZS": 2299.181418,
    "UAH": 25.676969,
    "UGX": 3695.187533,
    "USD": 1,
    "UYU": 34.870661,
    "UZS": 8582.336101,
    "VEF": 248487.642241,
    "VES": 7487.526397,
    "VND": 23269.350097,
    "VUV": 114.960514,
    "WST": 2.610054,
    "XAF": 585.189078,
    "XAG": 0.06112848,
    "XAU": 0.00070188,
    "XCD": 2.70255,
    "XDR": 0.723757,
    "XOF": 585.189078,
    "XPD": 0.00065314,
    "XPF": 106.457637,
    "XPT": 0.00117847,
    "YER": 250.300682,
    "ZAR": 13.863145,
    "ZMW": 12.825468,
    "ZWL": 322.000001
  }
}
`
