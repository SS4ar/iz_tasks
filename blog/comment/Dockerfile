FROM python:3.9-alpine
WORKDIR /code

RUN addgroup -g 2000 user \
    && adduser -u 2000 -G user -s /bin/sh -D user
RUN apk add curl
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
EXPOSE 80
COPY . .
RUN chmod +x ./entrypoint.sh

CMD ["./entrypoint.sh"]