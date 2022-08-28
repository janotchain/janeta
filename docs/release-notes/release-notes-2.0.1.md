janeta version 2.0.1 is now available from:

  https://github.com/janotchain/janeta/releases/tag/v2.0.1


Please report bugs using the issue tracker at github:

  https://github.com/janotchain/janeta/issues

How to Upgrade
===============

Please notice new version janeta path is $GOPATH/src/github.com/janeta/janeta if you install janeta from source code.  
If you are running an older version, shut it down. Wait until it has quited completely, and then run the new version janeta.
You can operate according to the user manual.[(janeta User Manual)](https://janeta.io/wp-content/themes/freddo/images/wallet/janetaUsermanualV1.0_en.pdf)


2.0.1 changelog
================
__janeta Node__

+ [`PR #1848`](https://github.com/janotchain/janeta/pull/1848/files)
    - Extend butxo struct with state data and add a lot of tx cases with butxo state.
+ [`PR #1851`](https://github.com/janotchain/janeta/pull/1852/files)
    - Add the new fork choice rule, which follow the chain containing the justified checkpoint of the greatest height.
+ [`PR #1854`](https://github.com/janotchain/janeta/pull/1854/files)
    - Replace uint256 with int64 when performing multiplication in the virtual machine.
+ [`PR #1855`](https://github.com/janotchain/janeta/pull/1855/files)
    - According to the number of validator mortgage and the number of votes received, a maximum of 10 validators are finally selected.
+ [`PR #1858`](https://github.com/janotchain/janeta/pull/1858/files) 
    - Add the structure of verification message, and auth whether a verification message is legal. 
+ [`PR #1865`](https://github.com/janotchain/janeta/pull/1865/files)
    - Save and delete contract registered by BCRP(janeta contract register protocol)
+ [`PR #1868`](https://github.com/janotchain/janeta/pull/1868/files) 
    - When a new block comes, consensus algorithm will update the information in the checkpoint and make it persistent.
+ [`PR #1891`](https://github.com/janotchain/janeta/pull/1891/files)
    - Convert call BCRP contract program to normal contract program
+ [`PR #1921`](https://github.com/janotchain/janeta/pull/1921/files)
    - Use 2d data to store butxo state.
+ [`PR #1930`](https://github.com/janotchain/janeta/pull/1930/files)
    - support executing state data in btm virtual machine.
+ [`PR #1935`](https://github.com/janotchain/janeta/pull/1868/files) 
    - In the new consensus algorithm, nodes take turns to generate blocks every 6 seconds.

Credits
--------

Thanks to everyone who directly contributed to this release:

- shenao78
- boomyl
- songxuexian
- DeKaiju
- Paladz

And everyone who helped test.