import json
import traceback
import cv2
import numpy as np
import json
import base64
import boto3
from botocore.exceptions import ClientError
from typing import List
import os
import subprocess

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

    # 'video-demo.mp4'라는 filename 생성(/tmp에 mp4 생성), X frame/sec
    # 'MPEG' - iOS에서 다운로드는 잘 되는데 재생 자체가 안됨(갤러리에 안 뜸)
    # mkv X264 인코더를 찾을 수 없다며 에러
    # mp4 X264 인코더를 찾을 수 없다며 에러
    # mp4 MJPG ?
    # mp4 MP4V tag 0x5634504d/'MP4V' is not supported with codec id 12 and format 'mp4 / MP4 (MPEG-4 Part 14)', 생성은 하는데 iOS에선 초록색
    # mp4 m,p,4,v 잘 만들어지긴 하는데 초록생
    # mp4 a,v,c,1
    # 아무래도 machine에 X264 decoder 가 설치되어있어야하는듯..?
    # avi XVID => avi는 iOS에서 재생이 안 되는

    output = cv2.VideoWriter('/tmp/video-demo.mp4', cv2.VideoWriter_fourcc(*'mp4v'), get_fps(event), size)
    os.system('ls /tmp')
    for image in images:
        output.write(image)
    output.release()

    output = subprocess.check_output('ffmpeg -y -i /tmp/video-demo.mp4 -vcodec libx264 /tmp/video-output.mp4', shell=True)
    print('output: ', output)
    os.system('ls /tmp')
    # 만든 mp4 파일 자체를 rb로 열어서 payload로 보내준다.
    mp4_file = open('/tmp/video-output.mp4', 'rb')

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

if __name__ == '__main__':
    size = (768, 1024)  # 기본 size
    user = 172
    download_video_request = DownloadVideoRequest([ "95edfbbc-cff3-49df-abbc-7ea295f4d5ad",
"1748b8a1-1bec-4300-961f-42df72d92498",], None)
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
                print("err")
            else:
                traceback.print_exc()

    # os.chdir('/tmp')
    # 'video-demo.mp4'라는 filename 생성(/tmp에 mp4 생성), X frame/sec
    # 'MPEG' - iOS에서 다운로드는 잘 되는데 재생 자체가 안됨(갤러리에 안 뜸)
    # mkv X264 인코더를 찾을 수 없다며 에러
    # mp4 X264 인코더를 찾을 수 없다며 에러
    # mp4 MJPG ?
    # mp4 MP4V tag 0x5634504d/'MP4V' is not supported with codec id 12 and format 'mp4 / MP4 (MPEG-4 Part 14)', 생성은 하는데 iOS에선 초록색
    # mp4 m,p,4,v 잘 만들어지긴 하는데 초록생
    # mp4 a,v,c,1
    # 아무래도 machine에 X264 decoder 가 설치되어있어야하는듯..?

    output = cv2.VideoWriter('/tmp/video-demo.mp4', cv2.VideoWriter_fourcc(*'mp4v'), 2, size)
    os.system('ls /tmp')
    for image in images:
        output.write(image)
    output.release()

    output = subprocess.check_output('/opt/ffmpeg -y -i /tmp/video-demo.mp4 -vcodec libx264 /tmp/video-output.mp4', shell=True)
    print('output: ', output)
    os.chdir('/tmp')

    # 만든 mp4 파일 자체를 rb로 열어서 payload로 보내준다.
    mp4_file = open('/tmp/video-output.mp4', 'rb')
