FROM docker.io/python:3-slim
WORKDIR /rofl-8004

COPY . /rofl-8004

RUN apt update && apt install -y git python3-dev gcc libc-dev && apt clean && rm -rf /var/lib/apt/lists/*
RUN pip install -r requirements.txt

ENTRYPOINT ["python", "main.py"]
