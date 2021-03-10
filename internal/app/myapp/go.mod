module myapp

go 1.14

require github.com/astaxie/beego v1.12.3
require github.com/jameszhangcn/redisApi v0.0.0-20210222

replace github.com/jameszhangcn/redisApi => ../../pkg/myprivlib/redis