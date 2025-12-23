FROM docker.io/python:3-slim

COPY . /rofl-8004

RUN apt update && apt install -y git python3-dev gcc libc-dev && apt clean && rm -rf /var/lib/apt/lists/*
RUN cd /rofl-8004 && pip install -r requirements.txt

WORKDIR /root
CMD ["python", "/rofl-8004/main.py"]
