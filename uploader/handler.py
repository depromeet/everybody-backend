import base64
import io
import json
import traceback
import uuid
import boto3
from PIL import Image
from requests_toolbelt.multipart import decoder

RESIZE_SIZES = ((48, 64), (192, 256), (768, 1024))
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
        print(event)
        user = event['headers']['user']
        body = base64.b64decode(event['body'])
        print('body', body)
        boundary = event['headers']['content-type']
        print('boundary', boundary)
        multipart_body = decoder.MultipartDecoder(body, boundary)

        keys = []

        for item in multipart_body.parts:
            print('item', item)
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
                key = str(uuid.uuid4())
                print(f'{part.filename}을 업로드합니다.')
                resized_result = generate_thumbnail(part.content)
                for w, img in resized_result.items():
                    tmp_file = io.BytesIO()
                    img.save(tmp_file, 'png', optimize=True)
                    bucket.put_object(
                        # ACL='public-read', # public 권한 필요
                        ContentType='image/png',
                        Key=f'{user}/{OUTPUT_OBJECT_PREFIX}{w}/{key}',
                        Body=tmp_file.getvalue(),
                    )
                keys.append(key)

            else:
                print(f'필드명이 image가 아님: {part.name}')

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
