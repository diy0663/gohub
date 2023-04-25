可视化:
参考: https://www.modb.pro/db/455206
1.安装:   embedded-struct-visualizer  
go install github.com/davidschlachter/embedded-struct-visualizer@latest

2.之后需要把gopath的bin加入到~/.zshrc里面
export PATH=$PATH:$(go env GOPATH)/bin

3. brew update

4.  brew install graphviz

5.项目根目录运行: embedded-struct-visualizer -out demo.gv 
生成了   demo.gv 文件

6.图形化展示 (生成demo.ps文件)
dot -Tps demo.gv -o demo.ps