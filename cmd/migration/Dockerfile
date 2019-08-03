FROM mysql
COPY cmd/migration/migration.sh .
COPY fixtures/init.sql .
RUN chmod 777 /migration.sh /init.sql
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git
RUN git clone https://github.com/vishnubob/wait-for-it.git
