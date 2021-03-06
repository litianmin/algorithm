位运算
==========================


盖世神功口诀
------------------------
> 清零取反要用与，某位置一可用或  
> 若要取反和交换，轻轻松松用异或


移位运算
------------------------
> - `<<`  左移，右边空出的位置补上 0，左边的位将从字头挤掉，其值相当于乘2  
> - `>>` 右移：右边的位被挤掉。对于左边移出的空位，如果是正数则空位补0，若为负数，可能补0或补1，这取决于所用的计算机系统  
> - 运算符，右边的位被挤掉，对于左边移出的空位一概补上0


运算符表
| 含义                        | 运算符 | 例子 |
| :---                         | :---:   | :---  |
| 左移                        | <<    |  0011 => 0110 |
| 右移                        | >>    |  0011 => 0110 |
| 按位或                      |  \|    |  0011 \| 1011 => 1011 |
| 按位与                      | &     |  0011 & 1011 => 0011 |
| 按位取反                    | ~     |  ~0011 => 1100 |
| 按位异或 (相同为零不同为一)   | ^     |  0011 ^ 1011 => 1000 |


按位与运算符 （&）
-------------------------
> - 运算规则  
>  0&0=0、0&1=0、1&0=0、1&1=1  

> - 注意  
> 负数按补码形式参与按位与运算  

> - 用途
>> - 清零  
>> 如果想将一个单元清零，即使其全部二进制位为0，只要与一个各位都为零的数值相与，结果为零  
>> - 取一个数的指定位  
>> 比如取数 X=1010 1110 的低4位，只需要另找一个数Y，令Y的低4位为1，其余位为0，即Y=0000 1111，然后将X与Y进行按位与运算（X&Y=0000 1110）即可得到X的指定位。  
>> - 判断奇偶  
>> 只要根据最未位是0还是1来决定，为0就是偶数，为1就是奇数。因此可以用if ((a & 1) == 0)代替if (a % 2 == 0)来判断a是不是偶数   
>> - 取余数  
>> key % M = key & （M-1)


按位或运算符（|） 
--------------------------
> - 运算规则  
> 0|0=0  0|1=1  1|0=1  1|1=1  

> - 用途  
>> - 常用来对一个数据的某些位设置为1  
>> 比如将数 X=1010 1110 的低4位设置为1，只需要另找一个数Y，令Y的低4位为1，其余位为0，即Y=0000 1111，然后将X与Y进行按位或运算（X|Y=1010 1111）即可得到。  

异或运算符（^）
--------------------------
> - 理解  
> 异或也叫半加运算，其运算法则相当于不带进位的二进制加法  

> - 性质  
>> - 交换律  
>> - 结合律 (a^b)^c == a^(b^c)  
>> - 对于任何数x，都有 x^x=0，x^0=x  
>> - 自反性: a^b^b=a^0=a;  

> - 用途  
>> - 翻转指定位  
>> 比如将数 X=1010 1110 的低4位进行翻转，只需要另找一个数Y，令Y的低4位为1，其余位为0，即Y=0000 1111，然后将X与Y进行异或运算（X^Y=1010 0001）即可得到。  
>> - 与0相异或值不变  
>> 例如：1010 1110 ^ 0000 0000 = 1010 1110  
>> - 交换两个数  
>> `a ^= b;  b ^= a; a ^= b;`  


