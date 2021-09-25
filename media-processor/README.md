# 에브리바디 media-processor

OpenCV와 NumPy를 이용해 사진을 편집하거나 영상을 생성합니다. Lambda + APIGateway를 통해 제공됩니다.

serverless framwork 참고 예시: https://github.com/serverless/examples/tree/master/aws-python-http-api

## 로컬에서 실행

```shell
$ serverless invoke local --function hello
$ serverless offline
```

## Deploy

```shell
# cat ~/.aws/config 를 통해 [profile depromeet-everybody] 프로파일이 올바르게 설정됐는지 확인한 뒤 진행하세요.

$ serverless deploy
```

