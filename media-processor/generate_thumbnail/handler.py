import io
from urllib import parse
import boto3
from PIL import Image
# numpy를 쓰는 이상 boto3를 또 dependency로 업로드하면
# 제한인 250MB를 넘어가려한다... ㅜㅜ

INPUT_OBJECT_PREFIX = 'original/'
# output path가 input path와 동일하면
# lambda 무한 loop trigger가 걸리니 절대 유의
OUTPUT_BUCKET_NAME = 'everybody-upload-output-dev-1'
OUTPUT_OBJECT_PREFIX = 'resized/'

def handle(event, context):
    source_bucket_name = event['Records'][0]['s3']['bucket']['name']

    # key는 url encoding 되어 전송될 수 있다. 한글이나 특문, 공백 등등에서는 url encoding 된 채로
    # s3에 접근하면 올바르게 object를 찾을 수 없다.
    source_object_key = parse.unquote(event['Records'][0]['s3']['object']['key'])
    print(f'썸네일 생성 for {source_object_key}')
    filename = source_object_key[len(INPUT_OBJECT_PREFIX):]
    s3 = boto3.resource('s3')

    obj = s3.Object(source_bucket_name, source_object_key)
    data = obj.get()['Body'].read()
    output = generate_thumbnail(data)

    output_bucket = s3.Bucket(OUTPUT_BUCKET_NAME)
    for size, img in output.items():
        tmp_file = io.BytesIO()
        img.save(tmp_file, 'png', optimize=True)
        output_bucket.put_object(Body=tmp_file.getvalue(), Key=f'{OUTPUT_OBJECT_PREFIX}{size}/{filename}', ContentType='image/png')

def generate_thumbnail(data) -> dict:
    img = Image.open(io.BytesIO(data))
    output = dict()
    for size in (32, 256, 1024):
        w = max(size, img.width)
        h = img.height * w // img.width
        resized = img.resize((w, h))
        output[str(size)] = resized

    # output은
    # key: str, value Image.Image
    return output
