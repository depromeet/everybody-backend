# uploader

전달 받은 이미지 파일을 리사이즈 후 S3에 업로드합니다.

Lambda와 API Gateway를 바탕으로 운영됩니다.

## 배포

```
# EVERYBODY_ENVIRONMENT 생략 시 기본값은 sandbox
# 디버깅용 샌드박스 환경 배포
$ EVERYBODY_ENVIRONMENT=sandbox sls deploy # or just $ sls deploy

# 개발 환경 배포
$ EVERYBODY_ENVIRONMENT=dev sls deploy
```