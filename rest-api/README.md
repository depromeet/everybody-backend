# 에브리바디 rest-api

API Gateway 뒷단에서 RESTful API를 제공합니다.

## ent ORM 사용법

### 새로운 테이블 정의

```shell
$ ent init User
```

### 정의한 스키마를 바탕으로 ORM 코드 생성

```shell
$ go generate ./ent
```

## 모킹(Mocking)

모킹은 단위 테스트 코드 작성 시 아주 편리한 테스트 기법이다. 
본 서비스는 [`stretchr/testify/mock`](https://github.com/stretchr/testify) 라는 모킹 프레임워크와
그 이용을 편리하게 해주는 CLI 도구 [`mockery`](https://github.com/vektra/mockery) 를 사용하고 있다.
새로운 인터페이스를 생성했고, 그에 대한 모킹 타입이 필요한 경우 다음의 커맨드를 통해 Mock type을 정의할 수 있다.

```shell
$ make mock
// check out mocks/
```

## 테스트 코드

```shell
$ go clean -testcache && go test ./...
#  or just
$ make test

# 다음과 같은 환경변수가 필요할 수 있습니다.
$ EVERYBODY_ENVIRONMENT=local EVERYBODY_REST_CONFIG_PATH=$(pwd)/config make test
```
