package main

import (
	"encoding/json"
    "fmt"
  //  "strconv"
  //  "strings"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type FoodChainCode struct{	
}

//food数据结构体
type ItemInfo struct{
    ItemID string `json:ItemID`                             //ItemID
    ItemProInfo ProInfo `json:ItemProInfo`                  //生产信息
    ItemIngInfo []IngInfo `json:ItemIngInfo`                //配料信息
    ItemLogInfo LogInfo `json:ItemLogInfo`                  //物流信息
}

type ItemAllInfo struct{
    ItemID string `json:ItemID`
    ItemProInfo ProInfo `json:ItemProInfo`
    ItemIngInfo []IngInfo `json:ItemIngInfo`
    ItemLogInfo []LogInfo `json:ItemLogInfo`
}

//生产信息
type ProInfo struct{
    ItemName string `json:ItemName`                         //Item名称
    ItemSpec string `json:ItemSpec`                         //Item规格
    ItemMFGDate string `json:ItemMFGDate`                   //Item日期
    ItemEXPDate string `json:ItemEXPDate`                   //Item保质期
    ItemLOT string `json:ItemLOT`                           //Item批次号
    ItemMFRID string `json:ItemMFRID`                       //Item生产商ID
    ItemLocation string `json:ItemLocation`                 //Item生产所在地
}
type IngInfo struct{
    IngID string `json:IngID`                               //配料ID
    IngName string `json:IngName`                           //配料名称
}

type LogInfo struct{
    LogDepartureTm string `json:LogDepartureTm`             //出发时间
    LogArrivalTm string `json:LogArrivalTm`                 //到达时间
    LogSender string `json:LogSender`                       //寄件方
    LogReceiver string `json:LogReceiver`                   //收件方
    LogMOT string `json:LogMOT`                             //运送方式
    LogCopID string `json:LogCopID`                         //物流公司名称
}

func (a *FoodChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (a *FoodChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()
    if fn == "addProInfo"{
        return a.addProInfo(stub,args)
    } else if fn == "addIngInfo"{
        return a.addIngInfo(stub,args)
    } else if fn == "getItemInfo"{
        return a.getItemInfo(stub,args)
    }else if fn == "addLogInfo"{
        return a.addLogInfo(stub,args)
    }else if fn == "getProInfo"{
        return a.getProInfo(stub,args)
    }else if fn == "getLogInfo"{
        return a.getLogInfo(stub,args)
    }else if fn == "getIngInfo"{
        return a.getIngInfo(stub,args)
    }else if fn == "getLogInfo_l"{
        return a.getLogInfo_l(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}

func (a *FoodChainCode) addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error 
    var ItemInfos ItemInfo

    if len(args)!=8{
        return shim.Error("Incorrect number of arguments.")
    }
    ItemInfos.ItemID = args[0]
    if ItemInfos.ItemID == ""{
        return shim.Error("ItemID can not be empty.")
    }
    
    ItemInfos.ItemProInfo.ItemName = args[1]
    ItemInfos.ItemProInfo.ItemSpec = args[2]
    ItemInfos.ItemProInfo.ItemMFGDate = args[3]
    ItemInfos.ItemProInfo.ItemEXPDate = args[4]
    ItemInfos.ItemProInfo.ItemLOT = args[5]
    ItemInfos.ItemProInfo.ItemMFRID = args[6]
    ItemInfos.ItemProInfo.ItemLocation = args[7]
    ProInfosJSONasBytes,err := json.Marshal(ItemInfos)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(ItemInfos.ItemID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

func(a *FoodChainCode) addIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
        
    var ItemInfos ItemInfo
    var IngInfoitem IngInfo

    if  (len(args)-1)%2 != 0 || len(args) == 1{
        return shim.Error("Incorrect number of arguments")
    }

    ItemID := args[0]
    for i :=1;i < len(args);{   
        IngInfoitem.IngID = args[i]
        IngInfoitem.IngName = args[i+1]
        ItemInfos.ItemIngInfo = append(ItemInfos.ItemIngInfo,IngInfoitem)
        i = i+2
    }
    
    
    ItemInfos.ItemID = ItemID
  /*  ItemInfos.ItemIngInfo = ItemIngInfo*/
    IngInfoJsonAsBytes,err := json.Marshal(ItemInfos)
    if err != nil {
    return shim.Error(err.Error())
    }

    err = stub.PutState(ItemInfos.ItemID,IngInfoJsonAsBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
        
}

func(a *FoodChainCode) addLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    var err error
    var ItemInfos ItemInfo

    if len(args)!=7{
        return shim.Error("Incorrect number of arguments.")
    }
    ItemInfos.ItemID = args[0]
    if ItemInfos.ItemID == ""{
        return shim.Error("ItemID can not be empty.")
    }
    ItemInfos.ItemLogInfo.LogDepartureTm = args[1]
    ItemInfos.ItemLogInfo.LogArrivalTm = args[2]
    ItemInfos.ItemLogInfo.LogSender = args[3]
    ItemInfos.ItemLogInfo.LogReceiver = args[4]
    ItemInfos.ItemLogInfo.LogMOT = args[5]
    ItemInfos.ItemLogInfo.LogCopID = args[6]
    
    LogInfosJSONasBytes,err := json.Marshal(ItemInfos)
    if err != nil{
        return shim.Error(err.Error())
    } 
    err = stub.PutState(ItemInfos.ItemID,LogInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}



func(a *FoodChainCode) getItemInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    ItemID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ItemID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var ItemAllinfo ItemAllInfo

    for resultsIterator.HasNext(){
        var ItemInfos ItemInfo
        response,err :=resultsIterator.Next()
        if err != nil {
             return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ItemInfos)
        if ItemInfos.ItemProInfo.ItemName !=""{
            ItemAllinfo.ItemProInfo = ItemInfos.ItemProInfo
        }else if ItemInfos.ItemIngInfo != nil{
            ItemAllinfo.ItemIngInfo = ItemInfos.ItemIngInfo
        }

    }
    
    jsonsAsBytes,err := json.Marshal(ItemAllinfo)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(jsonsAsBytes)
}
 

func(a *FoodChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    ItemID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ItemID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var ItemProInfo ProInfo

    for resultsIterator.HasNext(){
        var ItemInfos ItemInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ItemInfos)
        if ItemInfos.ItemProInfo.ItemName != ""{
            ItemProInfo = ItemInfos.ItemProInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(ItemProInfo)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    if len(args) !=1{
        return shim.Error("Incorrect number of arguments.")
    }
    ItemID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(ItemID)

    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var ItemIngInfo []IngInfo
    for resultsIterator.HasNext(){
        var ItemInfos ItemInfo
        response,err := resultsIterator.Next()
        if err != nil{
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ItemInfos)
        if ItemInfos.ItemIngInfo != nil{
            ItemIngInfo = ItemInfos.ItemIngInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(ItemIngInfo)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

    var LogInfos []LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    ItemID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(ItemID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var ItemInfos ItemInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ItemInfos)
        
    }
    jsonsAsBytes,err := json.Marshal(LogInfos)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    var Loginfo LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    ItemID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(ItemID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var ItemInfos ItemInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&ItemInfos)
        
    }
    jsonsAsBytes,err := json.Marshal(Loginfo)
    if err != nil{
        return shim.Error(err.Error ())
    }
    return shim.Success(jsonsAsBytes)
}


func main(){
     err := shim.Start(new(FoodChainCode))
     if err != nil {
         fmt.Printf("Error starting Food chaincode: %s ",err)
     }
}
