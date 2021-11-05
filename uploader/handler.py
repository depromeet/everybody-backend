import base64
import io
import json
import time
import traceback
import uuid
import boto3
from PIL import Image
from requests_toolbelt.multipart import decoder

RESIZE_SIZES = (
    (48, 64),
    (192, 256),
    (768, 1024)
)

OUTPUT_BUCKET_NAME = 'everybody-upload-output-dev-1'
OUTPUT_OBJECT_PREFIX = 'image/'

s3 = boto3.resource('s3')
bucket = s3.Bucket(OUTPUT_BUCKET_NAME)

class Multipart:
    def __init__(self, content, content_type, name, filename):
        self.content = content
        self.content_type = content_type
        self.name = name
        self.filename = filename

# api gateway body => S3 upload
# 참고: https://gist.github.com/tomfa/87947d2773b60fc3797491d6ef5e3d0e
# https://devlog-wjdrbs96.tistory.com/331
def handle(event, context):
    try:
        # debug 할 때에만 event 출력
        if event.get("queryStringParameters", {"debug": 'false'}) == 'true':
            print(event)
        # 기본적으로 random_key_enabled는 True
        random_key_enabled = event.get("queryStringParameters", {"random_key_enabled": 'true'}).get("random_key_enabled") == 'true'
        start_handle_time = time.time()
        user = event['headers'].get('user', '')
        if user == '' or not str(user).isdigit():
            return {
            'statusCode': 401,
            'headers': {
                'Content-Type': 'application/json'
            },
            'body': {
                'keys': None,
                'error': '유저 인증 정보가 존재하지 않습니다.'
            }
        }
        body = base64.b64decode(event['body'])
        finish_base64_decode_time = time.time()
        print(f'Elapsed(base64 디코딩 완료): {finish_base64_decode_time - start_handle_time}')
        # print('body', body)
        boundary = event['headers']['content-type']
        print('boundary', boundary)
        multipart_body = decoder.MultipartDecoder(body, boundary)

        keys = []

        for item in multipart_body.parts:
            # print('item', item)
            print('item.headers', item.headers)

            disposition = item.headers[b'Content-Disposition']
            params = {}
            # 참고: https://stackoverflow.com/a/6224nContent-Disposition4172/9471220
            for disposition_part in str(disposition).split(';'):
                kv = disposition_part.split('=', 2)
                params[str(kv[0]).strip()] = str(kv[1]).strip('\"\'\t \r\n') if len(kv) > 1 else str(kv[0]).strip()
            content_type = item.headers[b'Content-Type'].decode('utf-8') if b'Content-Type' in item.headers else None

            part = Multipart(item.content, content_type, params.get('name'), params.get('filename'))
            if part.name == 'image':
                key = str(uuid.uuid4()) if random_key_enabled else part.filename
                print(f'{part.filename}을 업로드합니다.')
                resized_result = generate_thumbnail(part.content)
                finish_resize_time = time.time()
                print(f'Elapsed(리사이징 완료): {finish_resize_time - start_handle_time}')
                
                for w, img in resized_result.items():
                    tmp_file = io.BytesIO()
                    img.save(tmp_file, 'png', optimize=False) # optimize를 True로 만들면 Lambda에서는 엄청나게 오랜 시간이 소요되어버림. 왜지?
                    finish_tmp_save = time.time()
                    print(f'Elapsed({w} 크기의 사진 임시 저장 완료): {finish_tmp_save - start_handle_time}')
                    bucket.put_object(
                        # ACL='public-read', # public 권한 필요
                        ContentType='image/png',
                        Key=f'{user}/{OUTPUT_OBJECT_PREFIX}{w}/{key}',
                        Body=tmp_file.getvalue(),
                    )
                    tmp_file.close()
                    finish_upload_time = time.time()
                    print(f'Elapsed({w} 크기의 사진 업로드 완료): {finish_upload_time - start_handle_time}')
                keys.append(key)

            else:
                print(f'필드명이 image가 아님: {part.name}')

        finish_request_time = time.time()
        print(f'Elapsed(요청 처리 완료): {finish_request_time - start_handle_time}')

        return {
            'statusCode': 200,
            'headers': {
                'Content-Type': 'application/json'
            },
            'body': json.dumps({
                'data': {
                    'keys': keys,
                },
                'error': None
            })
        }
    
    except Exception as e:
        traceback.print_exc()
        return {
            'statusCode': 500,
            'headers': {
                'Content-Type': 'application/json'
            },
            'body': {
                'keys': None,
                'error': str(e)
            }
        }

# 이미지 데이터를 받아 width와 사진의 dict로 리턴
def generate_thumbnail(data) -> dict:
    img = Image.open(io.BytesIO(data))
    output = dict()
    for w, h in RESIZE_SIZES:
        resized = img.resize((w, h))
        output[w] = resized

    # output은
    # key: str, value Image.Image
    return output

# 로컬 개발용
if __name__ == '__main__':
    from sample_data import sample_data
    print(handle(sample_data, {}))
