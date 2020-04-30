## 根据设计文档生成实体工具

#### 参数
+ “-s” 源文件路径，默认值“./source.xlsx”
+ “-t” 输出目标目录，默认值“./output”
+ “-l” 语言类型，目前支持Go与C#，默认值“CS”

#### 使用
使用powershell到执行文件目录执行

###### 默认会在执行文件同级下找到source文件，然后执行
```powershell
.\apidoc2entity.exe
```

###### 自定义参数启动
```powershell
.\apidoc2entity.exe -s "./source.xlsx" -t "./output" -l "GO"
```


