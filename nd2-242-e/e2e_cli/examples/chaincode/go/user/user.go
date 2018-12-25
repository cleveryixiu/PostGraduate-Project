package main

import (

    "crypto/md5"
    "encoding/json"
    "encoding/hex"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

type UserChainCode struct{
}

//User数据结构体
type UserInfo struct{
    UserName string `json:UserName`                             //用户名称
    UserType string `json:UserType`                             //用户类型
    UserPsw  string `json:UserPsw`                              //用户密码
    UserRpsw  string `json:UserRpsw`                            //用户确认密码
    UserID  string  `json:UserID`                               //用户唯一ID
}



func (a *UserChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (a *UserChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

    fn,args := stub.GetFunctionAndParameters()

    if fn == "addUserInfo"{
             return a.addUserInfo(stub,args)
    } else if fn == "getUserInfo"{
             return a.getUserInfo(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}


func (a *UserChainCode) addUserInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error 
    var UserInfo UserInfo

    if len(args)!=5{
        return shim.Error("Incorrect number of arguments.")
    }
    UserInfo.UserName = args[0]
    if UserInfo.UserName == ""{
        return shim.Error("UserName can not be empty.")
    }
    

    UserInfo.UserType = args[1]
    UserInfo.UserPsw = args[2]
    UserInfo.UserRpsw = args[3]
    UserInfo.UserID= args[4]

    if(UserInfo.UserPsw !=  UserInfo.UserRpsw){
        return shim.Error("The password entered two times is inconsistent.")
    }

    UserInfoJSONasBytes,err := json.Marshal(UserInfo)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(UserInfo.UserID,UserInfoJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    jsonsAsBytes,err := json.Marshal(UserInfo.UserID)
    if err != nil{
        return shim.Error(err.Error())
    }


    return shim.Success(jsonsAsBytes)
}




func(a *UserChainCode) getUserInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    UserID := args[0]

    Avalbytes, err := stub.GetState(UserID)
    if err != nil {
	jsonResp := "{\"Error\":\"Failed to get state for " + UserID + "\"}"
	return shim.Error(jsonResp)
    }

    if Avalbytes == nil {
	jsonResp := "{\"Error\":\"Nil amount for " + UserID + "\"}"
	return shim.Error(jsonResp)
    }

    jsonResp := "{\"UserID\":\"" + UserID + "\",\"UserInfo\":\"" + string(Avalbytes) + "\"}"
    fmt.Printf("Query Response:%s\n", jsonResp)
    return shim.Success(Avalbytes)
}
 


func main(){
     err := shim.Start(new(UserChainCode))
     if err != nil {
         fmt.Printf("Error starting Food chaincode: %s ",err)
     }
}
