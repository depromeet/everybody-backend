import json
import traceback

import cv2, os
import numpy as np
import json
import base64
import boto3
from botocore.exceptions import ClientError
from typing import List

INPUT_BUCKET_NAME = 'everybody-upload-output-dev-1'
AUTH_HEADER = 'user'

s3 = boto3.resource('s3')
bucket = s3.Bucket(INPUT_BUCKET_NAME)

def handle(event, context):
    print(event)
    download_video_request = DownloadVideoRequest(**json.loads(event['body']))
    size = (768, 1024) # 기본 size
    user = get_auth_user(event)
    if not user:
        return {
            "headers": {
                "Content-Type": "text/plain",
            },
            'isBase64Encoded': False,
            'statusCode': 401,
            'body': '유저 정보를 입력해주세요.',
        }
    images = []
    for k in download_video_request.keys:
        s3_absolute_key = f'{user}/image/768/{k}'
        print(f's3_absolute_key: {s3_absolute_key}')
        obj = bucket.Object(s3_absolute_key)
        try:
            # 참고: https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/s3.html#S3.Object.get
            encoded_img = np.fromstring(obj.get()['Body'].read(), dtype=np.uint8)
            img = cv2.imdecode(encoded_img, cv2.IMREAD_COLOR)
            # size는 (768, 1024)로 고정하는게 편함
            # height, width, layers = img.shape
            # size = (width, height)
            images.append(img)
        except ClientError as e:
            if e.response['Error']['Code'] == 'NoSuchKey':
                return {
                    "headers": {
                        "Content-Type": "text/plain",
                    },
                    'isBase64Encoded': False,
                    'statusCode': 403,
                    'body': '해당 파일에 접근할 권한이 없습니다. 존재하지 않는 파일이거나 본인의 파일이 아닙니다.',
                }
            else:
                traceback.print_exc()


    # os.chdir('/tmp')
    # 'video-demo.mp4'라는 filename 생성(/tmp에 mp4 생성), X frame/sec
    output = cv2.VideoWriter('/tmp/video-demo.mp4', cv2.VideoWriter_fourcc(*'mp4v'), get_fps(event), size)
    os.system('ls /tmp')
    for image in images:
        output.write(image)
    output.release()

    # 만든 mp4 파일 자체를 rb로 열어서 payload로 보내준다.
    mp4_file = open('/tmp/video-demo.mp4', 'rb')

    # S3에 업로드하지 않고 /tmp의 내용을 바로 base64를 통해 encode, decode하고
    # Content-Type을 정의함으로써 binary data를 전달할 수 있음.
    response = {
        "headers": {
            "Content-Type": "video/mp4",
        },
        'isBase64Encoded': True,
        "statusCode": 200,
        "body": base64.b64encode(mp4_file.read()).decode('utf-8'),
    }

    mp4_file.close()

    return response


def sample_download_video_request():
    return DownloadVideoRequest(["sample_images_09.jpg", "sample_images_10.jpg", "sample_images_11.jpg", "sample_images_12.jpg"], None)

def get_auth_user(event):
    return event.get('headers').get(AUTH_HEADER)

def get_fps(event):
    if 'body' not in event:
        raise RuntimeError('body를 정확히 전달해주세요.')
    body = json.loads(event['body'])
    # duration 단위는 초
    return 1 / float(body.get('duration', '0.25'))

class DownloadVideoRequest:
    def __init__(self, keys: List[str], duration: float):
        self.keys = keys if keys is not None else []
        self.duration = duration if duration is not None else 0.25

