package module

import (
	"fmt"
	"github.com/fofapro/fofa-go/fofa"
	jsoniter "github.com/json-iterator/go"
	"os"
	"sort"
	"strings"
)

var(
	json=jsoniter.ConfigCompatibleWithStandardLibrary
)
func isTakeIn(result []string,target ...string,)[]string{  //证书去重
	sort.Strings(result)
	for _,i:=range target{
		index:=sort.SearchStrings(result,i)
		if index<len(result) && result[index]==i{
			continue
		}
		result=append(result,i)
		sort.Strings(result)
	}
	return result
}
func parseResult(result1 []byte) ([][]string,[]string) {
	var temp Response
	err:=json.UnmarshalFromString(string(result1),&temp)
	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var protocolSlice []string
	var certSlice []string
	var excelSlice [][]string
	//只搜集url，协议部分暂不处理
	for _,tempItem:=range temp.Results{
		var tempSlice []string
		if len(tempItem[5])>0{     //搜集cert信息
			name1,name2:=GetCert(tempItem[5])
			certSlice=isTakeIn(certSlice,name1,name2)
		}
		tempStr:=""
		if len(tempItem[4])!=0{
			tempStr+=tempItem[0]+"\t"+tempItem[4]+"\t"+tempItem[2]
			protocolSlice=append(protocolSlice,tempStr)
		}else {
			if strings.HasPrefix(tempItem[0],"http"){
				tempSlice=append(tempSlice,tempItem[0],tempItem[2],tempItem[3])
				//urlAll+=tempItem[0]+"\r\n"
			}else {
				tempSlice=append(tempSlice,"http|https://"+tempItem[0],string(tempItem[2]),tempItem[3])
				//urlAll+="http|https://"+tempItem[0]+"\r\n"
			}
			excelSlice=append(excelSlice,tempSlice)
		}
	}
	//fmt.Println(protocolSlice)
	return excelSlice,certSlice
}
//分拣出cert数据，主要是Organization和CommonName字段
func GetCert(certString string)(string,string){
	tempA:=strings.Split(strings.TrimSpace(strings.Split(certString,"Subject:")[1]),"\n\nSubject Public Key Info:")[0]
	tempOrganization:=strings.Split(strings.TrimSpace(strings.Split(tempA,"Organization:")[1]),"CommonName:")[0]
	tempCommonName:=strings.TrimSpace(strings.Split(tempA,"CommonName:")[1])
	return strings.TrimRight(strings.TrimSpace(tempOrganization),"\n"),tempCommonName
}
func FofaMain(search string,myConfig Config)([][]string,[]string){
	clt:=fofa.NewFofaClient([]byte(myConfig.Fofa.Email),[]byte(myConfig.Fofa.Api))
	if clt==nil{
		fmt.Println("fail to create fofa client")
		os.Exit(1)
	}
	tempResult,err:=clt.QueryAsJSON(1,[]byte(search),[]byte("host,ip,port,title,protocol,cert"))
	if err!=nil{
		fmt.Println("fofa api request failed")
		os.Exit(1)
	}
	outStr,certTemp:=parseResult(tempResult)
	return outStr,certTemp
}