# file upload 시에
# API gateway가 lambda에게 보내는 event
# binary의 file data들도 base64 encoding 되어 string으로 담겨있음!
sample_data = {
    'version': '2.0',
    'routeKey': 'POST /upload',
    'rawPath': '/upload',
    'rawQueryString': 'random_key_enabled=false',
    'queryStringParameters': {'random_key_enabled': 'true'},
    'headers': {
        'user': 10,
        'accept': '*/*',
        'accept-encoding': 'gzip,deflate,br',
        'cache-control': 'no-cache',
        'content-length': '702',
        'content-type': 'multipart/form-data; boundary=--------------------------389603159367705259395659',
        'host': 'abbi52l09j.execute-api.ap-northeast-2.amazonaws.com',
        'postman-token': 'c96333a8-e789-4b1f-aeda-c86c4b8cf51f',
        'user-agent': 'PostmanRuntime/7.26.8',
        'x-amzn-trace-id': 'Root=1-616e31dc-47c5a0452694876b1af2a409',
        'x-forwarded-for': '124.50.93.166',
        'x-forwarded-port': '443',
        'x-forwarded-proto': 'https'
    },
    'requestContext': {
        'accountId': '070251821212',
        'apiId': 'abbi52l09j',
        'domainName': 'abbi52l09j.execute-api.ap-northeast-2.amazonaws.com',
        'domainPrefix': 'abbi52l09j',
        'http': {'method': 'POST',
        'path': '/upload',
        'protocol': 'HTTP/1.1',
        'sourceIp': '124.50.93.166',
        'userAgent': 'PostmanRuntime/7.26.8'},
        'requestId': 'Hby6ghgNIE0EJXg=',
        'routeKey': 'POST /upload',
        'stage': '$default',
        'time': '19/Oct/2021:02:47:56 +0000',
        'timeEpoch': 1634611676671
    },
    'body': 'LS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLTM4OTYwMzE1OTM2NzcwNTI1OTM5NTY1OQ0KQ29udGVudC1EaXNwb3NpdGlvbjogZm9ybS1kYXRhOyBuYW1lPSJpbWFnZSI7IGZpbGVuYW1lPSJTY3JlZW5zaG90IGZyb20gMjAyMS0wNi0yMCAxMy01NC0xNC5wbmciDQpDb250ZW50LVR5cGU6IGltYWdlL3BuZw0KDQqJUE5HDQoaCgAAAA1JSERSAAABLQAAAEsIBgAAACU50R0AAAAEc0JJVAgICAh8CGSIAAAAGXRFWHRTb2Z0d2FyZQBnbm9tZS1zY3JlZW5zaG907wO/PgAAAERpVFh0Q3JlYXRpb24gVGltZQAAAAAAMjAyMeuFhCAwNuyblCAyMOydvCAo7J28KSDsmKTtm4QgMDHsi5wgNTTrtoQgMTbstIhsvmx6AAABFUlEQVR4nO3UQQ0AIRDAwONkrH+foIEXaTKjoK+umdkfQMT/OgDghmkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKaYFpJgWkGJaQIppASmmBaSYFpBiWkCKaQEppgWkmBaQYlpAimkBKQcgowHjzv5XlgAAAABJRU5ErkJggg0KLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLTM4OTYwMzE1OTM2NzcwNTI1OTM5NTY1OS0tDQo=',
    'isBase64Encoded': True
}