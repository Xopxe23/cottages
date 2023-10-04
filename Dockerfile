FROM python:latest

WORKDIR /cottages

COPY . .

RUN pip install -r requirements.txt

RUN chmod a+x /cottages/docker/*.sh
