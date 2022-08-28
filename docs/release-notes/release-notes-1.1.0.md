janeta version 1.1.0 is now available from:

  https://github.com/janotchain/janeta/releases/tag/v1.1.0


Please report bugs using the issue tracker at github:

  https://github.com/janotchain/janeta/issues

How to Upgrade
===============

Please notice new version janeta path is $GOPATH/src/github.com/janotchain/janeta if you install janeta from source code.  
If you are running an older version, shut it down. Wait until it has quited completely, and then run the new version janeta.
You can operate according to the user manual.[(janeta User Manual)](https://janeta.io/wp-content/themes/freddo/images/wallet/janetaUsermanualV1.0_en.pdf)


1.1.0 changelog
================
__janeta Node__

+ [`PR #1805`](https://github.com/janotchain/janeta/pull/1805)
    - Correct janeta go import path to github.com/janotchain/janeta. Developer can use go module to manage dependency of janeta. 
+ [`PR #1815`](https://github.com/janotchain/janeta/pull/1815) 
    - Add asynchronous validate transactions function to optimize the performance of validating and saving block. 

__janeta Dashboard__

+ [`PR #1829`](https://github.com/janotchain/janeta/pull/1829) 
    - Fixed the decimals type string to integer in create asset page.

Credits
--------

Thanks to everyone who directly contributed to this release:

- DeKaiju
- iczc
- Paladz
- zcc0721
- ZhitingLin

And everyone who helped test.
