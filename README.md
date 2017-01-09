# swan-search

swan-search is a simple search engine based on mesos framework swan[https://github.com/Dataman-Cloud/swan]

# installation
1. get the source code first:
"""
go get github.com/Dataman-Cloud/swan-search
"""

2. prepare the config file:
"""
cp deploy/config.json.template deploy/config.json
"""
In config.json, Fill the clusters info. One swan service is one cluster.
In config.json.template:
"swan1" is cluster name, "http://172.28.128.4:9999" is cluster ips.
swan-search supports multi-manager mode. You can register manager ips using this format:
"""
{
	swan1: "http://172.28.128.4:9999, http://172.28.128.3:9999"
}
"""

3. run search:
"""
# make docker-build
# make docker-run
"""
then search service is running with addr: 0.0.0.0:9888

# how to use:
when search is running, you can call its ip using this command:
"""
curl 0.0.0.0:9888/search/v1/luckysearch?keyword=nginx0051
"""
result is like this:
"""
{"code":"0","data":[{"ID":"nginx0052-zliu-swan1","Name":"nginx0052","Type":"app","Param":{"AppId":"nginx0052-zliu-swan1"}},{"ID":"0-nginx0052-zliu-swan1","Name":"0-nginx0052-zliu-swan1","Type":"task","Param":{"AppId":"nginx0052-zliu-swan1","TaskId":"0"}},{"ID":"1-nginx0052-zliu-swan1","Name":"1-nginx0052-zliu-swan1","Type":"task","Param":{"AppId":"nginx0052-zliu-swan1","TaskId":"1"}}]}
"""
