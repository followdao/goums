# go-ums -- a golang based Terminal /Member / User Management Sub-system (UMS) with AAA



## 0. Status

this project back to active development, and re-design all.

Thie project aims for a Minimum Viable Product (MVP) or a prototype of UMS.

 

## 1. purpose

a general terminal / member management sub-system for TV-box ( android STB )  with AAA

> AAA --- Authentication（认证） / Authorization （授权） / Accounting (计费）

Phase 1:   go-ums is a MVP / prototype, for TV-box / STB terminal 


## 2. architecutre

![go-ums-all](./docs/go-ums-architecture-201912.png)



### 2.1 Business process / scenario:

1. The serial number generator  ( tvsn ) will generate the hardware serial number of the TV box / set-top box, import it into the database, and save it in an Excel file. The excel file is send to the factory,  and the serial number is burned to the TV box / set-top box when it is produced In the product as hardware ID
2. Mgn provides services for the background management UI ( admin web UI ) , and provides business integration adapter ( gRPC / RESTful ... ) , support  tools like tvsn  and 3rd system / application to manages TV box / set-top box terminals, including terminal activation / deactivation, member account lifecycle management corresponding to the box etc.
3. android apk inside TV box / set-top box, access AAA for register ( active ) / login ( authentication ) / get the TV-guide portal IP address and access token ( authorization ) 
4. AAA / UMS provide member magement and some business logic like AAA....
5. Session server provide session storage , and sync session status change to AAA local cache 



### 2.2. highlight ( maybe, it's will change everythings  until v1.0 release ):

* collect all business logic in UMS , support multiple AAA server with local session cache

* support multiple storage with interface ( adapter to multiple storage driver )

* support administrator / intergration API  via  gRPC and  RESTful

* add operator/administrator web UI for operation


## 3. tech stack

1. base on golang and high performance go module like gRPC / flatbuffers / fasthttp / fastcache ......
2. web UI base react javascript / HTML / css ......
3. postgresql 11+

   

## 4. License

MIT
