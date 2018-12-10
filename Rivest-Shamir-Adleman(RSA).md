公钥（非对称）加密概念

# RSA

- 1977
- Ron Rivest(MD5)、Adi Shamir、Leonard Adleman
- MIT

允许：

- 不必独立分发给参与各方密钥
- 不必验证消息来源

## 基本概念

- 有公钥和私钥
- 不能根据公钥计算出私钥
  - 反之亦然

```
Public-A Private-A for Alice
Public-B Private-B for Bob

明文：M

给Bob的密文：Cb = Public-B(M)
M = Private-B(Cb)
```

发给谁，就用谁的公开密钥加密，只有拥有对应的私钥的接收者才能解密。