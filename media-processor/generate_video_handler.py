try:
    import unzip_requirements
except ImportError:
    pass

import cv2, os
import numpy as np
import requests


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
    r = requests.put("/".join((tmp_s3_url, "video-demo.mp4")), data=mp4_file)
    print(r.status_code)

    response = {"statusCode": 200, "body": {"result": "generate video successfully"}}

    return response


def sample_images_list():
    tmp_image_keys = ["sample_images_09.jpg", "sample_images_10.jpg", "sample_images_11.jpg", "sample_images_12.jpg"]

    return tmp_image_keys