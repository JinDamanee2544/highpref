FROM python:3-slim

WORKDIR /app

RUN pip install ortools

COPY main.py prog

ENTRYPOINT ["python"]
CMD [ "prog" ]

