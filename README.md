# everybody-backend
에브리바리 쉑더바리 렛츠고바리 컴온바리 ~ ♪ 제주도엔 다금바리 ~ ♪ 디프만엔 에블바리 ~ ♪

🏋️ 눈바디를 기록하는 에브리바디 서비스의 Golang + Python 마이크로서비스 아키텍쳐 기반 백엔드 (디프만 10기 일조권침해조)

## 사용 기술 (예상)

**Backend** - Elastic Beanstalk (t2.micro 프리티어)
  * Golang API Gateway - 인증, 인가를 수행하는 API 게이트웨이 역할
  * Golang RESTful API - API Gateway 뒷단에서 RESTful API를 제공

**Media-Processor** - Python(Lambda 프리티어)
  * OpenCV와 NumPy를 이용해 사진을 편집하거나 영상을 생성

**Build** & **Deploy** - Github Action (공짜)

**Infra** - Elastic Beanstalk 바탕으로 인프라 구축
  * S3 - 프리티어
  * RDS - 프리티어
  * ALB - 프리티어
  * EC2 - 프리티어
  * Lambda - 프리티어

**Contact** - (Email) everybody.dev10@gmail.com
