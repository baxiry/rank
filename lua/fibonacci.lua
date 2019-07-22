#!/usr/bin/env luajit

-- or just lua without jit

local to = arg[1] and tonumber(arg[1]) or 10
local a,b = 0,1

for i=1, to do
    print(i, b)
    a,b = b, a+b
end
