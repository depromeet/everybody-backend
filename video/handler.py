import cv2, os
import numpy as np
import requests
import base64

def handle(event, context):
    img_array = []
    tmp_s3_url = "https://tmp-upload-1111.s3.ap-northeast-2.amazonaws.com"
    input_images_key = sample_images_list()
    print(input_images_key)

    for image_key in input_images_key:
        r = requests.get("/".join((tmp_s3_url, image_key)))
        print(r.status_code)
        encoded_img = np.fromstring(r.content, dtype=np.uint8)
        img = cv2.imdecode(encoded_img, cv2.IMREAD_COLOR)
        height, width, layers=img.shape
        size=(width, height)
        img_array.append(img)

    # os.chdir('/tmp')
    # 'video-demo.mp4'라는 filename 생성(/tmp에 mp4 생성), 0.5frame/sec
    output=cv2.VideoWriter('/tmp/video-demo.mp4', cv2.VideoWriter_fourcc(*'mp4v'), 0.5, size)
    os.system('ls /tmp')
    for i in range(len(img_array)):
        output.write(img_array[i])
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

    return response


def sample_images_list():
    tmp_image_keys = ["sample_images_09.jpg", "sample_images_10.jpg", "sample_images_11.jpg", "sample_images_12.jpg"]

    return tmp_image_keys
