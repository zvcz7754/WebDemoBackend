go語言實作一個各種主題的網站
以前後端分離這只有後端部分


框架:gin
資料庫:

cmd: 
router: 與socket對應/呼叫service函式/
service: 被呼叫後決定該如何呼叫repo
repo: ORM資料庫/不同service呼叫不同repo方法
entity: 接收資料庫回傳的資料

待辦:
body/josn的應對(在router)，要放哪裡
進入postBook要是結構還是slice
有可能一個有可能多個那我是不是預設多個

從資料庫來的資料出來都是struct要怎變成HTTP的body


->如果像get方法是參數放在URL就用傳參數，如果是放成body/json就用struct





Repo=不同qurey對應不同接口
router=URL對應啟動
servive=router呼叫的終點/匯入repo
