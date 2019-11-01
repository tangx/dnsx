# DNSx

DNS-extra


## Usage

### config

dnsx configure  -> interactive
domain: domain.com
provider:

```json
{
    "domain.com":{
        "provider":"aliyun",
        "AC":"value",
        "AK":"vallue",
    }, 
    "domain.org":{
        "provider":"qcloud",
        "AC":"value",
        "AK":"vallue",
    }, 
    "domain.cn":{
        "provider":"dnspod",
        "email":"value",
        "password":"vallue",
    }
}
```

### domain 
dnsx <action> <domain> <type> <value>

action: set,get,update,delete
type: a,txt,cname
value: value

