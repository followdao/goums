# go-UMS -- a golang based Terminal /Member / User Management Sub-system (UMS) with AAA



## 0. Status

this project back to active development, and re-design all.

Thie project aims for a Minimum Viable Product (MVP) or a prototype of UMS.

 

## 1. purpose

a general terminal / member management sub-system for TV-box ( android STB )  with AAA

> AAA --- Authentication（认证） / Authorization （授权） / Accounting (计费）

### 1.1  feature highlight ( maybe change  until v1.0.0 release ):

* collect all business logic in UMS , support multiple AAA server with local session cache

* support multiple storage with interface ( adapter to multiple storage driver )

* support administrator / intergration API  via  gRPC and  RESTful

* add operator/administrator web UI for operation

### 1.2  Phase 1 plan

  a MVP / prototype demo only, for TV-box / STB terminal 

## 2. design ( draft )

### 2.1 architecutre

![go-ums-all](./docs/go-ums-architecture-201912.png)

go-ums-interface-201912.png

### 2.2 Business process / scenario:

* In the diagram, mark 1 is  the serial number generator  ( tvsn ) 
will generate the hardware serial number of the TV box / set-top box, import it into the database, and save it in an Excel file. The excel file is send to the factory,  and the serial number is burned to the TV box / set-top box when it is produced In the product as hardware ID
*  mark 2 is Mgn, 
 provides services for the background management UI ( admin web UI ) , and provides business integration adapter ( gRPC / RESTful ... ) , support  tools like tvsn  and 3rd system / application to manages TV box / set-top box terminals, including terminal activation / deactivation, member account lifecycle management corresponding to the box etc.
*  mark 3 is android APK inside TV box / set-top box,
APK  access AAA for register ( active ) / login ( authentication ) / get the TV-guide portal IP address and access token ( authorization ) 
*  mark 4 and 5 , the AAA / UMS provide member magement and some business logic like AAA....
* mark 6 , the Session server provide session storage , and sync session status change to AAA local cache 



  


### 2.3. protocols between modules

![go-ums-interface](./docs/go-ums-interface-201912.png)

* In the diagram, mark 1 is the HTTP protocol ( RESTful ),  JSON encode 
* mark 2,   gRPC with protobuf encode 
* Mark 3 ,  gRPC with flatbuffers encode  (  used in all modules inside go-UMS )
* Mark 4,   TCP with bytes payload ( via customized encrypt  )
* Mark 5,   HTTP with bytes payload ( via customized encrypt )



### 2.4  data models 

#### 2.4.1  objects diagram

### 2.4.2  sql and funtions in postgresql 
### 2.4.3  IDL in protobuffers / flatbuffers 



## 4. tech stack

1. base on golang and high performance go module like gRPC / flatbuffers / fasthttp / fastcache ......
2. web UI base react javascript / HTML / css ......
3. postgresql 11+

   

## 4. License

MIT
