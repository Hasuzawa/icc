## Introduction

Experiments and utility functions for the ICC Specification v4.

## Apple ColorSync Utility

<img src="./image/ColorSync_Utility_icon.png" alt="icon of ColorSync Utility" width=128>

For Mac, the official application ColorSync Utility can registers different color profiles and transforms color between different color model.

<img src="./image/ColorSync_Utility_screenshot.png" alt="screen of ColorSync Utility">

## Specification

[Specification ICC.1:2022 (Profile version 4.4.0.0)](https://www.color.org/specification/ICC.1-2022-05.pdf)

## Documentation

[International Color Consortium](https://www.color.org/)

[ColorSync Utility](https://support.apple.com/guide/colorsync-utility/welcome/mac)

[Color profiles | Adobe](https://helpx.adobe.com/acrobat/using/color-profiles.html)

## Reference

[Manufacturer Registry](https://www.color.org/signatureRegistry/index.xalter)

## Bugs Known

These are bugs or discrepancies I discovered when I am developing the software.

1. In the official Private and ICC Tag and CMM Registry dated 4 March 2021, the CMM Signature for Konica Minolta is listed as "MCML" with ASCII of 0x4D43_4D44, but the ASCII value of "L" is 0x4C, the correct signature for "MCML" should be 0x4D43_4D4C. For backward compatibility, both values are decoded to "Konica Minolta".
