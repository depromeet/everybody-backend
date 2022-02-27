FROM public.ecr.aws/lambda/python:3.7
# ffmpeg 홈페이지: https://ffmpeg.org/download.html#build-linux
# 다운로드 페이지: https://johnvansickle.com/ffmpeg/
# intel 계열을 위한 컴파일된 static file
RUN yum install -y tar
ADD https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-amd64-static.tar.xz /tmp/ffmpeg.tar.xz
RUN cd /tmp &&  \
    tar Jxvf ffmpeg.tar.xz && \
    cp ffmpeg-git-20220108-amd64-static/ffmpeg /usr/local/bin/ffmpeg

#ADD https://johnvansickle.com/ffmpeg/builds/ffmpeg-git-arm64-static.tar.xz /tmp/ffmpeg.tar.xz
#RUN cd /tmp &&  \
#    tar Jxvf ffmpeg.tar.xz && \
#    cp ffmpeg-git-20220108-arm64-static/ffmpeg /usr/local/bin/ffmpeg \

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
COPY . .
CMD [ "handler.handle" ]

