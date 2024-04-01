FROM python:3

RUN apt-get update
RUN apt-get install libmcrypt-dev -y # TODO: revisar si se necesita

RUN pip install awscli
WORKDIR /
COPY entrypoint.sh entrypoint.sh
COPY main main
COPY conf/app.conf conf/app.conf
COPY static static
RUN chmod +x main entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]