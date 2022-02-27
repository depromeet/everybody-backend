# 에브리바디 video

OpenCV와 NumPy를 이용해 영상을 생성합니다. Lambda + APIGateway를 통해 제공됩니다.

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

**pip 의존성 관리**

아주 성가신 기술적 한계인 것 같은데 OpenCV, Numpy를 쓰면서 거의 Lambda의 layer 크기 제한 가까이 크기가 커진다.
그래서 추가적인 패키지를 설치하지 못하는 이슈가 있다.
미니멀하게 서비스를 구현해서 현재까지는 큰 문제는 없는데, 한 가지 번거로운 게 Lambda에는 boto3가 내장되어있어 의존성을 추가해주지 않아도 되는데,
`pip freeze`를 통해 로컬의 의존성을 전달할 때 boto3가 자동으로 삽입되고, 그럼 layer 크기 제한 오류가 발생한다.

따라서 `pip freeze`할 때마다 꼭 필요한 패키지만 명시했는지 확인해줘야한다 ㅜㅜ

## custom base image

```shell
aws ecr get-login-password --region ap-northeast-2 | docker login --username AWS --password-stdin 070251821212.dkr.ecr.ap-northeast-2.amazonaws.com 
docker build . -t 070251821212.dkr.ecr.ap-northeast-2.amazonaws.com/everybody-python
docker push 070251821212.dkr.ecr.ap-northeast-2.amazonaws.com/everybody-python
```