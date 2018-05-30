公开密钥算法可以抵挡通用户过跟踪验证的方式破解注册。

# 分类
- 内部注册机（内存注册机）：驻留内存、组织某些进程，查找正确注册序列。
- 外部注册机（算法注册机）：反汇编分析出注册算法。

## 基于ECC注册机制（一）

利用正负消元。

### 序列制作
- 选择一条椭圆曲线Ep(a,b)及基点G
- 选择私钥k（k小于G的阶n）
- 根据G计算公钥K=kG
- 产生随机整数r（r小于G的阶n）
- 计算点R=rG
- 将用户名和R坐标(x,y)作为参数，计算Hash=SHA(username,x,y)
- 计算sn=r-Hash*k (mod n)
- 将sn和Hash作为用户名的序列

### 序列验证
- 提取序列的sn和Hash
- 计算点R=sn*G+Hash*K (mod p)
  - sn*G+Hash*K
  - (r-Hash*k)*G+Hash*K
  - rG-Hash*kG+Hash*K
  - rG-Hash*K+Hash*K
  - rG
  - R
- 将用户名和R坐标(x,y)作为参数，计算H=SHA(username,x,y)
- 若H=Hash则注册成功

两个过程：
- 作者签名用到：Ep(a,b)，G，k（私钥）、r（随机数）
- 用户比对用到：Ep(a,b)，G，K（公钥）

可以保证：
- 序列是作者颁发的（签名）
- 序列是给用户名的（比对）

## 基于ECC注册机制（二）

利用因数消元。

## 签名
- R(x,y)=rG
- x'=x (mod n)
- sn=(Hash+x'*k)/r (mod n)

## 比对
- R=(Hash*G+x'*K)/sn
  - (Hash*G+x'*K)/[(Hash+x'*k)/r]
  - (Hash*G+x'*K)/[(Hash*G+x'*k*G)/rG]
  - rG*[(Hash*G+x'*K)/(Hash*G+x'k*G)]
  - rG*[(Hash*G+x'*K)/(Hash*G+x'K)]
  - rG
  - R(x,y)
