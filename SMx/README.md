# 概述
本章介绍的SMx系列算法，是中国国家商用密码管理办公室制定的一批密码标准，包含：

- SM1：对称加密算法，加密强度为128位，采用硬件实现。
- SM2：公钥加密算法，加密强度位256位。
- SM3：密码杂凑算法，杂凑长度为32字节——随SM2同期公布
- SM4：对称加密算法，加密强度位128位，随WAPI标准一起公布。
- SM7
- SM9

其中SM1和SM9算法未公开，需要调用严格管控的密码芯片接口来加密解密。

此外，还有ZUC（祖冲之算法）一并在本章讲解。

## SM2
- 数字签名算法
  - 摘要密文生成
  - 摘要密文验证
- 密钥交换算法
  - 会话密钥共享
- 公钥加密算法
  - 生成只能由唯一对应的私钥解密的密文

### 操作过程（Alice和Bob）
Alice（加密过程）:

- 用单向算法（FunctionAbstract）计算待发送的明文（Message）特征（用于完整性验证的指纹），得到指纹数据（FingerprintA）
- 用自己的私钥（AKey）加密数据的特征（签名）（Signature），得到带签名的指纹数据（SignedFingerprintA）
- 拼接明文和带签名的指纹数据，得到组合序列（MessageWithSignedFingerprintA）
- 用密文密码（AK）加密改组合序列，得到密文（Ciphertext=MessageWithSignedFingerprintAEncryptedByAK）
- 用Bob的公钥（BPubKey）加密密文密码（AK），得到只能由Bob的私钥（BKey）解密的密文密码（AKEncryptedByBPubKey）
- 将“Ciphertext”和“AKEncryptedByBPubKey”传送给对方（Bob）

Bob（解密过程）:

- 用自己的私钥（BKey）解密“AKEncryptedByBPubKey”得到“AK”
- 用“AK”解密“Ciphertext”得到“MessageWithSignedFingerprintA”
- 进而得到“Message”和“SignedFingerprintA”（根据拼接满足的某种规则直接拆分）
- 用单向算法（FunctionAbstract）计算接收到的明文“Message”特征，得到指纹数据（FingerprintB）
- 用Alice的公钥（APubKey）解密“SignedFingerprintA”得到“FingerprintA”（可解则说明是Alice传送的）
- 对比“FingerprintA”和“FingerprintB”（相同则说明明文无更改）

特点就是：同时实现数据加密、数据完整性验证、身份确认。

# FAQ

## SM2
- SM2和RSA的关系
  - SM2无国际化标准
- SM2较RSA的优势
- SM2和ECC的关系
- SM2签名支持多大数据、签名结果多少字节
  - 原始数据长度不限、签名结果是64字节
- SM2加密支持多大数据、加密结果多少字节
  - 支持128G字节数据加密、加密结果增加96个字节
