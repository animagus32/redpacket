# 实现一个简易的口令红包功能。

## 功能：

1.  用户可以发"口令红包"。发的红包不从余额内扣除。（假设用户都是通过银行卡发的红包）
2.  用户可以抢"口令红包"。每个用户对于每一个红包仅能抢一次。抢到的红包计算在余额内。
3.  "口令红包"可以指定发多少钱，以及红包的个数。每个红包的金额则是随机的。只要知道口令，无论谁都可以抢。
4.  口令由系统自动生成，仅包含大小写字母和数字，8位
5.  未抢完的红包在（发红包的）24小时后自动退回。
6.  用户可以查看自己的余额。
7.  用户可以查看自己抢到的红包列表。


## 接口
1. 注册 POST: api/v1/user/register  
    参数：
     * username 用户名
     * password 前端传md5后的密码
     
2. 登录 POST: api/v1/user/login  
    参数：
     * username 用户名
     * password 前端传md5后的密码

3. 获取余额 GET: api/v1/user/1/balance
   参数:
    * header: uid 和 token
 
4. 创建红包 POST: api/v1/redpacket/dispatch
    参数：
    * uid 
    * amount 
    * num
   
5. 抓红包  POST: api/v1/redpacket/grab
    参数：
     * secret
      
6. 获取红包  GET: api/v1/redpacket/list?uid=1
   参数：
   * uid
   

## 目前还剩功能 
 * 测试