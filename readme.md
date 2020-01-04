# 这是珍上完整的微服务的demo  
## 其中，  
   unit是单个的微服务，用于demo    
   gateway为微服务的测试网关，这不是每一个微服务必须实现的内容   
   unit/prog下为微服务可执行程序。  
   gateway下为网关可执行程序。  
   这个demo是一个独立的工程。  
   独立部署时，需要在demo目录下执行 go init mod
   獨立部署時，導入路徑 需要更改為：
 	//单独运行时导入改为这个
 	// or import "github.com/weihaoranW/vchat"
＃ 除了intf包外，該微服務沒有其它任何外部導出，    
   不需要暴露任何其它源代码级的共享。

＃ 關於測試變量的設置
在開發狀態下，config.yml可以指定一個固定的路徑，通過在環境變量中設置來達到效果：
此功能特别适用于单元测试或本地运行期的测试。
步骤：     
vim /etc/profile
export vchat_yml_path="/home/test/""
或更改 .bashrc变量(linux)。  
 ~/.bash_profile（maxOS）
验证变量是否设置成功：
echo $vchat_yml_path
要正常显示设置的路径 。

