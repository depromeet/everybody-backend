# everybody-backend

๐๏ธ ๋๋ฐ๋๋ฅผ ๊ธฐ๋กํ๋ ์๋ธ๋ฆฌ๋ฐ๋ ์๋น์ค์ Golang + Python ๋ง์ดํฌ๋ก์๋น์ค ์ํคํ์ณ ๊ธฐ๋ฐ ๋ฐฑ์๋ (๋ํ๋ง 10๊ธฐ ์ผ์กฐ๊ถ์นจํด์กฐ)

## ์ฌ์ฉ ๊ธฐ์  (์์)

**Backend** - Elastic Beanstalk (t2.micro ํ๋ฆฌํฐ์ด)
  * Golang API Gateway - ์ธ์ฆ, ์ธ๊ฐ๋ฅผ ์ํํ๋ API ๊ฒ์ดํธ์จ์ด ์ญํ 
  * Golang RESTful API - API Gateway ๋ท๋จ์์ RESTful API๋ฅผ ์ ๊ณต

**Media-Processor** - Python(Lambda ํ๋ฆฌํฐ์ด)
  * OpenCV์ NumPy๋ฅผ ์ด์ฉํด ์ฌ์ง์ ํธ์งํ๊ฑฐ๋ ์์์ ์์ฑ

**Build** & **Deploy** - Github Action (๊ณต์ง)

**Infra** - Elastic Beanstalk ๋ฐํ์ผ๋ก ์ธํ๋ผ ๊ตฌ์ถ
  * S3 - ํ๋ฆฌํฐ์ด
  * RDS - ํ๋ฆฌํฐ์ด
  * ALB - ํ๋ฆฌํฐ์ด
  * EC2 - ํ๋ฆฌํฐ์ด
  * Lambda - ํ๋ฆฌํฐ์ด

**Contact** - (Email) everybody.dev10@gmail.com
