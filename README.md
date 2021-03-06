# sap-api-integrations-product-master-class-reads
sap-api-integrations-product-master-class-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 品目クラスデータを取得するマイクロサービスです。      
sap-api-integrations-product-master-class-reads には、サンプルのAPI Json フォーマットが含まれています。     
sap-api-integrations-product-master-class-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。     
https://api.sap.com/api/OP_API_CLFN_PRODUCT_SRV/overview   

## 動作環境  
sap-api-integrations-product-master-class-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。    
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-product-master-class-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。    

## 本レポジトリ が 対応する API サービス
sap-api-integrations-product-master-class-reads が対応する APIサービス は、次のものです。  

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_CLFN_PRODUCT_SRV/overview    
* APIサービス名(=baseURL): API_CLFN_PRODUCT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-product-master-class-reads には、次の API をコールするためのリソースが含まれています。  

* A_ClfnProduct（品目クラス - 一般）※品目クラス関連データを取得するために、ToProductClass、ToClassDetails、ToProductCharc、と合わせて利用されます。
* ToProductClass（品目クラス - 品目クラス）
* ToClassDetails（品目クラス - クラス詳細）
* ToProductCharc（品目クラス - 品目特性）

## API への 値入力条件 の 初期値
sap-api-integrations-product-master-class-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ProductClass.Product（品目）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ProductGeneral" が指定されています。

```
	"api_schema": "A_ClfnProduct",
	"accepter": ["ProductGeneral"],
	"product_code": "AVC_RBT_APPL_UNIT",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "A_ClfnProduct",
	"accepter": ["All"],
	"product_code": "AVC_RBT_APPL_UNIT",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProductMasterClass(product string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ProductGeneral":
			func() {
				c.ProductGeneral(product)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP の 品目クラス　一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"Product" ～ "to_ProductCharc" は、/SAP_API_Output_Formatter/type.go 内 の type ProductGeneral{}による出力結果です。  
"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-product-master-class-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-product-master-class-reads/SAP_API_Caller.(*SAPAPICaller).ProductGeneral",
	"level": "INFO",
	"message": [
		{
			"Product": "AVC_RBT_APPL_UNIT",
			"ProductType": "KMAT",
			"CreationDate": "/Date(1587945600000)/",
			"LastChangeDate": "/Date(1619913600000)/",
			"IsMarkedForDeletion": false,
			"ProductGroup": "L004",
			"BaseUnit": "PC",
			"ProductHierarchy": "",
			"to_ProductClass": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_CLFN_PRODUCT_SRV/A_ClfnProduct('AVC_RBT_APPL_UNIT')/to_ProductClass",
			"to_ProductCharc": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_CLFN_PRODUCT_SRV/A_ClfnProduct('AVC_RBT_APPL_UNIT')/to_ProductCharc"
		}
	],
	"time": "2021-12-24T12:19:45.880643+09:00"
}

```
