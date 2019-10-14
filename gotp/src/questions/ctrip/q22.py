#!/bin/python
# -*- coding: utf8 -*-
import sys
import os
import re

#请完成下面这个函数，实现题目要求的功能
#当然，你也可以不按照下面这个模板来作答，完全按照自己的想法来 ^-^ 
#******************************开始写代码******************************


def  CheckBlackList(userIP, blackIP):
    isBlack = 1
    if "/" not in blackIP:
        if userIP != blackIP:
            isBlack = 0
    else:
        netIp, pcIp = blackIP.split("/")
        pcIp = int(pcIp)
        ip2bin = lambda s: "".join([("%8s" % bin(int(i)).replace("0b", "")).replace(" ", "0") for i in s.split(".")])
        netIPBin = ip2bin(netIp)
        userIPBin = ip2bin(userIP)
        for i in range(pcIp):
            if netIPBin[i] != userIPBin[i]:
                isBlack = 0
                break
    
    return isBlack

#******************************结束写代码******************************


try:
    _userIP = input()
except:
    _userIP = None

try:
    _blackIP = input()
except:
    _blackIP = None

  
res=CheckBlackList(_userIP, _blackIP)

print(str(int(res)) + "\n")