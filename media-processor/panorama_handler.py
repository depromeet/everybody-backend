try:
    import unzip_requirements
except ImportError:
    pass

import cv2
import numpy as np
import requests


# image concat을 통한 파노라마
# 참고: https://stackoverflow.com/questions/7589012/combining-two-images-with-opencv
def handle(event, context):
    tmp_s3_url = "https://tmp-upload-1111.s3.ap-northeast-2.amazonaws.com"
    tmp_s3_hosted_url = "http://tmp-upload-1111.s3-website.ap-northeast-2.amazonaws.com"
    # 미리 업로드해놨음 임시로
    tmp_image_keys = ["atg.jpg", "jinsu.jpg", "kitae.jpg"]
    tmp_images = []
    for image_key in tmp_image_keys:
      r = requests.get("/".join((tmp_s3_url, image_key)))
      print("/".join((tmp_s3_url, image_key)), r.status_code)
      image_np_array = np.asarray(bytearray(r.content), dtype=np.uint8)
      image = cv2.imdecode(image_np_array, cv2.IMREAD_COLOR)
      h, w = image.shape[:2]
      # center = (w / 2, h / 2)
      image = image[0:min(h, w), 0:min(h, w)]
      image = cv2.resize(image, dsize=(480, 480), interpolation=cv2.INTER_LINEAR)
      tmp_images.append(image)

    # np.ndarray인데 image 객체로 쓰일 수 있나봄
    transformed_result = np.concatenate(tmp_images, axis=1)
    body = cv2.imencode(".jpg", transformed_result)[1].tostring()
    requests.put("/".join((tmp_s3_url, "output.jpg")), data=body)

    # cv2.imwrite('out.png', transformed_result)

    response = {"statusCode": 200,
              "body": "/".join((tmp_s3_hosted_url, "output.jpg"))}

    return response