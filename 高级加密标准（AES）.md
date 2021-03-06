（美国）国家标准和技术协会（NIST）于1997年1月宣布创建新的密码算法的项目需求。

许多组织提议了各种密码算法。

各种算法通过在速度和安全系数的检验，在若干轮的学习和考核后，NIST最终选择叫作Rijndael的算法。

Rijndael因其安全、成本、弹性、完整性和算法监控的优点被选为最佳，因此，NIST选择Rijndael作为高级加密标准（AES），在2000年10月。

在2001年11月26日，AES成为一个FIPS（联邦信息处理标准）。

AES指定一个FIPS核准的密码学算法用于安全点数据。

美国政府（NSA）在2003年6月接收并宣布AES足够保证机密信息乃至绝密等级的安全。

Rijndael的命名源自两位比利时的密码学家，[Joan Daemen](https://en.wikipedia.org/wiki/Joan_Daemen)博士和[Vincent Rijmen](https://en.wikipedia.org/wiki/Vincent_Rijmen)博士，在“[比利时天主教鲁汶大学](https://www.kuleuven.be/kuleuven/)”电子工程系工作期间提出。

[Rijndael（AES）](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)是免专利费的，并且创造者还在公共领域给出各种参考实现。

## AES加密解密

[CODE](./AES.go)

## AES的高层设计

AES包含基于“[伽罗瓦域（Golois Field）](https://baike.baidu.com/item/%E6%9C%89%E9%99%90%E5%9F%9F/4273049)”的算术操作，有`GF(2^N)`结构，这里`N`为`8`。

AES是对称加密。

每次用128比特密钥把128比特明文加密为128比特的密文。

AES需要多余128比特实现加密和解密过程。

AEC的128比特加密有10轮（单步骤集的替代和组合网络设计）操作来执行加密和解密过程。

根据密钥的类型和轮数多少，有三个版本：AES-128（10轮）、AES-192（12轮）、AES-256（14轮）。

AES全部数据块，在每轮用完全相同的方式处理。

在AES中，明文经历`N*r`轮后变成密文。

再说，每轮由四个不同操作构成。

一个操作是组合，另外三个操作是替换。

有`SubBytes`、`ShiftRows`、`MixColumns`、`AddRoundKey`。

在AES中，加密过程用到的全部转换都将有逆转换，用于解密过程。

在AES中，解密的每轮就是用逆转换：`InvSubBytes`、`InvShiftRows`、`InvMixColumns`。

## 用AES强加密

AES在DES之后产出的，所有被识别的DES攻击，都在AES上做了演示，所有结果都是有效的。

相比DES，AES更有信心面对暴力攻击（brute-force），因为AES更大的可变密钥尺寸、块尺寸。

AES不易受到统计攻击的影响，并且已经测试过，常用技术给AES的密文做统计分析是无法实现的。

当今，没有特异的、线性的、成功的AES攻击被检测到。

AES最好的部分是所用的算法如此基础，以至于可以快速实现，用廉价的处理器和少量内存即可。

相比DES，AES需要更高处理和更多轮转换，这是能说的AES的“相对缺点”。