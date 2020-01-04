app=g
ext=.out
path=./
platform=$(P)
generate=$(g)
flag='all=-w -s'
export GO111MODULE=on

.PHONY:build,start,stop
build:
#	编译前先进行 go generate
ifeq ($(platform), linux)
	@CGO_ENABLED=0 GOOS="linux" GOARCH="amd64" go build -ldflags $(flag) -o $(path)$(app)_linux$(ext) ./
else
	@go build -ldflags $(flag) -o $(path)$(app)_darwin$(ext) ./
endif
start:
	./yh2_darwin.out restart && ./yh2_darwin.out restart -f worker
stop:
	./yh2_darwin.out stop && ./yh2_darwin.out stop -f worker
.PHONY:upx
upx:
ifeq ($(platform), mac)
	@upx -9 $(path)$(app)_darwin$(ext)
else ifeq ($(platform), n)
	@upx -9 $(path)$(app)*$(ext)
else
	@upx -9 $(path)$(app)_linux$(ext)
endif
.PHONY:build_race
build_race:
ifeq ($(platform), linux)
#需要安装 xgo
# docker pull karalabe/xgo-latest
# go get github.com/karalabe/xgo
# 编译之前 需要运行 go mod vendor
#	@xgo -go 1.12 -race -ldflags $(flag) --targets=linux/amd64 -out $(path)$(app) ./
#	@upx -9 *-linux-amd64
#不能用

else
	@go build -race -ldflags $(flag) -o $(path)$(app)_darwin$(ext) ./
endif
.PHONY:clean
clean:
	@rm -rf logs *$(ext) *.log

topicPath=$(p)
local=$(l)
.PHONY:sub
sub:
ifeq ($(local),1)
	mosquitto_sub -p 21883 -q 1 -h 127.0.0.1 -t "yh_root/$(topicPath)"
endif
	mosquitto_sub -p 1883 -q 1 -h 47.52.104.202 -t "yh_root/$(topicPath)"

# .PHONY:scpt
#scpt:
#	scp -P 39822 -i ~/.ssh/yhxilie.pem yh2_linux.out root@47.52.104.202:/opt/yh2
