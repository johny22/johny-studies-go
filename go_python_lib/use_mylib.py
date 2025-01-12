from ctypes import *

shared_lib = CDLL('./mylib.so')

name = input('What\' your name? ')
print('Hi, ', name)
