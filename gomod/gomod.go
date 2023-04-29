go mod init 生成go.mod文件
go mod download 下载go.mod文件中指明所有依赖
go mod tidy 整理现有依赖
go mod graph 查看现有依赖结构
go mod edit 编辑go.mod文件 
go mod vendor 导出项目所有的依赖到vendor目录
go mod verify 校验一个模块是否被篡改
go mod why 查看为什么要依赖某模块
